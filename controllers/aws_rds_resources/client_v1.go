// Copyright (c) 2023 Clumio All Rights Reserved

package awsrdsresources

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsRdsResourcesV1Client represents a custom type interface
type AwsRdsResourcesV1Client interface {
    // ListAwsRdsResources Retrieve a list of RDS resources.
    ListAwsRdsResources(
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string)(
        *models.ListRdsResourcesResponse,  *apiutils.APIError)
    
    // ReadAwsRdsResource Returns a representation of the specified RDS resource.
    ReadAwsRdsResource(
        resourceId string, 
        embed *string)(
        *models.ReadRdsResourceResponse,  *apiutils.APIError)
    
}

// NewAwsRdsResourcesV1 returns AwsRdsResourcesV1Client
func NewAwsRdsResourcesV1(config config.Config) AwsRdsResourcesV1Client {
    client := new(AwsRdsResourcesV1)
    client.config = config
    return client
}
