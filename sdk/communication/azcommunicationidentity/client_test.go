package azcommunicationidentity_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/communication/azcommunicationidentity"
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
	"github.com/stretchr/testify/require"
)

func TestNewClientManagedIdentity(t *testing.T) {
	client := createClientWithManagedIdentity()

	// Create a new identity
	clientCreateResponse, err := client.Create(context.TODO(), nil)
	require.Nil(t, err)
	require.NotNil(t, clientCreateResponse)

	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	require.True(t, strings.HasPrefix(id, "8:acs:"))

	// Clean up resources
	clientDeleteResponse, err := client.Delete(context.TODO(), id, nil)
	require.Nil(t, err)
	require.NotNil(t, clientDeleteResponse)
}

func TestGetTokenGeneratesTokenAndIdentityWithScopes(t *testing.T) {
	client := createClientWithManagedIdentity()

	// Create identity
	clientCreateResponse, err := client.Create(context.TODO(), nil)
	require.Nil(t, err)
	require.NotNil(t, clientCreateResponse)

	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	require.True(t, strings.HasPrefix(id, "8:acs:"))

	chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
	voipScope := azcommunicationidentity.CommunicationIdentityTokenScopeVoip
	expirationTime := int32(60)

	accessTokenRequest := azcommunicationidentity.AccessTokenRequest{
		Scopes:           []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope, &voipScope},
		ExpiresInMinutes: &expirationTime,
	}

	// Issue token for the created identity
	tokenResponse, err := client.IssueAccessToken(context.TODO(), id, accessTokenRequest, nil)
	require.Nil(t, err)
	require.NotNil(t, clientCreateResponse)
	require.NotNil(t, tokenResponse.AccessToken.Token)

	token := *tokenResponse.AccessToken.Token
	require.True(t, len(token) > 0)

	// Clean up resources
	client.Delete(context.TODO(), id, nil)
}

func TestRevokeAccessToken(t *testing.T) {
	client := createClientWithManagedIdentity()

	// Create identity
	clientCreateResponse, _ := client.Create(context.TODO(), nil)
	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
	voipScope := azcommunicationidentity.CommunicationIdentityTokenScopeVoip
	expirationTime := int32(60)

	accessTokenRequest := azcommunicationidentity.AccessTokenRequest{
		Scopes:           []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope, &voipScope},
		ExpiresInMinutes: &expirationTime,
	}

	// Issue token for the created identity
	tokenResponse, _ := client.IssueAccessToken(context.TODO(), id, accessTokenRequest, nil)
	token := *tokenResponse.AccessToken.Token
	require.True(t, len(token) > 0)

	revokeResponse, err := client.RevokeAccessTokens(context.TODO(), id, nil)
	require.Nil(t, err)
	require.NotNil(t, revokeResponse)

	// Clean up resources
	client.Delete(context.TODO(), id, nil)
}

func TestExchangeAccessToken(t *testing.T) {
	client := createClientWithManagedIdentity()

	exchangeTokenRequest, err := getTokenExchangeRequest()
	require.Nil(t, err)

	exchangeResponse, err := client.ExchangeTeamsUserAccessToken(context.TODO(), exchangeTokenRequest, nil)
	require.Nil(t, err)
	require.NotNil(t, exchangeResponse)

	token := *exchangeResponse.AccessToken.Token
	require.True(t, len(token) > 0)
}

func parseEndpoint(connectionString string) (string, error) {
	if !strings.HasPrefix(connectionString, "endpoint=") {
		return "", errors.New("connection string should follow the format: endpoint=url;accesskey=key")
	}

	connectionString = strings.TrimPrefix(connectionString, "endpoint=")
	return strings.Split(connectionString, ";")[0], nil
}

func createClientWithManagedIdentity() *azcommunicationidentity.Client {
	connectionString := os.Getenv("AZURE_COMMUNICATION_CONNECTION_STRING")

	endpoint, _ := parseEndpoint(connectionString)

	cred, _ := azidentity.NewDefaultAzureCredential(nil)

	return azcommunicationidentity.NewClient(endpoint, cred, nil)
}

func getTokenExchangeRequest() (azcommunicationidentity.TeamsUserExchangeTokenRequest, error) {
	appId := os.Getenv("COMMUNICATION_M365_APP_ID")
	authority := fmt.Sprintf("%v/%v", os.Getenv("COMMUNICATION_M365_AAD_AUTHORITY"), os.Getenv("COMMUNICATION_M365_AAD_TENANT"))

	publicClientApp, _ := public.New(appId, public.WithAuthority(authority))
	scopes := []string{"https://auth.msft.communication.azure.com/Teams.ManageCalls",
		"https://auth.msft.communication.azure.com/Teams.ManageChats"}

	result, err := publicClientApp.AcquireTokenByUsernamePassword(context.TODO(), scopes, os.Getenv("COMMUNICATION_MSAL_USERNAME"), os.Getenv("COMMUNICATION_MSAL_PASSWORD"))
	if err != nil {
		return azcommunicationidentity.TeamsUserExchangeTokenRequest{}, err
	}

	accountIds := strings.Split(result.Account.HomeAccountID, ".")

	return azcommunicationidentity.TeamsUserExchangeTokenRequest{
		AppID:  &appId,
		Token:  &result.AccessToken,
		UserID: &accountIds[0],
	}, nil
}
