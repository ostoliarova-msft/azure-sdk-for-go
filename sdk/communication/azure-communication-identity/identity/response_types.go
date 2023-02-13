//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package identity

// CommunicationIdentityClientCreateResponse contains the response from method CommunicationIdentityClient.Create.
type CommunicationIdentityClientCreateResponse struct {
	CommunicationIdentityAccessTokenResult
}

// CommunicationIdentityClientDeleteResponse contains the response from method CommunicationIdentityClient.Delete.
type CommunicationIdentityClientDeleteResponse struct {
	// placeholder for future response values
}

// CommunicationIdentityClientExchangeTeamsUserAccessTokenResponse contains the response from method CommunicationIdentityClient.ExchangeTeamsUserAccessToken.
type CommunicationIdentityClientExchangeTeamsUserAccessTokenResponse struct {
	CommunicationIdentityAccessToken
}

// CommunicationIdentityClientIssueAccessTokenResponse contains the response from method CommunicationIdentityClient.IssueAccessToken.
type CommunicationIdentityClientIssueAccessTokenResponse struct {
	CommunicationIdentityAccessToken
}

// CommunicationIdentityClientRevokeAccessTokensResponse contains the response from method CommunicationIdentityClient.RevokeAccessTokens.
type CommunicationIdentityClientRevokeAccessTokensResponse struct {
	// placeholder for future response values
}

