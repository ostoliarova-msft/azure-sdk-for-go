package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/communication/azcommunicationidentity"
)

func main() {
	runManagedIdentitySample()
	runHmacSample()
}

// This sample is using managed identities authentication
func runManagedIdentitySample() {
	// Get a connection string to our Azure Communication resource.
	connectionString := os.Getenv("AZURE_COMMUNICATION_CONNECTION_STRING")
	endpoint, err := parseEndpoint(connectionString)
	if err != nil {
		panic("[Managed identity] Invalid endpoint.")
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic("[Managed identity] Cannot create defasult credential.")
	}

	// Create the client that uses managed identities authentication
	client := azcommunicationidentity.NewClient(endpoint, credential, nil)
	fmt.Print("Created Communication client from managed identity. \n")

	// Create a new communication user.
	clientCreateResponse, err := client.CreateUser(context.TODO(), nil)
	if err != nil {
		panic("[Managed identity] Cannot create communication user.")
	}

	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	fmt.Printf("[Managed identity] Created Communication user: %s.\n", id)

	// Define scopes for the user token creation.
	chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
	voipScope := azcommunicationidentity.CommunicationIdentityTokenScopeVoip

	// Set token expiration time.
	expirationTime := int32(60)

	accessTokenRequest := azcommunicationidentity.AccessTokenRequest{
		Scopes:           []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope, &voipScope},
		ExpiresInMinutes: &expirationTime,
	}

	// Issue token for the created user.
	tokenResponse, err := client.GetToken(context.TODO(), id, accessTokenRequest, nil)
	if err != nil {
		panic("[Managed identity] Cannot issue a token.")
	}
	fmt.Printf("[Managed identity] Issued a token with the following scopes: %s, %s.\n", chatScope, voipScope)

	token := *tokenResponse.AccessToken.Token
	expiresOn := *tokenResponse.AccessToken.ExpiresOn
	fmt.Printf("Token: %s.\n", token)
	fmt.Printf("Expires on: %v.\n", expiresOn.Format(time.DateTime))

	//Revoke token.
	_, err = client.RevokeTokens(context.TODO(), id, nil)
	if err != nil {
		panic("[Managed identity] Cannot revoke a token.")
	}
	fmt.Print("[Managed identity] Revoked a token.\n")

	// Clean up resources
	_, err = client.DeleteUser(context.TODO(), id, nil)
	if err != nil {
		panic("[Managed identity] Cannot delete a user.")
	}
	fmt.Print("[Managed identity] Deleted a user.\n")
}

// This sample is using HMAC authentication
func runHmacSample() {
	// Get a connection string to our Azure Communication resource.
	connectionString := os.Getenv("AZURE_COMMUNICATION_CONNECTION_STRING")

	// Create the client that uses managed identities authentication
	client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
	fmt.Print("[HMAC] Created Communication client that uses HMAC authentication. \n")

	// Create a new communication user.
	clientCreateResponse, err := client.CreateUser(context.TODO(), nil)
	if err != nil {
		panic("[HMAC] Cannot create communication user.")
	}

	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	fmt.Printf("[HMAC] Created Communication user: %s.\n", id)

	// Define scopes for the user token creation.
	chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
	voipScope := azcommunicationidentity.CommunicationIdentityTokenScopeVoip

	// Set token expiration time.
	expirationTime := int32(60)

	accessTokenRequest := azcommunicationidentity.AccessTokenRequest{
		Scopes:           []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope, &voipScope},
		ExpiresInMinutes: &expirationTime,
	}

	// Issue token for the created user.
	tokenResponse, err := client.GetToken(context.TODO(), id, accessTokenRequest, nil)
	if err != nil {
		panic("[HMAC] Cannot issue a token.")
	}
	fmt.Printf("[HMAC] Issued a token with the following scopes: %s, %s.\n", chatScope, voipScope)

	token := *tokenResponse.AccessToken.Token
	expiresOn := *tokenResponse.AccessToken.ExpiresOn
	fmt.Printf("Token: %s.\n", token)
	fmt.Printf("Expires on: %v.\n", expiresOn.Format(time.DateTime))

	//Revoke token.
	_, err = client.RevokeTokens(context.TODO(), id, nil)
	if err != nil {
		panic("[HMAC] Cannot revoke a token.")
	}
	fmt.Print("[HMAC] Revoked a token.\n")

	// Clean up resources
	_, err = client.DeleteUser(context.TODO(), id, nil)
	if err != nil {
		panic("[HMAC] Cannot delete a user.")
	}
	fmt.Print("[HMAC] Deleted a user.\n")
}

func parseEndpoint(connectionString string) (string, error) {
	if !strings.HasPrefix(connectionString, "endpoint=") {
		return "", errors.New("connection string should follow the format: endpoint=url;accesskey=key")
	}

	connectionString = strings.TrimPrefix(connectionString, "endpoint=")
	return strings.Split(connectionString, ";")[0], nil
}
