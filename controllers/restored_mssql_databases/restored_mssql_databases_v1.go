// Copyright (c) 2021 Clumio All Rights Reserved

// Package restoredmssqldatabases contains methods related to RestoredMssqlDatabases
package restoredmssqldatabases

import (
    "encoding/json"
    "fmt"

    "github.com/clumio/clumiogosdk/api_utils"
    "github.com/clumio/clumiogosdk/config"
    "github.com/clumio/clumiogosdk/models"
    "github.com/go-resty/resty/v2"
)

// RestoredMssqlDatabasesV1 represents a custom type struct
type RestoredMssqlDatabasesV1 struct {
    config config.Config
}

//  RestoreMssqlDatabase Creates a restored MSSQL database from a given backup or to a specified point in time.
func (r *RestoredMssqlDatabasesV1) RestoreMssqlDatabase(
    embed *string, 
    body models.RestoreMssqlDatabaseV1Request)(
    *models.CreateMssqlDatabaseRestoreResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := r.config.BaseUrl + "/restores/mssql/databases"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/restored-mssql-databases=v1+json"
    var result *models.CreateMssqlDatabaseRestoreResponse
    client := resty.New()
    defaultString := "" 
    

    if embed == nil{
        embed = &defaultString
    }
    

    queryParams := map[string]string{
        "embed": *embed,
    }

    res, err := client.R().
        SetQueryParams(queryParams).
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
