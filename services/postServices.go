//what to do and how to do it
//sets the rules for the application
//validation checks

package services

import (
	"bloggers/models"
	"bloggers/repo"
	"errors"
	"strings"
)

type PostService struct {
	Repo repo.PostRepo
}

func (s *PostService) CreatePost(post *models.Post) error {

	if strings.TrimSpace(post.Title) == "" {
		return errors.New("title cannot be empty")
	}

	//save post to database
	err := s.Repo.CreatePost(post)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) GetPostByID(id uint) (*models.Post, error) {
	return s.Repo.GetPostByID(id)
}

func (s *PostService) EditPost(postID uint, updates map[string]interface{}) error {
	err := s.Repo.EditPost(postID, updates)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) DeletePost(post *models.Post, postID uint) error {
	err := s.Repo.DeletePost(post, postID)
	if err != nil {
		return err
	}

	return nil
}
