// Copyright (c) 2023 Clumio All Rights Reserved

// Package restoredawsdocumentdb contains methods related to RestoredAwsDocumentdb
package restoredawsdocumentdb

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredAwsDocumentdbV1 represents a custom type struct
type RestoredAwsDocumentdbV1 struct {
    config config.Config
}

// RestoreAwsDocumentdb Restores the specified DocumentDB cluster backup or snapshot to the specified target destination.
func (r *RestoredAwsDocumentdbV1) RestoreAwsDocumentdb(
    embed *string, 
    body models.RestoreAwsDocumentdbV1Request)(
    *models.CreateDocumentDBRestoreResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/aws/documentdb"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-aws-documentdb=v1+json"
    result := &models.CreateDocumentDBRestoreResponse{}
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
