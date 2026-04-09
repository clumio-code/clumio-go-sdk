// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupprotectiongroups contains methods related to BackupProtectionGroups
package backupprotectiongroups

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupProtectionGroupsV1 represents a custom type struct
type BackupProtectionGroupsV1 struct {
    config config.Config
}

// ListBackupProtectionGroups Retrieves a list of protection group backups.
func (b *BackupProtectionGroupsV1) ListBackupProtectionGroups(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string)(
    *models.ListProtectionGroupBackupsResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/protection-groups"

    
    header := "application/api.clumio.backup-protection-groups=v1+json"
    result := &models.ListProtectionGroupBackupsResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
    }
    if sort != nil {
        queryParams["sort"] = *sort
    }
    if filter != nil {
        queryParams["filter"] = *filter
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ListBackupProtectionGroupS3Assets Retrieves a list of protection group S3 asset backups.
func (b *BackupProtectionGroupsV1) ListBackupProtectionGroupS3Assets(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string)(
    *models.ListProtectionGroupS3AssetBackupsResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/protection-groups/s3-assets"

    
    header := "application/api.clumio.backup-protection-groups=v1+json"
    result := &models.ListProtectionGroupS3AssetBackupsResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
    }
    if sort != nil {
        queryParams["sort"] = *sort
    }
    if filter != nil {
        queryParams["filter"] = *filter
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ExportProtectionGroupS3AssetMalwareReport Exports the specified malware report for a protection group S3 asset.
func (b *BackupProtectionGroupsV1) ExportProtectionGroupS3AssetMalwareReport(
    embed *string, 
    body models.ExportProtectionGroupS3AssetMalwareReportV1Request)(
    *models.ExportMalwareReportResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/protection-groups/s3-assets/malware-report"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.backup-protection-groups=v1+json"
    result := &models.ExportMalwareReportResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadBackupProtectionGroupS3Asset Returns a representation of the specified protection group S3 asset backup.
func (b *BackupProtectionGroupsV1) ReadBackupProtectionGroupS3Asset(
    backupId string)(
    *models.ReadProtectionGroupS3AssetBackupResponse, *apiutils.APIError) {

    pathURL := "/backups/protection-groups/s3-assets/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-protection-groups=v1+json"
    result := &models.ReadProtectionGroupS3AssetBackupResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadBackupProtectionGroup Returns a representation of the specified protection group backup.
func (b *BackupProtectionGroupsV1) ReadBackupProtectionGroup(
    backupId string)(
    *models.ReadProtectionGroupBackupResponse, *apiutils.APIError) {

    pathURL := "/backups/protection-groups/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-protection-groups=v1+json"
    result := &models.ReadProtectionGroupBackupResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
