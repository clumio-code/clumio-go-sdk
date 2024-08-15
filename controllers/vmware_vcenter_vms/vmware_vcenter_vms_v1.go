// Copyright (c) 2023 Clumio All Rights Reserved

// Package vmwarevcentervms contains methods related to VmwareVcenterVms
package vmwarevcentervms

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// VmwareVcenterVmsV1 represents a custom type struct
type VmwareVcenterVmsV1 struct {
    config config.Config
}

// ListVmwareVcenterVms Returns a list of VMs in the specified vCenter server.
func (v *VmwareVcenterVmsV1) ListVmwareVcenterVms(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListVmsResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/vms"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-vms=v1+json"
    result := &models.ListVmsResponse{}
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


// ReadVmwareVcenterVm Returns a representation of the specified VM.
func (v *VmwareVcenterVmsV1) ReadVmwareVcenterVm(
    vcenterId string, 
    vmId string, 
    embed *string)(
    *models.ReadVmResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/vms/{vm_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "vm_id": vmId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-vms=v1+json"
    result := &models.ReadVmResponse{}
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
