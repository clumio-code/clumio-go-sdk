// Copyright (c) 2021 Clumio All Rights Reserved

// Package mssqlhosts contains methods related to MssqlHosts
package mssqlhosts

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// MssqlHostsV1 represents a custom type struct
type MssqlHostsV1 struct {
    config config.Config
}

//  ListMssqlHostConnections Returns a list of hosts
func (m *MssqlHostsV1) ListMssqlHostConnections(
    currentCount *int64, 
    filter *string, 
    limit *int64, 
    start *string)(
    *models.ListHcmHostsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := m.config.BaseUrl + "/connections/mssql/hosts"

    
    header := "application/mssql-hosts=v1+json"
    var result *models.ListHcmHostsResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if currentCount == nil{
        currentCount = &defaultInt64
    }
    if filter == nil{
        filter = &defaultString
    }
    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    

    queryParams := map[string]string{
        "currentCount": fmt.Sprintf("%v", *currentCount),
        "filter": *filter,
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
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


//  CreateMssqlHostConnections Create a MSSQL Connection.
func (m *MssqlHostsV1) CreateMssqlHostConnections(
    body *models.CreateMssqlHostConnectionsV1Request)(
    *models.CreateHcmHostResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := m.config.BaseUrl + "/connections/mssql/hosts"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/mssql-hosts=v1+json"
    var result *models.CreateHcmHostResponse
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetAuthToken(m.config.Token).
        SetBody(payload).
        SetResult(&result).
        Post(_queryBuilder)

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


//  DeleteMssqlHostConnections Delete the specified MSSQL host.
func (m *MssqlHostsV1) DeleteMssqlHostConnections(
    embed *string, 
    body *models.DeleteMssqlHostConnectionsV1Request)(
    *models.DeleteHcmHostResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := m.config.BaseUrl + "/connections/mssql/hosts"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/mssql-hosts=v1+json"
    var result *models.DeleteHcmHostResponse
    client := resty.New()
    defaultString := "" 
    

    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetAuthToken(m.config.Token).
        SetBody(payload).
        SetResult(&result).
        Delete(_queryBuilder)

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


//  MoveMssqlHostConnections Move the specified MSSQL hosts from a source Sub-Group to a destination Sub-Group.
func (m *MssqlHostsV1) MoveMssqlHostConnections(
    embed *string, 
    body *models.MoveMssqlHostConnectionsV1Request)(
    *models.MoveHcmHostsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := m.config.BaseUrl + "/connections/mssql/hosts"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/mssql-hosts=v1+json"
    var result *models.MoveHcmHostsResponse
    client := resty.New()
    defaultString := "" 
    

    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetAuthToken(m.config.Token).
        SetBody(payload).
        SetResult(&result).
        Patch(_queryBuilder)

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


//  CreateMssqlHostConnectionCredentials Create Edge Connector Credentials for the specified MSSQL host.
func (m *MssqlHostsV1) CreateMssqlHostConnectionCredentials(
    body *models.CreateMssqlHostConnectionCredentialsV1Request)(
    *models.CreateHostECCredentialsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := m.config.BaseUrl + "/connections/mssql/hosts/eccredentials"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/mssql-hosts=v1+json"
    var result *models.CreateHostECCredentialsResponse
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetAuthToken(m.config.Token).
        SetBody(payload).
        SetResult(&result).
        Post(_queryBuilder)

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


//  ReadMssqlHostConnections Returns a representation of the specified host.
func (m *MssqlHostsV1) ReadMssqlHostConnections(
    hostId string)(
    *models.ReadHcmHostResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/connections/mssql/hosts/{host_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "hostId": hostId,
    }
    _queryBuilder := m.config.BaseUrl + _pathURL

    
    header := "application/mssql-hosts=v1+json"
    var result *models.ReadHcmHostResponse
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


//  ListMssqlHosts Returns a list of hosts
func (m *MssqlHostsV1) ListMssqlHosts(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListMssqlHostsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := m.config.BaseUrl + "/datasources/mssql/hosts"

    
    header := "application/mssql-hosts=v1+json"
    var result *models.ListMssqlHostsResponse
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


//  ReadMssqlHosts Returns a representation of the specified host.
func (m *MssqlHostsV1) ReadMssqlHosts(
    hostId string)(
    *models.ReadMssqlHostResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/mssql/hosts/{host_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "hostId": hostId,
    }
    _queryBuilder := m.config.BaseUrl + _pathURL

    
    header := "application/mssql-hosts=v1+json"
    var result *models.ReadMssqlHostResponse
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
