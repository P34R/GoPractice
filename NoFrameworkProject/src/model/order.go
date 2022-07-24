package model

type status string

const (
	canceled      status = "canceled"
	in_processing        = "in processing"
	paid                 = "paid"
	done                 = "done"
)

type Order struct {
	Id     int    `json:"id"`
	Status status `json:"status"`
	Total  int    `json:"total"` //total cost
}
