// Copyright (c) 2023 Clumio All Rights Reserved

package gcplabels

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// GcpLabelsV1Client represents a custom type interface
type GcpLabelsV1Client interface {
    // ListGcpLabelKeys Returns a list of GCP label keys.
    ListGcpLabelKeys(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListGCPLabelKeysResponse,  *apiutils.APIError)
    
    // ListGcpLabelValues Returns a list of GCP label values for the specified label key.
    ListGcpLabelValues(
        labelKeyId string, 
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListGCPLabelValuesResponse,  *apiutils.APIError)
    
}

// NewGcpLabelsV1 returns GcpLabelsV1Client
func NewGcpLabelsV1(config config.Config) GcpLabelsV1Client {
    client := new(GcpLabelsV1)
    client.config = config
    return client
}
