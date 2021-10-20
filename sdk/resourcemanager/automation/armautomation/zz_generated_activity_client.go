//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ActivityClient contains the methods for the Activity group.
// Don't use this type directly, use NewActivityClient() instead.
type ActivityClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewActivityClient creates a new instance of ActivityClient with the specified values.
func NewActivityClient(con *arm.Connection, subscriptionID string) *ActivityClient {
	return &ActivityClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Get - Retrieve the activity in the module identified by module name and activity name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActivityClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, activityName string, options *ActivityGetOptions) (ActivityGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, activityName, options)
	if err != nil {
		return ActivityGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ActivityGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ActivityGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ActivityClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, activityName string, options *ActivityGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/activities/{activityName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if moduleName == "" {
		return nil, errors.New("parameter moduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moduleName}", url.PathEscape(moduleName))
	if activityName == "" {
		return nil, errors.New("parameter activityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activityName}", url.PathEscape(activityName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ActivityClient) getHandleResponse(resp *http.Response) (ActivityGetResponse, error) {
	result := ActivityGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Activity); err != nil {
		return ActivityGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ActivityClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByModule - Retrieve a list of activities in the module identified by module name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ActivityClient) ListByModule(resourceGroupName string, automationAccountName string, moduleName string, options *ActivityListByModuleOptions) *ActivityListByModulePager {
	return &ActivityListByModulePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByModuleCreateRequest(ctx, resourceGroupName, automationAccountName, moduleName, options)
		},
		advancer: func(ctx context.Context, resp ActivityListByModuleResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ActivityListResult.NextLink)
		},
	}
}

// listByModuleCreateRequest creates the ListByModule request.
func (client *ActivityClient) listByModuleCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, moduleName string, options *ActivityListByModuleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/modules/{moduleName}/activities"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if moduleName == "" {
		return nil, errors.New("parameter moduleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moduleName}", url.PathEscape(moduleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByModuleHandleResponse handles the ListByModule response.
func (client *ActivityClient) listByModuleHandleResponse(resp *http.Response) (ActivityListByModuleResponse, error) {
	result := ActivityListByModuleResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActivityListResult); err != nil {
		return ActivityListByModuleResponse{}, err
	}
	return result, nil
}

// listByModuleHandleError handles the ListByModule error response.
func (client *ActivityClient) listByModuleHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}