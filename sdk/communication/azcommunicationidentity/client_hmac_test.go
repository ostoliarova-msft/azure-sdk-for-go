package azcommunicationidentity_test

import (
	"context"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/communication/azcommunicationidentity"
	"github.com/stretchr/testify/assert"
)

var connectionString = os.Getenv("AZURE_COMMUNICATION_CONNECTION_STRING")

func TestIdentityClientUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	clientCreateResponse, _ := client.CreateUser(context.TODO(), nil)
	assert.NotNil(t, clientCreateResponse)
	assert.NotNil(t, clientCreateResponse.AccessTokenResult.Identity.ID)
}

func TestCreateUserAndTokenUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	var scope = azcommunicationidentity.CommunicationIdentityTokenScopeChat
	body := azcommunicationidentity.CreateRequest{
		Scopes: []*azcommunicationidentity.CommunicationIdentityTokenScope{&scope},
	}
	clientCreateOption := azcommunicationidentity.ClientCreateUserOptions{
		Body: &body,
	}
	clientCreateResponse, _ := client.CreateUser(context.TODO(), &clientCreateOption)
	assert.NotNil(t, clientCreateResponse)
	assert.NotNil(t, clientCreateResponse.AccessTokenResult.Identity.ID)
	assert.NotNil(t, clientCreateResponse.AccessTokenResult.AccessToken.Token)
}

func TestGetTokenUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	clientCreateResponse, _ := client.CreateUser(context.TODO(), nil)
	assert.NotNil(t, clientCreateResponse)
	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	assert.NotNil(t, id)

	chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
	accessTokenRequest := azcommunicationidentity.AccessTokenRequest{
		Scopes: []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope},
	}
	accessTokenResponse, _ := client.GetToken(context.TODO(), id, accessTokenRequest, nil)
	assert.NotNil(t, accessTokenResponse)
	assert.NotNil(t, accessTokenResponse.AccessToken.Token)
	assert.NotNil(t, accessTokenResponse.AccessToken.ExpiresOn)
}

func TestRevokeTokenUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	clientCreateResponse, _ := client.CreateUser(context.TODO(), nil)
	assert.NotNil(t, clientCreateResponse)
	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	assert.NotNil(t, id)

	chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
	accessTokenRequest := azcommunicationidentity.AccessTokenRequest{
		Scopes: []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope},
	}
	accessTokenResponse, _ := client.GetToken(context.TODO(), id, accessTokenRequest, nil)
	assert.NotNil(t, accessTokenResponse)
	assert.NotNil(t, accessTokenResponse.AccessToken.Token)

	revokeAccessTokenResponse, _ := client.RevokeTokens(context.TODO(), id, nil)
	assert.Empty(t, revokeAccessTokenResponse)
}

func TestDeleteUserUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	clientCreateResponse, _ := client.CreateUser(context.TODO(), nil)
	assert.NotNil(t, clientCreateResponse)
	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	assert.NotNil(t, id)

	clientDeleteResponse, _ := client.DeleteUser(context.TODO(), id, nil)
	assert.Empty(t, clientDeleteResponse)
}
