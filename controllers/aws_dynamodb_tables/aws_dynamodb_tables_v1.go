// Copyright (c) 2023 Clumio All Rights Reserved

// Package awsdynamodbtables contains methods related to AwsDynamodbTables
package awsdynamodbtables

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsDynamodbTablesV1 represents a custom type struct
type AwsDynamodbTablesV1 struct {
    config config.Config
}

// ListAwsDynamodbTables Retrieve a list of DynamoDB tables.
func (a *AwsDynamodbTablesV1) ListAwsDynamodbTables(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ListDynamoDBTableResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/datasources/aws/dynamodb-tables"

    
    header := "application/api.clumio.aws-dynamodb-tables=v1+json"
    result := &models.ListDynamoDBTableResponse{}
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


// ReadAwsDynamodbTable Returns a representation of specified DynamoDB table.
func (a *AwsDynamodbTablesV1) ReadAwsDynamodbTable(
    tableId string, 
    lookbackDays *int64, 
    embed *string)(
    *models.ReadDynamoDBTableResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/dynamodb-tables/{table_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "table_id": tableId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-dynamodb-tables=v1+json"
    result := &models.ReadDynamoDBTableResponse{}
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
