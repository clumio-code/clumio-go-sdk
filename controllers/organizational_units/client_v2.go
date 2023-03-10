// Copyright (c) 2021 Clumio All Rights Reserved

package organizationalunits

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// OrganizationalUnitsV2Client represents a custom type interface
type OrganizationalUnitsV2Client interface {
    // ListOrganizationalUnits Returns a list of organizational units.
    ListOrganizationalUnits(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListOrganizationalUnitsResponse,  *apiutils.APIError)
    
    // CreateOrganizationalUnit Create a new organizational unit. Adding entities to the OU is an asynchronous operation and has a task associated.
    //  When the request has entities to be added, the response has a task ID which can be used to
    //  track the progress of the operation.
    CreateOrganizationalUnit(
        embed *string, 
        body *models.CreateOrganizationalUnitV2Request)(
        *models.CreateOrganizationalUnitResponseWrapper,  *apiutils.APIError)
    
    // ReadOrganizationalUnit Returns a representation of the specified organizational unit.
    ReadOrganizationalUnit(
        id string, 
        embed *string)(
        *models.ReadOrganizationalUnitResponse,  *apiutils.APIError)
    
    // DeleteOrganizationalUnit Delete the specified organizational unit.
    DeleteOrganizationalUnit(
        id string, 
        embed *string)(
        *models.DeleteOrganizationalUnitResponse,  *apiutils.APIError)
    
    // PatchOrganizationalUnit Patch the specified organizational unit.
    //  The complete updated attribute(s) of the organizational unit have to be provided in the request.
    //  Adding or removing entities from the OU is an asynchronous operation and has a task associated.
    //  When the request has entities to be added or removed, the response has a task ID
    //  which can be used to track the progress of the operation.
    PatchOrganizationalUnit(
        id string, 
        embed *string, 
        body *models.PatchOrganizationalUnitV2Request)(
        *models.PatchOrganizationalUnitResponseWrapper,  *apiutils.APIError)
    
}

// NewOrganizationalUnitsV2 returns OrganizationalUnitsV2Client
func NewOrganizationalUnitsV2(config config.Config) OrganizationalUnitsV2Client {
    client := new(OrganizationalUnitsV2)
    client.config = config
    return client
}
