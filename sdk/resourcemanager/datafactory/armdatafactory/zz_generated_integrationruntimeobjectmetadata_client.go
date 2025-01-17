//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// IntegrationRuntimeObjectMetadataClient contains the methods for the IntegrationRuntimeObjectMetadata group.
// Don't use this type directly, use NewIntegrationRuntimeObjectMetadataClient() instead.
type IntegrationRuntimeObjectMetadataClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIntegrationRuntimeObjectMetadataClient creates a new instance of IntegrationRuntimeObjectMetadataClient with the specified values.
func NewIntegrationRuntimeObjectMetadataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IntegrationRuntimeObjectMetadataClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &IntegrationRuntimeObjectMetadataClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get a SSIS integration runtime object metadata by specified path. The return is pageable metadata list.
// If the operation fails it returns the *CloudError error type.
func (client *IntegrationRuntimeObjectMetadataClient) Get(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataGetOptions) (IntegrationRuntimeObjectMetadataGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimeObjectMetadataGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationRuntimeObjectMetadataGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeObjectMetadataGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationRuntimeObjectMetadataClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/getObjectMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.GetMetadataRequest != nil {
		return req, runtime.MarshalAsJSON(req, *options.GetMetadataRequest)
	}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationRuntimeObjectMetadataClient) getHandleResponse(resp *http.Response) (IntegrationRuntimeObjectMetadataGetResponse, error) {
	result := IntegrationRuntimeObjectMetadataGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SsisObjectMetadataListResponse); err != nil {
		return IntegrationRuntimeObjectMetadataGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IntegrationRuntimeObjectMetadataClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginRefresh - Refresh a SSIS integration runtime object metadata.
// If the operation fails it returns the *CloudError error type.
func (client *IntegrationRuntimeObjectMetadataClient) BeginRefresh(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataBeginRefreshOptions) (IntegrationRuntimeObjectMetadataRefreshPollerResponse, error) {
	resp, err := client.refresh(ctx, resourceGroupName, factoryName, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimeObjectMetadataRefreshPollerResponse{}, err
	}
	result := IntegrationRuntimeObjectMetadataRefreshPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IntegrationRuntimeObjectMetadataClient.Refresh", "", resp, client.pl, client.refreshHandleError)
	if err != nil {
		return IntegrationRuntimeObjectMetadataRefreshPollerResponse{}, err
	}
	result.Poller = &IntegrationRuntimeObjectMetadataRefreshPoller{
		pt: pt,
	}
	return result, nil
}

// Refresh - Refresh a SSIS integration runtime object metadata.
// If the operation fails it returns the *CloudError error type.
func (client *IntegrationRuntimeObjectMetadataClient) refresh(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataBeginRefreshOptions) (*http.Response, error) {
	req, err := client.refreshCreateRequest(ctx, resourceGroupName, factoryName, integrationRuntimeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.refreshHandleError(resp)
	}
	return resp, nil
}

// refreshCreateRequest creates the Refresh request.
func (client *IntegrationRuntimeObjectMetadataClient) refreshCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, integrationRuntimeName string, options *IntegrationRuntimeObjectMetadataBeginRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/integrationRuntimes/{integrationRuntimeName}/refreshObjectMetadata"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// refreshHandleError handles the Refresh error response.
func (client *IntegrationRuntimeObjectMetadataClient) refreshHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
