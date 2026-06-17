// Copyright (c) 2023 Clumio All Rights Reserved

package backupgcpprotectiongroupsgcsassets

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupGcpProtectionGroupsGcsAssetsV1Client represents a custom type interface
type BackupGcpProtectionGroupsGcsAssetsV1Client interface {
    // ListBackupGcpProtectionGroupGcsAssets Retrieves a list of asset backups.
    ListBackupGcpProtectionGroupGcsAssets(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string)(
        *models.ListGCSProtectionGroupAssetBackupsResponse,  *apiutils.APIError)
    
    // ReadBackupGcpProtectionGroupGcsAsset Returns a representation of the specified asset backup.
    ReadBackupGcpProtectionGroupGcsAsset(
        backupId string)(
        *models.ReadGCSProtectionGroupAssetBackupResponse,  *apiutils.APIError)
    
}

// NewBackupGcpProtectionGroupsGcsAssetsV1 returns BackupGcpProtectionGroupsGcsAssetsV1Client
func NewBackupGcpProtectionGroupsGcsAssetsV1(config config.Config) BackupGcpProtectionGroupsGcsAssetsV1Client {
    client := new(BackupGcpProtectionGroupsGcsAssetsV1)
    client.config = config
    return client
}
