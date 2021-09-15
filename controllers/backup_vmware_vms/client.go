// Copyright (c) 2021 Clumio All Rights Reserved

package backupvmwarevms

import (
     "github.com/clumio-code/clumiogosdk/api_utils"
     "github.com/clumio-code/clumiogosdk/config"
     "github.com/clumio-code/clumiogosdk/models"
)

// BackupVmwareVmsV1Client represents a custom type interface
type BackupVmwareVmsV1Client interface {
    //  Returns a list of VMware virtual machines (VMs) that have been backed up by Clumio. VM backups can be restored through the [POST /restores/vmware/vms](#operation/restore-vmware-vm) endpoint.
    ListBackupVmwareVms(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListVMBackupsResponse,  *apiutils.APIError)
    
    //  Performs an on-demand backup for the specified VM. The VM must be protected with a policy that includes a service level agreement (SLA) configured for on-demand backups.
    CreateBackupVmwareVm(
        body models.CreateBackupVmwareVmV1Request)(
        interface{},  *apiutils.APIError)
    
    //  Returns a representation of the specified VM backup.
    ReadBackupVmwareVm(
        backupId int64)(
        *models.ReadVMBackupResponse,  *apiutils.APIError)
    
}

// NewBackupVmwareVmsV1 returns BackupVmwareVmsV1Client
func NewBackupVmwareVmsV1(config config.Config) BackupVmwareVmsV1Client{
    client := new(BackupVmwareVmsV1)
    client.config = config
    return client
}
