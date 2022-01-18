// Copyright (c) 2021 Clumio All Rights Reserved

// Package roles contains methods related to Roles
package roles

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// RolesV1 represents a custom type struct
type RolesV1 struct {
    config config.Config
}

//  ListRoles Returns a list of roles that can be assigned to users, either while inviting users using the
//  [POST /users](#operation/create-user) API, or by updating the user using the
//  [PATCH /users/{user_id}](#operation/update-user) API.
func (r *RolesV1) ListRoles()(
    *models.ListRolesResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := r.config.BaseUrl + "/roles"

    
    header := "application/roles=v1+json"
    var result *models.ListRolesResponse
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetAuthToken(r.config.Token).
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


//  ReadRole Returns a representation of the specified role.
func (r *RolesV1) ReadRole(
    roleId string)(
    *models.ReadRoleResponse, *apiutils.APIError){

    var err error = nil
    pathURL := "/roles/{role_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "role_id": roleId,
    }
    queryBuilder := r.config.BaseUrl + pathURL

    
    header := "application/roles=v1+json"
    var result *models.ReadRoleResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(r.config.Token).
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
