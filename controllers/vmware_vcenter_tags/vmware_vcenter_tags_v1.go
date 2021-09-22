// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcentertags contains methods related to VmwareVcenterTags
package vmwarevcentertags

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterTagsV1 represents a custom type struct
type VmwareVcenterTagsV1 struct {
    config config.Config
}

//  ListVmwareVcenterTags Returns a list of tags in the specified vCenter server.
func (v *VmwareVcenterTagsV1) ListVmwareVcenterTags(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListTagsResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/tags"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-tags=v1+json"
    var result *models.ListTagsResponse
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
    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
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


//  ReadVmwareVcenterTag Returns a representation of the specified tag.
func (v *VmwareVcenterTagsV1) ReadVmwareVcenterTag(
    vcenterId string, 
    tagId string, 
    embed *string)(
    *models.ReadTagResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/tags/{tag_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "tag_id": tagId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-tags=v1+json"
    var result *models.ReadTagResponse
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
