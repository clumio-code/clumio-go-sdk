// Copyright (c) 2021 Clumio All Rights Reserved

package restoredmssqldatabases

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// RestoredMssqlDatabasesV1Client represents a custom type interface
type RestoredMssqlDatabasesV1Client interface {
    //  Creates a restored MSSQL database from a given backup or to a specified point in time.
    RestoreMssqlDatabase(
        embed *string, 
        body models.RestoreMssqlDatabaseV1Request)(
        *models.CreateMssqlDatabaseRestoreResponse,  *apiutils.APIError)
    
}

// NewRestoredMssqlDatabasesV1 returns RestoredMssqlDatabasesV1Client
func NewRestoredMssqlDatabasesV1(config config.Config) RestoredMssqlDatabasesV1Client{
    client := new(RestoredMssqlDatabasesV1)
    client.config = config
    return client
}
