// Copyright (c) 2021 Clumio All Rights Reserved

// Package postprocessawsconnection contains methods related to PostProcessAwsConnection
package postprocessawsconnection

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// PostProcessAwsConnectionV1 represents a custom type struct
type PostProcessAwsConnectionV1 struct {
    config config.Config
}

//  PostProcessAwsConnection Performs post-processing after AWS Connection Create, Update or Delete. This API should only be invoked by the Clumio Terraform provider and should not be invoked manually.
func (p *PostProcessAwsConnectionV1) PostProcessAwsConnection(
    body *models.PostProcessAwsConnectionV1Request)(
    interface{}, *apiutils.APIError){

    var err error = nil
    _queryBuilder := p.config.BaseUrl + "/connections/aws/post-process"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/post-process-aws-connection=v1+json"
    var result interface{}
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetAuthToken(p.config.Token).
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
