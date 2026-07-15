// Copyright (c) 2023 Clumio All Rights Reserved

// Package restoredawsneptune contains methods related to RestoredAwsNeptune
package restoredawsneptune

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredAwsNeptuneV1 represents a custom type struct
type RestoredAwsNeptuneV1 struct {
    config config.Config
}

// RestoreAwsNeptune Restores the specified Neptune cluster backup or snapshot to the specified target destination.
func (r *RestoredAwsNeptuneV1) RestoreAwsNeptune(
    embed *string, 
    body models.RestoreAwsNeptuneV1Request)(
    *models.CreateNeptuneRestoreResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/aws/neptune"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-aws-neptune=v1+json"
    result := &models.CreateNeptuneRestoreResponse{}
    queryParams := make(map[string]string)
    if embed != nil {
        queryParams["embed"] = *embed
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}
