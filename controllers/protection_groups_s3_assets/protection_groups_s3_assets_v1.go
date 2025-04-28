// Copyright (c) 2023 Clumio All Rights Reserved

// Package protectiongroupss3assets contains methods related to ProtectionGroupsS3Assets
package protectiongroupss3assets

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// ProtectionGroupsS3AssetsV1 represents a custom type struct
type ProtectionGroupsS3AssetsV1 struct {
    config config.Config
}

// ListProtectionGroupS3Assets Returns a list of protection group S3 assets.
func (p *ProtectionGroupsS3AssetsV1) ListProtectionGroupS3Assets(
    limit *int64, 
    start *string, 
    filter *string, 
    lookbackDays *int64)(
    *models.ListProtectionGroupS3AssetsResponse, *apiutils.APIError) {

    queryBuilder := p.config.BaseUrl + "/datasources/protection-groups/s3-assets"

    
    header := "application/api.clumio.protection-groups-s3-assets=v1+json"
    result := &models.ListProtectionGroupS3AssetsResponse{}
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
    if lookbackDays != nil {
        queryParams["lookback_days"] = fmt.Sprintf("%v", *lookbackDays)
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: p.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadProtectionGroupS3Asset Returns a representation of the specified protection group S3 asset.
func (p *ProtectionGroupsS3AssetsV1) ReadProtectionGroupS3Asset(
    protectionGroupS3AssetId string, 
    lookbackDays *int64)(
    *models.ReadProtectionGroupS3AssetResponse, *apiutils.APIError) {

    pathURL := "/datasources/protection-groups/s3-assets/{protection_group_s3_asset_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_s3_asset_id": protectionGroupS3AssetId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    
    header := "application/api.clumio.protection-groups-s3-assets=v1+json"
    result := &models.ReadProtectionGroupS3AssetResponse{}
    queryParams := make(map[string]string)
    if lookbackDays != nil {
        queryParams["lookback_days"] = fmt.Sprintf("%v", *lookbackDays)
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: p.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadProtectionGroupS3AssetContinuousBackupStats Returns continuous backup statistics of the specified protection group S3 asset.
func (p *ProtectionGroupsS3AssetsV1) ReadProtectionGroupS3AssetContinuousBackupStats(
    protectionGroupS3AssetId string, 
    bucketName *string, 
    bucketId *string, 
    beginTimestamp string, 
    endTimestamp string, 
    interval *string)(
    *models.ReadProtectionGroupS3AssetContinuousBackupStatsResponse, *apiutils.APIError) {

    pathURL := "/datasources/protection-groups/s3-assets/{protection_group_s3_asset_id}/continuous-backup-stats"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_s3_asset_id": protectionGroupS3AssetId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    
    header := "application/api.clumio.protection-groups-s3-assets=v1+json"
    result := &models.ReadProtectionGroupS3AssetContinuousBackupStatsResponse{}
    queryParams := make(map[string]string)
    if bucketName != nil {
        queryParams["bucket_name"] = *bucketName
    }
    if bucketId != nil {
        queryParams["bucket_id"] = *bucketId
    }
    if interval != nil {
        queryParams["interval"] = *interval
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: p.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ListProtectionGroupS3AssetPitrIntervals Returns a list of time intervals (start timestamp and end timestamp) in which the protection group S3 asset can be restored.
func (p *ProtectionGroupsS3AssetsV1) ListProtectionGroupS3AssetPitrIntervals(
    protectionGroupS3AssetId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListProtectionGroupS3AssetPitrIntervalsResponse, *apiutils.APIError) {

    pathURL := "/datasources/protection-groups/s3-assets/{protection_group_s3_asset_id}/pitr-intervals"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_s3_asset_id": protectionGroupS3AssetId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    
    header := "application/api.clumio.protection-groups-s3-assets=v1+json"
    result := &models.ListProtectionGroupS3AssetPitrIntervalsResponse{}
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
        Config: p.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
