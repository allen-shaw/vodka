package hello

import (
	"github.com/gin-gonic/gin"
	api "github.com/allen-shaw/vodka/demo/internal/model/apihello"
)

type HelloService interface {
	Hello(ctx *gin.Context, req *api.HelloReq) (*api.HelloResp, error)
}

func NewHelloService() HelloService {
	return &helloService{}
}

type helloService struct {
	// TODO: user implement
}

func (s *helloService) Hello(ctx *gin.Context, req *api.HelloReq) (*api.HelloResp, error) {
	// TODO: user implement
	panic("not implement")
}
