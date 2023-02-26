package internal

import "github.com/gin-gonic/gin"

func Register(router gin.IRouter, s *HelloService) {
	router.GET("/hello", s.Hello)
}
