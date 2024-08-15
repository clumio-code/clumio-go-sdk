// Copyright (c) 2023 Clumio All Rights Reserved

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

// OrganizationalUnitsV2 represents a custom type struct
type OrganizationalUnitsV2 struct {
    config config.Config
}

// ListOrganizationalUnits Returns a list of organizational units.
func (o *OrganizationalUnitsV2) ListOrganizationalUnits(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListOrganizationalUnitsResponse, *apiutils.APIError) {

    queryBuilder := o.config.BaseUrl + "/organizational-units"

    
    header := "application/api.clumio.organizational-units=v2+json"
    result := &models.ListOrganizationalUnitsResponse{}
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
func (o *OrganizationalUnitsV2) CreateOrganizationalUnit(
    embed *string, 
    body *models.CreateOrganizationalUnitV2Request)(
    *models.CreateOrganizationalUnitResponseWrapper, *apiutils.APIError) {

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
    header := "application/api.clumio.organizational-units=v2+json"
    result := &models.CreateOrganizationalUnitResponseWrapper{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
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
func (o *OrganizationalUnitsV2) ReadOrganizationalUnit(
    id string, 
    embed *string)(
    *models.ReadOrganizationalUnitResponse, *apiutils.APIError) {

    pathURL := "/organizational-units/{id}"
    //process optional template parameters
    pathParams := map[string]string{
        "id": id,
    }
    queryBuilder := o.config.BaseUrl + pathURL

    
    header := "application/api.clumio.organizational-units=v2+json"
    result := &models.ReadOrganizationalUnitResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
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
func (o *OrganizationalUnitsV2) DeleteOrganizationalUnit(
    id string, 
    embed *string)(
    *models.DeleteOrganizationalUnitResponse, *apiutils.APIError) {

    pathURL := "/organizational-units/{id}"
    //process optional template parameters
    pathParams := map[string]string{
        "id": id,
    }
    queryBuilder := o.config.BaseUrl + pathURL

    
    header := "application/api.clumio.organizational-units=v2+json"
    result := &models.DeleteOrganizationalUnitResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
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
//  The complete updated attribute(s) of the organizational unit must be provided in
//  the request. Adding or removing entities from the OU is an asynchronous operation
//  and has an associated task. When the request has entities to be added or removed,
//  the response contains a task ID that can be used to track the progress of the
//  operation.
func (o *OrganizationalUnitsV2) PatchOrganizationalUnit(
    id string, 
    embed *string, 
    body *models.PatchOrganizationalUnitV2Request)(
    *models.PatchOrganizationalUnitResponseWrapper, *apiutils.APIError) {

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
    header := "application/api.clumio.organizational-units=v2+json"
    result := &models.PatchOrganizationalUnitResponseWrapper{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
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
