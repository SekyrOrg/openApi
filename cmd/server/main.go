package main

import (
	"fmt"
	gateway "github.com/SekyrOrg/gateway/gen"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
)

func main() {
	SetupHandler()
}

type server struct {
	gateway.ServerInterface
}

func (s server) GetBeacons(ctx *gin.Context) {
	ctx.JSON(
		200,
		[]gateway.Beacon{
			{
				ID: &types.UUID{},
			},
		},
	)
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
	return r.Run(":8081")
}
