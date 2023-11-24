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
    filter *string)(
    *models.ListProtectionGroupS3AssetsResponse, *apiutils.APIError) {

    queryBuilder := p.config.BaseUrl + "/datasources/protection-groups/s3-assets"

    
    header := "application/api.clumio.protection-groups-s3-assets=v1+json"
    result := &models.ListProtectionGroupS3AssetsResponse{}
    defaultInt64 := int64(0)
    defaultString := "" 
    
    if limit == nil {
        limit = &defaultInt64
    }
    if start == nil {
        start = &defaultString
    }
    if filter == nil {
        filter = &defaultString
    }
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
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
    protectionGroupS3AssetId string)(
    *models.ReadProtectionGroupS3AssetResponse, *apiutils.APIError) {

    pathURL := "/datasources/protection-groups/s3-assets/{protection_group_s3_asset_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_s3_asset_id": protectionGroupS3AssetId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    
    header := "application/api.clumio.protection-groups-s3-assets=v1+json"
    result := &models.ReadProtectionGroupS3AssetResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: p.config,
        RequestUrl: queryBuilder,
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
    defaultString := "" 
    
    if bucketName == nil {
        bucketName = &defaultString
    }
    if bucketId == nil {
        bucketId = &defaultString
    }
    if interval == nil {
        interval = &defaultString
    }
    
    queryParams := map[string]string{
        "bucket_name": *bucketName,
        "bucket_id": *bucketId,
        "begin_timestamp": beginTimestamp,
        "end_timestamp": endTimestamp,
        "interval": *interval,
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
