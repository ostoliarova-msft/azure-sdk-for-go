# Azure Communication Identity client library for Go

Azure Communication Identity is managing identities and tokens for Azure Communication Services.

[Product documentation][product_docs]

## Getting started

### Use this command to generate Identity SDK code

autorest --input-file=https://raw.githubusercontent.com/Azure/azure-rest-api-specs/main/specification/communication/data-plane/Identity/stable/2022-10-01/CommunicationIdentity.json --go --license-header=MICROSOFT_APACHE_NO_VERSION --namespace=azcommunicationidentity --output-folder=./azcommunicationidentity --package-version=v1.0.0-beta --clear-output-folder --can-clear-output-folder

### Prerequisites

You need an [Azure subscription][azure_sub] and a [Communication Service Resource][communication_resource_docs] to use this package.

To create a new Communication Service, you can use the [Azure Portal][communication_resource_create_portal], the [Azure PowerShell][communication_resource_create_power_shell], or the [.NET management client library][communication_resource_create_net].

### Authenticate the client

The identity client can be authenticated using endpoint acquired from an Azure Communication Resources in the [Azure Portal][azure_portal] and valid Azure Active Directory token. This approach is using managed identities authentication.

```Go Snippet:Managed_identities_client
// Get the endpoint name value to our Azure Communication resource.
var endpoint = "<endpoint>"

// Get credential
credential, err := azidentity.NewDefaultAzureCredential(nil)

// Create the client that uses managed identities authentication
client := azcommunicationidentity.NewClient(endpoint, credential, nil)
```

Or alternatively using the connection string acquired from an Azure Communication Resource in the [Azure Portal][azure_portal]. This approach is using HMAC authentication.

```Go Snippet:Hmac_client
// Get a connection string to our Azure Communication resource.
connectionString := "<connection_string>"

// Create the client that uses managed identities authentication
client := azcommunicationidentity.NewClientWithHmacAuth(connectionString, nil)
```

### Key concepts

`azcommunicationidentity.Client` provides the functionalities to manage user access tokens: creating new ones and revoking them.

## Examples

### Creating a new user

```Go Snippet:CreateCommunicationUser
// Create a new communication user.
clientCreateResponse, err := client.CreateUser(context.TODO(), nil)
id := *clientCreateResponse.AccessTokenResult.Identity.ID
fmt.Printf("Created Communication user: %s.\n", id)
```

### Getting a token for an existing user

```Go Snippet:CreateCommunicationToken
// Define scopes for the user token creation.
chatScope := azcommunicationidentity.CommunicationIdentityTokenScopeChat
voipScope := azcommunicationidentity.CommunicationIdentityTokenScopeVoip

accessTokenRequest := azcommunicationidentity.AccessTokenRequest{
    Scopes:           []*azcommunicationidentity.CommunicationIdentityTokenScope{&chatScope, &voipScope},
}

// Issue token for the created user.
tokenResponse, err := client.GetToken(context.TODO(), id, accessTokenRequest, nil)
fmt.Printf("Issued a token with the following scopes: %s, %s.\n", chatScope, voipScope)
token := *tokenResponse.AccessToken.Token
expiresOn := *tokenResponse.AccessToken.ExpiresOn
fmt.Printf("Token: %s.\n", token)
fmt.Printf("Expires on: %v.\n", expiresOn.Format(time.DateTime))
```

It's also possible to create a Communication Identity access token by customizing the expiration time. Validity period of the token must be within [1,24] hours range. If not provided, the default value of 24 hours will be used.

```Go Snippet:CreateCommunicationTokenWithCustomExpiration
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
fmt.Printf("Issued a token with the following scopes: %s, %s.\n", chatScope, voipScope)
token := *tokenResponse.AccessToken.Token
expiresOn := *tokenResponse.AccessToken.ExpiresOn
fmt.Printf("Token: %s.\n", token)
fmt.Printf("Expires on: %v.\n", expiresOn.Format(time.DateTime))
```

### Revoking user tokens

We can revoke tokens using following code:

```Go Snippet:RevokeCommunicationUserToken
//Revoke token.
_, err = client.RevokeTokens(context.TODO(), id, nil)
fmt.Print("Revoked a token.\n")
```

### Deleting a user

```Go Snippet:DeleteCommunicationUser
// Delete user.
_, err = client.DeleteUser(context.TODO(), id, nil)
fmt.Print("Deleted a user.\n")
```

## Troubleshooting

All User token service operations will throw an error on failure.

```Go Snippet:CommunicationIdentityClient_Troubleshooting
// Create a new communication user.
clientCreateResponse, err := client.CreateUser(context.TODO(), nil)
if err != nil {
    panic("Cannot create communication user.")
}
```

## Next steps

[Read more about Communication user access tokens][user_access_token]

## Contributing

This project welcomes contributions and suggestions. Most contributions require you to agree to a Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us the rights to use your contribution. For details, visit [cla.microsoft.com][cla].

This project has adopted the [Microsoft Open Source Code of Conduct][coc]. For more information see the [Code of Conduct FAQ][coc_faq] or contact [opencode@microsoft.com][coc_contact] with any additional questions or comments.

<!-- LINKS -->

[azure_sub]: https://azure.microsoft.com/free/dotnet/
[azure_portal]: https://portal.azure.com
[cla]: https://cla.microsoft.com
[coc]: https://opensource.microsoft.com/codeofconduct/
[coc_faq]: https://opensource.microsoft.com/codeofconduct/faq/
[coc_contact]: mailto:opencode@microsoft.com
[product_docs]: https://docs.microsoft.com/azure/communication-services/overview
[user_access_token]: https://learn.microsoft.com/en-us/azure/communication-services/quickstarts/access-tokens?pivots=platform-azcli&tabs=windows
[communication_resource_docs]: https://docs.microsoft.com/azure/communication-services/quickstarts/create-communication-resource?tabs=windows&pivots=platform-azp
[communication_resource_create_portal]: https://docs.microsoft.com/azure/communication-services/quickstarts/create-communication-resource?tabs=windows&pivots=platform-azp
[communication_resource_create_power_shell]: https://docs.microsoft.com/powershell/module/az.communication/new-azcommunicationservice
[communication_resource_create_net]: https://docs.microsoft.com/azure/communication-services/quickstarts/create-communication-resource?tabs=windows&pivots=platform-net
