// Copyright (c) 2023 Clumio All Rights Reserved

// Package awsconnectiongroups contains methods related to AwsConnectionGroups
package awsconnectiongroups

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsConnectionGroupsV1 represents a custom type struct
type AwsConnectionGroupsV1 struct {
    config config.Config
}

// ListAwsConnectionGroups Returns a list of active connection groups that are managing AWS account connections.
func (a *AwsConnectionGroupsV1) ListAwsConnectionGroups(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListConnectionGroupsResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/connections/aws/connection-groups"

    
    header := "application/api.clumio.aws-connection-groups=v1+json"
    result := &models.ListConnectionGroupsResponse{}
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


// CreateAwsConnectionGroup Initiates a new connection group request, and in response returns a URL under field "deployment_template_url". Upon successfully deploying the cloudformation template
//  retrieved from the URL in the AWS account, the connection-group is created or the request is discarded.
//  
//  NOTE - The lifecycle of a connection-group is governed by its CloudFormation stack. If the CFT for a connection-group is deleted from AWS, the connection-group is automatically deleted.
//  The deletion does not affect any backups taken nor any connections it managed.
func (a *AwsConnectionGroupsV1) CreateAwsConnectionGroup(
    body *models.CreateAwsConnectionGroupV1Request)(
    *models.CreateConnectionGroupResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/connections/aws/connection-groups"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.aws-connection-groups=v1+json"
    result := &models.CreateConnectionGroupResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: a.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadAwsConnectionGroup Returns a representation of the specified connection group with the given ID.
func (a *AwsConnectionGroupsV1) ReadAwsConnectionGroup(
    connectionGroupId string, 
    embed *string, 
    returnExternalId *bool)(
    *models.ReadConnectionGroupResponse, *apiutils.APIError) {

    pathURL := "/connections/aws/connection-groups/{connection_group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "connection_group_id": connectionGroupId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-connection-groups=v1+json"
    result := &models.ReadConnectionGroupResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
    }
    if returnExternalId != nil {
        queryParams["return_external_id"] = fmt.Sprintf("%v", *returnExternalId)
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: a.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// UpdateAwsConnectionGroup Initiates an update to a connection group, and in response returns a URL under field "deployment_template_url". Upon successfully updating the connection-group's stack with the cloudformation template
//  retrieved from the URL, the connection-group update completes or the request is discarded.
//  
//  NOTE - The lifecycle of a connection-group is governed by its CloudFormation stack. If the CFT for a connection-group is deleted from AWS, the connection-group is automatically deleted.
//  The deletion does not affect any backups taken nor any connections it managed.
func (a *AwsConnectionGroupsV1) UpdateAwsConnectionGroup(
    connectionGroupId string, 
    body models.UpdateAwsConnectionGroupV1Request)(
    *models.UpdateConnectionGroupResponse, *apiutils.APIError) {

    pathURL := "/connections/aws/connection-groups/{connection_group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "connection_group_id": connectionGroupId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.aws-connection-groups=v1+json"
    result := &models.UpdateConnectionGroupResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: a.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Put,
    })

    return result, apiErr
}
