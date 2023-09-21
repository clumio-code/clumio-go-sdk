// Copyright (c) 2021 Clumio All Rights Reserved

package restoredprotectiongroupinstantaccessendpoints

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredProtectionGroupInstantAccessEndpointsV1Client represents a custom type interface
type RestoredProtectionGroupInstantAccessEndpointsV1Client interface {
    // ListProtectionGroupInstantAccessEndpoints Lists S3 instant access endpoints depending on the filters present in the body.
    ListProtectionGroupInstantAccessEndpoints(
        filter *string)(
        *models.ListS3InstantAccessEndpointsResponse,  *apiutils.APIError)
    
    // CreateProtectionGroupInstantAccessEndpoint Creates an endpoint that provides instant access to a protection group backup at a specific
    //  backup point or arbitrary point-in-time if the S3 continuous backup is enabled.
    CreateProtectionGroupInstantAccessEndpoint(
        body models.CreateProtectionGroupInstantAccessEndpointV1Request)(
        *models.CreateS3InstantAccessEndpointResponse,  *apiutils.APIError)
    
    // CostEstimatesProtectionGroupInstantAccessEndpoint Estimates cost for creating and maintaining a S3 instant access endpoint with given parameters.
    CostEstimatesProtectionGroupInstantAccessEndpoint(
        body models.CostEstimatesProtectionGroupInstantAccessEndpointV1Request)(
        *models.CostEstimatesProtectionGroupInstantAccessEndpointResponseWrapper,  *apiutils.APIError)
    
    // CostEstimatesDetailsProtectionGroupInstantAccessEndpoint Details of the cost estimate for an S3 instant access endpoint from a given estimate ID.
    CostEstimatesDetailsProtectionGroupInstantAccessEndpoint(
        estimateId string)(
        *models.EstimateCostDetailsS3InstantAccessEndpointResponse,  *apiutils.APIError)
    
    // ReadProtectionGroupInstantAccessEndpoint Return detailed information for a specified S3 instant access endpoint.
    ReadProtectionGroupInstantAccessEndpoint(
        endpointId string)(
        *models.ReadS3InstantAccessEndpointResponse,  *apiutils.APIError)
    
    // UpdateProtectionGroupInstantAccessEndpoint Updates an endpoint for S3 Instant Access.
    UpdateProtectionGroupInstantAccessEndpoint(
        endpointId string, 
        body models.UpdateProtectionGroupInstantAccessEndpointV1Request)(
        *models.UpdateS3InstantAccessEndpointResponse,  *apiutils.APIError)
    
    // DeleteProtectionGroupInstantAccessEndpoint Deletes a S3 instant access endpoint and all its details.
    DeleteProtectionGroupInstantAccessEndpoint(
        endpointId string)(
        interface{},  *apiutils.APIError)
    
    // ReadProtectionGroupInstantAccessEndpointUri Reads the URI of S3 instant access endpoint.
    ReadProtectionGroupInstantAccessEndpointUri(
        endpointId string)(
        *models.ReadS3InstantAccessEndpointUriResponse,  *apiutils.APIError)
    
    // AddProtectionGroupInstantAccessEndpointRole Adds a user-defined IAM role to allow access to an S3 Instant Access endpoint.
    AddProtectionGroupInstantAccessEndpointRole(
        endpointId string, 
        body models.AddProtectionGroupInstantAccessEndpointRoleV1Request)(
        *models.AddS3InstantAccessEndpointRoleResponse,  *apiutils.APIError)
    
    // ReadProtectionGroupInstantAccessEndpointRolePermission Reads the permissions JSON string to be attached to the user-defined IAM role that can access
    //  an S3 Instant Access endpoint.
    ReadProtectionGroupInstantAccessEndpointRolePermission(
        endpointId string)(
        *models.ReadS3InstantAccessEndpointRolePermissionResponse,  *apiutils.APIError)
    
    // UpdateProtectionGroupInstantAccessEndpointRole Updates a user-defined IAM role that is attached to an S3 Instant Access endpoint if any
    //  changes are made to that role.
    UpdateProtectionGroupInstantAccessEndpointRole(
        endpointId string, 
        roleId string, 
        body models.UpdateProtectionGroupInstantAccessEndpointRoleV1Request)(
        *models.UpdateS3InstantAccessEndpointRoleResponse,  *apiutils.APIError)
    
    // DeleteProtectionGroupInstantAccessEndpointRole Deletes a user-defined IAM role attached to an S3 Instant Access endpoint.
    DeleteProtectionGroupInstantAccessEndpointRole(
        endpointId string, 
        roleId string)(
        *models.DeleteS3InstantAccessEndpointRoleResponse,  *apiutils.APIError)
    
}

// NewRestoredProtectionGroupInstantAccessEndpointsV1 returns RestoredProtectionGroupInstantAccessEndpointsV1Client
func NewRestoredProtectionGroupInstantAccessEndpointsV1(config config.Config) RestoredProtectionGroupInstantAccessEndpointsV1Client {
    client := new(RestoredProtectionGroupInstantAccessEndpointsV1)
    client.config = config
    return client
}
