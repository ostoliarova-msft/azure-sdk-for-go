//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azcommunicationidentity

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ClientOptions contains the optional parameters when creating a Client.
type ClientOptions struct {
	azcore.ClientOptions
}

/*
NewClient creates a new instance of Azure Communication Identity Client with the specified values.
- endpoint - The communication resource, for example https://my-resource.communication.azure.com
- cred - an Azure AD credential, typically obtained via the azidentity module
- options - client options; pass nil to accept the default values
*/
func NewClient(endpoint string, credential azcore.TokenCredential, options *ClientOptions) *Client {
	authPolicy := runtime.NewBearerTokenPolicy(credential, []string{tokenScope}, nil)
	var conOptions = getClientOptions(options)
	conOptions.PerRetryPolicies = append(conOptions.PerRetryPolicies, authPolicy)
	pl := runtime.NewPipeline(moduleName, version, runtime.PipelineOptions{}, &conOptions.ClientOptions)
	return &Client{
		endpoint: endpoint,
		pl:       pl,
	}
}

/*
NewClientWithHmacAuth creates a new instance of Azure Communication Identity Client with the specified values using HMAC Authentication.
  - connectionString - The connection string for setting endpoint and initalizing Client, for example endpoint=https://REDACTED.communication.azure.com/;accesskey=<accessKey>"
  - options - client options; pass nil to accept the default values
*/
func NewClientWithHmacAuth(connectionString string, options *ClientOptions) *Client {
	connectionMap := parseConnectionString(connectionString)
	authPolicy := createCommunicationAccessKeyCredentialPolicy(connectionMap["accesskey"])
	var conOptions = getClientOptions(options)
	conOptions.PerRetryPolicies = append(conOptions.PerRetryPolicies, authPolicy)
	pl := runtime.NewPipeline(moduleName, version, runtime.PipelineOptions{}, &conOptions.ClientOptions)
	return &Client{
		endpoint: connectionMap["endpoint"],
		pl:       pl,
	}
}

func getClientOptions[T any](o *T) *T {
	if o == nil {
		return new(T)
	}
	return o
}
