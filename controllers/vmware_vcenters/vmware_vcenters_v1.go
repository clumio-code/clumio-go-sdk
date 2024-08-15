// Copyright (c) 2023 Clumio All Rights Reserved

// Package vmwarevcenters contains methods related to VmwareVcenters
package vmwarevcenters

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// VmwareVcentersV1 represents a custom type struct
type VmwareVcentersV1 struct {
    config config.Config
}

// ListVmwareVcenters Returns a list of vCenter servers.
func (v *VmwareVcentersV1) ListVmwareVcenters(
    limit *int64, 
    start *string, 
    embed *string)(
    *models.ListVcentersResponse, *apiutils.APIError) {

    queryBuilder := v.config.BaseUrl + "/datasources/vmware/vcenters"

    
    header := "application/api.clumio.vmware-vcenters=v1+json"
    result := &models.ListVcentersResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
    }
    if embed != nil {
        queryParams["embed"] = *embed
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: v.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// ReadVmwareVcenter Returns a representation of the specified vCenter server.
func (v *VmwareVcentersV1) ReadVmwareVcenter(
    vcenterId string, 
    embed *string)(
    *models.ReadVcenterResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenters=v1+json"
    result := &models.ReadVcenterResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
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
