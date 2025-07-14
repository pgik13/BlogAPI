package services

import (
	"bloggers/models"
	"bloggers/repo"
	"errors"
)

type CommentService struct {
	Repo repo.CommentRepo
}

func (s *CommentService) CreateComment(comment *models.Comment) error {

	ok := s.Repo.CheckDupComment(comment)
	if ok {
		return errors.New("duplicated comment")
	}

	err := s.Repo.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (s *CommentService) GetCommentByPostID(postID uint) ([]models.Comment, error) {
	return s.Repo.GetCommentByPostID(postID)
}
