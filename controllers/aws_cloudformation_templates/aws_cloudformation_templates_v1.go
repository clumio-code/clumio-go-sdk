// Copyright (c) 2021 Clumio All Rights Reserved

// Package awscloudformationtemplates contains methods related to AwsCloudformationTemplates
package awscloudformationtemplates

import (
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// AwsCloudformationTemplatesV1 represents a custom type struct
type AwsCloudformationTemplatesV1 struct {
    config config.Config
}

//  ReadAwsConnectionTemplates Returns the AWS CloudFormation templates available to install to connect
//  to Clumio. The provided URL will be pasted into the Amazon S3 URL field when
//  creating the CloudFormation stack.
func (a *AwsCloudformationTemplatesV1) ReadAwsConnectionTemplates()(
    *models.ReadAWSTemplatesResponse, *apiutils.APIError){

    var err error = nil
    queryBuilder := a.config.BaseUrl + "/connections/aws/cloudformation-templates"

    
    header := "application/aws-cloudformation-templates=v1+json"
    var result *models.ReadAWSTemplatesResponse
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
