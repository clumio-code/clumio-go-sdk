// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupawsneptune contains methods related to BackupAwsNeptune
package backupawsneptune

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsNeptuneV1 represents a custom type struct
type BackupAwsNeptuneV1 struct {
    config config.Config
}

// ListBackupAwsNeptune Retrieves a list of Neptune backups.
func (b *BackupAwsNeptuneV1) ListBackupAwsNeptune(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string)(
    *models.ListNeptuneBackupsResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/neptune"

    
    header := "application/api.clumio.backup-aws-neptune=v1+json"
    result := &models.ListNeptuneBackupsResponse{}
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


// CreateBackupAwsNeptune Performs an on-demand backup for the specified Neptune cluster.
func (b *BackupAwsNeptuneV1) CreateBackupAwsNeptune(
    embed *string, 
    body models.CreateBackupAwsNeptuneV1Request)(
    *models.OnDemandNeptuneBackupResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/neptune"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.backup-aws-neptune=v1+json"
    result := &models.OnDemandNeptuneBackupResponse{}
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


// ReadBackupAwsNeptune Returns a representation of the specified Neptune backup.
func (b *BackupAwsNeptuneV1) ReadBackupAwsNeptune(
    backupId string)(
    *models.ReadNeptuneBackupResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/neptune/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-aws-neptune=v1+json"
    result := &models.ReadNeptuneBackupResponse{}

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
