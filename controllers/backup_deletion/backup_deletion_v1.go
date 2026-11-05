// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupdeletion contains methods related to BackupDeletion
package backupdeletion

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupDeletionV1 represents a custom type struct
type BackupDeletionV1 struct {
    config config.Config
}

// ListBackupDeletions Lists all Backup Deletion Requests
func (b *BackupDeletionV1) ListBackupDeletions(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListBackupPurgeResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backup-deletions"

    
    header := "application/api.clumio.backup-deletion=v1+json"
    result := &models.ListBackupPurgeResponse{}
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
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreateBackupDeletion Creates a backup deletion request in draft state. Add the backups to delete with
//  POST /backup-deletions/{backup_deletion_id}/items. An organization may hold only
//  one draft request at a time.
func (b *BackupDeletionV1) CreateBackupDeletion(
    body models.CreateBackupDeletionV1Request)(
    *models.CreateBackupPurgeResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backup-deletions"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.backup-deletion=v1+json"
    result := &models.CreateBackupPurgeResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result201: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadBackupDeletion Returns a representation of the specified backup deletion request
func (b *BackupDeletionV1) ReadBackupDeletion(
    backupDeletionId string)(
    *models.ReadBackupPurgeResponse, *apiutils.APIError) {

    pathURL := "/backup-deletions/{backup_deletion_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_deletion_id": backupDeletionId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-deletion=v1+json"
    result := &models.ReadBackupPurgeResponse{}

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


// DeleteBackupDeletion Deletes a draft backup deletion request
func (b *BackupDeletionV1) DeleteBackupDeletion(
    backupDeletionId string)(
    *models.DeleteBackupPurgeResponse, *apiutils.APIError) {

    pathURL := "/backup-deletions/{backup_deletion_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_deletion_id": backupDeletionId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-deletion=v1+json"
    result := &models.DeleteBackupPurgeResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Delete,
    })

    return result, apiErr
}


// UpdateBackupDeletion Edits a draft backup deletion request, or submits it for signature.
//  
//  Send deletion_status "submitted" to submit the request for signature, or
//  "cancelled" to abandon it; both are allowed only from draft. Omit deletion_status
//  to edit the draft's name, requestor or approver.
func (b *BackupDeletionV1) UpdateBackupDeletion(
    backupDeletionId string, 
    body models.UpdateBackupDeletionV1Request)(
    *models.UpdateBackupPurgeResponse, *apiutils.APIError) {

    pathURL := "/backup-deletions/{backup_deletion_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_deletion_id": backupDeletionId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.backup-deletion=v1+json"
    result := &models.UpdateBackupPurgeResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Patch,
    })

    return result, apiErr
}


// ListBackupDeletionItems List items in backup deletion request
func (b *BackupDeletionV1) ListBackupDeletionItems(
    backupDeletionId string, 
    limit *int64, 
    start *string)(
    *models.ListBackupPurgeItemsResponse, *apiutils.APIError) {

    pathURL := "/backup-deletions/{backup_deletion_id}/items"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_deletion_id": backupDeletionId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-deletion=v1+json"
    result := &models.ListBackupPurgeItemsResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
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


// CreateBackupDeletionItems Create a backup deletion item and add to the backup deletion request
func (b *BackupDeletionV1) CreateBackupDeletionItems(
    backupDeletionId string, 
    body models.CreateBackupDeletionItemsV1Request)(
    *models.CreateBackupPurgeItemsResponse, *apiutils.APIError) {

    pathURL := "/backup-deletions/{backup_deletion_id}/items"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_deletion_id": backupDeletionId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.backup-deletion=v1+json"
    result := &models.CreateBackupPurgeItemsResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// CreateBackupDeletionItemsPreview Create a backup deletion items preview and returns the taskId for each item
func (b *BackupDeletionV1) CreateBackupDeletionItemsPreview(
    backupDeletionId string, 
    body models.CreateBackupDeletionItemsPreviewV1Request)(
    *models.CreateBackupPurgeItemsPreviewResponse, *apiutils.APIError) {

    pathURL := "/backup-deletions/{backup_deletion_id}/items/preview"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_deletion_id": backupDeletionId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.backup-deletion=v1+json"
    result := &models.CreateBackupPurgeItemsPreviewResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// DeleteBackupDeletionItem Remove a backup deletion item from backup deletion request
func (b *BackupDeletionV1) DeleteBackupDeletionItem(
    backupDeletionId string, 
    backupDeletionItemId string)(
    *models.RemoveBackupPurgeItemResponse, *apiutils.APIError) {

    pathURL := "/backup-deletions/{backup_deletion_id}/items/{backup_deletion_item_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_deletion_id": backupDeletionId,
        "backup_deletion_item_id": backupDeletionItemId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-deletion=v1+json"
    result := &models.RemoveBackupPurgeItemResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Delete,
    })

    return result, apiErr
}


// ListBackupDeletionItemsPreview Returns a representation of the specified backup deletion item preview
func (b *BackupDeletionV1) ListBackupDeletionItemsPreview(
    backupDeletionId string, 
    backupDeletionItemId string, 
    limit *int64, 
    start *string)(
    *models.ListBackupPurgeItemsPreviewResponse, *apiutils.APIError) {

    pathURL := "/backup-deletions/{backup_deletion_id}/items/{backup_deletion_item_id}/preview"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_deletion_id": backupDeletionId,
        "backup_deletion_item_id": backupDeletionItemId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-deletion=v1+json"
    result := &models.ListBackupPurgeItemsPreviewResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
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
