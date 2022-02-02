// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcenterfoldercompliancestats contains methods related to VmwareVcenterFolderComplianceStats
package vmwarevcenterfoldercompliancestats

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterFolderComplianceStatsV1 represents a custom type struct
type VmwareVcenterFolderComplianceStatsV1 struct {
    config config.Config
}

//  ReadVmwareVcenterFolderComplianceStats Returns the compliance statistics of VMs under folders and subfolders of the specified VMware folder.
func (v *VmwareVcenterFolderComplianceStatsV1) ReadVmwareVcenterFolderComplianceStats(
    vcenterId string, 
    folderId string)(
    *models.ReadVMwareFolderStatsResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/folders/{folder_id}/stats/compliance"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "folder_id": folderId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/vmware-vcenter-folder-compliance-stats=v1+json"
    var result *models.ReadVMwareFolderStatsResponse
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
