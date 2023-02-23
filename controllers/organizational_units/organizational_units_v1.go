// Copyright (c) 2021 Clumio All Rights Reserved

// Package organizationalunits contains methods related to OrganizationalUnits
package organizationalunits

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// OrganizationalUnitsV1 represents a custom type struct
type OrganizationalUnitsV1 struct {
    config config.Config
}

// ListOrganizationalUnits Returns a list of organizational units.
func (o *OrganizationalUnitsV1) ListOrganizationalUnits(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListOrganizationalUnitsResponseV1, *apiutils.APIError) {

    queryBuilder := o.config.BaseUrl + "/organizational-units"

    
    header := "application/api.clumio.organizational-units=v1+json"
    result := &models.ListOrganizationalUnitsResponseV1{}
    defaultInt64 := int64(0)
    defaultString := "" 
    
    if limit == nil {
        limit = &defaultInt64
    }
    if start == nil {
        start = &defaultString
    }
    if filter == nil {
        filter = &defaultString
    }
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: o.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreateOrganizationalUnit Create a new organizational unit. Adding entities to the OU is an asynchronous operation and has a task associated.
//  When the request has entities to be added, the response has a task ID which can be used to
//  track the progress of the operation.
func (o *OrganizationalUnitsV1) CreateOrganizationalUnit(
    embed *string, 
    body *models.CreateOrganizationalUnitV1Request)(
    *models.CreateOrganizationalUnitV1ResponseWrapper, *apiutils.APIError) {

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
    header := "application/api.clumio.organizational-units=v1+json"
    result := &models.CreateOrganizationalUnitV1ResponseWrapper{}
    defaultString := "" 
    
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
        "embed": *embed,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: o.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result.Http200,
        Result202: &result.Http202,
        RequestType: common.Post,
    })
    if result.Http200 != nil {
        result.StatusCode = 200
    } else if result.Http202 != nil {
        result.StatusCode = 202
    } 

    return result, apiErr
}


// ReadOrganizationalUnit Returns a representation of the specified organizational unit.
func (o *OrganizationalUnitsV1) ReadOrganizationalUnit(
    id string, 
    embed *string)(
    *models.ReadOrganizationalUnitResponseV1, *apiutils.APIError) {

    pathURL := "/organizational-units/{id}"
    //process optional template parameters
    pathParams := map[string]string{
        "id": id,
    }
    queryBuilder := o.config.BaseUrl + pathURL

    
    header := "application/api.clumio.organizational-units=v1+json"
    result := &models.ReadOrganizationalUnitResponseV1{}
    defaultString := "" 
    
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
        "embed": *embed,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: o.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// DeleteOrganizationalUnit Delete the specified organizational unit.
func (o *OrganizationalUnitsV1) DeleteOrganizationalUnit(
    id string, 
    embed *string)(
    *models.DeleteOrganizationalUnitResponse, *apiutils.APIError) {

    pathURL := "/organizational-units/{id}"
    //process optional template parameters
    pathParams := map[string]string{
        "id": id,
    }
    queryBuilder := o.config.BaseUrl + pathURL

    
    header := "application/api.clumio.organizational-units=v1+json"
    result := &models.DeleteOrganizationalUnitResponse{}
    defaultString := "" 
    
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
        "embed": *embed,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: o.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result202: &result,
        RequestType: common.Delete,
    })

    return result, apiErr
}


// PatchOrganizationalUnit Patch the specified organizational unit.
//  The complete updated attribute(s) of the organizational unit have to be provided in the request.
//  Adding or removing entities from the OU is an asynchronous operation and has a task associated.
//  When the request has entities to be added or removed, the response has a task ID
//  which can be used to track the progress of the operation.
func (o *OrganizationalUnitsV1) PatchOrganizationalUnit(
    id string, 
    embed *string, 
    body *models.PatchOrganizationalUnitV1Request)(
    *models.PatchOrganizationalUnitV1ResponseWrapper, *apiutils.APIError) {

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
    header := "application/api.clumio.organizational-units=v1+json"
    result := &models.PatchOrganizationalUnitV1ResponseWrapper{}
    defaultString := "" 
    
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
        "embed": *embed,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: o.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result.Http200,
        Result202: &result.Http202,
        RequestType: common.Patch,
    })
    if result.Http200 != nil {
        result.StatusCode = 200
    } else if result.Http202 != nil {
        result.StatusCode = 202
    } 

    return result, apiErr
}
