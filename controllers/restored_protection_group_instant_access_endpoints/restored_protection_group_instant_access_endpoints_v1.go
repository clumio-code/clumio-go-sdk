// Copyright (c) 2021 Clumio All Rights Reserved

// Package restoredprotectiongroupinstantaccessendpoints contains methods related to RestoredProtectionGroupInstantAccessEndpoints
package restoredprotectiongroupinstantaccessendpoints

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredProtectionGroupInstantAccessEndpointsV1 represents a custom type struct
type RestoredProtectionGroupInstantAccessEndpointsV1 struct {
    config config.Config
}

// ListProtectionGroupInstantAccessEndpoints Lists S3 instant access endpoints depending on the filters present in the body.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) ListProtectionGroupInstantAccessEndpoints(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListS3InstantAccessEndpointsResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/protection-groups/instant-access-endpoints"

    
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    result := &models.ListS3InstantAccessEndpointsResponse{}
    defaultInt64 := int64(0)
    defaultString := "" 
    
    if limit == nil {
        limit = &defaultInt64
    }
    if start == nil {
        start = &defaultString
    }
    if filter == nil {
        filter = &defaultString
    }
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreateProtectionGroupInstantAccessEndpoint Creates an endpoint that provides instant access to a protection group backup at a specific
//  backup point or arbitrary point-in-time if the S3 continuous backup is enabled.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) CreateProtectionGroupInstantAccessEndpoint(
    body models.CreateProtectionGroupInstantAccessEndpointV1Request)(
    *models.CreateS3InstantAccessEndpointResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/protection-groups/instant-access-endpoints"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    result := &models.CreateS3InstantAccessEndpointResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// CostEstimatesProtectionGroupInstantAccessEndpoint Estimates cost for creating and maintaining a S3 instant access endpoint with given parameters.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) CostEstimatesProtectionGroupInstantAccessEndpoint(
    body models.CostEstimatesProtectionGroupInstantAccessEndpointV1Request)(
    *models.CostEstimatesProtectionGroupInstantAccessEndpointResponseWrapper, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/protection-groups/instant-access-endpoints/cost-estimates"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    result := &models.CostEstimatesProtectionGroupInstantAccessEndpointResponseWrapper{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result200: &result.Http200,
        Result202: &result.Http202,
        RequestType: common.Post,
    })
    if result.Http200 != nil {
        result.StatusCode = 200
    } else if result.Http202 != nil {
        result.StatusCode = 202
    } 

    return result, apiErr
}


// CostEstimatesDetailsProtectionGroupInstantAccessEndpoint Details of the cost estimate for an S3 instant access endpoint from a given estimate ID.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) CostEstimatesDetailsProtectionGroupInstantAccessEndpoint(
    estimateId string)(
    *models.EstimateCostDetailsS3InstantAccessEndpointResponse, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/instant-access-endpoints/cost-estimates/{estimate_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "estimate_id": estimateId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    result := &models.EstimateCostDetailsS3InstantAccessEndpointResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadProtectionGroupInstantAccessEndpoint Return detailed information for a specified S3 instant access endpoint.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) ReadProtectionGroupInstantAccessEndpoint(
    endpointId string)(
    *models.ReadS3InstantAccessEndpointResponse, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/instant-access-endpoints/{endpoint_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "endpoint_id": endpointId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    result := &models.ReadS3InstantAccessEndpointResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// UpdateProtectionGroupInstantAccessEndpoint Updates an endpoint for S3 Instant Access.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) UpdateProtectionGroupInstantAccessEndpoint(
    endpointId string, 
    body models.UpdateProtectionGroupInstantAccessEndpointV1Request)(
    *models.UpdateS3InstantAccessEndpointResponse, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/instant-access-endpoints/{endpoint_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "endpoint_id": endpointId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    result := &models.UpdateS3InstantAccessEndpointResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Put,
    })

    return result, apiErr
}


// DeleteProtectionGroupInstantAccessEndpoint Deletes a S3 instant access endpoint and all its details.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) DeleteProtectionGroupInstantAccessEndpoint(
    endpointId string)(
    interface{}, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/instant-access-endpoints/{endpoint_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "endpoint_id": endpointId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    var result interface{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Delete,
    })

    return result, apiErr
}


// ReadProtectionGroupInstantAccessEndpointUri Reads the URI of S3 instant access endpoint.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) ReadProtectionGroupInstantAccessEndpointUri(
    endpointId string)(
    *models.ReadS3InstantAccessEndpointUriResponse, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/instant-access-endpoints/{endpoint_id}/_get_uri"
    //process optional template parameters
    pathParams := map[string]string{
        "endpoint_id": endpointId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    result := &models.ReadS3InstantAccessEndpointUriResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// AddProtectionGroupInstantAccessEndpointRole Adds a user-defined IAM role to allow access to an S3 Instant Access endpoint.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) AddProtectionGroupInstantAccessEndpointRole(
    endpointId string, 
    body models.AddProtectionGroupInstantAccessEndpointRoleV1Request)(
    *models.AddS3InstantAccessEndpointRoleResponse, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/instant-access-endpoints/{endpoint_id}/roles"
    //process optional template parameters
    pathParams := map[string]string{
        "endpoint_id": endpointId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    result := &models.AddS3InstantAccessEndpointRoleResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadProtectionGroupInstantAccessEndpointRolePermission Reads the permissions JSON string to be attached to the user-defined IAM role that can access
//  an S3 Instant Access endpoint.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) ReadProtectionGroupInstantAccessEndpointRolePermission(
    endpointId string)(
    *models.ReadS3InstantAccessEndpointRolePermissionResponse, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/instant-access-endpoints/{endpoint_id}/roles/permissions"
    //process optional template parameters
    pathParams := map[string]string{
        "endpoint_id": endpointId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    result := &models.ReadS3InstantAccessEndpointRolePermissionResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// UpdateProtectionGroupInstantAccessEndpointRole Updates a user-defined IAM role that is attached to an S3 Instant Access endpoint if any
//  changes are made to that role.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) UpdateProtectionGroupInstantAccessEndpointRole(
    endpointId string, 
    roleId string, 
    body models.UpdateProtectionGroupInstantAccessEndpointRoleV1Request)(
    *models.UpdateS3InstantAccessEndpointRoleResponse, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/instant-access-endpoints/{endpoint_id}/roles/{role_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "endpoint_id": endpointId,
        "role_id": roleId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    result := &models.UpdateS3InstantAccessEndpointRoleResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Put,
    })

    return result, apiErr
}


// DeleteProtectionGroupInstantAccessEndpointRole Deletes a user-defined IAM role attached to an S3 Instant Access endpoint.
func (r *RestoredProtectionGroupInstantAccessEndpointsV1) DeleteProtectionGroupInstantAccessEndpointRole(
    endpointId string, 
    roleId string)(
    *models.DeleteS3InstantAccessEndpointRoleResponse, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/instant-access-endpoints/{endpoint_id}/roles/{role_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "endpoint_id": endpointId,
        "role_id": roleId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.restored-protection-group-instant-access-endpoints=v1+json"
    result := &models.DeleteS3InstantAccessEndpointRoleResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Delete,
    })

    return result, apiErr
}
