// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupawsrdsresourcedatabasetables contains methods related to BackupAwsRdsResourceDatabaseTables
package backupawsrdsresourcedatabasetables

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsRdsResourceDatabaseTablesV1 represents a custom type struct
type BackupAwsRdsResourceDatabaseTablesV1 struct {
    config config.Config
}

// ListBackupAwsRdsResourceDatabaseTables Returns a list of RDS tables from the specified RDS backup.
func (b *BackupAwsRdsResourceDatabaseTablesV1) ListBackupAwsRdsResourceDatabaseTables(
    backupId string, 
    databaseName string, 
    currentCount *int64, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListRDSDatabaseTablesResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/rds-resources/{backup_id}/databases/{database_name}/tables"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
        "database_name": databaseName,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-aws-rds-resource-database-tables=v1+json"
    result := &models.ListRDSDatabaseTablesResponse{}
    queryParams := make(map[string]string)
    if currentCount != nil {
        queryParams["current_count"] = fmt.Sprintf("%v", *currentCount)
    }
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


// ReadBackupAwsRdsResourceDatabaseTable Returns a representation of the specified table from an RDS backup.
func (b *BackupAwsRdsResourceDatabaseTablesV1) ReadBackupAwsRdsResourceDatabaseTable(
    backupId string, 
    databaseName string, 
    tableId string, 
    embed *string)(
    *models.ReadRDSDatabaseTableResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/rds-resources/{backup_id}/databases/{database_name}/tables/{table_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
        "database_name": databaseName,
        "table_id": tableId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-aws-rds-resource-database-tables=v1+json"
    result := &models.ReadRDSDatabaseTableResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
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


// ReadBackupAwsRdsResourceDatabaseTableColumns Returns a list of columns within the specified table.
func (b *BackupAwsRdsResourceDatabaseTablesV1) ReadBackupAwsRdsResourceDatabaseTableColumns(
    backupId string, 
    databaseName string, 
    tableId string)(
    *models.ReadRDSDatabaseTableColumnsResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/rds-resources/{backup_id}/databases/{database_name}/tables/{table_id}/columns"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
        "database_name": databaseName,
        "table_id": tableId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-aws-rds-resource-database-tables=v1+json"
    result := &models.ReadRDSDatabaseTableColumnsResponse{}

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
