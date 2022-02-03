// Copyright (c) 2021 Clumio All Rights Reserved

// Package vmwarevcenterfolders contains methods related to VmwareVcenterFolders
package vmwarevcenterfolders

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// VmwareVcenterFoldersV1 represents a custom type struct
type VmwareVcenterFoldersV1 struct {
    config config.Config
}

//  ListVmwareVcenterFolders Returns a list of VMware folders in the specified vCenter server.
//  
//  The following table lists the supported Clumio folder types:
//  
//  
//  +-------------------------+---------------------------+
//  |       Folder Type       |        Description        |
//  +=========================+===========================+
//  | compute_resource_folder | Compute resource folders. |
//  +-------------------------+---------------------------+
//  | datacenter_folder       | Data center folders.      |
//  +-------------------------+---------------------------+
//  | vm_folder               | Virtual machine folders.  |
//  +-------------------------+---------------------------+
//  
//  
func (v *VmwareVcenterFoldersV1) ListVmwareVcenterFolders(
    vcenterId string, 
    limit *int64, 
    start *string, 
    filter *string, 
    embed *string)(
    *models.ListFoldersResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/folders"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/vmware-vcenter-folders=v1+json"
    var result *models.ListFoldersResponse
    client := resty.New()
    defaultInt64 := int64(0)
    defaultString := "" 
    

    if limit == nil{
        limit = &defaultInt64
    }
    if start == nil{
        start = &defaultString
    }
    if filter == nil{
        filter = &defaultString
    }
    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
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


//  ReadVmwareVcenterFolder Returns a representation of the specified VMware folder.
func (v *VmwareVcenterFoldersV1) ReadVmwareVcenterFolder(
    vcenterId string, 
    folderId string, 
    embed *string)(
    *models.ReadFolderResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/folders/{folder_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "folder_id": folderId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/vmware-vcenter-folders=v1+json"
    var result *models.ReadFolderResponse
    client := resty.New()
    defaultString := "" 
    

    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
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
