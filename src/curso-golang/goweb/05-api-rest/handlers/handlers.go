package handlers

import (
	"apirest/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request){

	// rw.Header().Set("Content-Type", "application/json")
	/*
	// db.Connect()
	// users, _ := models.ListUsers()
	// db.Close()
	
	// output, _ := json.Marshal(users)
	// fmt.Fprintln(rw, string(output)) */

	if users, err := models.ListUsers(); err != nil{
		models.SendNotFound(rw)
	}else{
		models.SendData(rw, users)
	}
}

func GetUser(rw http.ResponseWriter, r *http.Request){

	// rw.Header().Set("Content-Type", "application/json")
	/*
	// // Obtener id
	// vars := mux.Vars(r)

	// userId, _ := strconv.Atoi(vars["id"])

	// db.Connect()
	// user, _ := models.GetUser(userId)
	// db.Close()
	
	// output, _ := json.Marshal(user)
	// fmt.Fprintln(rw, string(output)) */

	if user, err := getUserByRequest(r); err != nil{
		models.SendNotFound(rw)
	}else{
		models.SendData(rw, user)
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request){
	// rw.Header().Set("Content-Type", "application/json")
	
	// Obtener registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil{
		models.SendUnprocessableEntity(rw)
	}else{
		user.Save()
		models.SendData(rw, user)
	}
		
}

func UpdateUser(rw http.ResponseWriter, r *http.Request){
	// rw.Header().Set("Content-Type", "application/json")

	// Obtener registro
	// user := models.User{}
	var userId int64

	if user, err := getUserByRequest(r); err != nil{
		models.SendNotFound(rw)
	}else{
		userId = user.Id	
	}
	
	user := models.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil{
		models.SendUnprocessableEntity(rw)
	}else{
		user.Id = userId
		user.Save()
		models.SendData(rw, user)
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request){
	/*rw.Header().Set("Content-Type", "application/json")

	// Obtener id
	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user, _ := models.GetUser(userId)
	user.Delete()
	db.Close()
	
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))*/
	
	if user, err := getUserByRequest(r); err != nil{
		models.SendNotFound(rw)
	}else{
		user.Delete()
		models.SendData(rw, user)
	}
}

func getUserByRequest(r *http.Request) (models.User, error){
	// Obtener ID
	vars := mux.Vars(r)

	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil{
		return *user, err
	}else{
		return *user, nil
	}
}