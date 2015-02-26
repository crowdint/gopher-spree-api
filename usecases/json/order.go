package json

import (
	"github.com/jinzhu/copier"

	"github.com/crowdint/gopher-spree-api/configs/spree"
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	. "github.com/crowdint/gopher-spree-api/utils"
)

type OrderInteractor struct {
	AssetInteractor       *AssetInteractor
	AdjustmentRepository  *repositories.AdjustmentRepository
	BaseRepository        *repositories.DbRepository
	OrderRepository       *repositories.OrderRepository
	OptionValueRepository *repositories.OptionValueRepository
	ShipmentRepository    *repositories.ShipmentRepository
}

func (this *OrderInteractor) Show(o *domain.Order, u *domain.User) (*domain.Order, error) {
	order := domain.Order{}
	copier.Copy(&order, o)

	this.setComputedValues(&order, u)

	variantsMap, productsMap, pricesMap, stockItemsMap := this.getAssociationMaps(&order)

	for i, lineItem := range *order.LineItems {
		variant := variantsMap[lineItem.VariantId].(domain.Variant)
		product := productsMap[variant.ProductId].(domain.Product)
		price := pricesMap[variant.Id].(domain.Price)

		variant.Name = product.Name
		variant.Description = product.Description
		variant.Slug = product.Slug
		variant.Price = price.Amount

		for _, stockItem := range stockItemsMap[variant.Id].([]interface{}) {
			si := stockItem.(domain.StockItem)
			variant.StockItems = append(variant.StockItems, &si)
		}

		variant.SetInventoryValues()
		variant.Images = this.getVariantImages(variant.Id)
		variant.OptionValues = this.OptionValueRepository.AllByVariantAssociation(&variant)

		(*order.LineItems)[i].Variant = &variant
		(*order.LineItems)[i].Adjustments = this.AdjustmentRepository.AllByAdjustable(lineItem.Id, lineItem.SpreeClass())
	}

	this.setPayments(&order)
	order.Shipments = this.ShipmentRepository.AllByOrder(&order)
	order.Adjustments = this.AdjustmentRepository.AllByAdjustable(order.Id, order.SpreeClass())

	return &order, nil
}

func (this *OrderInteractor) GetResponse(currentPage, perPage int, params ResponseParameters) (ContentResponse, error) {
	orders := []domain.Order{}
	ordersJson := []domain.Order{}

	query, gparams, err := params.GetGransakParams()

	if err != nil {
		return &OrderResponse{}, err
	}

	err = this.BaseRepository.All(&orders, map[string]interface{}{"limit": perPage, "offset": currentPage}, query, gparams)

	if err != nil {
		return &OrderResponse{}, err
	}

	if len(orders) == 0 {
		return &OrderResponse{data: &ordersJson}, nil
	}

	var orderIds []int64
	for _, order := range orders {
		orderIds = append(orderIds, order.Id)
	}

	quantities, err := this.OrderRepository.SumLineItemsQuantityByOrderIds(orderIds)
	for index, order := range orders {
		orders[index].Quantity = quantities[order.Id]
	}

	copier.Copy(&ordersJson, &orders)

	return &OrderResponse{data: &ordersJson}, nil
}

func (this *OrderInteractor) GetShowResponse(params ResponseParameters) (interface{}, error) {
	return nil, nil
}

func (this *OrderInteractor) GetTotalCount(params ResponseParameters) (int64, error) {
	query, gparams, err := params.GetGransakParams()

	if err != nil {
		return 0, err
	}

	return this.BaseRepository.Count(domain.Order{}, query, gparams)
}

func (this *OrderInteractor) setPayments(order *domain.Order) {
	payments := []domain.Payment{}
	this.BaseRepository.All(&payments, map[string]interface{}{
		"order": "created_at",
	}, "order_id = ?", order.Id)
	order.Payments = payments
}

