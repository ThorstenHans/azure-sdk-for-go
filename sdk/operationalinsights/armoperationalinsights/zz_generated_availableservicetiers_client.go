// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// AvailableServiceTiersClient contains the methods for the AvailableServiceTiers group.
// Don't use this type directly, use NewAvailableServiceTiersClient() instead.
type AvailableServiceTiersClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewAvailableServiceTiersClient creates a new instance of AvailableServiceTiersClient with the specified values.
func NewAvailableServiceTiersClient(con *armcore.Connection, subscriptionID string) *AvailableServiceTiersClient {
	return &AvailableServiceTiersClient{con: con, subscriptionID: subscriptionID}
}

// ListByWorkspace - Gets the available service tiers for the workspace.
// If the operation fails it returns a generic error.
func (client *AvailableServiceTiersClient) ListByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, options *AvailableServiceTiersListByWorkspaceOptions) (AvailableServiceTiersListByWorkspaceResponse, error) {
	req, err := client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return AvailableServiceTiersListByWorkspaceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AvailableServiceTiersListByWorkspaceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return AvailableServiceTiersListByWorkspaceResponse{}, client.listByWorkspaceHandleError(resp)
	}
	return client.listByWorkspaceHandleResponse(resp)
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *AvailableServiceTiersClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *AvailableServiceTiersListByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/availableServiceTiers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *AvailableServiceTiersClient) listByWorkspaceHandleResponse(resp *azcore.Response) (AvailableServiceTiersListByWorkspaceResponse, error) {
	result := AvailableServiceTiersListByWorkspaceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.AvailableServiceTierArray); err != nil {
		return AvailableServiceTiersListByWorkspaceResponse{}, err
	}
	return result, nil
}

// listByWorkspaceHandleError handles the ListByWorkspace error response.
func (client *AvailableServiceTiersClient) listByWorkspaceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
