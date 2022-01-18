// Copyright (c) 2021 Clumio All Rights Reserved

// Package organizationalunits contains methods related to OrganizationalUnits
package organizationalunits

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// OrganizationalUnitsV1 represents a custom type struct
type OrganizationalUnitsV1 struct {
    config config.Config
}

//  ListOrganizationalUnits Returns a list of organizational units.
func (o *OrganizationalUnitsV1) ListOrganizationalUnits(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListOrganizationalUnitsResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := o.config.BaseUrl + "/organizational-units"

    
    header := "application/organizational-units=v1+json"
    var result *models.ListOrganizationalUnitsResponse
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
        SetAuthToken(o.config.Token).
        SetResult(&result).
        Get(queryBuilder)

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


//  CreateOrganizationalUnit Create a new organizational unit.
func (o *OrganizationalUnitsV1) CreateOrganizationalUnit(
    embed *string, 
    body *models.CreateOrganizationalUnitV1Request)(
    *models.CreateOrganizationalUnitResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := o.config.BaseUrl + "/organizational-units"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/organizational-units=v1+json"
    var result *models.CreateOrganizationalUnitResponse
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
        SetHeader("Accept", header).
        SetAuthToken(o.config.Token).
        SetBody(payload).
        SetResult(&result).
        Post(queryBuilder)

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


//  ReadOrganizationalUnit Returns a representation of the specified organizational unit.
func (o *OrganizationalUnitsV1) ReadOrganizationalUnit(
    id string)(
    *models.ReadOrganizationalUnitResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/organizational-units/{id}"
    //process optional template parameters
    pathParams := map[string]string{
        "id": id,
    }
    queryBuilder := o.config.BaseUrl + pathURL

    
    header := "application/organizational-units=v1+json"
    var result *models.ReadOrganizationalUnitResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(o.config.Token).
        SetResult(&result).
        Get(queryBuilder)

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


//  DeleteOrganizationalUnit Delete the specified organizational unit.
func (o *OrganizationalUnitsV1) DeleteOrganizationalUnit(
    id string, 
    embed *string)(
    *models.DeleteOrganizationalUnitResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/organizational-units/{id}"
    //process optional template parameters
    pathParams := map[string]string{
        "id": id,
    }
    queryBuilder := o.config.BaseUrl + pathURL

    
    header := "application/organizational-units=v1+json"
    var result *models.DeleteOrganizationalUnitResponse
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
        SetAuthToken(o.config.Token).
        SetResult(&result).
        Delete(queryBuilder)

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


//  PatchOrganizationalUnit Patch the specified organizational unit.
//  The complete updated attribute(s) of the organizational unit has to be provided in the request.
func (o *OrganizationalUnitsV1) PatchOrganizationalUnit(
    id string, 
    embed *string, 
    body *models.PatchOrganizationalUnitV1Request)(
    *models.PatchOrganizationalUnitResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/organizational-units/{id}"
    //process optional template parameters
    pathParams := map[string]string{
        "id": id,
    }
    queryBuilder := o.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/organizational-units=v1+json"
    var result *models.PatchOrganizationalUnitResponse
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
        SetAuthToken(o.config.Token).
        SetBody(payload).
        SetResult(&result).
        Patch(queryBuilder)

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
