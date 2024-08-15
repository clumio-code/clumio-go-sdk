// Copyright (c) 2023 Clumio All Rights Reserved

// Package mssqlhosts contains methods related to MssqlHosts
package mssqlhosts

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// MssqlHostsV1 represents a custom type struct
type MssqlHostsV1 struct {
    config config.Config
}

// ListMssqlHostConnections Returns a list of hosts
func (m *MssqlHostsV1) ListMssqlHostConnections(
    currentCount *int64, 
    filter *string, 
    limit *int64, 
    start *string)(
    *models.ListHcmHostsResponse, *apiutils.APIError) {

    queryBuilder := m.config.BaseUrl + "/connections/mssql/hosts"

    
    header := "application/api.clumio.mssql-hosts=v1+json"
    result := &models.ListHcmHostsResponse{}
    queryParams := make(map[string]string)
    if currentCount != nil {
        queryParams["current_count"] = fmt.Sprintf("%v", *currentCount)
    }
    if filter != nil {
        queryParams["filter"] = *filter
    }
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
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


// CreateMssqlHostConnections Create a MSSQL Connection.
func (m *MssqlHostsV1) CreateMssqlHostConnections(
    body *models.CreateMssqlHostConnectionsV1Request)(
    *models.CreateHcmHostResponse, *apiutils.APIError) {

    queryBuilder := m.config.BaseUrl + "/connections/mssql/hosts"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.mssql-hosts=v1+json"
    result := &models.CreateHcmHostResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: m.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// DeleteMssqlHostConnections Delete the specified MSSQL host.
func (m *MssqlHostsV1) DeleteMssqlHostConnections(
    embed *string, 
    body *models.DeleteMssqlHostConnectionsV1Request)(
    *models.DeleteHcmHostResponse, *apiutils.APIError) {

    queryBuilder := m.config.BaseUrl + "/connections/mssql/hosts"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.mssql-hosts=v1+json"
    result := &models.DeleteHcmHostResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: m.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Delete,
    })

    return result, apiErr
}


// MoveMssqlHostConnections Move the specified MSSQL hosts from a source Sub-Group to a destination Sub-Group.
func (m *MssqlHostsV1) MoveMssqlHostConnections(
    embed *string, 
    body *models.MoveMssqlHostConnectionsV1Request)(
    *models.MoveHcmHostsResponse, *apiutils.APIError) {

    queryBuilder := m.config.BaseUrl + "/connections/mssql/hosts"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.mssql-hosts=v1+json"
    result := &models.MoveHcmHostsResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: m.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Patch,
    })

    return result, apiErr
}


// CreateMssqlHostConnectionCredentials Create Edge Connector Credentials for the specified MSSQL host.
func (m *MssqlHostsV1) CreateMssqlHostConnectionCredentials(
    body *models.CreateMssqlHostConnectionCredentialsV1Request)(
    *models.CreateHostECCredentialsResponse, *apiutils.APIError) {

    queryBuilder := m.config.BaseUrl + "/connections/mssql/hosts/eccredentials"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.mssql-hosts=v1+json"
    result := &models.CreateHostECCredentialsResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: m.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadMssqlHostConnections Returns a representation of the specified host.
func (m *MssqlHostsV1) ReadMssqlHostConnections(
    hostId string)(
    *models.ReadHcmHostResponse, *apiutils.APIError) {

    pathURL := "/connections/mssql/hosts/{host_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "host_id": hostId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    
    header := "application/api.clumio.mssql-hosts=v1+json"
    result := &models.ReadHcmHostResponse{}

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


// ListMssqlHosts Returns a list of hosts
func (m *MssqlHostsV1) ListMssqlHosts(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListMssqlHostsResponse, *apiutils.APIError) {

    queryBuilder := m.config.BaseUrl + "/datasources/mssql/hosts"

    
    header := "application/api.clumio.mssql-hosts=v1+json"
    result := &models.ListMssqlHostsResponse{}
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


// ReadMssqlHosts Returns a representation of the specified host.
func (m *MssqlHostsV1) ReadMssqlHosts(
    hostId string)(
    *models.ReadMssqlHostResponse, *apiutils.APIError) {

    pathURL := "/datasources/mssql/hosts/{host_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "host_id": hostId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    
    header := "application/api.clumio.mssql-hosts=v1+json"
    result := &models.ReadMssqlHostResponse{}

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
