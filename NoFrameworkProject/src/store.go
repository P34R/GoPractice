package src

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
}

/*
"host=localhost dbname=projectdb port=5432 user=docker_user password=admin sslmode=disable",
"postgres://docker_user:admin@localhost/projectdb?sslmode=disable"

host.docker.internal

"user=pqgotest dbname=pqgotest sslmode=verify-full"
"postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
*/
func NewStore() *Store {
	store := &Store{
		url: "host=postgres dbname=projectdb port=5432 user=docker_user password=admin sslmode=disable",
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
func (s *Store) Open() {
	db, err := sql.Open("postgres", s.url)
	if err != nil {
		fmt.Print("Open error ")
		panic(err)
	}
	if err := db.Ping(); err != nil {
		fmt.Print("Ping error ")
		panic(err)
	}
	s.db = db
	s.Init()
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
