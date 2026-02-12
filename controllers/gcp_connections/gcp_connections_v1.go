// Copyright (c) 2023 Clumio All Rights Reserved

// Package gcpconnections contains methods related to GcpConnections
package gcpconnections

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpConnectionsV1 represents a custom type struct
type GcpConnectionsV1 struct {
    config config.Config
}

// ListGcpConnections Lists GCP Connections for a particular org
func (g *GcpConnectionsV1) ListGcpConnections(
    limit *int64, 
    start *string)(
    *models.ListGCPConnectionsResponse, *apiutils.APIError) {

    queryBuilder := g.config.BaseUrl + "/connections/gcp"

    
    header := "application/api.clumio.gcp-connections=v1+json"
    result := &models.ListGCPConnectionsResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
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


// CreateGcpConnection Create a new GCP Connection. This API should only be invoked by the Clumio Terraform provider and should not be invoked manually.
func (g *GcpConnectionsV1) CreateGcpConnection(
    body *models.CreateGcpConnectionV1Request)(
    *models.CreateGCPConnectionResponse, *apiutils.APIError) {

    queryBuilder := g.config.BaseUrl + "/connections/gcp"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.gcp-connections=v1+json"
    result := &models.CreateGCPConnectionResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// PostProcessGcpConnection Performs post-processing after GCP Connection Create, Update or Delete. This API should only be invoked by the Clumio Terraform provider and should not be invoked manually.
func (g *GcpConnectionsV1) PostProcessGcpConnection(
    body *models.PostProcessGcpConnectionV1Request)(
    interface{}, *apiutils.APIError) {

    queryBuilder := g.config.BaseUrl + "/connections/gcp/_post_process"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.gcp-connections=v1+json"
    var result interface{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadGcpConnection Reads a GCP Connection from the given project id
func (g *GcpConnectionsV1) ReadGcpConnection(
    projectId string)(
    *models.ReadGCPConnectionResponse, *apiutils.APIError) {

    pathURL := "/connections/gcp/{project_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "project_id": projectId,
    }
    queryBuilder := g.config.BaseUrl + pathURL

    
    header := "application/api.clumio.gcp-connections=v1+json"
    result := &models.ReadGCPConnectionResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// DeleteGcpConnection Deletes a GCP Connection
func (g *GcpConnectionsV1) DeleteGcpConnection(
    projectId string)(
    interface{}, *apiutils.APIError) {

    pathURL := "/connections/gcp/{project_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "project_id": projectId,
    }
    queryBuilder := g.config.BaseUrl + pathURL

    
    header := "application/api.clumio.gcp-connections=v1+json"
    var result interface{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Delete,
    })

    return result, apiErr
}


// UpdateGcpConnection Updates a GCP Connection having the given project id. This API should only be invoked by the Clumio Terraform provider and should not be invoked manually.
func (g *GcpConnectionsV1) UpdateGcpConnection(
    projectId string, 
    body *models.UpdateGcpConnectionV1Request)(
    *models.UpdateGCPConnectionResponse, *apiutils.APIError) {

    pathURL := "/connections/gcp/{project_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "project_id": projectId,
    }
    queryBuilder := g.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.gcp-connections=v1+json"
    result := &models.UpdateGCPConnectionResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: g.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Patch,
    })

    return result, apiErr
}
