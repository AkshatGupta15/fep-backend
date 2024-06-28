package application

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func openConnection() {
	host := viper.GetString("DATABASE.HOST")
	port := viper.GetString("DATABASE.PORT")
	password := viper.GetString("DATABASE.PASSWORD")

	dbName := viper.GetString("DBNAME.APPLICATION")
	user := viper.GetString("DATABASE.USER")

	dsn := "host=" + host + " user=" + user + " password=" + password
	dsn += " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Kolkata"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		logrus.Fatal("Failed to connect to project database: ", err)
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&Application{})
	if err != nil {
		logrus.Fatal("Failed to migrate application database: ", err)
		panic(err)
	}

	logrus.Info("Connected to application database")
}

func init() {
	openConnection()
}
