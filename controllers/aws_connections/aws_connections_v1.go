// Copyright (c) 2021 Clumio All Rights Reserved

// Package awsconnections contains methods related to AwsConnections
package awsconnections

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumiogosdk/api_utils"
    "github.com/clumio-code/clumiogosdk/config"
    "github.com/clumio-code/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// AwsConnectionsV1 represents a custom type struct
type AwsConnectionsV1 struct {
    config config.Config
}

//  ListAwsConnections Returns a list of AWS Connections
func (a *AwsConnectionsV1) ListAwsConnections(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListAWSConnectionsResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := a.config.BaseUrl + "/connections/aws"

    
    header := "application/aws-connections=v1+json"
    var result *models.ListAWSConnectionsResponse
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
        SetAuthToken(a.config.Token).
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


//  CreateAwsConnection Initiate a new AWS connection.
func (a *AwsConnectionsV1) CreateAwsConnection(
    body *models.CreateAwsConnectionV1Request)(
    *models.CreateAWSConnectionResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := a.config.BaseUrl + "/connections/aws"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/aws-connections=v1+json"
    var result *models.CreateAWSConnectionResponse
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetAuthToken(a.config.Token).
        SetBody(payload).
        SetResult(&result).
        Post(_queryBuilder)

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


//  ReadAwsConnection Returns a representation of the specified AWS connection.
func (a *AwsConnectionsV1) ReadAwsConnection(
    connectionId string)(
    *models.ReadAWSConnectionResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/connections/aws/{connection_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "connectionId": connectionId,
    }
    _queryBuilder := a.config.BaseUrl + _pathURL

    
    header := "application/aws-connections=v1+json"
    var result *models.ReadAWSConnectionResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(a.config.Token).
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


//  DeleteAwsConnection Delete the specified AWS connection.
func (a *AwsConnectionsV1) DeleteAwsConnection(
    connectionId string)(
    interface{}, *apiutils.APIError){

    var err error = nil
    _pathURL := "/connections/aws/{connection_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "connectionId": connectionId,
    }
    _queryBuilder := a.config.BaseUrl + _pathURL

    
    header := "application/aws-connections=v1+json"
    var result interface{}
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(a.config.Token).
        SetResult(&result).
        Delete(_queryBuilder)

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


//  UpdateAwsConnection Returns a new template url for the specified configuration.
func (a *AwsConnectionsV1) UpdateAwsConnection(
    connectionId string, 
    body models.UpdateAwsConnectionV1Request)(
    *models.UpdateAWSConnectionResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/connections/aws/{connection_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "connectionId": connectionId,
    }
    _queryBuilder := a.config.BaseUrl + _pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/aws-connections=v1+json"
    var result *models.UpdateAWSConnectionResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(a.config.Token).
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
