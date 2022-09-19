package models

import (
	"apirest/db"
)

type User struct {
	Id       int64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

// Constructor de usuario
func NewUser(userName, password, email string) *User {

	user := &User{
		Username: userName,
		Password: password,
		Email:    email,
	}

	return user
}

// Crear usuario e insertarlo en la bd
func CreateUser(userName, password, email string) *User{
	user := NewUser(userName, password, email)
	user.Save()
	return user
}

// Insertar un registro a la bd
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _:= db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

// Listar los registros
func ListUsers() (Users, error){
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	
	rows, err := db.Query(sql)

	for rows.Next(){
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users, err
}

// Listar un registro
func GetUser(id int) (*User, error){
	user := NewUser("", "", "")
	
	sql := "SELECT id, username, password, email FROM users WHERE id=?"

	if rows, err := db.Query(sql, id); err != nil{
		return nil, err
	}else{
		for rows.Next(){
			rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		}
		return user, nil
	}


	
}

// Actualizar registro
func (user *User) update(){
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

// Guardar o editar registro
func (user *User) Save(){
	if user.Id == 0 {
		user.insert()
	}else{
		user.update()
	}
}

// Eliminar un registro
func (user *User) Delete(){
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
}

