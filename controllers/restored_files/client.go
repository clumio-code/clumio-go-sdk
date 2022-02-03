// Copyright (c) 2021 Clumio All Rights Reserved

package restoredfiles

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredFilesV1Client represents a custom type interface
type RestoredFilesV1Client interface {
    // ListRestoredFiles Gets the list of active restored files for an asset.
    ListRestoredFiles(
        limit *int64, 
        start *string, 
        filter string)(
        *models.RestoredFilesResponse,  *apiutils.APIError)
    
    // RestoreFiles Restores one or more files from the specified backup.
    RestoreFiles(
        body models.RestoreFilesV1Request)(
        *models.RestoreFileResponse,  *apiutils.APIError)
    
}

// NewRestoredFilesV1 returns RestoredFilesV1Client
func NewRestoredFilesV1(config config.Config) RestoredFilesV1Client{
    client := new(RestoredFilesV1)
    client.config = config
    return client
}
