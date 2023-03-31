package main

import (
	"fmt"
	gateway "github.com/SekyrOrg/openApi/gen/gateway"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	url2 "net/url"
)

func main() {
	SetupHandler()
}

type server struct {
	gateway.ServerInterface
}

func (s server) GetBeacons(ctx *gin.Context) {
	targetUrl, err := url2.Parse("http://localhost:8081")
	if err != nil {
		panic(err)
	}
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(ctx.Writer, ctx.Request)
}

func SetupHandler() error {
	var myApi server
	swagger, err := gateway.GetSwagger()
	if err != nil {
		return fmt.Errorf("could not get swagger: %w", err)
	}

	r := gin.Default()
	r.Use(middleware.OapiRequestValidator(swagger))
	r = gateway.RegisterHandlers(r, myApi)
	return r.Run(":8080")
}
