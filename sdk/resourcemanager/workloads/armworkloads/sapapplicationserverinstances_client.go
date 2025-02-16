//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armworkloads

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SAPApplicationServerInstancesClient contains the methods for the SAPApplicationServerInstances group.
// Don't use this type directly, use NewSAPApplicationServerInstancesClient() instead.
type SAPApplicationServerInstancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSAPApplicationServerInstancesClient creates a new instance of SAPApplicationServerInstancesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSAPApplicationServerInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SAPApplicationServerInstancesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SAPApplicationServerInstancesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Puts the SAP Application Server Instance.
// This will be used by service only. PUT by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// applicationInstanceName - The name of SAP Application Server instance.
// body - The SAP Application Server instance request body.
// options - SAPApplicationServerInstancesClientBeginCreateOptions contains the optional parameters for the SAPApplicationServerInstancesClient.BeginCreate
// method.
func (client *SAPApplicationServerInstancesClient) BeginCreate(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body SAPApplicationServerInstance, options *SAPApplicationServerInstancesClientBeginCreateOptions) (*runtime.Poller[SAPApplicationServerInstancesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SAPApplicationServerInstancesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SAPApplicationServerInstancesClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Puts the SAP Application Server Instance.
// This will be used by service only. PUT by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
func (client *SAPApplicationServerInstancesClient) create(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body SAPApplicationServerInstance, options *SAPApplicationServerInstancesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *SAPApplicationServerInstancesClient) createCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body SAPApplicationServerInstance, options *SAPApplicationServerInstancesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances/{applicationInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if applicationInstanceName == "" {
		return nil, errors.New("parameter applicationInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationInstanceName}", url.PathEscape(applicationInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Deletes the SAP Application Server Instance.
// This operation will be used by service only. Delete by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// applicationInstanceName - The name of SAP Application Server instance.
// options - SAPApplicationServerInstancesClientBeginDeleteOptions contains the optional parameters for the SAPApplicationServerInstancesClient.BeginDelete
// method.
func (client *SAPApplicationServerInstancesClient) BeginDelete(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginDeleteOptions) (*runtime.Poller[SAPApplicationServerInstancesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SAPApplicationServerInstancesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SAPApplicationServerInstancesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the SAP Application Server Instance.
// This operation will be used by service only. Delete by end user will return a Bad Request error.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
func (client *SAPApplicationServerInstancesClient) deleteOperation(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SAPApplicationServerInstancesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances/{applicationInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if applicationInstanceName == "" {
		return nil, errors.New("parameter applicationInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationInstanceName}", url.PathEscape(applicationInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the SAP Application Server Instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// applicationInstanceName - The name of SAP Application Server instance.
// options - SAPApplicationServerInstancesClientGetOptions contains the optional parameters for the SAPApplicationServerInstancesClient.Get
// method.
func (client *SAPApplicationServerInstancesClient) Get(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientGetOptions) (SAPApplicationServerInstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, options)
	if err != nil {
		return SAPApplicationServerInstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SAPApplicationServerInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SAPApplicationServerInstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SAPApplicationServerInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *SAPApplicationServerInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances/{applicationInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if applicationInstanceName == "" {
		return nil, errors.New("parameter applicationInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationInstanceName}", url.PathEscape(applicationInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SAPApplicationServerInstancesClient) getHandleResponse(resp *http.Response) (SAPApplicationServerInstancesClientGetResponse, error) {
	result := SAPApplicationServerInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPApplicationServerInstance); err != nil {
		return SAPApplicationServerInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the SAP Application server Instances in an SVI.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// options - SAPApplicationServerInstancesClientListOptions contains the optional parameters for the SAPApplicationServerInstancesClient.List
// method.
func (client *SAPApplicationServerInstancesClient) NewListPager(resourceGroupName string, sapVirtualInstanceName string, options *SAPApplicationServerInstancesClientListOptions) *runtime.Pager[SAPApplicationServerInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SAPApplicationServerInstancesClientListResponse]{
		More: func(page SAPApplicationServerInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SAPApplicationServerInstancesClientListResponse) (SAPApplicationServerInstancesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SAPApplicationServerInstancesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SAPApplicationServerInstancesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SAPApplicationServerInstancesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SAPApplicationServerInstancesClient) listCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, options *SAPApplicationServerInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SAPApplicationServerInstancesClient) listHandleResponse(resp *http.Response) (SAPApplicationServerInstancesClientListResponse, error) {
	result := SAPApplicationServerInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SAPApplicationServerInstanceList); err != nil {
		return SAPApplicationServerInstancesClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Puts the SAP Application Server Instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// sapVirtualInstanceName - The name of the Virtual Instances for SAP.
// applicationInstanceName - The name of SAP Application Server instance.
// body - The SAP Application Server instance request body.
// options - SAPApplicationServerInstancesClientBeginUpdateOptions contains the optional parameters for the SAPApplicationServerInstancesClient.BeginUpdate
// method.
func (client *SAPApplicationServerInstancesClient) BeginUpdate(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body UpdateSAPApplicationInstanceRequest, options *SAPApplicationServerInstancesClientBeginUpdateOptions) (*runtime.Poller[SAPApplicationServerInstancesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SAPApplicationServerInstancesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SAPApplicationServerInstancesClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Puts the SAP Application Server Instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-12-01-preview
func (client *SAPApplicationServerInstancesClient) update(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body UpdateSAPApplicationInstanceRequest, options *SAPApplicationServerInstancesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sapVirtualInstanceName, applicationInstanceName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SAPApplicationServerInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body UpdateSAPApplicationInstanceRequest, options *SAPApplicationServerInstancesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Workloads/sapVirtualInstances/{sapVirtualInstanceName}/applicationInstances/{applicationInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sapVirtualInstanceName == "" {
		return nil, errors.New("parameter sapVirtualInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sapVirtualInstanceName}", url.PathEscape(sapVirtualInstanceName))
	if applicationInstanceName == "" {
		return nil, errors.New("parameter applicationInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationInstanceName}", url.PathEscape(applicationInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
