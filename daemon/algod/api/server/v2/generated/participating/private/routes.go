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
