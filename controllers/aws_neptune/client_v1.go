// Copyright (c) 2023 Clumio All Rights Reserved

package awsneptune

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsNeptuneV1Client represents a custom type interface
type AwsNeptuneV1Client interface {
    // ListAwsNeptune Retrieve a list of Neptune clusters.
    ListAwsNeptune(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string, 
        embed *string, 
        lookbackDays *int64)(
        *models.ListNeptuneResponse,  *apiutils.APIError)
    
    // ReadAwsNeptune Returns a representation of the specified Neptune cluster.
    ReadAwsNeptune(
        resourceId string, 
        lookbackDays *int64, 
        embed *string)(
        *models.ReadNeptuneResponse,  *apiutils.APIError)
    
}

// NewAwsNeptuneV1 returns AwsNeptuneV1Client
func NewAwsNeptuneV1(config config.Config) AwsNeptuneV1Client {
    client := new(AwsNeptuneV1)
    client.config = config
    return client
}
