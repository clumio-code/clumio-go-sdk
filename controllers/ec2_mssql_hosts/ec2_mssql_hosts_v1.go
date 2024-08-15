// Copyright (c) 2023 Clumio All Rights Reserved

// Package ec2mssqlhosts contains methods related to Ec2MssqlHosts
package ec2mssqlhosts

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlHostsV1 represents a custom type struct
type Ec2MssqlHostsV1 struct {
    config config.Config
}

// ListEc2MssqlHosts Returns a list of EC2 MSSQL hosts
func (e *Ec2MssqlHostsV1) ListEc2MssqlHosts(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListEC2MSSQLInvHostsResponse, *apiutils.APIError) {

    queryBuilder := e.config.BaseUrl + "/datasources/aws/ec2-mssql/hosts"

    
    header := "application/api.clumio.ec2-mssql-hosts=v1+json"
    result := &models.ListEC2MSSQLInvHostsResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
    }
    if filter != nil {
        queryParams["filter"] = *filter
    }
    if embed != nil {
        queryParams["embed"] = *embed
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


// ReadEc2MssqlHost Returns a representation of the specified host.
func (e *Ec2MssqlHostsV1) ReadEc2MssqlHost(
    hostId string)(
    *models.ReadEC2MSSQLInvHostResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/ec2-mssql/hosts/{host_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "host_id": hostId,
    }
    queryBuilder := e.config.BaseUrl + pathURL

    
    header := "application/api.clumio.ec2-mssql-hosts=v1+json"
    result := &models.ReadEC2MSSQLInvHostResponse{}

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
