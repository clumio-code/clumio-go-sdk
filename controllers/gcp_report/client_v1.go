// Copyright (c) 2023 Clumio All Rights Reserved

package gcpreport

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpReportV1Client represents a custom type interface
type GcpReportV1Client interface {
    // GcsAssetErrorReport Returns task ID
    GcsAssetErrorReport(
        embed *string, 
        body models.GcsAssetErrorReportV1Request)(
        *models.GcsAssetErrorReportV1Response,  *apiutils.APIError)
    
    // GcsProtectionGroupErrorReport Returns task ID
    GcsProtectionGroupErrorReport(
        embed *string, 
        body models.GcsProtectionGroupErrorReportV1Request)(
        *models.GcsProtectionGroupErrorReportV1Response,  *apiutils.APIError)
    
}

// NewGcpReportV1 returns GcpReportV1Client
func NewGcpReportV1(config config.Config) GcpReportV1Client {
    client := new(GcpReportV1)
    client.config = config
    return client
}
