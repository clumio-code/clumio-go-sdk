// Copyright (c) 2023 Clumio All Rights Reserved

package backupawsneptune

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsNeptuneV1Client represents a custom type interface
type BackupAwsNeptuneV1Client interface {
    // ListBackupAwsNeptune Retrieves a list of Neptune backups.
    ListBackupAwsNeptune(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string)(
        *models.ListNeptuneBackupsResponse,  *apiutils.APIError)
    
    // CreateBackupAwsNeptune Performs an on-demand backup for the specified Neptune cluster.
    CreateBackupAwsNeptune(
        embed *string, 
        body models.CreateBackupAwsNeptuneV1Request)(
        *models.OnDemandNeptuneBackupResponse,  *apiutils.APIError)
    
    // ReadBackupAwsNeptune Returns a representation of the specified Neptune backup.
    ReadBackupAwsNeptune(
        backupId string)(
        *models.ReadNeptuneBackupResponse,  *apiutils.APIError)
    
}

// NewBackupAwsNeptuneV1 returns BackupAwsNeptuneV1Client
func NewBackupAwsNeptuneV1(config config.Config) BackupAwsNeptuneV1Client {
    client := new(BackupAwsNeptuneV1)
    client.config = config
    return client
}
