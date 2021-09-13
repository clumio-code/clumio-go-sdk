// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcentercategories contains methods related to VmwareVcenterCategories
package vmwarevcentercategories

import (
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterCategoriesV1 represents a custom type struct
type VmwareVcenterCategoriesV1 struct {
    config config.Config
}

//  ListVmwareVcenterCategories Returns a list of tag categories in the specified vCenter server.
func (v *VmwareVcenterCategoriesV1) ListVmwareVcenterCategories(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListTagCategories2Response, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/categories"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenterId": vcenterId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-categories=v1+json"
    var result *models.ListTagCategories2Response
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


//  ReadVmwareVcenterCategory Returns a representation of the specified tag category.
func (v *VmwareVcenterCategoriesV1) ReadVmwareVcenterCategory(
    vcenterId string, 
    categoryId string)(
    *models.ReadTagCategory2Response, *apiutils.APIError){

    var err error = nil
    _pathURL := "/datasources/vmware/vcenters/{vcenter_id}/categories/{category_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenterId": vcenterId,
        "categoryId": categoryId,
    }
    _queryBuilder := v.config.BaseUrl + _pathURL

    
    header := "application/vmware-vcenter-categories=v1+json"
    var result *models.ReadTagCategory2Response
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
