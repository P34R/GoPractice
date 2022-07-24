package store

import (
	"GoTask1/src/model"
	"fmt"
)

type DishRepository struct {
	Store *Store
}

func (r *DishRepository) Create(u *model.Dish) error {
	fmt.Println(u)
	err := r.Store.db.QueryRow("INSERT INTO \"dish\" (\"calories\", \"category\", \"price\") VALUES ($1,$2,$3) RETURNING \"id\"", u.Calories, u.Category, u.Price).Scan(&u.Id)
	if err != nil {
		return err
	}
	return nil
}

//Update method requires Dish model to have dish id, what will be updated
func (r *DishRepository) Update(u *model.Dish) error {
	n, err := r.Read(u.Id)
	if err != nil {
		return err
	}
	if u.Category == 0 {
		u.Category = n.Category
	}
	if u.Price == 0 {
		u.Price = n.Price
	}
	if u.Calories == 0 {
		u.Calories = n.Calories
	}
	_, err = r.Store.db.Exec("UPDATE \"dish\" SET \"calories\"=$1,\"category\"=$2,\"price\"=$3 WHERE \"id\"=$4", u.Calories, u.Category, u.Price, u.Id)
	if err != nil {
		return err
	}
	if u.Price != n.Price {
		r.Store.Order().findDishUsage(u.Id)
	}
	return nil
}
func (r *DishRepository) Read(id int) (*model.Dish, error) {
	var u model.Dish
	err := r.Store.db.QueryRow("SELECT * FROM \"dish\" WHERE \"id\"=$1", id).Scan(&u.Id, &u.Calories, &u.Category, &u.Price)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *DishRepository) Delete(id int) error {
	_, err := r.Store.db.Exec("DELETE FROM \"dish\" WHERE \"id\"=$1", id)
	if err != nil {
		return err
	}
	return nil
}
