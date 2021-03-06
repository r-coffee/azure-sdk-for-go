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
)

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	con *armcore.Connection
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
func NewOperationsClient(con *armcore.Connection) OperationsClient {
	return OperationsClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client OperationsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// List - Lists all of the available Storage Rest API operations.
func (client OperationsClient) List(ctx context.Context, options *OperationsListOptions) (OperationListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return OperationListResultResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return OperationListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return OperationListResultResponse{}, client.listHandleError(resp)
	}
	result, err := client.listHandleResponse(resp)
	if err != nil {
		return OperationListResultResponse{}, err
	}
	return result, nil
}

// listCreateRequest creates the List request.
func (client OperationsClient) listCreateRequest(ctx context.Context, options *OperationsListOptions) (*azcore.Request, error) {
	urlPath := "/providers/Microsoft.Storage/operations"
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

// listHandleResponse handles the List response.
func (client OperationsClient) listHandleResponse(resp *azcore.Response) (OperationListResultResponse, error) {
	result := OperationListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.OperationListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client OperationsClient) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
