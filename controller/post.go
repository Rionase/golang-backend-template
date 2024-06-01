package controller

import (
	"golang-backend-template/model"
	"golang-backend-template/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IPostController interface {
	SeePostById(*gin.Context)
	SeeAllPost(*gin.Context)
	AddPost(*gin.Context)
	EditPost(*gin.Context)
	DeletePost(*gin.Context)
}

type postController struct {
	postService service.IPostService
}

func NewPostController(postService service.IPostService) *postController {
	return &postController{postService}
}

func (p postController) SeePostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "wrong post_id param format"})
		return
	}
	postData, err := p.postService.SeePostById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": postData})
}

func (p postController) SeeAllPost(c *gin.Context) {
	filter := model.PostSearchFilter{}

	username := c.Query("username")
	if username != "" {
		filter.Username = username
	}

	postsData, err := p.postService.SeeAllPost(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": postsData})
}

func (p postController) AddPost(c *gin.Context) {
	user_id := c.MustGet("id").(uint)

	var postData model.PostBody
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	if err := p.postService.AddPost(postData, user_id); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "success add post"})
}

func (p postController) EditPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "wrong post_id param format"})
		return
	}

	user_id := c.MustGet("id").(uint)
	role := c.MustGet("role").(string)

	var postData model.PostBody
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	if err := p.postService.EditPost(uint(id), user_id, role, postData); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "success edit post"})
}

func (p postController) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: "wrong post_id param format"})
		return
	}

	user_id := c.MustGet("id").(uint)
	role := c.MustGet("role").(string)
	if err := p.postService.DeletePost(uint(id), user_id, role); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "success delete post"})
}
