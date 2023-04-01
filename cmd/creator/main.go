package main

import (
	"fmt"
	creator "github.com/SekyrOrg/openApi/gen/creator"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	var myApi openApi
	swagger, err := creator.GetSwagger()
	if err != nil {
		panic(fmt.Errorf("could not get swagger: %w", err))
	}

	r := gin.Default()
	r.Use(middleware.OapiRequestValidator(swagger))
	r = creator.RegisterHandlers(r, myApi)
	r.Run(":8081")
}

var _ creator.ServerInterface = openApi{}

type openApi struct {
}

func (o openApi) GetDistlist(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) GetHealthz(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) PostCreator(c *gin.Context, params creator.PostCreatorParams) {
	//TODO implement me
	panic("implement me")
}
