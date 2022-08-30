// Copyright (c) 2021 Clumio All Rights Reserved

// Package reportdownloads contains methods related to ReportDownloads
package reportdownloads

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// ReportDownloadsV1 represents a custom type struct
type ReportDownloadsV1 struct {
    config config.Config
}

// ListReportDownloads Returns a list of unexpired, generated reports.
func (r *ReportDownloadsV1) ListReportDownloads(
    body models.ListReportDownloadsV1Request)(
    *models.ListReportDownloadsResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/reports/downloads"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.report-downloads=v1+json"
    var result *models.ListReportDownloadsResponse

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreateReportDownload Generates a report of a specified type given certain general conditions such as time range and other type-specific filters.
func (r *ReportDownloadsV1) CreateReportDownload(
    body models.CreateReportDownloadV1Request)(
    *models.CreateReportDownloadResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/reports/downloads"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.report-downloads=v1+json"
    var result *models.CreateReportDownloadResponse

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}
