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
	// Get a LedgerStateDelta object for a given transaction group
	// (GET /v2/deltas/txn/group/{id})
	GetLedgerStateDeltaForTransactionGroup(ctx echo.Context, id string, params GetLedgerStateDeltaForTransactionGroupParams) error
	// Get a LedgerStateDelta object for a given round
	// (GET /v2/deltas/{round})
	GetLedgerStateDelta(ctx echo.Context, round uint64, params GetLedgerStateDeltaParams) error
	// Get LedgerStateDelta objects for all transaction groups in a given round
	// (GET /v2/deltas/{round}/txn/group)
	GetTransactionGroupLedgerStateDeltasForRound(ctx echo.Context, round uint64, params GetTransactionGroupLedgerStateDeltasForRoundParams) error
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

// GetLedgerStateDeltaForTransactionGroup converts echo context to params.
func (w *ServerInterfaceWrapper) GetLedgerStateDeltaForTransactionGroup(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLedgerStateDeltaForTransactionGroupParams
	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetLedgerStateDeltaForTransactionGroup(ctx, id, params)
	return err
}

// GetLedgerStateDelta converts echo context to params.
func (w *ServerInterfaceWrapper) GetLedgerStateDelta(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameterWithLocation("simple", false, "round", runtime.ParamLocationPath, ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLedgerStateDeltaParams
	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetLedgerStateDelta(ctx, round, params)
	return err
}

// GetTransactionGroupLedgerStateDeltasForRound converts echo context to params.
func (w *ServerInterfaceWrapper) GetTransactionGroupLedgerStateDeltasForRound(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameterWithLocation("simple", false, "round", runtime.ParamLocationPath, ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTransactionGroupLedgerStateDeltasForRoundParams
	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTransactionGroupLedgerStateDeltasForRound(ctx, round, params)
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

	router.GET(baseURL+"/v2/deltas/txn/group/:id", wrapper.GetLedgerStateDeltaForTransactionGroup, m...)
	router.GET(baseURL+"/v2/deltas/:round", wrapper.GetLedgerStateDelta, m...)
	router.GET(baseURL+"/v2/deltas/:round/txn/group", wrapper.GetTransactionGroupLedgerStateDeltasForRound, m...)
	router.DELETE(baseURL+"/v2/ledger/sync", wrapper.UnsetSyncRound, m...)
	router.GET(baseURL+"/v2/ledger/sync", wrapper.GetSyncRound, m...)
	router.POST(baseURL+"/v2/ledger/sync/:round", wrapper.SetSyncRound, m...)

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
