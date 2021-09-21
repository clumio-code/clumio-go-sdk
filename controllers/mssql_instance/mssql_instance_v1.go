// Copyright (c) 2021 Clumio All Rights Reserved

// Package mssqlinstance contains methods related to MssqlInstance
package mssqlinstance

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// MssqlInstanceV1 represents a custom type struct
type MssqlInstanceV1 struct {
    config config.Config
}

//  ListMssqlInstance Returns a list of Instances
func (m *MssqlInstanceV1) ListMssqlInstance(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListMssqlInstancesResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := m.config.BaseUrl + "/datasources/mssql/instances"

    
    header := "application/mssql-instance=v1+json"
    var result *models.ListMssqlInstancesResponse
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
        SetHeader("Accept", header).
        SetAuthToken(m.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

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


//  ReadMssqlInstance Returns a representation of the specified instance.
func (m *MssqlInstanceV1) ReadMssqlInstance(
    instanceId string)(
    *models.ReadMssqlInstanceResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/mssql/instances/{instance_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "instanceId": instanceId,
    }
    _queryBuilder := m.config.BaseUrl + _pathURL

    
    header := "application/mssql-instance=v1+json"
    var result *models.ReadMssqlInstanceResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(m.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

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
