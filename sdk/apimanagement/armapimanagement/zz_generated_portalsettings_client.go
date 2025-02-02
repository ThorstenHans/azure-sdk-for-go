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
	"strings"
)

// PortalSettingsClient contains the methods for the PortalSettings group.
// Don't use this type directly, use NewPortalSettingsClient() instead.
type PortalSettingsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewPortalSettingsClient creates a new instance of PortalSettingsClient with the specified values.
func NewPortalSettingsClient(con *armcore.Connection, subscriptionID string) *PortalSettingsClient {
	return &PortalSettingsClient{con: con, subscriptionID: subscriptionID}
}

// ListByService - Lists a collection of portalsettings defined within a service instance..
// If the operation fails it returns the *ErrorResponse error type.
func (client *PortalSettingsClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string, options *PortalSettingsListByServiceOptions) (PortalSettingsListByServiceResponse, error) {
	req, err := client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return PortalSettingsListByServiceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PortalSettingsListByServiceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PortalSettingsListByServiceResponse{}, client.listByServiceHandleError(resp)
	}
	return client.listByServiceHandleResponse(resp)
}

// listByServiceCreateRequest creates the ListByService request.
func (client *PortalSettingsClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *PortalSettingsListByServiceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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

// listByServiceHandleResponse handles the ListByService response.
func (client *PortalSettingsClient) listByServiceHandleResponse(resp *azcore.Response) (PortalSettingsListByServiceResponse, error) {
	result := PortalSettingsListByServiceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PortalSettingsCollection); err != nil {
		return PortalSettingsListByServiceResponse{}, err
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *PortalSettingsClient) listByServiceHandleError(resp *azcore.Response) error {
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
