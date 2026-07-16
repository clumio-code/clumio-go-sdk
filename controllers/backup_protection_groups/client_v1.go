// Copyright (c) 2023 Clumio All Rights Reserved

package backupprotectiongroups

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupProtectionGroupsV1Client represents a custom type interface
type BackupProtectionGroupsV1Client interface {
    // ListBackupProtectionGroups Retrieves a list of protection group backups.
    ListBackupProtectionGroups(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string)(
        *models.ListProtectionGroupBackupsResponse,  *apiutils.APIError)
    
    // ListBackupProtectionGroupS3Assets Retrieves a list of protection group S3 asset backups.
    ListBackupProtectionGroupS3Assets(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string)(
        *models.ListProtectionGroupS3AssetBackupsResponse,  *apiutils.APIError)
    
    // ExportProtectionGroupS3AssetMalwareReport Exports the specified malware report for a protection group S3 asset. This endpoint is deprecated; use [POST /backups/protection-groups/s3-assets/threat-report](#operation/export-protection-group-s3-asset-threat-report) instead.
    ExportProtectionGroupS3AssetMalwareReport(
        embed *string, 
        body models.ExportProtectionGroupS3AssetMalwareReportV1Request)(
        *models.ExportMalwareReportResponse,  *apiutils.APIError)
    
    // ExportProtectionGroupS3AssetThreatReport Exports the specified threat report for a protection group S3 asset.
    ExportProtectionGroupS3AssetThreatReport(
        embed *string, 
        body models.ExportProtectionGroupS3AssetThreatReportV1Request)(
        *models.ExportThreatReportResponse,  *apiutils.APIError)
    
    // ReadBackupProtectionGroupS3Asset Returns a representation of the specified protection group S3 asset backup.
    ReadBackupProtectionGroupS3Asset(
        backupId string)(
        *models.ReadProtectionGroupS3AssetBackupResponse,  *apiutils.APIError)
    
    // ReadBackupProtectionGroup Returns a representation of the specified protection group backup.
    ReadBackupProtectionGroup(
        backupId string)(
        *models.ReadProtectionGroupBackupResponse,  *apiutils.APIError)
    
}

// NewBackupProtectionGroupsV1 returns BackupProtectionGroupsV1Client
func NewBackupProtectionGroupsV1(config config.Config) BackupProtectionGroupsV1Client {
    client := new(BackupProtectionGroupsV1)
    client.config = config
    return client
}
