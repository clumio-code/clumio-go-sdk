// Copyright (c) 2021 Clumio All Rights Reserved

// Package restoredawsebsvolumes contains methods related to RestoredAwsEbsVolumes
package restoredawsebsvolumes

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// RestoredAwsEbsVolumesV1 represents a custom type struct
type RestoredAwsEbsVolumesV1 struct {
    config config.Config
}

//  RestoreAwsEbsVolume Restores the specified source EBS volume backup to the specified target destination. The source EBS volume must be one that was backup up by Clumio.
func (r *RestoredAwsEbsVolumesV1) RestoreAwsEbsVolume(
    body models.RestoreAwsEbsVolumeV1Request)(
    interface{}, *apiutils.APIError){

    var err error = nil
    _queryBuilder := r.config.BaseUrl + "/restores/aws/ebs-volumes"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/restored-aws-ebs-volumes=v1+json"
    var result interface{}
    client := resty.New()

    res, err := client.R().
        SetHeader("Accept", header).
        SetAuthToken(r.config.Token).
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
