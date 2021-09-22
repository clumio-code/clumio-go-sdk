// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcenterdatastores contains methods related to VmwareVcenterDatastores
package vmwarevcenterdatastores

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterDatastoresV1 represents a custom type struct
type VmwareVcenterDatastoresV1 struct {
    config config.Config
}

//  ListVmwareVcenterDatastores Returns a list of datastores in the specified vCenter server.
//  
//  
//  Supported Datastore Types
//  Clumio supports the following VMware datastore types:
//  
//  Network File System (nfs)
//  vSAN (virtual_storage_area_network)
//  vSphere Virtual Machine File System (vmfs)
//  Virtual Volumes (virtual_volume)
//  
func (v *VmwareVcenterDatastoresV1) ListVmwareVcenterDatastores(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListVMwareDatastoresResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/datastores"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-datastores=v1+json"
    var result *models.ListVMwareDatastoresResponse
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


//  ReadVmwareVcenterDatastore Returns a representation of the specified datastore.
func (v *VmwareVcenterDatastoresV1) ReadVmwareVcenterDatastore(
    vcenterId string, 
    datastoreId string)(
    *models.ReadVMwareDatastoreResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/datastores/{datastore_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "datastore_id": datastoreId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-datastores=v1+json"
    var result *models.ReadVMwareDatastoreResponse
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
