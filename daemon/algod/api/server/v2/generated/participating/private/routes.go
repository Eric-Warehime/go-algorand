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
	// Removes minimum sync round restriction from the ledger.
	// (DELETE /v2/ledger/sync)
	UnsetSyncRound(ctx echo.Context) error
	// Returns the minimum sync round the ledger is keeping in cache.
	// (GET /v2/ledger/sync)
	GetSyncRound(ctx echo.Context) error
	// Given a round, tells the ledger to keep that round in its cache.
	// (POST /v2/ledger/sync/{round})
	SetSyncRound(ctx echo.Context, round uint64) error
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

	router.DELETE(baseURL+"/v2/ledger/sync", wrapper.UnsetSyncRound, m...)
	router.GET(baseURL+"/v2/ledger/sync", wrapper.GetSyncRound, m...)
	router.POST(baseURL+"/v2/ledger/sync/:round", wrapper.SetSyncRound, m...)
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
