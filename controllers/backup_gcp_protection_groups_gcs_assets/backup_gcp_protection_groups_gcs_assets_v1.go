// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupgcpprotectiongroupsgcsassets contains methods related to BackupGcpProtectionGroupsGcsAssets
package backupgcpprotectiongroupsgcsassets

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupGcpProtectionGroupsGcsAssetsV1 represents a custom type struct
type BackupGcpProtectionGroupsGcsAssetsV1 struct {
    config config.Config
}

// ListBackupGcpProtectionGroupGcsAssets Retrieves a list of asset backups.
func (b *BackupGcpProtectionGroupsGcsAssetsV1) ListBackupGcpProtectionGroupGcsAssets(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string)(
    *models.ListGCSProtectionGroupAssetBackupsResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/gcp/protection-groups/assets"

    
    header := "application/api.clumio.backup-gcp-protection-groups-gcs-assets=v1+json"
    result := &models.ListGCSProtectionGroupAssetBackupsResponse{}
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


// ReadBackupGcpProtectionGroupGcsAsset Returns a representation of the specified asset backup.
func (b *BackupGcpProtectionGroupsGcsAssetsV1) ReadBackupGcpProtectionGroupGcsAsset(
    backupId string)(
    *models.ReadGCSProtectionGroupAssetBackupResponse, *apiutils.APIError) {

    pathURL := "/backups/gcp/protection-groups/assets/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-gcp-protection-groups-gcs-assets=v1+json"
    result := &models.ReadGCSProtectionGroupAssetBackupResponse{}

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
