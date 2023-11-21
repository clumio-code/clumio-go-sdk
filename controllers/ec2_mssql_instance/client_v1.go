// Copyright (c) 2023 Clumio All Rights Reserved

package ec2mssqlinstance

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlInstanceV1Client represents a custom type interface
type Ec2MssqlInstanceV1Client interface {
    // ListEc2MssqlInstances Returns a list of Instances
    ListEc2MssqlInstances(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListEC2MSSQLInstancesResponse,  *apiutils.APIError)
    
    // ReadEc2MssqlInstance Returns a representation of the specified instance.
    ReadEc2MssqlInstance(
        instanceId string)(
        *models.ReadEC2MSSQLInstanceResponse,  *apiutils.APIError)
    
}

// NewEc2MssqlInstanceV1 returns Ec2MssqlInstanceV1Client
func NewEc2MssqlInstanceV1(config config.Config) Ec2MssqlInstanceV1Client {
    client := new(Ec2MssqlInstanceV1)
    client.config = config
    return client
}
