package userstore

import (
	"migrationtest/db"
	"migrationtest/models"
)

type UserStorer interface {
	InsertUser(*models.User) error
	GetUsers() ([]*models.User, error)
}

type userStore struct {
	db.Database
}

func NewUserStore(db db.Database) *userStore {
	return &userStore{
		db,
	}
}

func (u *userStore) InsertUser(user *models.User) error {
	stmt, err := u.DB.Prepare("INSERT INTO USERS(USERNAME,PASSWORD,EMAIL) VALUES ($1,$2,$3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Username, user.Password, user.Email)
	return nil
}

func (u *userStore) GetUsers() ([]*models.User, error) {
	rows, err := u.DB.Query("SELECT * FROM USERS;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.User_ID, &user.Username, &user.Password, &user.Email, &user.Mood); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
