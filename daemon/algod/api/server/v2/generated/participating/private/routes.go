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
	// Get a RoundDeltas object for a given round
	// (GET /v2/deltas/{round})
	GetRoundDeltas(ctx echo.Context, round uint64) error
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

// GetRoundDeltas converts echo context to params.
func (w *ServerInterfaceWrapper) GetRoundDeltas(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameterWithLocation("simple", false, "round", runtime.ParamLocationPath, ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRoundDeltas(ctx, round)
	return err
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

	router.GET(baseURL+"/v2/deltas/:round", wrapper.GetRoundDeltas, m...)
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
	"W9UAAA==",
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
