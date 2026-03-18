package use_case

import (
	"github.com/S-L-T/demo-microservice/domain/entity"
	"github.com/S-L-T/demo-microservice/domain/repository"
)

func NewUserUseCase(r repository.UserRepository) User {
	return User{
		r: r,
	}
}

type User struct {
	r repository.UserRepository
}

func (u *User) AddUser(ue entity.User) (string, error) {
	return u.r.AddUser(ue)
}

func (u *User) UpdateUser(ue entity.User) error {
	return u.r.UpdateUser(ue)
}

func (u *User) DeleteUser(id string) error {
	return u.r.DeleteUser(id)
}

func (u *User) GetPaginatedUsers(f entity.Filter, pageNum uint64, pageSize uint64) ([]entity.User, error) {
	return u.r.GetPaginatedUsers(f, pageNum, pageSize)
}
