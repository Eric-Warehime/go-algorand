// Package public provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package public

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
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Broadcasts a raw transaction to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameterWithLocation("simple", false, "txid", runtime.ParamLocationPath, ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
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

	router.GET(baseURL+"/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.POST(baseURL+"/v2/transactions", wrapper.RawTransaction, m...)
	router.GET(baseURL+"/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET(baseURL+"/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

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
