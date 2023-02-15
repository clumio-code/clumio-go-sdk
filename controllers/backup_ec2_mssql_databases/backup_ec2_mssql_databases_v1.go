// Copyright (c) 2021 Clumio All Rights Reserved

// Package backupec2mssqldatabases contains methods related to BackupEc2MssqlDatabases
package backupec2mssqldatabases

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupEc2MssqlDatabasesV1 represents a custom type struct
type BackupEc2MssqlDatabasesV1 struct {
    config config.Config
}

// ListBackupEc2MssqlDatabases Retrieve a list of EC2 MSSQL database backups.
func (b *BackupEc2MssqlDatabasesV1) ListBackupEc2MssqlDatabases(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string, 
    embed *string)(
    *models.ListEC2MSSQLDatabaseBackupsResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/ec2-mssql/databases"

    
    header := "application/api.clumio.backup-ec2-mssql-databases=v1+json"
    result := &models.ListEC2MSSQLDatabaseBackupsResponse{}
    defaultInt64 := int64(0)
    defaultString := "" 
    
    if limit == nil {
        limit = &defaultInt64
    }
    if start == nil {
        start = &defaultString
    }
    if sort == nil {
        sort = &defaultString
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
        "sort": *sort,
        "filter": *filter,
        "embed": *embed,
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


// CreateBackupEc2MssqlDatabase Performs an on-demand backup for the specified EC2 MSSQL asset.
func (b *BackupEc2MssqlDatabasesV1) CreateBackupEc2MssqlDatabase(
    embed *string, 
    body models.CreateBackupEc2MssqlDatabaseV1Request)(
    *models.OnDemandEC2MSSQLDatabaseBackupResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/ec2-mssql/databases"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.backup-ec2-mssql-databases=v1+json"
    result := &models.OnDemandEC2MSSQLDatabaseBackupResponse{}
    defaultString := "" 
    
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
        "embed": *embed,
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


// ReadBackupEc2MssqlDatabase Returns a representation of the specified EC2 MSSQL database backup.
func (b *BackupEc2MssqlDatabasesV1) ReadBackupEc2MssqlDatabase(
    backupId string)(
    *models.ReadEC2MSSQLDatabaseBackupResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/ec2-mssql/databases/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-ec2-mssql-databases=v1+json"
    result := &models.ReadEC2MSSQLDatabaseBackupResponse{}

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
