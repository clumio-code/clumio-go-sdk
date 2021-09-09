// Copyright (c) 2021 Clumio All Rights Reserved

// Package backupmssqldatabases contains methods related to BackupMssqlDatabases
package backupmssqldatabases

import (
    "encoding/json"
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// BackupMssqlDatabasesV1 represents a custom type struct
type BackupMssqlDatabasesV1 struct {
    config config.Config
}

//  ListBackupMssqlDatabases Retrieve a list of MSSQL database backups.
func (b *BackupMssqlDatabasesV1) ListBackupMssqlDatabases(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListMssqlDatabaseBackupsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := b.config.BaseUrl + "/backups/mssql/databases"

    
    header := "application/backup-mssql-databases=v1+json"
    var result *models.ListMssqlDatabaseBackupsResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    if filter == nil{
        filter = &defaultString
    }
    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetAuthToken(b.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}


//  CreateBackupMssqlDatabase Performs an on-demand backup for the specified MSSQL asset. The MSSQL asset must be protected with a policy that includes a service level agreement (SLA) configured for on-demand backups.
func (b *BackupMssqlDatabasesV1) CreateBackupMssqlDatabase(
    embed *string, 
    body models.CreateBackupMssqlDatabaseV1Request)(
    *models.OnDemandMssqlBackupResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := b.config.BaseUrl + "/backups/mssql/databases"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/backup-mssql-databases=v1+json"
    var result *models.OnDemandMssqlBackupResponse
    client := resty.New()
    defaultString := "" 
    

    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetAuthToken(b.config.Token).
        SetBody(payload).
        SetResult(&result).
        Post(_queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}


//  ReadBackupMssqlDatabase Returns a representation of the specified MSSQL database backup.
func (b *BackupMssqlDatabasesV1) ReadBackupMssqlDatabase(
    backupId string)(
    *models.ReadMssqlDatabaseBackupResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/backups/mssql/databases/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backupId": backupId,
    }
    _queryBuilder := b.config.BaseUrl + _pathURL

    
    header := "application/backup-mssql-databases=v1+json"
    var result *models.ReadMssqlDatabaseBackupResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(b.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}
