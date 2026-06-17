// Copyright (c) 2023 Clumio All Rights Reserved

// Package gcpreport contains methods related to GcpReport
package gcpreport

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpReportV1 represents a custom type struct
type GcpReportV1 struct {
    config config.Config
}

// GcsAssetErrorReport Returns task ID
func (g *GcpReportV1) GcsAssetErrorReport(
    embed *string, 
    body models.GcsAssetErrorReportV1Request)(
    *models.GcsAssetErrorReportV1Response, *apiutils.APIError) {

    queryBuilder := g.config.BaseUrl + "/datasources/gcp/report/gcs/gcs-asset"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.gcp-report=v1+json"
    result := &models.GcsAssetErrorReportV1Response{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// GcsProtectionGroupErrorReport Returns task ID
func (g *GcpReportV1) GcsProtectionGroupErrorReport(
    embed *string, 
    body models.GcsProtectionGroupErrorReportV1Request)(
    *models.GcsProtectionGroupErrorReportV1Response, *apiutils.APIError) {

    queryBuilder := g.config.BaseUrl + "/datasources/gcp/report/gcs/protection-group"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.gcp-report=v1+json"
    result := &models.GcsProtectionGroupErrorReportV1Response{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}
