// Copyright (c) 2021 Clumio All Rights Reserved

// Package individualalerts contains methods related to IndividualAlerts
package individualalerts

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// IndividualAlertsV1 represents a custom type struct
type IndividualAlertsV1 struct {
    config config.Config
}

//  ListIndividualAlerts Returns a list of individual alerts.
//  
//  Each alert is associated with a cause, which represents the issue that generated the alert,
//  and each cause belongs to a general alert alert type. Some alert types may be associated with multiple causes.
//  
//  The following table lists the Clumio alert types:
//  
//  
//  +--------------------------+---------------------------------------------------+
//  |        Alert Type        |                    Description                    |
//  +==========================+===================================================+
//  | policy_violated          | An entity's scheduled backup failed.              |
//  +--------------------------+---------------------------------------------------+
//  | restore_failed           | An entity restore failed.                         |
//  +--------------------------+---------------------------------------------------+
//  | file_restore_failed      | A file restore failed.                            |
//  +--------------------------+---------------------------------------------------+
//  | tag_conflict             | An EBS volume has multiple associated tags        |
//  |                          | with different protection policies applied at the |
//  |                          | tag level.                                        |
//  +--------------------------+---------------------------------------------------+
//  | aws_account_disconnected | The connection                                    |
//  |                          | between Clumio Cloud Connector and the user's AWS |
//  |                          | account failed.                                   |
//  +--------------------------+---------------------------------------------------+
//  | enc_key_inaccessible     | An issue is blocking encryption key access.       |
//  +--------------------------+---------------------------------------------------+
//  
func (i *IndividualAlertsV1) ListIndividualAlerts(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string, 
    embed *string)(
    *models.ListAlertsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := i.config.BaseUrl + "/alerts/individual"

    
    header := "application/individual-alerts=v1+json"
    var result *models.ListAlertsResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    if sort == nil{
        sort = &defaultString
    }
    if filter == nil{
        filter = &defaultString
    }
    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "sort": *sort,
        "filter": *filter,
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader("Accept", header).
        SetAuthToken(i.config.Token).
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


//  ReadIndividualAlert Returns a representation of the specified individual alert.
func (i *IndividualAlertsV1) ReadIndividualAlert(
    individualAlertId string, 
    embed *string)(
    *models.ReadAlertResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/alerts/individual/{individual_alert_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "individual_alert_id": individualAlertId,
    }
    _queryBuilder := i.config.BaseUrl + _pathURL

    
    header := "application/individual-alerts=v1+json"
    var result *models.ReadAlertResponse
    client := resty.New()
    defaultString := "" 
    

    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(i.config.Token).
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


//  UpdateIndividualAlert Manages an existing individual alert. Managing an individual alert includes clearing the alert and adding notes to the specified alert.
func (i *IndividualAlertsV1) UpdateIndividualAlert(
    individualAlertId string, 
    embed *string, 
    body *models.UpdateIndividualAlertV1Request)(
    *models.UpdateAlertResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/alerts/individual/{individual_alert_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "individual_alert_id": individualAlertId,
    }
    _queryBuilder := i.config.BaseUrl + _pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/individual-alerts=v1+json"
    var result *models.UpdateAlertResponse
    client := resty.New()
    defaultString := "" 
    

    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(i.config.Token).
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
