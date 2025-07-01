package part4

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func P4Run() {
	r := gin.Default()

	// register cross domain handler globally
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},                         // 允许跨域发来请求的网站
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool { // 自定义过滤源站的方法
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/registry", Register)
	r.POST("/login", Login)

	// 登录验证
	r.Use(JWTAuth())

	r.POST("/post/create", postCreate)
	r.POST("/post/queryById", postRead)
	r.POST("/post/updateById", postUpdate)
	r.POST("/post/deleteById", postDelete)
	r.POST("/comment/create", CommentCreate)
	r.GET("/comment/query", CommentsQuery)

	// support any type of Method, eg. POST、GET、PUT...
	r.Any("/any", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	r.Run(":9091")
}
