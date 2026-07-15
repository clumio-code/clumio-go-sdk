// Copyright (c) 2023 Clumio All Rights Reserved

package restoredawsneptune

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredAwsNeptuneV1Client represents a custom type interface
type RestoredAwsNeptuneV1Client interface {
    // RestoreAwsNeptune Restores the specified Neptune cluster backup or snapshot to the specified target destination.
    RestoreAwsNeptune(
        embed *string, 
        body models.RestoreAwsNeptuneV1Request)(
        *models.CreateNeptuneRestoreResponse,  *apiutils.APIError)
    
}

// NewRestoredAwsNeptuneV1 returns RestoredAwsNeptuneV1Client
func NewRestoredAwsNeptuneV1(config config.Config) RestoredAwsNeptuneV1Client {
    client := new(RestoredAwsNeptuneV1)
    client.config = config
    return client
}
