package repo

import (
	"bloggers/database"
	"bloggers/models"
)

type CommentRepository interface {
	CheckDupComment(comment *models.Comment) bool
	CreateComment(comment *models.Comment) error
	GetCommentByPostID(comment *models.Comment, postID *models.Post) (*models.Comment, error)
}

type CommentRepo struct {
}

func (r *CommentRepo) CheckDupComment(comment *models.Comment) bool {
	var count int64
	database.DB.Model(&models.Comment{}).Where(&models.Comment{Content: comment.Content}).Count(&count) //counts how many lines exist in the database with the same content
	return count > 0
}

func (r *CommentRepo) CreateComment(comment *models.Comment) error {
	return database.DB.Create(comment).Error //handles error handling as well
}

func (r *CommentRepo) GetCommentByPostID(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := database.DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error //preload allows for gorm to find the db relations through the foreign key
	if err != nil {
		return nil, err
	}

	return comments, nil
}
