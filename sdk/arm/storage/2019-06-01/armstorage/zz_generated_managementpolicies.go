// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorage

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// ManagementPoliciesClient contains the methods for the ManagementPolicies group.
// Don't use this type directly, use NewManagementPoliciesClient() instead.
type ManagementPoliciesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewManagementPoliciesClient creates a new instance of ManagementPoliciesClient with the specified values.
func NewManagementPoliciesClient(con *armcore.Connection, subscriptionID string) *ManagementPoliciesClient {
	return &ManagementPoliciesClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Sets the managementpolicy to the specified storage account.
func (client *ManagementPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, properties ManagementPolicy, options *ManagementPoliciesCreateOrUpdateOptions) (ManagementPolicyResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, managementPolicyName, properties, options)
	if err != nil {
		return ManagementPolicyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ManagementPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ManagementPolicyResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagementPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, properties ManagementPolicy, options *ManagementPoliciesCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{managementPolicyName}", url.PathEscape(string(managementPolicyName)))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(properties)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagementPoliciesClient) createOrUpdateHandleResponse(resp *azcore.Response) (ManagementPolicyResponse, error) {
	var val *ManagementPolicy
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ManagementPolicyResponse{}, err
	}
	return ManagementPolicyResponse{RawResponse: resp.Response, ManagementPolicy: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ManagementPoliciesClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Delete - Deletes the managementpolicy associated with the specified storage account.
func (client *ManagementPoliciesClient) Delete(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, options *ManagementPoliciesDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, managementPolicyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagementPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, options *ManagementPoliciesDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{managementPolicyName}", url.PathEscape(string(managementPolicyName)))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ManagementPoliciesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets the managementpolicy associated with the specified storage account.
func (client *ManagementPoliciesClient) Get(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, options *ManagementPoliciesGetOptions) (ManagementPolicyResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, managementPolicyName, options)
	if err != nil {
		return ManagementPolicyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ManagementPolicyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ManagementPolicyResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagementPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, managementPolicyName ManagementPolicyName, options *ManagementPoliciesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/managementPolicies/{managementPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{managementPolicyName}", url.PathEscape(string(managementPolicyName)))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagementPoliciesClient) getHandleResponse(resp *azcore.Response) (ManagementPolicyResponse, error) {
	var val *ManagementPolicy
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ManagementPolicyResponse{}, err
	}
	return ManagementPolicyResponse{RawResponse: resp.Response, ManagementPolicy: val}, nil
}

// getHandleError handles the Get error response.
func (client *ManagementPoliciesClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}