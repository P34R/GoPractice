package main

import (
	"GoTask1/src"
	"GoTask1/src/model"
	"fmt"
)

func main() {
	store := src.NewStore()
	store.Open()
	store.Init()
	categoryToAdd := &model.Category{
		Name: "{Meat}",
	}
	dishToAdd := &model.Dish{
		Calories: 500,
		Category: 1,
		Price:    200,
	}
	u, err := store.Category().Create(categoryToAdd)
	fmt.Println(u)
	fmt.Println(err)
	store.Dish().Create(dishToAdd)
	var dishmodel *model.Dish
	dishmodel, _ = store.Dish().Read(1)
	fmt.Println(dishmodel)
	dishmodel.Calories = 400
	store.Dish().Update(dishmodel)
	store.Close()
}
