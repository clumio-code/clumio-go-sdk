// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupfilesystemdirectories contains methods related to BackupFilesystemDirectories
package backupfilesystemdirectories

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupFilesystemDirectoriesV1 represents a custom type struct
type BackupFilesystemDirectoriesV1 struct {
    config config.Config
}

// ReadBackupFilesystemDirectory Browse files in the directory with the specified ID.
func (b *BackupFilesystemDirectoriesV1) ReadBackupFilesystemDirectory(
    backupId string, 
    filesystemId string, 
    directoryId string, 
    limit *int64, 
    start *string)(
    *models.ReadDirectoryResponse, *apiutils.APIError) {

    pathURL := "/backups/{backup_id}/filesystems/{filesystem_id}/directories/{directory_id}/browse"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
        "filesystem_id": filesystemId,
        "directory_id": directoryId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-filesystem-directories=v1+json"
    result := &models.ReadDirectoryResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
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
