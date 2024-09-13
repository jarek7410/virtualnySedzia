package model

import (
	"github.com/guregu/null/v5"
	"gorm.io/gorm"
	"virtualnySedziaServer/database"
)

type Issue struct {
	gorm.Model
	ID         uint   `gorm:"primarykey"`
	ShortTitle string `json:"short_title"`
	Title      string `json:"title"`
	Deal
	Text     string    `json:"text,required"`
	Author   string    `json:"author,omitempty"`
	UserID   null.Int  `json:"user_id,omitempty"`
	Comments []Comment `json:"comments"`
}
type Deal struct {
	Vulnerable string   `json:"vulnerable"`
	Dealer     string   `json:"dealer"`
	HandN      Hand     `gorm:"foreignKey:HandNID" json:"handN,omitempty"`
	HandS      Hand     `gorm:"foreignKey:HandSID" json:"handS,omitempty"`
	HandW      Hand     `gorm:"foreignKey:HandWID" json:"handW,omitempty"`
	HandE      Hand     `gorm:"foreignKey:HandEID" json:"handE,omitempty"`
	HandNID    null.Int `json:"N,omitempty"`
	HandSID    null.Int `json:"S,omitempty"`
	HandWID    null.Int `json:"W,omitempty"`
	HandEID    null.Int `json:"E,omitempty"`
	Minmax     string   `json:"minmax"`
}
type Hand struct {
	ID       uint   `gorm:"primarykey"`
	Spike    string `json:"spike"`
	Hearts   string `json:"hearts"`
	Diamonds string `json:"diamonds"`
	Clubs    string `json:"clubs"`
}

func (I *Issue) Save() error {
	if err := database.Re.DB.Save(&I).Error; err != nil {
		return err
	}
	return nil
}

func (I *Issue) GetByID() error {
	if err := database.Re.DB.Preload("HandN").Preload("HandS").Preload("HandW").Preload("HandE").
		First(&I).Error; err != nil {
		return err
	}
	return nil
}
func (I *Issue) LoadWithComments() error {
	if err := database.Re.DB.Preload("Comments").Preload("Comments.Comments").Preload("HandN").Preload("HandS").Preload("HandW").Preload("HandE").
		First(&I).Error; err != nil {
		return err
	}
	return nil
}

func (I *Issue) Delete() error {
	if err := database.Re.DB.Delete(&I).Error; err != nil {
		return err
	}
	return nil
}

func IssueGetWithOffsetAndLimit(offset, limit int) ([]Issue, error) {
	var I []Issue
	if err := database.Re.DB.Model(&Issue{}).Preload("HandN").Preload("HandS").Preload("HandW").Preload("HandE").Limit(limit).Offset(offset).
		Find(&I).Error; err != nil {
		return nil, err
	}
	return I, nil
}
