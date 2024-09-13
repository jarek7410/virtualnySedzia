package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guregu/null/v5"
	"net/http"
	"strconv"
	"virtualnySedziaServer/model"
	"virtualnySedziaServer/securiry"
)

func GetComment(ctx *gin.Context) {
	id, erratoi := strconv.Atoi(ctx.Param("id"))
	if erratoi != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id should be uint"})
		return
	}
	comment := model.Comment{}
	comment.ID = uint(id)
	if err := comment.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	ctx.JSON(http.StatusOK, comment)
}

func PostComment(ctx *gin.Context) {
	comment := model.Comment{}
	if err := ctx.ShouldBind(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := securiry.CurrentUser(ctx)
	if err == nil {
		comment.UserID = null.IntFrom(int64(user.ID))
	}

	if err := comment.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, comment)

}

func DeleteComment(ctx *gin.Context) {
	id, erratoi := strconv.Atoi(ctx.Param("id"))
	if erratoi != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id should be uint"})
		return
	}
	comment := model.Comment{}
	comment.ID = uint(id)

	if err := comment.Delete(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, comment)

}
