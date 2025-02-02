// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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
)

// APIReleaseClient contains the methods for the APIRelease group.
// Don't use this type directly, use NewAPIReleaseClient() instead.
type APIReleaseClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewAPIReleaseClient creates a new instance of APIReleaseClient with the specified values.
func NewAPIReleaseClient(con *armcore.Connection, subscriptionID string) *APIReleaseClient {
	return &APIReleaseClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates a new Release for the API.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIReleaseClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, parameters APIReleaseContract, options *APIReleaseCreateOrUpdateOptions) (APIReleaseCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, apiID, releaseID, parameters, options)
	if err != nil {
		return APIReleaseCreateOrUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APIReleaseCreateOrUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return APIReleaseCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *APIReleaseClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, parameters APIReleaseContract, options *APIReleaseCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if releaseID == "" {
		return nil, errors.New("parameter releaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{releaseId}", url.PathEscape(releaseID))
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
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *APIReleaseClient) createOrUpdateHandleResponse(resp *azcore.Response) (APIReleaseCreateOrUpdateResponse, error) {
	result := APIReleaseCreateOrUpdateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.APIReleaseContract); err != nil {
		return APIReleaseCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *APIReleaseClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Deletes the specified release in the API.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIReleaseClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, ifMatch string, options *APIReleaseDeleteOptions) (APIReleaseDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, apiID, releaseID, ifMatch, options)
	if err != nil {
		return APIReleaseDeleteResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APIReleaseDeleteResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return APIReleaseDeleteResponse{}, client.deleteHandleError(resp)
	}
	return APIReleaseDeleteResponse{RawResponse: resp.Response}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *APIReleaseClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, ifMatch string, options *APIReleaseDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if releaseID == "" {
		return nil, errors.New("parameter releaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{releaseId}", url.PathEscape(releaseID))
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
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *APIReleaseClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Returns the details of an API release.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIReleaseClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, options *APIReleaseGetOptions) (APIReleaseGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, apiID, releaseID, options)
	if err != nil {
		return APIReleaseGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APIReleaseGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return APIReleaseGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *APIReleaseClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, options *APIReleaseGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if releaseID == "" {
		return nil, errors.New("parameter releaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{releaseId}", url.PathEscape(releaseID))
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
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *APIReleaseClient) getHandleResponse(resp *azcore.Response) (APIReleaseGetResponse, error) {
	result := APIReleaseGetResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.APIReleaseContract); err != nil {
		return APIReleaseGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *APIReleaseClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetEntityTag - Returns the etag of an API release.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIReleaseClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, options *APIReleaseGetEntityTagOptions) (APIReleaseGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, apiID, releaseID, options)
	if err != nil {
		return APIReleaseGetEntityTagResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APIReleaseGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *APIReleaseClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, options *APIReleaseGetEntityTagOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if releaseID == "" {
		return nil, errors.New("parameter releaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{releaseId}", url.PathEscape(releaseID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *APIReleaseClient) getEntityTagHandleResponse(resp *azcore.Response) (APIReleaseGetEntityTagResponse, error) {
	result := APIReleaseGetEntityTagResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists all releases of an API. An API release is created when making an API Revision current. Releases are also used to rollback to previous
// revisions. Results will be paged and can be constrained by
// the $top and $skip parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIReleaseClient) ListByService(resourceGroupName string, serviceName string, apiID string, options *APIReleaseListByServiceOptions) APIReleaseListByServicePager {
	return &apiReleaseListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, apiID, options)
		},
		advancer: func(ctx context.Context, resp APIReleaseListByServiceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.APIReleaseCollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *APIReleaseClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, options *APIReleaseListByServiceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *APIReleaseClient) listByServiceHandleResponse(resp *azcore.Response) (APIReleaseListByServiceResponse, error) {
	result := APIReleaseListByServiceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.APIReleaseCollection); err != nil {
		return APIReleaseListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *APIReleaseClient) listByServiceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Update - Updates the details of the release of the API specified by its identifier.
// If the operation fails it returns the *ErrorResponse error type.
func (client *APIReleaseClient) Update(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, ifMatch string, parameters APIReleaseContract, options *APIReleaseUpdateOptions) (APIReleaseUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, apiID, releaseID, ifMatch, parameters, options)
	if err != nil {
		return APIReleaseUpdateResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return APIReleaseUpdateResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return APIReleaseUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *APIReleaseClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, releaseID string, ifMatch string, parameters APIReleaseContract, options *APIReleaseUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/releases/{releaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if releaseID == "" {
		return nil, errors.New("parameter releaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{releaseId}", url.PathEscape(releaseID))
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
	reqQP.Set("api-version", "2020-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("If-Match", ifMatch)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client *APIReleaseClient) updateHandleResponse(resp *azcore.Response) (APIReleaseUpdateResponse, error) {
	result := APIReleaseUpdateResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := resp.UnmarshalAsJSON(&result.APIReleaseContract); err != nil {
		return APIReleaseUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *APIReleaseClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
