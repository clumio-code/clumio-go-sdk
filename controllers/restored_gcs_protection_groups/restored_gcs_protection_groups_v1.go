// Copyright (c) 2023 Clumio All Rights Reserved

// Package restoredgcsprotectiongroups contains methods related to RestoredGcsProtectionGroups
package restoredgcsprotectiongroups

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredGcsProtectionGroupsV1 represents a custom type struct
type RestoredGcsProtectionGroupsV1 struct {
    config config.Config
}

// RestoreGcsProtectionGroup Restores the specified protection group backup to the specified target destination.
func (r *RestoredGcsProtectionGroupsV1) RestoreGcsProtectionGroup(
    embed *string, 
    body models.RestoreGcsProtectionGroupV1Request)(
    *models.RestoreGCSProtectionGroupResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/gcp/protection-groups"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-gcs-protection-groups=v1+json"
    result := &models.RestoreGCSProtectionGroupResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
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


// RestoreGcsProtectionGroupObjects Restores the specified list of objects to the specified target destination.
func (r *RestoredGcsProtectionGroupsV1) RestoreGcsProtectionGroupObjects(
    gcsProtectionGroupId string, 
    embed *string, 
    body models.RestoreGcsProtectionGroupObjectsV1Request)(
    *models.RestoreGCSObjectsResponse, *apiutils.APIError) {

    pathURL := "/restores/gcp/protection-groups/{gcs_protection_group_id}/gcs-objects"
    //process optional template parameters
    pathParams := map[string]string{
        "gcs_protection_group_id": gcsProtectionGroupId,
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
    header := "application/api.clumio.restored-gcs-protection-groups=v1+json"
    result := &models.RestoreGCSObjectsResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
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


// PreviewGcsProtectionGroup Preview a protection group restore.
func (r *RestoredGcsProtectionGroupsV1) PreviewGcsProtectionGroup(
    gcsProtectionGroupId string, 
    body models.PreviewGcsProtectionGroupV1Request)(
    *models.PreviewGCSProtectionGroupAsyncResponse, *apiutils.APIError) {

    pathURL := "/restores/gcp/protection-groups/{gcs_protection_group_id}/previews"
    //process optional template parameters
    pathParams := map[string]string{
        "gcs_protection_group_id": gcsProtectionGroupId,
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
    header := "application/api.clumio.restored-gcs-protection-groups=v1+json"
    result := &models.PreviewGCSProtectionGroupAsyncResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// PreviewDetailsGcsProtectionGroup Details for protection group restore preview
func (r *RestoredGcsProtectionGroupsV1) PreviewDetailsGcsProtectionGroup(
    gcsProtectionGroupId string, 
    previewId string)(
    *models.PreviewDetailsGCSProtectionGroupResponse, *apiutils.APIError) {

    pathURL := "/restores/gcp/protection-groups/{gcs_protection_group_id}/previews/{preview_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "gcs_protection_group_id": gcsProtectionGroupId,
        "preview_id": previewId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.restored-gcs-protection-groups=v1+json"
    result := &models.PreviewDetailsGCSProtectionGroupResponse{}

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
