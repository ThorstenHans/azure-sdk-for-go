// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armweb

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

// KubeEnvironmentsClient contains the methods for the KubeEnvironments group.
// Don't use this type directly, use NewKubeEnvironmentsClient() instead.
type KubeEnvironmentsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewKubeEnvironmentsClient creates a new instance of KubeEnvironmentsClient with the specified values.
func NewKubeEnvironmentsClient(con *armcore.Connection, subscriptionID string) *KubeEnvironmentsClient {
	return &KubeEnvironmentsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Description for Creates or updates a Kubernetes Environment.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *KubeEnvironmentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, name string, kubeEnvironmentEnvelope KubeEnvironment, options *KubeEnvironmentsBeginCreateOrUpdateOptions) (KubeEnvironmentsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, name, kubeEnvironmentEnvelope, options)
	if err != nil {
		return KubeEnvironmentsCreateOrUpdatePollerResponse{}, err
	}
	result := KubeEnvironmentsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("KubeEnvironmentsClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return KubeEnvironmentsCreateOrUpdatePollerResponse{}, err
	}
	poller := &kubeEnvironmentsCreateOrUpdatePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (KubeEnvironmentsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new KubeEnvironmentsCreateOrUpdatePoller from the specified resume token.
// token - The value must come from a previous call to KubeEnvironmentsCreateOrUpdatePoller.ResumeToken().
func (client *KubeEnvironmentsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (KubeEnvironmentsCreateOrUpdatePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("KubeEnvironmentsClient.CreateOrUpdate", token, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return KubeEnvironmentsCreateOrUpdatePollerResponse{}, err
	}
	poller := &kubeEnvironmentsCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return KubeEnvironmentsCreateOrUpdatePollerResponse{}, err
	}
	result := KubeEnvironmentsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (KubeEnvironmentsCreateOrUpdateResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Description for Creates or updates a Kubernetes Environment.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *KubeEnvironmentsClient) createOrUpdate(ctx context.Context, resourceGroupName string, name string, kubeEnvironmentEnvelope KubeEnvironment, options *KubeEnvironmentsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, kubeEnvironmentEnvelope, options)
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
func (client *KubeEnvironmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, kubeEnvironmentEnvelope KubeEnvironment, options *KubeEnvironmentsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/kubeEnvironments/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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
	reqQP.Set("api-version", "2021-01-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(kubeEnvironmentEnvelope)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *KubeEnvironmentsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Description for Delete a Kubernetes Environment.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *KubeEnvironmentsClient) BeginDelete(ctx context.Context, resourceGroupName string, name string, options *KubeEnvironmentsBeginDeleteOptions) (KubeEnvironmentsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, name, options)
	if err != nil {
		return KubeEnvironmentsDeletePollerResponse{}, err
	}
	result := KubeEnvironmentsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("KubeEnvironmentsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return KubeEnvironmentsDeletePollerResponse{}, err
	}
	poller := &kubeEnvironmentsDeletePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (KubeEnvironmentsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new KubeEnvironmentsDeletePoller from the specified resume token.
// token - The value must come from a previous call to KubeEnvironmentsDeletePoller.ResumeToken().
func (client *KubeEnvironmentsClient) ResumeDelete(ctx context.Context, token string) (KubeEnvironmentsDeletePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("KubeEnvironmentsClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return KubeEnvironmentsDeletePollerResponse{}, err
	}
	poller := &kubeEnvironmentsDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return KubeEnvironmentsDeletePollerResponse{}, err
	}
	result := KubeEnvironmentsDeletePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (KubeEnvironmentsDeleteResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Description for Delete a Kubernetes Environment.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *KubeEnvironmentsClient) deleteOperation(ctx context.Context, resourceGroupName string, name string, options *KubeEnvironmentsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, options)
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
func (client *KubeEnvironmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, options *KubeEnvironmentsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/kubeEnvironments/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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
	reqQP.Set("api-version", "2021-01-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *KubeEnvironmentsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Description for Get the properties of a Kubernetes Environment.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *KubeEnvironmentsClient) Get(ctx context.Context, resourceGroupName string, name string, options *KubeEnvironmentsGetOptions) (KubeEnvironmentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, options)
	if err != nil {
		return KubeEnvironmentsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return KubeEnvironmentsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return KubeEnvironmentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *KubeEnvironmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, options *KubeEnvironmentsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/kubeEnvironments/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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
	reqQP.Set("api-version", "2021-01-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *KubeEnvironmentsClient) getHandleResponse(resp *azcore.Response) (KubeEnvironmentsGetResponse, error) {
	result := KubeEnvironmentsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.KubeEnvironment); err != nil {
		return KubeEnvironmentsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *KubeEnvironmentsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByResourceGroup - Description for Get all the Kubernetes Environments in a resource group.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *KubeEnvironmentsClient) ListByResourceGroup(resourceGroupName string, options *KubeEnvironmentsListByResourceGroupOptions) KubeEnvironmentsListByResourceGroupPager {
	return &kubeEnvironmentsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp KubeEnvironmentsListByResourceGroupResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.KubeEnvironmentCollection.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *KubeEnvironmentsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *KubeEnvironmentsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/kubeEnvironments"
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
	reqQP.Set("api-version", "2021-01-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *KubeEnvironmentsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (KubeEnvironmentsListByResourceGroupResponse, error) {
	result := KubeEnvironmentsListByResourceGroupResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.KubeEnvironmentCollection); err != nil {
		return KubeEnvironmentsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *KubeEnvironmentsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListBySubscription - Description for Get all Kubernetes Environments for a subscription.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *KubeEnvironmentsClient) ListBySubscription(options *KubeEnvironmentsListBySubscriptionOptions) KubeEnvironmentsListBySubscriptionPager {
	return &kubeEnvironmentsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp KubeEnvironmentsListBySubscriptionResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.KubeEnvironmentCollection.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *KubeEnvironmentsClient) listBySubscriptionCreateRequest(ctx context.Context, options *KubeEnvironmentsListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/kubeEnvironments"
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
	reqQP.Set("api-version", "2021-01-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *KubeEnvironmentsClient) listBySubscriptionHandleResponse(resp *azcore.Response) (KubeEnvironmentsListBySubscriptionResponse, error) {
	result := KubeEnvironmentsListBySubscriptionResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.KubeEnvironmentCollection); err != nil {
		return KubeEnvironmentsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *KubeEnvironmentsClient) listBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Update - Description for Creates or updates a Kubernetes Environment.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *KubeEnvironmentsClient) Update(ctx context.Context, resourceGroupName string, name string, kubeEnvironmentEnvelope KubeEnvironmentPatchResource, options *KubeEnvironmentsUpdateOptions) (KubeEnvironmentsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, name, kubeEnvironmentEnvelope, options)
	if err != nil {
		return KubeEnvironmentsUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return KubeEnvironmentsUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return KubeEnvironmentsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *KubeEnvironmentsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, name string, kubeEnvironmentEnvelope KubeEnvironmentPatchResource, options *KubeEnvironmentsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/kubeEnvironments/{name}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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
	reqQP.Set("api-version", "2021-01-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(kubeEnvironmentEnvelope)
}

// updateHandleResponse handles the Update response.
func (client *KubeEnvironmentsClient) updateHandleResponse(resp *azcore.Response) (KubeEnvironmentsUpdateResponse, error) {
	result := KubeEnvironmentsUpdateResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.KubeEnvironment); err != nil {
		return KubeEnvironmentsUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *KubeEnvironmentsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
