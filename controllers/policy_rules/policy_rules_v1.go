// Copyright (c) 2021 Clumio All Rights Reserved

// Package policyrules contains methods related to PolicyRules
package policyrules

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
    "github.com/go-resty/resty/v2"
)

// PolicyRulesV1 represents a custom type struct
type PolicyRulesV1 struct {
    config config.Config
}

//  CreatePolicyRule Creates a new policy rule. Policy rules determine how a policy should be assigned to assets.
func (p *PolicyRulesV1) CreatePolicyRule(
    body *models.CreatePolicyRuleV1Request)(
    *models.CreateRuleResponse, *apiutils.APIError){

    var err error = nil
    _queryBuilder := p.config.BaseUrl + "/policies/rules"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/policy-rules=v1+json"
    var result *models.CreateRuleResponse
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


//  ReadPolicyRule Returns a representation of the specified policy rule.
func (p *PolicyRulesV1) ReadPolicyRule(
    ruleId string)(
    *models.ReadRuleResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/policies/rules/{rule_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "rule_id": ruleId,
    }
    _queryBuilder := p.config.BaseUrl + _pathURL

    
    header := "application/policy-rules=v1+json"
    var result *models.ReadRuleResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(p.config.Token).
        SetResult(&result).
        Get(_queryBuilder)

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


//  UpdatePolicyRule Updates an existing policy rule.
func (p *PolicyRulesV1) UpdatePolicyRule(
    ruleId string, 
    body *models.UpdatePolicyRuleV1Request)(
    *models.UpdateRuleResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/policies/rules/{rule_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "rule_id": ruleId,
    }
    _queryBuilder := p.config.BaseUrl + _pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/policy-rules=v1+json"
    var result *models.UpdateRuleResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(p.config.Token).
        SetBody(payload).
        SetResult(&result).
        Put(_queryBuilder)

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


//  DeletePolicyRule Deletes the specified policy rule.
func (p *PolicyRulesV1) DeletePolicyRule(
    ruleId string, 
    body *models.DeletePolicyRuleV1Request)(
    *models.DeleteRuleResponse, *apiutils.APIError){

    var err error = nil
    _pathURL := "/policies/rules/{rule_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "rule_id": ruleId,
    }
    _queryBuilder := p.config.BaseUrl + _pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/policy-rules=v1+json"
    var result *models.DeleteRuleResponse
    client := resty.New()

    res, err := client.R().
        SetPathParams(pathParams).
        SetHeader("Accept", header).
        SetAuthToken(p.config.Token).
        SetBody(payload).
        SetResult(&result).
        Delete(_queryBuilder)

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
