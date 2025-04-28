// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupawsdynamodbtables contains methods related to BackupAwsDynamodbTables
package backupawsdynamodbtables

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsDynamodbTablesV1 represents a custom type struct
type BackupAwsDynamodbTablesV1 struct {
    config config.Config
}

// ListBackupAwsDynamodbTables Retrieves a list of DynamoDB table backups.
func (b *BackupAwsDynamodbTablesV1) ListBackupAwsDynamodbTables(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string)(
    *models.ListDynamoDBTableBackupsResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/dynamodb-tables"

    
    header := "application/api.clumio.backup-aws-dynamodb-tables=v1+json"
    result := &models.ListDynamoDBTableBackupsResponse{}
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
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreateBackupAwsDynamodbTable Performs an on-demand backup for the specified DynamoDB table.
func (b *BackupAwsDynamodbTablesV1) CreateBackupAwsDynamodbTable(
    embed *string, 
    body models.CreateBackupAwsDynamodbTableV1Request)(
    *models.OnDemandDynamoDBBackupResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/dynamodb-tables"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.backup-aws-dynamodb-tables=v1+json"
    result := &models.OnDemandDynamoDBBackupResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadBackupAwsDynamodbTable Returns a representation of the specified DynamoDB table backup.
func (b *BackupAwsDynamodbTablesV1) ReadBackupAwsDynamodbTable(
    backupId string)(
    *models.ReadDynamoDBTableBackupResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/dynamodb-tables/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-aws-dynamodb-tables=v1+json"
    result := &models.ReadDynamoDBTableBackupResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
