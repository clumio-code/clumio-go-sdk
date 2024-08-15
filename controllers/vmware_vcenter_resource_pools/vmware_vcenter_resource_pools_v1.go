// Copyright (c) 2023 Clumio All Rights Reserved

// Package vmwarevcenterresourcepools contains methods related to VmwareVcenterResourcePools
package vmwarevcenterresourcepools

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// VmwareVcenterResourcePoolsV1 represents a custom type struct
type VmwareVcenterResourcePoolsV1 struct {
    config config.Config
}

// ListVmwareVcenterResourcePools Returns a list of resource pools in the specified vCenter server.
func (v *VmwareVcenterResourcePoolsV1) ListVmwareVcenterResourcePools(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListResourcePoolsResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/resource-pools"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-resource-pools=v1+json"
    result := &models.ListResourcePoolsResponse{}
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


// ReadVmwareVcenterResourcePool Returns a representation of the specified resource pool.
func (v *VmwareVcenterResourcePoolsV1) ReadVmwareVcenterResourcePool(
    vcenterId string, 
    resourcePoolId string)(
    *models.ReadResourcePoolResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/resource-pools/{resource_pool_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "resource_pool_id": resourcePoolId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-resource-pools=v1+json"
    result := &models.ReadResourcePoolResponse{}

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
