package model

import (
	"github.com/guregu/null/v5"
	"gorm.io/gorm"
	"virtualnySedziaServer/database"
)

type Comment struct {
	gorm.Model
	ShortTile string    `json:"short_tile"`
	Title     string    `json:"title"`
	Text      string    `json:"text,required"`
	UserID    null.Int  `json:"user_id"`
	IssueID   null.Int  `json:"issue_id"`
	Comments  []Comment `json:"comments"`
	CommentID null.Int  `json:"comment_id"`
}

func (c *Comment) Save() error {
	if err := database.Re.DB.Save(&c).Error; err != nil {
		return err
	}
	return nil
}
func (c *Comment) GetByID() error {
	if err := database.Re.DB.First(&c).Error; err != nil {
		return err
	}
	return nil
}
func (c *Comment) Delete() error {
	if err := database.Re.DB.Delete(&c).Error; err != nil {
		return err
	}
	return nil
}
