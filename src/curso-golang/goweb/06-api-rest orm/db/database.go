package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:123456789@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"

var Database = func () (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil{
		fmt.Println("Error en la conexion a la base de datos", err)
		panic(err)
	}else{
		fmt.Println("Conexion exitosa")
		return db
	}
}()


