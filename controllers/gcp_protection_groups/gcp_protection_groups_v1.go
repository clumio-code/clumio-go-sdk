// Copyright (c) 2023 Clumio All Rights Reserved

// Package gcpprotectiongroups contains methods related to GcpProtectionGroups
package gcpprotectiongroups

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpProtectionGroupsV1 represents a custom type struct
type GcpProtectionGroupsV1 struct {
    config config.Config
}

// ListGcpProtectionGroups Returns a list of GCP protection groups.
func (g *GcpProtectionGroupsV1) ListGcpProtectionGroups(
    limit *int64, 
    start *string, 
    filter *string, 
    bucketUuidDetail *string, 
    lookbackDays *int64)(
    *models.ListGCPProtectionGroupsResponse, *apiutils.APIError) {

    queryBuilder := g.config.BaseUrl + "/datasources/gcp/protection-groups"

    
    header := "application/api.clumio.gcp-protection-groups=v1+json"
    result := &models.ListGCPProtectionGroupsResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
    }
    if filter != nil {
        queryParams["filter"] = *filter
    }
    if bucketUuidDetail != nil {
        queryParams["bucket_uuid_detail"] = *bucketUuidDetail
    }
    if lookbackDays != nil {
        queryParams["lookback_days"] = fmt.Sprintf("%v", *lookbackDays)
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreateGcpProtectionGroup Creates a new GCP protection group.
func (g *GcpProtectionGroupsV1) CreateGcpProtectionGroup(
    body models.CreateGcpProtectionGroupV1Request)(
    *models.CreateGCPProtectionGroupResponse, *apiutils.APIError) {

    queryBuilder := g.config.BaseUrl + "/datasources/gcp/protection-groups"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.gcp-protection-groups=v1+json"
    result := &models.CreateGCPProtectionGroupResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result201: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadGcpProtectionGroup Returns a representation of the specified GCP protection group.
func (g *GcpProtectionGroupsV1) ReadGcpProtectionGroup(
    protectionGroupId string, 
    lookbackDays *int64)(
    *models.ReadGCPProtectionGroupResponse, *apiutils.APIError) {

    pathURL := "/datasources/gcp/protection-groups/{protection_group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_id": protectionGroupId,
    }
    queryBuilder := g.config.BaseUrl + pathURL

    
    header := "application/api.clumio.gcp-protection-groups=v1+json"
    result := &models.ReadGCPProtectionGroupResponse{}
    queryParams := make(map[string]string)
    if lookbackDays != nil {
        queryParams["lookback_days"] = fmt.Sprintf("%v", *lookbackDays)
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// DeleteGcpProtectionGroup Deletes the specified GCP protection group.
func (g *GcpProtectionGroupsV1) DeleteGcpProtectionGroup(
    protectionGroupId string)(
    interface{}, *apiutils.APIError) {

    pathURL := "/datasources/gcp/protection-groups/{protection_group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_id": protectionGroupId,
    }
    queryBuilder := g.config.BaseUrl + pathURL

    
    header := "application/api.clumio.gcp-protection-groups=v1+json"
    var result interface{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        RequestType: common.Delete,
    })

    return result, apiErr
}


// UpdateGcpProtectionGroup Updates an existing GCP protection group.
func (g *GcpProtectionGroupsV1) UpdateGcpProtectionGroup(
    protectionGroupId string, 
    body models.UpdateGcpProtectionGroupV1Request)(
    *models.UpdateGCPProtectionGroupResponse, *apiutils.APIError) {

    pathURL := "/datasources/gcp/protection-groups/{protection_group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_id": protectionGroupId,
    }
    queryBuilder := g.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.gcp-protection-groups=v1+json"
    result := &models.UpdateGCPProtectionGroupResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Patch,
    })

    return result, apiErr
}
