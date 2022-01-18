// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcentercomputeresourcecompliancestats contains methods related to VmwareVcenterComputeResourceComplianceStats
package vmwarevcentercomputeresourcecompliancestats

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterComputeResourceComplianceStatsV1 represents a custom type struct
type VmwareVcenterComputeResourceComplianceStatsV1 struct {
    config config.Config
}

//  ReadVmwareVcenterComputeResourceComplianceStats Returns a representation of the compliance statistics of the specified VMware compute resource.
func (v *VmwareVcenterComputeResourceComplianceStatsV1) ReadVmwareVcenterComputeResourceComplianceStats(
    vcenterId string, 
    computeResourceId string)(
    *models.ReadVMwareComputeResourceStatsResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/compute-resources/{compute_resource_id}/stats/compliance"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "compute_resource_id": computeResourceId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/vmware-vcenter-compute-resource-compliance-stats=v1+json"
    var result *models.ReadVMwareComputeResourceStatsResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(v.config.Token).
        SetResult(&result).
        Get(queryBuilder)

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
