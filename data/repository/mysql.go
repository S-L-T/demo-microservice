package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/S-L-T/demo-microservice/domain/entity"
	presentation_http "github.com/S-L-T/demo-microservice/presentation/http"
	"time"
)

func NewMySQLUserRepository() (MySQLUserRepository, error) {
	db, err := sql.Open("mysql", "demo:demo@tcp(db:3306)/users?parseTime=true")
	if err != nil {
		return MySQLUserRepository{}, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return MySQLUserRepository{
		db: db,
	}, nil
}

type MySQLUserRepository struct {
	db *sql.DB
}

func (r *MySQLUserRepository) AddUser(u entity.User) (string, error) {
	stmt, err := r.db.Prepare(
		"INSERT INTO users(first_name,last_name,nickname,password,email,country) VALUES (?,?,?,?,?,?);")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	res, err := stmt.Query(u.FirstName,
		u.LastName,
		u.Nickname,
		u.Password,
		u.Email,
		u.Country,
	)
	if err != nil {
		return "", err
	}

	var iu presentation_http.User
	if res.Next() {
		err = res.Scan(&iu.ID, &iu.FirstName, &iu.LastName, &iu.Nickname, &iu.Password, &iu.Email, &iu.Country, &iu.CreatedAt, &iu.UpdatedAt)
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%x", iu.ID), nil
}

func (r *MySQLUserRepository) UpdateUser(u entity.User) error {
	stmt, err := r.db.Prepare(
		"UPDATE users SET first_name=?, last_name=?, nickname=?, password=?, email=?, country=? WHERE id=UUID_TO_BIN(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.FirstName,
		u.LastName,
		u.Nickname,
		u.Password,
		u.Email,
		u.Country,
		u.ID,
	)
	if err != nil {
		return err
	}

	ra, err := res.RowsAffected()
	if ra != 1 {
		return errors.New("No rows affected")
	}

	return nil
}

func (r *MySQLUserRepository) DeleteUser(id string) error {
	stmt, err := r.db.Prepare("DELETE FROM users WHERE id=UUID_TO_BIN(?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	ra, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if ra == 0 {
		return errors.New("No rows affected")
	}

	return nil
}

func (r *MySQLUserRepository) GetPaginatedUsers(filter entity.Filter, pageNum uint64, pageSize uint64) ([]entity.User, error) {
	stmt, err := r.db.Prepare("SELECT * FROM users LIMIT ?, ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Query((pageNum-1)*pageSize, pageSize)
	var users []entity.User
	for res.Next() {
		var u presentation_http.User
		err = res.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Nickname, &u.Password, &u.Email, &u.Country, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		tc, err	:= time.Parse(time.RFC3339, u.CreatedAt)
		if err != nil {
			return nil, err
		}

		tu, err := time.Parse(time.RFC3339, u.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, entity.User{
			ID:        fmt.Sprintf("%x", u.ID),
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Nickname:  u.Nickname,
			Password:  u.Password,
			Email:     u.Email,
			Country:   u.Country,
			CreatedAt: tc,
			UpdatedAt: tu,
		})
	}

	return users, nil
}
