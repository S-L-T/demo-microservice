package repository

import (
	"github.com/S-L-T/demo-microservice/domain/entity"
)

type UserRepository interface {
	AddUser(user entity.User) (string, error)
	UpdateUser(user entity.User) error
	DeleteUser(id string) error
	GetPaginatedUsers(filter entity.Filter, pageNum uint64, pageSize uint64) ([]entity.User, error)
}
