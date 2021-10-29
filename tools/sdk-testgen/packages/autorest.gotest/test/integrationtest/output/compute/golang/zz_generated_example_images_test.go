//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package golang_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/CreateAnImageFromABlobWithDiskEncryptionSet.json
func ExampleImagesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := golang.NewImagesClient(con,
		"<subscription-id>")
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<image-name>",
		golang.Image{
			Resource: golang.Resource{
				Location: to.StringPtr("<location>"),
			},
			Properties: &golang.ImageProperties{
				StorageProfile: &golang.ImageStorageProfile{
					OSDisk: &golang.ImageOSDisk{
						ImageDisk: golang.ImageDisk{
							BlobURI: to.StringPtr("<blob-uri>"),
							DiskEncryptionSet: &golang.DiskEncryptionSetParameters{
								SubResource: golang.SubResource{
									ID: to.StringPtr("<id>"),
								},
							},
						},
						OSState: golang.OperatingSystemStateTypesGeneralized.ToPtr(),
						OSType:  golang.OperatingSystemTypesLinux.ToPtr(),
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
	log.Printf("Image.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/UpdateImage.json
func ExampleImagesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := golang.NewImagesClient(con,
		"<subscription-id>")
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<image-name>",
		golang.ImageUpdate{
			UpdateResource: golang.UpdateResource{
				Tags: map[string]*string{
					"department": to.StringPtr("HR"),
				},
			},
			Properties: &golang.ImageProperties{
				HyperVGeneration: golang.HyperVGenerationTypesV1.ToPtr(),
				SourceVirtualMachine: &golang.SubResource{
					ID: to.StringPtr("<id>"),
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
	log.Printf("Image.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetInformationAboutAnImage.json
func ExampleImagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := golang.NewImagesClient(con,
		"<subscription-id>")
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<image-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Image.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListImagesInAResourceGroup.json
func ExampleImagesClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := golang.NewImagesClient(con,
		"<subscription-id>")
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Image.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ListImagesInASubscription.json
func ExampleImagesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	con := arm.NewDefaultConnection(cred, nil)
	ctx := context.Background()
	client := golang.NewImagesClient(con,
		"<subscription-id>")
	pager := client.List(nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Image.ID: %s\n", *v.ID)
		}
	}
}