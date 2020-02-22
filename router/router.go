package router
/**
 * <p> @Description TODO
 * @author weijian
 * @date 2020-02-22 18:18
 */

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.New()
	apiVersionOne := router.Group("/api/v1/")
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
