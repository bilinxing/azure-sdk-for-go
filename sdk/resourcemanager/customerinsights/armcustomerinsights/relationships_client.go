//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcustomerinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RelationshipsClient contains the methods for the Relationships group.
// Don't use this type directly, use NewRelationshipsClient() instead.
type RelationshipsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRelationshipsClient creates a new instance of RelationshipsClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRelationshipsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RelationshipsClient, error) {
	cl, err := arm.NewClient(moduleName+".RelationshipsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RelationshipsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a relationship or updates an existing relationship within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - relationshipName - The name of the Relationship.
//   - parameters - Parameters supplied to the CreateOrUpdate Relationship operation.
//   - options - RelationshipsClientBeginCreateOrUpdateOptions contains the optional parameters for the RelationshipsClient.BeginCreateOrUpdate
//     method.
func (client *RelationshipsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, relationshipName string, parameters RelationshipResourceFormat, options *RelationshipsClientBeginCreateOrUpdateOptions) (*runtime.Poller[RelationshipsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, hubName, relationshipName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RelationshipsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[RelationshipsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a relationship or updates an existing relationship within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
func (client *RelationshipsClient) createOrUpdate(ctx context.Context, resourceGroupName string, hubName string, relationshipName string, parameters RelationshipResourceFormat, options *RelationshipsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hubName, relationshipName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RelationshipsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, relationshipName string, parameters RelationshipResourceFormat, options *RelationshipsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if relationshipName == "" {
		return nil, errors.New("parameter relationshipName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationshipName}", url.PathEscape(relationshipName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a relationship within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - relationshipName - The name of the relationship.
//   - options - RelationshipsClientBeginDeleteOptions contains the optional parameters for the RelationshipsClient.BeginDelete
//     method.
func (client *RelationshipsClient) BeginDelete(ctx context.Context, resourceGroupName string, hubName string, relationshipName string, options *RelationshipsClientBeginDeleteOptions) (*runtime.Poller[RelationshipsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, hubName, relationshipName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[RelationshipsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[RelationshipsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a relationship within a hub.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
func (client *RelationshipsClient) deleteOperation(ctx context.Context, resourceGroupName string, hubName string, relationshipName string, options *RelationshipsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hubName, relationshipName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RelationshipsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hubName string, relationshipName string, options *RelationshipsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if relationshipName == "" {
		return nil, errors.New("parameter relationshipName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationshipName}", url.PathEscape(relationshipName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets information about the specified relationship.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - relationshipName - The name of the relationship.
//   - options - RelationshipsClientGetOptions contains the optional parameters for the RelationshipsClient.Get method.
func (client *RelationshipsClient) Get(ctx context.Context, resourceGroupName string, hubName string, relationshipName string, options *RelationshipsClientGetOptions) (RelationshipsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hubName, relationshipName, options)
	if err != nil {
		return RelationshipsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RelationshipsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RelationshipsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RelationshipsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hubName string, relationshipName string, options *RelationshipsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships/{relationshipName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if relationshipName == "" {
		return nil, errors.New("parameter relationshipName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationshipName}", url.PathEscape(relationshipName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RelationshipsClient) getHandleResponse(resp *http.Response) (RelationshipsClientGetResponse, error) {
	result := RelationshipsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RelationshipResourceFormat); err != nil {
		return RelationshipsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByHubPager - Gets all relationships in the hub.
//
// Generated from API version 2017-04-26
//   - resourceGroupName - The name of the resource group.
//   - hubName - The name of the hub.
//   - options - RelationshipsClientListByHubOptions contains the optional parameters for the RelationshipsClient.NewListByHubPager
//     method.
func (client *RelationshipsClient) NewListByHubPager(resourceGroupName string, hubName string, options *RelationshipsClientListByHubOptions) *runtime.Pager[RelationshipsClientListByHubResponse] {
	return runtime.NewPager(runtime.PagingHandler[RelationshipsClientListByHubResponse]{
		More: func(page RelationshipsClientListByHubResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RelationshipsClientListByHubResponse) (RelationshipsClientListByHubResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByHubCreateRequest(ctx, resourceGroupName, hubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RelationshipsClientListByHubResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RelationshipsClientListByHubResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RelationshipsClientListByHubResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByHubHandleResponse(resp)
		},
	})
}

// listByHubCreateRequest creates the ListByHub request.
func (client *RelationshipsClient) listByHubCreateRequest(ctx context.Context, resourceGroupName string, hubName string, options *RelationshipsClientListByHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}/relationships"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHubHandleResponse handles the ListByHub response.
func (client *RelationshipsClient) listByHubHandleResponse(resp *http.Response) (RelationshipsClientListByHubResponse, error) {
	result := RelationshipsClientListByHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RelationshipListResult); err != nil {
		return RelationshipsClientListByHubResponse{}, err
	}
	return result, nil
}