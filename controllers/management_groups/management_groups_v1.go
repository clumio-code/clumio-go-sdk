// Copyright (c) 2021 Clumio All Rights Reserved

// Package managementgroups contains methods related to ManagementGroups
package managementgroups

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// ManagementGroupsV1 represents a custom type struct
type ManagementGroupsV1 struct {
    config config.Config
}

//  ListManagementGroups Returns a list of management groups.
func (m *ManagementGroupsV1) ListManagementGroups(
    limit *int64, 
    start *string)(
    *models.ListManagementGroupsResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := m.config.BaseUrl + "/management-groups"

    
    header := "application/management-groups=v1+json"
    var result *models.ListManagementGroupsResponse
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
        SetHeader("Accept", header).
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


//  ReadManagementGroup Returns a representation of the specified management group. Management groups are used to
//  manage the SQL hosts and cloud connectors deployed in vCenter servers.
//  
//  Returns a representation of the specified management-groups.
func (m *ManagementGroupsV1) ReadManagementGroup(
    groupId string)(
    *models.ReadManagementGroupResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/management-groups/{group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "group_id": groupId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    
    header := "application/management-groups=v1+json"
    var result *models.ReadManagementGroupResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
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


//  UpdateManagementGroup Update the specified management group.
func (m *ManagementGroupsV1) UpdateManagementGroup(
    groupId string, 
    body *models.UpdateManagementGroupV1Request)(
    *models.UpdateManagementGroupResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/management-groups/{group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "group_id": groupId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/management-groups=v1+json"
    var result *models.UpdateManagementGroupResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(m.config.Token).
        SetBody(payload).
        SetResult(&result).
        Put(queryBuilder)

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
