// Copyright (c) 2021 Clumio All Rights Reserved

package ec2mssqlfailovercluster

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlFailoverClusterV1Client represents a custom type interface
type Ec2MssqlFailoverClusterV1Client interface {
    // ReadEc2MssqlFailoverCluster Returns a representation of the specified failover cluster.
    ReadEc2MssqlFailoverCluster(
        failoverClusterId string)(
        *models.ReadEC2MSSQLFCIResponse,  *apiutils.APIError)
    
}

// NewEc2MssqlFailoverClusterV1 returns Ec2MssqlFailoverClusterV1Client
func NewEc2MssqlFailoverClusterV1(config config.Config) Ec2MssqlFailoverClusterV1Client {
    client := new(Ec2MssqlFailoverClusterV1)
    client.config = config
    return client
}
