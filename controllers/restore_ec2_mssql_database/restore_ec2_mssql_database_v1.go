// Copyright (c) 2021 Clumio All Rights Reserved

// Package restoreec2mssqldatabase contains methods related to RestoreEc2MssqlDatabase
package restoreec2mssqldatabase

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// RestoreEc2MssqlDatabaseV1 represents a custom type struct
type RestoreEc2MssqlDatabaseV1 struct {
    config config.Config
}

// RestoreEc2MssqlDatabase Restores an EC2 MSSQL database from a given backup or to a specified point in time.
func (r *RestoreEc2MssqlDatabaseV1) RestoreEc2MssqlDatabase(
    embed *string, 
    body models.RestoreEc2MssqlDatabaseV1Request)(
    *models.CreateEC2MSSQLDatabaseRestoreResponse, *apiutils.APIError) {

    queryBuilder := r.config.BaseUrl + "/restores/aws/ec2-mssql/databases"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.restore-ec2-mssql-database=v1+json"
    result := &models.CreateEC2MSSQLDatabaseRestoreResponse{}
    defaultString := "" 
    
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
        "embed": *embed,
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
