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
	// Get a connection string to our Azure Communication resource.
	connectionString := os.Getenv("AZURE_COMMUNICATION_CONNECTION_STRING")
	endpoint, err := parseEndpoint(connectionString)
	if err != nil {
		panic("Invalid endpoint.")
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic("Cannot create defasult credential.")
	}

	client := azcommunicationidentity.NewClient(endpoint, credential, nil)
	fmt.Print("Created Communication client from managed identity. \n")

	// Create a new communication user.
	clientCreateResponse, err := client.CreateUser(context.TODO(), nil)
	if err != nil {
		panic("Cannot create communication user.")
	}

	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	fmt.Printf("Created Communication user: %s.\n", id)

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
		panic("Cannot issue a token.")
	}
	fmt.Printf("Issued a token with the following scopes: %s, %s.\n", chatScope, voipScope)

	token := *tokenResponse.AccessToken.Token
	expiresOn := *tokenResponse.AccessToken.ExpiresOn
	fmt.Printf("Token: %s.\n", token)
	fmt.Printf("Expires on: %v.\n", expiresOn.Format(time.DateTime))

	//Revoke token.
	_, err = client.RevokeTokens(context.TODO(), id, nil)
	if err != nil {
		panic("Cannot revoke a token.")
	}
	fmt.Print("Revoked a token.\n")

	// Clean up resources
	_, err = client.DeleteUser(context.TODO(), id, nil)
	if err != nil {
		panic("Cannot delete a user.")
	}
	fmt.Print("Deleted a user.\n")
}

func parseEndpoint(connectionString string) (string, error) {
	if !strings.HasPrefix(connectionString, "endpoint=") {
		return "", errors.New("connection string should follow the format: endpoint=url;accesskey=key")
	}

	connectionString = strings.TrimPrefix(connectionString, "endpoint=")
	return strings.Split(connectionString, ";")[0], nil
}
