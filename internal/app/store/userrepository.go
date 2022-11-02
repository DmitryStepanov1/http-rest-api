package store

import "github.com/DmitryStepanov1/http-rest-api.git/internal/app/model"

// UserRepository ...
type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	return nil, nil
}

func (r *UserRepository) FindByEmail(email string) (model.User, error) {
	return nil, nil
}
