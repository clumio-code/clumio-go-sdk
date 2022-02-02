// Copyright (c) 2021 Clumio All Rights Reserved

// Package mssqldatabases contains methods related to MssqlDatabases
package mssqldatabases

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// MssqlDatabasesV1 represents a custom type struct
type MssqlDatabasesV1 struct {
    config config.Config
}

//  ListMssqlDatabases Returns a list of Databases
func (m *MssqlDatabasesV1) ListMssqlDatabases(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListMssqlDatabasesResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := m.config.BaseUrl + "/datasources/mssql/databases"

    
    header := "application/mssql-databases=v1+json"
    var result *models.ListMssqlDatabasesResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    if filter == nil{
        filter = &defaultString
    }
    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetHeader("x-clumio-organizationalunit-context", m.config.OrganizationalUnitContext).
        SetAuthToken(m.config.Token).
        SetResult(&result).
        Get(queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}


//  ReadMssqlDatabases Returns a representation of the specified database.
func (m *MssqlDatabasesV1) ReadMssqlDatabases(
    databaseId string)(
    *models.ReadMssqlDatabaseResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/mssql/databases/{database_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "database_id": databaseId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    
    header := "application/mssql-databases=v1+json"
    var result *models.ReadMssqlDatabaseResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetHeader("x-clumio-organizationalunit-context", m.config.OrganizationalUnitContext).
        SetAuthToken(m.config.Token).
        SetResult(&result).
        Get(queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}


//  ListMssqlDatabasePitrIntervals Returns restorable times as a list of intervals.
func (m *MssqlDatabasesV1) ListMssqlDatabasePitrIntervals(
    databaseId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListMssqlDatabasePitrIntervalsResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/mssql/databases/{database_id}/pitr-intervals"
    //process optional template parameters
    pathParams := map[string]string{
        "database_id": databaseId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    
    header := "application/mssql-databases=v1+json"
    var result *models.ListMssqlDatabasePitrIntervalsResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    if filter == nil{
        filter = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetHeader("x-clumio-organizationalunit-context", m.config.OrganizationalUnitContext).
        SetAuthToken(m.config.Token).
        SetResult(&result).
        Get(queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}
