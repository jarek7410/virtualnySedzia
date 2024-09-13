package model

import (
	"afmib_server/database"
	"gorm.io/gorm"
)

// Role model
type Role struct {
	gorm.Model
	ID          uint8  `gorm:"primary_key"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `gorm:"size:255;not null" json:"description"`
}

// Create a role
func CreateRole(Role *Role) (err error) {
	err = database.Re.DB.Create(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get all roles
func GetRoles(Role *[]Role) (err error) {
	err = database.Re.DB.Find(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Get role by id
func GetRole(Role *Role, id int) (err error) {
	err = database.Re.DB.Where("id = ?", id).First(Role).Error
	if err != nil {
		return err
	}
	return nil
}

// Update role
func UpdateRole(Role *Role) (err error) {
	database.Re.DB.Save(Role)
	return nil
}
