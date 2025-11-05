package routers

import (
	"tutorial/api/article"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))
	r.POST("/addArticle", article.AddArticle)
	r.GET("/article", article.CheckAnItem)
	r.GET("/listArticle", article.ListNArticles)
	r.PATCH("/editArticle", article.EditArticle)
	return r
}
