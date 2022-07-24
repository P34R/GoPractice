package handlers

import (
	"GoTask1/src/model"
	"GoTask1/src/parsers"
	"GoTask1/src/store"
	"encoding/json"
	"log"
	"net/http"
)

type Dish struct {
	l *log.Logger
	s *store.Store
}

func NewDishHandler(l *log.Logger, s *store.Store) *Dish {
	return &Dish{l, s}
}
func (h *Dish) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		if !HandleBasicAuth(r, h.s, 1) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.GetDish(rw, r)

	case http.MethodPost:
		if !HandleBasicAuth(r, h.s, 3) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.UpdateDish(rw, r)

	case http.MethodDelete:
		if !HandleBasicAuth(r, h.s, 3) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.DeleteDish(rw, r)

	default:
		h.l.Println(http.StatusNotFound)
	}

}

func (h *Dish) UpdateDish(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
	}
	a := model.Dish{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	a.Id = ints[0]
	if err := h.s.Dish().Update(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}
func (h *Dish) CreateDish(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	if !HandleBasicAuth(r, h.s, 3) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("username or password are incorrect"))
		return
	}
	a := model.Dish{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.s.Dish().Create(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}
func (h *Dish) GetDish(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
		return
	}
	mod, err := h.s.Dish().Read(ints[0])
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
func (h *Dish) DeleteDish(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
		return
	}
	if err := h.s.Dish().Delete(ints[0]); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}
