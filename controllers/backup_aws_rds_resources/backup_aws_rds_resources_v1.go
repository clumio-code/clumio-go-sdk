// Copyright (c) 2021 Clumio All Rights Reserved

// Package backupawsrdsresources contains methods related to BackupAwsRdsResources
package backupawsrdsresources

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsRdsResourcesV1 represents a custom type struct
type BackupAwsRdsResourcesV1 struct {
    config config.Config
}

// ListBackupAwsRdsResources Retrieves a list of RDS database backups.
func (b *BackupAwsRdsResourcesV1) ListBackupAwsRdsResources(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string)(
    *models.ListRdsDatabaseBackupsResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/rds-resources"

    
    header := "application/api.clumio.backup-aws-rds-resources=v1+json"
    result := &models.ListRdsDatabaseBackupsResponse{}
    defaultInt64 := int64(0)
    defaultString := "" 
    
    if limit == nil {
        limit = &defaultInt64
    }
    if start == nil {
        start = &defaultString
    }
    if sort == nil {
        sort = &defaultString
    }
    if filter == nil {
        filter = &defaultString
    }
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "sort": *sort,
        "filter": *filter,
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


// ReadBackupAwsRdsResource Returns a representation of the specified RDS database backup.
func (b *BackupAwsRdsResourcesV1) ReadBackupAwsRdsResource(
    backupId string)(
    *models.ReadRdsDatabaseBackupResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/rds-resources/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-aws-rds-resources=v1+json"
    result := &models.ReadRdsDatabaseBackupResponse{}

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


// ListAwsRdsResourcesOptionGroups Retrieves a list of RDS option groups which are superset of persistent and permanent
//  options present in the backup snapshot for a given environment.
func (b *BackupAwsRdsResourcesV1) ListAwsRdsResourcesOptionGroups(
    backupId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListRdsOptionGroupsResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/rds-resources/{backup_id}/option-groups"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-aws-rds-resources=v1+json"
    result := &models.ListRdsOptionGroupsResponse{}
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
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
