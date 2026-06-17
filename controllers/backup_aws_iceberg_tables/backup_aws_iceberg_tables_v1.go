// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupawsicebergtables contains methods related to BackupAwsIcebergTables
package backupawsicebergtables

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsIcebergTablesV1 represents a custom type struct
type BackupAwsIcebergTablesV1 struct {
    config config.Config
}

// ListBackupAwsIcebergTables Retrieves a list of Iceberg table backups.
func (b *BackupAwsIcebergTablesV1) ListBackupAwsIcebergTables(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string)(
    *models.ListIcebergTableBackupsResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/iceberg-tables"

    
    header := "application/api.clumio.backup-aws-iceberg-tables=v1+json"
    result := &models.ListIcebergTableBackupsResponse{}
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


// CreateBackupAwsIcebergTable Performs an on-demand backup for the specified AWS Iceberg Table.
func (b *BackupAwsIcebergTablesV1) CreateBackupAwsIcebergTable(
    embed *string, 
    body models.CreateBackupAwsIcebergTableV1Request)(
    *models.OnDemandAWSIcebergTableBackupResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/iceberg-tables"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.backup-aws-iceberg-tables=v1+json"
    result := &models.OnDemandAWSIcebergTableBackupResponse{}
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


// ReadBackupAwsIcebergTable Returns a representation of the specified Iceberg table backup.
func (b *BackupAwsIcebergTablesV1) ReadBackupAwsIcebergTable(
    backupId string)(
    *models.ReadIcebergTableBackupResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/iceberg-tables/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-aws-iceberg-tables=v1+json"
    result := &models.ReadIcebergTableBackupResponse{}

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
