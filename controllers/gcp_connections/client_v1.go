// Copyright (c) 2023 Clumio All Rights Reserved

package gcpconnections

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpConnectionsV1Client represents a custom type interface
type GcpConnectionsV1Client interface {
    // ListGcpConnections Lists GCP Connections for a particular org
    ListGcpConnections(
        limit *int64, 
        start *string)(
        *models.ListGCPConnectionsResponse,  *apiutils.APIError)
    
    // CreateGcpConnection Create a new GCP Connection. This API should only be invoked by the Clumio Terraform provider and should not be invoked manually.
    CreateGcpConnection(
        body *models.CreateGcpConnectionV1Request)(
        *models.CreateGCPConnectionResponse,  *apiutils.APIError)
    
    // PostProcessGcpConnection Performs post-processing after GCP Connection Create, Update or Delete. This API should only be invoked by the Clumio Terraform provider and should not be invoked manually.
    PostProcessGcpConnection(
        body *models.PostProcessGcpConnectionV1Request)(
        interface{},  *apiutils.APIError)
    
    // ReadGcpConnection Reads a GCP Connection from the given project id
    ReadGcpConnection(
        projectId string)(
        *models.ReadGCPConnectionResponse,  *apiutils.APIError)
    
    // DeleteGcpConnection Deletes a GCP Connection
    DeleteGcpConnection(
        projectId string)(
        interface{},  *apiutils.APIError)
    
    // UpdateGcpConnection Updates a GCP Connection having the given project id. This API should only be invoked by the Clumio Terraform provider and should not be invoked manually.
    UpdateGcpConnection(
        projectId string, 
        body *models.UpdateGcpConnectionV1Request)(
        *models.UpdateGCPConnectionResponse,  *apiutils.APIError)
    
}

// NewGcpConnectionsV1 returns GcpConnectionsV1Client
func NewGcpConnectionsV1(config config.Config) GcpConnectionsV1Client {
    client := new(GcpConnectionsV1)
    client.config = config
    return client
}
