package handlers

import (
	"GoTask1/src/model"
	"GoTask1/src/parsers"
	"GoTask1/src/store"
	"encoding/json"
	"log"
	"net/http"
)

type Order struct {
	l *log.Logger
	s *store.Store
}

func NewOrderHandler(l *log.Logger, s *store.Store) *Order {
	return &Order{l, s}
}
func (h *Order) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		if !HandleBasicAuth(r, h.s, 1) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.GetOrder(rw, r)

	case http.MethodPost:
		if !HandleBasicAuth(r, h.s, 2) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.AddElement(rw, r)

	case http.MethodPatch:
		if !HandleBasicAuth(r, h.s, 1) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.ChangeStatus(rw, r)

	default:
		h.l.Println(http.StatusNotFound)
	}

}

func (h *Order) CreateOrder(rw http.ResponseWriter, r *http.Request) {
	if !HandleBasicAuth(r, h.s, 1) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("username or password are incorrect"))
		return
	}
	a, err := h.s.Order().Create()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
	encoder := json.NewEncoder(rw)
	if err := encoder.Encode(a); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}
func (h *Order) GetOrder(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
		return
	}
	mod, err := h.s.Order().GetOrder(ints[0])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(rw)
	if err = encoder.Encode(&mod); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}
func (h *Order) GetElement(rw http.ResponseWriter, r *http.Request) {
	if !HandleBasicAuth(r, h.s, 1) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("username or password are incorrect"))
		return
	}
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
		return
	}
	mod, err := h.s.Order().GetOrderElement(ints[0])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(rw)
	if err = encoder.Encode(&mod); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}
func (h *Order) AddElement(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
		return
	}
	type req struct {
		DishId       int `json:"dish_id"`
		DishQuantity int `json:"dish_quantity"`
	}
	var a req
	oe := model.OrderElement{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	oe.OrderId = ints[0]
	oe.DishQuantity = a.DishQuantity
	oe.DishId = a.DishId
	if err := h.s.Order().AddElement(&oe); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
	if err := json.NewEncoder(rw).Encode(oe); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}
func (h *Order) DeleteElement(rw http.ResponseWriter, r *http.Request) {
	if !HandleBasicAuth(r, h.s, 2) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("username or password are incorrect"))
		return
	}
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
		return
	}
	if err := h.s.Order().DeleteElement(ints[0]); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}
func (h *Order) ChangeStatus(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
		return
	}
	a := model.Order{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	a.Id = ints[0]
	if err := h.s.Order().ChangeStatus(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}
