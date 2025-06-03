// Copyright (c) 2023 Clumio All Rights Reserved

package reportcompliance

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// ReportComplianceV1Client represents a custom type interface
type ReportComplianceV1Client interface {
    // ListComplianceReportConfigurations Get a list of all the compliance report configurations.
    ListComplianceReportConfigurations(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListComplianceConfigurationsResponse,  *apiutils.APIError)
    
    // CreateComplianceReportConfiguration Create a compliance report configuration.
    CreateComplianceReportConfiguration(
        body *models.CreateComplianceReportConfigurationV1Request)(
        *models.CreateComplianceConfigurationResponse,  *apiutils.APIError)
    
    // ReadComplianceReportConfiguration Returns a representation of the specified compliance report configuration.
    ReadComplianceReportConfiguration(
        configurationId string)(
        *models.ReadComplianceConfigurationResponse,  *apiutils.APIError)
    
    // UpdateComplianceReportConfiguration Update a compliance report configuration with the id {configuration_id}.
    UpdateComplianceReportConfiguration(
        configurationId string, 
        body *models.UpdateComplianceReportConfigurationV1Request)(
        *models.UpdateComplianceConfigurationResponse,  *apiutils.APIError)
    
    // DeleteComplianceReportConfiguration Delete a compliance report configuration with the id {configuration_id}.
    DeleteComplianceReportConfiguration(
        configurationId string)(
        interface{},  *apiutils.APIError)
    
}

// NewReportComplianceV1 returns ReportComplianceV1Client
func NewReportComplianceV1(config config.Config) ReportComplianceV1Client {
    client := new(ReportComplianceV1)
    client.config = config
    return client
}
