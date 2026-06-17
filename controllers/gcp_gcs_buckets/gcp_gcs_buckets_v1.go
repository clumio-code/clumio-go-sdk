// Copyright (c) 2023 Clumio All Rights Reserved

// Package gcpgcsbuckets contains methods related to GcpGcsBuckets
package gcpgcsbuckets

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpGcsBucketsV1 represents a custom type struct
type GcpGcsBucketsV1 struct {
    config config.Config
}

// ListGcpGcsBuckets Returns a list of GCS buckets.
func (g *GcpGcsBucketsV1) ListGcpGcsBuckets(
    limit *int64, 
    start *string, 
    filter *string, 
    bucketMatcher *string)(
    *models.ListGCSBucketsResponse, *apiutils.APIError) {

    queryBuilder := g.config.BaseUrl + "/datasources/gcp/gcs-buckets"

    
    header := "application/api.clumio.gcp-gcs-buckets=v1+json"
    result := &models.ListGCSBucketsResponse{}
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
    if bucketMatcher != nil {
        queryParams["bucket_matcher"] = *bucketMatcher
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


// ReadGcpGcsBucket Returns a representation of the specified GCS bucket.
func (g *GcpGcsBucketsV1) ReadGcpGcsBucket(
    bucketId string)(
    *models.ReadGCSBucketResponse, *apiutils.APIError) {

    pathURL := "/datasources/gcp/gcs-buckets/{bucket_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "bucket_id": bucketId,
    }
    queryBuilder := g.config.BaseUrl + pathURL

    
    header := "application/api.clumio.gcp-gcs-buckets=v1+json"
    result := &models.ReadGCSBucketResponse{}

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
