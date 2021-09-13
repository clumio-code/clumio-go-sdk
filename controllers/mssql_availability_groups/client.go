// Copyright (c) 2021 Clumio All Rights Reserved

package mssqlavailabilitygroups

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// MssqlAvailabilityGroupsV1Client represents a custom type interface
type MssqlAvailabilityGroupsV1Client interface {
    //  Returns a list of Availability Groups.
    ListMssqlAvailabilityGroups(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListMssqlAGsResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified availability group.
    ReadMssqlAvailabilityGroup(
        availabilityGroupId string)(
        *models.ReadMssqlAGResponse,  *apiutils.APIError)
    
}

// NewMssqlAvailabilityGroupsV1 returns MssqlAvailabilityGroupsV1Client
func NewMssqlAvailabilityGroupsV1(config config.Config) MssqlAvailabilityGroupsV1Client{
    client := new(MssqlAvailabilityGroupsV1)
    client.config = config
    return client
}
