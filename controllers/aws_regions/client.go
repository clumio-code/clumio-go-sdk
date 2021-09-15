// Copyright (c) 2021 Clumio All Rights Reserved

package awsregions

import (
     "github.com/clumio-code/clumiogosdk/api_utils"
     "github.com/clumio-code/clumiogosdk/config"
     "github.com/clumio-code/clumiogosdk/models"
)

// AwsRegionsV1Client represents a custom type interface
type AwsRegionsV1Client interface {
    //  Returns a list of valid regions for creating AWS connections
    ListConnectionAwsRegions(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListAWSRegionsResponse,  *apiutils.APIError)
    
}

// NewAwsRegionsV1 returns AwsRegionsV1Client
func NewAwsRegionsV1(config config.Config) AwsRegionsV1Client{
    client := new(AwsRegionsV1)
    client.config = config
    return client
}
