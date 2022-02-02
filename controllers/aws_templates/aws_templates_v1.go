// Copyright (c) 2021 Clumio All Rights Reserved

// Package awstemplates contains methods related to AwsTemplates
package awstemplates

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// AwsTemplatesV1 represents a custom type struct
type AwsTemplatesV1 struct {
    config config.Config
}

//  ReadConnectionTemplates Returns the AWS CloudFormation and Terraform templates available to install to connect
//  to Clumio.
func (a *AwsTemplatesV1) ReadConnectionTemplates()(
    *models.ReadAWSTemplatesV2Response, *apiutils.APIError){

    var err error = nil
    queryBuilder := a.config.BaseUrl + "/connections/aws/templates"

    
    header := "application/aws-templates=v1+json"
    var result *models.ReadAWSTemplatesV2Response
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetHeader("x-clumio-organizationalunit-context", a.config.OrganizationalUnitContext).
        SetAuthToken(a.config.Token).
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


//  CreateConnectionTemplate Returns the URLs for AWS CloudFormation and terraform templates  corresponding
//  to a given configuration of asset types.
func (a *AwsTemplatesV1) CreateConnectionTemplate(
    body *models.CreateConnectionTemplateV1Request)(
    *models.CreateAWSTemplateV2Response, *apiutils.APIError){

    var err error = nil
    queryBuilder := a.config.BaseUrl + "/connections/aws/templates"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/aws-templates=v1+json"
    var result *models.CreateAWSTemplateV2Response
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetHeader("x-clumio-organizationalunit-context", a.config.OrganizationalUnitContext).
        SetAuthToken(a.config.Token).
        SetBody(payload).
        SetResult(&result).
        Post(queryBuilder)

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
