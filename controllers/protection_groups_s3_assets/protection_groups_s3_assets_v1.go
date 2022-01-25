// Copyright (c) 2021 Clumio All Rights Reserved

// Package protectiongroupss3assets contains methods related to ProtectionGroupsS3Assets
package protectiongroupss3assets

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// ProtectionGroupsS3AssetsV1 represents a custom type struct
type ProtectionGroupsS3AssetsV1 struct {
    config config.Config
}

//  ListProtectionGroupS3Assets Returns a list of protection group S3 assets.
func (p *ProtectionGroupsS3AssetsV1) ListProtectionGroupS3Assets(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListProtectionGroupS3AssetsResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := p.config.BaseUrl + "/datasources/protection-groups/s3-assets"

    
    header := "application/protection-groups-s3-assets=v1+json"
    var result *models.ListProtectionGroupS3AssetsResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    if filter == nil{
        filter = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetAuthToken(p.config.Token).
        SetResult(&result).
        Get(queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}


//  ReadProtectionGroupS3Asset Returns a representation of the specified protection group S3 asset.
func (p *ProtectionGroupsS3AssetsV1) ReadProtectionGroupS3Asset(
    protectionGroupS3AssetId string)(
    *models.ReadProtectionGroupS3AssetResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/protection-groups/s3-assets/{protection_group_s3_asset_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "protection_group_s3_asset_id": protectionGroupS3AssetId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    
    header := "application/protection-groups-s3-assets=v1+json"
    var result *models.ReadProtectionGroupS3AssetResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(p.config.Token).
        SetResult(&result).
        Get(queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}
