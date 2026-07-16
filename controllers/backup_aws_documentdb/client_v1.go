// Copyright (c) 2023 Clumio All Rights Reserved

package backupawsdocumentdb

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsDocumentdbV1Client represents a custom type interface
type BackupAwsDocumentdbV1Client interface {
    // ListBackupAwsDocumentdb Retrieves a list of DocumentDB backups.
    ListBackupAwsDocumentdb(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string)(
        *models.ListDocumentDBBackupsResponse,  *apiutils.APIError)
    
    // CreateBackupAwsDocumentdb Performs an on-demand backup for the specified DocumentDB cluster.
    CreateBackupAwsDocumentdb(
        embed *string, 
        body models.CreateBackupAwsDocumentdbV1Request)(
        *models.OnDemandDocumentDBBackupResponse,  *apiutils.APIError)
    
    // ReadBackupAwsDocumentdb Returns a representation of the specified DocumentDB backup.
    ReadBackupAwsDocumentdb(
        backupId string)(
        *models.ReadDocumentDBBackupResponse,  *apiutils.APIError)
    
}

// NewBackupAwsDocumentdbV1 returns BackupAwsDocumentdbV1Client
func NewBackupAwsDocumentdbV1(config config.Config) BackupAwsDocumentdbV1Client {
    client := new(BackupAwsDocumentdbV1)
    client.config = config
    return client
}
