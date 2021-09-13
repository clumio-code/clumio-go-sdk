// Copyright (c) 2021 Clumio All Rights Reserved

package managementgroups

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// ManagementGroupsV1Client represents a custom type interface
type ManagementGroupsV1Client interface {
    //  Returns a list of management groups.
    ListManagementGroups(
        limit *int64, 
        start *string)(
        *models.ListManagementGroupsResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified management group. Management groups are used to
    //  manage the SQL hosts and cloud connectors deployed in vCenter servers.
    //  
    //  Returns a representation of the specified management-groups.
    ReadManagementGroup(
        groupId string)(
        *models.ReadManagementGroupResponse,  *apiutils.APIError)
    
    //  Update the specified management group.
    UpdateManagementGroup(
        groupId string, 
        body *models.UpdateManagementGroupV1Request)(
        *models.UpdateManagementGroupResponse,  *apiutils.APIError)
    
}

// NewManagementGroupsV1 returns ManagementGroupsV1Client
func NewManagementGroupsV1(config config.Config) ManagementGroupsV1Client{
    client := new(ManagementGroupsV1)
    client.config = config
    return client
}
