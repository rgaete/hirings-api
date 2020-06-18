package utils

import (
	"log"

	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// GetConnection obtiene una conexión a la base de datos
func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:my-secret-pw@tcp(192.168.64.1:3306)/contacts")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
