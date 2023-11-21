// Copyright (c) 2023 Clumio All Rights Reserved

// Package ec2mssqldatabases contains methods related to Ec2MssqlDatabases
package ec2mssqldatabases

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlDatabasesV1 represents a custom type struct
type Ec2MssqlDatabasesV1 struct {
    config config.Config
}

// ListEc2MssqlDatabases Returns a list of Databases
func (e *Ec2MssqlDatabasesV1) ListEc2MssqlDatabases(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListEC2MSSQLDatabasesResponse, *apiutils.APIError) {

    queryBuilder := e.config.BaseUrl + "/datasources/aws/ec2-mssql/databases"

    
    header := "application/api.clumio.ec2-mssql-databases=v1+json"
    result := &models.ListEC2MSSQLDatabasesResponse{}
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


// ReadEc2MssqlDatabase Returns a representation of the specified database.
func (e *Ec2MssqlDatabasesV1) ReadEc2MssqlDatabase(
    databaseId string)(
    *models.ReadEC2MSSQLDatabaseResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/ec2-mssql/databases/{database_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "database_id": databaseId,
    }
    queryBuilder := e.config.BaseUrl + pathURL

    
    header := "application/api.clumio.ec2-mssql-databases=v1+json"
    result := &models.ReadEC2MSSQLDatabaseResponse{}

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


// ListEc2MssqlDatabasePitrIntervals Returns a list of time intervals (start timestamp and end timestamp) in which the MSSQL database can be restored.
func (e *Ec2MssqlDatabasesV1) ListEc2MssqlDatabasePitrIntervals(
    databaseId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListEC2MssqlDatabasePitrIntervalsResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/ec2-mssql/databases/{database_id}/pitr-intervals"
    //process optional template parameters
    pathParams := map[string]string{
        "database_id": databaseId,
    }
    queryBuilder := e.config.BaseUrl + pathURL

    
    header := "application/api.clumio.ec2-mssql-databases=v1+json"
    result := &models.ListEC2MssqlDatabasePitrIntervalsResponse{}
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
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
