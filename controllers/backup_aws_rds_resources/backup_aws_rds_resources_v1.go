// Copyright (c) 2023 Clumio All Rights Reserved

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
