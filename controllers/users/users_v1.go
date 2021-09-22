// Copyright (c) 2021 Clumio All Rights Reserved

// Package users contains methods related to Users
package users

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// UsersV1 represents a custom type struct
type UsersV1 struct {
    config config.Config
}

//  ListUsers Returns a list of Clumio users.
func (u *UsersV1) ListUsers(
    limit *int64, 
    start *string, 
    filter *string)(
    *models.ListUsersResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := u.config.BaseUrl + "/users"

    
    header := "application/users=v1+json"
    var result *models.ListUsersResponse
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
        SetHeader("Accept", header).
        SetAuthToken(u.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

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


//  CreateUser Creates a new user. Specify the user's full name and email address to generate an email message that is sent to the user with an invitation to activate their Clumio account.
func (u *UsersV1) CreateUser(
    body *models.CreateUserV1Request)(
    *models.CreateUserResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := u.config.BaseUrl + "/users"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/users=v1+json"
    var result *models.CreateUserResponse
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetAuthToken(u.config.Token).
        SetBody(payload).
        SetResult(&result).
        Post(_queryBuilder)

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


//  UpdateUserProfile Manages the current user's profile, such as changing the user's full name.
func (u *UsersV1) UpdateUserProfile(
    body *models.UpdateUserProfileV1Request)(
    *models.EditProfileResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := u.config.BaseUrl + "/users/my-profile"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/users=v1+json"
    var result *models.EditProfileResponse
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetAuthToken(u.config.Token).
        SetBody(payload).
        SetResult(&result).
        Patch(_queryBuilder)

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


//  ReadUser Returns a representation of the specified Clumio user.
func (u *UsersV1) ReadUser(
    userId int64)(
    *models.ReadUserResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/users/{user_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "user_id": fmt.Sprintf("%v", userId),
    }
    _queryBuilder := u.config.BaseUrl + _pathURL

    
    header := "application/users=v1+json"
    var result *models.ReadUserResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(u.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

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


//  DeleteUser Deletes an existing user from Clumio, revoking the user's access to Clumio. A deleted user cannot be recovered.
func (u *UsersV1) DeleteUser(
    userId int64)(
    interface{}, *apiutils.APIError){

    var err error = nil
    _pathURL := "/users/{user_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "user_id": fmt.Sprintf("%v", userId),
    }
    _queryBuilder := u.config.BaseUrl + _pathURL

    
    header := "application/users=v1+json"
    var result interface{}
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(u.config.Token).
        SetResult(&result).
        Delete(_queryBuilder)

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


//  UpdateUser Manages an existing user. Managing a user includes enabling or disabling the user, or changing the user's full name.
func (u *UsersV1) UpdateUser(
    userId int64, 
    body *models.UpdateUserV1Request)(
    *models.UpdateUserResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/users/{user_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "user_id": fmt.Sprintf("%v", userId),
    }
    _queryBuilder := u.config.BaseUrl + _pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/users=v1+json"
    var result *models.UpdateUserResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(u.config.Token).
        SetBody(payload).
        SetResult(&result).
        Patch(_queryBuilder)

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


//  ChangePassword Change the password of the specified user. Users can change their own passwords.
func (u *UsersV1) ChangePassword(
    userId int64, 
    body *models.ChangePasswordV1Request)(
    interface{}, *apiutils.APIError){

    var err error = nil
    _pathURL := "/users/{user_id}/password"
    //process optional template parameters
    pathParams := map[string]string{
        "user_id": fmt.Sprintf("%v", userId),
    }
    _queryBuilder := u.config.BaseUrl + _pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/users=v1+json"
    var result interface{}
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(u.config.Token).
        SetBody(payload).
        SetResult(&result).
        Put(_queryBuilder)

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
