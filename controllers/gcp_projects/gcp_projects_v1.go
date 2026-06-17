// Copyright (c) 2023 Clumio All Rights Reserved

// Package gcpprojects contains methods related to GcpProjects
package gcpprojects

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpProjectsV1 represents a custom type struct
type GcpProjectsV1 struct {
    config config.Config
}

// ListGcpProjects Returns a list of GCP projects. By default only active projects are
//  returned; use the is_deleted filter to return deleted projects.
func (g *GcpProjectsV1) ListGcpProjects(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListGCPProjectsResponse, *apiutils.APIError) {

    queryBuilder := g.config.BaseUrl + "/datasources/gcp/projects"

    
    header := "application/api.clumio.gcp-projects=v1+json"
    result := &models.ListGCPProjectsResponse{}
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
