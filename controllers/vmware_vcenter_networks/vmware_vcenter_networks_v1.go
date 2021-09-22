// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcenternetworks contains methods related to VmwareVcenterNetworks
package vmwarevcenternetworks

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterNetworksV1 represents a custom type struct
type VmwareVcenterNetworksV1 struct {
    config config.Config
}

//  ListVmwareVcenterNetworks Returns a list of networks in the specified vCenter server.
func (v *VmwareVcenterNetworksV1) ListVmwareVcenterNetworks(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListVMwareVCenterNetworksResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/networks"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-networks=v1+json"
    var result *models.ListVMwareVCenterNetworksResponse
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


//  ReadVmwareVcenterNetwork Returns a representation of the specified network.
func (v *VmwareVcenterNetworksV1) ReadVmwareVcenterNetwork(
    vcenterId string, 
    networkId string)(
    *models.ReadVMwareVCenterNetworkResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/networks/{network_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "network_id": networkId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-networks=v1+json"
    var result *models.ReadVMwareVCenterNetworkResponse
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
