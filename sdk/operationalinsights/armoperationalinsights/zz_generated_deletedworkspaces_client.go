// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// DeletedWorkspacesClient contains the methods for the DeletedWorkspaces group.
// Don't use this type directly, use NewDeletedWorkspacesClient() instead.
type DeletedWorkspacesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewDeletedWorkspacesClient creates a new instance of DeletedWorkspacesClient with the specified values.
func NewDeletedWorkspacesClient(con *armcore.Connection, subscriptionID string) *DeletedWorkspacesClient {
	return &DeletedWorkspacesClient{con: con, subscriptionID: subscriptionID}
}

// List - Gets recently deleted workspaces in a subscription, available for recovery.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DeletedWorkspacesClient) List(ctx context.Context, options *DeletedWorkspacesListOptions) (DeletedWorkspacesListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return DeletedWorkspacesListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DeletedWorkspacesListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DeletedWorkspacesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *DeletedWorkspacesClient) listCreateRequest(ctx context.Context, options *DeletedWorkspacesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.OperationalInsights/deletedWorkspaces"
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
	reqQP.Set("api-version", "2020-10-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeletedWorkspacesClient) listHandleResponse(resp *azcore.Response) (DeletedWorkspacesListResponse, error) {
	result := DeletedWorkspacesListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.WorkspaceListResult); err != nil {
		return DeletedWorkspacesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *DeletedWorkspacesClient) listHandleError(resp *azcore.Response) error {
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

// ListByResourceGroup - Gets recently deleted workspaces in a resource group, available for recovery.
// If the operation fails it returns the *ErrorResponse error type.
func (client *DeletedWorkspacesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *DeletedWorkspacesListByResourceGroupOptions) (DeletedWorkspacesListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return DeletedWorkspacesListByResourceGroupResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DeletedWorkspacesListByResourceGroupResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DeletedWorkspacesListByResourceGroupResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DeletedWorkspacesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DeletedWorkspacesListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/deletedWorkspaces"
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
	reqQP.Set("api-version", "2020-10-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DeletedWorkspacesClient) listByResourceGroupHandleResponse(resp *azcore.Response) (DeletedWorkspacesListByResourceGroupResponse, error) {
	result := DeletedWorkspacesListByResourceGroupResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.WorkspaceListResult); err != nil {
		return DeletedWorkspacesListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DeletedWorkspacesClient) listByResourceGroupHandleError(resp *azcore.Response) error {
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
