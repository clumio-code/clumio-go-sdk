// Copyright (c) 2021 Clumio All Rights Reserved

// Package restoredfiles contains methods related to RestoredFiles
package restoredfiles

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredFilesV1 represents a custom type struct
type RestoredFilesV1 struct {
    config config.Config
}

// ListRestoredFiles Gets the list of active restored files for an asset.
func (r *RestoredFilesV1) ListRestoredFiles(
    limit *int64, 
    start *string, 
    filter string)(
    *models.RestoredFilesResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/files"

    
    header := "application/api.clumio.restored-files=v1+json"
    var result *models.RestoredFilesResponse
    defaultInt64 := int64(0)
    defaultString := "" 
    
    if limit == nil {
        limit = &defaultInt64
    }
    if start == nil {
        start = &defaultString
    }
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": filter,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// RestoreFiles Restores one or more files from the specified backup.
func (r *RestoredFilesV1) RestoreFiles(
    body models.RestoreFilesV1Request)(
    *models.RestoreFileResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/files"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-files=v1+json"
    var result *models.RestoreFileResponse

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
