//data access layer
//talks directly to the database
//handles the things the service layer doesn't need to know

package repo

import (
	"bloggers/database"
	"bloggers/models"
)

type PostRepository interface {
}

type PostRepo struct{}

func (r *PostRepo) CreatePost(post *models.Post) error {
	err := database.DB.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepo) GetPostByID(id uint) (*models.Post, error) { //return pointer to avoid copies - can be resued anytime i need to get a post
	var post models.Post
	err := database.DB.Preload("User").Where("id = ?", id).First(&post).Error
	if err != nil {
		return nil, err
	}

	return &post, nil
}

/*save will overwrite fields that aren't updated with empty values
func (r *PostRepo) EditPost(post *models.Post) error {
	err := database.DB.Save(post).Error
	if err != nil {
		return err
	}
	return nil
} */

func (r *PostRepo) EditPost(postID uint, updates map[string]interface{}) error {
	//updates only fields provided
	err := database.DB.Model(&models.Post{}).
		Where("id = ?", postID).
		Updates(updates).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepo) DeletePost(post *models.Post, postID uint) error {
	err := database.DB.Delete(&post, postID).Error

	if err != nil {
		return err
	}

	return nil
}
