// Copyright (c) 2021 Clumio All Rights Reserved

package organizationalunits

import (
     "github.com/clumio-code/clumiogosdk/api_utils"
     "github.com/clumio-code/clumiogosdk/config"
     "github.com/clumio-code/clumiogosdk/models"
)

// OrganizationalUnitsV1Client represents a custom type interface
type OrganizationalUnitsV1Client interface {
    //  Returns a list of organizational units.
    ListOrganizationalUnits(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListOrganizationalUnitsResponse,  *apiutils.APIError)
    
    //  Create a new organizational unit.
    CreateOrganizationalUnit(
        embed *string, 
        body *models.CreateOrganizationalUnitV1Request)(
        *models.CreateOrganizationalUnitResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified organizational unit.
    ReadOrganizationalUnit(
        id string)(
        *models.ReadOrganizationalUnitResponse,  *apiutils.APIError)
    
    //  Delete the specified organizational unit.
    DeleteOrganizationalUnit(
        id string, 
        embed *string)(
        *models.DeleteOrganizationalUnitResponse,  *apiutils.APIError)
    
    //  Patch the specified organizational unit.
    //  The complete updated attribute(s) of the organizational unit has to be provided in the request.
    PatchOrganizationalUnit(
        id string, 
        embed *string, 
        body *models.PatchOrganizationalUnitV1Request)(
        *models.PatchOrganizationalUnitResponse,  *apiutils.APIError)
    
}

// NewOrganizationalUnitsV1 returns OrganizationalUnitsV1Client
func NewOrganizationalUnitsV1(config config.Config) OrganizationalUnitsV1Client{
    client := new(OrganizationalUnitsV1)
    client.config = config
    return client
}
