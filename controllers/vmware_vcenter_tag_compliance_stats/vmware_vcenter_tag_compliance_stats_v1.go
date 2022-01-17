// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcentertagcompliancestats contains methods related to VmwareVcenterTagComplianceStats
package vmwarevcentertagcompliancestats

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterTagComplianceStatsV1 represents a custom type struct
type VmwareVcenterTagComplianceStatsV1 struct {
    config config.Config
}

//  ReadVmwareVcenterTagComplianceStats Returns the compliance statistics of the specified tag.
func (v *VmwareVcenterTagComplianceStatsV1) ReadVmwareVcenterTagComplianceStats(
    vcenterId string, 
    tagId string)(
    *models.ReadVMwareTagStatsResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/tags/{tag_id}/stats/compliance"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "tag_id": tagId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/vmware-vcenter-tag-compliance-stats=v1+json"
    var result *models.ReadVMwareTagStatsResponse
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
