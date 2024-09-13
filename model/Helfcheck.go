package model

import "afmib_server/database"

func HealthCheck() error {
	db, _ := database.Re.DB.DB()
	return db.Ping()

}
