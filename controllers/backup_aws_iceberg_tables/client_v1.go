// Copyright (c) 2023 Clumio All Rights Reserved

package backupawsicebergtables

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsIcebergTablesV1Client represents a custom type interface
type BackupAwsIcebergTablesV1Client interface {
    // ListBackupAwsIcebergTables Retrieves a list of Iceberg table backups.
    ListBackupAwsIcebergTables(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string)(
        *models.ListIcebergTableBackupsResponse,  *apiutils.APIError)
    
    // CreateBackupAwsIcebergTable Performs an on-demand backup for the specified AWS Iceberg Table.
    CreateBackupAwsIcebergTable(
        embed *string, 
        body models.CreateBackupAwsIcebergTableV1Request)(
        *models.OnDemandAWSIcebergTableBackupResponse,  *apiutils.APIError)
    
    // ReadBackupAwsIcebergTable Returns a representation of the specified Iceberg table backup.
    ReadBackupAwsIcebergTable(
        backupId string)(
        *models.ReadIcebergTableBackupResponse,  *apiutils.APIError)
    
}

// NewBackupAwsIcebergTablesV1 returns BackupAwsIcebergTablesV1Client
func NewBackupAwsIcebergTablesV1(config config.Config) BackupAwsIcebergTablesV1Client {
    client := new(BackupAwsIcebergTablesV1)
    client.config = config
    return client
}
