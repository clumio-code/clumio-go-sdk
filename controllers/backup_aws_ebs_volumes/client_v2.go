// Copyright (c) 2023 Clumio All Rights Reserved

package backupawsebsvolumes

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsEbsVolumesV2Client represents a custom type interface
type BackupAwsEbsVolumesV2Client interface {
    // ListBackupAwsEbsVolumes Returns a list of EBS volumes that have been backed up by Clumio. EBS volume backups can be restored through the [POST /restores/aws/ebs-volumes](#operation/restore-aws-ebs-volume) endpoint.
    ListBackupAwsEbsVolumes(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string)(
        *models.ListEBSBackupsResponse,  *apiutils.APIError)
    
    // CreateBackupAwsEbsVolume Performs an on-demand backup for the specified EBS volume.
    CreateBackupAwsEbsVolume(
        embed *string, 
        body models.CreateBackupAwsEbsVolumeV2Request)(
        *models.OnDemandEBSBackupResponse,  *apiutils.APIError)
    
    // ReadBackupAwsEbsVolume Returns a representation of the specified EBS volume backup.
    ReadBackupAwsEbsVolume(
        backupId string)(
        *models.ReadEBSBackupResponse,  *apiutils.APIError)
    
}

// NewBackupAwsEbsVolumesV2 returns BackupAwsEbsVolumesV2Client
func NewBackupAwsEbsVolumesV2(config config.Config) BackupAwsEbsVolumesV2Client {
    client := new(BackupAwsEbsVolumesV2)
    client.config = config
    return client
}
