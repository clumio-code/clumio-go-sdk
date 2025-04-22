// Copyright (c) 2023 Clumio All Rights Reserved

// Package restoredawss3buckets contains methods related to RestoredAwsS3Buckets
package restoredawss3buckets

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredAwsS3BucketsV1 represents a custom type struct
type RestoredAwsS3BucketsV1 struct {
    config config.Config
}

// RestoreAwsS3Bucket Restores the specified S3 bucket to the specified target destination using Clumio Backtrack.
func (r *RestoredAwsS3BucketsV1) RestoreAwsS3Bucket(
    body models.RestoreAwsS3BucketV1Request)(
    *models.RestoreS3BucketResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/aws/s3-buckets"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-aws-s3-buckets=v1+json"
    result := &models.RestoreS3BucketResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// PreviewAwsS3Bucket Previews the objects that will be restored by S3 bucket restore.
func (r *RestoredAwsS3BucketsV1) PreviewAwsS3Bucket(
    bucketId string, 
    body models.PreviewAwsS3BucketV1Request)(
    *models.PreviewS3BucketResponse, *apiutils.APIError) {

    pathURL := "/restores/aws/s3-buckets/{bucket_id}/previews"
    //process optional template parameters
    pathParams := map[string]string{
        "bucket_id": bucketId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-aws-s3-buckets=v1+json"
    result := &models.PreviewS3BucketResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// PreviewDetailsAwsS3Bucket Details of an S3 bucket restore preview task, started by PreviewS3Bucket API.
func (r *RestoredAwsS3BucketsV1) PreviewDetailsAwsS3Bucket(
    bucketId string, 
    previewId string)(
    *models.PreviewDetailsS3BucketResponse, *apiutils.APIError) {

    pathURL := "/restores/aws/s3-buckets/{bucket_id}/previews/{preview_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "bucket_id": bucketId,
        "preview_id": previewId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.restored-aws-s3-buckets=v1+json"
    result := &models.PreviewDetailsS3BucketResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
