package store

import "github.com/DmitryStepanov1/http-rest-api/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
