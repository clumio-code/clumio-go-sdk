// Copyright (c) 2023 Clumio All Rights Reserved

// Package awsrdsresources contains methods related to AwsRdsResources
package awsrdsresources

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsRdsResourcesV1 represents a custom type struct
type AwsRdsResourcesV1 struct {
    config config.Config
}

// ListAwsRdsResources Retrieve a list of RDS resources.
func (a *AwsRdsResourcesV1) ListAwsRdsResources(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ListRdsResourcesResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/datasources/aws/rds-resources"

    
    header := "application/api.clumio.aws-rds-resources=v1+json"
    result := &models.ListRdsResourcesResponse{}
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
    if embed == nil {
        embed = &defaultString
    }
    if lookbackDays == nil {
        lookbackDays = &defaultInt64
    }
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
        "embed": *embed,
        "lookback_days": fmt.Sprintf("%v", *lookbackDays),
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: a.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadAwsRdsResource Returns a representation of the specified RDS resource.
func (a *AwsRdsResourcesV1) ReadAwsRdsResource(
    resourceId string, 
    lookbackDays *int64, 
    embed *string)(
    *models.ReadRdsResourceResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/rds-resources/{resource_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "resource_id": resourceId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-rds-resources=v1+json"
    result := &models.ReadRdsResourceResponse{}
    defaultInt64 := int64(0)
    defaultString := "" 
    
    if lookbackDays == nil {
        lookbackDays = &defaultInt64
    }
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
        "lookback_days": fmt.Sprintf("%v", *lookbackDays),
        "embed": *embed,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: a.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
