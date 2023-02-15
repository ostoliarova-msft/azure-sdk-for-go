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

// NewClient creates a new instance of Azure Communication Identity Client with the specified values.
//   - endpoint - The communication resource, for example https://my-resource.communication.azure.com
//   - cred - an Azure AD credential, typically obtained via the azidentity module
//   - options - client options; pass nil to accept the default values
func NewClient(endpoint string, credential azcore.TokenCredential, options *ClientOptions) *Client {
	authPolicy := runtime.NewBearerTokenPolicy(credential, []string{"https://communication.azure.com//.default"}, nil)
	var conOptions = getClientOptions(options)
	conOptions.PerRetryPolicies = append(conOptions.PerRetryPolicies, authPolicy)
	pl := runtime.NewPipeline(moduleName, version, runtime.PipelineOptions{}, &conOptions.ClientOptions)
	return &Client{
		endpoint: endpoint,
		pl:       pl,
	}
}

func getClientOptions[T any](o *T) *T {
	if o == nil {
		return new(T)
	}
	return o
}
