// Copyright (c) 2021 Clumio All Rights Reserved

package reportdownloads

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// ReportDownloadsV1Client represents a custom type interface
type ReportDownloadsV1Client interface {
    //  List unexpired report downloads.
    ListReportDownloads(
        body models.ListReportDownloadsV1Request)(
        *models.ListReportDownloadsResponse,  *apiutils.APIError)
    
    //  Create a new Report download.
    CreateReportDownload(
        body models.CreateReportDownloadV1Request)(
        *models.CreateReportDownloadResponse,  *apiutils.APIError)
    
}

// NewReportDownloadsV1 returns ReportDownloadsV1Client
func NewReportDownloadsV1(config config.Config) ReportDownloadsV1Client{
    client := new(ReportDownloadsV1)
    client.config = config
    return client
}
