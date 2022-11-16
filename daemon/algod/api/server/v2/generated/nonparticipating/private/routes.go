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

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"

	"github.com/algorand/oapi-codegen/pkg/runtime"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameterWithLocation("simple", false, "catchpoint", runtime.ParamLocationPath, ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
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

	router.POST(baseURL+"/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.POST(baseURL+"/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"tWs3n2/+bwAAAP//nLAtUwG3AAA=",
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
