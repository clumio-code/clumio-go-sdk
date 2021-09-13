// Copyright (c) 2021 Clumio All Rights Reserved

// Package awsenvironments contains methods related to AwsEnvironments
package awsenvironments

import (
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// AwsEnvironmentsV1 represents a custom type struct
type AwsEnvironmentsV1 struct {
    config config.Config
}

//  ListAwsEnvironments Returns a list of AWS environments.
func (a *AwsEnvironmentsV1) ListAwsEnvironments(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListAWSEnvironmentsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := a.config.BaseUrl + "/datasources/aws/environments"

    
    header := "application/aws-environments=v1+json"
    var result *models.ListAWSEnvironmentsResponse
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
    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetAuthToken(a.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}


//  ReadAwsEnvironment Returns a representation of the specified AWS environment.
func (a *AwsEnvironmentsV1) ReadAwsEnvironment(
    environmentId string, 
    embed *string)(
    *models.ReadAWSEnvironmentResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/aws/environments/{environment_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "environmentId": environmentId,
    }
    _queryBuilder := a.config.BaseUrl + _pathURL

    
    header := "application/aws-environments=v1+json"
    var result *models.ReadAWSEnvironmentResponse
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
        SetHeader("Accept", header).
        SetAuthToken(a.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}
