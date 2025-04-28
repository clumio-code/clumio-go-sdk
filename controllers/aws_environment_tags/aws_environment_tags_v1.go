// Copyright (c) 2023 Clumio All Rights Reserved

// Package awsenvironmenttags contains methods related to AwsEnvironmentTags
package awsenvironmenttags

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsEnvironmentTagsV1 represents a custom type struct
type AwsEnvironmentTagsV1 struct {
    config config.Config
}

// ListAwsEnvironmentTags Returns a list of AWS tags in the specified environment.
func (a *AwsEnvironmentTagsV1) ListAwsEnvironmentTags(
    environmentId string, 
    currentCount *int64, 
    limit *int64, 
    totalCount *int64, 
    totalPagesCount *int64, 
    start *string, 
    filter *string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ListAwsTagsResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/environments/{environment_id}/tags"
    //process optional template parameters
    pathParams := map[string]string{
        "environment_id": environmentId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-environment-tags=v1+json"
    result := &models.ListAwsTagsResponse{}
    queryParams := make(map[string]string)
    if currentCount != nil {
        queryParams["current_count"] = fmt.Sprintf("%v", *currentCount)
    }
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if totalCount != nil {
        queryParams["total_count"] = fmt.Sprintf("%v", *totalCount)
    }
    if totalPagesCount != nil {
        queryParams["total_pages_count"] = fmt.Sprintf("%v", *totalPagesCount)
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
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadAwsEnvironmentTag Returns a representation of the specified AWS tag in the specified environment.
func (a *AwsEnvironmentTagsV1) ReadAwsEnvironmentTag(
    environmentId string, 
    tagId string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ReadAwsTagResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/environments/{environment_id}/tags/{tag_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "environment_id": environmentId,
        "tag_id": tagId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-environment-tags=v1+json"
    result := &models.ReadAwsTagResponse{}
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


// ReadAwsEnvironmentTagEbsVolumesProtectionStats Returns the specified AWS tag's EBS protection statistics.
func (a *AwsEnvironmentTagsV1) ReadAwsEnvironmentTagEbsVolumesProtectionStats(
    environmentId string, 
    tagId string)(
    *models.ReadEbsTagProtectionStatsResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/environments/{environment_id}/tags/{tag_id}/stats/protection/ebs-volumes"
    //process optional template parameters
    pathParams := map[string]string{
        "environment_id": environmentId,
        "tag_id": tagId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-environment-tags=v1+json"
    result := &models.ReadEbsTagProtectionStatsResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: a.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
