package initializer

import (
	// "blog-crud/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=pg-17b26418-zinpainghtet-24dc.i.aivencloud.com user=avnadmin password=AVNS_AzpKrCNsRXGopmeiyYp dbname=defaultdb port=21847  sslmode=require"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connection to database")
		return
	}
	fmt.Println("DB connection established ", DB)

}
