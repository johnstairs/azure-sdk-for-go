// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// DiagnosticSettingsClient contains the methods for the DiagnosticSettings group.
// Don't use this type directly, use NewDiagnosticSettingsClient() instead.
type DiagnosticSettingsClient struct {
	con *armcore.Connection
}

// NewDiagnosticSettingsClient creates a new instance of DiagnosticSettingsClient with the specified values.
func NewDiagnosticSettingsClient(con *armcore.Connection) *DiagnosticSettingsClient {
	return &DiagnosticSettingsClient{con: con}
}

// CreateOrUpdate - Creates or updates diagnostic settings for the specified resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DiagnosticSettingsClient) CreateOrUpdate(ctx context.Context, resourceURI string, name string, parameters DiagnosticSettingsResource, options *DiagnosticSettingsCreateOrUpdateOptions) (DiagnosticSettingsResourceResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceURI, name, parameters, options)
	if err != nil {
		return DiagnosticSettingsResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DiagnosticSettingsResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DiagnosticSettingsResourceResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DiagnosticSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceURI string, name string, parameters DiagnosticSettingsResource, options *DiagnosticSettingsCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DiagnosticSettingsClient) createOrUpdateHandleResponse(resp *azcore.Response) (DiagnosticSettingsResourceResponse, error) {
	var val *DiagnosticSettingsResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiagnosticSettingsResourceResponse{}, err
	}
	return DiagnosticSettingsResourceResponse{RawResponse: resp.Response, DiagnosticSettingsResource: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DiagnosticSettingsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes existing diagnostic settings for the specified resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DiagnosticSettingsClient) Delete(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceURI, name, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DiagnosticSettingsClient) deleteCreateRequest(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DiagnosticSettingsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the active diagnostic settings for the specified resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DiagnosticSettingsClient) Get(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsGetOptions) (DiagnosticSettingsResourceResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceURI, name, options)
	if err != nil {
		return DiagnosticSettingsResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DiagnosticSettingsResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DiagnosticSettingsResourceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiagnosticSettingsClient) getCreateRequest(ctx context.Context, resourceURI string, name string, options *DiagnosticSettingsGetOptions) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiagnosticSettingsClient) getHandleResponse(resp *azcore.Response) (DiagnosticSettingsResourceResponse, error) {
	var val *DiagnosticSettingsResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiagnosticSettingsResourceResponse{}, err
	}
	return DiagnosticSettingsResourceResponse{RawResponse: resp.Response, DiagnosticSettingsResource: val}, nil
}

// getHandleError handles the Get error response.
func (client *DiagnosticSettingsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets the active diagnostic settings list for the specified resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DiagnosticSettingsClient) List(ctx context.Context, resourceURI string, options *DiagnosticSettingsListOptions) (DiagnosticSettingsResourceCollectionResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return DiagnosticSettingsResourceCollectionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DiagnosticSettingsResourceCollectionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DiagnosticSettingsResourceCollectionResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DiagnosticSettingsClient) listCreateRequest(ctx context.Context, resourceURI string, options *DiagnosticSettingsListOptions) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/diagnosticSettings"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2017-05-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiagnosticSettingsClient) listHandleResponse(resp *azcore.Response) (DiagnosticSettingsResourceCollectionResponse, error) {
	var val *DiagnosticSettingsResourceCollection
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DiagnosticSettingsResourceCollectionResponse{}, err
	}
	return DiagnosticSettingsResourceCollectionResponse{RawResponse: resp.Response, DiagnosticSettingsResourceCollection: val}, nil
}

// listHandleError handles the List error response.
func (client *DiagnosticSettingsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}