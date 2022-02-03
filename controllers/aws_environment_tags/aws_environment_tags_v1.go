// Copyright (c) 2021 Clumio All Rights Reserved

// Package awsenvironmenttags contains methods related to AwsEnvironmentTags
package awsenvironmenttags

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// AwsEnvironmentTagsV1 represents a custom type struct
type AwsEnvironmentTagsV1 struct {
    config config.Config
}

//  ListAwsEnvironmentTags Returns a list of AWS tags in the specified environment.
func (a *AwsEnvironmentTagsV1) ListAwsEnvironmentTags(
    environmentId string, 
    currentCount *int64, 
    limit *int64, 
    totalCount *int64, 
    totalPagesCount *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListAwsTagsResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/aws/environments/{environment_id}/tags"
    //process optional template parameters
    pathParams := map[string]string{
        "environment_id": environmentId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/aws-environment-tags=v1+json"
    var result *models.ListAwsTagsResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if currentCount == nil{
        currentCount = &defaultInt64
    }
    if limit == nil{
        limit = &defaultInt64
    }
    if totalCount == nil{
        totalCount = &defaultInt64
    }
    if totalPagesCount == nil{
        totalPagesCount = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    if filter == nil{
        filter = &defaultString
    }
    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "currentCount": fmt.Sprintf("%v", *currentCount),
        "limit": fmt.Sprintf("%v", *limit),
        "totalCount": fmt.Sprintf("%v", *totalCount),
        "totalPagesCount": fmt.Sprintf("%v", *totalPagesCount),
        "start": *start,
        "filter": *filter,
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetPathParams(pathParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, a.config.OrganizationalUnitContext).
        SetAuthToken(a.config.Token).
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


//  ReadAwsEnvironmentTag Returns a representation of the specified AWS tag in the specified environment.
func (a *AwsEnvironmentTagsV1) ReadAwsEnvironmentTag(
    environmentId string, 
    tagId string, 
    embed *string)(
    *models.ReadAwsTagResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/aws/environments/{environment_id}/tags/{tag_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "environment_id": environmentId,
        "tag_id": tagId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/aws-environment-tags=v1+json"
    var result *models.ReadAwsTagResponse
    client := resty.New()
    defaultString := "" 
    

    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetPathParams(pathParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, a.config.OrganizationalUnitContext).
        SetAuthToken(a.config.Token).
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


//  ReadAwsEnvironmentTagEbsVolumesComplianceStats Returns the specified AWS tag's EBS compliance statistics.
func (a *AwsEnvironmentTagsV1) ReadAwsEnvironmentTagEbsVolumesComplianceStats(
    environmentId string, 
    tagId string)(
    *models.ReadEbsTagComplianceStatsResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/aws/environments/{environment_id}/tags/{tag_id}/stats/compliance/ebs-volumes"
    //process optional template parameters
    pathParams := map[string]string{
        "environment_id": environmentId,
        "tag_id": tagId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/aws-environment-tags=v1+json"
    var result *models.ReadEbsTagComplianceStatsResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, a.config.OrganizationalUnitContext).
        SetAuthToken(a.config.Token).
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
