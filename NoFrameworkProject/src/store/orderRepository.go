package store

import (
	"GoTask1/src/model"
	"database/sql"
	"fmt"
)

/*
Запросы добавления заказа,
добавления блюда в заказ,
удаления блюда из заказа,
изменения статуса заказа.
*/

type OrderRepository struct {
	Store *Store
}

func (r *OrderRepository) Create() (*model.Order, error) {
	o := &model.Order{
		Status: "in processing",
		Total:  0,
	}
	err := r.Store.db.QueryRow("INSERT INTO \"order\" (\"status\",\"total\") VALUES ($1,$2) RETURNING \"id\"", o.Status, o.Total).Scan(&o.Id)
	if err != nil {
		return nil, err
	}
	return o, nil
}
func (r *OrderRepository) GetOrder(id int) (*model.Order, error) {
	var o model.Order
	err := r.Store.db.QueryRow("SELECT * FROM \"order\" WHERE \"id\"=$1", id).Scan(&o.Id, &o.Status, &o.Total)
	fmt.Println(&o)
	if err != nil {
		return nil, err
	}
	return &o, nil
}
func (r *OrderRepository) GetOrderElement(id int) (*model.OrderElement, error) {
	var e model.OrderElement
	err := r.Store.db.QueryRow("SELECT * FROM \"order_element\" WHERE \"id\"=$1", id).Scan(&e.Id, &e.OrderId, &e.DishId, &e.DishQuantity)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *OrderRepository) AddElement(e *model.OrderElement) error {
	err := r.Store.db.QueryRow("INSERT INTO \"order_element\" (\"order_id\",\"dish_id\",\"dish_quantity\") VALUES ($1,$2,$3) RETURNING \"id\"", e.OrderId, e.DishId, e.DishQuantity).Scan(&e.Id)
	if err != nil {
		return err
	}
	o, err := r.GetOrder(e.OrderId)
	fmt.Println(err)
	fmt.Println("order: ", o)
	fmt.Println(e)
	r.updateCost(o)
	return nil
}
func (r *OrderRepository) changeNull() {
	r.Store.db.Exec("UPDATE \"order\" set \"total\"=0 WHERE \"id\"=(SELECT \"id\" from \"order\" WHERE \"total\" IS NULL)")
}
func (r *OrderRepository) DeleteElement(id int) error {
	e, err := r.GetOrderElement(id)
	if err != nil {
		return err
	}
	o, _ := r.GetOrder(e.OrderId)
	_, err = r.Store.db.Exec("DELETE FROM \"order_element\" WHERE \"id\"=$1", id)
	if err != nil {
		return err
	}
	r.updateCost(o)
	return nil
}
func (r *OrderRepository) ChangeStatus(o *model.Order) error {
	_, err := r.Store.db.Exec("UPDATE \"order\" SET \"status\"=$1 WHERE \"id\"=$2", o.Status, o.Id)
	if err != nil {
		return err
	}
	return nil
}
func (r *OrderRepository) updateCost(o *model.Order) {
	fmt.Println(o)
	if o.Status != "in processing" {
		return
	}
	r.Store.db.QueryRow("UPDATE \"order\" SET \"total\"= GREATEST(0,(SELECT sum(\"dish\".\"price\"*\"order_element\".\"dish_quantity\") from \"dish\", \"order_element\" WHERE \"order_element\".\"dish_id\"=\"dish\".\"id\" AND \"order_id\"=$1)) WHERE \"id\"=$1 ", o.Id)
}

func (r *OrderRepository) findDishUsage(id int) {
	rows, _ := r.Store.db.Query("SELECT \"order_id\" FROM \"order_element\" WHERE \"dish_id\"=$1", id)
	for rows.Next() {
		var id int
		_ = rows.Scan(&id)
		if s, err := r.GetOrder(id); err != sql.ErrNoRows {
			r.updateCost(s)
		}
	}
}
