// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupawsdocumentdb contains methods related to BackupAwsDocumentdb
package backupawsdocumentdb

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsDocumentdbV1 represents a custom type struct
type BackupAwsDocumentdbV1 struct {
    config config.Config
}

// ListBackupAwsDocumentdb Retrieves a list of DocumentDB backups.
func (b *BackupAwsDocumentdbV1) ListBackupAwsDocumentdb(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string)(
    *models.ListDocumentDBBackupsResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/documentdb"

    
    header := "application/api.clumio.backup-aws-documentdb=v1+json"
    result := &models.ListDocumentDBBackupsResponse{}
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


// CreateBackupAwsDocumentdb Performs an on-demand backup for the specified DocumentDB cluster.
func (b *BackupAwsDocumentdbV1) CreateBackupAwsDocumentdb(
    embed *string, 
    body models.CreateBackupAwsDocumentdbV1Request)(
    *models.OnDemandDocumentDBBackupResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/documentdb"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.backup-aws-documentdb=v1+json"
    result := &models.OnDemandDocumentDBBackupResponse{}
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


// ReadBackupAwsDocumentdb Returns a representation of the specified DocumentDB backup.
func (b *BackupAwsDocumentdbV1) ReadBackupAwsDocumentdb(
    backupId string)(
    *models.ReadDocumentDBBackupResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/documentdb/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-aws-documentdb=v1+json"
    result := &models.ReadDocumentDBBackupResponse{}

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
