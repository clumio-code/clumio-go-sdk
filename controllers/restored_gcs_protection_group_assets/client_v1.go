// Copyright (c) 2023 Clumio All Rights Reserved

package restoredgcsprotectiongroupassets

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredGcsProtectionGroupAssetsV1Client represents a custom type interface
type RestoredGcsProtectionGroupAssetsV1Client interface {
    // RestoreGcsProtectionGroupAsset Restores the specified asset backup to the specified target destination.
    RestoreGcsProtectionGroupAsset(
        embed *string, 
        body models.RestoreGcsProtectionGroupAssetV1Request)(
        *models.RestoreGCSProtectionGroupBucketResponse,  *apiutils.APIError)
    
    // PreviewGcsProtectionGroupAsset Preview an asset restore
    PreviewGcsProtectionGroupAsset(
        gcsAssetId string, 
        body models.PreviewGcsProtectionGroupAssetV1Request)(
        *models.PreviewGCSProtectionGroupAssetAsyncResponse,  *apiutils.APIError)
    
    // PreviewDetailsGcsProtectionGroupAsset Details for asset restore preview
    PreviewDetailsGcsProtectionGroupAsset(
        gcsAssetId string, 
        previewId string)(
        *models.PreviewGCSProtectionGroupAssetDetailsResponse,  *apiutils.APIError)
    
}

// NewRestoredGcsProtectionGroupAssetsV1 returns RestoredGcsProtectionGroupAssetsV1Client
func NewRestoredGcsProtectionGroupAssetsV1(config config.Config) RestoredGcsProtectionGroupAssetsV1Client {
    client := new(RestoredGcsProtectionGroupAssetsV1)
    client.config = config
    return client
}
