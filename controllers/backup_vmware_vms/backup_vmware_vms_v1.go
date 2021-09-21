// Copyright (c) 2021 Clumio All Rights Reserved

// Package backupvmwarevms contains methods related to BackupVmwareVms
package backupvmwarevms

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// BackupVmwareVmsV1 represents a custom type struct
type BackupVmwareVmsV1 struct {
    config config.Config
}

//  ListBackupVmwareVms Returns a list of VMware virtual machines (VMs) that have been backed up by Clumio. VM backups can be restored through the [POST /restores/vmware/vms](#operation/restore-vmware-vm) endpoint.
func (b *BackupVmwareVmsV1) ListBackupVmwareVms(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListVMBackupsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := b.config.BaseUrl + "/backups/vmware/vms"

    
    header := "application/backup-vmware-vms=v1+json"
    var result *models.ListVMBackupsResponse
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


//  CreateBackupVmwareVm Performs an on-demand backup for the specified VM. The VM must be protected with a policy that includes a service level agreement (SLA) configured for on-demand backups.
func (b *BackupVmwareVmsV1) CreateBackupVmwareVm(
    body models.CreateBackupVmwareVmV1Request)(
    interface{}, *apiutils.APIError){

    var err error = nil
    _queryBuilder := b.config.BaseUrl + "/backups/vmware/vms"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/backup-vmware-vms=v1+json"
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


//  ReadBackupVmwareVm Returns a representation of the specified VM backup.
func (b *BackupVmwareVmsV1) ReadBackupVmwareVm(
    backupId int64)(
    *models.ReadVMBackupResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/backups/vmware/vms/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backupId": fmt.Sprintf("%v", backupId),
    }
    _queryBuilder := b.config.BaseUrl + _pathURL

    
    header := "application/backup-vmware-vms=v1+json"
    var result *models.ReadVMBackupResponse
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
