// Copyright (c) 2023 Clumio All Rights Reserved

// Package awss3buckets contains methods related to AwsS3Buckets
package awss3buckets

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsS3BucketsV1 represents a custom type struct
type AwsS3BucketsV1 struct {
    config config.Config
}

// ListAwsS3Buckets Returns a list of S3 buckets.
func (a *AwsS3BucketsV1) ListAwsS3Buckets(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListBucketsResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/datasources/aws/s3-buckets"

    
    header := "application/api.clumio.aws-s3-buckets=v1+json"
    result := &models.ListBucketsResponse{}
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


// ReadAwsS3Bucket Returns a representation of the specified S3 bucket.
func (a *AwsS3BucketsV1) ReadAwsS3Bucket(
    bucketId string)(
    *models.ReadBucketResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/s3-buckets/{bucket_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "bucket_id": bucketId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-s3-buckets=v1+json"
    result := &models.ReadBucketResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: a.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
