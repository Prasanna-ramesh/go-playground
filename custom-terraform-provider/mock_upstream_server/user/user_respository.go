package user

import (
	"errors"
	"fmt"
	"slices"
)

type User struct {
	id   uint
	name string
	age  uint8
}

var users []User

func save(user User) User {
	users = append(users, user)
	return user
}

func findUserById(id uint) *User {
	index := slices.IndexFunc(users, func(user User) bool {
		return user.id == id
	})
	if index == -1 {
		return nil
	}

	return &users[index]
}

func deleteUserById(id uint) error {
	index := slices.IndexFunc(users, func(user User) bool {
		return user.id == id
	})
	if index == -1 {
		return errors.New(fmt.Sprintf("invalid Id %d", id))
	}

	users = append(users[:index], users[index+1:]...)
	return nil
}

func updateUser(user User) *User {
	index := slices.IndexFunc(users, func(existingUser User) bool {
		return existingUser.id == user.id
	})
	if index == -1 {
		return nil
	}

	users[index] = user
	return &users[index]
}

func getAllUsers() []User {
	return users
}
