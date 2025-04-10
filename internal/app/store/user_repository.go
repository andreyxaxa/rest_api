package store

import "github.com/andreyxaxa/rest_api/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	Find(id int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
	Delete(id int) error
}
