// Copyright (c) 2023 Clumio All Rights Reserved

// Package restoredgcsprotectiongroupassets contains methods related to RestoredGcsProtectionGroupAssets
package restoredgcsprotectiongroupassets

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredGcsProtectionGroupAssetsV1 represents a custom type struct
type RestoredGcsProtectionGroupAssetsV1 struct {
    config config.Config
}

// RestoreGcsProtectionGroupAsset Restores the specified asset backup to the specified target destination.
func (r *RestoredGcsProtectionGroupAssetsV1) RestoreGcsProtectionGroupAsset(
    embed *string, 
    body models.RestoreGcsProtectionGroupAssetV1Request)(
    *models.RestoreGCSProtectionGroupBucketResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/gcp/protection-groups/assets"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-gcs-protection-group-assets=v1+json"
    result := &models.RestoreGCSProtectionGroupBucketResponse{}
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


// PreviewGcsProtectionGroupAsset Preview an asset restore
func (r *RestoredGcsProtectionGroupAssetsV1) PreviewGcsProtectionGroupAsset(
    gcsAssetId string, 
    body models.PreviewGcsProtectionGroupAssetV1Request)(
    *models.PreviewGCSProtectionGroupAssetAsyncResponse, *apiutils.APIError) {

    pathURL := "/restores/gcp/protection-groups/assets/{gcs_asset_id}/previews"
    //process optional template parameters
    pathParams := map[string]string{
        "gcs_asset_id": gcsAssetId,
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
    header := "application/api.clumio.restored-gcs-protection-group-assets=v1+json"
    result := &models.PreviewGCSProtectionGroupAssetAsyncResponse{}

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


// PreviewDetailsGcsProtectionGroupAsset Details for asset restore preview
func (r *RestoredGcsProtectionGroupAssetsV1) PreviewDetailsGcsProtectionGroupAsset(
    gcsAssetId string, 
    previewId string)(
    *models.PreviewGCSProtectionGroupAssetDetailsResponse, *apiutils.APIError) {

    pathURL := "/restores/gcp/protection-groups/assets/{gcs_asset_id}/previews/{preview_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "gcs_asset_id": gcsAssetId,
        "preview_id": previewId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.restored-gcs-protection-group-assets=v1+json"
    result := &models.PreviewGCSProtectionGroupAssetDetailsResponse{}

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
