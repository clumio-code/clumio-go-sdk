// Copyright (c) 2023 Clumio All Rights Reserved

package gcpgcsbuckets

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpGcsBucketsV1Client represents a custom type interface
type GcpGcsBucketsV1Client interface {
    // ListGcpGcsBuckets Returns a list of GCS buckets.
    ListGcpGcsBuckets(
        limit *int64, 
        start *string, 
        filter *string, 
        bucketMatcher *string)(
        *models.ListGCSBucketsResponse,  *apiutils.APIError)
    
    // ReadGcpGcsBucket Returns a representation of the specified GCS bucket.
    ReadGcpGcsBucket(
        bucketId string)(
        *models.ReadGCSBucketResponse,  *apiutils.APIError)
    
}

// NewGcpGcsBucketsV1 returns GcpGcsBucketsV1Client
func NewGcpGcsBucketsV1(config config.Config) GcpGcsBucketsV1Client {
    client := new(GcpGcsBucketsV1)
    client.config = config
    return client
}
