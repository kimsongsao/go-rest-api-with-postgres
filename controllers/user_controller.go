package controllers

import (
	"fmt"
	"golangrestapi/config"
	"golangrestapi/models"
	"golangrestapi/requests"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/swaggo/swag/example/celler/httputil"
	"golang.org/x/crypto/bcrypt"
)

// SignUp godoc
//
//	@Summary		SignUp
//	@Description	Create new an user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		requests.SignupRequest	true	"Signup Data"
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/users/signup [post]
func Signup(ctx *gin.Context) {
	var body requests.SignupRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	// var body struct {
	// 	Name     string
	// 	Email    string
	// 	Password string
	// }

	// if ctx.BindJSON(&body) != nil {
	// 	ctx.JSON(400, gin.H{"error": "Bad request"})
	// 	return
	// }

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}

	result := config.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(200, gin.H{"data": user})

}

// Login godoc
//
//	@Summary		Login
//	@Description	Login
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		requests.LoginRequest	true	"Login Data"
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/users/login [post]
func Login(ctx *gin.Context) {
	var body requests.LoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	// var body struct {
	// 	Email    string
	// 	Password string
	// }

	// if ctx.BindJSON(&body) != nil {
	// 	ctx.JSON(400, gin.H{"error": "bad request"})
	// 	return
	// }

	var user models.User

	result := config.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(500, gin.H{"error": "error signing token"})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "localhost", false, true)
	ctx.JSON(200, gin.H{"data": "You are logged in!", "token": tokenString})
}

// Validate godoc
//
//	@Summary		Validate
//	@Description	Validate
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200		{object} 	models.User
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/users/auth [post]
func Validate(ctx *gin.Context) {
	user, err := ctx.Get("user")
	if !err {
		ctx.JSON(500, gin.H{"error": err})
		return
	}
	ctx.JSON(200, gin.H{"data": "You are logged in!", "user": user})
}

// Logout godoc
//
//	@Summary		Logout
//	@Description	Logout
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/users/logout [post]

func Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "localhost", false, true)
	ctx.JSON(200, gin.H{"data": "You are logged out!"})
}

// GetUsers godoc
//
//	@Summary		Show all Users
//	@Description	get users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			q	query		string	false	"name search by q"
//	@Success		200	{array}		models.User
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Security		ApiKeyAuth
//	@Router			/users [get]
func GetUsers(ctx *gin.Context) {
	var users []models.User

	err := config.DB.Model(&models.User{}).Preload("Posts").Find(&users).Error

	if err != nil {
		fmt.Println(err)
		ctx.JSON(500, gin.H{"error": "error getting users"})
		return
	}

	ctx.JSON(200, gin.H{"data": users})

}
