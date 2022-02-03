// Copyright (c) 2021 Clumio All Rights Reserved

// Package backupvmwarevms contains methods related to BackupVmwareVms
package backupvmwarevms

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupVmwareVmsV1 represents a custom type struct
type BackupVmwareVmsV1 struct {
    config config.Config
}

// ListBackupVmwareVms Returns a list of VMware virtual machines (VMs) that have been backed up by Clumio. VM backups can be restored through the [POST /restores/vmware/vms](#operation/restore-vmware-vm) endpoint.
func (b *BackupVmwareVmsV1) ListBackupVmwareVms(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListVMBackupsResponse, *apiutils.APIError){

    queryBuilder := b.config.BaseUrl + "/backups/vmware/vms"

    
    header := "application/backup-vmware-vms=v1+json"
    var result *models.ListVMBackupsResponse
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

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreateBackupVmwareVm Performs an on-demand backup for the specified VM. The VM must be protected with a policy that includes a service level agreement (SLA) configured for on-demand backups.
func (b *BackupVmwareVmsV1) CreateBackupVmwareVm(
    body models.CreateBackupVmwareVmV1Request)(
    interface{}, *apiutils.APIError){

    queryBuilder := b.config.BaseUrl + "/backups/vmware/vms"

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

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadBackupVmwareVm Returns a representation of the specified VM backup.
func (b *BackupVmwareVmsV1) ReadBackupVmwareVm(
    backupId int64)(
    *models.ReadVMBackupResponse, *apiutils.APIError){

    pathURL := "/backups/vmware/vms/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": fmt.Sprintf("%v", backupId),
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/backup-vmware-vms=v1+json"
    var result *models.ReadVMBackupResponse

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
