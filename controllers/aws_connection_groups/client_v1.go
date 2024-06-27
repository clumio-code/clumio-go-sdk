// Copyright (c) 2023 Clumio All Rights Reserved

package awsconnectiongroups

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsConnectionGroupsV1Client represents a custom type interface
type AwsConnectionGroupsV1Client interface {
    // ListAwsConnectionGroups Returns a list of active connection groups that are managing AWS account connections.
    ListAwsConnectionGroups(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListConnectionGroupsResponse,  *apiutils.APIError)
    
    // CreateAwsConnectionGroup Initiates a new connection group request, and in response returns a URL under field "deployment_template_url". Upon successfully deploying the cloudformation template
    //  retrieved from the URL in the AWS account, the connection-group is created or the request is discarded.
    //  
    //  NOTE - The lifecycle of a connection-group is governed by its CloudFormation stack. If the CFT for a connection-group is deleted from AWS, the connection-group is automatically deleted.
    //  The deletion does not affect any backups taken nor any connections it managed.
    CreateAwsConnectionGroup(
        body *models.CreateAwsConnectionGroupV1Request)(
        *models.CreateConnectionGroupResponse,  *apiutils.APIError)
    
    // ReadAwsConnectionGroup Returns a representation of the specified connection group with the given ID.
    ReadAwsConnectionGroup(
        connectionGroupId string, 
        embed *string, 
        returnExternalId *bool)(
        *models.ReadConnectionGroupResponse,  *apiutils.APIError)
    
    // UpdateAwsConnectionGroup Initiates an update to a connection group, and in response returns a URL under field "deployment_template_url". Upon successfully updating the connection-group's stack with the cloudformation template
    //  retrieved from the URL, the connection-group update completes or the request is discarded.
    //  
    //  NOTE - The lifecycle of a connection-group is governed by its CloudFormation stack. If the CFT for a connection-group is deleted from AWS, the connection-group is automatically deleted.
    //  The deletion does not affect any backups taken nor any connections it managed.
    UpdateAwsConnectionGroup(
        connectionGroupId string, 
        body models.UpdateAwsConnectionGroupV1Request)(
        *models.UpdateConnectionGroupResponse,  *apiutils.APIError)
    
}

// NewAwsConnectionGroupsV1 returns AwsConnectionGroupsV1Client
func NewAwsConnectionGroupsV1(config config.Config) AwsConnectionGroupsV1Client {
    client := new(AwsConnectionGroupsV1)
    client.config = config
    return client
}
