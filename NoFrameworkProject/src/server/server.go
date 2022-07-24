package server

import (
	"GoTask1/src/server/handlers"
	"GoTask1/src/store"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	store    *store.Store
	ServeMux *http.ServeMux
	Server   *http.Server
}

func newServer() *Server {
	ServeMux := http.NewServeMux()
	s := store.NewStore()
	return &Server{
		store:    s,
		ServeMux: ServeMux,
		Server: &http.Server{
			Addr:         ":9091",
			Handler:      ServeMux,
			IdleTimeout:  120 * time.Second,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}
}
func (s *Server) configure() {
	l := log.New(os.Stdout, "project-api ", log.LstdFlags)
	// /category/id
	// /category
	// /dish/id
	// /dish
	// /order/id/dish
	// /order
	// /order/id
	ch := handlers.NewCategoryHandler(l, s.store)
	dh := handlers.NewDishHandler(l, s.store)
	oh := handlers.NewOrderHandler(l, s.store)
	oeh := handlers.NewOrderElementHandler(l, s.store)
	uh := handlers.NewUserHandler(l, s.store)
	s.ServeMux.Handle("/category/", ch)
	s.ServeMux.Handle("/dish/", dh)
	s.ServeMux.Handle("/order/", oh)
	s.ServeMux.Handle("/order_element/", oeh)
	s.ServeMux.Handle("/user", uh)
	s.ServeMux.HandleFunc("/order", oh.CreateOrder)
	s.ServeMux.HandleFunc("/category", ch.CreateCategory)
	s.ServeMux.HandleFunc("/dish", dh.CreateDish)

}
func Start() {
	s := newServer()
	s.configure()
	if err := s.store.Open(); err != nil {
		panic(err)
	}
	defer s.store.Close()
	if err := s.Server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// Not used currently
func ServerAnswerJSON(rw http.ResponseWriter, name string, text string) error {
	type req struct {
		name string
		text string
	}
	r := &req{
		name: name,
		text: text,
	}
	encod := json.NewEncoder(rw)
	if err := encod.Encode(r); err != nil {
		return err
	}
	return nil
}
