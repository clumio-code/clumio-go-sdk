// Copyright (c) 2023 Clumio All Rights Reserved

// Package audittrails contains methods related to AuditTrails
package audittrails

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AuditTrailsV1 represents a custom type struct
type AuditTrailsV1 struct {
    config config.Config
}

// ListAuditTrails Returns a list of audit trails.
func (a *AuditTrailsV1) ListAuditTrails(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListAuditTrailsResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/audit-trails"

    
    header := "application/api.clumio.audit-trails=v1+json"
    result := &models.ListAuditTrailsResponse{}
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
        Config: a.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
