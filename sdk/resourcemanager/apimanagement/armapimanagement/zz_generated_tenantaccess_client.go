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
	"strings"
)

// TenantAccessClient contains the methods for the TenantAccess group.
// Don't use this type directly, use NewTenantAccessClient() instead.
type TenantAccessClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTenantAccessClient creates a new instance of TenantAccessClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTenantAccessClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TenantAccessClient, error) {
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
	client := &TenantAccessClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Create - Update tenant access information details.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// accessName - The identifier of the Access configuration.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// parameters - Parameters supplied to retrieve the Tenant Access Information.
// options - TenantAccessClientCreateOptions contains the optional parameters for the TenantAccessClient.Create method.
func (client *TenantAccessClient) Create(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, ifMatch string, parameters AccessInformationCreateParameters, options *TenantAccessClientCreateOptions) (TenantAccessClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, serviceName, accessName, ifMatch, parameters, options)
	if err != nil {
		return TenantAccessClientCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TenantAccessClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *TenantAccessClient) createCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, ifMatch string, parameters AccessInformationCreateParameters, options *TenantAccessClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
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
	req.Raw().Header.Set("If-Match", ifMatch)
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *TenantAccessClient) createHandleResponse(resp *http.Response) (TenantAccessClientCreateResponse, error) {
	result := TenantAccessClientCreateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessInformationContract); err != nil {
		return TenantAccessClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Get tenant access information details without secrets.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// accessName - The identifier of the Access configuration.
// options - TenantAccessClientGetOptions contains the optional parameters for the TenantAccessClient.Get method.
func (client *TenantAccessClient) Get(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientGetOptions) (TenantAccessClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TenantAccessClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TenantAccessClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}"
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
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
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
func (client *TenantAccessClient) getHandleResponse(resp *http.Response) (TenantAccessClientGetResponse, error) {
	result := TenantAccessClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessInformationContract); err != nil {
		return TenantAccessClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Tenant access metadata
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// accessName - The identifier of the Access configuration.
// options - TenantAccessClientGetEntityTagOptions contains the optional parameters for the TenantAccessClient.GetEntityTag
// method.
func (client *TenantAccessClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientGetEntityTagOptions) (TenantAccessClientGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessClientGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessClientGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *TenantAccessClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}"
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
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
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
func (client *TenantAccessClient) getEntityTagHandleResponse(resp *http.Response) (TenantAccessClientGetEntityTagResponse, error) {
	result := TenantAccessClientGetEntityTagResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// NewListByServicePager - Returns list of access infos - for Git and Management endpoints.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// options - TenantAccessClientListByServiceOptions contains the optional parameters for the TenantAccessClient.ListByService
// method.
func (client *TenantAccessClient) NewListByServicePager(resourceGroupName string, serviceName string, options *TenantAccessClientListByServiceOptions) *runtime.Pager[TenantAccessClientListByServiceResponse] {
	return runtime.NewPager(runtime.PageProcessor[TenantAccessClientListByServiceResponse]{
		More: func(page TenantAccessClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TenantAccessClientListByServiceResponse) (TenantAccessClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TenantAccessClientListByServiceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return TenantAccessClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TenantAccessClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *TenantAccessClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *TenantAccessClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant"
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
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *TenantAccessClient) listByServiceHandleResponse(resp *http.Response) (TenantAccessClientListByServiceResponse, error) {
	result := TenantAccessClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessInformationCollection); err != nil {
		return TenantAccessClientListByServiceResponse{}, err
	}
	return result, nil
}

// ListSecrets - Get tenant access information details.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// accessName - The identifier of the Access configuration.
// options - TenantAccessClientListSecretsOptions contains the optional parameters for the TenantAccessClient.ListSecrets
// method.
func (client *TenantAccessClient) ListSecrets(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientListSecretsOptions) (TenantAccessClientListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessClientListSecretsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TenantAccessClientListSecretsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *TenantAccessClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/listSecrets"
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
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *TenantAccessClient) listSecretsHandleResponse(resp *http.Response) (TenantAccessClientListSecretsResponse, error) {
	result := TenantAccessClientListSecretsResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessInformationSecretsContract); err != nil {
		return TenantAccessClientListSecretsResponse{}, err
	}
	return result, nil
}

// RegeneratePrimaryKey - Regenerate primary access key
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// accessName - The identifier of the Access configuration.
// options - TenantAccessClientRegeneratePrimaryKeyOptions contains the optional parameters for the TenantAccessClient.RegeneratePrimaryKey
// method.
func (client *TenantAccessClient) RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientRegeneratePrimaryKeyOptions) (TenantAccessClientRegeneratePrimaryKeyResponse, error) {
	req, err := client.regeneratePrimaryKeyCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessClientRegeneratePrimaryKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessClientRegeneratePrimaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return TenantAccessClientRegeneratePrimaryKeyResponse{}, runtime.NewResponseError(resp)
	}
	return TenantAccessClientRegeneratePrimaryKeyResponse{}, nil
}

// regeneratePrimaryKeyCreateRequest creates the RegeneratePrimaryKey request.
func (client *TenantAccessClient) regeneratePrimaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientRegeneratePrimaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/regeneratePrimaryKey"
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
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// RegenerateSecondaryKey - Regenerate secondary access key
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// accessName - The identifier of the Access configuration.
// options - TenantAccessClientRegenerateSecondaryKeyOptions contains the optional parameters for the TenantAccessClient.RegenerateSecondaryKey
// method.
func (client *TenantAccessClient) RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientRegenerateSecondaryKeyOptions) (TenantAccessClientRegenerateSecondaryKeyResponse, error) {
	req, err := client.regenerateSecondaryKeyCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessClientRegenerateSecondaryKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessClientRegenerateSecondaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return TenantAccessClientRegenerateSecondaryKeyResponse{}, runtime.NewResponseError(resp)
	}
	return TenantAccessClientRegenerateSecondaryKeyResponse{}, nil
}

// regenerateSecondaryKeyCreateRequest creates the RegenerateSecondaryKey request.
func (client *TenantAccessClient) regenerateSecondaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientRegenerateSecondaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/regenerateSecondaryKey"
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
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Update - Update tenant access information details.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// accessName - The identifier of the Access configuration.
// ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
// it should be * for unconditional update.
// parameters - Parameters supplied to retrieve the Tenant Access Information.
// options - TenantAccessClientUpdateOptions contains the optional parameters for the TenantAccessClient.Update method.
func (client *TenantAccessClient) Update(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, ifMatch string, parameters AccessInformationUpdateParameters, options *TenantAccessClientUpdateOptions) (TenantAccessClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, accessName, ifMatch, parameters, options)
	if err != nil {
		return TenantAccessClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TenantAccessClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *TenantAccessClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, ifMatch string, parameters AccessInformationUpdateParameters, options *TenantAccessClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
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
func (client *TenantAccessClient) updateHandleResponse(resp *http.Response) (TenantAccessClientUpdateResponse, error) {
	result := TenantAccessClientUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessInformationContract); err != nil {
		return TenantAccessClientUpdateResponse{}, err
	}
	return result, nil
}
