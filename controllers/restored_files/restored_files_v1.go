// Copyright (c) 2021 Clumio All Rights Reserved

// Package restoredfiles contains methods related to RestoredFiles
package restoredfiles

import (
    "encoding/json"
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// RestoredFilesV1 represents a custom type struct
type RestoredFilesV1 struct {
    config config.Config
}

//  ListRestoredFiles Gets the list of active restored files for an asset.
func (r *RestoredFilesV1) ListRestoredFiles(
    limit *int64, 
    start *string, 
    filter string)(
    *models.RestoredFilesResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := r.config.BaseUrl + "/restores/files"

    
    header := "application/restored-files=v1+json"
    var result *models.RestoredFilesResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": filter,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetAuthToken(r.config.Token).
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


//  RestoreFiles Restores one or more files from the specified backup.
func (r *RestoredFilesV1) RestoreFiles(
    body models.RestoreFilesV1Request)(
    *models.RestoreFileResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := r.config.BaseUrl + "/restores/files"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/restored-files=v1+json"
    var result *models.RestoreFileResponse
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
