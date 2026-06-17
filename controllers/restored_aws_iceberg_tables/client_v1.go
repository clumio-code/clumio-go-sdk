// Copyright (c) 2023 Clumio All Rights Reserved

package restoredawsicebergtables

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredAwsIcebergTablesV1Client represents a custom type interface
type RestoredAwsIcebergTablesV1Client interface {
    // RestoreAwsIcebergTable Restores the specified source AWS Iceberg Table snapshots to the specified target destination. The source AWS Iceberg Table must be one that was backed up by Clumio.
    RestoreAwsIcebergTable(
        body models.RestoreAwsIcebergTableV1Request)(
        *models.RestoreAWSIcebergTableResponse,  *apiutils.APIError)
    
}

// NewRestoredAwsIcebergTablesV1 returns RestoredAwsIcebergTablesV1Client
func NewRestoredAwsIcebergTablesV1(config config.Config) RestoredAwsIcebergTablesV1Client {
    client := new(RestoredAwsIcebergTablesV1)
    client.config = config
    return client
}
