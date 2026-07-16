// Copyright (c) 2023 Clumio All Rights Reserved

package restoredawsdocumentdb

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredAwsDocumentdbV1Client represents a custom type interface
type RestoredAwsDocumentdbV1Client interface {
    // RestoreAwsDocumentdb Restores the specified DocumentDB cluster backup or snapshot to the specified target destination.
    RestoreAwsDocumentdb(
        embed *string, 
        body models.RestoreAwsDocumentdbV1Request)(
        *models.CreateDocumentDBRestoreResponse,  *apiutils.APIError)
    
}

// NewRestoredAwsDocumentdbV1 returns RestoredAwsDocumentdbV1Client
func NewRestoredAwsDocumentdbV1(config config.Config) RestoredAwsDocumentdbV1Client {
    client := new(RestoredAwsDocumentdbV1)
    client.config = config
    return client
}
