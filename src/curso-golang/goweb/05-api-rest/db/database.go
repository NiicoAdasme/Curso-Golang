package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(localhost:3306)/database
const url = "root:123456789@tcp(localhost:3306)/goweb_db"

// Guarda la conexion
var db *sql.DB

// Realizar la conexion
func Connect(){
	connection, err := sql.Open("mysql",url)
	if err != nil{
		panic(err)
	}

	fmt.Println("Conexion exitosa")
	db = connection


}

// Cerrar la conexion
func Close(){
	db.Close()
}

// Verificar la conexion
func Ping(){
	if err := db.Ping(); err != nil{
		panic(err)
	}
}

// Verificar si una tabla existe
func ExistsTable(tableName string) bool{
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil{
		fmt.Println("Error:", err)
	}

	return rows.Next()
}

// Crea una tabla
func CreateTable(schema string, name string){
	if !ExistsTable(name){
		_, err := Exec(schema)
		if err != nil{
			fmt.Println(err)
		}
	}
	
}

// Reiniciar registro de una tabla
func TruncateTable(tableName string){
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}



// Polimorfismo de Exec
func Exec(query string, args ...any) (sql.Result, error){
	Connect()

	result, err := db.Exec(query, args...)
	Close()
	if err != nil{
		fmt.Println(err)
	}

	return result, err
}

// Polimorfismo de Query
func Query(query string, args ...any) (*sql.Rows, error){
	Connect()
	rows, err := db.Query(query, args...)
	Close()

	if err != nil{
		fmt.Println(err)
	}

	return rows, err
}
