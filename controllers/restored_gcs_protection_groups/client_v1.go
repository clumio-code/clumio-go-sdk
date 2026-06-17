// Copyright (c) 2023 Clumio All Rights Reserved

package restoredgcsprotectiongroups

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredGcsProtectionGroupsV1Client represents a custom type interface
type RestoredGcsProtectionGroupsV1Client interface {
    // RestoreGcsProtectionGroup Restores the specified protection group backup to the specified target destination.
    RestoreGcsProtectionGroup(
        embed *string, 
        body models.RestoreGcsProtectionGroupV1Request)(
        *models.RestoreGCSProtectionGroupResponse,  *apiutils.APIError)
    
    // RestoreGcsProtectionGroupObjects Restores the specified list of objects to the specified target destination.
    RestoreGcsProtectionGroupObjects(
        gcsProtectionGroupId string, 
        embed *string, 
        body models.RestoreGcsProtectionGroupObjectsV1Request)(
        *models.RestoreGCSObjectsResponse,  *apiutils.APIError)
    
    // PreviewGcsProtectionGroup Preview a protection group restore.
    PreviewGcsProtectionGroup(
        gcsProtectionGroupId string, 
        body models.PreviewGcsProtectionGroupV1Request)(
        *models.PreviewGCSProtectionGroupAsyncResponse,  *apiutils.APIError)
    
    // PreviewDetailsGcsProtectionGroup Details for protection group restore preview
    PreviewDetailsGcsProtectionGroup(
        gcsProtectionGroupId string, 
        previewId string)(
        *models.PreviewDetailsGCSProtectionGroupResponse,  *apiutils.APIError)
    
}

// NewRestoredGcsProtectionGroupsV1 returns RestoredGcsProtectionGroupsV1Client
func NewRestoredGcsProtectionGroupsV1(config config.Config) RestoredGcsProtectionGroupsV1Client {
    client := new(RestoredGcsProtectionGroupsV1)
    client.config = config
    return client
}
