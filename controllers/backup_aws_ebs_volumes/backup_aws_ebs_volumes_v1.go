// Copyright (c) 2021 Clumio All Rights Reserved

// Package backupawsebsvolumes contains methods related to BackupAwsEbsVolumes
package backupawsebsvolumes

import (
    "encoding/json"
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// BackupAwsEbsVolumesV1 represents a custom type struct
type BackupAwsEbsVolumesV1 struct {
    config config.Config
}

//  ListBackupAwsEbsVolumes Returns a list of EBS volumes that have been backed up by Clumio. EBS volume backups can be restored through the [POST /restores/aws/ebs-volumes](#operation/restore-aws-ebs-volume) endpoint.
func (b *BackupAwsEbsVolumesV1) ListBackupAwsEbsVolumes(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListEBSBackupsResponseV1, *apiutils.APIError){

    var err error = nil
    _queryBuilder := b.config.BaseUrl + "/backups/aws/ebs-volumes"

    
    header := "application/backup-aws-ebs-volumes=v1+json"
    var result *models.ListEBSBackupsResponseV1
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
        SetAuthToken(b.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

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


//  CreateBackupAwsEbsVolume Performs an on-demand backup for the specified EBS volume. The EBS volume must be protected with a policy that includes a service level agreement (SLA) configured for on-demand backups.
func (b *BackupAwsEbsVolumesV1) CreateBackupAwsEbsVolume(
    body models.CreateBackupAwsEbsVolumeV1Request)(
    interface{}, *apiutils.APIError){

    var err error = nil
    _queryBuilder := b.config.BaseUrl + "/backups/aws/ebs-volumes"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/backup-aws-ebs-volumes=v1+json"
    var result interface{}
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetAuthToken(b.config.Token).
        SetBody(payload).
        SetResult(&result).
        Post(_queryBuilder)

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


//  ReadBackupAwsEbsVolume Returns a representation of the specified EBS volume backup.
func (b *BackupAwsEbsVolumesV1) ReadBackupAwsEbsVolume(
    backupId string)(
    *models.ReadEBSBackupResponseV1, *apiutils.APIError){

    var err error = nil
    _pathURL := "/backups/aws/ebs-volumes/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backupId": backupId,
    }
    _queryBuilder := b.config.BaseUrl + _pathURL

    
    header := "application/backup-aws-ebs-volumes=v1+json"
    var result *models.ReadEBSBackupResponseV1
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(b.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

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
