package model

import "virtualnySedziaServer/database"

func HealthCheck() error {
	db, _ := database.Re.DB.DB()
	return db.Ping()

}
