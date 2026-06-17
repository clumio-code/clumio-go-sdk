// Copyright (c) 2023 Clumio All Rights Reserved

// Package awsicebergtables contains methods related to AwsIcebergTables
package awsicebergtables

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsIcebergTablesV1 represents a custom type struct
type AwsIcebergTablesV1 struct {
    config config.Config
}

// ListAwsIcebergTables Returns a list of Iceberg Tables.
func (a *AwsIcebergTablesV1) ListAwsIcebergTables(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ListIcebergTablesResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/datasources/aws/iceberg-tables"

    
    header := "application/api.clumio.aws-iceberg-tables=v1+json"
    result := &models.ListIcebergTablesResponse{}
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


// ReadAwsIcebergTable Returns a representation of the specified Iceberg Table.
func (a *AwsIcebergTablesV1) ReadAwsIcebergTable(
    tableId string, 
    lookbackDays *int64, 
    embed *string)(
    *models.ReadIcebergTableResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/iceberg-tables/{table_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "table_id": tableId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-iceberg-tables=v1+json"
    result := &models.ReadIcebergTableResponse{}
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
