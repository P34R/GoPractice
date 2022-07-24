package store

import "GoTask1/src/model"

type CategoryRepository struct {
	Store *Store
}

func (r *CategoryRepository) Create(u *model.Category) error {
	err := r.Store.db.QueryRow("INSERT INTO \"category\" (\"name\") VALUES ($1) RETURNING \"id\"", u.Name).Scan(&u.Id)
	if err != nil {
		return err
	}
	return nil
}

//Update method requires Category model to have category id, what will be updated
func (r *CategoryRepository) Update(u *model.Category) error {
	_, err := r.Store.db.Exec("UPDATE \"category\" SET \"name\"=$1 WHERE \"id\"=$2", u.Name, u.Id)
	if err != nil {
		return err
	}
	return nil
}
func (r *CategoryRepository) Read(id int) (*model.Category, error) {
	var u model.Category
	err := r.Store.db.QueryRow("SELECT * FROM \"category\" WHERE \"id\"=($1)", id).Scan(&u.Id, &u.Name)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (r *CategoryRepository) Delete(id int) error {
	_, err := r.Store.db.Exec("DELETE FROM \"category\" WHERE \"id\"=$1", id)
	if err != nil {
		return err
	}
	return nil
}
