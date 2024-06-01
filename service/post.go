package service

import (
	"errors"
	"golang-backend-template/model"
	"golang-backend-template/repository"
)

type IPostService interface {
	SeePostById(id uint) (model.JoinedReadPost, error)
	SeeAllPost(filter model.PostSearchFilter) ([]model.JoinedReadPost, error)
	AddPost(postData model.PostBody, ownerUserId uint) error
	EditPost(id uint, userId uint, role string, postData model.PostBody) error
	DeletePost(id uint, userId uint, role string) error
}

type postService struct {
	postRepo repository.IPostRepository
}

func NewPostService(postRepo repository.IPostRepository) *postService {
	return &postService{postRepo}
}

func (p postService) SeePostById(id uint) (model.JoinedReadPost, error) {
	postData, err := p.postRepo.SeePostById(id)
	if err != nil {
		return model.JoinedReadPost{}, err
	}
	return postData, nil
}

func (p postService) SeeAllPost(filter model.PostSearchFilter) ([]model.JoinedReadPost, error) {
	postsData, err := p.postRepo.SeeAllPost(filter)
	if err != nil {
		return nil, err
	}
	return postsData, nil
}

func (p postService) AddPost(postData model.PostBody, ownerUserId uint) error {
	if postData.Title == "" || postData.Description == "" {
		return errors.New("title or description shouldn't be empty")
	}
	if err := p.postRepo.AddPost(postData, ownerUserId); err != nil {
		return err
	}
	return nil
}

func (p postService) EditPost(id uint, userId uint, userRole string, postData model.PostBody) error {
	if postData.Title == "" && postData.Description == "" {
		return errors.New("title and description can't be both empty")
	}
	editedData, err := p.postRepo.ReadPostById(id)
	if err != nil {
		return err
	}
	// NEED TO BE THE OWNER OF THE POST OR ADMIN TO EDIT THE POST
	if editedData.UserId != userId && userRole != "admin" {
		return errors.New("not the owner of this post")
	}
	if postData.Title != "" {
		editedData.Title = postData.Title
	}
	if postData.Description != "" {
		editedData.Description = postData.Description
	}
	if err := p.postRepo.EditPost(editedData); err != nil {
		return err
	}
	return nil
}

func (p postService) DeletePost(id uint, userId uint, userRole string) error {
	deletedData, err := p.postRepo.ReadPostById(id)
	if err != nil {
		return err
	}
	// NEED TO BE THE OWNER OF THE POST OR ADMIN TO EDIT THE POST
	if deletedData.UserId != userId && userRole != "admin" {
		return errors.New("not the owner of this post")
	}
	if err := p.postRepo.DeletePost(id); err != nil {
		return err
	}
	return nil
}
