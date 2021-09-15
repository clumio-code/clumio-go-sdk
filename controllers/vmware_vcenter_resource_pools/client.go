// Copyright (c) 2021 Clumio All Rights Reserved

package vmwarevcenterresourcepools

import (
     "github.com/clumio-code/clumiogosdk/api_utils"
     "github.com/clumio-code/clumiogosdk/config"
     "github.com/clumio-code/clumiogosdk/models"
)

// VmwareVcenterResourcePoolsV1Client represents a custom type interface
type VmwareVcenterResourcePoolsV1Client interface {
    //  Returns a list of resource pools in the specified vCenter server.
    ListVmwareVcenterResourcePools(
        vcenterId string, 
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListResourcePoolsResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified resource pool.
    ReadVmwareVcenterResourcePool(
        vcenterId string, 
        resourcePoolId string)(
        *models.ReadResourcePoolResponse,  *apiutils.APIError)
    
}

// NewVmwareVcenterResourcePoolsV1 returns VmwareVcenterResourcePoolsV1Client
func NewVmwareVcenterResourcePoolsV1(config config.Config) VmwareVcenterResourcePoolsV1Client{
    client := new(VmwareVcenterResourcePoolsV1)
    client.config = config
    return client
}
