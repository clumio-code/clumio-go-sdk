// Copyright (c) 2021 Clumio All Rights Reserved

package backupawsec2instances

import (
     "github.com/clumio-code/clumio-go-sdk/api_utils"
     "github.com/clumio-code/clumio-go-sdk/config"
     "github.com/clumio-code/clumio-go-sdk/models"
)

// BackupAwsEc2InstancesV1Client represents a custom type interface
type BackupAwsEc2InstancesV1Client interface {
    // ListBackupAwsEc2Instances Returns a list of EC2 instances that have been backed up by Clumio. EC2 instance backups can be restored through the [POST /restores/aws/ec2-instances](#operation/restore-aws-ec2-instance) endpoint.
    ListBackupAwsEc2Instances(
        limit *int64, 
        start *string, 
        sort *string, 
        filter *string)(
        *models.ListEC2BackupsResponse,  *apiutils.APIError)
    
    // CreateBackupAwsEc2Instance Performs an on-demand backup for the specified EC2 instance.
    CreateBackupAwsEc2Instance(
        embed *string, 
        body models.CreateBackupAwsEc2InstanceV1Request)(
        *models.OnDemandEC2BackupResponse,  *apiutils.APIError)
    
    // ReadBackupAwsEc2Instance Returns a representation of the specified EC2 instance backup.
    ReadBackupAwsEc2Instance(
        backupId string)(
        *models.ReadEC2BackupResponse,  *apiutils.APIError)
    
}

// NewBackupAwsEc2InstancesV1 returns BackupAwsEc2InstancesV1Client
func NewBackupAwsEc2InstancesV1(config config.Config) BackupAwsEc2InstancesV1Client {
    client := new(BackupAwsEc2InstancesV1)
    client.config = config
    return client
}
