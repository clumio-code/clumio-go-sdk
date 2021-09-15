// Copyright (c) 2021 Clumio All Rights Reserved

package backupmssqldatabases

import (
     "github.com/clumio-code/clumiogosdk/api_utils"
     "github.com/clumio-code/clumiogosdk/config"
     "github.com/clumio-code/clumiogosdk/models"
)

// BackupMssqlDatabasesV1Client represents a custom type interface
type BackupMssqlDatabasesV1Client interface {
    //  Retrieve a list of MSSQL database backups.
    ListBackupMssqlDatabases(
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string)(
        *models.ListMssqlDatabaseBackupsResponse,  *apiutils.APIError)
    
    //  Performs an on-demand backup for the specified MSSQL asset. The MSSQL asset must be protected with a policy that includes a service level agreement (SLA) configured for on-demand backups.
    CreateBackupMssqlDatabase(
        embed *string, 
        body models.CreateBackupMssqlDatabaseV1Request)(
        *models.OnDemandMssqlBackupResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified MSSQL database backup.
    ReadBackupMssqlDatabase(
        backupId string)(
        *models.ReadMssqlDatabaseBackupResponse,  *apiutils.APIError)
    
}

// NewBackupMssqlDatabasesV1 returns BackupMssqlDatabasesV1Client
func NewBackupMssqlDatabasesV1(config config.Config) BackupMssqlDatabasesV1Client{
    client := new(BackupMssqlDatabasesV1)
    client.config = config
    return client
}
