package json

import (
	"github.com/jinzhu/copier"

	"github.com/crowdint/gopher-spree-api/configs/spree"
	"github.com/crowdint/gopher-spree-api/domain/json"
	"github.com/crowdint/gopher-spree-api/domain/models"
	"github.com/crowdint/gopher-spree-api/interfaces/repositories"
	. "github.com/crowdint/gopher-spree-api/utils"
)

type OrderInteractor struct {
	Repository *repositories.DbRepo
}

func (this *OrderInteractor) Show(o *models.Order, u *models.User) (*json.Order, error) {
	order := json.Order{}
	copier.Copy(&order, o)

	order.Permissions = this.getPermissions(&order, u)
	order.Quantity = this.getQuantity(&order)
	order.BillAddress = this.getAddress(&order, "BillAddressId")
	order.ShipAddress = this.getAddress(&order, "ShipAddressId")
	order.LineItems = this.getLineItems(&order)

	variantIds := Collect(*order.LineItems, "VariantId")
	var variants []json.Variant
	this.Repository.All(&variants, nil, "id IN(?)", variantIds)

	productIds := Collect(variants, "ProductId")
	var products []json.Product
	this.Repository.All(&products, nil, "id IN(?)", productIds)

	var prices []models.Price
	this.Repository.All(&prices, nil, "currency = ? AND variant_id IN(?)", spree.Get(spree.CURRENCY), variantIds)

	var stockLocations []json.StockLocation
	this.Repository.All(&stockLocations, nil, map[string]interface{}{"active": true})
	stockLocationIds := Collect(stockLocations, "Id")

	var stockItems []json.StockItem
	this.Repository.All(&stockItems, nil, "variant_id IN(?) AND stock_location_id IN(?)", variantIds, stockLocationIds)

	variantsMap := ToMap(variants, "Id", false)
	productsMap := ToMap(products, "Id", false)
	pricesMap := ToMap(prices, "VariantId", false)
	stockItemsMap := ToMap(stockItems, "VariantId", true)

	for i, lineItem := range *order.LineItems {
		variant := variantsMap[lineItem.VariantId].(json.Variant)
		product := productsMap[variant.Id].(json.Product)
		price := pricesMap[variant.Id].(models.Price)

		variant.Name = product.Name
		variant.Description = product.Description
		variant.Slug = product.Slug
		variant.Price = price.Amount

		for _, stockItem := range stockItemsMap[variant.Id].([]interface{}) {
			si := stockItem.(json.StockItem)
			variant.StockItems = append(variant.StockItems, &si)
		}

		variant.SetInventoryValues()
		(*order.LineItems)[i].Variant = &variant
	}

	return &order, nil
}

func (this *OrderInteractor) getQuantity(order *json.Order) int64 {
	quantities, _ := this.Repository.SumLineItemsQuantityByOrderIds([]int64{order.Id})
	return quantities[order.Id]
}

func (this *OrderInteractor) getPermissions(order *json.Order, user *models.User) *json.Permissions {
	updatePermission := user.HasRole("admin") || (*order.UserId == user.Id)
	permissions := &json.Permissions{CanUpdate: &updatePermission}
	return permissions
}

func (this *OrderInteractor) getLineItems(order *json.Order) *[]json.LineItem {
	lineItems := &[]json.LineItem{}
	this.Repository.Association(order, lineItems, "OrderId")
	return lineItems
}

func (this *OrderInteractor) getAddress(order *json.Order, id string) *json.Address {
	address := &json.Address{}

	this.Repository.Association(order, address, id)

	address.Country = &json.Country{}
	this.Repository.Association(address, address.Country, "CountryId")

	address.State = &json.State{}
	this.Repository.Association(address, address.State, "StateId")
	address.StateName = address.State.Name
	address.StateText = address.State.Abbr

	return address
}

func (this *OrderInteractor) GetShowResponse(params ResponseParameters) (interface{}, error) {
	return nil, nil
}

func (this *OrderInteractor) GetResponse(currentPage, perPage int, params ResponseParameters) (ContentResponse, error) {
	orders := []models.Order{}
	ordersJson := []json.Order{}

	query, gparams, err := params.GetGransakParams()

	if err != nil {
		return &OrderResponse{}, err
	}

	err = this.Repository.All(&orders, map[string]interface{}{"limit": perPage, "offset": currentPage}, query, gparams)

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

	quantities, err := this.Repository.SumLineItemsQuantityByOrderIds(orderIds)
	for index, order := range orders {
		orders[index].Quantity = quantities[order.Id]
	}

	copier.Copy(&ordersJson, &orders)

	return &OrderResponse{data: &ordersJson}, nil
}

func (this *OrderInteractor) GetTotalCount(params ResponseParameters) (int64, error) {
	query, gparams, err := params.GetGransakParams()

	if err != nil {
		return 0, err
	}

	return this.Repository.Count(models.Order{}, query, gparams)
}

func NewOrderInteractor() *OrderInteractor {
	return &OrderInteractor{
		Repository: repositories.NewDatabaseRepository(),
	}
}

type OrderResponse struct {
	data *[]json.Order
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
