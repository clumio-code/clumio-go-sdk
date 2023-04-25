// Copyright (c) 2021 Clumio All Rights Reserved

package ec2mssqlhosts

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlHostsV1Client represents a custom type interface
type Ec2MssqlHostsV1Client interface {
    // ListEc2MssqlHosts Returns a list of EC2 MSSQL hosts
    ListEc2MssqlHosts(
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string)(
        *models.ListEC2MSSQLInvHostsResponse,  *apiutils.APIError)
    
    // ReadEc2MssqlHost Returns a representation of the specified host.
    ReadEc2MssqlHost(
        hostId string)(
        *models.ReadEC2MSSQLInvHostResponse,  *apiutils.APIError)
    
}

// NewEc2MssqlHostsV1 returns Ec2MssqlHostsV1Client
func NewEc2MssqlHostsV1(config config.Config) Ec2MssqlHostsV1Client {
    client := new(Ec2MssqlHostsV1)
    client.config = config
    return client
}
