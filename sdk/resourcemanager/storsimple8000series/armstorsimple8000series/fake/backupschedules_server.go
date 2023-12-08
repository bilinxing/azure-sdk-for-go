//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple8000series/armstorsimple8000series"
	"net/http"
	"net/url"
	"regexp"
)

// BackupSchedulesServer is a fake server for instances of the armstorsimple8000series.BackupSchedulesClient type.
type BackupSchedulesServer struct {
	// BeginCreateOrUpdate is the fake for method BackupSchedulesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, parameters armstorsimple8000series.BackupSchedule, options *armstorsimple8000series.BackupSchedulesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armstorsimple8000series.BackupSchedulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method BackupSchedulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, options *armstorsimple8000series.BackupSchedulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armstorsimple8000series.BackupSchedulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BackupSchedulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, deviceName string, backupPolicyName string, backupScheduleName string, resourceGroupName string, managerName string, options *armstorsimple8000series.BackupSchedulesClientGetOptions) (resp azfake.Responder[armstorsimple8000series.BackupSchedulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByBackupPolicyPager is the fake for method BackupSchedulesClient.NewListByBackupPolicyPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBackupPolicyPager func(deviceName string, backupPolicyName string, resourceGroupName string, managerName string, options *armstorsimple8000series.BackupSchedulesClientListByBackupPolicyOptions) (resp azfake.PagerResponder[armstorsimple8000series.BackupSchedulesClientListByBackupPolicyResponse])
}

// NewBackupSchedulesServerTransport creates a new instance of BackupSchedulesServerTransport with the provided implementation.
// The returned BackupSchedulesServerTransport instance is connected to an instance of armstorsimple8000series.BackupSchedulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupSchedulesServerTransport(srv *BackupSchedulesServer) *BackupSchedulesServerTransport {
	return &BackupSchedulesServerTransport{
		srv:                        srv,
		beginCreateOrUpdate:        newTracker[azfake.PollerResponder[armstorsimple8000series.BackupSchedulesClientCreateOrUpdateResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armstorsimple8000series.BackupSchedulesClientDeleteResponse]](),
		newListByBackupPolicyPager: newTracker[azfake.PagerResponder[armstorsimple8000series.BackupSchedulesClientListByBackupPolicyResponse]](),
	}
}

// BackupSchedulesServerTransport connects instances of armstorsimple8000series.BackupSchedulesClient to instances of BackupSchedulesServer.
// Don't use this type directly, use NewBackupSchedulesServerTransport instead.
type BackupSchedulesServerTransport struct {
	srv                        *BackupSchedulesServer
	beginCreateOrUpdate        *tracker[azfake.PollerResponder[armstorsimple8000series.BackupSchedulesClientCreateOrUpdateResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armstorsimple8000series.BackupSchedulesClientDeleteResponse]]
	newListByBackupPolicyPager *tracker[azfake.PagerResponder[armstorsimple8000series.BackupSchedulesClientListByBackupPolicyResponse]]
}

// Do implements the policy.Transporter interface for BackupSchedulesServerTransport.
func (b *BackupSchedulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BackupSchedulesClient.BeginCreateOrUpdate":
		resp, err = b.dispatchBeginCreateOrUpdate(req)
	case "BackupSchedulesClient.BeginDelete":
		resp, err = b.dispatchBeginDelete(req)
	case "BackupSchedulesClient.Get":
		resp, err = b.dispatchGet(req)
	case "BackupSchedulesClient.NewListByBackupPolicyPager":
		resp, err = b.dispatchNewListByBackupPolicyPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BackupSchedulesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := b.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<backupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules/(?P<backupScheduleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstorsimple8000series.BackupSchedule](req)
		if err != nil {
			return nil, err
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		backupPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupPolicyName")])
		if err != nil {
			return nil, err
		}
		backupScheduleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupScheduleName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginCreateOrUpdate(req.Context(), deviceNameParam, backupPolicyNameParam, backupScheduleNameParam, resourceGroupNameParam, managerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		b.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		b.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		b.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (b *BackupSchedulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if b.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := b.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<backupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules/(?P<backupScheduleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		backupPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupPolicyName")])
		if err != nil {
			return nil, err
		}
		backupScheduleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupScheduleName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginDelete(req.Context(), deviceNameParam, backupPolicyNameParam, backupScheduleNameParam, resourceGroupNameParam, managerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		b.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		b.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		b.beginDelete.remove(req)
	}

	return resp, nil
}

func (b *BackupSchedulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<backupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules/(?P<backupScheduleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
	if err != nil {
		return nil, err
	}
	backupPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupPolicyName")])
	if err != nil {
		return nil, err
	}
	backupScheduleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupScheduleName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), deviceNameParam, backupPolicyNameParam, backupScheduleNameParam, resourceGroupNameParam, managerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BackupSchedule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BackupSchedulesServerTransport) dispatchNewListByBackupPolicyPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListByBackupPolicyPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBackupPolicyPager not implemented")}
	}
	newListByBackupPolicyPager := b.newListByBackupPolicyPager.get(req)
	if newListByBackupPolicyPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorSimple/managers/(?P<managerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/devices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<backupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		deviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		backupPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupPolicyName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managerName")])
		if err != nil {
			return nil, err
		}
		resp := b.srv.NewListByBackupPolicyPager(deviceNameParam, backupPolicyNameParam, resourceGroupNameParam, managerNameParam, nil)
		newListByBackupPolicyPager = &resp
		b.newListByBackupPolicyPager.add(req, newListByBackupPolicyPager)
	}
	resp, err := server.PagerResponderNext(newListByBackupPolicyPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListByBackupPolicyPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBackupPolicyPager) {
		b.newListByBackupPolicyPager.remove(req)
	}
	return resp, nil
}