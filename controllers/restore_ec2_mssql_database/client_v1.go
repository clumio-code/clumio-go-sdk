// Copyright (c) 2021 Clumio All Rights Reserved

package restoreec2mssqldatabase

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoreEc2MssqlDatabaseV1Client represents a custom type interface
type RestoreEc2MssqlDatabaseV1Client interface {
    // RestoreEc2MssqlDatabase Restores an EC2 MSSQL database from a given backup or to a specified point in time.
    RestoreEc2MssqlDatabase(
        embed *string, 
        body models.RestoreEc2MssqlDatabaseV1Request)(
        *models.CreateEC2MSSQLDatabaseRestoreResponse,  *apiutils.APIError)
    
}

// NewRestoreEc2MssqlDatabaseV1 returns RestoreEc2MssqlDatabaseV1Client
func NewRestoreEc2MssqlDatabaseV1(config config.Config) RestoreEc2MssqlDatabaseV1Client {
    client := new(RestoreEc2MssqlDatabaseV1)
    client.config = config
    return client
}
