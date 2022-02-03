// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcentercompliancestats contains methods related to VmwareVcenterComplianceStats
package vmwarevcentercompliancestats

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterComplianceStatsV1 represents a custom type struct
type VmwareVcenterComplianceStatsV1 struct {
    config config.Config
}

//  ReadVmwareVcenterComplianceStats Returns a representation of the compliance statistics of VMs in the specified vCenter server.
func (v *VmwareVcenterComplianceStatsV1) ReadVmwareVcenterComplianceStats(
    vcenterId string)(
    *models.ReadVMwareVCenterProtectionStatsResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/stats/compliance"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/vmware-vcenter-compliance-stats=v1+json"
    var result *models.ReadVMwareVCenterProtectionStatsResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, v.config.OrganizationalUnitContext).
        SetAuthToken(v.config.Token).
        SetResult(&result).
        Get(queryBuilder)

    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       common.InternalServerError,
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    if !res.IsSuccess(){
        return nil, &apiutils.APIError{
            ResponseCode: res.RawResponse.StatusCode,
            Reason:       common.NonSuccessStatusCodeError,
            Response:     res.Body(),
        }
    }
    return result, nil
}
