// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package aznetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// SecurityPartnerProvidersOperations contains the methods for the SecurityPartnerProviders group.
type SecurityPartnerProvidersOperations interface {
	// BeginCreateOrUpdate - Creates or updates the specified Security Partner Provider.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters SecurityPartnerProvider) (*SecurityPartnerProviderResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (SecurityPartnerProviderPoller, error)
	// BeginDelete - Deletes the specified Security Partner Provider.
	BeginDelete(ctx context.Context, resourceGroupName string, securityPartnerProviderName string) (*HTTPResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified Security Partner Provider.
	Get(ctx context.Context, resourceGroupName string, securityPartnerProviderName string) (*SecurityPartnerProviderResponse, error)
	// List - Gets all the Security Partner Providers in a subscription.
	List() (SecurityPartnerProviderListResultPager, error)
	// ListByResourceGroup - Lists all Security Partner Providers in a resource group.
	ListByResourceGroup(resourceGroupName string) (SecurityPartnerProviderListResultPager, error)
	// UpdateTags - Updates tags of a Security Partner Provider resource.
	UpdateTags(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters TagsObject) (*SecurityPartnerProviderResponse, error)
}

// securityPartnerProvidersOperations implements the SecurityPartnerProvidersOperations interface.
type securityPartnerProvidersOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates the specified Security Partner Provider.
func (client *securityPartnerProvidersOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters SecurityPartnerProvider) (*SecurityPartnerProviderResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, securityPartnerProviderName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("securityPartnerProvidersOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &securityPartnerProviderPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*SecurityPartnerProviderResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *securityPartnerProvidersOperations) ResumeCreateOrUpdate(token string) (SecurityPartnerProviderPoller, error) {
	pt, err := resumePollingTracker("securityPartnerProvidersOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &securityPartnerProviderPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *securityPartnerProvidersOperations) createOrUpdateCreateRequest(resourceGroupName string, securityPartnerProviderName string, parameters SecurityPartnerProvider) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *securityPartnerProvidersOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*SecurityPartnerProviderResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := SecurityPartnerProviderResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SecurityPartnerProvider)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *securityPartnerProvidersOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified Security Partner Provider.
func (client *securityPartnerProvidersOperations) BeginDelete(ctx context.Context, resourceGroupName string, securityPartnerProviderName string) (*HTTPResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, securityPartnerProviderName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("securityPartnerProvidersOperations.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *securityPartnerProvidersOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("securityPartnerProvidersOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *securityPartnerProvidersOperations) deleteCreateRequest(resourceGroupName string, securityPartnerProviderName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *securityPartnerProvidersOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *securityPartnerProvidersOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified Security Partner Provider.
func (client *securityPartnerProvidersOperations) Get(ctx context.Context, resourceGroupName string, securityPartnerProviderName string) (*SecurityPartnerProviderResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, securityPartnerProviderName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *securityPartnerProvidersOperations) getCreateRequest(resourceGroupName string, securityPartnerProviderName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *securityPartnerProvidersOperations) getHandleResponse(resp *azcore.Response) (*SecurityPartnerProviderResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := SecurityPartnerProviderResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SecurityPartnerProvider)
}

// getHandleError handles the Get error response.
func (client *securityPartnerProvidersOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all the Security Partner Providers in a subscription.
func (client *securityPartnerProvidersOperations) List() (SecurityPartnerProviderListResultPager, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	return &securityPartnerProviderListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *SecurityPartnerProviderListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.SecurityPartnerProviderListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.SecurityPartnerProviderListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *securityPartnerProvidersOperations) listCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/securityPartnerProviders"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *securityPartnerProvidersOperations) listHandleResponse(resp *azcore.Response) (*SecurityPartnerProviderListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := SecurityPartnerProviderListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SecurityPartnerProviderListResult)
}

// listHandleError handles the List error response.
func (client *securityPartnerProvidersOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Lists all Security Partner Providers in a resource group.
func (client *securityPartnerProvidersOperations) ListByResourceGroup(resourceGroupName string) (SecurityPartnerProviderListResultPager, error) {
	req, err := client.listByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &securityPartnerProviderListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByResourceGroupHandleResponse,
		advancer: func(resp *SecurityPartnerProviderListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.SecurityPartnerProviderListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.SecurityPartnerProviderListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *securityPartnerProvidersOperations) listByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *securityPartnerProvidersOperations) listByResourceGroupHandleResponse(resp *azcore.Response) (*SecurityPartnerProviderListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByResourceGroupHandleError(resp)
	}
	result := SecurityPartnerProviderListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SecurityPartnerProviderListResult)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *securityPartnerProvidersOperations) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates tags of a Security Partner Provider resource.
func (client *securityPartnerProvidersOperations) UpdateTags(ctx context.Context, resourceGroupName string, securityPartnerProviderName string, parameters TagsObject) (*SecurityPartnerProviderResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, securityPartnerProviderName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *securityPartnerProvidersOperations) updateTagsCreateRequest(resourceGroupName string, securityPartnerProviderName string, parameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/securityPartnerProviders/{securityPartnerProviderName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{securityPartnerProviderName}", url.PathEscape(securityPartnerProviderName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *securityPartnerProvidersOperations) updateTagsHandleResponse(resp *azcore.Response) (*SecurityPartnerProviderResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := SecurityPartnerProviderResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SecurityPartnerProvider)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *securityPartnerProvidersOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}