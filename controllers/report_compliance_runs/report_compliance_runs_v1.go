// Copyright (c) 2023 Clumio All Rights Reserved

// Package reportcomplianceruns contains methods related to ReportComplianceRuns
package reportcomplianceruns

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// ReportComplianceRunsV1 represents a custom type struct
type ReportComplianceRunsV1 struct {
    config config.Config
}

// ListComplianceReportRuns Get a list of all the compliance report runs belonging to the configuration.
func (r *ReportComplianceRunsV1) ListComplianceReportRuns(
    configurationId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListComplianceRunsResponse, *apiutils.APIError) {

    pathURL := "/reports/compliance/configurations/{configuration_id}/runs"
    //process optional template parameters
    pathParams := map[string]string{
        "configuration_id": configurationId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.report-compliance-runs=v1+json"
    result := &models.ListComplianceRunsResponse{}
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
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreateComplianceReportRun Create a new compliance report run.
func (r *ReportComplianceRunsV1) CreateComplianceReportRun(
    configurationId string, 
    body *models.CreateComplianceReportRunV1Request)(
    *models.CreateComplianceRunResponse, *apiutils.APIError) {

    pathURL := "/reports/compliance/configurations/{configuration_id}/runs"
    //process optional template parameters
    pathParams := map[string]string{
        "configuration_id": configurationId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.report-compliance-runs=v1+json"
    result := &models.CreateComplianceRunResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// DeleteComplianceReportRun Delete a compliance report run.
func (r *ReportComplianceRunsV1) DeleteComplianceReportRun(
    configurationId string, 
    runId string)(
    interface{}, *apiutils.APIError) {

    pathURL := "/reports/compliance/configurations/{configuration_id}/runs/{run_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "configuration_id": configurationId,
        "run_id": runId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.report-compliance-runs=v1+json"
    var result interface{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Delete,
    })

    return result, apiErr
}


// SendComplianceReportRunEmail Send a compliance report run to the given emails.
func (r *ReportComplianceRunsV1) SendComplianceReportRunEmail(
    configurationId string, 
    runId string, 
    body *models.SendComplianceReportRunEmailV1Request)(
    *models.SendComplianceRunEmailResponse, *apiutils.APIError) {

    pathURL := "/reports/compliance/configurations/{configuration_id}/runs/{run_id}/_notify"
    //process optional template parameters
    pathParams := map[string]string{
        "configuration_id": configurationId,
        "run_id": runId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.report-compliance-runs=v1+json"
    result := &models.SendComplianceRunEmailResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}
