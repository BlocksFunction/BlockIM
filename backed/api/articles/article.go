package articles

import (
	"Backed/database"
	"Backed/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	errorKey       = "error"
	internalError  = "服务器内部错误，请稍后再试"
	duplicateError = "该名称已被占用，请换一个吧!"
)

func GetList(c *gin.Context) {
	articlesDB, err := database.UseArticleData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: internalError})
		return
	}
	defer articlesDB.Close()

	articles, err := articlesDB.ListArticles(10000, 0, "", "", false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: internalError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"articles": articles})
}

func AddArticle(c *gin.Context) {
	articlesDB, err := database.UseArticleData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: internalError})
		return
	}
	defer articlesDB.Close()

	claims := utils.GetClaims(c)
	if claims == nil { // 实际上是不会发生的
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无法获取用户信息"})
		return
	}

	var formData database.Article
	formData.Title = c.PostForm("title")
	formData.Excerpt = c.PostForm("excerpt")
	formData.Author = claims.Subject
	formData.ReadTime, err = strconv.Atoi(c.PostForm("read_time"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: "输入格式错误，请重试"})
		return
	}
	formData.Likes, err = strconv.Atoi(c.PostForm("likes"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: "输入格式错误，请重试"})
		return
	}
	formData.Comments, err = strconv.Atoi(c.PostForm("comments"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: "输入格式错误，请重试"})
		return
	}
	formData.Views, err = strconv.Atoi(c.PostForm("views"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: "输入格式错误，请重试"})
		return
	}
	formData.Category = c.PostForm("category")
	tags := c.PostForm("tags")
	formData.Tags = strings.Split(tags, ",")
	formData.Featured, err = strconv.ParseBool(c.PostForm("featured"))
	article, err := articlesDB.InsertArticle(&formData)
	if err != nil {
		if articlesDB.IsDuplicateError(err) {
			c.JSON(http.StatusInternalServerError, gin.H{errorKey: duplicateError})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{errorKey: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "id": article.ID})
}

func DeleteArticle(c *gin.Context) {
	articlesDB, err := database.UseArticleData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
		return
	}
	defer articlesDB.Close()

	claims := utils.GetClaims(c)
	if claims == nil { // 实际上是不会发生的
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无法获取用户信息"})
		return
	}
	username := claims.Subject

	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: "输入格式错误，请重试"})
		return
	}

	var article *database.Article
	article, err = articlesDB.GetArticleByID(int64(articleID))
	if article == nil {
		c.JSON(http.StatusUnauthorized, gin.H{errorKey: "该文章不存在"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: internalError})
		return
	}

	if article.Author == username {
		_, err = articlesDB.DeleteArticle(int64(articleID))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{errorKey: internalError})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "这篇文章不是你的哦"})
	}
}

func GetArticle(c *gin.Context) {
	articlesDB, err := database.UseArticleData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: internalError})
	}
	defer articlesDB.Close()

	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: internalError})
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: "输入格式错误，请重试"})
		return
	}

	var article *database.Article
	article, err = articlesDB.GetArticleByID(int64(articleID))
	if article == nil {
		c.JSON(http.StatusUnauthorized, gin.H{errorKey: "该文章不存在"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{errorKey: internalError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"article": article})
}
