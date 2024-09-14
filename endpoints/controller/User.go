package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	dto2 "virtualnySedziaServer/dto"
	model2 "virtualnySedziaServer/model"
	"virtualnySedziaServer/securiry"
)

// Register user
func Register(context *gin.Context) {
	var input dto2.Register

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model2.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   3,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})

}

// User Login
func Login(context *gin.Context) {
	var input dto2.Login

	if err := context.ShouldBindJSON(&input); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	user, err := model2.GetUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusGone, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidateUserPassword(input.Password)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	jwt, err := securiry.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": jwt, "username": input.Username, "message": "Successfully logged in"})

}

// get all users
func GetUsers(context *gin.Context) {
	var user []model2.User
	err := model2.GetUsers(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}

// get all users
func GetUsersPublic(context *gin.Context) {

	var user []model2.User
	err := model2.GetUsers(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	var userPublic []dto2.PublicUserInfo
	users, _ := json.Marshal(user)
	err = json.Unmarshal(users, &userPublic)

	if err != nil {
		context.JSON(501, gin.H{"error": err})
		return
	}

	context.JSON(http.StatusOK, userPublic)
}

// get user by id
func GetUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var user model2.User
	err := model2.GetUser(&user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}
func GetMyUser(ctx *gin.Context) {
	user, err := securiry.CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	userMe := dto2.UserDataMe{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		PID:      user.PID,
		Name:     user.Name,
		Surname:  user.Surname,
	}

	ctx.JSON(http.StatusOK, userMe)
}
func ChangeMyUser(ctx *gin.Context) {
	user, err := securiry.CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	changesUser := dto2.UserChangeDataMe{}
	if err := ctx.ShouldBindJSON(&changesUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"changeUser": changesUser, "User": user})
		return
	}
	log.Println(changesUser)
	m, e := json.Marshal(changesUser)
	if e != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	e2 := json.Unmarshal(m, &user)
	if e2 != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	err1 := user.Update()
	if err1 != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"changeUser": changesUser, "User": user})
}

// update user
func UpdateUser(c *gin.Context) {
	//var input model.Update
	var User model2.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := model2.GetUser(&User, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = c.BindJSON(&User)
	if err != nil {
		return
	}
	err = User.UpdateAsAdmin()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, User)
}

func GerUserActions(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	user := model2.User{ID: uint(id)}
	if err := user.GetActions(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func ChangeMyUserPassword(ctx *gin.Context) {
	user, err := securiry.CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}
	changesPassword := dto2.ChangePassword{}
	if err := ctx.ShouldBindJSON(&changesPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	user.Password = changesPassword.NewPassword

	if err := user.ChangePassword(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
