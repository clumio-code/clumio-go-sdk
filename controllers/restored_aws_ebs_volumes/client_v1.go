// Copyright (c) 2023 Clumio All Rights Reserved

package restoredawsebsvolumes

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredAwsEbsVolumesV1Client represents a custom type interface
type RestoredAwsEbsVolumesV1Client interface {
    // RestoreAwsEbsVolume TODO: Add comment
    RestoreAwsEbsVolume(
        body models.RestoreAwsEbsVolumeV1Request)(
        *models.RestoreEBSResponseV1,  *apiutils.APIError)
    
}

// NewRestoredAwsEbsVolumesV1 returns RestoredAwsEbsVolumesV1Client
func NewRestoredAwsEbsVolumesV1(config config.Config) RestoredAwsEbsVolumesV1Client {
    client := new(RestoredAwsEbsVolumesV1)
    client.config = config
    return client
}
