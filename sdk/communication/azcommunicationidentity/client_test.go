package azcommunicationidentity_test

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/communication/azcommunicationidentity"
	"github.com/stretchr/testify/require"
)

func TestNewClientManagedIdentity(t *testing.T) {
	connectionString := os.Getenv("AZURE_COMMUNICATION_CONNECTION_STRING")

	endpoint, err := parseEndpoint(connectionString)
	require.Nil(t, err)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	require.Nil(t, err)

	client := azcommunicationidentity.NewClient(endpoint, cred, nil)
	clientCreateResponse, err := client.Create(context.TODO(), nil)
	require.Nil(t, err)
	require.NotNil(t, clientCreateResponse)

	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	require.True(t, strings.HasPrefix(id, "8:acs:"))

	clientDeleteResponse, err := client.Delete(context.TODO(), id, nil)
	require.Nil(t, err)
	require.NotNil(t, clientDeleteResponse)
}

func parseEndpoint(connectionString string) (string, error) {
	if !strings.HasPrefix(connectionString, "endpoint=") {
		return "", errors.New("connection string should follow the format: endpoint=url;accesskey=key")
	}

	connectionString = strings.TrimPrefix(connectionString, "endpoint=")
	return strings.Split(connectionString, ";")[0], nil
}
