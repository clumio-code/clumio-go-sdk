// Copyright (c) 2021 Clumio All Rights Reserved

package vmwarevcentertags

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// VmwareVcenterTagsV1Client represents a custom type interface
type VmwareVcenterTagsV1Client interface {
    //  Returns a list of tags in the specified vCenter server.
    ListVmwareVcenterTags(
        vcenterId string, 
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string)(
        *models.ListTagsResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified tag.
    ReadVmwareVcenterTag(
        vcenterId string, 
        tagId string, 
        embed *string)(
        *models.ReadTagResponse,  *apiutils.APIError)
    
}

// NewVmwareVcenterTagsV1 returns VmwareVcenterTagsV1Client
func NewVmwareVcenterTagsV1(config config.Config) VmwareVcenterTagsV1Client{
    client := new(VmwareVcenterTagsV1)
    client.config = config
    return client
}
