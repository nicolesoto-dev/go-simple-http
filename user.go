package main

import "fmt"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"first_name"`
}

var users = []User{}

func InsertUser(u User) error {
	if u.ID == 0 || u.Name == "" {
		return fmt.Errorf("invalid user: %v", u)
	}

	for _, user := range users {
		if user.ID == u.ID {
			return fmt.Errorf("user already exists: %v", u)
		}
	}

	users = append(users, u)

	return nil
}
