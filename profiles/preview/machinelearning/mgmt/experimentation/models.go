// +build go1.9

// Copyright 2019 Microsoft Corporation
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

package experimentation

import original "github.com/Azure/azure-sdk-for-go/services/machinelearning/mgmt/2017-05-01-preview/experimentation"

type AccountsClient = original.AccountsClient

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type ProvisioningState = original.ProvisioningState

const (
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type Account = original.Account
type AccountListResult = original.AccountListResult
type AccountListResultIterator = original.AccountListResultIterator
type AccountListResultPage = original.AccountListResultPage
type AccountProperties = original.AccountProperties
type AccountPropertiesUpdateParameters = original.AccountPropertiesUpdateParameters
type AccountUpdateParameters = original.AccountUpdateParameters
type ErrorResponse = original.ErrorResponse
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type Project = original.Project
type ProjectListResult = original.ProjectListResult
type ProjectListResultIterator = original.ProjectListResultIterator
type ProjectListResultPage = original.ProjectListResultPage
type ProjectProperties = original.ProjectProperties
type ProjectPropertiesUpdateParameters = original.ProjectPropertiesUpdateParameters
type ProjectUpdateParameters = original.ProjectUpdateParameters
type Resource = original.Resource
type StorageAccountProperties = original.StorageAccountProperties
type Workspace = original.Workspace
type WorkspaceListResult = original.WorkspaceListResult
type WorkspaceListResultIterator = original.WorkspaceListResultIterator
type WorkspaceListResultPage = original.WorkspaceListResultPage
type WorkspaceProperties = original.WorkspaceProperties
type WorkspacePropertiesUpdateParameters = original.WorkspacePropertiesUpdateParameters
type WorkspaceUpdateParameters = original.WorkspaceUpdateParameters
type OperationsClient = original.OperationsClient
type ProjectsClient = original.ProjectsClient
type WorkspacesClient = original.WorkspacesClient

func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProjectsClient(subscriptionID string) ProjectsClient {
	return original.NewProjectsClient(subscriptionID)
}
func NewProjectsClientWithBaseURI(baseURI string, subscriptionID string) ProjectsClient {
	return original.NewProjectsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}