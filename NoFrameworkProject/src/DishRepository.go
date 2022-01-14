package src

import (
	"GoTask1/src/model"
)

type DishRepository struct {
	Store *Store
}

func (r *DishRepository) Create(u *model.Dish) (*model.Dish, error) {
	err := r.Store.db.QueryRow("INSERT INTO \"dish\" (\"calories\", \"category\", \"price\") VALUES ($1,$2,$3) RETURNING \"id\"", u.Calories, u.Category, u.Price).Scan(&u.Id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//Update method requires Dish model to have dish id, what will be updated
func (r *DishRepository) Update(u *model.Dish) (*model.Dish, error) {
	_, err := r.Store.db.Exec("UPDATE \"dish\" SET \"calories\"=$1,\"category\"=$2,\"price\"=$3 WHERE \"id\"=$4", u.Calories, u.Category, u.Price, u.Id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (r *DishRepository) Read(id int) (*model.Dish, error) {
	var u model.Dish
	row := r.Store.db.QueryRow("SELECT * FROM \"dish\" WHERE \"id\"=$1", id)
	err := row.Scan(&u.Id, &u.Calories, &u.Category, &u.Price)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *DishRepository) Delete(id int) error {
	_, err := r.Store.db.Exec("DELETE FROM \"dish\" WHERE \"id\"=$1", id)
	if err != nil {
		panic(err)
	}
	return nil
}
