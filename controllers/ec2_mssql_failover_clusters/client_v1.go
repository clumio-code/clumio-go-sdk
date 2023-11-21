// Copyright (c) 2023 Clumio All Rights Reserved

package ec2mssqlfailoverclusters

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// Ec2MssqlFailoverClustersV1Client represents a custom type interface
type Ec2MssqlFailoverClustersV1Client interface {
    // ListEc2MssqlFailoverClusters Returns a list of failover clusters.
    ListEc2MssqlFailoverClusters(
        limit *int64, 
        start *string, 
        filter *string, 
        embed *string)(
        *models.ListEC2MSSQLFCIsResponse,  *apiutils.APIError)
    
}

// NewEc2MssqlFailoverClustersV1 returns Ec2MssqlFailoverClustersV1Client
func NewEc2MssqlFailoverClustersV1(config config.Config) Ec2MssqlFailoverClustersV1Client {
    client := new(Ec2MssqlFailoverClustersV1)
    client.config = config
    return client
}
