// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ApplicationSecurityGroupsClient contains the methods for the ApplicationSecurityGroups group.
// Don't use this type directly, use NewApplicationSecurityGroupsClient() instead.
type ApplicationSecurityGroupsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewApplicationSecurityGroupsClient creates a new instance of ApplicationSecurityGroupsClient with the specified values.
func NewApplicationSecurityGroupsClient(con *armcore.Connection, subscriptionID string) *ApplicationSecurityGroupsClient {
	return &ApplicationSecurityGroupsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates an application security group.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationSecurityGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsBeginCreateOrUpdateOptions) (ApplicationSecurityGroupPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
	if err != nil {
		return ApplicationSecurityGroupPollerResponse{}, err
	}
	result := ApplicationSecurityGroupPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ApplicationSecurityGroupsClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ApplicationSecurityGroupPollerResponse{}, err
	}
	poller := &applicationSecurityGroupPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ApplicationSecurityGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new ApplicationSecurityGroupPoller from the specified resume token.
// token - The value must come from a previous call to ApplicationSecurityGroupPoller.ResumeToken().
func (client *ApplicationSecurityGroupsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (ApplicationSecurityGroupPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ApplicationSecurityGroupsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ApplicationSecurityGroupPollerResponse{}, err
	}
	poller := &applicationSecurityGroupPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ApplicationSecurityGroupPollerResponse{}, err
	}
	result := ApplicationSecurityGroupPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ApplicationSecurityGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an application security group.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationSecurityGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationSecurityGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup, options *ApplicationSecurityGroupsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ApplicationSecurityGroupsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes the specified application security group.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationSecurityGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, applicationSecurityGroupName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ApplicationSecurityGroupsClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ApplicationSecurityGroupsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ApplicationSecurityGroupsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified application security group.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationSecurityGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationSecurityGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ApplicationSecurityGroupsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets information about the specified application security group.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationSecurityGroupsClient) Get(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsGetOptions) (ApplicationSecurityGroupResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, options)
	if err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ApplicationSecurityGroupResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationSecurityGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, options *ApplicationSecurityGroupsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationSecurityGroupsClient) getHandleResponse(resp *azcore.Response) (ApplicationSecurityGroupResponse, error) {
	var val *ApplicationSecurityGroup
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	return ApplicationSecurityGroupResponse{RawResponse: resp.Response, ApplicationSecurityGroup: val}, nil
}

// getHandleError handles the Get error response.
func (client *ApplicationSecurityGroupsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets all the application security groups in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationSecurityGroupsClient) List(resourceGroupName string, options *ApplicationSecurityGroupsListOptions) ApplicationSecurityGroupListResultPager {
	return &applicationSecurityGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ApplicationSecurityGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ApplicationSecurityGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *ApplicationSecurityGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationSecurityGroupsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ApplicationSecurityGroupsClient) listHandleResponse(resp *azcore.Response) (ApplicationSecurityGroupListResultResponse, error) {
	var val *ApplicationSecurityGroupListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ApplicationSecurityGroupListResultResponse{}, err
	}
	return ApplicationSecurityGroupListResultResponse{RawResponse: resp.Response, ApplicationSecurityGroupListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ApplicationSecurityGroupsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListAll - Gets all application security groups in a subscription.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationSecurityGroupsClient) ListAll(options *ApplicationSecurityGroupsListAllOptions) ApplicationSecurityGroupListResultPager {
	return &applicationSecurityGroupListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp ApplicationSecurityGroupListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ApplicationSecurityGroupListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *ApplicationSecurityGroupsClient) listAllCreateRequest(ctx context.Context, options *ApplicationSecurityGroupsListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationSecurityGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *ApplicationSecurityGroupsClient) listAllHandleResponse(resp *azcore.Response) (ApplicationSecurityGroupListResultResponse, error) {
	var val *ApplicationSecurityGroupListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ApplicationSecurityGroupListResultResponse{}, err
	}
	return ApplicationSecurityGroupListResultResponse{RawResponse: resp.Response, ApplicationSecurityGroupListResult: val}, nil
}

// listAllHandleError handles the ListAll error response.
func (client *ApplicationSecurityGroupsClient) listAllHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// UpdateTags - Updates an application security group's tags.
// If the operation fails it returns the *CloudError error type.
func (client *ApplicationSecurityGroupsClient) UpdateTags(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters TagsObject, options *ApplicationSecurityGroupsUpdateTagsOptions) (ApplicationSecurityGroupResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, applicationSecurityGroupName, parameters, options)
	if err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ApplicationSecurityGroupResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ApplicationSecurityGroupsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters TagsObject, options *ApplicationSecurityGroupsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationSecurityGroupName == "" {
		return nil, errors.New("parameter applicationSecurityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ApplicationSecurityGroupsClient) updateTagsHandleResponse(resp *azcore.Response) (ApplicationSecurityGroupResponse, error) {
	var val *ApplicationSecurityGroup
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ApplicationSecurityGroupResponse{}, err
	}
	return ApplicationSecurityGroupResponse{RawResponse: resp.Response, ApplicationSecurityGroup: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *ApplicationSecurityGroupsClient) updateTagsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
