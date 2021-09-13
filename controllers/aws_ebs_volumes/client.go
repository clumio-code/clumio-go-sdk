// Copyright (c) 2021 Clumio All Rights Reserved

package awsebsvolumes

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// AwsEbsVolumesV1Client represents a custom type interface
type AwsEbsVolumesV1Client interface {
    //  Returns a list of EBS volumes.
    ListAwsEbsVolumes(
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string)(
        *models.ListEbsVolumesResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified EBS volume.
    ReadAwsEbsVolume(
        volumeId string, 
        embed *string)(
        *models.ReadEbsVolumeResponse,  *apiutils.APIError)
    
}

// NewAwsEbsVolumesV1 returns AwsEbsVolumesV1Client
func NewAwsEbsVolumesV1(config config.Config) AwsEbsVolumesV1Client{
    client := new(AwsEbsVolumesV1)
    client.config = config
    return client
}
