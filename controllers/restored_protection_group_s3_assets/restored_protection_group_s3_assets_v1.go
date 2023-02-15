// Copyright (c) 2021 Clumio All Rights Reserved

// Package restoredprotectiongroups3assets contains methods related to RestoredProtectionGroupS3Assets
package restoredprotectiongroups3assets

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredProtectionGroupS3AssetsV1 represents a custom type struct
type RestoredProtectionGroupS3AssetsV1 struct {
    config config.Config
}

// RestoreProtectionGroupS3Asset Restores the specified protection group S3 asset backup to the specified target destination.
func (r *RestoredProtectionGroupS3AssetsV1) RestoreProtectionGroupS3Asset(
    embed *string, 
    body models.RestoreProtectionGroupS3AssetV1Request)(
    *models.RestoreProtectionGroupS3AssetResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/protection-groups/s3-assets"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-protection-group-s3-assets=v1+json"
    result := &models.RestoreProtectionGroupS3AssetResponse{}
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


// PreviewProtectionGroupS3Asset Preview a protection group S3 asset restore
func (r *RestoredProtectionGroupS3AssetsV1) PreviewProtectionGroupS3Asset(
    protectionGroupS3AssetId string, 
    body models.PreviewProtectionGroupS3AssetV1Request)(
    *models.PreviewProtectionGroupS3AssetResponseWrapper, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/s3-assets/{protection_group_s3_asset_id}/previews"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_s3_asset_id": protectionGroupS3AssetId,
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
    header := "application/api.clumio.restored-protection-group-s3-assets=v1+json"
    result := &models.PreviewProtectionGroupS3AssetResponseWrapper{}

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


// PreviewDetailsProtectionGroupS3Asset Details for protection group S3 asset restore preview
func (r *RestoredProtectionGroupS3AssetsV1) PreviewDetailsProtectionGroupS3Asset(
    protectionGroupS3AssetId string, 
    previewId string)(
    *models.PreviewProtectionGroupS3AssetDetailsResponse, *apiutils.APIError) {

    pathURL := "/restores/protection-groups/s3-assets/{protection_group_s3_asset_id}/previews/{preview_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_s3_asset_id": protectionGroupS3AssetId,
        "preview_id": previewId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.restored-protection-group-s3-assets=v1+json"
    result := &models.PreviewProtectionGroupS3AssetDetailsResponse{}

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
