// Copyright (c) 2021 Clumio All Rights Reserved

// Package backupfilesystems contains methods related to BackupFilesystems
package backupfilesystems

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// BackupFilesystemsV1 represents a custom type struct
type BackupFilesystemsV1 struct {
    config config.Config
}

//  ListBackupFilesystems Returns a list of filesystems.
func (b *BackupFilesystemsV1) ListBackupFilesystems(
    backupId string, 
    limit *int64, 
    start *string)(
    *models.ListFileSystemsResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/backups/{backup_id}/filesystems"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/backup-filesystems=v1+json"
    var result *models.ListFileSystemsResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetHeader("x-clumio-organizationalunit-context", b.config.OrganizationalUnitContext).
        SetAuthToken(b.config.Token).
        SetResult(&result).
        Get(queryBuilder)

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


//  ReadFilesystem Returns a representation of the specified filesystem.
func (b *BackupFilesystemsV1) ReadFilesystem(
    filesystemId string, 
    backupId string)(
    *models.ReadFileSystemResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/backups/{backup_id}/filesystems/{filesystem_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "filesystem_id": filesystemId,
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/backup-filesystems=v1+json"
    var result *models.ReadFileSystemResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetHeader("x-clumio-organizationalunit-context", b.config.OrganizationalUnitContext).
        SetAuthToken(b.config.Token).
        SetResult(&result).
        Get(queryBuilder)

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
