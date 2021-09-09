// Copyright (c) 2021 Clumio All Rights Reserved

// Package backupfilesystemdirectories contains methods related to BackupFilesystemDirectories
package backupfilesystemdirectories

import (
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// BackupFilesystemDirectoriesV1 represents a custom type struct
type BackupFilesystemDirectoriesV1 struct {
    config config.Config
}

//  ReadBackupFilesystemDirectory Browse files in the directory with the specified ID.
func (b *BackupFilesystemDirectoriesV1) ReadBackupFilesystemDirectory(
    backupId string, 
    filesystemId string, 
    directoryId string, 
    limit *int64, 
    start *string)(
    *models.ReadDirectoryResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/backups/{backup_id}/filesystems/{filesystem_id}/directories/{directory_id}/browse"
    //process optional template parameters
    pathParams := map[string]string{
        "backupId": backupId,
        "filesystemId": filesystemId,
        "directoryId": directoryId,
    }
    _queryBuilder := b.config.BaseUrl + _pathURL

    
    header := "application/backup-filesystem-directories=v1+json"
    var result *models.ReadDirectoryResponse
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
