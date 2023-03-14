package internal

import (
	"errors"
	"net/http"

	"github.com/allen-shaw/vodka/demo/internal/api/hello"
	api "github.com/allen-shaw/vodka/demo/internal/model/api/hello"
	"github.com/gin-gonic/gin"
)

type HelloService struct {
	server hello.HelloService
}

func NewHelloService() *HelloService {
	return &HelloService{server: hello.NewHelloService()}
}

func (s *HelloService) Error(c *gin.Context, code int, msg string) {
	c.AbortWithError(code, errors.New(msg))
}
func (s *HelloService) ErrorWithResp(c *gin.Context, code int, resp any) {
	c.AbortWithStatusJSON(code, resp)
}
func (s *HelloService) Success(c *gin.Context, resp any) {
	c.JSON(http.StatusOK, resp)
}

func (s *HelloService) Hello(c *gin.Context) {
	var req api.HelloReq

	c.ShouldBindJSON()

	err := c.ShouldBindQuery(&req)
	if err != nil {
		s.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	resp, err := s.server.Hello(c, &req)
	if err != nil {
		s.ErrorWithResp(c, http.StatusInternalServerError, resp)
	}

	s.Success(c, resp)
}
