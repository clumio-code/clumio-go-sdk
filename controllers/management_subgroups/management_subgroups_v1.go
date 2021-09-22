// Copyright (c) 2021 Clumio All Rights Reserved

// Package managementsubgroups contains methods related to ManagementSubgroups
package managementsubgroups

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// ManagementSubgroupsV1 represents a custom type struct
type ManagementSubgroupsV1 struct {
    config config.Config
}

//  ListManagementSubgroups Returns a list of subgroups.
func (m *ManagementSubgroupsV1) ListManagementSubgroups(
    groupId string, 
    limit *int64, 
    start *string)(
    *models.ListSubgroupsResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/management-groups/{group_id}/subgroups"
    //process optional template parameters
    pathParams := map[string]string{
        "group_id": groupId,
    }
    _queryBuilder := m.config.BaseUrl + _pathURL

    
    header := "application/management-subgroups=v1+json"
    var result *models.ListSubgroupsResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
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


//  ReadManagementSubgroup Subgroups are used to manage cloud connectors and SQL hosts residing in the same vCenter server.
//  
//  Returns a representation of the specified subgroups.
func (m *ManagementSubgroupsV1) ReadManagementSubgroup(
    subgroupId string, 
    groupId string)(
    *models.ReadSubgroupResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/management-groups/{group_id}/subgroups/{subgroup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "subgroup_id": subgroupId,
        "group_id": groupId,
    }
    _queryBuilder := m.config.BaseUrl + _pathURL

    
    header := "application/management-subgroups=v1+json"
    var result *models.ReadSubgroupResponse
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


//  UpdateManagementSubgroup Update the specified subgroup.
func (m *ManagementSubgroupsV1) UpdateManagementSubgroup(
    subgroupId string, 
    groupId string, 
    body *models.UpdateManagementSubgroupV1Request)(
    *models.UpdateSubgroupResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/management-groups/{group_id}/subgroups/{subgroup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "subgroup_id": subgroupId,
        "group_id": groupId,
    }
    _queryBuilder := m.config.BaseUrl + _pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/management-subgroups=v1+json"
    var result *models.UpdateSubgroupResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(m.config.Token).
        SetBody(payload).
        SetResult(&result).
        Put(_queryBuilder)

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
