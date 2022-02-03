// Copyright (c) 2021 Clumio All Rights Reserved

// Package protectiongroups contains methods related to ProtectionGroups
package protectiongroups

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// ProtectionGroupsV1 represents a custom type struct
type ProtectionGroupsV1 struct {
    config config.Config
}

//  ListProtectionGroups Returns a list of protection groups.
func (p *ProtectionGroupsV1) ListProtectionGroups(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListProtectionGroupsResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := p.config.BaseUrl + "/datasources/protection-groups"

    
    header := "application/protection-groups=v1+json"
    var result *models.ListProtectionGroupsResponse
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
    

    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "filter": *filter,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, p.config.OrganizationalUnitContext).
        SetAuthToken(p.config.Token).
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


//  ReadProtectionGroup Returns a representation of the specified protection group.
func (p *ProtectionGroupsV1) ReadProtectionGroup(
    groupId string)(
    *models.ReadProtectionGroupResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/datasources/protection-groups/{group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "group_id": groupId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    
    header := "application/protection-groups=v1+json"
    var result *models.ReadProtectionGroupResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, p.config.OrganizationalUnitContext).
        SetAuthToken(p.config.Token).
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


//  CreateProtectionGroup Creates a new protection group by specifying object filters. Appearance in
//  datasources/protection-groups read/listing is asynchronous and may take a few
//  seconds to minutes at most. As a result, the protection group won't be protectable
//  via /policies/assignments until it appears in the /datasources/protection-groups
//  endpoint.
func (p *ProtectionGroupsV1) CreateProtectionGroup(
    body models.CreateProtectionGroupV1Request)(
    *models.CreateProtectionGroupResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := p.config.BaseUrl + "/protection-groups"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/protection-groups=v1+json"
    var result *models.CreateProtectionGroupResponse
    client := resty.New()

    res, err := client.R().
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, p.config.OrganizationalUnitContext).
        SetAuthToken(p.config.Token).
        SetBody(payload).
        SetResult(&result).
        Post(queryBuilder)

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


//  UpdateProtectionGroup Updates a protection group by modifying object filters. Appearance in
//  datasources/protection-groups read/listing is asynchronous and may take a few
//  seconds to minutes at most. Must be in the same OU context as the creator of this
//  protection group.
func (p *ProtectionGroupsV1) UpdateProtectionGroup(
    groupId string, 
    body *models.UpdateProtectionGroupV1Request)(
    *models.UpdateProtectionGroupResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/protection-groups/{group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "group_id": groupId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/protection-groups=v1+json"
    var result *models.UpdateProtectionGroupResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, p.config.OrganizationalUnitContext).
        SetAuthToken(p.config.Token).
        SetBody(payload).
        SetResult(&result).
        Put(queryBuilder)

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


//  DeleteProtectionGroup Marks a protection group as deleted by taking the protection group ID. Appearance
//  in /datasources/protection-groups read/listing is asynchronous and may take a few
//  seconds to minutes at most. Must be in the same OU context as the creator of this
//  protection group.
func (p *ProtectionGroupsV1) DeleteProtectionGroup(
    groupId string)(
    interface{}, *apiutils.APIError){

    var err error = nil
    pathURL := "/protection-groups/{group_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "group_id": groupId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    
    header := "application/protection-groups=v1+json"
    var result interface{}
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, p.config.OrganizationalUnitContext).
        SetAuthToken(p.config.Token).
        SetResult(&result).
        Delete(queryBuilder)

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


//  AddBucketProtectionGroup Adds a bucket to protection group and creates a child protection group S3 asset.
//  Appearance in /datasources/protection-groups/s3-assets read/listing is asynchronous
//  and may take a few seconds to minutes at most. Must be in the same OU context as
//  the creator of this protection group. Bucket ID in body can be found in
//  datasources/aws/s3.
func (p *ProtectionGroupsV1) AddBucketProtectionGroup(
    groupId string, 
    body models.AddBucketProtectionGroupV1Request)(
    *models.AddBucketToProtectionGroupResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/protection-groups/{group_id}/buckets"
    //process optional template parameters
    pathParams := map[string]string{
        "group_id": groupId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/protection-groups=v1+json"
    var result *models.AddBucketToProtectionGroupResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, p.config.OrganizationalUnitContext).
        SetAuthToken(p.config.Token).
        SetBody(payload).
        SetResult(&result).
        Post(queryBuilder)

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


//  DeleteBucketProtectionGroup Deletes a bucket from a protection group and the child protection group S3 asset.
//  Appearance in /datasources/protection-groups/s3-assets read/listing is asynchronous
//  and may take a few seconds to minutes at most. Must be in the same OU context as
//  the creator of this protection group. Bucket ID in path can be found in
//  datasources/aws/s3.
func (p *ProtectionGroupsV1) DeleteBucketProtectionGroup(
    groupId string, 
    bucketId string)(
    *models.DeleteBucketFromProtectionGroupResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/protection-groups/{group_id}/buckets/{bucket_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "group_id": groupId,
        "bucket_id": bucketId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    
    header := "application/protection-groups=v1+json"
    var result *models.DeleteBucketFromProtectionGroupResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader(common.AcceptHeader, header).
        SetHeader(common.OrgUnitContextHeader, p.config.OrganizationalUnitContext).
        SetAuthToken(p.config.Token).
        SetResult(&result).
        Delete(queryBuilder)

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
