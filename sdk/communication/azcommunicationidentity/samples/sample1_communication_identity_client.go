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
	"github.com/AzureAD/microsoft-authentication-library-for-go/apps/public"
)

const managedIdentitySting = "[Managed identity]"
const hmacString = "[HMAC]"

type TokenExchangeBag struct {
	appId        string
	aadAuthority string
	aadTenant    string
	msalUserName string
	msalPassword string
}

func main() {
	connectionString := os.Getenv("AZURE_COMMUNICATION_CONNECTION_STRING")

	appId := os.Getenv("COMMUNICATION_M365_APP_ID")
	aadAuthority := os.Getenv("COMMUNICATION_M365_AAD_AUTHORITY")
	aadTenant := os.Getenv("COMMUNICATION_M365_AAD_TENANT")
	msalUserName := os.Getenv("COMMUNICATION_MSAL_USERNAME")
	msalPassword := os.Getenv("COMMUNICATION_MSAL_PASSWORD")

	tokenExchangeBag := TokenExchangeBag{
		appId:        appId,
		aadAuthority: aadAuthority,
		aadTenant:    aadTenant,
		msalUserName: msalUserName,
		msalPassword: msalPassword,
	}

	// Create the client that uses managed identities authentication
	managedIdentityClient := createManagedIdentityClient(connectionString)
	runSample(*managedIdentityClient, tokenExchangeBag, managedIdentitySting)

	// Create the client that uses HMAC authentication
	hmacClient := createHmacClient(connectionString)
	runSample(*hmacClient, tokenExchangeBag, hmacString)
}

func runSample(client azcommunicationidentity.Client, tokenExchangeBag TokenExchangeBag, authenticationPrefix string) {
	// Create a new communication user.
	clientCreateResponse, err := client.CreateUser(context.TODO(), nil)
	if err != nil {
		panic(fmt.Sprintf("%s Cannot create communication user.", authenticationPrefix))
	}

	id := *clientCreateResponse.AccessTokenResult.Identity.ID
	fmt.Printf("%s Created Communication user: %s.\n", authenticationPrefix, id)

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
		panic(fmt.Sprintf("%s Cannot issue a token.", authenticationPrefix))
	}
	fmt.Printf("%s Issued a token with the following scopes: %s, %s.\n", authenticationPrefix, chatScope, voipScope)

	token := *tokenResponse.AccessToken.Token
	expiresOn := *tokenResponse.AccessToken.ExpiresOn
	fmt.Printf("Token: %s.\n", token)
	fmt.Printf("Expires on: %v.\n", expiresOn.Format(time.DateTime))

	//Revoke token.
	_, err = client.RevokeTokens(context.TODO(), id, nil)
	if err != nil {
		panic(fmt.Sprintf("%s Cannot revoke a token.", authenticationPrefix))
	}
	fmt.Printf("%s Revoked a token.\n", authenticationPrefix)

	// Clean up resources
	_, err = client.DeleteUser(context.TODO(), id, nil)
	if err != nil {
		panic(fmt.Sprintf("%s Cannot delete a user.", authenticationPrefix))
	}
	fmt.Printf("%s Deleted a user.\n", authenticationPrefix)

	// <--- Exchange token part --->
	exchangeTokenRequest, err := getTokenExchangeRequest(tokenExchangeBag)
	fmt.Printf("%s Exchanging token for a Teams user.\n", authenticationPrefix)
	exchangeResponse, err := client.GetTokenForTeamsUser(context.TODO(), exchangeTokenRequest, nil)
	if err != nil {
		panic(fmt.Sprintf("%s Cannot exchange a token.", authenticationPrefix))
	}

	exchangedToken := *exchangeResponse.AccessToken.Token
	fmt.Printf("%s Successfully exchanged token: %s.\n", authenticationPrefix, exchangedToken)
}

func createManagedIdentityClient(connectionString string) *azcommunicationidentity.Client {
	endpoint, err := parseEndpoint(connectionString)
	if err != nil {
		panic(fmt.Sprintf("%s Invalid endpoint.", managedIdentitySting))
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		panic(fmt.Sprintf("%s Cannot create default credential.", managedIdentitySting))
	}

	return azcommunicationidentity.NewClient(endpoint, credential, nil)
}

func createHmacClient(connectionString string) *azcommunicationidentity.Client {
	return azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
}

func parseEndpoint(connectionString string) (string, error) {
	if !strings.HasPrefix(connectionString, "endpoint=") {
		return "", errors.New("connection string should follow the format: endpoint=url;accesskey=key")
	}

	connectionString = strings.TrimPrefix(connectionString, "endpoint=")
	return strings.Split(connectionString, ";")[0], nil
}

func getTokenExchangeRequest(tokenExchangeBag TokenExchangeBag) (azcommunicationidentity.GetTokenForTeamsUserRequest, error) {
	appId := tokenExchangeBag.appId
	authority := fmt.Sprintf("%v/%v", tokenExchangeBag.aadAuthority, tokenExchangeBag.aadTenant)

	publicClientApp, _ := public.New(appId, public.WithAuthority(authority))
	scopes := []string{"https://auth.msft.communication.azure.com/Teams.ManageCalls",
		"https://auth.msft.communication.azure.com/Teams.ManageChats"}

	result, err := publicClientApp.AcquireTokenByUsernamePassword(context.TODO(), scopes, tokenExchangeBag.msalUserName, tokenExchangeBag.msalPassword)
	if err != nil {
		return azcommunicationidentity.GetTokenForTeamsUserRequest{}, err
	}

	accountIds := strings.Split(result.Account.HomeAccountID, ".")

	return azcommunicationidentity.GetTokenForTeamsUserRequest{
		ClientID:          &appId,
		TeamsUserAADToken: &result.AccessToken,
		UserObjectID:      &accountIds[0],
	}, nil
}
