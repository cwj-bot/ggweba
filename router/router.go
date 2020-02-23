package router
/**
 * <p> @Description TODO
 * @author weijian
 * @date 2020-02-22 18:18
 */

import (
	_ "github.com/detectiveHLH/go-backend-starter/docs"
	"github.com/detectiveHLH/go-backend-starter/middleware/jwt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)


func InitRouter() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/login", Login)


	apiVersionOne := router.Group("/api/v1/")
	apiVersionOne.Use(jwt.Jwt())

	apiVersionOne.GET("hello1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"code": 200,
			"message": "hello1",
			"data": nil,
		})
	})
	apiVersionOne.GET("hello2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
			"code": 200,
			"message": "hello2",
			"data": nil,
		})
	})
	return router
}
