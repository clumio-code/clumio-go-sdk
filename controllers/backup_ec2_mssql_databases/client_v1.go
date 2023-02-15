// Copyright (c) 2021 Clumio All Rights Reserved

package backupec2mssqldatabases

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupEc2MssqlDatabasesV1Client represents a custom type interface
type BackupEc2MssqlDatabasesV1Client interface {
    // ListBackupEc2MssqlDatabases Retrieve a list of EC2 MSSQL database backups.
    ListBackupEc2MssqlDatabases(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string, 
        embed *string)(
        *models.ListEC2MSSQLDatabaseBackupsResponse,  *apiutils.APIError)
    
    // CreateBackupEc2MssqlDatabase Performs an on-demand backup for the specified EC2 MSSQL asset.
    CreateBackupEc2MssqlDatabase(
        embed *string, 
        body models.CreateBackupEc2MssqlDatabaseV1Request)(
        *models.OnDemandEC2MSSQLDatabaseBackupResponse,  *apiutils.APIError)
    
    // ReadBackupEc2MssqlDatabase Returns a representation of the specified EC2 MSSQL database backup.
    ReadBackupEc2MssqlDatabase(
        backupId string)(
        *models.ReadEC2MSSQLDatabaseBackupResponse,  *apiutils.APIError)
    
}

// NewBackupEc2MssqlDatabasesV1 returns BackupEc2MssqlDatabasesV1Client
func NewBackupEc2MssqlDatabasesV1(config config.Config) BackupEc2MssqlDatabasesV1Client {
    client := new(BackupEc2MssqlDatabasesV1)
    client.config = config
    return client
}
