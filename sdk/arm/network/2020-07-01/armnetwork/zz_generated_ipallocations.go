// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IPAllocationsClient contains the methods for the IPAllocations group.
// Don't use this type directly, use NewIPAllocationsClient() instead.
type IPAllocationsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewIPAllocationsClient creates a new instance of IPAllocationsClient with the specified values.
func NewIPAllocationsClient(con *armcore.Connection, subscriptionID string) IPAllocationsClient {
	return IPAllocationsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client IPAllocationsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// BeginCreateOrUpdate - Creates or updates an IpAllocation in the specified resource group.
func (client IPAllocationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsBeginCreateOrUpdateOptions) (IPAllocationPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, ipAllocationName, parameters, options)
	if err != nil {
		return IPAllocationPollerResponse{}, err
	}
	result := IPAllocationPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("IPAllocationsClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return IPAllocationPollerResponse{}, err
	}
	poller := &ipAllocationPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (IPAllocationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new IPAllocationPoller from the specified resume token.
// token - The value must come from a previous call to IPAllocationPoller.ResumeToken().
func (client IPAllocationsClient) ResumeCreateOrUpdate(token string) (IPAllocationPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("IPAllocationsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &ipAllocationPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Creates or updates an IpAllocation in the specified resource group.
func (client IPAllocationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, ipAllocationName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client IPAllocationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters IPAllocation, options *IPAllocationsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client IPAllocationsClient) createOrUpdateHandleResponse(resp *azcore.Response) (IPAllocationResponse, error) {
	result := IPAllocationResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.IPAllocation)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client IPAllocationsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified IpAllocation.
func (client IPAllocationsClient) BeginDelete(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.delete(ctx, resourceGroupName, ipAllocationName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("IPAllocationsClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client IPAllocationsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("IPAllocationsClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// Delete - Deletes the specified IpAllocation.
func (client IPAllocationsClient) delete(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, ipAllocationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client IPAllocationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client IPAllocationsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified IpAllocation by resource group.
func (client IPAllocationsClient) Get(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsGetOptions) (IPAllocationResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ipAllocationName, options)
	if err != nil {
		return IPAllocationResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return IPAllocationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IPAllocationResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return IPAllocationResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client IPAllocationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, options *IPAllocationsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client IPAllocationsClient) getHandleResponse(resp *azcore.Response) (IPAllocationResponse, error) {
	result := IPAllocationResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.IPAllocation)
	return result, err
}

// getHandleError handles the Get error response.
func (client IPAllocationsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all IpAllocations in a subscription.
func (client IPAllocationsClient) List(options *IPAllocationsListOptions) IPAllocationListResultPager {
	return &ipAllocationListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp IPAllocationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.IPAllocationListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client IPAllocationsClient) listCreateRequest(ctx context.Context, options *IPAllocationsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/IpAllocations"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client IPAllocationsClient) listHandleResponse(resp *azcore.Response) (IPAllocationListResultResponse, error) {
	result := IPAllocationListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.IPAllocationListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client IPAllocationsClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByResourceGroup - Gets all IpAllocations in a resource group.
func (client IPAllocationsClient) ListByResourceGroup(resourceGroupName string, options *IPAllocationsListByResourceGroupOptions) IPAllocationListResultPager {
	return &ipAllocationListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp IPAllocationListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.IPAllocationListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client IPAllocationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *IPAllocationsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client IPAllocationsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (IPAllocationListResultResponse, error) {
	result := IPAllocationListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.IPAllocationListResult)
	return result, err
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client IPAllocationsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates a IpAllocation tags.
func (client IPAllocationsClient) UpdateTags(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters TagsObject, options *IPAllocationsUpdateTagsOptions) (IPAllocationResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, ipAllocationName, parameters, options)
	if err != nil {
		return IPAllocationResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return IPAllocationResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IPAllocationResponse{}, client.updateTagsHandleError(resp)
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return IPAllocationResponse{}, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client IPAllocationsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, ipAllocationName string, parameters TagsObject, options *IPAllocationsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/IpAllocations/{ipAllocationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ipAllocationName}", url.PathEscape(ipAllocationName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client IPAllocationsClient) updateTagsHandleResponse(resp *azcore.Response) (IPAllocationResponse, error) {
	result := IPAllocationResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.IPAllocation)
	return result, err
}

// updateTagsHandleError handles the UpdateTags error response.
func (client IPAllocationsClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
