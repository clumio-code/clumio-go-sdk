// Copyright (c) 2023 Clumio All Rights Reserved

package reportcomplianceruns

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// ReportComplianceRunsV1Client represents a custom type interface
type ReportComplianceRunsV1Client interface {
    // ListComplianceReportRuns Get a list of all the compliance report runs belonging to the configuration.
    ListComplianceReportRuns(
        configurationId string, 
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListComplianceRunsResponse,  *apiutils.APIError)
    
    // CreateComplianceReportRun Create a new compliance report run.
    CreateComplianceReportRun(
        configurationId string, 
        body *models.CreateComplianceReportRunV1Request)(
        *models.CreateComplianceRunResponse,  *apiutils.APIError)
    
    // DeleteComplianceReportRun Delete a compliance report run.
    DeleteComplianceReportRun(
        configurationId string, 
        runId string)(
        interface{},  *apiutils.APIError)
    
    // SendComplianceReportRunEmail Send a compliance report run to the given emails.
    SendComplianceReportRunEmail(
        configurationId string, 
        runId string, 
        body *models.SendComplianceReportRunEmailV1Request)(
        *models.SendComplianceRunEmailResponse,  *apiutils.APIError)
    
}

// NewReportComplianceRunsV1 returns ReportComplianceRunsV1Client
func NewReportComplianceRunsV1(config config.Config) ReportComplianceRunsV1Client {
    client := new(ReportComplianceRunsV1)
    client.config = config
    return client
}
