// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcenterdatacentercompliancestats contains methods related to VmwareVcenterDatacenterComplianceStats
package vmwarevcenterdatacentercompliancestats

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterDatacenterComplianceStatsV1 represents a custom type struct
type VmwareVcenterDatacenterComplianceStatsV1 struct {
    config config.Config
}

//  ReadVmwareVcenterDatacenterComplianceStats Returns a representation of the compliance statistics of the specified data center.
func (v *VmwareVcenterDatacenterComplianceStatsV1) ReadVmwareVcenterDatacenterComplianceStats(
    vcenterId string, 
    datacenterId string)(
    *models.ReadVMwareDatacenterStatsResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/datacenters/{datacenter_id}/stats/compliance"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "datacenter_id": datacenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/vmware-vcenter-datacenter-compliance-stats=v1+json"
    var result *models.ReadVMwareDatacenterStatsResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetHeader("x-clumio-organizationalunit-context", v.config.OrganizationalUnitContext).
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
