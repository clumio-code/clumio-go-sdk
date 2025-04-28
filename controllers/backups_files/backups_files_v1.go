// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupsfiles contains methods related to BackupsFiles
package backupsfiles

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupsFilesV1 represents a custom type struct
type BackupsFilesV1 struct {
    config config.Config
}

// ListFiles Retrieve the list of files whose name matches a given regex pattern.
func (b *BackupsFilesV1) ListFiles(
    limit *int64, 
    start *string, 
    filter string)(
    *models.FileSearchResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/files/search"

    
    header := "application/api.clumio.backups-files=v1+json"
    result := &models.FileSearchResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ListFileVersions Retrieve the list of versions of the file.
func (b *BackupsFilesV1) ListFileVersions(
    searchResultId string, 
    limit *int64, 
    start *string)(
    *models.FileListResponse, *apiutils.APIError) {

    pathURL := "/backups/files/search/{search_result_id}/versions"
    //process optional template parameters
    pathParams := map[string]string{
        "search_result_id": searchResultId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backups-files=v1+json"
    result := &models.FileListResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
