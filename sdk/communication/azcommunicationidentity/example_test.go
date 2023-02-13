
package azcommunicationidentity_test

import (
	"fmt"
	"context"
    "testing"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/communication/azcommunicationidentity"
)

func TestExampleNewServiceClient(t *testing.T) {
	endpoint := ""

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(err)
	}
	client := azcommunicationidentity.NewClient(endpoint, cred, nil)
	clientCreateResponse, _ := client.Create(context.TODO(), nil)
	fmt.Printf("response>>%#v", clientCreateResponse)
	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	fmt.Printf("\nID>>%#v", id)
	clientDeleteResponse, _ := client.Delete(context.TODO(), id, nil)	
	fmt.Printf("\ndelete response>>%#v", clientDeleteResponse)	
}