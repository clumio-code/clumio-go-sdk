// Copyright (c) 2021 Clumio All Rights Reserved

// Package ec2mssqlfailovercluster contains methods related to Ec2MssqlFailoverCluster
package ec2mssqlfailovercluster

import (
    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlFailoverClusterV1 represents a custom type struct
type Ec2MssqlFailoverClusterV1 struct {
    config config.Config
}

// ReadEc2MssqlFailoverCluster Returns a representation of the specified failover cluster.
func (e *Ec2MssqlFailoverClusterV1) ReadEc2MssqlFailoverCluster(
    failoverClusterId string)(
    *models.ReadEC2MSSQLFCIResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/ec2-mssql/failover-clusters/{failover_cluster_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "failover_cluster_id": failoverClusterId,
    }
    queryBuilder := e.config.BaseUrl + pathURL

    
    header := "application/api.clumio.ec2-mssql-failover-cluster=v1+json"
    result := &models.ReadEC2MSSQLFCIResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: e.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
