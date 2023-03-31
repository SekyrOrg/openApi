package main

import (
	"fmt"
	server "github.com/SekyrOrg/openApi/gen/server"
	middleware "github.com/deepmap/oapi-codegen/pkg/gin-middleware"
	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
)

func main() {
	var myApi openApi
	swagger, err := server.GetSwagger()
	if err != nil {
		panic(fmt.Errorf("could not get swagger: %w", err))
	}

	r := gin.Default()
	r.Use(middleware.OapiRequestValidator(swagger))
	r = server.RegisterHandlers(r, myApi)
	r.Run(":8081")
}

var _ server.ServerInterface = openApi{}

type openApi struct {
}

func (o openApi) GetBeacons(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) DeleteBeaconsBeaconId(c *gin.Context, beaconId types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) GetBeaconsBeaconId(c *gin.Context, beaconId types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) GetGroups(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) PostGroups(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) DeleteGroupsGroupId(c *gin.Context, groupId types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) GetGroupsGroupId(c *gin.Context, groupId types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) PutGroupsGroupId(c *gin.Context, groupId types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) GetRecipients(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) PostRecipients(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) DeleteRecipientsRecipientId(c *gin.Context, recipientId types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) GetRecipientsRecipientId(c *gin.Context, recipientId types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (o openApi) PutRecipientsRecipientId(c *gin.Context, recipientId types.UUID) {
	//TODO implement me
	panic("implement me")
}
