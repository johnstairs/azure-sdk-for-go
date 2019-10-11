package databoxapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/databox/mgmt/2019-09-01/databox"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result databox.OperationListPage, err error)
}

var _ OperationsClientAPI = (*databox.OperationsClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	BookShipmentPickUp(ctx context.Context, resourceGroupName string, jobName string, shipmentPickUpRequest databox.ShipmentPickUpRequest) (result databox.ShipmentPickUpResponse, err error)
	Cancel(ctx context.Context, resourceGroupName string, jobName string, cancellationReason databox.CancellationReason) (result autorest.Response, err error)
	Create(ctx context.Context, resourceGroupName string, jobName string, jobResource databox.JobResource) (result databox.JobsCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, jobName string) (result databox.JobsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, jobName string, expand string) (result databox.JobResource, err error)
	List(ctx context.Context, skipToken string) (result databox.JobResourceListPage, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, skipToken string) (result databox.JobResourceListPage, err error)
	ListCredentials(ctx context.Context, resourceGroupName string, jobName string) (result databox.UnencryptedCredentialsList, err error)
	Update(ctx context.Context, resourceGroupName string, jobName string, jobResourceUpdateParameter databox.JobResourceUpdateParameter, ifMatch string) (result databox.JobsUpdateFuture, err error)
}

var _ JobsClientAPI = (*databox.JobsClient)(nil)

// ServiceClientAPI contains the set of methods on the ServiceClient type.
type ServiceClientAPI interface {
	ListAvailableSkus(ctx context.Context, location string, availableSkuRequest databox.AvailableSkuRequest) (result databox.AvailableSkusResultPage, err error)
	ListAvailableSkusByResourceGroup(ctx context.Context, resourceGroupName string, location string, availableSkuRequest databox.AvailableSkuRequest) (result databox.AvailableSkusResultPage, err error)
	RegionConfiguration(ctx context.Context, location string, regionConfigurationRequest databox.RegionConfigurationRequest) (result databox.RegionConfigurationResponse, err error)
	ValidateAddressMethod(ctx context.Context, location string, validateAddress databox.ValidateAddress) (result databox.AddressValidationOutput, err error)
	ValidateInputs(ctx context.Context, location string, validationRequest databox.BasicValidationRequest) (result databox.ValidationResponse, err error)
	ValidateInputsByResourceGroup(ctx context.Context, resourceGroupName string, location string, validationRequest databox.BasicValidationRequest) (result databox.ValidationResponse, err error)
}

var _ ServiceClientAPI = (*databox.ServiceClient)(nil)