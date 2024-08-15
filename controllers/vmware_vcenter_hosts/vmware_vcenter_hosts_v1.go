// Copyright (c) 2023 Clumio All Rights Reserved

// Package vmwarevcenterhosts contains methods related to VmwareVcenterHosts
package vmwarevcenterhosts

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// VmwareVcenterHostsV1 represents a custom type struct
type VmwareVcenterHostsV1 struct {
    config config.Config
}

// ListVmwareVcenterHosts Returns a list of hosts in the specified vCenter server.
func (v *VmwareVcenterHostsV1) ListVmwareVcenterHosts(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListHostsResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/hosts"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-hosts=v1+json"
    result := &models.ListHostsResponse{}
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


// ReadVmwareVcenterHost Returns a representation of the specified host.
func (v *VmwareVcenterHostsV1) ReadVmwareVcenterHost(
    vcenterId string, 
    hostId string)(
    *models.ReadHostResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/hosts/{host_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "host_id": hostId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-hosts=v1+json"
    result := &models.ReadHostResponse{}

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
