package internal

import "github.com/gin-gonic/gin"

func Register(router gin.IRouter, s *HelloService) {
	hello := router.Group("hellp")
	hello.GET("/hello", s.Hello)
}
