// Copyright (c) 2023 Clumio All Rights Reserved

// Package policyrules contains methods related to PolicyRules
package policyrules

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// PolicyRulesV1 represents a custom type struct
type PolicyRulesV1 struct {
    config config.Config
}

// ListPolicyRules Returns a list of policy rules.
func (p *PolicyRulesV1) ListPolicyRules(
    limit *int64, 
    start *string, 
    organizationalUnitId *string, 
    sort *string, 
    filter *string)(
    *models.ListRulesResponse, *apiutils.APIError) {

    queryBuilder := p.config.BaseUrl + "/policies/rules"

    
    header := "application/api.clumio.policy-rules=v1+json"
    result := &models.ListRulesResponse{}
    queryParams := make(map[string]string)
    if limit != nil {
        queryParams["limit"] = fmt.Sprintf("%v", *limit)
    }
    if start != nil {
        queryParams["start"] = *start
    }
    if organizationalUnitId != nil {
        queryParams["organizational_unit_id"] = *organizationalUnitId
    }
    if sort != nil {
        queryParams["sort"] = *sort
    }
    if filter != nil {
        queryParams["filter"] = *filter
    }
    

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: p.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreatePolicyRule Creates a new policy rule. Policy rules determine how a policy should be assigned to assets.
//  Additionally, to create a rule in the context of another Organizational Unit, refer to the
//  Getting Started documentation.
func (p *PolicyRulesV1) CreatePolicyRule(
    body *models.CreatePolicyRuleV1Request)(
    *models.CreateRuleResponse, *apiutils.APIError) {

    queryBuilder := p.config.BaseUrl + "/policies/rules"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.policy-rules=v1+json"
    result := &models.CreateRuleResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: p.config,
        RequestUrl: queryBuilder,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadPolicyRule Returns a representation of the specified policy rule.
func (p *PolicyRulesV1) ReadPolicyRule(
    ruleId string)(
    *models.ReadRuleResponse, *apiutils.APIError) {

    pathURL := "/policies/rules/{rule_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "rule_id": ruleId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    
    header := "application/api.clumio.policy-rules=v1+json"
    result := &models.ReadRuleResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: p.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// UpdatePolicyRule Updates an existing policy rule.
func (p *PolicyRulesV1) UpdatePolicyRule(
    ruleId string, 
    body *models.UpdatePolicyRuleV1Request)(
    *models.UpdateRuleResponse, *apiutils.APIError) {

    pathURL := "/policies/rules/{rule_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "rule_id": ruleId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.policy-rules=v1+json"
    result := &models.UpdateRuleResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: p.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Body: payload,
        Result200: &result,
        RequestType: common.Put,
    })

    return result, apiErr
}


// DeletePolicyRule Deletes the specified policy rule.
func (p *PolicyRulesV1) DeletePolicyRule(
    ruleId string)(
    *models.DeleteRuleResponse, *apiutils.APIError) {

    pathURL := "/policies/rules/{rule_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "rule_id": ruleId,
    }
    queryBuilder := p.config.BaseUrl + pathURL

    
    header := "application/api.clumio.policy-rules=v1+json"
    result := &models.DeleteRuleResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: p.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Delete,
    })

    return result, apiErr
}
