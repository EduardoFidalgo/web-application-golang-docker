package models

import (
	"log"

	"github.com/webservice-golang-api/database"
)

type User struct {
	Id    int
	Name  string
	Email string
	Phone string
}

func GetUsers() []User {
	db := database.Connect()
	defer db.Close()

	sql := "SELECT * FROM users ORDER BY id ASC"

	all, err := db.Query(sql)
	if err != nil {
		panic(err.Error())
	}

	var usersList []User

	for all.Next() {
		user := User{}

		err := all.Scan(&user.Id, &user.Name, &user.Email, &user.Phone)
		if err != nil {
			panic(err.Error())
		}

		usersList = append(usersList, user)
	}

	return usersList
}

func StoreUser(name string, email string, phone string) {
	db := database.Connect()

	sql := "INSERT INTO users (name, email, phone) values (?,?,?)"

	insert, err := db.Prepare(sql)
	if err != nil {
		log.Panic(err.Error())
	}

	insert.Exec(name, email, phone)

	defer db.Close()
}

func UpdateView(id int) User {
	db := database.Connect()

	userData, err := db.Query("SELECT * from users WHERE id = ?", id)
	if err != nil {
		log.Panic(err.Error())
	}

	userToUpdate := User{}

	for userData.Next() {
		var id int
		var name, email, phone string

		err = userData.Scan(&id, &name, &email, &phone)
		if err != nil {
			log.Panic(err.Error())
		}

		userToUpdate.Id = id
		userToUpdate.Name = name
		userToUpdate.Email = email
		userToUpdate.Phone = phone
	}

	defer db.Close()

	return userToUpdate
}

func Update(id int, name string, email string, phone string) {
	db := database.Connect()

	sql := "UPDATE users SET name= ?, email= ?, phone= ? WHERE id = ?"

	updateData, err := db.Prepare(sql)
	if err != nil {
		log.Panic(err.Error())
	}

	updateData.Exec(name, email, phone, id)

	defer db.Close()
}

func Destroy(id int) {
	db := database.Connect()

	sql := "DELETE FROM users WHERE id = ?"

	destroy, err := db.Prepare(sql)
	if err != nil {
		log.Panic(err.Error())
	}

	destroy.Exec(id)

	defer db.Close()
}
