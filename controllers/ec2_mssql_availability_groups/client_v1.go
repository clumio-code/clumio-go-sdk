// Copyright (c) 2021 Clumio All Rights Reserved

package ec2mssqlavailabilitygroups

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlAvailabilityGroupsV1Client represents a custom type interface
type Ec2MssqlAvailabilityGroupsV1Client interface {
    // ListEc2MssqlAvailabilityGroups Returns a list of Availability Groups.
    ListEc2MssqlAvailabilityGroups(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListEC2MssqlAGsResponse,  *apiutils.APIError)
    
    // ReadEc2MssqlAvailabilityGroup Returns a representation of the specified availability group.
    ReadEc2MssqlAvailabilityGroup(
        availabilityGroupId string)(
        *models.ReadEC2MssqlAGResponse,  *apiutils.APIError)
    
}

// NewEc2MssqlAvailabilityGroupsV1 returns Ec2MssqlAvailabilityGroupsV1Client
func NewEc2MssqlAvailabilityGroupsV1(config config.Config) Ec2MssqlAvailabilityGroupsV1Client {
    client := new(Ec2MssqlAvailabilityGroupsV1)
    client.config = config
    return client
}
