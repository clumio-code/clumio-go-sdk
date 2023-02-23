// Copyright (c) 2021 Clumio All Rights Reserved

package organizationalunits

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// OrganizationalUnitsV1Client represents a custom type interface
type OrganizationalUnitsV1Client interface {
    // ListOrganizationalUnits Returns a list of organizational units.
    ListOrganizationalUnits(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListOrganizationalUnitsResponseV1,  *apiutils.APIError)
    
    // CreateOrganizationalUnit Create a new organizational unit. Adding entities to the OU is an asynchronous operation and has a task associated.
    //  When the request has entities to be added, the response has a task ID which can be used to
    //  track the progress of the operation.
    CreateOrganizationalUnit(
        embed *string, 
        body *models.CreateOrganizationalUnitV1Request)(
        *models.CreateOrganizationalUnitV1ResponseWrapper,  *apiutils.APIError)
    
    // ReadOrganizationalUnit Returns a representation of the specified organizational unit.
    ReadOrganizationalUnit(
        id string, 
        embed *string)(
        *models.ReadOrganizationalUnitResponseV1,  *apiutils.APIError)
    
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
        body *models.PatchOrganizationalUnitV1Request)(
        *models.PatchOrganizationalUnitV1ResponseWrapper,  *apiutils.APIError)
    
}

// NewOrganizationalUnitsV1 returns OrganizationalUnitsV1Client
func NewOrganizationalUnitsV1(config config.Config) OrganizationalUnitsV1Client {
    client := new(OrganizationalUnitsV1)
    client.config = config
    return client
}
