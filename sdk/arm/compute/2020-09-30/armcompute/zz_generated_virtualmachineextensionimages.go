// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// VirtualMachineExtensionImagesClient contains the methods for the VirtualMachineExtensionImages group.
// Don't use this type directly, use NewVirtualMachineExtensionImagesClient() instead.
type VirtualMachineExtensionImagesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineExtensionImagesClient creates a new instance of VirtualMachineExtensionImagesClient with the specified values.
func NewVirtualMachineExtensionImagesClient(con *armcore.Connection, subscriptionID string) VirtualMachineExtensionImagesClient {
	return VirtualMachineExtensionImagesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VirtualMachineExtensionImagesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Gets a virtual machine extension image.
func (client VirtualMachineExtensionImagesClient) Get(ctx context.Context, location string, publisherName string, typeParameter string, version string, options *VirtualMachineExtensionImagesGetOptions) (VirtualMachineExtensionImageResponse, error) {
	req, err := client.getCreateRequest(ctx, location, publisherName, typeParameter, version, options)
	if err != nil {
		return VirtualMachineExtensionImageResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineExtensionImageResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineExtensionImageResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return VirtualMachineExtensionImageResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client VirtualMachineExtensionImagesClient) getCreateRequest(ctx context.Context, location string, publisherName string, typeParameter string, version string, options *VirtualMachineExtensionImagesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions/{version}"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	urlPath = strings.ReplaceAll(urlPath, "{type}", url.PathEscape(typeParameter))
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client VirtualMachineExtensionImagesClient) getHandleResponse(resp *azcore.Response) (VirtualMachineExtensionImageResponse, error) {
	result := VirtualMachineExtensionImageResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineExtensionImage)
	return result, err
}

// getHandleError handles the Get error response.
func (client VirtualMachineExtensionImagesClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListTypes - Gets a list of virtual machine extension image types.
func (client VirtualMachineExtensionImagesClient) ListTypes(ctx context.Context, location string, publisherName string, options *VirtualMachineExtensionImagesListTypesOptions) (VirtualMachineExtensionImageArrayResponse, error) {
	req, err := client.listTypesCreateRequest(ctx, location, publisherName, options)
	if err != nil {
		return VirtualMachineExtensionImageArrayResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineExtensionImageArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineExtensionImageArrayResponse{}, client.listTypesHandleError(resp)
	}
	result, err := client.listTypesHandleResponse(resp)
	if err != nil {
		return VirtualMachineExtensionImageArrayResponse{}, err
	}
	return result, nil
}

// listTypesCreateRequest creates the ListTypes request.
func (client VirtualMachineExtensionImagesClient) listTypesCreateRequest(ctx context.Context, location string, publisherName string, options *VirtualMachineExtensionImagesListTypesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listTypesHandleResponse handles the ListTypes response.
func (client VirtualMachineExtensionImagesClient) listTypesHandleResponse(resp *azcore.Response) (VirtualMachineExtensionImageArrayResponse, error) {
	result := VirtualMachineExtensionImageArrayResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineExtensionImageArray)
	return result, err
}

// listTypesHandleError handles the ListTypes error response.
func (client VirtualMachineExtensionImagesClient) listTypesHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListVersions - Gets a list of virtual machine extension image versions.
func (client VirtualMachineExtensionImagesClient) ListVersions(ctx context.Context, location string, publisherName string, typeParameter string, options *VirtualMachineExtensionImagesListVersionsOptions) (VirtualMachineExtensionImageArrayResponse, error) {
	req, err := client.listVersionsCreateRequest(ctx, location, publisherName, typeParameter, options)
	if err != nil {
		return VirtualMachineExtensionImageArrayResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineExtensionImageArrayResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineExtensionImageArrayResponse{}, client.listVersionsHandleError(resp)
	}
	result, err := client.listVersionsHandleResponse(resp)
	if err != nil {
		return VirtualMachineExtensionImageArrayResponse{}, err
	}
	return result, nil
}

// listVersionsCreateRequest creates the ListVersions request.
func (client VirtualMachineExtensionImagesClient) listVersionsCreateRequest(ctx context.Context, location string, publisherName string, typeParameter string, options *VirtualMachineExtensionImagesListVersionsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/publishers/{publisherName}/artifacttypes/vmextension/types/{type}/versions"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	urlPath = strings.ReplaceAll(urlPath, "{type}", url.PathEscape(typeParameter))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Filter != nil {
		query.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		query.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		query.Set("$orderby", *options.Orderby)
	}
	query.Set("api-version", "2020-06-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listVersionsHandleResponse handles the ListVersions response.
func (client VirtualMachineExtensionImagesClient) listVersionsHandleResponse(resp *azcore.Response) (VirtualMachineExtensionImageArrayResponse, error) {
	result := VirtualMachineExtensionImageArrayResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VirtualMachineExtensionImageArray)
	return result, err
}

// listVersionsHandleError handles the ListVersions error response.
func (client VirtualMachineExtensionImagesClient) listVersionsHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
