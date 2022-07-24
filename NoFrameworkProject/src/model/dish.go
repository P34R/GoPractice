package model

type Dish struct {
	Id       int `json:"id"`
	Calories int `json:"calories"`
	Category int `json:"category"`
	Price    int `json:"price"`
}
