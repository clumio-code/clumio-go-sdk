// Copyright (c) 2021 Clumio All Rights Reserved

// Package awss3buckets contains methods related to AwsS3Buckets
package awss3buckets

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// AwsS3BucketsV1 represents a custom type struct
type AwsS3BucketsV1 struct {
    config config.Config
}

//  ListAwsS3Buckets Returns a list of S3 buckets.
func (a *AwsS3BucketsV1) ListAwsS3Buckets(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListBucketsResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := a.config.BaseUrl + "/datasources/aws/s3-buckets"

    
    header := "application/aws-s3-buckets=v1+json"
    var result *models.ListBucketsResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    if filter == nil{
        filter = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, a.config.OrganizationalUnitContext).
        SetAuthToken(a.config.Token).
        SetResult(&result).
        Get(queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       common.InternalServerError,
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       common.NonSuccessStatusCodeError,
            Response:     res.Body(),
        }
    }
    return result, nil
}


//  ReadAwsS3Bucket Returns a representation of the specified S3 bucket.
func (a *AwsS3BucketsV1) ReadAwsS3Bucket(
    bucketId string)(
    *models.ReadBucketResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/aws/s3-buckets/{bucket_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "bucket_id": bucketId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/aws-s3-buckets=v1+json"
    var result *models.ReadBucketResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, a.config.OrganizationalUnitContext).
        SetAuthToken(a.config.Token).
        SetResult(&result).
        Get(queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       common.InternalServerError,
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       common.NonSuccessStatusCodeError,
            Response:     res.Body(),
        }
    }
    return result, nil
}
