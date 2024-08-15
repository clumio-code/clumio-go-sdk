// Copyright (c) 2023 Clumio All Rights Reserved

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
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListReportDownloadsResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/reports/downloads"

    
    header := "application/api.clumio.report-downloads=v1+json"
    result := &models.ListReportDownloadsResponse{}
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
        Config: r.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
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
    result := &models.CreateReportDownloadResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}
