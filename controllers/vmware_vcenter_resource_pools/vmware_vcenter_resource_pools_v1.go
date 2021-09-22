// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcenterresourcepools contains methods related to VmwareVcenterResourcePools
package vmwarevcenterresourcepools

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterResourcePoolsV1 represents a custom type struct
type VmwareVcenterResourcePoolsV1 struct {
    config config.Config
}

//  ListVmwareVcenterResourcePools Returns a list of resource pools in the specified vCenter server.
func (v *VmwareVcenterResourcePoolsV1) ListVmwareVcenterResourcePools(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListResourcePoolsResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/resource-pools"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-resource-pools=v1+json"
    var result *models.ListResourcePoolsResponse
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
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(v.config.Token).
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


//  ReadVmwareVcenterResourcePool Returns a representation of the specified resource pool.
func (v *VmwareVcenterResourcePoolsV1) ReadVmwareVcenterResourcePool(
    vcenterId string, 
    resourcePoolId string)(
    *models.ReadResourcePoolResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/resource-pools/{resource_pool_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "resource_pool_id": resourcePoolId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-resource-pools=v1+json"
    var result *models.ReadResourcePoolResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(v.config.Token).
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
