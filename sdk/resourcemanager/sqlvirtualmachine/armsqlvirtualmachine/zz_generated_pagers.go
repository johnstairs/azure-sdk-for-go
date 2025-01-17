//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsqlvirtualmachine

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AvailabilityGroupListenersListByGroupPager provides operations for iterating over paged responses.
type AvailabilityGroupListenersListByGroupPager struct {
	client    *AvailabilityGroupListenersClient
	current   AvailabilityGroupListenersListByGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AvailabilityGroupListenersListByGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AvailabilityGroupListenersListByGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AvailabilityGroupListenersListByGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AvailabilityGroupListenerListResult.NextLink == nil || len(*p.current.AvailabilityGroupListenerListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AvailabilityGroupListenersListByGroupResponse page.
func (p *AvailabilityGroupListenersListByGroupPager) PageResponse() AvailabilityGroupListenersListByGroupResponse {
	return p.current
}

// OperationsListPager provides operations for iterating over paged responses.
type OperationsListPager struct {
	client    *OperationsClient
	current   OperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsListResponse page.
func (p *OperationsListPager) PageResponse() OperationsListResponse {
	return p.current
}

// SQLVirtualMachineGroupsListByResourceGroupPager provides operations for iterating over paged responses.
type SQLVirtualMachineGroupsListByResourceGroupPager struct {
	client    *SQLVirtualMachineGroupsClient
	current   SQLVirtualMachineGroupsListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SQLVirtualMachineGroupsListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SQLVirtualMachineGroupsListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SQLVirtualMachineGroupsListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SQLVirtualMachineGroupListResult.NextLink == nil || len(*p.current.SQLVirtualMachineGroupListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SQLVirtualMachineGroupsListByResourceGroupResponse page.
func (p *SQLVirtualMachineGroupsListByResourceGroupPager) PageResponse() SQLVirtualMachineGroupsListByResourceGroupResponse {
	return p.current
}

// SQLVirtualMachineGroupsListPager provides operations for iterating over paged responses.
type SQLVirtualMachineGroupsListPager struct {
	client    *SQLVirtualMachineGroupsClient
	current   SQLVirtualMachineGroupsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SQLVirtualMachineGroupsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SQLVirtualMachineGroupsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SQLVirtualMachineGroupsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SQLVirtualMachineGroupListResult.NextLink == nil || len(*p.current.SQLVirtualMachineGroupListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SQLVirtualMachineGroupsListResponse page.
func (p *SQLVirtualMachineGroupsListPager) PageResponse() SQLVirtualMachineGroupsListResponse {
	return p.current
}

// SQLVirtualMachinesListByResourceGroupPager provides operations for iterating over paged responses.
type SQLVirtualMachinesListByResourceGroupPager struct {
	client    *SQLVirtualMachinesClient
	current   SQLVirtualMachinesListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SQLVirtualMachinesListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SQLVirtualMachinesListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SQLVirtualMachinesListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SQLVirtualMachineListResult.NextLink == nil || len(*p.current.SQLVirtualMachineListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SQLVirtualMachinesListByResourceGroupResponse page.
func (p *SQLVirtualMachinesListByResourceGroupPager) PageResponse() SQLVirtualMachinesListByResourceGroupResponse {
	return p.current
}

// SQLVirtualMachinesListBySQLVMGroupPager provides operations for iterating over paged responses.
type SQLVirtualMachinesListBySQLVMGroupPager struct {
	client    *SQLVirtualMachinesClient
	current   SQLVirtualMachinesListBySQLVMGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SQLVirtualMachinesListBySQLVMGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SQLVirtualMachinesListBySQLVMGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SQLVirtualMachinesListBySQLVMGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SQLVirtualMachineListResult.NextLink == nil || len(*p.current.SQLVirtualMachineListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySQLVMGroupHandleError(resp)
		return false
	}
	result, err := p.client.listBySQLVMGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SQLVirtualMachinesListBySQLVMGroupResponse page.
func (p *SQLVirtualMachinesListBySQLVMGroupPager) PageResponse() SQLVirtualMachinesListBySQLVMGroupResponse {
	return p.current
}

// SQLVirtualMachinesListPager provides operations for iterating over paged responses.
type SQLVirtualMachinesListPager struct {
	client    *SQLVirtualMachinesClient
	current   SQLVirtualMachinesListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SQLVirtualMachinesListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SQLVirtualMachinesListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SQLVirtualMachinesListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SQLVirtualMachineListResult.NextLink == nil || len(*p.current.SQLVirtualMachineListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SQLVirtualMachinesListResponse page.
func (p *SQLVirtualMachinesListPager) PageResponse() SQLVirtualMachinesListResponse {
	return p.current
}
