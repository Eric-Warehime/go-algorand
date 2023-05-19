// Package private provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST(baseURL+"/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE(baseURL+"/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET(baseURL+"/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST(baseURL+"/v2/participation/:participation-id", wrapper.AppendKeys, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"wpHC9Yfr/wkAAP//n/TLvvHVAAA=",
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
