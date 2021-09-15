// Copyright (c) 2021 Clumio All Rights Reserved

// Package audittrails contains methods related to AuditTrails
package audittrails

import (
    "fmt"

    "github.com/clumio-code/clumiogosdk/api_utils"
    "github.com/clumio-code/clumiogosdk/config"
    "github.com/clumio-code/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// AuditTrailsV1 represents a custom type struct
type AuditTrailsV1 struct {
    config config.Config
}

//  ListAuditTrails Returns a list of audit trails.
func (a *AuditTrailsV1) ListAuditTrails(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListAuditTrailsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := a.config.BaseUrl + "/audit-trails"

    
    header := "application/audit-trails=v1+json"
    var result *models.ListAuditTrailsResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    if filter == nil{
        filter = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetAuthToken(a.config.Token).
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
