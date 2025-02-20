// Package data provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package data

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
	// Removes minimum sync round restriction from the ledger.
	// (DELETE /v2/ledger/sync)
	UnsetSyncRound(ctx echo.Context) error
	// Returns the minimum sync round the ledger is keeping in cache.
	// (GET /v2/ledger/sync)
	GetSyncRound(ctx echo.Context) error
	// Given a round, tells the ledger to keep that round in its cache.
	// (POST /v2/ledger/sync/{round})
	SetSyncRound(ctx echo.Context, round uint64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// UnsetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) UnsetSyncRound(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UnsetSyncRound(ctx)
	return err
}

// GetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) GetSyncRound(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSyncRound(ctx)
	return err
}

// SetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) SetSyncRound(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameterWithLocation("simple", false, "round", runtime.ParamLocationPath, ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SetSyncRound(ctx, round)
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

	router.DELETE(baseURL+"/v2/ledger/sync", wrapper.UnsetSyncRound, m...)
	router.GET(baseURL+"/v2/ledger/sync", wrapper.GetSyncRound, m...)
	router.POST(baseURL+"/v2/ledger/sync/:round", wrapper.SetSyncRound, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AQAA//+92+QcXcsAAA==",
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
