// Copyright (c) 2023 Clumio All Rights Reserved

// Package mssqldatabases contains methods related to MssqlDatabases
package mssqldatabases

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// MssqlDatabasesV1 represents a custom type struct
type MssqlDatabasesV1 struct {
    config config.Config
}

// ListMssqlDatabases Returns a list of Databases
func (m *MssqlDatabasesV1) ListMssqlDatabases(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListMssqlDatabasesResponse, *apiutils.APIError) {

    queryBuilder := m.config.BaseUrl + "/datasources/mssql/databases"

    
    header := "application/api.clumio.mssql-databases=v1+json"
    result := &models.ListMssqlDatabasesResponse{}
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
        Config: m.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadMssqlDatabases Returns a representation of the specified database.
func (m *MssqlDatabasesV1) ReadMssqlDatabases(
    databaseId string)(
    *models.ReadMssqlDatabaseResponse, *apiutils.APIError) {

    pathURL := "/datasources/mssql/databases/{database_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "database_id": databaseId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    
    header := "application/api.clumio.mssql-databases=v1+json"
    result := &models.ReadMssqlDatabaseResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: m.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ListMssqlDatabasePitrIntervals Returns restorable times as a list of intervals.
func (m *MssqlDatabasesV1) ListMssqlDatabasePitrIntervals(
    databaseId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListMssqlDatabasePitrIntervalsResponse, *apiutils.APIError) {

    pathURL := "/datasources/mssql/databases/{database_id}/pitr-intervals"
    //process optional template parameters
    pathParams := map[string]string{
        "database_id": databaseId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    
    header := "application/api.clumio.mssql-databases=v1+json"
    result := &models.ListMssqlDatabasePitrIntervalsResponse{}
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
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: m.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
