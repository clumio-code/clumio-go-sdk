// Copyright (c) 2021 Clumio All Rights Reserved

package users

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// UsersV2Client represents a custom type interface
type UsersV2Client interface {
    // ListUsers Returns a list of Clumio users.
    ListUsers(
        limit *int64, 
        start *string, 
        filter *string)(
        *models.ListUsersResponse,  *apiutils.APIError)
    
    // CreateUser Creates a new user. Specify the user's full name and email address to generate an email message that is sent to the user with an invitation to activate their Clumio account.
    CreateUser(
        body *models.CreateUserV2Request)(
        *models.CreateUserResponse,  *apiutils.APIError)
    
    // ChangePassword Change the password of the current user. Users can only change their own passwords.
    ChangePassword(
        body *models.ChangePasswordV2Request)(
        *models.ChangePasswordResponse,  *apiutils.APIError)
    
    // UpdateUserProfile Manages the current user's profile, such as changing the user's full name.
    UpdateUserProfile(
        body *models.UpdateUserProfileV2Request)(
        *models.EditProfileResponse,  *apiutils.APIError)
    
    // ReadUser Returns a representation of the specified Clumio user.
    ReadUser(
        userId int64)(
        *models.ReadUserResponse,  *apiutils.APIError)
    
    // DeleteUser Deletes an existing user from Clumio, revoking the user's access to Clumio. A deleted user cannot be recovered.
    DeleteUser(
        userId int64)(
        interface{},  *apiutils.APIError)
    
    // UpdateUser Manages an existing user. Managing a user includes enabling or disabling the user,
    //  changing the user's full name or updating the user's access control.
    UpdateUser(
        userId int64, 
        body *models.UpdateUserV2Request)(
        *models.UpdateUserResponse,  *apiutils.APIError)
    
}

// NewUsersV2 returns UsersV2Client
func NewUsersV2(config config.Config) UsersV2Client {
    client := new(UsersV2)
    client.config = config
    return client
}
