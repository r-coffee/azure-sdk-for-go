package kustoapi

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
	"github.com/Azure/azure-sdk-for-go/services/kusto/mgmt/2019-05-15/kusto"
)

// ClustersClientAPI contains the set of methods on the ClustersClient type.
type ClustersClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, clusterName kusto.ClusterCheckNameRequest) (result kusto.CheckNameResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters kusto.Cluster) (result kusto.ClustersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ClustersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.Cluster, err error)
	List(ctx context.Context) (result kusto.ClusterListResult, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result kusto.ClusterListResult, err error)
	ListSkus(ctx context.Context) (result kusto.SkuDescriptionList, err error)
	ListSkusByResource(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ListResourceSkusResult, err error)
	Start(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ClustersStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.ClustersStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, parameters kusto.ClusterUpdate) (result kusto.ClustersUpdateFuture, err error)
}

var _ ClustersClientAPI = (*kusto.ClustersClient)(nil)

// DatabasesClientAPI contains the set of methods on the DatabasesClient type.
type DatabasesClientAPI interface {
	AddPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToAdd kusto.DatabasePrincipalListRequest) (result kusto.DatabasePrincipalListResult, err error)
	CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, databaseName kusto.DatabaseCheckNameRequest) (result kusto.CheckNameResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters kusto.Database) (result kusto.DatabasesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.DatabasesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.Database, err error)
	ListByCluster(ctx context.Context, resourceGroupName string, clusterName string) (result kusto.DatabaseListResult, err error)
	ListPrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.DatabasePrincipalListResult, err error)
	RemovePrincipals(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, databasePrincipalsToRemove kusto.DatabasePrincipalListRequest) (result kusto.DatabasePrincipalListResult, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters kusto.DatabaseUpdate) (result kusto.DatabasesUpdateFuture, err error)
}

var _ DatabasesClientAPI = (*kusto.DatabasesClient)(nil)

// DataConnectionsClientAPI contains the set of methods on the DataConnectionsClient type.
type DataConnectionsClientAPI interface {
	CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName kusto.DataConnectionCheckNameRequest) (result kusto.CheckNameResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters kusto.BasicDataConnection) (result kusto.DataConnectionsCreateOrUpdateFuture, err error)
	DataConnectionValidationMethod(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters kusto.DataConnectionValidation) (result kusto.DataConnectionValidationListResult, err error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (result kusto.DataConnectionsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (result kusto.DataConnectionModel, err error)
	ListByDatabase(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result kusto.DataConnectionListResult, err error)
	Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters kusto.BasicDataConnection) (result kusto.DataConnectionsUpdateFuture, err error)
}

var _ DataConnectionsClientAPI = (*kusto.DataConnectionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result kusto.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result kusto.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*kusto.OperationsClient)(nil)
