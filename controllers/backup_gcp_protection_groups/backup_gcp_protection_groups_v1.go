// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupgcpprotectiongroups contains methods related to BackupGcpProtectionGroups
package backupgcpprotectiongroups

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupGcpProtectionGroupsV1 represents a custom type struct
type BackupGcpProtectionGroupsV1 struct {
    config config.Config
}

// ListBackupGcpProtectionGroups Retrieves a list of protection group backups.
func (b *BackupGcpProtectionGroupsV1) ListBackupGcpProtectionGroups(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string)(
    *models.ListGCSProtectionGroupBackupsResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/gcp/protection-groups"

    
    header := "application/api.clumio.backup-gcp-protection-groups=v1+json"
    result := &models.ListGCSProtectionGroupBackupsResponse{}
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


// ReadBackupGcpProtectionGroup Returns a representation of the specified protection group backup.
func (b *BackupGcpProtectionGroupsV1) ReadBackupGcpProtectionGroup(
    backupId string)(
    *models.ReadGCSProtectionGroupBackupResponse, *apiutils.APIError) {

    pathURL := "/backups/gcp/protection-groups/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-gcp-protection-groups=v1+json"
    result := &models.ReadGCSProtectionGroupBackupResponse{}

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
