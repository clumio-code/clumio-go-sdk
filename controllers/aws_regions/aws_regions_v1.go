// Copyright (c) 2021 Clumio All Rights Reserved

// Package awsregions contains methods related to AwsRegions
package awsregions

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsRegionsV1 represents a custom type struct
type AwsRegionsV1 struct {
    config config.Config
}

// ListConnectionAwsRegions Returns a list of valid regions for creating AWS connections
func (a *AwsRegionsV1) ListConnectionAwsRegions(
    limit *int64, 
    start *string)(
    *models.ListAWSRegionsResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/connections/aws/regions"

    
    header := "application/api.clumio.aws-regions=v1+json"
    result := &models.ListAWSRegionsResponse{}
    defaultInt64 := int64(0)
    defaultString := "" 
    
    if limit == nil {
        limit = &defaultInt64
    }
    if start == nil {
        start = &defaultString
    }
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
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
