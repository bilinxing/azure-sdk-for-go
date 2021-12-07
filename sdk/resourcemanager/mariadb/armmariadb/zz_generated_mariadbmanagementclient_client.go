//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmariadb

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MariaDBManagementClient contains the methods for the MariaDBManagementClient group.
// Don't use this type directly, use NewMariaDBManagementClient() instead.
type MariaDBManagementClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewMariaDBManagementClient creates a new instance of MariaDBManagementClient with the specified values.
func NewMariaDBManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MariaDBManagementClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &MariaDBManagementClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateRecommendedActionSession - Create recommendation action session for the advisor.
// If the operation fails it returns a generic error.
func (client *MariaDBManagementClient) BeginCreateRecommendedActionSession(ctx context.Context, resourceGroupName string, serverName string, advisorName string, databaseName string, options *MariaDBManagementClientBeginCreateRecommendedActionSessionOptions) (MariaDBManagementClientCreateRecommendedActionSessionPollerResponse, error) {
	resp, err := client.createRecommendedActionSession(ctx, resourceGroupName, serverName, advisorName, databaseName, options)
	if err != nil {
		return MariaDBManagementClientCreateRecommendedActionSessionPollerResponse{}, err
	}
	result := MariaDBManagementClientCreateRecommendedActionSessionPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MariaDBManagementClient.CreateRecommendedActionSession", "", resp, client.pl, client.createRecommendedActionSessionHandleError)
	if err != nil {
		return MariaDBManagementClientCreateRecommendedActionSessionPollerResponse{}, err
	}
	result.Poller = &MariaDBManagementClientCreateRecommendedActionSessionPoller{
		pt: pt,
	}
	return result, nil
}

// CreateRecommendedActionSession - Create recommendation action session for the advisor.
// If the operation fails it returns a generic error.
func (client *MariaDBManagementClient) createRecommendedActionSession(ctx context.Context, resourceGroupName string, serverName string, advisorName string, databaseName string, options *MariaDBManagementClientBeginCreateRecommendedActionSessionOptions) (*http.Response, error) {
	req, err := client.createRecommendedActionSessionCreateRequest(ctx, resourceGroupName, serverName, advisorName, databaseName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createRecommendedActionSessionHandleError(resp)
	}
	return resp, nil
}

// createRecommendedActionSessionCreateRequest creates the CreateRecommendedActionSession request.
func (client *MariaDBManagementClient) createRecommendedActionSessionCreateRequest(ctx context.Context, resourceGroupName string, serverName string, advisorName string, databaseName string, options *MariaDBManagementClientBeginCreateRecommendedActionSessionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/advisors/{advisorName}/createRecommendedActionSession"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	reqQP.Set("databaseName", databaseName)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// createRecommendedActionSessionHandleError handles the CreateRecommendedActionSession error response.
func (client *MariaDBManagementClient) createRecommendedActionSessionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ResetQueryPerformanceInsightData - Reset data for Query Performance Insight.
// If the operation fails it returns the *CloudError error type.
func (client *MariaDBManagementClient) ResetQueryPerformanceInsightData(ctx context.Context, resourceGroupName string, serverName string, options *MariaDBManagementClientResetQueryPerformanceInsightDataOptions) (MariaDBManagementClientResetQueryPerformanceInsightDataResponse, error) {
	req, err := client.resetQueryPerformanceInsightDataCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return MariaDBManagementClientResetQueryPerformanceInsightDataResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MariaDBManagementClientResetQueryPerformanceInsightDataResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MariaDBManagementClientResetQueryPerformanceInsightDataResponse{}, client.resetQueryPerformanceInsightDataHandleError(resp)
	}
	return client.resetQueryPerformanceInsightDataHandleResponse(resp)
}

// resetQueryPerformanceInsightDataCreateRequest creates the ResetQueryPerformanceInsightData request.
func (client *MariaDBManagementClient) resetQueryPerformanceInsightDataCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *MariaDBManagementClientResetQueryPerformanceInsightDataOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/resetQueryPerformanceInsightData"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// resetQueryPerformanceInsightDataHandleResponse handles the ResetQueryPerformanceInsightData response.
func (client *MariaDBManagementClient) resetQueryPerformanceInsightDataHandleResponse(resp *http.Response) (MariaDBManagementClientResetQueryPerformanceInsightDataResponse, error) {
	result := MariaDBManagementClientResetQueryPerformanceInsightDataResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryPerformanceInsightResetDataResult); err != nil {
		return MariaDBManagementClientResetQueryPerformanceInsightDataResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// resetQueryPerformanceInsightDataHandleError handles the ResetQueryPerformanceInsightData error response.
func (client *MariaDBManagementClient) resetQueryPerformanceInsightDataHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
