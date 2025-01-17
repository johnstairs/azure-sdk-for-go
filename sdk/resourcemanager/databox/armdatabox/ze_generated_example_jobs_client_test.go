//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabox_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databox/armdatabox"
)

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsList.json
func ExampleJobsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	pager := client.List(&armdatabox.JobsListOptions{SkipToken: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("JobResource.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/MarkDevicesShipped.json
func ExampleJobsClient_MarkDevicesShipped() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	_, err = client.MarkDevicesShipped(ctx,
		"<job-name>",
		"<resource-group-name>",
		armdatabox.MarkDevicesShippedRequest{
			DeliverToDcPackageDetails: &armdatabox.PackageCarrierInfo{
				CarrierName: to.StringPtr("<carrier-name>"),
				TrackingID:  to.StringPtr("<tracking-id>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsListByResourceGroup.json
func ExampleJobsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		&armdatabox.JobsListByResourceGroupOptions{SkipToken: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("JobResource.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsGet.json
func ExampleJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<job-name>",
		&armdatabox.JobsGetOptions{Expand: to.StringPtr("<expand>")})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("JobResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsCreate.json
func ExampleJobsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<job-name>",
		armdatabox.JobResource{
			Resource: armdatabox.Resource{
				Location: to.StringPtr("<location>"),
				SKU: &armdatabox.SKU{
					Name: armdatabox.SKUNameDataBox.ToPtr(),
				},
			},
			Properties: &armdatabox.JobProperties{
				TransferType: armdatabox.TransferTypeImportToAzure.ToPtr(),
				Details: &armdatabox.DataBoxJobDetails{
					JobDetails: armdatabox.JobDetails{
						ContactDetails: &armdatabox.ContactDetails{
							ContactName: to.StringPtr("<contact-name>"),
							EmailList: []*string{
								to.StringPtr("testing@microsoft.com")},
							Phone:          to.StringPtr("<phone>"),
							PhoneExtension: to.StringPtr("<phone-extension>"),
						},
						DataImportDetails: []*armdatabox.DataImportDetails{
							{
								AccountDetails: &armdatabox.StorageAccountDetails{
									DataAccountDetails: armdatabox.DataAccountDetails{
										DataAccountType: armdatabox.DataAccountTypeStorageAccount.ToPtr(),
									},
									StorageAccountID: to.StringPtr("<storage-account-id>"),
								},
							}},
						JobDetailsType: armdatabox.ClassDiscriminatorDataBox.ToPtr(),
						ShippingAddress: &armdatabox.ShippingAddress{
							AddressType:     armdatabox.AddressTypeCommercial.ToPtr(),
							City:            to.StringPtr("<city>"),
							CompanyName:     to.StringPtr("<company-name>"),
							Country:         to.StringPtr("<country>"),
							PostalCode:      to.StringPtr("<postal-code>"),
							StateOrProvince: to.StringPtr("<state-or-province>"),
							StreetAddress1:  to.StringPtr("<street-address1>"),
							StreetAddress2:  to.StringPtr("<street-address2>"),
						},
					},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("JobResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsDelete.json
func ExampleJobsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<job-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsPatch.json
func ExampleJobsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<job-name>",
		armdatabox.JobResourceUpdateParameter{
			Properties: &armdatabox.UpdateJobProperties{
				Details: &armdatabox.UpdateJobDetails{
					ContactDetails: &armdatabox.ContactDetails{
						ContactName: to.StringPtr("<contact-name>"),
						EmailList: []*string{
							to.StringPtr("testing@microsoft.com")},
						Phone:          to.StringPtr("<phone>"),
						PhoneExtension: to.StringPtr("<phone-extension>"),
					},
					ShippingAddress: &armdatabox.ShippingAddress{
						AddressType:     armdatabox.AddressTypeCommercial.ToPtr(),
						City:            to.StringPtr("<city>"),
						CompanyName:     to.StringPtr("<company-name>"),
						Country:         to.StringPtr("<country>"),
						PostalCode:      to.StringPtr("<postal-code>"),
						StateOrProvince: to.StringPtr("<state-or-province>"),
						StreetAddress1:  to.StringPtr("<street-address1>"),
						StreetAddress2:  to.StringPtr("<street-address2>"),
					},
				},
			},
		},
		&armdatabox.JobsBeginUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("JobResource.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/BookShipmentPickupPost.json
func ExampleJobsClient_BookShipmentPickUp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	_, err = client.BookShipmentPickUp(ctx,
		"<resource-group-name>",
		"<job-name>",
		armdatabox.ShipmentPickUpRequest{
			EndTime:          to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-22T18:30:00Z"); return t }()),
			ShipmentLocation: to.StringPtr("<shipment-location>"),
			StartTime:        to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-09-20T18:30:00Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsCancelPost.json
func ExampleJobsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	_, err = client.Cancel(ctx,
		"<resource-group-name>",
		"<job-name>",
		armdatabox.CancellationReason{
			Reason: to.StringPtr("<reason>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsListCredentials.json
func ExampleJobsClient_ListCredentials() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatabox.NewJobsClient("<subscription-id>", cred, nil)
	_, err = client.ListCredentials(ctx,
		"<resource-group-name>",
		"<job-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
