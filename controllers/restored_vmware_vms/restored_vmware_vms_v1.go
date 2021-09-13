// Copyright (c) 2021 Clumio All Rights Reserved

// Package restoredvmwarevms contains methods related to RestoredVmwareVms
package restoredvmwarevms

import (
    "encoding/json"
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// RestoredVmwareVmsV1 represents a custom type struct
type RestoredVmwareVmsV1 struct {
    config config.Config
}

//  RestoreVmwareVm Restores the specified source VM backup to the specified target destination. The source VM must be one that was backed up by Clumio.
func (r *RestoredVmwareVmsV1) RestoreVmwareVm(
    body models.RestoreVmwareVmV1Request)(
    *models.RestoreVMwareVMResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := r.config.BaseUrl + "/restores/vmware/vms"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/restored-vmware-vms=v1+json"
    var result *models.RestoreVMwareVMResponse
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
