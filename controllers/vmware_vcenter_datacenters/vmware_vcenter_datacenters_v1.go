// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcenterdatacenters contains methods related to VmwareVcenterDatacenters
package vmwarevcenterdatacenters

import (
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterDatacentersV1 represents a custom type struct
type VmwareVcenterDatacentersV1 struct {
    config config.Config
}

//  ListVmwareVcenterDatacenters Returns a list of VMware data centers in the specified vCenter server.
func (v *VmwareVcenterDatacentersV1) ListVmwareVcenterDatacenters(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string, 
    Embed *string)(
    *models.ListDatacentersResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/datacenters"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenterId": vcenterId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-datacenters=v1+json"
    var result *models.ListDatacentersResponse
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
    if Embed == nil{
        Embed = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
        "Embed": *Embed,
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


//  ReadVmwareVcenterDatacenter Returns a representation of the specified VMware data center within the specified vCenter server.
func (v *VmwareVcenterDatacentersV1) ReadVmwareVcenterDatacenter(
    vcenterId string, 
    datacenterId string, 
    embed *string)(
    *models.ReadDatacenterResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/datacenters/{datacenter_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenterId": vcenterId,
        "datacenterId": datacenterId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-datacenters=v1+json"
    var result *models.ReadDatacenterResponse
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
