//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// DeletedWorkspacesClient contains the methods for the DeletedWorkspaces group.
// Don't use this type directly, use NewDeletedWorkspacesClient() instead.
type DeletedWorkspacesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDeletedWorkspacesClient creates a new instance of DeletedWorkspacesClient with the specified values.
func NewDeletedWorkspacesClient(con *arm.Connection, subscriptionID string) *DeletedWorkspacesClient {
	return &DeletedWorkspacesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Gets recently deleted workspaces in a subscription, available for recovery.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DeletedWorkspacesClient) List(ctx context.Context, options *DeletedWorkspacesListOptions) (DeletedWorkspacesListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return DeletedWorkspacesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeletedWorkspacesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeletedWorkspacesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DeletedWorkspacesClient) listCreateRequest(ctx context.Context, options *DeletedWorkspacesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/deletedWorkspaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeletedWorkspacesClient) listHandleResponse(resp *http.Response) (DeletedWorkspacesListResponse, error) {
	result := DeletedWorkspacesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceListResult); err != nil {
		return DeletedWorkspacesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DeletedWorkspacesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Gets recently deleted workspaces in a resource group, available for recovery.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DeletedWorkspacesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *DeletedWorkspacesListByResourceGroupOptions) (DeletedWorkspacesListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return DeletedWorkspacesListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeletedWorkspacesListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeletedWorkspacesListByResourceGroupResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DeletedWorkspacesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DeletedWorkspacesListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/deletedWorkspaces"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DeletedWorkspacesClient) listByResourceGroupHandleResponse(resp *http.Response) (DeletedWorkspacesListByResourceGroupResponse, error) {
	result := DeletedWorkspacesListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceListResult); err != nil {
		return DeletedWorkspacesListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DeletedWorkspacesClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}