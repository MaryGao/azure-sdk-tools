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

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/examples/GetDiskRestorePointResources.json
func ExampleDiskRestorePointClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := golang.NewDiskRestorePointClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<restore-point-collection-name>",
		"<vm-restore-point-name>",
		"<disk-restore-point-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DiskRestorePoint.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2020-12-01/examples/ListDiskRestorePointsInVmRestorePoint.json
func ExampleDiskRestorePointClient_ListByRestorePoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := golang.NewDiskRestorePointClient("<subscription-id>", cred, nil)
	pager := client.ListByRestorePoint("<resource-group-name>",
		"<restore-point-collection-name>",
		"<vm-restore-point-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("DiskRestorePoint.ID: %s\n", *v.ID)
		}
	}
}