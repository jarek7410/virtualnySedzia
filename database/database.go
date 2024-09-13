package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type Repo struct {
	DB *gorm.DB
}

var Re *Repo

func InitDatabase() *Repo {
	var db = connectDB()
	Re = &Repo{
		DB: db,
	}
	return Re
}

func connectDB() *gorm.DB {
	log.Println("loading environments")
	var dbHost = os.Getenv("DB_HOST")
	var dbPost = os.Getenv("DB_PORT")
	var dbUser = os.Getenv("DB_USER")
	var dbName = os.Getenv("DB_NAME")
	var dbPass = os.Getenv("DB_PASSWORD")

	log.Println("Connecting to db")

	psqlconn :=
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPost, dbUser, dbPass, dbName)
	log.Println(psqlconn)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err, "fail to connect to db")
	}

	log.Println("`Successfully connected to the database")
	return db
}

func TestConnection(db *gorm.DB) {
	s, err := db.DB()
	if err != nil {
		return
	}
	go func() {
		for {
			errPing := s.Ping()
			if errPing != nil {
				log.Println("db connection error: ", err)
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

//
//func (r Repo) Migrate() error {
//	err := r.DB.AutoMigrate(&model2.Role{}, &model2.User{})
//	if err != nil {
//		return err
//	}
//	return nil
//}
