package controllers

import (
	"fmt"
	"golangrestapi/config"
	"golangrestapi/models"
	"golangrestapi/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

// Post godoc
//
//	@Summary		Show all posts
//	@Description	get posts
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			q	query		string	false	"name search by q"
//	@Success		200	{array}		models.Post
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/posts [get]
func GetPosts(ctx *gin.Context) {
	var posts []models.Post
	result := config.DB.Find(&posts)
	if result.Error != nil {
		httputil.NewError(ctx, http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": posts})

}

// Post godoc
//
//	@Summary		Show a post
//	@Description	get posts by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Post Id"
//	@Success		200	{object}	models.Post
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/posts/{id} [get]
func GetPost(ctx *gin.Context) {
	var post models.Post
	result := config.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		httputil.NewError(ctx, http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": post})
}

// Post godoc
//
//	@Summary		Create a post
//	@Description	Create post
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			request	body		requests.PostRequest	true	"Post Data"
//	@Success		200		{object}	models.Post
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/posts [post]
func CreatePost(ctx *gin.Context) {
	var body requests.PostRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	// same as before...
	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
		return
	}

	body.UserID = user.(models.User).ID
	post := models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Draft: body.Draft, Author: body.Author, UserID: body.UserID}

	fmt.Println(post)
	result := config.DB.Create(&post)
	if result.Error != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": post})

}

// Post godoc
//
//	@Summary		Update a post
//	@Description	Update post by ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int						true	"Post Id"
//	@Param			request	body		requests.PostRequest	true	"Post Update Data"
//	@Success		200		{object}	models.Post
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/posts/{id} [put]
func UpdatePost(ctx *gin.Context) {
	var body requests.PostRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var post models.Post

	result := config.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		httputil.NewError(ctx, http.StatusInternalServerError, result.Error)
		return
	}

	config.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body, Likes: body.Likes, Draft: body.Draft, Author: body.Author})

	ctx.JSON(http.StatusOK, gin.H{"data": post})

}

// DeletePost godoc
//
//	@Summary		Delete a post
//	@Description	Delete by post ID
//	@Tags			posts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Post ID"	Format(int64)
//	@Success		204	{object}	models.Post
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/posts/{id} [delete]
func DeletePost(ctx *gin.Context) {

	id := ctx.Param("id")

	config.DB.Delete(&models.Post{}, id)

	ctx.JSON(http.StatusOK, gin.H{"data": "post has been deleted successfully"})

}
