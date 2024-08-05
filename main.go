package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/msi/armmsi"
)

func main() {
	// 1. Create a DefaultAzureCredential

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating credential: %v\n", err)
		os.Exit(1)
	}

	// 2. Create a Managed Identities Client

	clientFactory, err := armmsi.NewClientFactory("6185011d-554a-41cc-9a78-ae6d0a808a54", cred, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create client: %v", err)
	}
	pager := clientFactory.NewUserAssignedIdentitiesClient().NewListBySubscriptionPager(nil)
	ctx := context.Background()
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
			fmt.Printf("FederatedIdentityCredentialsClient: %v\n", *v.Name)

		}

	}
}
