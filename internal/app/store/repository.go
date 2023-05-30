package store

import "github.com/DmitryStepanov1/http-rest-api.git/internal/app/model"

// UserRepository repository
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
