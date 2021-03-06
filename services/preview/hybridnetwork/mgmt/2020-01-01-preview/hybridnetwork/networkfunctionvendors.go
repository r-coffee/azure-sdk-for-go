package hybridnetwork

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// NetworkFunctionVendorsClient is the client for the NetworkFunctionVendors methods of the Hybridnetwork service.
type NetworkFunctionVendorsClient struct {
	BaseClient
}

// NewNetworkFunctionVendorsClient creates an instance of the NetworkFunctionVendorsClient client.
func NewNetworkFunctionVendorsClient(subscriptionID string) NetworkFunctionVendorsClient {
	return NewNetworkFunctionVendorsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNetworkFunctionVendorsClientWithBaseURI creates an instance of the NetworkFunctionVendorsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewNetworkFunctionVendorsClientWithBaseURI(baseURI string, subscriptionID string) NetworkFunctionVendorsClient {
	return NetworkFunctionVendorsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all the available vendor and sku information.
func (client NetworkFunctionVendorsClient) List(ctx context.Context) (result NetworkFunctionVendorListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkFunctionVendorsClient.List")
		defer func() {
			sc := -1
			if result.nfvlr.Response.Response != nil {
				sc = result.nfvlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.NetworkFunctionVendorsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.NetworkFunctionVendorsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.nfvlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybridnetwork.NetworkFunctionVendorsClient", "List", resp, "Failure sending request")
		return
	}

	result.nfvlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.NetworkFunctionVendorsClient", "List", resp, "Failure responding to request")
	}
	if result.nfvlr.hasNextLink() && result.nfvlr.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListPreparer prepares the List request.
func (client NetworkFunctionVendorsClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/networkFunctionVendors", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkFunctionVendorsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client NetworkFunctionVendorsClient) ListResponder(resp *http.Response) (result NetworkFunctionVendorListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client NetworkFunctionVendorsClient) listNextResults(ctx context.Context, lastResults NetworkFunctionVendorListResult) (result NetworkFunctionVendorListResult, err error) {
	req, err := lastResults.networkFunctionVendorListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybridnetwork.NetworkFunctionVendorsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybridnetwork.NetworkFunctionVendorsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.NetworkFunctionVendorsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client NetworkFunctionVendorsClient) ListComplete(ctx context.Context) (result NetworkFunctionVendorListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkFunctionVendorsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx)
	return
}
