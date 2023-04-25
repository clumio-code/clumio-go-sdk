// Copyright (c) 2021 Clumio All Rights Reserved

// Package ec2mssqlinstance contains methods related to Ec2MssqlInstance
package ec2mssqlinstance

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlInstanceV1 represents a custom type struct
type Ec2MssqlInstanceV1 struct {
    config config.Config
}

// ListEc2MssqlInstances Returns a list of Instances
func (e *Ec2MssqlInstanceV1) ListEc2MssqlInstances(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListEC2MSSQLInstancesResponse, *apiutils.APIError) {

    queryBuilder := e.config.BaseUrl + "/datasources/aws/ec2-mssql/instances"

    
    header := "application/api.clumio.ec2-mssql-instance=v1+json"
    result := &models.ListEC2MSSQLInstancesResponse{}
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
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
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


// ReadEc2MssqlInstance Returns a representation of the specified instance.
func (e *Ec2MssqlInstanceV1) ReadEc2MssqlInstance(
    instanceId string)(
    *models.ReadEC2MSSQLInstanceResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/ec2-mssql/instances/{instance_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "instance_id": instanceId,
    }
    queryBuilder := e.config.BaseUrl + pathURL

    
    header := "application/api.clumio.ec2-mssql-instance=v1+json"
    result := &models.ReadEC2MSSQLInstanceResponse{}

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
