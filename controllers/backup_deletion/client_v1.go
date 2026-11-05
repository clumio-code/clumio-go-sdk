// Copyright (c) 2023 Clumio All Rights Reserved

package backupdeletion

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupDeletionV1Client represents a custom type interface
type BackupDeletionV1Client interface {
    // ListBackupDeletions Lists all Backup Deletion Requests
    ListBackupDeletions(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListBackupPurgeResponse,  *apiutils.APIError)
    
    // CreateBackupDeletion Creates a backup deletion request in draft state. Add the backups to delete with
    //  POST /backup-deletions/{backup_deletion_id}/items. An organization may hold only
    //  one draft request at a time.
    CreateBackupDeletion(
        body models.CreateBackupDeletionV1Request)(
        *models.CreateBackupPurgeResponse,  *apiutils.APIError)
    
    // ReadBackupDeletion Returns a representation of the specified backup deletion request
    ReadBackupDeletion(
        backupDeletionId string)(
        *models.ReadBackupPurgeResponse,  *apiutils.APIError)
    
    // DeleteBackupDeletion Deletes a draft backup deletion request
    DeleteBackupDeletion(
        backupDeletionId string)(
        *models.DeleteBackupPurgeResponse,  *apiutils.APIError)
    
    // UpdateBackupDeletion Edits a draft backup deletion request, or submits it for signature.
    //  
    //  Send deletion_status "submitted" to submit the request for signature, or
    //  "cancelled" to abandon it; both are allowed only from draft. Omit deletion_status
    //  to edit the draft's name, requestor or approver.
    UpdateBackupDeletion(
        backupDeletionId string, 
        body models.UpdateBackupDeletionV1Request)(
        *models.UpdateBackupPurgeResponse,  *apiutils.APIError)
    
    // ListBackupDeletionItems List items in backup deletion request
    ListBackupDeletionItems(
        backupDeletionId string, 
        limit *int64, 
        start *string)(
        *models.ListBackupPurgeItemsResponse,  *apiutils.APIError)
    
    // CreateBackupDeletionItems Create a backup deletion item and add to the backup deletion request
    CreateBackupDeletionItems(
        backupDeletionId string, 
        body models.CreateBackupDeletionItemsV1Request)(
        *models.CreateBackupPurgeItemsResponse,  *apiutils.APIError)
    
    // CreateBackupDeletionItemsPreview Create a backup deletion items preview and returns the taskId for each item
    CreateBackupDeletionItemsPreview(
        backupDeletionId string, 
        body models.CreateBackupDeletionItemsPreviewV1Request)(
        *models.CreateBackupPurgeItemsPreviewResponse,  *apiutils.APIError)
    
    // DeleteBackupDeletionItem Remove a backup deletion item from backup deletion request
    DeleteBackupDeletionItem(
        backupDeletionId string, 
        backupDeletionItemId string)(
        *models.RemoveBackupPurgeItemResponse,  *apiutils.APIError)
    
    // ListBackupDeletionItemsPreview Returns a representation of the specified backup deletion item preview
    ListBackupDeletionItemsPreview(
        backupDeletionId string, 
        backupDeletionItemId string, 
        limit *int64, 
        start *string)(
        *models.ListBackupPurgeItemsPreviewResponse,  *apiutils.APIError)
    
}

// NewBackupDeletionV1 returns BackupDeletionV1Client
func NewBackupDeletionV1(config config.Config) BackupDeletionV1Client {
    client := new(BackupDeletionV1)
    client.config = config
    return client
}
