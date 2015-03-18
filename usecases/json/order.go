package json

import (
	"log"

	"github.com/crowdint/gopher-spree-api/cache"
	"github.com/crowdint/gopher-spree-api/configs/spree"
	"github.com/crowdint/gopher-spree-api/domain"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	"github.com/crowdint/gopher-spree-api/utils"
)

type OrderInteractor struct {
	AssetInteractor       *AssetInteractor
	AdjustmentRepository  *repositories.AdjustmentRepository
	OrderRepository       *repositories.OrderRepository
	OptionValueRepository *repositories.OptionValueRepository
	ShipmentRepository    *repositories.ShipmentRepository
}

func NewOrderInteractor() *OrderInteractor {
	return &OrderInteractor{
		AssetInteractor:       NewAssetInteractor(),
		AdjustmentRepository:  repositories.NewAdjustmentRepository(),
		OrderRepository:       repositories.NewOrderRepository(),
		OptionValueRepository: repositories.NewOptionValueRepo(),
		ShipmentRepository:    repositories.NewShipmentRepository(),
	}
}

func (this *OrderInteractor) Show(order *domain.Order, u *domain.User) (*domain.Order, error) {
	if err := cache.Find(order); err != nil {
		utils.LogrusError("Show", err)

		this.setComputedValues(order, u)

		variantsMap, productsMap, pricesMap, stockItemsMap := this.getAssociationMaps(order)

		for i, lineItem := range *order.LineItems {
			variant := variantsMap[lineItem.VariantId].(domain.Variant)
			product := productsMap[variant.ProductId].(domain.Product)
			price := pricesMap[variant.Id].(domain.Price)

			variant.Name = product.Name
			variant.Description = product.Description
			variant.Slug = product.Slug
			variant.Price = &price.Amount

			for _, stockItem := range stockItemsMap[variant.Id].([]interface{}) {
				si := stockItem.(domain.StockItem)
				variant.StockItems = append(variant.StockItems, &si)
			}

			variant.SetComputedValues()
			variant.Images = this.getVariantImages(variant.Id)
			variant.OptionValues = this.OptionValueRepository.AllByVariantAssociation(&variant)

			(*order.LineItems)[i].Variant = &variant
			(*order.LineItems)[i].Adjustments = this.AdjustmentRepository.AllByAdjustable(lineItem)
		}

		this.setPayments(order)
		this.setShipments(order)

		order.Adjustments = this.AdjustmentRepository.AllByAdjustable(order)

		if err := cache.Set(order); err != nil {
			utils.LogrusError("Show", err)
			log.Println("An error occurred while setting the cache: ", err.Error())
		}
	}

	return order, nil
}

func (this *OrderInteractor) GetResponse(currentPage, perPage int, params ResponseParameters) (ContentResponse, error) {
	orders := []*domain.Order{}

	queryData, err := params.GetQuery()
	if err != nil {
		return &OrderResponse{}, err
	}

	query := queryData.Query
	gparams := queryData.Params

	err = this.OrderRepository.All(&orders, map[string]interface{}{"limit": perPage, "offset": currentPage}, query, gparams)
	if err != nil {
		utils.LogrusError("GetResponse", err)

		return &OrderResponse{}, err
	}

	ordersCached := this.toCacheData(orders)
	missingOrdersCached, _ := cache.FetchMultiWithPrefix("index", ordersCached)
	if len(missingOrdersCached) == 0 {
		return OrderResponse{data: orders}, nil
	}

	var orderIds []int64
	for _, order := range orders {
		orderIds = append(orderIds, order.Id)
	}

	quantities, err := this.OrderRepository.SumLineItemsQuantityByOrderIds(orderIds)
	for _, order := range orders {
		order.Quantity = quantities[order.Id]
		order.SetComputedValues()
		cache.SetWithPrefix("index", order)
	}

	return &OrderResponse{data: orders}, nil
}

func (this *OrderInteractor) GetShowResponse(params ResponseParameters) (interface{}, error) {
	return nil, nil
}

func (this *OrderInteractor) GetCreateResponse(params ResponseParameters) (interface{}, interface{}, error) {
	return nil, nil, nil
}

func (this *OrderInteractor) GetTotalCount(params ResponseParameters) (int64, error) {
	queryData, err := params.GetQuery()
	if err != nil {
		utils.LogrusError("GetTotalCount", err)

		return 0, err
	}

	query := queryData.Query
	gparams := queryData.Params

	return this.OrderRepository.Count(domain.Order{}, query, gparams)
}

func (this *OrderInteractor) toCacheData(orderSlice []*domain.Order) (ordersCached []cache.Cacheable) {
	for _, order := range orderSlice {
		ordersCached = append(ordersCached, order)
	}
	return
}

func (this *OrderInteractor) setPayments(order *domain.Order) {
	payments := []*domain.Payment{}
	this.OrderRepository.All(&payments, map[string]interface{}{
		"order": "created_at",
	}, "order_id = ?", order.Id)

	for _, payment := range payments {
		payment.Order = order
		payment.SetComputedValues()
	}

	order.Payments = payments
}

func (this *OrderInteractor) setShipments(order *domain.Order) {
	order.Shipments = this.ShipmentRepository.AllByOrder(order)
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

	this.OrderRepository.Association(order, address, id)

	if address.Id != 0 {
		address.Country = &domain.Country{}
		this.OrderRepository.Association(address, address.Country, "CountryId")

		address.State = &domain.State{}
		this.OrderRepository.Association(address, address.State, "StateId")
		address.StateName = address.State.Name
		address.StateText = address.State.Abbr
	} else {
		address = nil
	}

	return address
}

func (this *OrderInteractor) getAssociationMaps(order *domain.Order) (varm, prom, prim, stim map[int64]interface{}) {
	variantIds := utils.Collect(*order.LineItems, "VariantId")
	var variants []domain.Variant
	this.OrderRepository.All(&variants, nil, "id IN(?)", variantIds)
	varm = utils.ToMap(variants, "Id", false)

	productIds := utils.Collect(variants, "ProductId")
	var products []domain.Product
	this.OrderRepository.All(&products, nil, "id IN(?)", productIds)
	prom = utils.ToMap(products, "Id", false)

	var prices []domain.Price
	this.OrderRepository.All(&prices, nil, "currency = ? AND variant_id IN(?)", spree.Get(spree.CURRENCY), variantIds)
	prim = utils.ToMap(prices, "VariantId", false)

	var stockLocations []domain.StockLocation
	this.OrderRepository.All(&stockLocations, nil, map[string]interface{}{"active": true})
	stockLocationIds := utils.Collect(stockLocations, "Id")

	var stockItems []domain.StockItem
	this.OrderRepository.All(&stockItems, nil, "variant_id IN(?) AND stock_location_id IN(?)", variantIds, stockLocationIds)
	stim = utils.ToMap(stockItems, "VariantId", true)

	return
}

func (this *OrderInteractor) getLineItems(order *domain.Order) *[]domain.LineItem {
	lineItems := &[]domain.LineItem{}
	this.OrderRepository.Association(order, lineItems, "OrderId")
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
	data []*domain.Order
}

func (this OrderResponse) GetCount() int {
	return len(this.data)
}

func (this OrderResponse) GetData() interface{} {
	return this.data
}

func (this OrderResponse) GetTag() string {
	return "orders"
}
