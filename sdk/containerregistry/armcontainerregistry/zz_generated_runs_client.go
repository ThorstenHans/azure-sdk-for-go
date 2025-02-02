// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// RunsClient contains the methods for the Runs group.
// Don't use this type directly, use NewRunsClient() instead.
type RunsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewRunsClient creates a new instance of RunsClient with the specified values.
func NewRunsClient(con *armcore.Connection, subscriptionID string) *RunsClient {
	return &RunsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCancel - Cancel an existing run.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) BeginCancel(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsBeginCancelOptions) (HTTPPollerResponse, error) {
	resp, err := client.cancel(ctx, resourceGroupName, registryName, runID, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("RunsClient.Cancel", "", resp, client.con.Pipeline(), client.cancelHandleError)
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

// ResumeCancel creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *RunsClient) ResumeCancel(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("RunsClient.Cancel", token, client.con.Pipeline(), client.cancelHandleError)
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

// Cancel - Cancel an existing run.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) cancel(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsBeginCancelOptions) (*azcore.Response, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, registryName, runID, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.cancelHandleError(resp)
	}
	return resp, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *RunsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsBeginCancelOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *RunsClient) cancelHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the detailed information for a given run.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) Get(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsGetOptions) (RunResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, runID, options)
	if err != nil {
		return RunResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RunResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RunResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RunsClient) getHandleResponse(resp *azcore.Response) (RunResponse, error) {
	var val *Run
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RunResponse{}, err
	}
	return RunResponse{RawResponse: resp.Response, Run: val}, nil
}

// getHandleError handles the Get error response.
func (client *RunsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetLogSasURL - Gets a link to download the run logs.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) GetLogSasURL(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsGetLogSasURLOptions) (RunGetLogResultResponse, error) {
	req, err := client.getLogSasURLCreateRequest(ctx, resourceGroupName, registryName, runID, options)
	if err != nil {
		return RunGetLogResultResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RunGetLogResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RunGetLogResultResponse{}, client.getLogSasURLHandleError(resp)
	}
	return client.getLogSasURLHandleResponse(resp)
}

// getLogSasURLCreateRequest creates the GetLogSasURL request.
func (client *RunsClient) getLogSasURLCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, options *RunsGetLogSasURLOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}/listLogSasUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getLogSasURLHandleResponse handles the GetLogSasURL response.
func (client *RunsClient) getLogSasURLHandleResponse(resp *azcore.Response) (RunGetLogResultResponse, error) {
	var val *RunGetLogResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RunGetLogResultResponse{}, err
	}
	return RunGetLogResultResponse{RawResponse: resp.Response, RunGetLogResult: val}, nil
}

// getLogSasURLHandleError handles the GetLogSasURL error response.
func (client *RunsClient) getLogSasURLHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Gets all the runs for a registry.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) List(resourceGroupName string, registryName string, options *RunsListOptions) RunListResultPager {
	return &runListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp RunListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RunListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *RunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *RunsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RunsClient) listHandleResponse(resp *azcore.Response) (RunListResultResponse, error) {
	var val *RunListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RunListResultResponse{}, err
	}
	return RunListResultResponse{RawResponse: resp.Response, RunListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *RunsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginUpdate - Patch the run properties.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, runID string, runUpdateParameters RunUpdateParameters, options *RunsBeginUpdateOptions) (RunPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, registryName, runID, runUpdateParameters, options)
	if err != nil {
		return RunPollerResponse{}, err
	}
	result := RunPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("RunsClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return RunPollerResponse{}, err
	}
	poller := &runPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (RunResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new RunPoller from the specified resume token.
// token - The value must come from a previous call to RunPoller.ResumeToken().
func (client *RunsClient) ResumeUpdate(ctx context.Context, token string) (RunPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("RunsClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return RunPollerResponse{}, err
	}
	poller := &runPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return RunPollerResponse{}, err
	}
	result := RunPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (RunResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Patch the run properties.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RunsClient) update(ctx context.Context, resourceGroupName string, registryName string, runID string, runUpdateParameters RunUpdateParameters, options *RunsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, runID, runUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *RunsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, runID string, runUpdateParameters RunUpdateParameters, options *RunsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/runs/{runId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(runUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *RunsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
