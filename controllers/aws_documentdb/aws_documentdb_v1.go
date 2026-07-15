// Copyright (c) 2023 Clumio All Rights Reserved

// Package awsdocumentdb contains methods related to AwsDocumentdb
package awsdocumentdb

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsDocumentdbV1 represents a custom type struct
type AwsDocumentdbV1 struct {
    config config.Config
}

// ListAwsDocumentdb Retrieve a list of DocumentDB clusters.
func (a *AwsDocumentdbV1) ListAwsDocumentdb(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ListDocumentDBResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/datasources/aws/documentdb"

    
    header := "application/api.clumio.aws-documentdb=v1+json"
    result := &models.ListDocumentDBResponse{}
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


// ReadAwsDocumentdb Returns a representation of the specified DocumentDB cluster.
func (a *AwsDocumentdbV1) ReadAwsDocumentdb(
    resourceId string, 
    lookbackDays *int64, 
    embed *string)(
    *models.ReadDocumentDBResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/documentdb/{resource_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "resource_id": resourceId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-documentdb=v1+json"
    result := &models.ReadDocumentDBResponse{}
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
