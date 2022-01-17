// Copyright (c) 2021 Clumio All Rights Reserved

// Package backupsfiles contains methods related to BackupsFiles
package backupsfiles

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// BackupsFilesV1 represents a custom type struct
type BackupsFilesV1 struct {
    config config.Config
}

//  ListFiles Retrieve the list of files whose name matches a given regex pattern.
func (b *BackupsFilesV1) ListFiles(
    limit *int64, 
    start *string, 
    filter string)(
    *models.FileSearchResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := b.config.BaseUrl + "/backups/files/search"

    
    header := "application/backups-files=v1+json"
    var result *models.FileSearchResponse
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
        SetAuthToken(b.config.Token).
        SetResult(&result).
        Get(queryBuilder)

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


//  ListFileVersions Retrieve the list of versions of the file.
func (b *BackupsFilesV1) ListFileVersions(
    searchResultId string, 
    limit *int64, 
    start *string)(
    *models.FileListResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/backups/files/search/{search_result_id}/versions"
    //process optional template parameters
    pathParams := map[string]string{
        "search_result_id": searchResultId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/backups-files=v1+json"
    var result *models.FileListResponse
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
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(b.config.Token).
        SetResult(&result).
        Get(queryBuilder)

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
