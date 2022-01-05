// Copyright (c) 2021 Clumio All Rights Reserved

package policyrules

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// PolicyRulesV1Client represents a custom type interface
type PolicyRulesV1Client interface {
    //  Creates a new policy rule. Policy rules determine how a policy should be assigned to assets.
    CreatePolicyRule(
        body *models.CreatePolicyRuleV1Request)(
        *models.CreateRuleResponse,  *apiutils.APIError)
    
    //  Returns a representation of the specified policy rule.
    ReadPolicyRule(
        ruleId string)(
        *models.ReadRuleResponse,  *apiutils.APIError)
    
    //  Updates an existing policy rule.
    UpdatePolicyRule(
        ruleId string, 
        body *models.UpdatePolicyRuleV1Request)(
        *models.UpdateRuleResponse,  *apiutils.APIError)
    
    //  Deletes the specified policy rule.
    DeletePolicyRule(
        ruleId string, 
        body *models.DeletePolicyRuleV1Request)(
        *models.DeleteRuleResponse,  *apiutils.APIError)
    
}

// NewPolicyRulesV1 returns PolicyRulesV1Client
func NewPolicyRulesV1(config config.Config) PolicyRulesV1Client{
    client := new(PolicyRulesV1)
    client.config = config
    return client
}
