// Copyright (c) 2023 Clumio All Rights Reserved

// Package awsec2instances contains methods related to AwsEc2Instances
package awsec2instances

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsEc2InstancesV1 represents a custom type struct
type AwsEc2InstancesV1 struct {
    config config.Config
}

// ListAwsEc2Instances Returns a list of EC2 instances.
func (a *AwsEc2InstancesV1) ListAwsEc2Instances(
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string, 
    lookbackDays *int64)(
    *models.ListEc2InstancesResponse, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/datasources/aws/ec2-instances"

    
    header := "application/api.clumio.aws-ec2-instances=v1+json"
    result := &models.ListEc2InstancesResponse{}
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


// ReadAwsEc2Instance Returns a representation of the specified EC2 instance.
func (a *AwsEc2InstancesV1) ReadAwsEc2Instance(
    instanceId string, 
    lookbackDays *int64, 
    embed *string)(
    *models.ReadEc2InstanceResponse, *apiutils.APIError) {

    pathURL := "/datasources/aws/ec2-instances/{instance_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "instance_id": instanceId,
    }
    queryBuilder := a.config.BaseUrl + pathURL

    
    header := "application/api.clumio.aws-ec2-instances=v1+json"
    result := &models.ReadEc2InstanceResponse{}
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
