// Copyright (c) 2023 Clumio All Rights Reserved

// Package awsneptune contains methods related to AwsNeptune
package awsneptune

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsNeptuneV1 represents a custom type struct
type AwsNeptuneV1 struct {
    config config.Config
}

// ListAwsNeptune Retrieve a list of Neptune clusters.
func (a *AwsNeptuneV1) ListAwsNeptune(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ListNeptuneResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/datasources/aws/neptune"

    
    header := "application/api.clumio.aws-neptune=v1+json"
    result := &models.ListNeptuneResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
    }
    if sort != nil {
        queryParams["sort"] = *sort
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


// ReadAwsNeptune Returns a representation of the specified Neptune cluster.
func (a *AwsNeptuneV1) ReadAwsNeptune(
    resourceId string, 
    lookbackDays *int64, 
    embed *string)(
    *models.ReadNeptuneResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/neptune/{resource_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "resource_id": resourceId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-neptune=v1+json"
    result := &models.ReadNeptuneResponse{}
    queryParams := make(map[string]string)
    if lookbackDays != nil {
        queryParams["lookback_days"] = fmt.Sprintf("%v", *lookbackDays)
    }
    if embed != nil {
        queryParams["embed"] = *embed
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
