package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guregu/null/v5"
	"net/http"
	"strconv"
	"virtualnySedziaServer/model"
	"virtualnySedziaServer/securiry"
)

func GetIssue(ctx *gin.Context) {
	id, erratoi := strconv.Atoi(ctx.Param("id"))
	if erratoi != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id should be uint"})
		return
	}
	issue := model.Issue{}
	issue.ID = uint(id)

	if err := issue.GetByID(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, issue)
}
func PostIssue(ctx *gin.Context) {
	issue := model.Issue{}
	if err := ctx.ShouldBind(&issue); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := securiry.CurrentUser(ctx)
	if err == nil {
		issue.UserID = null.IntFrom(int64(user.ID))
	}

	if err := issue.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, issue)
}

func DeleteIssue(ctx *gin.Context) {
	id, erratoi := strconv.Atoi(ctx.Param("id"))
	if erratoi != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id should be uint"})
		return
	}
	issue := model.Issue{}
	issue.ID = uint(id)

	if err := issue.Delete(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, issue)
}

func GetIssueComments(ctx *gin.Context) {
	id, erratoi := strconv.Atoi(ctx.Param("id"))
	if erratoi != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id should be uint"})
		return
	}
	issue := model.Issue{}
	issue.ID = uint(id)

	if err := issue.LoadWithComments(); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, issue)

}

func GetIssueQuery(ctx *gin.Context) {
	offset, erratoio := strconv.Atoi(ctx.Query("o"))
	limit, erratoil := strconv.Atoi(ctx.Query("l"))

	if erratoio != nil || erratoil != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "query for offset and limit should be uint"})
		return
	}
	var issues []model.Issue
	issues, err := model.IssueGetWithOffsetAndLimit(offset, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, issues)
}
