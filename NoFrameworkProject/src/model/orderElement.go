package model

type OrderElement struct {
	Id           int `json:"id"`
	OrderId      int `json:"order_id"`
	DishId       int `json:"dish_id"`
	DishQuantity int `json:"dish_quantity"`
}
