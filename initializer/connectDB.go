package initializer

import (
	// "blog-crud/models"
	"fmt"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// var err error
	// db connection on read me

	// if err != nil {
	// 	fmt.Println("Error connection to database")
	// 	return
	// }
	fmt.Println("DB connection established ", DB)

}
