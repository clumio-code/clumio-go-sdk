// Copyright (c) 2023 Clumio All Rights Reserved

// Package awstemplates contains methods related to AwsTemplates
package awstemplates

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// AwsTemplatesV1 represents a custom type struct
type AwsTemplatesV1 struct {
    config config.Config
}

// ReadConnectionTemplates Returns the AWS CloudFormation and Terraform templates available to install to connect
//  to Clumio.
func (a *AwsTemplatesV1) ReadConnectionTemplates()(
    *models.ReadAWSTemplatesV2Response, *apiutils.APIError) {

    queryBuilder := a.config.BaseUrl + "/connections/aws/templates"

    
    header := "application/api.clumio.aws-templates=v1+json"
    result := &models.ReadAWSTemplatesV2Response{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: a.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreateConnectionTemplate Returns the URLs for AWS CloudFormation and terraform templates  corresponding
//  to a given configuration of asset types.
func (a *AwsTemplatesV1) CreateConnectionTemplate(
    returnGroupToken *bool, 
    body *models.CreateConnectionTemplateV1Request)(
    *models.CreateAWSTemplateV2Response, *apiutils.APIError) {

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
    header := "application/api.clumio.aws-templates=v1+json"
    result := &models.CreateAWSTemplateV2Response{}
    queryParams := make(map[string]string)
    if returnGroupToken != nil {
        queryParams["return_group_token"] = fmt.Sprintf("%v", *returnGroupToken)
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: a.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}
