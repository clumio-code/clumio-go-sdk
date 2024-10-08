// Copyright (c) 2023 Clumio All Rights Reserved

// Package mssqlavailabilitygroups contains methods related to MssqlAvailabilityGroups
package mssqlavailabilitygroups

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// MssqlAvailabilityGroupsV1 represents a custom type struct
type MssqlAvailabilityGroupsV1 struct {
    config config.Config
}

// ListMssqlAvailabilityGroups Returns a list of Availability Groups.
func (m *MssqlAvailabilityGroupsV1) ListMssqlAvailabilityGroups(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListMssqlAGsResponse, *apiutils.APIError) {

    queryBuilder := m.config.BaseUrl + "/datasources/mssql/availability-groups"

    
    header := "application/api.clumio.mssql-availability-groups=v1+json"
    result := &models.ListMssqlAGsResponse{}
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
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadMssqlAvailabilityGroup Returns a representation of the specified availability group.
func (m *MssqlAvailabilityGroupsV1) ReadMssqlAvailabilityGroup(
    availabilityGroupId string)(
    *models.ReadMssqlAGResponse, *apiutils.APIError) {

    pathURL := "/datasources/mssql/availability-groups/{availability_group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "availability_group_id": availabilityGroupId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    
    header := "application/api.clumio.mssql-availability-groups=v1+json"
    result := &models.ReadMssqlAGResponse{}

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
