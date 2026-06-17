// Copyright (c) 2023 Clumio All Rights Reserved

package awsicebergtables

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsIcebergTablesV1Client represents a custom type interface
type AwsIcebergTablesV1Client interface {
    // ListAwsIcebergTables Returns a list of Iceberg Tables.
    ListAwsIcebergTables(
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string, 
        lookbackDays *int64)(
        *models.ListIcebergTablesResponse,  *apiutils.APIError)
    
    // ReadAwsIcebergTable Returns a representation of the specified Iceberg Table.
    ReadAwsIcebergTable(
        tableId string, 
        lookbackDays *int64, 
        embed *string)(
        *models.ReadIcebergTableResponse,  *apiutils.APIError)
    
}

// NewAwsIcebergTablesV1 returns AwsIcebergTablesV1Client
func NewAwsIcebergTablesV1(config config.Config) AwsIcebergTablesV1Client {
    client := new(AwsIcebergTablesV1)
    client.config = config
    return client
}
