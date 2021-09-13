// Copyright (c) 2021 Clumio All Rights Reserved

package backupfilesystems

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// BackupFilesystemsV1Client represents a custom type interface
type BackupFilesystemsV1Client interface {
    //  Returns a list of filesystems.
    ListBackupFilesystems(
        backupId string, 
        limit *int64, 
        start *string)(
        *models.ListFileSystemsResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified filesystem.
    ReadFilesystem(
        filesystemId string, 
        backupId string)(
        *models.ReadFileSystemResponse,  *apiutils.APIError)
    
}

// NewBackupFilesystemsV1 returns BackupFilesystemsV1Client
func NewBackupFilesystemsV1(config config.Config) BackupFilesystemsV1Client{
    client := new(BackupFilesystemsV1)
    client.config = config
    return client
}
