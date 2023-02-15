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
	verifyUser(t, clientCreateResponse)
}

func TestCreateUserAndTokenUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
	clientCreateOption := azcommunicationidentity.ClientCreateUserOptions{
		Body: &azcommunicationidentity.CreateRequest{
			Scopes: []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope},
		},
	}
	clientCreateResponse, _ := client.CreateUser(context.TODO(), &clientCreateOption)
	verifyUser(t, clientCreateResponse)
	verifyAccessToken(t, *clientCreateResponse.AccessToken)
}

func TestCreateUserAndTokenWithCustomExpirationUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
	expirationTime := int32(60)
	clientCreateOption := azcommunicationidentity.ClientCreateUserOptions{
		Body: &azcommunicationidentity.CreateRequest{
			Scopes:           []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope},
			ExpiresInMinutes: &expirationTime,
		},
	}
	clientCreateResponse, _ := client.CreateUser(context.TODO(), &clientCreateOption)
	verifyUser(t, clientCreateResponse)
	verifyAccessToken(t, *clientCreateResponse.AccessToken)
}

func TestGetTokenUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	clientCreateResponse, _ := client.CreateUser(context.TODO(), nil)
	verifyUser(t, clientCreateResponse)
	id := *clientCreateResponse.AccessTokenResult.Identity.ID

	chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
	accessTokenRequest := azcommunicationidentity.AccessTokenRequest{
		Scopes: []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope},
	}
	accessTokenResponse, _ := client.GetToken(context.TODO(), id, accessTokenRequest, nil)
	assert.NotNil(t, accessTokenResponse)
	verifyAccessToken(t, accessTokenResponse.AccessToken)
}

func TestGetTokenWithCustomExpirationUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	clientCreateResponse, _ := client.CreateUser(context.TODO(), nil)
	verifyUser(t, clientCreateResponse)
	id := *clientCreateResponse.AccessTokenResult.Identity.ID

	chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
	expirationTime := int32(60)
	accessTokenRequest := azcommunicationidentity.AccessTokenRequest{
		Scopes:           []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope},
		ExpiresInMinutes: &expirationTime,
	}
	accessTokenResponse, _ := client.GetToken(context.TODO(), id, accessTokenRequest, nil)
	assert.NotNil(t, accessTokenResponse)
	verifyAccessToken(t, accessTokenResponse.AccessToken)
}

func TestRevokeTokenUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	clientCreateResponse, _ := client.CreateUser(context.TODO(), nil)
	verifyUser(t, clientCreateResponse)
	id := *clientCreateResponse.AccessTokenResult.Identity.ID

	chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
	accessTokenRequest := azcommunicationidentity.AccessTokenRequest{
		Scopes: []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope},
	}
	accessTokenResponse, _ := client.GetToken(context.TODO(), id, accessTokenRequest, nil)
	assert.NotNil(t, accessTokenResponse)
	verifyAccessToken(t, accessTokenResponse.AccessToken)

	revokeAccessTokenResponse, _ := client.RevokeTokens(context.TODO(), id, nil)
	assert.Empty(t, revokeAccessTokenResponse)
}

func TestDeleteUserUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	clientCreateResponse, _ := client.CreateUser(context.TODO(), nil)
	verifyUser(t, clientCreateResponse)
	id := *clientCreateResponse.AccessTokenResult.Identity.ID

	clientDeleteResponse, _ := client.DeleteUser(context.TODO(), id, nil)
	assert.Empty(t, clientDeleteResponse)
}

func TestExchangeAccessTokenForTeamsUserUsingConnectionString(t *testing.T) {
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	exchangeTokenRequest, _ := getTokenExchangeRequest()

	tokenExchangeResponse, _ := client.GetTokenForTeamsUser(context.TODO(), exchangeTokenRequest, nil)
	assert.NotNil(t, tokenExchangeResponse)
	verifyAccessToken(t, tokenExchangeResponse.AccessToken)
}

func verifyUser(t *testing.T, clientCreateResponse azcommunicationidentity.ClientCreateUserResponse) {
	assert.NotNil(t, clientCreateResponse)
	assert.NotNil(t, clientCreateResponse.AccessTokenResult.Identity.ID)
	assert.NotEmpty(t, clientCreateResponse.AccessTokenResult.Identity.ID)
}

func verifyAccessToken(t *testing.T, accessToken azcommunicationidentity.AccessToken) {
	assert.NotNil(t, accessToken)
	assert.NotNil(t, accessToken.Token)
	assert.NotEmpty(t, accessToken.Token)
	assert.NotNil(t, accessToken.ExpiresOn)
}
