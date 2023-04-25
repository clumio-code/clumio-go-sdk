// Copyright (c) 2021 Clumio All Rights Reserved

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
    filter *string)(
    *models.ListEC2MssqlAGsResponse, *apiutils.APIError) {

    queryBuilder := e.config.BaseUrl + "/datasources/aws/ec2-mssql/availability-groups"

    
    header := "application/api.clumio.ec2-mssql-availability-groups=v1+json"
    result := &models.ListEC2MssqlAGsResponse{}
    defaultInt64 := int64(0)
    defaultString := "" 
    
    if limit == nil {
        limit = &defaultInt64
    }
    if start == nil {
        start = &defaultString
    }
    if filter == nil {
        filter = &defaultString
    }
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
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
    availabilityGroupId string)(
    *models.ReadEC2MssqlAGResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/ec2-mssql/availability-groups/{availability_group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "availability_group_id": availabilityGroupId,
    }
    queryBuilder := e.config.BaseUrl + pathURL

    
    header := "application/api.clumio.ec2-mssql-availability-groups=v1+json"
    result := &models.ReadEC2MssqlAGResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: e.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
