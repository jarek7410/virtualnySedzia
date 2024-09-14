package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"virtualnySedziaServer/database"
	"virtualnySedziaServer/endpoints"
	model2 "virtualnySedziaServer/model"
)

func main() {
	loadEnv()

	_ = loadDatabase()

	serveApplication()

}

func loadEnv() {
	//undo comment for debug and development
	err := godotenv.Load(".env")
	if err != nil {

		log.Println("Error loading .env file")
		shell := os.Getenv("SHELL")
		log.Println(shell)

	}
	log.Println(".env file loaded successfully")
}

func loadDatabase() *database.Repo {
	r := database.InitDatabase()
	err := Migrate(r)
	if err != nil {
		log.Fatalln("database migration do not work")
	}

	seedData(r)

	database.TestConnection(r.DB)

	return r
}
func seedData(r *database.Repo) {
	var roles = []model2.Role{{Name: "admin", Description: "Administrator role"}, {Name: "customer", Description: "Authenticated customer role"}, {Name: "anonymous", Description: "Unauthenticated customer role"}}
	var user = []model2.User{{
		Username: os.Getenv("ADMIN_USERNAME"),
		Email:    os.Getenv("ADMIN_EMAIL"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		RoleID:   1,
		Name:     "adminName",
		Surname:  "adminSurname",
		PID:      "-1",
	}}
	r.DB.Save(&roles)
	r.DB.Save(&user)
}
func Migrate(r *database.Repo) error {
	err := r.DB.AutoMigrate(
		&model2.Role{},
		&model2.User{},
		&model2.Comment{},
		&model2.Issue{},
		&model2.Hand{},
	)
	if err != nil {
		return err
	}
	return nil
}

func serveApplication() {
	routes := endpoints.NewRouts()
	routes.AddPaths()
	routes.AddAuthPaths()
	routes.ServeApplication()

	routes.Start(os.Getenv("PORT"))
	fmt.Println("Server running on port 8000")
}
