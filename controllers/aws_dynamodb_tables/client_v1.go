// Copyright (c) 2023 Clumio All Rights Reserved

package awsdynamodbtables

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsDynamodbTablesV1Client represents a custom type interface
type AwsDynamodbTablesV1Client interface {
    // ListAwsDynamodbTables Retrieve a list of DynamoDB tables.
    ListAwsDynamodbTables(
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string, 
        lookbackDays *int64)(
        *models.ListDynamoDBTableResponse,  *apiutils.APIError)
    
    // ReadAwsDynamodbTable Returns a representation of specified DynamoDB table.
    ReadAwsDynamodbTable(
        tableId string, 
        lookbackDays *int64, 
        embed *string)(
        *models.ReadDynamoDBTableResponse,  *apiutils.APIError)
    
}

// NewAwsDynamodbTablesV1 returns AwsDynamodbTablesV1Client
func NewAwsDynamodbTablesV1(config config.Config) AwsDynamodbTablesV1Client {
    client := new(AwsDynamodbTablesV1)
    client.config = config
    return client
}
