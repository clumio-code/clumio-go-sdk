// Copyright (c) 2023 Clumio All Rights Reserved

// Package awsenvironments contains methods related to AwsEnvironments
package awsenvironments

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsEnvironmentsV1 represents a custom type struct
type AwsEnvironmentsV1 struct {
    config config.Config
}

// ListAwsEnvironments Returns a list of AWS environments.
func (a *AwsEnvironmentsV1) ListAwsEnvironments(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ListAWSEnvironmentsResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/datasources/aws/environments"

    
    header := "application/api.clumio.aws-environments=v1+json"
    result := &models.ListAWSEnvironmentsResponse{}
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
        Config: a.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadAwsEnvironment Returns a representation of the specified AWS environment.
func (a *AwsEnvironmentsV1) ReadAwsEnvironment(
    environmentId string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ReadAWSEnvironmentResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/environments/{environment_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "environment_id": environmentId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-environments=v1+json"
    result := &models.ReadAWSEnvironmentResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
    }
    if lookbackDays != nil {
        queryParams["lookback_days"] = fmt.Sprintf("%v", *lookbackDays)
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
