// Copyright (c) 2023 Clumio All Rights Reserved

// Package vmwarevcentercomputeresources contains methods related to VmwareVcenterComputeResources
package vmwarevcentercomputeresources

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// VmwareVcenterComputeResourcesV1 represents a custom type struct
type VmwareVcenterComputeResourcesV1 struct {
    config config.Config
}

// ListVmwareVcenterComputeResources Returns a list of VMware compute resources in the specified vCenter server.
//  
//  The following table lists the supported Clumio protection statuses:
//  
//  
//  +-------------------+-------------------------------------------------------+
//  | Protection Status |                        Values                         |
//  +===================+=======================================================+
//  | protected         | A compute resource that is protected by a policy.     |
//  +-------------------+-------------------------------------------------------+
//  | unprotected       | A compute resource that is not protected by a policy. |
//  +-------------------+-------------------------------------------------------+
//  
func (v *VmwareVcenterComputeResourcesV1) ListVmwareVcenterComputeResources(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListComputeResourcesResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/compute-resources"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-compute-resources=v1+json"
    result := &models.ListComputeResourcesResponse{}
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


// ReadVmwareVcenterComputeResource Returns a representation of the specified VMware compute resource.
func (v *VmwareVcenterComputeResourcesV1) ReadVmwareVcenterComputeResource(
    vcenterId string, 
    computeResourceId string, 
    embed *string)(
    *models.ReadComputeResourceResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/compute-resources/{compute_resource_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "compute_resource_id": computeResourceId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-compute-resources=v1+json"
    result := &models.ReadComputeResourceResponse{}
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
