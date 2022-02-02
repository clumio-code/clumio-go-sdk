// Copyright (c) 2021 Clumio All Rights Reserved

// Package generalsettings contains methods related to GeneralSettings
package generalsettings

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// GeneralSettingsV2 represents a custom type struct
type GeneralSettingsV2 struct {
    config config.Config
}

//  ReadGeneralSettings Retrieves organization-wide setting details, including password and security settings.
func (g *GeneralSettingsV2) ReadGeneralSettings()(
    *models.ReadGeneralSettingsResponseV2, *apiutils.APIError){

    var err error = nil
    queryBuilder := g.config.BaseUrl + "/settings/general"

    
    header := "application/general-settings=v2+json"
    var result *models.ReadGeneralSettingsResponseV2
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetHeader("x-clumio-organizationalunit-context", g.config.OrganizationalUnitContext).
        SetAuthToken(g.config.Token).
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


//  UpdateGeneralSettings Updates organization-wide settings, including password and security settings.
func (g *GeneralSettingsV2) UpdateGeneralSettings(
    body *models.UpdateGeneralSettingsV2Request)(
    *models.PatchGeneralSettingsResponseV2, *apiutils.APIError){

    var err error = nil
    queryBuilder := g.config.BaseUrl + "/settings/general"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/general-settings=v2+json"
    var result *models.PatchGeneralSettingsResponseV2
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetHeader("x-clumio-organizationalunit-context", g.config.OrganizationalUnitContext).
        SetAuthToken(g.config.Token).
        SetBody(payload).
        SetResult(&result).
        Patch(queryBuilder)

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
