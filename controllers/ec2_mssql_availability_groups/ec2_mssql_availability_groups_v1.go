// Copyright (c) 2023 Clumio All Rights Reserved

// Package ec2mssqlavailabilitygroups contains methods related to Ec2MssqlAvailabilityGroups
package ec2mssqlavailabilitygroups

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlAvailabilityGroupsV1 represents a custom type struct
type Ec2MssqlAvailabilityGroupsV1 struct {
    config config.Config
}

// ListEc2MssqlAvailabilityGroups Returns a list of Availability Groups.
func (e *Ec2MssqlAvailabilityGroupsV1) ListEc2MssqlAvailabilityGroups(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ListEC2MssqlAGsResponse, *apiutils.APIError) {

    queryBuilder := e.config.BaseUrl + "/datasources/aws/ec2-mssql/availability-groups"

    
    header := "application/api.clumio.ec2-mssql-availability-groups=v1+json"
    result := &models.ListEC2MssqlAGsResponse{}
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
    if lookbackDays != nil {
        queryParams["lookback_days"] = fmt.Sprintf("%v", *lookbackDays)
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: e.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadEc2MssqlAvailabilityGroup Returns a representation of the specified availability group.
func (e *Ec2MssqlAvailabilityGroupsV1) ReadEc2MssqlAvailabilityGroup(
    availabilityGroupId string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ReadEC2MssqlAGResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/ec2-mssql/availability-groups/{availability_group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "availability_group_id": availabilityGroupId,
    }
    queryBuilder := e.config.BaseUrl + pathURL

    
    header := "application/api.clumio.ec2-mssql-availability-groups=v1+json"
    result := &models.ReadEC2MssqlAGResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
    }
    if lookbackDays != nil {
        queryParams["lookback_days"] = fmt.Sprintf("%v", *lookbackDays)
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: e.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
