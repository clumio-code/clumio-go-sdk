// Copyright (c) 2021 Clumio All Rights Reserved

// Package consolidatedalerts contains methods related to ConsolidatedAlerts
package consolidatedalerts

import (
    "encoding/json"
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// ConsolidatedAlertsV1 represents a custom type struct
type ConsolidatedAlertsV1 struct {
    config config.Config
}

//  ListConsolidatedAlerts Returns a list of consolidated alerts.
func (c *ConsolidatedAlertsV1) ListConsolidatedAlerts(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListConsolidatedAlertsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := c.config.BaseUrl + "/alerts/consolidated"

    
    header := "application/consolidated-alerts=v1+json"
    var result *models.ListConsolidatedAlertsResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    if filter == nil{
        filter = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetAuthToken(c.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}


//  ReadConsolidatedAlert Returns a representation of the specified consolidated alert.
func (c *ConsolidatedAlertsV1) ReadConsolidatedAlert(
    id string)(
    *models.ReadConsolidatedAlertResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/alerts/consolidated/{id}"
    //process optional template parameters
    pathParams := map[string]string{
        "id": id,
    }
    _queryBuilder := c.config.BaseUrl + _pathURL

    
    header := "application/consolidated-alerts=v1+json"
    var result *models.ReadConsolidatedAlertResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(c.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}


//  UpdateConsolidatedAlert Manages the specified consolidated alert. Managing a consolidated alert includes clearing the alert and adding notes to the specified consolidated alert.
func (c *ConsolidatedAlertsV1) UpdateConsolidatedAlert(
    id string, 
    body *models.UpdateConsolidatedAlertV1Request)(
    *models.UpdateConsolidatedAlertResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/alerts/consolidated/{id}"
    //process optional template parameters
    pathParams := map[string]string{
        "id": id,
    }
    _queryBuilder := c.config.BaseUrl + _pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/consolidated-alerts=v1+json"
    var result *models.UpdateConsolidatedAlertResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(c.config.Token).
        SetBody(payload).
        SetResult(&result).
        Patch(_queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       "Internal Server Error",
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       "Non-success status code returned.",
            Response:     res.Body(),
        }
    }
    return result, nil
}
