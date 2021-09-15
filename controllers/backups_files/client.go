// Copyright (c) 2021 Clumio All Rights Reserved

package backupsfiles

import (
     "github.com/clumio-code/clumiogosdk/api_utils"
     "github.com/clumio-code/clumiogosdk/config"
     "github.com/clumio-code/clumiogosdk/models"
)

// BackupsFilesV1Client represents a custom type interface
type BackupsFilesV1Client interface {
    //  Retrieve the list of files whose name matches a given regex pattern.
    ListFiles(
        limit *int64, 
        start *string, 
        filter string)(
        *models.FileSearchResponse,  *apiutils.APIError)
    
    //  Retrieve the list of versions of the file.
    ListFileVersions(
        searchResultId string, 
        limit *int64, 
        start *string)(
        *models.FileListResponse,  *apiutils.APIError)
    
}

// NewBackupsFilesV1 returns BackupsFilesV1Client
func NewBackupsFilesV1(config config.Config) BackupsFilesV1Client{
    client := new(BackupsFilesV1)
    client.config = config
    return client
}
