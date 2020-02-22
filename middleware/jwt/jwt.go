package jwt

/**
 * <p> @Description TODO
 * @author weijian
 * @date 2020-02-22 18:15
 */

import (
	_"github.com/detectiveHLH/go-backend-starter/consts"
	_"github.com/detectiveHLH/go-backend-starter/util"
	"github.com/gin-gonic/gin"
	_"github.com/gin-gonic/gin"
	_"net/http"
	_"time"
)

func Jwt() gin.HandlerFunc {
	//return func(c *gin.Context) {
	//	var code int
	//	var data interface{}
	//
	//	code = consts.SUCCESS
	//	token := c.Query("token")
	//	if token == "" {
	//		code = INVALID_PARAMS
	//	} else {
	//		claims, err := util.ParseToken(token)
	//		if err != nil {
	//			code = ERROR_AUTH_CHECK_TOKEN_FAIL
	//		} else if time.Now().Unix() > claims.ExpiresAt {
	//			code = ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	//		}
	//	}
	//
	//	if code != SUCCESS {
	//		c.JSON(http.StatusUnauthorized, gin.H{
	//			"code": code,
	//			"msg":  GetMsg(code),
	//			"data": data,
	//		})
	//
	//		c.Abort()
	//		return
	//	}
	//
	//	c.Next()
	//}
}
