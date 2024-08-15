// Copyright (c) 2023 Clumio All Rights Reserved

// Package vmwarevcenterfolders contains methods related to VmwareVcenterFolders
package vmwarevcenterfolders

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// VmwareVcenterFoldersV1 represents a custom type struct
type VmwareVcenterFoldersV1 struct {
    config config.Config
}

// ListVmwareVcenterFolders Returns a list of VMware folders in the specified vCenter server.
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
    *models.ListFoldersResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/folders"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-folders=v1+json"
    result := &models.ListFoldersResponse{}
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


// ReadVmwareVcenterFolder Returns a representation of the specified VMware folder.
func (v *VmwareVcenterFoldersV1) ReadVmwareVcenterFolder(
    vcenterId string, 
    folderId string, 
    embed *string)(
    *models.ReadFolderResponse, *apiutils.APIError) {

    pathURL := "/datasources/vmware/vcenters/{vcenter_id}/folders/{folder_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "vcenter_id": vcenterId,
        "folder_id": folderId,
    }
    queryBuilder := v.config.BaseUrl + pathURL

    
    header := "application/api.clumio.vmware-vcenter-folders=v1+json"
    result := &models.ReadFolderResponse{}
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
