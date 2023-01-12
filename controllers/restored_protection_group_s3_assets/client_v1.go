// Copyright (c) 2021 Clumio All Rights Reserved

package restoredprotectiongroups3assets

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredProtectionGroupS3AssetsV1Client represents a custom type interface
type RestoredProtectionGroupS3AssetsV1Client interface {
    // RestoreProtectionGroupS3Asset Restores the specified protection group S3 asset backup to the specified target destination.
    RestoreProtectionGroupS3Asset(
        embed *string, 
        body models.RestoreProtectionGroupS3AssetV1Request)(
        *models.RestoreProtectionGroupS3AssetResponse,  *apiutils.APIError)
    
    // PreviewProtectionGroupS3Asset Preview a protection group S3 asset restore
    PreviewProtectionGroupS3Asset(
        protectionGroupS3AssetId string, 
        body models.PreviewProtectionGroupS3AssetV1Request)(
        *models.PreviewProtectionGroupS3AssetSyncResponse,  *apiutils.APIError)
    
    // PreviewDetailsProtectionGroupS3Asset Details for protection group S3 asset restore preview
    PreviewDetailsProtectionGroupS3Asset(
        protectionGroupS3AssetId string, 
        previewId string)(
        *models.PreviewProtectionGroupS3AssetDetailsResponse,  *apiutils.APIError)
    
}

// NewRestoredProtectionGroupS3AssetsV1 returns RestoredProtectionGroupS3AssetsV1Client
func NewRestoredProtectionGroupS3AssetsV1(config config.Config) RestoredProtectionGroupS3AssetsV1Client {
    client := new(RestoredProtectionGroupS3AssetsV1)
    client.config = config
    return client
}
