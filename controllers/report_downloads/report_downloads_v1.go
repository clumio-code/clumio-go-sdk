// Copyright (c) 2021 Clumio All Rights Reserved

// Package reportdownloads contains methods related to ReportDownloads
package reportdownloads

import (
    "encoding/json"
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// ReportDownloadsV1 represents a custom type struct
type ReportDownloadsV1 struct {
    config config.Config
}

//  ListReportDownloads List unexpired report downloads.
func (r *ReportDownloadsV1) ListReportDownloads(
    body models.ListReportDownloadsV1Request)(
    *models.ListReportDownloadsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := r.config.BaseUrl + "/reports/downloads"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/report-downloads=v1+json"
    var result *models.ListReportDownloadsResponse
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetAuthToken(r.config.Token).
        SetBody(payload).
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


//  CreateReportDownload Create a new Report download.
func (r *ReportDownloadsV1) CreateReportDownload(
    body models.CreateReportDownloadV1Request)(
    *models.CreateReportDownloadResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := r.config.BaseUrl + "/reports/downloads"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/report-downloads=v1+json"
    var result *models.CreateReportDownloadResponse
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetAuthToken(r.config.Token).
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
