// Copyright (c) 2021 Clumio All Rights Reserved

package vmwarevcenterdatacenters

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// VmwareVcenterDatacentersV1Client represents a custom type interface
type VmwareVcenterDatacentersV1Client interface {
    //  Returns a list of VMware data centers in the specified vCenter server.
    ListVmwareVcenterDatacenters(
        vcenterId string, 
        limit *int64, 
        start *string, 
        filter *string, 
        Embed *string)(
        *models.ListDatacentersResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified VMware data center within the specified vCenter server.
    ReadVmwareVcenterDatacenter(
        vcenterId string, 
        datacenterId string, 
        embed *string)(
        *models.ReadDatacenterResponse,  *apiutils.APIError)
    
}

// NewVmwareVcenterDatacentersV1 returns VmwareVcenterDatacentersV1Client
func NewVmwareVcenterDatacentersV1(config config.Config) VmwareVcenterDatacentersV1Client{
    client := new(VmwareVcenterDatacentersV1)
    client.config = config
    return client
}
