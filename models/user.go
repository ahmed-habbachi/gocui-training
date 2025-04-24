package models

import "slices"

type User struct {
	Id    int
	Name  string
	Age   int
	Email string
}

var (
	users  []User
	nextID = 1
)

func GetUsers() []User {
	return users
}

func GetUser(id int) *User {
	index := slices.IndexFunc(users, func(u User) bool {
		return u.Id == id
	})
	if index == -1 {
		return nil
	}
	return &users[index]
}

func AddUser(name string, age int, email string) User {
	user := User{
		Id:    nextID,
		Name:  name,
		Age:   age,
		Email: email,
	}
	nextID++
	users = append(users, user)
	return user
}

func UpdateUser(id int, user User) {
	u := GetUser(id)
	if u != nil {
		u.Name = user.Name
		u.Age = user.Age
		u.Email = user.Email
	}
}

func DeleteUser(id int) {
	for i, user := range users {
		if user.Id == id {
			users = slices.Delete(users, i, i+1)
			break
		}
	}
}
