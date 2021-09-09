// Copyright (c) 2021 Clumio All Rights Reserved

// Package awsregions contains methods related to AwsRegions
package awsregions

import (
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// AwsRegionsV1 represents a custom type struct
type AwsRegionsV1 struct {
    config config.Config
}

//  ListConnectionAwsRegions Returns a list of valid regions for creating AWS connections
func (a *AwsRegionsV1) ListConnectionAwsRegions(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListAWSRegionsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := a.config.BaseUrl + "/connections/aws/regions"

    
    header := "application/aws-regions=v1+json"
    var result *models.ListAWSRegionsResponse
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
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
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
