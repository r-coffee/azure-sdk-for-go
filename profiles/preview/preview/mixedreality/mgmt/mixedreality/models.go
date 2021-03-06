// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package mixedreality

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/mixedreality/mgmt/2020-05-01-preview/mixedreality"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type NameAvailability = original.NameAvailability

const (
	False NameAvailability = original.False
	True  NameAvailability = original.True
)

type NameUnavailableReason = original.NameUnavailableReason

const (
	AlreadyExists NameUnavailableReason = original.AlreadyExists
	Invalid       NameUnavailableReason = original.Invalid
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type SkuTier = original.SkuTier

const (
	Basic    SkuTier = original.Basic
	Free     SkuTier = original.Free
	Premium  SkuTier = original.Premium
	Standard SkuTier = original.Standard
)

type AccountKeyRegenerateRequest = original.AccountKeyRegenerateRequest
type AccountKeys = original.AccountKeys
type AccountProperties = original.AccountProperties
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResponse = original.CheckNameAvailabilityResponse
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type Identity = original.Identity
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationPage = original.OperationPage
type OperationPageIterator = original.OperationPageIterator
type OperationPagePage = original.OperationPagePage
type OperationsClient = original.OperationsClient
type Plan = original.Plan
type ProxyResource = original.ProxyResource
type RemoteRenderingAccount = original.RemoteRenderingAccount
type RemoteRenderingAccountIdentity = original.RemoteRenderingAccountIdentity
type RemoteRenderingAccountPage = original.RemoteRenderingAccountPage
type RemoteRenderingAccountPageIterator = original.RemoteRenderingAccountPageIterator
type RemoteRenderingAccountPagePage = original.RemoteRenderingAccountPagePage
type RemoteRenderingAccountsClient = original.RemoteRenderingAccountsClient
type Resource = original.Resource
type ResourceModelWithAllowedPropertySet = original.ResourceModelWithAllowedPropertySet
type ResourceModelWithAllowedPropertySetIdentity = original.ResourceModelWithAllowedPropertySetIdentity
type ResourceModelWithAllowedPropertySetPlan = original.ResourceModelWithAllowedPropertySetPlan
type ResourceModelWithAllowedPropertySetSku = original.ResourceModelWithAllowedPropertySetSku
type Sku = original.Sku
type SpatialAnchorsAccount = original.SpatialAnchorsAccount
type SpatialAnchorsAccountPage = original.SpatialAnchorsAccountPage
type SpatialAnchorsAccountPageIterator = original.SpatialAnchorsAccountPageIterator
type SpatialAnchorsAccountPagePage = original.SpatialAnchorsAccountPagePage
type SpatialAnchorsAccountsClient = original.SpatialAnchorsAccountsClient
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationPageIterator(page OperationPagePage) OperationPageIterator {
	return original.NewOperationPageIterator(page)
}
func NewOperationPagePage(cur OperationPage, getNextPage func(context.Context, OperationPage) (OperationPage, error)) OperationPagePage {
	return original.NewOperationPagePage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewRemoteRenderingAccountPageIterator(page RemoteRenderingAccountPagePage) RemoteRenderingAccountPageIterator {
	return original.NewRemoteRenderingAccountPageIterator(page)
}
func NewRemoteRenderingAccountPagePage(cur RemoteRenderingAccountPage, getNextPage func(context.Context, RemoteRenderingAccountPage) (RemoteRenderingAccountPage, error)) RemoteRenderingAccountPagePage {
	return original.NewRemoteRenderingAccountPagePage(cur, getNextPage)
}
func NewRemoteRenderingAccountsClient(subscriptionID string) RemoteRenderingAccountsClient {
	return original.NewRemoteRenderingAccountsClient(subscriptionID)
}
func NewRemoteRenderingAccountsClientWithBaseURI(baseURI string, subscriptionID string) RemoteRenderingAccountsClient {
	return original.NewRemoteRenderingAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSpatialAnchorsAccountPageIterator(page SpatialAnchorsAccountPagePage) SpatialAnchorsAccountPageIterator {
	return original.NewSpatialAnchorsAccountPageIterator(page)
}
func NewSpatialAnchorsAccountPagePage(cur SpatialAnchorsAccountPage, getNextPage func(context.Context, SpatialAnchorsAccountPage) (SpatialAnchorsAccountPage, error)) SpatialAnchorsAccountPagePage {
	return original.NewSpatialAnchorsAccountPagePage(cur, getNextPage)
}
func NewSpatialAnchorsAccountsClient(subscriptionID string) SpatialAnchorsAccountsClient {
	return original.NewSpatialAnchorsAccountsClient(subscriptionID)
}
func NewSpatialAnchorsAccountsClientWithBaseURI(baseURI string, subscriptionID string) SpatialAnchorsAccountsClient {
	return original.NewSpatialAnchorsAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleNameAvailabilityValues() []NameAvailability {
	return original.PossibleNameAvailabilityValues()
}
func PossibleNameUnavailableReasonValues() []NameUnavailableReason {
	return original.PossibleNameUnavailableReasonValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
