package project

import (
	"github.com/pclubiitk/fep-backend/student"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		logrus.Fatal("Failed to connect to project database: ", err)
		panic(err)
	}

	db = database
	err = db.AutoMigrate(&Project{}, &Application{}, &student.Student{}, &SelectedStudent{})
	if err != nil {
		logrus.Fatal("Failed to migrate project ,application and student database: ", err)
		panic(err)
	}

	logrus.Info("Connected to project and application database")
}

func init() {
	openConnection()
}
