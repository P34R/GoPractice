package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store struct {
	url      string
	db       *sql.DB
	dish     *DishRepository
	category *CategoryRepository
	order    *OrderRepository
	user     *UserRepository
}

func NewStore() *Store {

	store := &Store{
		url: "host=localhost dbname=projectdb port=5432 user=postgres password=admin sslmode=disable",
	}

	return store
}
func (s *Store) Init() {
	_, err := s.db.Exec("DROP TABLE IF EXISTS \"dish\"")
	if err != nil {
		panic(err)
	}
	_, err = s.db.Exec("DROP TABLE IF EXISTS \"category\"")
	if err != nil {
		panic(err)
	}
	_, err = s.db.Exec("CREATE TABLE \"category\"(id SERIAL NOT NULL PRIMARY KEY, name varchar[50] NOT NULL)")
	if err != nil {
		panic(err)
	}

	_, err = s.db.Exec("CREATE TABLE \"dish\"(id SERIAL NOT NULL PRIMARY KEY, calories int NOT NULL, category int not null, price int not null, FOREIGN KEY (category)  REFERENCES \"category\" (id))")
	if err != nil {
		panic(err)
	}

}
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.url)
	if err != nil {
		fmt.Print("Open error ")
		return err
	}
	if err := db.Ping(); err != nil {
		fmt.Print("Ping error ")
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
func (s *Store) Dish() *DishRepository {
	if s.dish == nil {
		s.dish = &DishRepository{
			Store: s,
		}
	}
	return s.dish
}

func (s *Store) Category() *CategoryRepository {
	if s.category == nil {
		s.category = &CategoryRepository{
			Store: s,
		}
	}
	return s.category
}
func (s *Store) Order() *OrderRepository {
	if s.order == nil {
		s.order = &OrderRepository{
			Store: s,
		}
	}
	return s.order
}
func (s *Store) User() *UserRepository {
	if s.user == nil {
		s.user = &UserRepository{
			Store: s,
		}
	}
	return s.user
}
