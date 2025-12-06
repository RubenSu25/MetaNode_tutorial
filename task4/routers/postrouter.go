package routers

import "github.com/gin-gonic/gin"

func Postrouters(r *gin.Engine) {

	postrouters := r.Group("/post")
	{
		postrouters.GET("/list", )
	}

}
