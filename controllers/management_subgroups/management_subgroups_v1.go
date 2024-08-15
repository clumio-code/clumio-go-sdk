// Copyright (c) 2023 Clumio All Rights Reserved

// Package managementsubgroups contains methods related to ManagementSubgroups
package managementsubgroups

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// ManagementSubgroupsV1 represents a custom type struct
type ManagementSubgroupsV1 struct {
    config config.Config
}

// ListManagementSubgroups Returns a list of subgroups.
func (m *ManagementSubgroupsV1) ListManagementSubgroups(
    groupId string, 
    limit *int64, 
    start *string)(
    *models.ListSubgroupsResponse, *apiutils.APIError) {

    pathURL := "/management-groups/{group_id}/subgroups"
    //process optional template parameters
    pathParams := map[string]string{
        "group_id": groupId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    
    header := "application/api.clumio.management-subgroups=v1+json"
    result := &models.ListSubgroupsResponse{}
    queryParams := make(map[string]string)
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
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadManagementSubgroup Subgroups are used to manage cloud connectors and SQL hosts residing in the same vCenter server.
//  
//  Returns a representation of the specified subgroups.
func (m *ManagementSubgroupsV1) ReadManagementSubgroup(
    subgroupId string, 
    groupId string)(
    *models.ReadSubgroupResponse, *apiutils.APIError) {

    pathURL := "/management-groups/{group_id}/subgroups/{subgroup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "subgroup_id": subgroupId,
        "group_id": groupId,
    }
    queryBuilder := m.config.BaseUrl + pathURL

    
    header := "application/api.clumio.management-subgroups=v1+json"
    result := &models.ReadSubgroupResponse{}

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
