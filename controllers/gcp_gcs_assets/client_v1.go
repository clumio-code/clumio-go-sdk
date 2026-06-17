// Copyright (c) 2023 Clumio All Rights Reserved

package gcpgcsassets

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpGcsAssetsV1Client represents a custom type interface
type GcpGcsAssetsV1Client interface {
    // ListGcpGcsAssets Returns a list of GCP GCS assets.
    ListGcpGcsAssets(
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string)(
        *models.ListGCPGCSAssetsResponse,  *apiutils.APIError)
    
    // ReadGcpGcsAsset Returns a representation of the specified GCP GCS asset.
    ReadGcpGcsAsset(
        gcsAssetId string, 
        embed *string)(
        *models.ReadGCPGCSAssetResponse,  *apiutils.APIError)
    
    // ReadGcpGcsAssetContinuousBackupStats Returns continuous backup statistics of the specified GCP GCS asset.
    ReadGcpGcsAssetContinuousBackupStats(
        gcsAssetId string, 
        beginTimestamp string, 
        endTimestamp string, 
        interval *string)(
        *models.ReadGCPGCSAssetContinuousBackupStatsResponse,  *apiutils.APIError)
    
    // ListGcpGcsAssetPitrIntervals Returns a list of time intervals (start timestamp and end timestamp) in which the GCS asset can be restored.
    ListGcpGcsAssetPitrIntervals(
        gcsAssetId string, 
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListGCPGCSAssetPitrIntervalsResponse,  *apiutils.APIError)
    
}

// NewGcpGcsAssetsV1 returns GcpGcsAssetsV1Client
func NewGcpGcsAssetsV1(config config.Config) GcpGcsAssetsV1Client {
    client := new(GcpGcsAssetsV1)
    client.config = config
    return client
}
