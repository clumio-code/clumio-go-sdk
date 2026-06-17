// Copyright (c) 2023 Clumio All Rights Reserved

package gcpprotectiongroups

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpProtectionGroupsV1Client represents a custom type interface
type GcpProtectionGroupsV1Client interface {
    // ListGcpProtectionGroups Returns a list of GCP protection groups.
    ListGcpProtectionGroups(
        limit *int64, 
        start *string, 
        filter *string, 
        bucketUuidDetail *string, 
        lookbackDays *int64)(
        *models.ListGCPProtectionGroupsResponse,  *apiutils.APIError)
    
    // CreateGcpProtectionGroup Creates a new GCP protection group.
    CreateGcpProtectionGroup(
        body models.CreateGcpProtectionGroupV1Request)(
        *models.CreateGCPProtectionGroupResponse,  *apiutils.APIError)
    
    // ReadGcpProtectionGroup Returns a representation of the specified GCP protection group.
    ReadGcpProtectionGroup(
        protectionGroupId string, 
        lookbackDays *int64)(
        *models.ReadGCPProtectionGroupResponse,  *apiutils.APIError)
    
    // DeleteGcpProtectionGroup Deletes the specified GCP protection group.
    DeleteGcpProtectionGroup(
        protectionGroupId string)(
        interface{},  *apiutils.APIError)
    
    // UpdateGcpProtectionGroup Updates an existing GCP protection group.
    UpdateGcpProtectionGroup(
        protectionGroupId string, 
        body models.UpdateGcpProtectionGroupV1Request)(
        *models.UpdateGCPProtectionGroupResponse,  *apiutils.APIError)
    
}

// NewGcpProtectionGroupsV1 returns GcpProtectionGroupsV1Client
func NewGcpProtectionGroupsV1(config config.Config) GcpProtectionGroupsV1Client {
    client := new(GcpProtectionGroupsV1)
    client.config = config
    return client
}
