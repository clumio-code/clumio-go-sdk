// Copyright (c) 2021 Clumio All Rights Reserved

package vmwarevcenters

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// VmwareVcentersV1Client represents a custom type interface
type VmwareVcentersV1Client interface {
    //  Returns a list of vCenter servers.
    ListVmwareVcenters(
        limit *int64, 
        start *string, 
        embed *string)(
        *models.ListVcentersResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified vCenter server.
    ReadVmwareVcenter(
        vcenterId string, 
        embed *string)(
        *models.ReadVcenterResponse,  *apiutils.APIError)
    
}

// NewVmwareVcentersV1 returns VmwareVcentersV1Client
func NewVmwareVcentersV1(config config.Config) VmwareVcentersV1Client{
    client := new(VmwareVcentersV1)
    client.config = config
    return client
}
