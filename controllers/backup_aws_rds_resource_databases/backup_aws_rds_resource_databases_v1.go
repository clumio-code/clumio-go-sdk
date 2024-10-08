// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupawsrdsresourcedatabases contains methods related to BackupAwsRdsResourceDatabases
package backupawsrdsresourcedatabases

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsRdsResourceDatabasesV1 represents a custom type struct
type BackupAwsRdsResourceDatabasesV1 struct {
    config config.Config
}

// ListBackupAwsRdsResourceDatabases Retrieves a list of RDS databases from an RDS backup.
func (b *BackupAwsRdsResourceDatabasesV1) ListBackupAwsRdsResourceDatabases(
    backupId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListRDSBackupDatabasesResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/rds-resources/{backup_id}/databases"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-aws-rds-resource-databases=v1+json"
    result := &models.ListRDSBackupDatabasesResponse{}
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
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
