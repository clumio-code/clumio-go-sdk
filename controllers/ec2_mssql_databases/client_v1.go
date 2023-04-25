// Copyright (c) 2021 Clumio All Rights Reserved

package ec2mssqldatabases

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlDatabasesV1Client represents a custom type interface
type Ec2MssqlDatabasesV1Client interface {
    // ListEc2MssqlDatabases Returns a list of Databases
    ListEc2MssqlDatabases(
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string)(
        *models.ListEC2MSSQLDatabasesResponse,  *apiutils.APIError)
    
    // ReadEc2MssqlDatabase Returns a representation of the specified database.
    ReadEc2MssqlDatabase(
        databaseId string)(
        *models.ReadEC2MSSQLDatabaseResponse,  *apiutils.APIError)
    
    // ListEc2MssqlDatabasePitrIntervals Returns a list of time intervals (start timestamp and end timestamp) in which the MSSQL database can be restored.
    ListEc2MssqlDatabasePitrIntervals(
        databaseId string, 
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListEC2MssqlDatabasePitrIntervalsResponse,  *apiutils.APIError)
    
}

// NewEc2MssqlDatabasesV1 returns Ec2MssqlDatabasesV1Client
func NewEc2MssqlDatabasesV1(config config.Config) Ec2MssqlDatabasesV1Client {
    client := new(Ec2MssqlDatabasesV1)
    client.config = config
    return client
}
