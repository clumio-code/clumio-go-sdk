// Copyright (c) 2023 Clumio All Rights Reserved

package backupawsrdsresourcedatabasetables

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsRdsResourceDatabaseTablesV1Client represents a custom type interface
type BackupAwsRdsResourceDatabaseTablesV1Client interface {
    // ListBackupAwsRdsResourceDatabaseTables Returns a list of RDS tables from the specified RDS backup.
    ListBackupAwsRdsResourceDatabaseTables(
        backupId string, 
        databaseName string, 
        currentCount *int64, 
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListRDSDatabaseTablesResponse,  *apiutils.APIError)
    
    // ReadBackupAwsRdsResourceDatabaseTable Returns a representation of the specified table from an RDS backup.
    ReadBackupAwsRdsResourceDatabaseTable(
        backupId string, 
        databaseName string, 
        tableId string, 
        embed *string)(
        *models.ReadRDSDatabaseTableResponse,  *apiutils.APIError)
    
    // ReadBackupAwsRdsResourceDatabaseTableColumns Returns a list of columns within the specified table.
    ReadBackupAwsRdsResourceDatabaseTableColumns(
        backupId string, 
        databaseName string, 
        tableId string)(
        *models.ReadRDSDatabaseTableColumnsResponse,  *apiutils.APIError)
    
}

// NewBackupAwsRdsResourceDatabaseTablesV1 returns BackupAwsRdsResourceDatabaseTablesV1Client
func NewBackupAwsRdsResourceDatabaseTablesV1(config config.Config) BackupAwsRdsResourceDatabaseTablesV1Client {
    client := new(BackupAwsRdsResourceDatabaseTablesV1)
    client.config = config
    return client
}
