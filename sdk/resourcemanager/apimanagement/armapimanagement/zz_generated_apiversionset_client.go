//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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
	"strconv"
	"strings"
)

// APIVersionSetClient contains the methods for the APIVersionSet group.
// Don't use this type directly, use NewAPIVersionSetClient() instead.
type APIVersionSetClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAPIVersionSetClient creates a new instance of APIVersionSetClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAPIVersionSetClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APIVersionSetClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &APIVersionSetClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or Updates a Api Version Set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// versionSetID - Api Version Set identifier. Must be unique in the current API Management service instance.
// parameters - Create or update parameters.
// options - APIVersionSetClientCreateOrUpdateOptions contains the optional parameters for the APIVersionSetClient.CreateOrUpdate
// method.
func (client *APIVersionSetClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, parameters APIVersionSetContract, options *APIVersionSetClientCreateOrUpdateOptions) (APIVersionSetClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, versionSetID, parameters, options)
	if err != nil {
		return APIVersionSetClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionSetClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return APIVersionSetClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APIVersionSetClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, parameters APIVersionSetContract, options *APIVersionSetClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *APIVersionSetClient) createOrUpdateHandleResponse(resp *http.Response) (APIVersionSetClientCreateOrUpdateResponse, error) {
	result := APIVersionSetClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIVersionSetContract); err != nil {
		return APIVersionSetClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specific Api Version Set.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// versionSetID - Api Version Set identifier. Must be unique in the current API Management service instance.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// options - APIVersionSetClientDeleteOptions contains the optional parameters for the APIVersionSetClient.Delete method.
func (client *APIVersionSetClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, ifMatch string, options *APIVersionSetClientDeleteOptions) (APIVersionSetClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, versionSetID, ifMatch, options)
	if err != nil {
		return APIVersionSetClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionSetClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return APIVersionSetClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return APIVersionSetClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APIVersionSetClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, ifMatch string, options *APIVersionSetClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the details of the Api Version Set specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// versionSetID - Api Version Set identifier. Must be unique in the current API Management service instance.
// options - APIVersionSetClientGetOptions contains the optional parameters for the APIVersionSetClient.Get method.
func (client *APIVersionSetClient) Get(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, options *APIVersionSetClientGetOptions) (APIVersionSetClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, versionSetID, options)
	if err != nil {
		return APIVersionSetClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionSetClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionSetClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APIVersionSetClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, options *APIVersionSetClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIVersionSetClient) getHandleResponse(resp *http.Response) (APIVersionSetClientGetResponse, error) {
	result := APIVersionSetClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIVersionSetContract); err != nil {
		return APIVersionSetClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the Api Version Set specified by its identifier.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// versionSetID - Api Version Set identifier. Must be unique in the current API Management service instance.
// options - APIVersionSetClientGetEntityTagOptions contains the optional parameters for the APIVersionSetClient.GetEntityTag
// method.
func (client *APIVersionSetClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, options *APIVersionSetClientGetEntityTagOptions) (APIVersionSetClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, versionSetID, options)
	if err != nil {
		return APIVersionSetClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionSetClientGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *APIVersionSetClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, options *APIVersionSetClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *APIVersionSetClient) getEntityTagHandleResponse(resp *http.Response) (APIVersionSetClientGetEntityTagResponse, error) {
	result := APIVersionSetClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// NewListByServicePager - Lists a collection of API Version Sets in the specified service instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - APIVersionSetClientListByServiceOptions contains the optional parameters for the APIVersionSetClient.ListByService
// method.
func (client *APIVersionSetClient) NewListByServicePager(resourceGroupName string, serviceName string, options *APIVersionSetClientListByServiceOptions) *runtime.Pager[APIVersionSetClientListByServiceResponse] {
	return runtime.NewPager(runtime.PageProcessor[APIVersionSetClientListByServiceResponse]{
		More: func(page APIVersionSetClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APIVersionSetClientListByServiceResponse) (APIVersionSetClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return APIVersionSetClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return APIVersionSetClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return APIVersionSetClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *APIVersionSetClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *APIVersionSetClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *APIVersionSetClient) listByServiceHandleResponse(resp *http.Response) (APIVersionSetClientListByServiceResponse, error) {
	result := APIVersionSetClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIVersionSetCollection); err != nil {
		return APIVersionSetClientListByServiceResponse{}, err
	}
	return result, nil
}

// Update - Updates the details of the Api VersionSet specified by its identifier.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// versionSetID - Api Version Set identifier. Must be unique in the current API Management service instance.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// parameters - Update parameters.
// options - APIVersionSetClientUpdateOptions contains the optional parameters for the APIVersionSetClient.Update method.
func (client *APIVersionSetClient) Update(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, ifMatch string, parameters APIVersionSetUpdateParameters, options *APIVersionSetClientUpdateOptions) (APIVersionSetClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, versionSetID, ifMatch, parameters, options)
	if err != nil {
		return APIVersionSetClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return APIVersionSetClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return APIVersionSetClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *APIVersionSetClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, versionSetID string, ifMatch string, parameters APIVersionSetUpdateParameters, options *APIVersionSetClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apiVersionSets/{versionSetId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if versionSetID == "" {
		return nil, errors.New("parameter versionSetID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{versionSetId}", url.PathEscape(versionSetID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *APIVersionSetClient) updateHandleResponse(resp *http.Response) (APIVersionSetClientUpdateResponse, error) {
	result := APIVersionSetClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIVersionSetContract); err != nil {
		return APIVersionSetClientUpdateResponse{}, err
	}
	return result, nil
}
