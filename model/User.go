package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"html"
	"strings"
	"virtualnySedziaServer/database"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	RoleID   uint8  `gorm:"not null;DEFAULT:3" json:"role_id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Name     string `gorm:"size:255" json:"name"`
	Surname  string `gorm:"size:255" json:"surname"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	PID      string `gorm:"size:20" json:"pid"`
	Role     Role   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Comments []Comment
	Issues   []Issue
}

// Save user details
func (user *User) Save() (*User, error) {
	if err := database.Re.DB.Create(&user).Error; err != nil {
		return &User{}, err
	}
	return user, nil
}

// Generate encrypted password
func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

// Validate user password
func (user *User) ValidateUserPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// Get all users
func GetUsers(User *[]User) (err error) {
	err = database.Re.DB.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Get user by username
func GetUserByUsername(username string) (User, error) {
	var user User
	err := database.Re.DB.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Get user by id
func GetUserById(id uint) (User, error) {
	var user User
	err := database.Re.DB.Where("id=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// Get user by id
func GetUser(User *User, id int) (err error) {
	err = database.Re.DB.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

// Update user
func (u *User) Update() (err error) {
	err = database.Re.DB.Omit("password").Save(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (user *User) GetActions() error {
	if err := database.Re.DB.Preload("Issues").Preload("Comments").First(&user).Error; err != nil {
		return err
	}
	return nil
}
