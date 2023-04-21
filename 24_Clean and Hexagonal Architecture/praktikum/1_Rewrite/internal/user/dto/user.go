package dto

import "rewrite/pkg/entity"

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

func (u *User) FromEntity(entity *entity.User) {
	u.ID = entity.ID
	u.Email = entity.Email
	u.Password = entity.Password
}

func (u *Users) ToEntity() entity.Users {
	var users entity.Users
	for _, user := range *u {
		users = append(users, *user.ToEntity())
	}
	return users
}

func (u *Users) FromEntity(entities entity.Users) {
	for _, entity := range entities {
		var user User
		user.FromEntity(&entity)
		*u = append(*u, user)
	}
}
