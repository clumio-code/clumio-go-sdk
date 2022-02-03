// Copyright (c) 2021 Clumio All Rights Reserved

// Package mssqlavailabilitygroups contains methods related to MssqlAvailabilityGroups
package mssqlavailabilitygroups

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// MssqlAvailabilityGroupsV1 represents a custom type struct
type MssqlAvailabilityGroupsV1 struct {
    config config.Config
}

//  ListMssqlAvailabilityGroups Returns a list of Availability Groups.
func (m *MssqlAvailabilityGroupsV1) ListMssqlAvailabilityGroups(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListMssqlAGsResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := m.config.BaseUrl + "/datasources/mssql/availability-groups"

    
    header := "application/mssql-availability-groups=v1+json"
    var result *models.ListMssqlAGsResponse
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
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, m.config.OrganizationalUnitContext).
        SetAuthToken(m.config.Token).
        SetResult(&result).
        Get(queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       common.InternalServerError,
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       common.NonSuccessStatusCodeError,
            Response:     res.Body(),
        }
    }
    return result, nil
}


//  ReadMssqlAvailabilityGroup Returns a representation of the specified availability group.
func (m *MssqlAvailabilityGroupsV1) ReadMssqlAvailabilityGroup(
    availabilityGroupId string)(
    *models.ReadMssqlAGResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/mssql/availability-groups/{availability_group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "availability_group_id": availabilityGroupId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    
    header := "application/mssql-availability-groups=v1+json"
    var result *models.ReadMssqlAGResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, m.config.OrganizationalUnitContext).
        SetAuthToken(m.config.Token).
        SetResult(&result).
        Get(queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       common.InternalServerError,
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       common.NonSuccessStatusCodeError,
            Response:     res.Body(),
        }
    }
    return result, nil
}
