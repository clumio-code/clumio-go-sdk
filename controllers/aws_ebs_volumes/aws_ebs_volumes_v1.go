// Copyright (c) 2023 Clumio All Rights Reserved

// Package awsebsvolumes contains methods related to AwsEbsVolumes
package awsebsvolumes

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsEbsVolumesV1 represents a custom type struct
type AwsEbsVolumesV1 struct {
    config config.Config
}

// ListAwsEbsVolumes Returns a list of EBS volumes.
func (a *AwsEbsVolumesV1) ListAwsEbsVolumes(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ListEbsVolumesResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/datasources/aws/ebs-volumes"

    
    header := "application/api.clumio.aws-ebs-volumes=v1+json"
    result := &models.ListEbsVolumesResponse{}
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
    if embed != nil {
        queryParams["embed"] = *embed
    }
    if lookbackDays != nil {
        queryParams["lookback_days"] = fmt.Sprintf("%v", *lookbackDays)
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


// ReadAwsEbsVolume Returns a representation of the specified EBS volume.
func (a *AwsEbsVolumesV1) ReadAwsEbsVolume(
    volumeId string, 
    lookbackDays *int64, 
    embed *string)(
    *models.ReadEbsVolumeResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/ebs-volumes/{volume_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "volume_id": volumeId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-ebs-volumes=v1+json"
    result := &models.ReadEbsVolumeResponse{}
    queryParams := make(map[string]string)
    if lookbackDays != nil {
        queryParams["lookback_days"] = fmt.Sprintf("%v", *lookbackDays)
    }
    if embed != nil {
        queryParams["embed"] = *embed
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: a.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
