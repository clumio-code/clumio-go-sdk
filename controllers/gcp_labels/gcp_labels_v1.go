// Copyright (c) 2023 Clumio All Rights Reserved

// Package gcplabels contains methods related to GcpLabels
package gcplabels

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpLabelsV1 represents a custom type struct
type GcpLabelsV1 struct {
    config config.Config
}

// ListGcpLabelKeys Returns a list of GCP label keys.
func (g *GcpLabelsV1) ListGcpLabelKeys(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListGCPLabelKeysResponse, *apiutils.APIError) {

    queryBuilder := g.config.BaseUrl + "/datasources/gcp/label-keys"

    
    header := "application/api.clumio.gcp-labels=v1+json"
    result := &models.ListGCPLabelKeysResponse{}
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
        Config: g.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ListGcpLabelValues Returns a list of GCP label values for the specified label key.
func (g *GcpLabelsV1) ListGcpLabelValues(
    labelKeyId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListGCPLabelValuesResponse, *apiutils.APIError) {

    pathURL := "/datasources/gcp/label-keys/{label_key_id}/label-values"
    //process optional template parameters
    pathParams := map[string]string{
        "label_key_id": labelKeyId,
    }
    queryBuilder := g.config.BaseUrl + pathURL

    
    header := "application/api.clumio.gcp-labels=v1+json"
    result := &models.ListGCPLabelValuesResponse{}
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
        Config: g.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
