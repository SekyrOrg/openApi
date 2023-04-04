// Package creator provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package creator

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Dist defines model for Dist.
type Dist struct {
	Arch *string `json:"arch,omitempty"`
	Os   *string `json:"os,omitempty"`
}

// Arch defines model for Arch.
type Arch = string

// BeaconUUID defines model for BeaconUUID.
type BeaconUUID = openapi_types.UUID

// Debug defines model for Debug.
type Debug = bool

// GroupUUID defines model for GroupUUID.
type GroupUUID = openapi_types.UUID

// Gzip defines model for Gzip.
type Gzip = bool

// Lldflags defines model for Lldflags.
type Lldflags = string

// OS defines model for OS.
type OS = string

// ReportAddr defines model for ReportAddr.
type ReportAddr = string

// Static defines model for Static.
type Static = bool

// Transport defines model for Transport.
type Transport = string

// Upx defines model for Upx.
type Upx = bool

// UpxLevel defines model for UpxLevel.
type UpxLevel = int

// Error defines model for Error.
type Error struct {
	// Code A unique error code for the specific type of error
	Code string `json:"code"`

	// Message A human-readable error message describing the error
	Message string `json:"message"`
}

// PostCreatorParams defines parameters for PostCreator.
type PostCreatorParams struct {
	// ReportAddr The URL of the report server.
	ReportAddr ReportAddr `form:"report_addr" json:"report_addr"`

	// Os The operating system of the beacon.
	Os OS `form:"os" json:"os"`

	// Arch The architecture of the beacon.
	Arch Arch `form:"arch" json:"arch"`

	// BeaconUuid The UUID of the beacon.
	BeaconUuid *BeaconUUID `form:"beacon_uuid,omitempty" json:"beacon_uuid,omitempty"`

	// GroupUuid The UUID of the group.
	GroupUuid *GroupUUID `form:"group_uuid,omitempty" json:"group_uuid,omitempty"`

	// Static Indicates if the beacon is static.
	Static *Static `form:"static,omitempty" json:"static,omitempty"`

	// Upx Indicates if the beacon is compressed using UPX.
	Upx *Upx `form:"upx,omitempty" json:"upx,omitempty"`

	// UpxLevel The compression level used by UPX.
	UpxLevel *UpxLevel `form:"upx_level,omitempty" json:"upx_level,omitempty"`

	// Gzip Indicates if the beacon is compressed using Gzip.
	Gzip *Gzip `form:"gzip,omitempty" json:"gzip,omitempty"`

	// Debug Include debug information in the beacon
	Debug *Debug `form:"debug,omitempty" json:"debug,omitempty"`

	// Lldflags The lldflags used to build the beacon.
	Lldflags *Lldflags `form:"lldflags,omitempty" json:"lldflags,omitempty"`

	// Transport The transport protocol used by the beacon.
	Transport *Transport `form:"transport,omitempty" json:"transport,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new beacon.
	// (POST /creator)
	PostCreator(c *gin.Context, params PostCreatorParams)
	// List all supported OS and Arch combinations
	// (GET /creator/distlist)
	GetCreatorDistlist(c *gin.Context)
	// Check the health of the server.
	// (GET /healthz)
	GetHealthz(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// PostCreator operation middleware
func (siw *ServerInterfaceWrapper) PostCreator(c *gin.Context) {

	var err error

	c.Set(BearerAuthScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PostCreatorParams

	// ------------- Required query parameter "report_addr" -------------

	if paramValue := c.Query("report_addr"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument report_addr is required, but not found: %s", err), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "report_addr", c.Request.URL.Query(), &params.ReportAddr)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter report_addr: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "os" -------------

	if paramValue := c.Query("os"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument os is required, but not found: %s", err), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "os", c.Request.URL.Query(), &params.Os)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter os: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "arch" -------------

	if paramValue := c.Query("arch"); paramValue != "" {

	} else {
		siw.ErrorHandler(c, fmt.Errorf("Query argument arch is required, but not found: %s", err), http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "arch", c.Request.URL.Query(), &params.Arch)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter arch: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "beacon_uuid" -------------

	err = runtime.BindQueryParameter("form", true, false, "beacon_uuid", c.Request.URL.Query(), &params.BeaconUuid)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter beacon_uuid: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "group_uuid" -------------

	err = runtime.BindQueryParameter("form", true, false, "group_uuid", c.Request.URL.Query(), &params.GroupUuid)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter group_uuid: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "static" -------------

	err = runtime.BindQueryParameter("form", true, false, "static", c.Request.URL.Query(), &params.Static)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter static: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "upx" -------------

	err = runtime.BindQueryParameter("form", true, false, "upx", c.Request.URL.Query(), &params.Upx)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter upx: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "upx_level" -------------

	err = runtime.BindQueryParameter("form", true, false, "upx_level", c.Request.URL.Query(), &params.UpxLevel)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter upx_level: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "gzip" -------------

	err = runtime.BindQueryParameter("form", true, false, "gzip", c.Request.URL.Query(), &params.Gzip)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter gzip: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "debug" -------------

	err = runtime.BindQueryParameter("form", true, false, "debug", c.Request.URL.Query(), &params.Debug)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter debug: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "lldflags" -------------

	err = runtime.BindQueryParameter("form", true, false, "lldflags", c.Request.URL.Query(), &params.Lldflags)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter lldflags: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "transport" -------------

	err = runtime.BindQueryParameter("form", true, false, "transport", c.Request.URL.Query(), &params.Transport)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter transport: %s", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.PostCreator(c, params)
}

// GetCreatorDistlist operation middleware
func (siw *ServerInterfaceWrapper) GetCreatorDistlist(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{""})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetCreatorDistlist(c)
}

// GetHealthz operation middleware
func (siw *ServerInterfaceWrapper) GetHealthz(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
	}

	siw.Handler.GetHealthz(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router *gin.Engine, si ServerInterface) *gin.Engine {
	return RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router *gin.Engine, si ServerInterface, options GinServerOptions) *gin.Engine {

	errorHandler := options.ErrorHandler

	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.POST(options.BaseURL+"/creator", wrapper.PostCreator)

	router.GET(options.BaseURL+"/creator/distlist", wrapper.GetCreatorDistlist)

	router.GET(options.BaseURL+"/healthz", wrapper.GetHealthz)

	return router
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xXXW/bNhT9KwS3Rzl2Usdr/eYuXec1QIOmQQsEQUCJ1xY7iWT5kcYN/N+HS1KWHcup",
	"iqB7MkweXh6e+6kHWqhaKwnSWTp9oJoZVoMDE/7NTFHiLwdbGKGdUJJO6ccSCDNFKRwUzhsgakFcCSQH",
	"Vih5RDMqEPbVg1nRjEpWA51SPEEzauCrFwY4nTrjIaO2KKFmeAncs1pXCL1/ObmdjGlG3Urjf+uMkEu6",
	"Xmf0dbjj6mp+1s0Ld/rxibu33gtOu2kcn7yA8enkjwG8fJUPjk/4iwEbn04G45PJ5PR0PB6NRiOa0YUy",
	"NXN0SpOpfdJnkPvlPt+5LCrPgXDcJkJGQ0JJIuTWAw7wD6d2mHNYMF+5RtnNQ+LfRCtXqgImA6+3Rnnd",
	"T8slQg9JGTb/FyXffhe6S0guCubAErHteSIsweg2YC1w4q2QS4IWDr4DrT9L0fOKLyq2tN2CVmmXeOTj",
	"FMm9qHiPWG0OdpOjA0sG3+gWP1o6pzsFfH/ZTU1pMMyhQHZlHdT9ckjZnhn9SUiuvtlOSh9AK+NmnJsD",
	"YfjhvGFjApRYMHdgDpGKoFuGBvuxQ7Wmw2FaOCpUPYxGOvleOuZE8VNRaMORQ4TjbrdvF6yyvSLvo2HS",
	"Bs6dIrpmm2ijnCpUFYMwX/Vw8+bwgfjj0vaLvit9/7zsvbr4fIik1/fPlPBK35/DHVTdCjZcsD5XCNsI",
	"+DSp2wDupna8RevVhpOQDpZg6BpJGbBaSQuhpLwxRoUsKZR0IIOvmdYV6ieUHH6xSPdh6yptMLOdiOcL",
	"xWH/dTPipfjqgQCaJwgiC2WCL6yGQixEQZAbpmHA7HhbyDtWCX6LmQa2I2cyWoO1bNl5delrJgcGGGd5",
	"1VBIeBLBOXoeyezfPY93k3Q32RpfuuKvLQbXUYuW2s0Gr/IvULio/iO2MvFTReEN2kFMFDvoeyas21ed",
	"pSnqxwNOhhW1d+F8zDejFgpvhFtdIqV0uRbvYDXzLlAIMVoC42DaIP08mF3MB+9g1d4STyGhHJgB05yP",
	"//5q+vQ/nz42gZ1GKhMMJyuhDgQdcbjZ9/6b+Ewyu5iHgKuZZEt0dqwA4dXCBR3i3IdImtE7MDZaOD4a",
	"HY2Cbhok04JO6YuwlFHNXBkUGBYGmIt5o1V0UGp3Ss45ndILZd2fCZTtjMDXD/R3Aws6pb8N20F52EKG",
	"W91rnf0Q/f6yDyqM3T1wW8NwD3Q77fUApybXA4k1vR8sFtc+THEO64GLc3UP4GYs64FtG+n6JlYMsO61",
	"4qsn6q4qHLiBdQZYvVt/NyNtLiQzq448Xj+eUR4X/ZPR6NddvZuQrOm9IWWAk4VRdewDPq+Fw6VkLZxN",
	"baxb080jhrFtbReokFnbpeX6BtW2vq7R+JSGdATCiIRvW9OJC7P1NW2y9QaNNgk+5MK6KpXgJXQk+lto",
	"8vysgf6U2Pv9VTio7SEFUmcYhrbQFmxmDFt19heClLDLWq8xBIGT95eESU6wJOAIguoj2v5qB5wjE1ZV",
	"fagc8ksJrHLl96fc8XeCPNMNuw0XB2pvd2fUeNFqZ3po137cXPechTNh/BDBSbVMptYZPY3cn0JL5doT",
	"6524L6H4N2Rc3G++fNpPnkbqSB6VXm8W98ar1FhDguw21tT7G4dhYTxwFhmFD8NIP/JqLSQe65v1fwEA",
	"AP//y9Xu60wSAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
