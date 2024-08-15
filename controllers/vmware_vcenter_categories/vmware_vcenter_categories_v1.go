// Copyright (c) 2023 Clumio All Rights Reserved

// Package vmwarevcentercategories contains methods related to VmwareVcenterCategories
package vmwarevcentercategories

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// VmwareVcenterCategoriesV1 represents a custom type struct
type VmwareVcenterCategoriesV1 struct {
    config config.Config
}

// ListVmwareVcenterCategories Returns a list of tag categories in the specified vCenter server.
func (v *VmwareVcenterCategoriesV1) ListVmwareVcenterCategories(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListTagCategories2Response, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/categories"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-categories=v1+json"
    result := &models.ListTagCategories2Response{}
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
        Config: v.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadVmwareVcenterCategory Returns a representation of the specified tag category.
func (v *VmwareVcenterCategoriesV1) ReadVmwareVcenterCategory(
    vcenterId string, 
    categoryId string)(
    *models.ReadTagCategory2Response, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/categories/{category_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "category_id": categoryId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-categories=v1+json"
    result := &models.ReadTagCategory2Response{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: v.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
