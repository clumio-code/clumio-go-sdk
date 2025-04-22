// Copyright (c) 2023 Clumio All Rights Reserved

package restoredawss3buckets

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredAwsS3BucketsV1Client represents a custom type interface
type RestoredAwsS3BucketsV1Client interface {
    // RestoreAwsS3Bucket Restores the specified S3 bucket to the specified target destination using Clumio Backtrack.
    RestoreAwsS3Bucket(
        body models.RestoreAwsS3BucketV1Request)(
        *models.RestoreS3BucketResponse,  *apiutils.APIError)
    
    // PreviewAwsS3Bucket Previews the objects that will be restored by S3 bucket restore.
    PreviewAwsS3Bucket(
        bucketId string, 
        body models.PreviewAwsS3BucketV1Request)(
        *models.PreviewS3BucketResponse,  *apiutils.APIError)
    
    // PreviewDetailsAwsS3Bucket Details of an S3 bucket restore preview task, started by PreviewS3Bucket API.
    PreviewDetailsAwsS3Bucket(
        bucketId string, 
        previewId string)(
        *models.PreviewDetailsS3BucketResponse,  *apiutils.APIError)
    
}

// NewRestoredAwsS3BucketsV1 returns RestoredAwsS3BucketsV1Client
func NewRestoredAwsS3BucketsV1(config config.Config) RestoredAwsS3BucketsV1Client {
    client := new(RestoredAwsS3BucketsV1)
    client.config = config
    return client
}
