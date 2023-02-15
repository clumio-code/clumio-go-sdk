// Copyright (c) 2021 Clumio All Rights Reserved

package backupawsrdsresourcedatabases

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsRdsResourceDatabasesV1Client represents a custom type interface
type BackupAwsRdsResourceDatabasesV1Client interface {
    // ListBackupAwsRdsResourceDatabases Retrieves a list of RDS databases from an RDS backup.
    ListBackupAwsRdsResourceDatabases(
        backupId string, 
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListRDSBackupDatabasesResponse,  *apiutils.APIError)
    
}

// NewBackupAwsRdsResourceDatabasesV1 returns BackupAwsRdsResourceDatabasesV1Client
func NewBackupAwsRdsResourceDatabasesV1(config config.Config) BackupAwsRdsResourceDatabasesV1Client {
    client := new(BackupAwsRdsResourceDatabasesV1)
    client.config = config
    return client
}
