// Copyright (c) 2021 Clumio All Rights Reserved

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
    embed *string)(
    *models.ListDynamoDBTableResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/datasources/aws/dynamodb-tables"

    
    header := "application/api.clumio.aws-dynamodb-tables=v1+json"
    result := &models.ListDynamoDBTableResponse{}
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
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
        "embed": *embed,
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
    defaultString := "" 
    
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
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
