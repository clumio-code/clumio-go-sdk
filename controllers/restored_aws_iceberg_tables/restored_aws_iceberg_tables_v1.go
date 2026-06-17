// Copyright (c) 2023 Clumio All Rights Reserved

// Package restoredawsicebergtables contains methods related to RestoredAwsIcebergTables
package restoredawsicebergtables

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoredAwsIcebergTablesV1 represents a custom type struct
type RestoredAwsIcebergTablesV1 struct {
    config config.Config
}

// RestoreAwsIcebergTable Restores the specified source AWS Iceberg Table snapshots to the specified target destination. The source AWS Iceberg Table must be one that was backed up by Clumio.
func (r *RestoredAwsIcebergTablesV1) RestoreAwsIcebergTable(
    body models.RestoreAwsIcebergTableV1Request)(
    *models.RestoreAWSIcebergTableResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/aws/iceberg-tables"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restored-aws-iceberg-tables=v1+json"
    result := &models.RestoreAWSIcebergTableResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: r.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}
