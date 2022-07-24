package handlers

import (
	"GoTask1/src/parsers"
	"GoTask1/src/store"
	"encoding/json"
	"log"
	"net/http"
)

type OrderElement struct {
	l *log.Logger
	s *store.Store
}

func NewOrderElementHandler(l *log.Logger, s *store.Store) *OrderElement {
	return &OrderElement{l, s}
}
func (h *OrderElement) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if !HandleBasicAuth(r, h.s, 1) {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("username or password are incorrect"))
	}
	switch r.Method {

	case http.MethodGet:
		h.GetElement(rw, r)

	case http.MethodDelete:
		h.DeleteElement(rw, r)

	default:
		h.l.Println(http.StatusNotFound)
	}

}

func (h *OrderElement) GetElement(rw http.ResponseWriter, r *http.Request) {
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
func (h *OrderElement) DeleteElement(rw http.ResponseWriter, r *http.Request) {
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
