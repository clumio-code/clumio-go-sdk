// Copyright (c) 2023 Clumio All Rights Reserved

// Package reportcompliance contains methods related to ReportCompliance
package reportcompliance

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// ReportComplianceV1 represents a custom type struct
type ReportComplianceV1 struct {
    config config.Config
}

// ListComplianceReportConfigurations Get a list of all the compliance report configurations.
func (r *ReportComplianceV1) ListComplianceReportConfigurations(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListComplianceConfigurationsResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/reports/compliance/configurations"

    
    header := "application/api.clumio.report-compliance=v1+json"
    result := &models.ListComplianceConfigurationsResponse{}
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
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreateComplianceReportConfiguration Create a compliance report configuration.
func (r *ReportComplianceV1) CreateComplianceReportConfiguration(
    body *models.CreateComplianceReportConfigurationV1Request)(
    *models.CreateComplianceConfigurationResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/reports/compliance/configurations"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.report-compliance=v1+json"
    result := &models.CreateComplianceConfigurationResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadComplianceReportConfiguration Returns a representation of the specified compliance report configuration.
func (r *ReportComplianceV1) ReadComplianceReportConfiguration(
    configurationId string)(
    *models.ReadComplianceConfigurationResponse, *apiutils.APIError) {

    pathURL := "/reports/compliance/configurations/{configuration_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "configuration_id": configurationId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.report-compliance=v1+json"
    result := &models.ReadComplianceConfigurationResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// UpdateComplianceReportConfiguration Update a compliance report configuration with the id {configuration_id}.
func (r *ReportComplianceV1) UpdateComplianceReportConfiguration(
    configurationId string, 
    body *models.UpdateComplianceReportConfigurationV1Request)(
    *models.UpdateComplianceConfigurationResponse, *apiutils.APIError) {

    pathURL := "/reports/compliance/configurations/{configuration_id}"
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
    header := "application/api.clumio.report-compliance=v1+json"
    result := &models.UpdateComplianceConfigurationResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Put,
    })

    return result, apiErr
}


// DeleteComplianceReportConfiguration Delete a compliance report configuration with the id {configuration_id}.
func (r *ReportComplianceV1) DeleteComplianceReportConfiguration(
    configurationId string)(
    interface{}, *apiutils.APIError) {

    pathURL := "/reports/compliance/configurations/{configuration_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "configuration_id": configurationId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/api.clumio.report-compliance=v1+json"
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
