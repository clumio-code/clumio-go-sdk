// Copyright (c) 2021 Clumio All Rights Reserved

package vmwarevcenterdatastores

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// VmwareVcenterDatastoresV1Client represents a custom type interface
type VmwareVcenterDatastoresV1Client interface {
    //  Returns a list of datastores in the specified vCenter server.
    //  
    //  
    //  Supported Datastore Types
    //  Clumio supports the following VMware datastore types:
    //  
    //  Network File System (nfs)
    //  vSAN (virtual_storage_area_network)
    //  vSphere Virtual Machine File System (vmfs)
    //  Virtual Volumes (virtual_volume)
    //  
    ListVmwareVcenterDatastores(
        vcenterId string, 
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListVMwareDatastoresResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified datastore.
    ReadVmwareVcenterDatastore(
        vcenterId string, 
        datastoreId string)(
        *models.ReadVMwareDatastoreResponse,  *apiutils.APIError)
    
}

// NewVmwareVcenterDatastoresV1 returns VmwareVcenterDatastoresV1Client
func NewVmwareVcenterDatastoresV1(config config.Config) VmwareVcenterDatastoresV1Client{
    client := new(VmwareVcenterDatastoresV1)
    client.config = config
    return client
}
