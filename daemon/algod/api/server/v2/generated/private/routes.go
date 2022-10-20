// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Aborts a catchpoint catchup.
	// (DELETE /v2/catchup/{catchpoint})
	AbortCatchup(ctx echo.Context, catchpoint string) error
	// Starts a catchpoint catchup.
	// (POST /v2/catchup/{catchpoint})
	StartCatchup(ctx echo.Context, catchpoint string) error
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

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AbortCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) AbortCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AbortCatchup(ctx, catchpoint)
	return err
}

// StartCatchup converts echo context to params.
func (w *ServerInterfaceWrapper) StartCatchup(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "catchpoint" -------------
	var catchpoint string

	err = runtime.BindStyledParameter("simple", false, "catchpoint", ctx.Param("catchpoint"), &catchpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter catchpoint: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.StartCatchup(ctx, catchpoint)
	return err
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameter("simple", false, "participation-id", ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty":  true,
		"timeout": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE("/v2/catchup/:catchpoint", wrapper.AbortCatchup, m...)
	router.POST("/v2/catchup/:catchpoint", wrapper.StartCatchup, m...)
	router.GET("/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST("/v2/participation", wrapper.AddParticipationKey, m...)
	router.DELETE("/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET("/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST("/v2/participation/:participation-id", wrapper.AppendKeys, m...)
	router.POST("/v2/shutdown", wrapper.ShutdownNode, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
