// Copyright (c) 2021 Clumio All Rights Reserved

// Package restoredprotectiongroups contains methods related to RestoredProtectionGroups
package restoredprotectiongroups

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredProtectionGroupsV1 represents a custom type struct
type RestoredProtectionGroupsV1 struct {
    config config.Config
}

// RestoreProtectionGroup Restores the specified protection group backup to the specified target destination.
func (r *RestoredProtectionGroupsV1) RestoreProtectionGroup(
    embed *string, 
    body models.RestoreProtectionGroupV1Request)(
    *models.RestoreProtectionGroupResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/protection-groups"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-protection-groups=v1+json"
    result := &models.RestoreProtectionGroupResponse{}
    defaultString := "" 
    
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
        "embed": *embed,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// PreviewProtectionGroup Preview a protection group restore.
func (r *RestoredProtectionGroupsV1) PreviewProtectionGroup(
    protectionGroupId string, 
    body models.PreviewProtectionGroupV1Request)(
    *models.PreviewProtectionGroupResponseWrapper, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/{protection_group_id}/previews"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_id": protectionGroupId,
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
    header := "application/api.clumio.restored-protection-groups=v1+json"
    result := &models.PreviewProtectionGroupResponseWrapper{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
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


// PreviewDetailsProtectionGroup Details for protection group bucket restore preview
func (r *RestoredProtectionGroupsV1) PreviewDetailsProtectionGroup(
    protectionGroupId string, 
    previewId string)(
    *models.PreviewDetailsProtectionGroupResponse, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/{protection_group_id}/previews/{preview_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_id": protectionGroupId,
        "preview_id": previewId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.restored-protection-groups=v1+json"
    result := &models.PreviewDetailsProtectionGroupResponse{}

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


// RestoreProtectionGroupS3Objects Restores the specified list of S3 objects to the specified target destination.
func (r *RestoredProtectionGroupsV1) RestoreProtectionGroupS3Objects(
    protectionGroupId string, 
    embed *string, 
    body models.RestoreProtectionGroupS3ObjectsV1Request)(
    *models.RestoreObjectsResponse, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/{protection_group_id}/s3-objects"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_id": protectionGroupId,
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
    header := "application/api.clumio.restored-protection-groups=v1+json"
    result := &models.RestoreObjectsResponse{}
    defaultString := "" 
    
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
        "embed": *embed,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}
