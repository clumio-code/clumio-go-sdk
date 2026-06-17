// Copyright (c) 2023 Clumio All Rights Reserved

// Package gcpgcsassets contains methods related to GcpGcsAssets
package gcpgcsassets

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpGcsAssetsV1 represents a custom type struct
type GcpGcsAssetsV1 struct {
    config config.Config
}

// ListGcpGcsAssets Returns a list of GCP GCS assets.
func (g *GcpGcsAssetsV1) ListGcpGcsAssets(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListGCPGCSAssetsResponse, *apiutils.APIError) {

    queryBuilder := g.config.BaseUrl + "/datasources/gcp/gcs-assets"

    
    header := "application/api.clumio.gcp-gcs-assets=v1+json"
    result := &models.ListGCPGCSAssetsResponse{}
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
    if embed != nil {
        queryParams["embed"] = *embed
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


// ReadGcpGcsAsset Returns a representation of the specified GCP GCS asset.
func (g *GcpGcsAssetsV1) ReadGcpGcsAsset(
    gcsAssetId string, 
    embed *string)(
    *models.ReadGCPGCSAssetResponse, *apiutils.APIError) {

    pathURL := "/datasources/gcp/gcs-assets/{gcs_asset_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "gcs_asset_id": gcsAssetId,
    }
    queryBuilder := g.config.BaseUrl + pathURL

    
    header := "application/api.clumio.gcp-gcs-assets=v1+json"
    result := &models.ReadGCPGCSAssetResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
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


// ReadGcpGcsAssetContinuousBackupStats Returns continuous backup statistics of the specified GCP GCS asset.
func (g *GcpGcsAssetsV1) ReadGcpGcsAssetContinuousBackupStats(
    gcsAssetId string, 
    beginTimestamp string, 
    endTimestamp string, 
    interval *string)(
    *models.ReadGCPGCSAssetContinuousBackupStatsResponse, *apiutils.APIError) {

    pathURL := "/datasources/gcp/gcs-assets/{gcs_asset_id}/continuous-backup-stats"
    //process optional template parameters
    pathParams := map[string]string{
        "gcs_asset_id": gcsAssetId,
    }
    queryBuilder := g.config.BaseUrl + pathURL

    
    header := "application/api.clumio.gcp-gcs-assets=v1+json"
    result := &models.ReadGCPGCSAssetContinuousBackupStatsResponse{}
    queryParams := make(map[string]string)
    if interval != nil {
        queryParams["interval"] = *interval
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


// ListGcpGcsAssetPitrIntervals Returns a list of time intervals (start timestamp and end timestamp) in which the GCS asset can be restored.
func (g *GcpGcsAssetsV1) ListGcpGcsAssetPitrIntervals(
    gcsAssetId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListGCPGCSAssetPitrIntervalsResponse, *apiutils.APIError) {

    pathURL := "/datasources/gcp/gcs-assets/{gcs_asset_id}/pitr-intervals"
    //process optional template parameters
    pathParams := map[string]string{
        "gcs_asset_id": gcsAssetId,
    }
    queryBuilder := g.config.BaseUrl + pathURL

    
    header := "application/api.clumio.gcp-gcs-assets=v1+json"
    result := &models.ListGCPGCSAssetPitrIntervalsResponse{}
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
