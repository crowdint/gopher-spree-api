package repositories

type OrderRepository struct {
	DbRepository
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		DbRepository{Spree_db},
	}
}

func (this *OrderRepository) SumLineItemsQuantityByOrderIds(ids []int64) (map[int64]int64, error) {
	orderQuantities := []struct {
		Id  int64
		Sum int64
	}{}

	err := this.dbHandler.
		Table("spree_line_items").
		Select("order_id AS id, SUM(quantity) AS sum").
		Where("order_id IN (?)", ids).
		Group("order_id").
		Scan(&orderQuantities).
		Error

	orderQuantitiesMap := map[int64]int64{}
	for _, oq := range orderQuantities {
		orderQuantitiesMap[oq.Id] = oq.Sum
	}

	return orderQuantitiesMap, err
}
