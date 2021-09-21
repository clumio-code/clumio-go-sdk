// Copyright (c) 2021 Clumio All Rights Reserved

// Package awsebsvolumes contains methods related to AwsEbsVolumes
package awsebsvolumes

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// AwsEbsVolumesV1 represents a custom type struct
type AwsEbsVolumesV1 struct {
    config config.Config
}

//  ListAwsEbsVolumes Returns a list of EBS volumes.
func (a *AwsEbsVolumesV1) ListAwsEbsVolumes(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListEbsVolumesResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := a.config.BaseUrl + "/datasources/aws/ebs-volumes"

    
    header := "application/aws-ebs-volumes=v1+json"
    var result *models.ListEbsVolumesResponse
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
    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetAuthToken(a.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}


//  ReadAwsEbsVolume Returns a representation of the specified EBS volume.
func (a *AwsEbsVolumesV1) ReadAwsEbsVolume(
    volumeId string, 
    embed *string)(
    *models.ReadEbsVolumeResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/aws/ebs-volumes/{volume_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "volumeId": volumeId,
    }
    _queryBuilder := a.config.BaseUrl + _pathURL

    
    header := "application/aws-ebs-volumes=v1+json"
    var result *models.ReadEbsVolumeResponse
    client := resty.New()
    defaultString := "" 
    

    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(a.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}
