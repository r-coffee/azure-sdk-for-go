package apimanagement

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

// APIIssueClient is the apiManagement Client
type APIIssueClient struct {
	BaseClient
}

// NewAPIIssueClient creates an instance of the APIIssueClient client.
func NewAPIIssueClient(subscriptionID string) APIIssueClient {
	return NewAPIIssueClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAPIIssueClientWithBaseURI creates an instance of the APIIssueClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewAPIIssueClientWithBaseURI(baseURI string, subscriptionID string) APIIssueClient {
	return APIIssueClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates a new Issue for an API or updates an existing one.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// parameters - create parameters.
// ifMatch - eTag of the Issue Entity. ETag should match the current entity state from the header response of
// the GET request or it should be * for unconditional update.
func (client APIIssueClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, parameters IssueContract, ifMatch string) (result IssueContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.IssueContractProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.IssueContractProperties.Title", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.IssueContractProperties.Description", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.IssueContractProperties.UserID", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serviceName, apiid, issueID, parameters, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client APIIssueClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, parameters IssueContract, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client APIIssueClient) CreateOrUpdateResponder(resp *http.Response) (result IssueContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified Issue from an API.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
// ifMatch - eTag of the Issue Entity. ETag should match the current entity state from the header response of
// the GET request or it should be * for unconditional update.
func (client APIIssueClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, ifMatch string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, serviceName, apiid, issueID, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client APIIssueClient) DeletePreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client APIIssueClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the Issue for an API specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
func (client APIIssueClient) Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string) (result IssueContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, serviceName, apiid, issueID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client APIIssueClient) GetPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client APIIssueClient) GetResponder(resp *http.Response) (result IssueContract, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Head gets the entity state (Etag) version of the Issue for an API specified by its identifier.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// apiid - API identifier. Must be unique in the current API Management service instance.
// issueID - issue identifier. Must be unique in the current API Management service instance.
func (client APIIssueClient) Head(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/APIIssueClient.Head")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: apiid,
			Constraints: []validation.Constraint{{Target: "apiid", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "apiid", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "apiid", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}},
		{TargetValue: issueID,
			Constraints: []validation.Constraint{{Target: "issueID", Name: validation.MaxLength, Rule: 256, Chain: nil},
				{Target: "issueID", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "issueID", Name: validation.Pattern, Rule: `^[^*#&+:<>?]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.APIIssueClient", "Head", err.Error())
	}

	req, err := client.HeadPreparer(ctx, resourceGroupName, serviceName, apiid, issueID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Head", nil, "Failure preparing request")
		return
	}

	resp, err := client.HeadSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Head", resp, "Failure sending request")
		return
	}

	result, err = client.HeadResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.APIIssueClient", "Head", resp, "Failure responding to request")
	}

	return
}

// HeadPreparer prepares the Head request.
func (client APIIssueClient) HeadPreparer(ctx context.Context, resourceGroupName string, serviceName string, apiid string, issueID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiId":             autorest.Encode("path", apiid),
		"issueId":           autorest.Encode("path", issueID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/issues/{issueId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// HeadSender sends the Head request. The method will close the
// http.Response Body if it receives an error.
func (client APIIssueClient) HeadSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// HeadResponder handles the response to the Head request. The method always
// closes the http.Response Body.
func (client APIIssueClient) HeadResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
