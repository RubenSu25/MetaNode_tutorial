package middlewares

import (
	"task4/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {

				logger.Log.WithFields(logrus.Fields{
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
					"err":    r,
					"ip":     c.ClientIP(),
				}).Error("服务器内部异常")

				c.AbortWithStatusJSON(500, gin.H{
					"code":    10000,
					"message": "服务器内部错误",
				})

				// 中止请求处理
				c.Abort()
			}
		}()

		c.Next()
	}

}
