//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// AccessReviewInstanceMyDecisionsClient contains the methods for the AccessReviewInstanceMyDecisions group.
// Don't use this type directly, use NewAccessReviewInstanceMyDecisionsClient() instead.
type AccessReviewInstanceMyDecisionsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewAccessReviewInstanceMyDecisionsClient creates a new instance of AccessReviewInstanceMyDecisionsClient with the specified values.
func NewAccessReviewInstanceMyDecisionsClient(con *arm.Connection) *AccessReviewInstanceMyDecisionsClient {
	return &AccessReviewInstanceMyDecisionsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// GetByID - Get my single access review instance decision.
// If the operation fails it returns the *ErrorDefinition error type.
func (client *AccessReviewInstanceMyDecisionsClient) GetByID(ctx context.Context, scheduleDefinitionID string, id string, decisionID string, options *AccessReviewInstanceMyDecisionsGetByIDOptions) (AccessReviewInstanceMyDecisionsGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, scheduleDefinitionID, id, decisionID, options)
	if err != nil {
		return AccessReviewInstanceMyDecisionsGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewInstanceMyDecisionsGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessReviewInstanceMyDecisionsGetByIDResponse{}, client.getByIDHandleError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *AccessReviewInstanceMyDecisionsClient) getByIDCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, decisionID string, options *AccessReviewInstanceMyDecisionsGetByIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/decisions/{decisionId}"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	if decisionID == "" {
		return nil, errors.New("parameter decisionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{decisionId}", url.PathEscape(decisionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *AccessReviewInstanceMyDecisionsClient) getByIDHandleResponse(resp *http.Response) (AccessReviewInstanceMyDecisionsGetByIDResponse, error) {
	result := AccessReviewInstanceMyDecisionsGetByIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewDecision); err != nil {
		return AccessReviewInstanceMyDecisionsGetByIDResponse{}, err
	}
	return result, nil
}

// getByIDHandleError handles the GetByID error response.
func (client *AccessReviewInstanceMyDecisionsClient) getByIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDefinition{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Get my access review instance decisions.
// If the operation fails it returns the *ErrorDefinition error type.
func (client *AccessReviewInstanceMyDecisionsClient) List(scheduleDefinitionID string, id string, options *AccessReviewInstanceMyDecisionsListOptions) *AccessReviewInstanceMyDecisionsListPager {
	return &AccessReviewInstanceMyDecisionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, scheduleDefinitionID, id, options)
		},
		advancer: func(ctx context.Context, resp AccessReviewInstanceMyDecisionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.AccessReviewDecisionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *AccessReviewInstanceMyDecisionsClient) listCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *AccessReviewInstanceMyDecisionsListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/decisions"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessReviewInstanceMyDecisionsClient) listHandleResponse(resp *http.Response) (AccessReviewInstanceMyDecisionsListResponse, error) {
	result := AccessReviewInstanceMyDecisionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewDecisionListResult); err != nil {
		return AccessReviewInstanceMyDecisionsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *AccessReviewInstanceMyDecisionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDefinition{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Patch - Record a decision.
// If the operation fails it returns the *ErrorDefinition error type.
func (client *AccessReviewInstanceMyDecisionsClient) Patch(ctx context.Context, scheduleDefinitionID string, id string, decisionID string, properties AccessReviewDecisionProperties, options *AccessReviewInstanceMyDecisionsPatchOptions) (AccessReviewInstanceMyDecisionsPatchResponse, error) {
	req, err := client.patchCreateRequest(ctx, scheduleDefinitionID, id, decisionID, properties, options)
	if err != nil {
		return AccessReviewInstanceMyDecisionsPatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AccessReviewInstanceMyDecisionsPatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AccessReviewInstanceMyDecisionsPatchResponse{}, client.patchHandleError(resp)
	}
	return client.patchHandleResponse(resp)
}

// patchCreateRequest creates the Patch request.
func (client *AccessReviewInstanceMyDecisionsClient) patchCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, decisionID string, properties AccessReviewDecisionProperties, options *AccessReviewInstanceMyDecisionsPatchOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/decisions/{decisionId}"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	if decisionID == "" {
		return nil, errors.New("parameter decisionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{decisionId}", url.PathEscape(decisionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, properties)
}

// patchHandleResponse handles the Patch response.
func (client *AccessReviewInstanceMyDecisionsClient) patchHandleResponse(resp *http.Response) (AccessReviewInstanceMyDecisionsPatchResponse, error) {
	result := AccessReviewInstanceMyDecisionsPatchResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewDecision); err != nil {
		return AccessReviewInstanceMyDecisionsPatchResponse{}, err
	}
	return result, nil
}

// patchHandleError handles the Patch error response.
func (client *AccessReviewInstanceMyDecisionsClient) patchHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorDefinition{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}