// Copyright (c) 2021 Clumio All Rights Reserved

package backupawsrdsresources

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsRdsResourcesV1Client represents a custom type interface
type BackupAwsRdsResourcesV1Client interface {
    // ListBackupAwsRdsResources Retrieves a list of RDS database backups.
    ListBackupAwsRdsResources(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string)(
        *models.ListRdsDatabaseBackupsResponse,  *apiutils.APIError)
    
    // ReadBackupAwsRdsResource Returns a representation of the specified RDS database backup.
    ReadBackupAwsRdsResource(
        backupId string)(
        *models.ReadRdsDatabaseBackupResponse,  *apiutils.APIError)
    
    // ListAwsRdsResourcesOptionGroups Retrieves a list of RDS option groups which are superset of persistent and permanent
    //  options present in the backup snapshot for a given environment.
    ListAwsRdsResourcesOptionGroups(
        backupId string, 
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListRdsOptionGroupsResponse,  *apiutils.APIError)
    
}

// NewBackupAwsRdsResourcesV1 returns BackupAwsRdsResourcesV1Client
func NewBackupAwsRdsResourcesV1(config config.Config) BackupAwsRdsResourcesV1Client {
    client := new(BackupAwsRdsResourcesV1)
    client.config = config
    return client
}
