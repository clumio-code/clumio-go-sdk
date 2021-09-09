// Copyright (c) 2021 Clumio All Rights Reserved

package vmwarevcenternetworks

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// VmwareVcenterNetworksV1Client represents a custom type interface
type VmwareVcenterNetworksV1Client interface {
    //  Returns a list of networks in the specified vCenter server.
    ListVmwareVcenterNetworks(
        vcenterId string, 
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListVMwareVCenterNetworksResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified network.
    ReadVmwareVcenterNetwork(
        vcenterId string, 
        networkId string)(
        *models.ReadVMwareVCenterNetworkResponse,  *apiutils.APIError)
    
}

// NewVmwareVcenterNetworksV1 returns VmwareVcenterNetworksV1Client
func NewVmwareVcenterNetworksV1(config config.Config) VmwareVcenterNetworksV1Client{
    client := new(VmwareVcenterNetworksV1)
    client.config = config
    return client
}
