package routers

import (
	"tutorial/api/article"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/addArticle", article.AddArticle)
	r.GET("/article", article.CheckAnItem)
	r.GET("/listArticle", article.ListNArticles)
	return r
}