func (this *OrderInteractor) getVariantImages(variantId int64) []*domain.Asset {
	jsonImages := []*domain.Asset{}

	modelImages, err := this.AssetInteractor.Repository.AllImagesByVariantId(variantId)
	if err == nil {
		jsonImages = this.AssetInteractor.toJsonAssets(modelImages)
	}

	return jsonImages
}

func (this *OrderInteractor) getAddress(order *domain.Order, id string) *domain.Address {
	address := &domain.Address{}

	this.BaseRepository.Association(order, address, id)

	if address.Id != 0 {
		address.Country = &domain.Country{}
		this.BaseRepository.Association(address, address.Country, "CountryId")

		address.State = &domain.State{}
		this.BaseRepository.Association(address, address.State, "StateId")
		address.StateName = address.State.Name
		address.StateText = address.State.Abbr
	} else {
		address = nil
	}

	return address
}

func (this *OrderInteractor) getAssociationMaps(order *domain.Order) (varm, prom, prim, stim map[int64]interface{}) {
	variantIds := Collect(*order.LineItems, "VariantId")
	var variants []domain.Variant
	this.BaseRepository.All(&variants, nil, "id IN(?)", variantIds)
	varm = ToMap(variants, "Id", false)

	productIds := Collect(variants, "ProductId")
	var products []domain.Product
	this.BaseRepository.All(&products, nil, "id IN(?)", productIds)
	prom = ToMap(products, "Id", false)

	var prices []domain.Price
	this.BaseRepository.All(&prices, nil, "currency = ? AND variant_id IN(?)", spree.Get(spree.CURRENCY), variantIds)
	prim = ToMap(prices, "VariantId", false)

	var stockLocations []domain.StockLocation
	this.BaseRepository.All(&stockLocations, nil, map[string]interface{}{"active": true})
	stockLocationIds := Collect(stockLocations, "Id")

	var stockItems []domain.StockItem
	this.BaseRepository.All(&stockItems, nil, "variant_id IN(?) AND stock_location_id IN(?)", variantIds, stockLocationIds)
	stim = ToMap(stockItems, "VariantId", true)

	return
}

func (this *OrderInteractor) getLineItems(order *domain.Order) *[]domain.LineItem {
	lineItems := &[]domain.LineItem{}
	this.BaseRepository.Association(order, lineItems, "OrderId")
	return lineItems
}

func (this *OrderInteractor) getPermissions(order *domain.Order, user *domain.User) *domain.Permissions {
	updatePermission := user.HasRole("admin") || (*order.UserId == user.Id)
	permissions := &domain.Permissions{CanUpdate: &updatePermission}
	return permissions
}

func (this *OrderInteractor) getQuantity(order *domain.Order) int64 {
	quantities, _ := this.OrderRepository.SumLineItemsQuantityByOrderIds([]int64{order.Id})
	return quantities[order.Id]
}

func (this *OrderInteractor) setComputedValues(order *domain.Order, user *domain.User) {
	order.Permissions = this.getPermissions(order, user)
	order.Quantity = this.getQuantity(order)
	order.BillAddress = this.getAddress(order, "BillAddressId")
	order.ShipAddress = this.getAddress(order, "ShipAddressId")
	order.LineItems = this.getLineItems(order)
}

type OrderResponse struct {
	data *[]domain.Order
}

func (this *OrderResponse) GetCount() int {
	return len(*this.data)
}

func (this *OrderResponse) GetData() interface{} {
	return this.data
}

func (this OrderResponse) GetTag() string {
	return "orders"
}

func NewOrderInteractor() *OrderInteractor {
	return &OrderInteractor{
		AssetInteractor:       NewAssetInteractor(),
		AdjustmentRepository:  repositories.NewAdjustmentRepository(),
		BaseRepository:        repositories.NewDatabaseRepository(),
		OrderRepository:       repositories.NewOrderRepository(),
		OptionValueRepository: repositories.NewOptionValueRepo(),
		ShipmentRepository:    repositories.NewShipmentRepository(),
	}
}
