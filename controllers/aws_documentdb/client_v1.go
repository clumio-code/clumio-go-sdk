// Copyright (c) 2023 Clumio All Rights Reserved

package awsdocumentdb

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsDocumentdbV1Client represents a custom type interface
type AwsDocumentdbV1Client interface {
    // ListAwsDocumentdb Retrieve a list of DocumentDB clusters.
    ListAwsDocumentdb(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string, 
        embed *string, 
        lookbackDays *int64)(
        *models.ListDocumentDBResponse,  *apiutils.APIError)
    
    // ReadAwsDocumentdb Returns a representation of the specified DocumentDB cluster.
    ReadAwsDocumentdb(
        resourceId string, 
        lookbackDays *int64, 
        embed *string)(
        *models.ReadDocumentDBResponse,  *apiutils.APIError)
    
}

// NewAwsDocumentdbV1 returns AwsDocumentdbV1Client
func NewAwsDocumentdbV1(config config.Config) AwsDocumentdbV1Client {
    client := new(AwsDocumentdbV1)
    client.config = config
    return client
}
