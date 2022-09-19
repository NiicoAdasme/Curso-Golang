package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	
	// fmt.Println(db.ExistsTable("users"))
	// db.TruncateTable("users")
	// db.CreateTable(models.UserSchema, "users")
	user := models.CreateUser("alex","alex123","alex@gmail.com")
	fmt.Println(user)

	// user := models.GetUser(2)
	// fmt.Println(user)

	// user.Username = "carlos"
	// user.Password = "carlos123"
	// user.Email = "carlos@gmail.com"
	// user.Save()
	// user.Delete()

	// db.TruncateTable("users")
	fmt.Println(models.ListUsers())


	db.Close()
	// db.Ping()
}