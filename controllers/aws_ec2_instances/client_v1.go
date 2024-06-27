// Copyright (c) 2023 Clumio All Rights Reserved

package awsec2instances

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsEc2InstancesV1Client represents a custom type interface
type AwsEc2InstancesV1Client interface {
    // ListAwsEc2Instances Returns a list of EC2 instances.
    ListAwsEc2Instances(
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string, 
        lookbackDays *int64)(
        *models.ListEc2InstancesResponse,  *apiutils.APIError)
    
    // ReadAwsEc2Instance Returns a representation of the specified EC2 instance.
    ReadAwsEc2Instance(
        instanceId string, 
        lookbackDays *int64, 
        embed *string)(
        *models.ReadEc2InstanceResponse,  *apiutils.APIError)
    
}

// NewAwsEc2InstancesV1 returns AwsEc2InstancesV1Client
func NewAwsEc2InstancesV1(config config.Config) AwsEc2InstancesV1Client {
    client := new(AwsEc2InstancesV1)
    client.config = config
    return client
}
