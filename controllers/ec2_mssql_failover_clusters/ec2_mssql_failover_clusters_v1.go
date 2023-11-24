// Copyright (c) 2023 Clumio All Rights Reserved

// Package ec2mssqlfailoverclusters contains methods related to Ec2MssqlFailoverClusters
package ec2mssqlfailoverclusters

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlFailoverClustersV1 represents a custom type struct
type Ec2MssqlFailoverClustersV1 struct {
    config config.Config
}

// ListEc2MssqlFailoverClusters Returns a list of failover clusters.
func (e *Ec2MssqlFailoverClustersV1) ListEc2MssqlFailoverClusters(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListEC2MSSQLFCIsResponse, *apiutils.APIError) {

    queryBuilder := e.config.BaseUrl + "/datasources/aws/ec2-mssql/failover-clusters"

    
    header := "application/api.clumio.ec2-mssql-failover-clusters=v1+json"
    result := &models.ListEC2MSSQLFCIsResponse{}
    defaultInt64 := int64(0)
    defaultString := "" 
    
    if limit == nil {
        limit = &defaultInt64
    }
    if start == nil {
        start = &defaultString
    }
    if filter == nil {
        filter = &defaultString
    }
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
        "embed": *embed,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: e.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
