// Copyright (c) 2021 Clumio All Rights Reserved

package awsconnections

import (
     "github.com/clumio/clumiogosdk/api_utils"
     "github.com/clumio/clumiogosdk/config"
     "github.com/clumio/clumiogosdk/models"
)

// AwsConnectionsV1Client represents a custom type interface
type AwsConnectionsV1Client interface {
    //  Returns a list of AWS Connections
    ListAwsConnections(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListAWSConnectionsResponse,  *apiutils.APIError)
    
    //  Initiate a new AWS connection.
    CreateAwsConnection(
        body *models.CreateAwsConnectionV1Request)(
        *models.CreateAWSConnectionResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified AWS connection.
    ReadAwsConnection(
        connectionId string)(
        *models.ReadAWSConnectionResponse,  *apiutils.APIError)
    
    //  Delete the specified AWS connection.
    DeleteAwsConnection(
        connectionId string)(
        interface{},  *apiutils.APIError)
    
    //  Returns a new template url for the specified configuration.
    UpdateAwsConnection(
        connectionId string, 
        body models.UpdateAwsConnectionV1Request)(
        *models.UpdateAWSConnectionResponse,  *apiutils.APIError)
    
}

// NewAwsConnectionsV1 returns AwsConnectionsV1Client
func NewAwsConnectionsV1(config config.Config) AwsConnectionsV1Client{
    client := new(AwsConnectionsV1)
    client.config = config
    return client
}
