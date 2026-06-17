// Copyright (c) 2023 Clumio All Rights Reserved

package gcpprojects

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpProjectsV1Client represents a custom type interface
type GcpProjectsV1Client interface {
    // ListGcpProjects Returns a list of GCP projects. By default only active projects are
    //  returned; use the is_deleted filter to return deleted projects.
    ListGcpProjects(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListGCPProjectsResponse,  *apiutils.APIError)
    
}

// NewGcpProjectsV1 returns GcpProjectsV1Client
func NewGcpProjectsV1(config config.Config) GcpProjectsV1Client {
    client := new(GcpProjectsV1)
    client.config = config
    return client
}
