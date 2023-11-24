// Copyright (c) 2023 Clumio All Rights Reserved

// Package backupawsec2instances contains methods related to BackupAwsEc2Instances
package backupawsec2instances

import (
    "encoding/json"
    "fmt"

    "github.com/clumio-code/clumio-go-sdk/api_utils"
    "github.com/clumio-code/clumio-go-sdk/common"
    "github.com/clumio-code/clumio-go-sdk/config"
    "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsEc2InstancesV1 represents a custom type struct
type BackupAwsEc2InstancesV1 struct {
    config config.Config
}

// ListBackupAwsEc2Instances Returns a list of EC2 instances that have been backed up by Clumio. EC2 instance backups can be restored through the [POST /restores/aws/ec2-instances](#operation/restore-aws-ec2-instance) endpoint.
func (b *BackupAwsEc2InstancesV1) ListBackupAwsEc2Instances(
    limit *int64, 
    start *string, 
    sort *string, 
    filter *string)(
    *models.ListEC2BackupsResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/ec2-instances"

    
    header := "application/api.clumio.backup-aws-ec2-instances=v1+json"
    result := &models.ListEC2BackupsResponse{}
    defaultInt64 := int64(0)
    defaultString := "" 
    
    if limit == nil {
        limit = &defaultInt64
    }
    if start == nil {
        start = &defaultString
    }
    if sort == nil {
        sort = &defaultString
    }
    if filter == nil {
        filter = &defaultString
    }
    
    queryParams := map[string]string{
        "limit": fmt.Sprintf("%v", *limit),
        "start": *start,
        "sort": *sort,
        "filter": *filter,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}


// CreateBackupAwsEc2Instance Performs an on-demand backup for the specified EC2 instance.
func (b *BackupAwsEc2InstancesV1) CreateBackupAwsEc2Instance(
    embed *string, 
    body models.CreateBackupAwsEc2InstanceV1Request)(
    *models.OnDemandEC2BackupResponse, *apiutils.APIError) {

    queryBuilder := b.config.BaseUrl + "/backups/aws/ec2-instances"

    bytes, err := json.Marshal(body)
    if err != nil {
        return nil, &apiutils.APIError{
            ResponseCode: 500,
            Reason:       fmt.Sprintf("Failed to Marshal Request Body %v", body),
            Response:     []byte(fmt.Sprintf("%v", err)),
        }
    }
    payload := string(bytes)
    header := "application/api.clumio.backup-aws-ec2-instances=v1+json"
    result := &models.OnDemandEC2BackupResponse{}
    defaultString := "" 
    
    if embed == nil {
        embed = &defaultString
    }
    
    queryParams := map[string]string{
        "embed": *embed,
    }

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        QueryParams: queryParams,
        AcceptHeader: header,
        Body: payload,
        Result202: &result,
        RequestType: common.Post,
    })

    return result, apiErr
}


// ReadBackupAwsEc2Instance Returns a representation of the specified EC2 instance backup.
func (b *BackupAwsEc2InstancesV1) ReadBackupAwsEc2Instance(
    backupId string)(
    *models.ReadEC2BackupResponse, *apiutils.APIError) {

    pathURL := "/backups/aws/ec2-instances/{backup_id}"
    //process optional template parameters
    pathParams := map[string]string{
        "backup_id": backupId,
    }
    queryBuilder := b.config.BaseUrl + pathURL

    
    header := "application/api.clumio.backup-aws-ec2-instances=v1+json"
    result := &models.ReadEC2BackupResponse{}

    apiErr := common.InvokeAPI(&common.InvokeAPIRequest{
        Config: b.config,
        RequestUrl: queryBuilder,
        PathParams: pathParams,
        AcceptHeader: header,
        Result200: &result,
        RequestType: common.Get,
    })

    return result, apiErr
}
