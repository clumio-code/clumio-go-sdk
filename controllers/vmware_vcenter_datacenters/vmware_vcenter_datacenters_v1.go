// Copyright (c) 2023 Clumio All Rights Reserved

// Package vmwarevcenterdatacenters contains methods related to VmwareVcenterDatacenters
package vmwarevcenterdatacenters

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// VmwareVcenterDatacentersV1 represents a custom type struct
type VmwareVcenterDatacentersV1 struct {
    config config.Config
}

// ListVmwareVcenterDatacenters Returns a list of VMware data centers in the specified vCenter server.
func (v *VmwareVcenterDatacentersV1) ListVmwareVcenterDatacenters(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string, 
    Embed *string)(
    *models.ListDatacentersResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/datacenters"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-datacenters=v1+json"
    result := &models.ListDatacentersResponse{}
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
    if Embed != nil {
        queryParams["_embed"] = *Embed
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


// ReadVmwareVcenterDatacenter Returns a representation of the specified VMware data center within the specified vCenter server.
func (v *VmwareVcenterDatacentersV1) ReadVmwareVcenterDatacenter(
    vcenterId string, 
    datacenterId string, 
    embed *string)(
    *models.ReadDatacenterResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/datacenters/{datacenter_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "datacenter_id": datacenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-datacenters=v1+json"
    result := &models.ReadDatacenterResponse{}
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
