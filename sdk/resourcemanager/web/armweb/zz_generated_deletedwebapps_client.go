//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armweb

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

// DeletedWebAppsClient contains the methods for the DeletedWebApps group.
// Don't use this type directly, use NewDeletedWebAppsClient() instead.
type DeletedWebAppsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDeletedWebAppsClient creates a new instance of DeletedWebAppsClient with the specified values.
func NewDeletedWebAppsClient(con *arm.Connection, subscriptionID string) *DeletedWebAppsClient {
	return &DeletedWebAppsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// GetDeletedWebAppByLocation - Description for Get deleted app for a subscription at location.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *DeletedWebAppsClient) GetDeletedWebAppByLocation(ctx context.Context, location string, deletedSiteID string, options *DeletedWebAppsGetDeletedWebAppByLocationOptions) (DeletedWebAppsGetDeletedWebAppByLocationResponse, error) {
	req, err := client.getDeletedWebAppByLocationCreateRequest(ctx, location, deletedSiteID, options)
	if err != nil {
		return DeletedWebAppsGetDeletedWebAppByLocationResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DeletedWebAppsGetDeletedWebAppByLocationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeletedWebAppsGetDeletedWebAppByLocationResponse{}, client.getDeletedWebAppByLocationHandleError(resp)
	}
	return client.getDeletedWebAppByLocationHandleResponse(resp)
}

// getDeletedWebAppByLocationCreateRequest creates the GetDeletedWebAppByLocation request.
func (client *DeletedWebAppsClient) getDeletedWebAppByLocationCreateRequest(ctx context.Context, location string, deletedSiteID string, options *DeletedWebAppsGetDeletedWebAppByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/deletedSites/{deletedSiteId}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if deletedSiteID == "" {
		return nil, errors.New("parameter deletedSiteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deletedSiteId}", url.PathEscape(deletedSiteID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDeletedWebAppByLocationHandleResponse handles the GetDeletedWebAppByLocation response.
func (client *DeletedWebAppsClient) getDeletedWebAppByLocationHandleResponse(resp *http.Response) (DeletedWebAppsGetDeletedWebAppByLocationResponse, error) {
	result := DeletedWebAppsGetDeletedWebAppByLocationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedSite); err != nil {
		return DeletedWebAppsGetDeletedWebAppByLocationResponse{}, err
	}
	return result, nil
}

// getDeletedWebAppByLocationHandleError handles the GetDeletedWebAppByLocation error response.
func (client *DeletedWebAppsClient) getDeletedWebAppByLocationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Description for Get all deleted apps for a subscription.
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *DeletedWebAppsClient) List(options *DeletedWebAppsListOptions) *DeletedWebAppsListPager {
	return &DeletedWebAppsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DeletedWebAppsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DeletedWebAppCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *DeletedWebAppsClient) listCreateRequest(ctx context.Context, options *DeletedWebAppsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeletedWebAppsClient) listHandleResponse(resp *http.Response) (DeletedWebAppsListResponse, error) {
	result := DeletedWebAppsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedWebAppCollection); err != nil {
		return DeletedWebAppsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DeletedWebAppsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByLocation - Description for Get all deleted apps for a subscription at location
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *DeletedWebAppsClient) ListByLocation(location string, options *DeletedWebAppsListByLocationOptions) *DeletedWebAppsListByLocationPager {
	return &DeletedWebAppsListByLocationPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByLocationCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp DeletedWebAppsListByLocationResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DeletedWebAppCollection.NextLink)
		},
	}
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *DeletedWebAppsClient) listByLocationCreateRequest(ctx context.Context, location string, options *DeletedWebAppsListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/deletedSites"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *DeletedWebAppsClient) listByLocationHandleResponse(resp *http.Response) (DeletedWebAppsListByLocationResponse, error) {
	result := DeletedWebAppsListByLocationResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedWebAppCollection); err != nil {
		return DeletedWebAppsListByLocationResponse{}, err
	}
	return result, nil
}

// listByLocationHandleError handles the ListByLocation error response.
func (client *DeletedWebAppsClient) listByLocationHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}