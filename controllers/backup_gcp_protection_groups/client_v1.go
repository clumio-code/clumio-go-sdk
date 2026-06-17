// Copyright (c) 2023 Clumio All Rights Reserved

package backupgcpprotectiongroups

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupGcpProtectionGroupsV1Client represents a custom type interface
type BackupGcpProtectionGroupsV1Client interface {
    // ListBackupGcpProtectionGroups Retrieves a list of protection group backups.
    ListBackupGcpProtectionGroups(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string)(
        *models.ListGCSProtectionGroupBackupsResponse,  *apiutils.APIError)
    
    // ReadBackupGcpProtectionGroup Returns a representation of the specified protection group backup.
    ReadBackupGcpProtectionGroup(
        backupId string)(
        *models.ReadGCSProtectionGroupBackupResponse,  *apiutils.APIError)
    
}

// NewBackupGcpProtectionGroupsV1 returns BackupGcpProtectionGroupsV1Client
func NewBackupGcpProtectionGroupsV1(config config.Config) BackupGcpProtectionGroupsV1Client {
    client := new(BackupGcpProtectionGroupsV1)
    client.config = config
    return client
}
