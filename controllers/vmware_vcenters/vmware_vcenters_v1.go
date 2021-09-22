// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcenters contains methods related to VmwareVcenters
package vmwarevcenters

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcentersV1 represents a custom type struct
type VmwareVcentersV1 struct {
    config config.Config
}

//  ListVmwareVcenters Returns a list of vCenter servers.
func (v *VmwareVcentersV1) ListVmwareVcenters(
    limit *int64, 
    start *string, 
    embed *string)(
    *models.ListVcentersResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := v.config.BaseUrl + "/datasources/vmware/vcenters"

    
    header := "application/vmware-vcenters=v1+json"
    var result *models.ListVcentersResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
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


//  ReadVmwareVcenter Returns a representation of the specified vCenter server.
func (v *VmwareVcentersV1) ReadVmwareVcenter(
    vcenterId string, 
    embed *string)(
    *models.ReadVcenterResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenters=v1+json"
    var result *models.ReadVcenterResponse
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
