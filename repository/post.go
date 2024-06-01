package repository

import (
	"golang-backend-template/model"

	"gorm.io/gorm"
)

type IPostRepository interface {
	ReadPostById(id uint) (model.Post, error)
	SeePostById(id uint) (model.JoinedReadPost, error)
	// '/post', '/post?user' AND ALL OTHER FILTER WILL CALL SeeAllPost
	SeeAllPost(filter model.PostSearchFilter) ([]model.JoinedReadPost, error)
	AddPost(postData model.PostBody, ownerUserId uint) error
	EditPost(postData model.Post) error
	DeletePost(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{db}
}

func (p postRepository) ReadPostById(id uint) (model.Post, error) {
	postData := model.Post{}
	if err := p.db.Where("id = ?", id).First(&postData).Error; err != nil {
		return model.Post{}, err
	}
	return postData, nil
}

func (p postRepository) SeePostById(id uint) (model.JoinedReadPost, error) {
	postData := model.JoinedReadPost{}
	if err := p.db.Table("posts").Select("posts.id, users.username, posts.title, posts.description, posts.created_at").
		Joins("inner join users on users.id = posts.user_id").
		// JOIN DOESN'T FILTER SOFT DELETED DATA
		Where("posts.id = ? AND posts.deleted_at IS NULL", id).First(&postData).Error; err != nil {
		return model.JoinedReadPost{}, err
	}
	return postData, nil
}

func (p postRepository) SeeAllPost(filter model.PostSearchFilter) ([]model.JoinedReadPost, error) {
	postsData := []model.JoinedReadPost{}
	query := p.db.Table("posts").Select("posts.id, users.username, posts.title, posts.description, posts.created_at").
		Joins("inner join users on users.id = posts.user_id").
		// JOIN DOESN'T FILTER SOFT DELETED DATA
		Where("posts.deleted_at IS NULL")
	if filter.Username != "" {
		query = query.Where("username = ?", filter.Username)
	}
	if err := query.Scan(&postsData).Error; err != nil {
		return nil, err
	}
	return postsData, nil
}

func (p postRepository) AddPost(postData model.PostBody, ownerUserId uint) error {
	newPostData := model.Post{Title: postData.Title, Description: postData.Description, UserId: ownerUserId}
	if err := p.db.Create(&newPostData).Error; err != nil {
		return err
	}
	return nil
}

func (p postRepository) EditPost(postData model.Post) error {
	if err := p.db.Save(&postData).Error; err != nil {
		return err
	}
	return nil
}

func (p postRepository) DeletePost(id uint) error {
	if err := p.db.Where("id = ?", id).Delete(&model.Post{}).Error; err != nil {
		return err
	}
	return nil
}
