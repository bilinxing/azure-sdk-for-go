//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armorbital

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

// SpacecraftsClient contains the methods for the Spacecrafts group.
// Don't use this type directly, use NewSpacecraftsClient() instead.
type SpacecraftsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSpacecraftsClient creates a new instance of SpacecraftsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSpacecraftsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SpacecraftsClient, error) {
	cl, err := arm.NewClient(moduleName+".SpacecraftsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SpacecraftsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a spacecraft resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - spacecraftName - Spacecraft ID.
//   - parameters - The parameters to provide for the created spacecraft.
//   - options - SpacecraftsClientBeginCreateOrUpdateOptions contains the optional parameters for the SpacecraftsClient.BeginCreateOrUpdate
//     method.
func (client *SpacecraftsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, spacecraftName string, parameters Spacecraft, options *SpacecraftsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SpacecraftsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, spacecraftName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SpacecraftsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SpacecraftsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a spacecraft resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *SpacecraftsClient) createOrUpdate(ctx context.Context, resourceGroupName string, spacecraftName string, parameters Spacecraft, options *SpacecraftsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, spacecraftName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SpacecraftsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, parameters Spacecraft, options *SpacecraftsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a specified spacecraft resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - spacecraftName - Spacecraft ID.
//   - options - SpacecraftsClientBeginDeleteOptions contains the optional parameters for the SpacecraftsClient.BeginDelete method.
func (client *SpacecraftsClient) BeginDelete(ctx context.Context, resourceGroupName string, spacecraftName string, options *SpacecraftsClientBeginDeleteOptions) (*runtime.Poller[SpacecraftsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, spacecraftName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SpacecraftsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SpacecraftsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a specified spacecraft resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *SpacecraftsClient) deleteOperation(ctx context.Context, resourceGroupName string, spacecraftName string, options *SpacecraftsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, spacecraftName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SpacecraftsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, options *SpacecraftsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified spacecraft in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - spacecraftName - Spacecraft ID.
//   - options - SpacecraftsClientGetOptions contains the optional parameters for the SpacecraftsClient.Get method.
func (client *SpacecraftsClient) Get(ctx context.Context, resourceGroupName string, spacecraftName string, options *SpacecraftsClientGetOptions) (SpacecraftsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, spacecraftName, options)
	if err != nil {
		return SpacecraftsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SpacecraftsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SpacecraftsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SpacecraftsClient) getCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, options *SpacecraftsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SpacecraftsClient) getHandleResponse(resp *http.Response) (SpacecraftsClientGetResponse, error) {
	result := SpacecraftsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Spacecraft); err != nil {
		return SpacecraftsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns list of spacecrafts by resource group.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - SpacecraftsClientListOptions contains the optional parameters for the SpacecraftsClient.NewListPager method.
func (client *SpacecraftsClient) NewListPager(resourceGroupName string, options *SpacecraftsClientListOptions) *runtime.Pager[SpacecraftsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SpacecraftsClientListResponse]{
		More: func(page SpacecraftsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SpacecraftsClientListResponse) (SpacecraftsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SpacecraftsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SpacecraftsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SpacecraftsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SpacecraftsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *SpacecraftsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SpacecraftsClient) listHandleResponse(resp *http.Response) (SpacecraftsClientListResponse, error) {
	result := SpacecraftsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpacecraftListResult); err != nil {
		return SpacecraftsClientListResponse{}, err
	}
	return result, nil
}

// BeginListAvailableContacts - Returns list of available contacts. A contact is available if the spacecraft is visible from
// the ground station for more than the minimum viable contact duration provided in the contact profile.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - spacecraftName - Spacecraft ID.
//   - parameters - The parameters to provide for the contacts.
//   - options - SpacecraftsClientBeginListAvailableContactsOptions contains the optional parameters for the SpacecraftsClient.BeginListAvailableContacts
//     method.
func (client *SpacecraftsClient) BeginListAvailableContacts(ctx context.Context, resourceGroupName string, spacecraftName string, parameters ContactParameters, options *SpacecraftsClientBeginListAvailableContactsOptions) (*runtime.Poller[*runtime.Pager[SpacecraftsClientListAvailableContactsResponse]], error) {
	pager := runtime.NewPager(runtime.PagingHandler[SpacecraftsClientListAvailableContactsResponse]{
		More: func(page SpacecraftsClientListAvailableContactsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SpacecraftsClientListAvailableContactsResponse) (SpacecraftsClientListAvailableContactsResponse, error) {
			req, err := runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			if err != nil {
				return SpacecraftsClientListAvailableContactsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SpacecraftsClientListAvailableContactsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SpacecraftsClientListAvailableContactsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAvailableContactsHandleResponse(resp)
		},
	})
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listAvailableContacts(ctx, resourceGroupName, spacecraftName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[*runtime.Pager[SpacecraftsClientListAvailableContactsResponse]]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Response:      &pager,
		})
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[*runtime.Pager[SpacecraftsClientListAvailableContactsResponse]]{
			Response: &pager,
		})
	}
}

// ListAvailableContacts - Returns list of available contacts. A contact is available if the spacecraft is visible from the
// ground station for more than the minimum viable contact duration provided in the contact profile.
//
// Generated from API version 2022-11-01
func (client *SpacecraftsClient) listAvailableContacts(ctx context.Context, resourceGroupName string, spacecraftName string, parameters ContactParameters, options *SpacecraftsClientBeginListAvailableContactsOptions) (*http.Response, error) {
	req, err := client.listAvailableContactsCreateRequest(ctx, resourceGroupName, spacecraftName, parameters, options)
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

// listAvailableContactsCreateRequest creates the ListAvailableContacts request.
func (client *SpacecraftsClient) listAvailableContactsCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, parameters ContactParameters, options *SpacecraftsClientBeginListAvailableContactsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}/listAvailableContacts"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// listAvailableContactsHandleResponse handles the ListAvailableContacts response.
func (client *SpacecraftsClient) listAvailableContactsHandleResponse(resp *http.Response) (SpacecraftsClientListAvailableContactsResponse, error) {
	result := SpacecraftsClientListAvailableContactsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableContactsListResult); err != nil {
		return SpacecraftsClientListAvailableContactsResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Returns list of spacecrafts by subscription.
//
// Generated from API version 2022-11-01
//   - options - SpacecraftsClientListBySubscriptionOptions contains the optional parameters for the SpacecraftsClient.NewListBySubscriptionPager
//     method.
func (client *SpacecraftsClient) NewListBySubscriptionPager(options *SpacecraftsClientListBySubscriptionOptions) *runtime.Pager[SpacecraftsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[SpacecraftsClientListBySubscriptionResponse]{
		More: func(page SpacecraftsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SpacecraftsClientListBySubscriptionResponse) (SpacecraftsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SpacecraftsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SpacecraftsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SpacecraftsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SpacecraftsClient) listBySubscriptionCreateRequest(ctx context.Context, options *SpacecraftsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Orbital/spacecrafts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SpacecraftsClient) listBySubscriptionHandleResponse(resp *http.Response) (SpacecraftsClientListBySubscriptionResponse, error) {
	result := SpacecraftsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SpacecraftListResult); err != nil {
		return SpacecraftsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdateTags - Updates the specified spacecraft tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - spacecraftName - Spacecraft ID.
//   - parameters - Parameters supplied to update spacecraft tags.
//   - options - SpacecraftsClientBeginUpdateTagsOptions contains the optional parameters for the SpacecraftsClient.BeginUpdateTags
//     method.
func (client *SpacecraftsClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, spacecraftName string, parameters TagsObject, options *SpacecraftsClientBeginUpdateTagsOptions) (*runtime.Poller[SpacecraftsClientUpdateTagsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTags(ctx, resourceGroupName, spacecraftName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SpacecraftsClientUpdateTagsResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SpacecraftsClientUpdateTagsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateTags - Updates the specified spacecraft tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01
func (client *SpacecraftsClient) updateTags(ctx context.Context, resourceGroupName string, spacecraftName string, parameters TagsObject, options *SpacecraftsClientBeginUpdateTagsOptions) (*http.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, spacecraftName, parameters, options)
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

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SpacecraftsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, spacecraftName string, parameters TagsObject, options *SpacecraftsClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Orbital/spacecrafts/{spacecraftName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if spacecraftName == "" {
		return nil, errors.New("parameter spacecraftName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{spacecraftName}", url.PathEscape(spacecraftName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}