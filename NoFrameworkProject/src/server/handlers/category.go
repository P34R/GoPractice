package handlers

import (
	"GoTask1/src/model"
	"GoTask1/src/parsers"
	"GoTask1/src/store"
	"encoding/json"
	"log"
	"net/http"
)

type Category struct {
	l *log.Logger
	s *store.Store
}

func NewCategoryHandler(l *log.Logger, s *store.Store) *Category {
	return &Category{l, s}
}
func (h *Category) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		if !HandleBasicAuth(r, h.s, 1) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.GetCategory(rw, r)

	case http.MethodPost:
		if !HandleBasicAuth(r, h.s, 3) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.UpdateCategory(rw, r)

	case http.MethodDelete:
		if !HandleBasicAuth(r, h.s, 3) {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("username or password are incorrect"))
			return
		}
		h.DeleteCategory(rw, r)

	default:
		h.l.Println(http.StatusNotFound)
	}

}

func (h *Category) UpdateCategory(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
	}
	a := model.Category{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	a.SetName(a.Name)
	a.Id = ints[0]
	if err := h.s.Category().Update(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}
func (h *Category) CreateCategory(rw http.ResponseWriter, r *http.Request) {
	if !HandleBasicAuth(r, h.s, 3) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("username or password are incorrect"))
		return
	}
	a := model.Category{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	a.SetName(a.Name)
	if err := h.s.Category().Create(&a); err != nil {
		http.Error(rw, err.Error(), http.StatusConflict)
		return
	}
}
func (h *Category) GetCategory(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
		return
	}
	mod, err := h.s.Category().Read(ints[0])
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
func (h *Category) DeleteCategory(rw http.ResponseWriter, r *http.Request) {
	ints := parsers.GetIntsFromURL(r.URL.Path)
	if len(ints) == 0 {
		http.Error(rw, "Id was not provided", http.StatusBadRequest)
		return
	}
	if err := h.s.Category().Delete(ints[0]); err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}
