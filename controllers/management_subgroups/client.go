// Copyright (c) 2021 Clumio All Rights Reserved

package managementsubgroups

import (
     "github.com/clumio-code/clumiogosdk/api_utils"
     "github.com/clumio-code/clumiogosdk/config"
     "github.com/clumio-code/clumiogosdk/models"
)

// ManagementSubgroupsV1Client represents a custom type interface
type ManagementSubgroupsV1Client interface {
    //  Returns a list of subgroups.
    ListManagementSubgroups(
        groupId string, 
        limit *int64, 
        start *string)(
        *models.ListSubgroupsResponse,  *apiutils.APIError)
    
    //  Subgroups are used to manage cloud connectors and SQL hosts residing in the same vCenter server.
    //  
    //  Returns a representation of the specified subgroups.
    ReadManagementSubgroup(
        subgroupId string, 
        groupId string)(
        *models.ReadSubgroupResponse,  *apiutils.APIError)
    
    //  Update the specified subgroup.
    UpdateManagementSubgroup(
        subgroupId string, 
        groupId string, 
        body *models.UpdateManagementSubgroupV1Request)(
        *models.UpdateSubgroupResponse,  *apiutils.APIError)
    
}

// NewManagementSubgroupsV1 returns ManagementSubgroupsV1Client
func NewManagementSubgroupsV1(config config.Config) ManagementSubgroupsV1Client{
    client := new(ManagementSubgroupsV1)
    client.config = config
    return client
}
