package config

import (
	"fmt"
	"os"

	"MyGO.com/m/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DBUser string
	DBHost string
	DBName string
	DBPass string
}

func SetupDBConnection() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error on loading env or there is not .env file in your project")
	}

	var dbConfig DBConfig
	dbConfig.DBUser = os.Getenv("DBUSER")
	dbConfig.DBHost = os.Getenv("DBHOST")
	dbConfig.DBName = os.Getenv("DBNAME")
	dbConfig.DBPass = os.Getenv("DBPASSWORD")

	fmt.Println("Here is DB name in db config >>>>>>>>>>>>>>", dbConfig.DBUser)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBHost, dbConfig.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	errMigrate := db.AutoMigrate(&model.User{})
	if errMigrate != nil {
		return nil
	}

	fmt.Println("Connect DB success fully")

	return db
}
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	panic(dbSQL.Close())
}
