// Copyright (c) 2021 Clumio All Rights Reserved

// Package models has the structs representing requests
package models

// UpdateConsolidatedAlertV1Request represents a custom type struct
type UpdateConsolidatedAlertV1Request struct {
    // A record of user-provided information about the alert. The note must be less than 1024 characters in length. Adding a new note overwrites any existing notes.
    Notes  *string `json:"notes"`
    // Manually clears an active alert. To clear the active alert, set the parameter to "cleared". Once an alert is cleared,
    // the status of the alert changes from "active" to "cleared".
    // If the alert is already in "cleared" status, this action is ignored.
    // An alert that is in "cleared" status cannot be changed to "active" status.
    Status *string `json:"status"`
}

// UpdateIndividualAlertV1Request represents a custom type struct
type UpdateIndividualAlertV1Request struct {
    // A record of information about the alert. The note must be less than 1024 characters in length. Adding a new note overwrites any existing notes.
    Notes  *string `json:"notes"`
    // Manually clears an active alert. To clear the alert, set to "cleared"`". Once an alert is cleared,
    // the status of the alert changes from "active" to "cleared".
    // If the alert is already in "cleared" status, this action is ignored.
    // An alert that is in "cleared" status cannot be changed to "active" status.
    Status *string `json:"status"`
}

// CreateBackupAwsDynamodbTableV1Request represents a custom type struct
type CreateBackupAwsDynamodbTableV1Request struct {
    // Settings for requesting on-demand backup directly.
    Settings   *OnDemandSetting `json:"settings"`
    // Performs the operation on the DynamoDB table with the specified Clumio-assigned ID.
    TableId    *string          `json:"table_id"`
    // The type of the backup. Possible values - `clumio_backup`, `aws_snapshot`. The
    // type will be assumed as `aws_snapshot` if the field is left empty.
    ClumioType *string          `json:"type"`
}

// CreateBackupAwsEbsVolumeV1Request represents a custom type struct
type CreateBackupAwsEbsVolumeV1Request struct {
    // Settings for requesting on-demand backup directly.
    Settings *OnDemandSetting `json:"settings"`
    // Performs the operation on the EBS volume with the specified Clumio-assigned ID.
    VolumeId *string          `json:"volume_id"`
}

// CreateBackupAwsEbsVolumeV2Request represents a custom type struct
type CreateBackupAwsEbsVolumeV2Request struct {
    // Settings for requesting on-demand backup directly.
    Settings   *OnDemandSetting `json:"settings"`
    // The type of the backup. Possible values - `clumio_backup`, `aws_snapshot`. The
    // type will be assumed as `clumio_backup` if the field is left empty.
    ClumioType *string          `json:"type"`
    // Performs the operation on the EBS volume with the specified Clumio-assigned ID.
    VolumeId   *string          `json:"volume_id"`
}

// CreateBackupAwsEc2InstanceV1Request represents a custom type struct
type CreateBackupAwsEc2InstanceV1Request struct {
    // Performs the operation on the EC2 instance with the specified Clumio-assigned ID.
    InstanceId *string          `json:"instance_id"`
    // Settings for requesting on-demand backup directly.
    Settings   *OnDemandSetting `json:"settings"`
    // The type of the backup. Possible values - `clumio_backup`, `aws_snapshot`. The
    // type will be assumed as `clumio_backup` if the field is left empty.
    ClumioType *string          `json:"type"`
}

// CreateBackupEc2MssqlDatabaseV1Request represents a custom type struct
type CreateBackupEc2MssqlDatabaseV1Request struct {
    // Performs the operation on the MSSQL asset with the specified Clumio-assigned ID.
    AssetId    *string          `json:"asset_id"`
    // Settings for requesting on-demand backup directly.
    Settings   *OnDemandSetting `json:"settings"`
    // The type of the backup.
    ClumioType *string          `json:"type"`
}

// CreateAwsConnectionV1Request represents a custom type struct
type CreateAwsConnectionV1Request struct {
    // The AWS-assigned ID of the account associated with the connection.
    AccountNativeId          *string   `json:"account_native_id"`
    // The AWS region associated with the connection. For example, `us-east-1`.
    AwsRegion                *string   `json:"aws_region"`
    // The user-provided description for this connection.
    Description              *string   `json:"description"`
    // The Clumio-assigned ID of the organizational unit associated with the
    // AWS environment. If this parameter is not provided, then the value
    // defaults to the first organizational unit assigned to the requesting
    // user. For more information about organizational units, refer to the
    // Organizational-Units documentation.
    OrganizationalUnitId     *string   `json:"organizational_unit_id"`
    // The asset types enabled for protect.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS",
    // "IcebergOnGlue", "IcebergOnS3Tables", "FSX"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    ProtectAssetTypesEnabled []*string `json:"protect_asset_types_enabled"`
    // The services to be enabled for this configuration. Valid values are
    // ["discover"], ["discover", "protect"]. This is only set when the
    // registration is created, the enabled services are obtained directly from
    // the installed template after that. (Deprecated as all connections will have
    // both discover and protect enabled)
    ServicesEnabled          []*string `json:"services_enabled"`
}

// CreateAwsConnectionGroupV1Request represents a custom type struct.
// The body of the request.
type CreateAwsConnectionGroupV1Request struct {
    // The AWS-assigned ID of the account to be associated with the Connection Group.
    AccountNativeId       *string   `json:"account_native_id"`
    // The asset types to be connected via the connection-group.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS",
    // "IcebergOnGlue", "IcebergOnS3Tables", "FSX"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    AssetTypes            []*string `json:"asset_types"`
    // DEPRECATED, use "asset_types" instead.
    // 
    // The asset types to be connected via the connection-group.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS",
    // "IcebergOnGlue", "IcebergOnS3Tables", "FSX"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    AssetTypesEnabled     []*string `json:"asset_types_enabled"`
    // The AWS regions to be associated with the Connection Group.
    AwsRegions            []*string `json:"aws_regions"`
    // Description for this connection group.
    Description           *string   `json:"description"`
    // The AWS Account that manages the connection-group's stack.
    // If the provided master_aws_account_id different than the account_native_id then
    // use service managed permissions while deploying stack.
    MasterAwsAccountId    *string   `json:"master_aws_account_id"`
    // The AWS Region that manages the connection-group's stack.
    MasterRegion          *string   `json:"master_region"`
    // The Clumio-assigned ID of the organizational unit associated with the
    // AWS environment. If this parameter is not provided, then the value
    // defaults to the first organizational unit assigned to the requesting
    // user. For more information about organizational units, refer to the
    // Organizational-Units documentation.
    OrganizationalUnitId  *string   `json:"organizational_unit_id"`
    // TODO: Add struct field description
    TemplatePermissionSet *string   `json:"template_permission_set"`
}

// UpdateAwsConnectionGroupV1Request represents a custom type struct
type UpdateAwsConnectionGroupV1Request struct {
    // The asset types to be connected via the connection-group.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS",
    // "IcebergOnGlue", "IcebergOnS3Tables", "FSX"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    AssetTypes            []*string `json:"asset_types"`
    // DEPRECATED, use "asset_types" instead.
    // 
    // 
    // The asset types to be connected via the connection-group.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS",
    // "IcebergOnGlue", "IcebergOnS3Tables", "FSX"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    AssetTypesEnabled     []*string `json:"asset_types_enabled"`
    // The AWS regions to be associated with the Connection Group.
    AwsRegions            []*string `json:"aws_regions"`
    // Description for this connection group.
    Description           *string   `json:"description"`
    // TODO: Add struct field description
    TemplatePermissionSet *string   `json:"template_permission_set"`
}

// PostProcessAwsConnectionV1Request represents a custom type struct.
// The body of the request.
type PostProcessAwsConnectionV1Request struct {
    // The AWS-assigned ID of the account associated with the connection.
    AccountNativeId     *string            `json:"account_native_id"`
    // The AWS region associated with the connection. For example, `us-east-1`.
    AwsRegion           *string            `json:"aws_region"`
    // ClumioEventPubId is the Clumio Event Pub SNS topic ID.
    ClumioEventPubId    *string            `json:"clumio_event_pub_id"`
    // Configuration represents the AWS connection configuration in json string format
    Configuration       *string            `json:"configuration"`
    // Role arn to be assumed before accessing ClumioRole in customer account
    IntermediateRoleArn *string            `json:"intermediate_role_arn"`
    // Properties is a key value map meant to be used for passing additional information
    // like resource IDs/ARNs.
    Properties          map[string]*string `json:"properties"`
    // RequestType indicates whether this is a Create, Update or Delete request
    RequestType         *string            `json:"request_type"`
    // RoleArn is the ARN of the ClumioIAMRole created in the customer account
    RoleArn             *string            `json:"role_arn"`
    // Role External Id is the unique string passed while creating the AWS resources .
    RoleExternalId      *string            `json:"role_external_id"`
    // The 36-character Clumio AWS integration ID token used to identify the
    // installation of the CloudFormation/Terraform template on the account.
    Token               *string            `json:"token"`
}

// CreateConnectionTemplateV1Request represents a custom type struct
type CreateConnectionTemplateV1Request struct {
    // The asset types for which the template is to be generated.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS", "IcebergOnGlue", "IcebergOnS3Tables"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    AssetTypesEnabled     []*string `json:"asset_types_enabled"`
    // Account ID for the AWS environment to be connected
    // Mandatory to pass a 12 digit string if show_manual_resources is set to true
    AwsAccountId          *string   `json:"aws_account_id"`
    // AWS Region of the AWS environment to be connected
    // Mandatory to pass a non-empty string if show_manual_resources is set to true
    AwsRegion             *string   `json:"aws_region"`
    // Returns the resources to be created manually if set to true
    ShowManualResources   *bool     `json:"show_manual_resources"`
    // TODO: Add struct field description
    TemplatePermissionSet *string   `json:"template_permission_set"`
}

// UpdateAwsConnectionV1Request represents a custom type struct
type UpdateAwsConnectionV1Request struct {
    // Asset types enabled with the given resource ARNs.
    // This field is only applicable to manually configured connections.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS", "IcebergOnGlue", "IcebergOnS3Tables"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    AssetTypesEnabled []*string  `json:"asset_types_enabled"`
    // An optional, user-provided description for this connection.
    Description       *string    `json:"description"`
    // Partial updates are not supported, therefore you must provide ARNs for all configured resources,
    // including those for resources that are not being updated.
    Resources         *Resources `json:"resources"`
}

// SetBucketPropertiesV1Request represents a custom type struct.
// The set of properties that are being updated for the given bucket.
type SetBucketPropertiesV1Request struct {
    // If true, enables continuous backup for the given bucket.
    // If false, disables continuous backup for the given bucket.
    // If not set, does not update EventBridge.
    EventBridgeEnabled              *bool `json:"event_bridge_enabled"`
    // If true, tries to disable EventBridge notification for the given bucket.
    // It may override the existing bucket notification configuration in the customer's account.
    // This takes effect only when `event_bridge_enabled` is set to `false`.
    EventBridgeNotificationDisabled *bool `json:"event_bridge_notification_disabled"`
}

// UpdateManagementGroupV1Request represents a custom type struct
type UpdateManagementGroupV1Request struct {
    // Determines whether backups are allowed to occur across different subgroups or cloud connectors.
    BackupAcrossSubgroups *bool   `json:"backup_across_subgroups"`
    // The name of the management group.
    Name                  *string `json:"name"`
}

// CreateOrganizationalUnitV2Request represents a custom type struct
type CreateOrganizationalUnitV2Request struct {
    // A description of the organizational unit.
    Description *string         `json:"description"`
    // entityModel denotes the entityModel
    Entities    []*EntityModel  `json:"entities"`
    // Unique name assigned to the organizational unit.
    Name        *string         `json:"name"`
    // The Clumio-assigned ID of the parent organizational unit under which the new organizational unit is to be created.
    // If absent, the new organizational unit is created under the current organizational unit.
    ParentId    *string         `json:"parent_id"`
    // A user along with a role.
    Users       []*UserWithRole `json:"users"`
}

// CreateOrganizationalUnitV1Request represents a custom type struct
type CreateOrganizationalUnitV1Request struct {
    // A description of the organizational unit.
    Description *string        `json:"description"`
    // entityModel denotes the entityModel
    Entities    []*EntityModel `json:"entities"`
    // Unique name assigned to the organizational unit.
    Name        *string        `json:"name"`
    // The Clumio-assigned ID of the parent organizational unit under which the new organizational unit is to be created.
    // If absent, the new organizational unit is created under the current organizational unit.
    ParentId    *string        `json:"parent_id"`
    // List of user IDs to assign this organizational unit.
    Users       []*string      `json:"users"`
}

// PatchOrganizationalUnitV2Request represents a custom type struct
type PatchOrganizationalUnitV2Request struct {
    // A description of the organizational unit.
    Description      *string                           `json:"description"`
    // Updates to the entities in the organizational unit.
    // Adding or removing entities from the OU is an asynchronous operation.
    // The response has a task ID which can be used to track the progress of the operation.
    Entities         *UpdateEntities                   `json:"entities"`
    // Unique name assigned to the organizational unit.
    Name             *string                           `json:"name"`
    // UpdateProtectionGroupAssignments denotes the protection groups to be assigned or
    // unassigned.
    // Updates to the protection group assignments.
    ProtectionGroups *UpdateProtectionGroupAssignments `json:"protection_groups"`
    // UpdateUserAssignmentsWithRole denotes the users to be added or removed along with the role.
    Users            *UpdateUserAssignmentsWithRole    `json:"users"`
}

// PatchOrganizationalUnitV1Request represents a custom type struct
type PatchOrganizationalUnitV1Request struct {
    // A description of the organizational unit.
    Description *string                `json:"description"`
    // Updates to the entities in the organizational unit.
    // Adding or removing entities from the OU is an asynchronous operation.
    // The response has a task ID which can be used to track the progress of the operation.
    Entities    *UpdateEntities        `json:"entities"`
    // Unique name assigned to the organizational unit.
    Name        *string                `json:"name"`
    // Updates to the user assignments.
    Users       *UpdateUserAssignments `json:"users"`
}

// SetPolicyAssignmentsV1Request represents a custom type struct
type SetPolicyAssignmentsV1Request struct {
    // The portion of the policy assignment used for updates/creations
    Items []*AssignmentInputModel `json:"items"`
}

// CreatePolicyDefinitionV1Request represents a custom type struct
type CreatePolicyDefinitionV1Request struct {
    // The status of the policy.
    // Refer to the Policy Activation Status table for a complete list of policy statuses.
    ActivationStatus     *string                 `json:"activation_status"`
    // The user-provided name of the policy.
    Name                 *string                 `json:"name"`
    // TODO: Add struct field description
    Operations           []*PolicyOperationInput `json:"operations"`
    // The Clumio-assigned ID of the organizational unit associated with the policy.
    OrganizationalUnitId *string                 `json:"organizational_unit_id"`
    // The policy-level timezone is deprecated, as the operation-level timezone should be used instead.
    // The timezone must be a valid location name from the IANA Time Zone database.
    // For instance, "America/New_York", "US/Central", "UTC".
    // Deprecated: true
    Timezone             *string                 `json:"timezone"`
}

// UpdatePolicyDefinitionV1Request represents a custom type struct
type UpdatePolicyDefinitionV1Request struct {
    // The status of the policy.
    // Refer to the Policy Activation Status table for a complete list of policy statuses.
    ActivationStatus     *string                 `json:"activation_status"`
    // The user-provided name of the policy.
    Name                 *string                 `json:"name"`
    // TODO: Add struct field description
    Operations           []*PolicyOperationInput `json:"operations"`
    // The Clumio-assigned ID of the organizational unit associated with the policy.
    OrganizationalUnitId *string                 `json:"organizational_unit_id"`
    // The policy-level timezone is deprecated, as the operation-level timezone should be used instead.
    // The timezone must be a valid location name from the IANA Time Zone database.
    // For instance, "America/New_York", "US/Central", "UTC".
    // Deprecated: true
    Timezone             *string                 `json:"timezone"`
}

// CreatePolicyRuleV1Request represents a custom type struct
type CreatePolicyRuleV1Request struct {
    // An action to be applied subject to the rule criteria.
    Action    *RuleAction   `json:"action"`
    // The following table describes the possible conditions for a rule.
    // 
    // +-----------------------+---------------------------+--------------------------+
    // |         Field         |      Rule Condition       |       Description        |
    // +=======================+===========================+==========================+
    // | aws_account_native_id | $eq, $in                  | Denotes the AWS account  |
    // |                       |                           | to conditionalize on     |
    // |                       |                           |                          |
    // |                       |                           | {"aws_account_native_id" |
    // |                       |                           | :{"$eq":"111111111111"}} |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_account_native_id" |
    // |                       |                           | :{"$in":["111111111111", |
    // |                       |                           | "222222222222"]}}        |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // +-----------------------+---------------------------+--------------------------+
    // | aws_region            | $eq, $in                  | Denotes the AWS region   |
    // |                       |                           | to conditionalize on     |
    // |                       |                           |                          |
    // |                       |                           | {"aws_region":{"$eq":"us |
    // |                       |                           | -west-2"}}               |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_region":{"$in":["u |
    // |                       |                           | s-west-2", "us-          |
    // |                       |                           | east-1"]}}               |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // +-----------------------+---------------------------+--------------------------+
    // | aws_tag               | $eq, $in, $all,           | Denotes the AWS tag(s)   |
    // |                       | $contains, $not_eq,       | to conditionalize on.    |
    // |                       | $not_in, $not_all,        | Max 100 tags allowed in  |
    // |                       | $not_contains             | each rule                |
    // |                       |                           | and tag key can be upto  |
    // |                       |                           | 128 characters and value |
    // |                       |                           | can be upto 256          |
    // |                       |                           | characters long.         |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$eq":{"key" |
    // |                       |                           | :"Environment",          |
    // |                       |                           | "value":"Prod"}}}        |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$in":[{"key |
    // |                       |                           | ":"Environment",         |
    // |                       |                           | "value":"Prod"},         |
    // |                       |                           | {"key":"Hello",          |
    // |                       |                           | "value":"World"}]}}      |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$all":[{"ke |
    // |                       |                           | y":"Environment",        |
    // |                       |                           | "value":"Prod"},         |
    // |                       |                           | {"key":"Hello",          |
    // |                       |                           | "value":"World"}]}}      |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$contains": |
    // |                       |                           | {"key":"Environment",    |
    // |                       |                           | "value":"Prod"}}}        |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$not_eq":{" |
    // |                       |                           | key":"Environment",      |
    // |                       |                           | "value":"Prod"}}}        |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$not_in":[{ |
    // |                       |                           | "key":"Environment",     |
    // |                       |                           | "value":"Prod"},         |
    // |                       |                           | {"key":"Hello",          |
    // |                       |                           | "value":"World"}]}}      |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$not_all":[ |
    // |                       |                           | {"key":"Environment",    |
    // |                       |                           | "value":"Prod"},         |
    // |                       |                           | {"key":"Hello",          |
    // |                       |                           | "value":"World"}]}}      |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$not_contai |
    // |                       |                           | ns":{"key":"Environment" |
    // |                       |                           | , "value":"Prod"}}}      |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // +-----------------------+---------------------------+--------------------------+
    // | entity_type           | $eq, $in                  | Denotes the AWS entity   |
    // |                       |                           | type to conditionalize   |
    // |                       |                           | on. (Required)           |
    // |                       |                           |                          |
    // |                       |                           | {"entity_type":{"$eq":"a |
    // |                       |                           | ws_rds_instance"}}       |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"entity_type":{"$in":[" |
    // |                       |                           | aws_rds_instance",       |
    // |                       |                           | "aws_ebs_volume", "aws_e |
    // |                       |                           | c2_instance","aws_dynamo |
    // |                       |                           | db_table",               |
    // |                       |                           | "aws_rds_cluster"]}}     |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // +-----------------------+---------------------------+--------------------------+
    // 
    Condition *string       `json:"condition"`
    // Name of the rule. Max 100 characters.
    Name      *string       `json:"name"`
    // A priority relative to other rules.
    Priority  *RulePriority `json:"priority"`
}

// UpdatePolicyRuleV1Request represents a custom type struct
type UpdatePolicyRuleV1Request struct {
    // An action to be applied subject to the rule criteria.
    Action    *RuleAction   `json:"action"`
    // The following table describes the possible conditions for a rule.
    // 
    // +-----------------------+---------------------------+--------------------------+
    // |         Field         |      Rule Condition       |       Description        |
    // +=======================+===========================+==========================+
    // | aws_account_native_id | $eq, $in                  | Denotes the AWS account  |
    // |                       |                           | to conditionalize on     |
    // |                       |                           |                          |
    // |                       |                           | {"aws_account_native_id" |
    // |                       |                           | :{"$eq":"111111111111"}} |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_account_native_id" |
    // |                       |                           | :{"$in":["111111111111", |
    // |                       |                           | "222222222222"]}}        |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // +-----------------------+---------------------------+--------------------------+
    // | aws_region            | $eq, $in                  | Denotes the AWS region   |
    // |                       |                           | to conditionalize on     |
    // |                       |                           |                          |
    // |                       |                           | {"aws_region":{"$eq":"us |
    // |                       |                           | -west-2"}}               |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_region":{"$in":["u |
    // |                       |                           | s-west-2", "us-          |
    // |                       |                           | east-1"]}}               |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // +-----------------------+---------------------------+--------------------------+
    // | aws_tag               | $eq, $in, $all,           | Denotes the AWS tag(s)   |
    // |                       | $contains, $not_eq,       | to conditionalize on.    |
    // |                       | $not_in, $not_all,        | Max 100 tags allowed in  |
    // |                       | $not_contains             | each rule                |
    // |                       |                           | and tag key can be upto  |
    // |                       |                           | 128 characters and value |
    // |                       |                           | can be upto 256          |
    // |                       |                           | characters long.         |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$eq":{"key" |
    // |                       |                           | :"Environment",          |
    // |                       |                           | "value":"Prod"}}}        |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$in":[{"key |
    // |                       |                           | ":"Environment",         |
    // |                       |                           | "value":"Prod"},         |
    // |                       |                           | {"key":"Hello",          |
    // |                       |                           | "value":"World"}]}}      |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$all":[{"ke |
    // |                       |                           | y":"Environment",        |
    // |                       |                           | "value":"Prod"},         |
    // |                       |                           | {"key":"Hello",          |
    // |                       |                           | "value":"World"}]}}      |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$contains": |
    // |                       |                           | {"key":"Environment",    |
    // |                       |                           | "value":"Prod"}}}        |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$not_eq":{" |
    // |                       |                           | key":"Environment",      |
    // |                       |                           | "value":"Prod"}}}        |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$not_in":[{ |
    // |                       |                           | "key":"Environment",     |
    // |                       |                           | "value":"Prod"},         |
    // |                       |                           | {"key":"Hello",          |
    // |                       |                           | "value":"World"}]}}      |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$not_all":[ |
    // |                       |                           | {"key":"Environment",    |
    // |                       |                           | "value":"Prod"},         |
    // |                       |                           | {"key":"Hello",          |
    // |                       |                           | "value":"World"}]}}      |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"aws_tag":{"$not_contai |
    // |                       |                           | ns":{"key":"Environment" |
    // |                       |                           | , "value":"Prod"}}}      |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // +-----------------------+---------------------------+--------------------------+
    // | entity_type           | $eq, $in                  | Denotes the AWS entity   |
    // |                       |                           | type to conditionalize   |
    // |                       |                           | on. (Required)           |
    // |                       |                           |                          |
    // |                       |                           | {"entity_type":{"$eq":"a |
    // |                       |                           | ws_rds_instance"}}       |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // |                       |                           | {"entity_type":{"$in":[" |
    // |                       |                           | aws_rds_instance",       |
    // |                       |                           | "aws_ebs_volume", "aws_e |
    // |                       |                           | c2_instance","aws_dynamo |
    // |                       |                           | db_table",               |
    // |                       |                           | "aws_rds_cluster"]}}     |
    // |                       |                           |                          |
    // |                       |                           |                          |
    // +-----------------------+---------------------------+--------------------------+
    // 
    Condition *string       `json:"condition"`
    // Name of the rule. Max 100 characters.
    Name      *string       `json:"name"`
    // A priority relative to other rules.
    Priority  *RulePriority `json:"priority"`
}

// CreateProtectionGroupV1Request represents a custom type struct
type CreateProtectionGroupV1Request struct {
    // The following table describes the possible conditions for a bucket to be
    // automatically added to a protection group. 
    // Denotes the properties to conditionalize on. For `$eq`, `$not_eq`, `$contains` and `$not_contains` a single element is provided: `{'$eq':{'key':'Environment', 'value':'Prod'}}`. For all other operations, a list is provided: `{'$in':[{'key':'Environment','value':'Prod'}, {'key':'Hello', 'value':'World'}]}`.
    // 
    // +-------------------+-----------------------------+----------------------------+
    // |       Field       |       Rule Condition        |        Description         |
    // +===================+=============================+============================+
    // | aws_tag           | $eq, $not_eq, $contains,    | Supports filtering by AWS  |
    // |                   | $not_contains, $all,        | tag(s) using the following |
    // |                   | $not_all, $in, $not_in      | operators. For example,    |
    // |                   |                             |                            |
    // |                   |                             | {"aws_tag":{"$eq":{"key":" |
    // |                   |                             | Environment",              |
    // |                   |                             | "value":"Prod"}}}          |
    // |                   |                             |                            |
    // |                   |                             |                            |
    // +-------------------+-----------------------------+----------------------------+
    // | account_native_id | $eq, $in                    |                            |
    // |                   |                             | This will be deprecated    |
    // |                   |                             | and use                    |
    // |                   |                             | aws_account_native_id      |
    // |                   |                             | instead.                   |
    // |                   |                             | Supports filtering by AWS  |
    // |                   |                             | account(s) using the       |
    // |                   |                             | following operators. For   |
    // |                   |                             | example,                   |
    // |                   |                             |                            |
    // |                   |                             | {"account_native_id":{"$in |
    // |                   |                             | ":["111111111111"]}}       |
    // |                   |                             |                            |
    // |                   |                             |                            |
    // +-------------------+-----------------------------+----------------------------+
    // | aws_region        | $eq, $in                    | Supports filtering by AWS  |
    // |                   |                             | region(s) using the        |
    // |                   |                             | following operators. For   |
    // |                   |                             | example,                   |
    // |                   |                             |                            |
    // |                   |                             | {"aws_region":{"$eq":"us-  |
    // |                   |                             | west-2"}}                  |
    // |                   |                             |                            |
    // |                   |                             |                            |
    // +-------------------+-----------------------------+----------------------------+
    // 
    BucketRule   *string       `json:"bucket_rule"`
    // The user-assigned description of the protection group.
    Description  *string       `json:"description"`
    // The user-assigned name of the protection group.
    Name         *string       `json:"name"`
    // ObjectFilter
    // defines which objects will be backed up.
    ObjectFilter *ObjectFilter `json:"object_filter"`
}

// UpdateProtectionGroupV1Request represents a custom type struct
type UpdateProtectionGroupV1Request struct {
    // The following table describes the possible conditions for a bucket to be
    // automatically added to a protection group. 
    // Denotes the properties to conditionalize on. For `$eq`, `$not_eq`, `$contains` and `$not_contains` a single element is provided: `{'$eq':{'key':'Environment', 'value':'Prod'}}`. For all other operations, a list is provided: `{'$in':[{'key':'Environment','value':'Prod'}, {'key':'Hello', 'value':'World'}]}`.
    // 
    // +-------------------+-----------------------------+----------------------------+
    // |       Field       |       Rule Condition        |        Description         |
    // +===================+=============================+============================+
    // | aws_tag           | $eq, $not_eq, $contains,    | Supports filtering by AWS  |
    // |                   | $not_contains, $all,        | tag(s) using the following |
    // |                   | $not_all, $in, $not_in      | operators. For example,    |
    // |                   |                             |                            |
    // |                   |                             | {"aws_tag":{"$eq":{"key":" |
    // |                   |                             | Environment",              |
    // |                   |                             | "value":"Prod"}}}          |
    // |                   |                             |                            |
    // |                   |                             |                            |
    // +-------------------+-----------------------------+----------------------------+
    // | account_native_id | $eq, $in                    |                            |
    // |                   |                             | This will be deprecated    |
    // |                   |                             | and use                    |
    // |                   |                             | aws_account_native_id      |
    // |                   |                             | instead.                   |
    // |                   |                             | Supports filtering by AWS  |
    // |                   |                             | account(s) using the       |
    // |                   |                             | following operators. For   |
    // |                   |                             | example,                   |
    // |                   |                             |                            |
    // |                   |                             | {"account_native_id":{"$in |
    // |                   |                             | ":["111111111111"]}}       |
    // |                   |                             |                            |
    // |                   |                             |                            |
    // +-------------------+-----------------------------+----------------------------+
    // | aws_region        | $eq, $in                    | Supports filtering by AWS  |
    // |                   |                             | region(s) using the        |
    // |                   |                             | following operators. For   |
    // |                   |                             | example,                   |
    // |                   |                             |                            |
    // |                   |                             | {"aws_region":{"$eq":"us-  |
    // |                   |                             | west-2"}}                  |
    // |                   |                             |                            |
    // |                   |                             |                            |
    // +-------------------+-----------------------------+----------------------------+
    // 
    BucketRule   *string       `json:"bucket_rule"`
    // The user-assigned description of the protection group.
    Description  *string       `json:"description"`
    // The user-assigned name of the protection group.
    Name         *string       `json:"name"`
    // ObjectFilter
    // defines which objects will be backed up.
    ObjectFilter *ObjectFilter `json:"object_filter"`
}

// AddBucketProtectionGroupV1Request represents a custom type struct
type AddBucketProtectionGroupV1Request struct {
    // TODO: Add struct field description
    BucketId *string `json:"bucket_id"`
}

// CreateComplianceReportConfigurationV1Request represents a custom type struct
type CreateComplianceReportConfigurationV1Request struct {
    // The user-provided description of the compliance report configuration.
    Description  *string              `json:"description"`
    // The user-provided name of the compliance report configuration.
    Name         *string              `json:"name"`
    // Notification channels to send the generated report runs.
    Notification *NotificationSetting `json:"notification"`
    // Filter and control parameters of compliance report.
    Parameter    *Parameter           `json:"parameter"`
    // When the report will be generated and sent. If the schedule is not provided then a
    // default value will be used.
    Schedule     *ScheduleSetting     `json:"schedule"`
}

// UpdateComplianceReportConfigurationV1Request represents a custom type struct
type UpdateComplianceReportConfigurationV1Request struct {
    // The user-provided description of the compliance report configuration.
    Description  *string              `json:"description"`
    // The user-provided name of the compliance report configuration.
    Name         *string              `json:"name"`
    // Notification channels to send the generated report runs.
    Notification *NotificationSetting `json:"notification"`
    // Filter and control parameters of compliance report.
    Parameter    *Parameter           `json:"parameter"`
    // When the report will be generated and sent. If the schedule is not provided then a
    // default value will be used.
    Schedule     *ScheduleSetting     `json:"schedule"`
}

// CreateComplianceReportRunV1Request represents a custom type struct
type CreateComplianceReportRunV1Request struct {
    // Name of the new compliance report run that will be created.
    // If not given, default uses `{configuration ID} - {MM/DD/YYYY(Created time)}`.
    Name *string `json:"name"`
}

// SendComplianceReportRunEmailV1Request represents a custom type struct
type SendComplianceReportRunEmailV1Request struct {
    // List of emails to be notified.
    EmailList []*string `json:"email_list"`
}

// CreateReportDownloadV1Request represents a custom type struct
type CreateReportDownloadV1Request struct {
    // The name of the report. Field cannot be empty.
    FileName   *string `json:"file_name"`
    // 
    // +----------------------+------------------+-------------+----------------------+
    // |        Field         | Filter Condition | Report Type |     Description      |
    // +======================+==================+=============+======================+
    // | activity_start_times | $gte, $lt, $eq   | Activity    | Activity start       |
    // | tamp                 |                  |             | timestamp denotes    |
    // |                      |                  |             | the time filter for  |
    // |                      |                  |             | Activity reports.    |
    // |                      |                  |             | $gte and $lt accept  |
    // |                      |                  |             | RFC-3339 timestamps  |
    // |                      |                  |             | and $eq accepts a    |
    // |                      |                  |             | unix timestamp       |
    // |                      |                  |             | denoting the offset  |
    // |                      |                  |             | from the current     |
    // |                      |                  |             | time. $eq takes      |
    // |                      |                  |             | precedence over both |
    // |                      |                  |             | $gte and $lt so if   |
    // |                      |                  |             | $eq is used, the     |
    // |                      |                  |             | backend will use the |
    // |                      |                  |             | relative time filter |
    // |                      |                  |             | instead of absolute  |
    // |                      |                  |             | time filters.For     |
    // |                      |                  |             | example,             |
    // |                      |                  |             |                      |
    // |                      |                  |             | "filter":"{"activity |
    // |                      |                  |             | _start_timestamp":{" |
    // |                      |                  |             | $eq":86400}}"        |
    // |                      |                  |             |                      |
    // |                      |                  |             |                      |
    // +----------------------+------------------+-------------+----------------------+
    // | timestamp            | $gte, $lte, $eq  | Consumption | timestamp denotes    |
    // |                      |                  |             | the time filter for  |
    // |                      |                  |             | Consumption reports. |
    // |                      |                  |             | $gte and $lte accept |
    // |                      |                  |             | RFC-3339 timestamps  |
    // |                      |                  |             | and $eq accepts a    |
    // |                      |                  |             | duration in seconds  |
    // |                      |                  |             | denoting the offset  |
    // |                      |                  |             | from the current     |
    // |                      |                  |             | time. $eq takes      |
    // |                      |                  |             | precedence over both |
    // |                      |                  |             | $gte and $lte so if  |
    // |                      |                  |             | $eq is used, the     |
    // |                      |                  |             | backend will use the |
    // |                      |                  |             | relative time filter |
    // |                      |                  |             | instead of absolute  |
    // |                      |                  |             | time filters.        |
    // |                      |                  |             |                      |
    // |                      |                  |             | "filter": "{\"timest |
    // |                      |                  |             | amp\":{\"$gte\":\"20 |
    // |                      |                  |             | 22-07-               |
    // |                      |                  |             | 27T14:35:33.735Z\",\ |
    // |                      |                  |             | "$lte\":\"2022-08-   |
    // |                      |                  |             | 03T14:35:33.735Z\"}} |
    // |                      |                  |             | "                    |
    // |                      |                  |             |                      |
    // |                      |                  |             |                      |
    // +----------------------+------------------+-------------+----------------------+
    // | organizational_unit_ | $in              | Consumption | Organizational Unit  |
    // | id                   |                  |             | ID (OU ID) filters   |
    // |                      |                  |             | the consumption data |
    // |                      |                  |             | generated for the    |
    // |                      |                  |             | report to the given  |
    // |                      |                  |             | OU IDs and its       |
    // |                      |                  |             | children.            |
    // |                      |                  |             |                      |
    // |                      |                  |             | "filter": "{\"organi |
    // |                      |                  |             | zational_unit_id\":{ |
    // |                      |                  |             | \"$in\":[\"00000000- |
    // |                      |                  |             | 0000-0000-0000-      |
    // |                      |                  |             | 000000000000\"]}}"   |
    // |                      |                  |             |                      |
    // |                      |                  |             |                      |
    // +----------------------+------------------+-------------+----------------------+
    // | entity_type          | $in              | Consumption |                      |
    // |                      |                  |             | Entity type filters  |
    // |                      |                  |             | the consumption data |
    // |                      |                  |             | generated for the    |
    // |                      |                  |             | report to the given  |
    // |                      |                  |             | entity types.        |
    // |                      |                  |             | If filter is empty,  |
    // |                      |                  |             | it shows consumption |
    // |                      |                  |             | report of the all    |
    // |                      |                  |             | entity types.        |
    // |                      |                  |             |                      |
    // |                      |                  |             | filter={"entity_type |
    // |                      |                  |             | ":{"$in":["LocalProt |
    // |                      |                  |             | ectionGroup,         |
    // |                      |                  |             | DynamoDB"]}}         |
    // |                      |                  |             |                      |
    // |                      |                  |             |                      |
    // +----------------------+------------------+-------------+----------------------+
    // | task                 | $in              | Activity    |  Possible values for |
    // |                      |                  |             | task include backup  |
    // |                      |                  |             | and                  |
    // |                      |                  |             | restore.             |
    // |                      |                  |             | For example,         |
    // |                      |                  |             |                      |
    // |                      |                  |             | "filter":"{"task":{" |
    // |                      |                  |             | $in":["ebs_increment |
    // |                      |                  |             | al_backup"]}}"       |
    // |                      |                  |             |                      |
    // |                      |                  |             |                      |
    // +----------------------+------------------+-------------+----------------------+
    // | status               | $in              | Activity    |  Filter on activity  |
    // |                      |                  |             | status of entity.    |
    // |                      |                  |             | Possible values for  |
    // |                      |                  |             | activity status      |
    // |                      |                  |             | include aborted,     |
    // |                      |                  |             | failed, and success  |
    // |                      |                  |             | For example,         |
    // |                      |                  |             |                      |
    // |                      |                  |             | "filter": "{"status" |
    // |                      |                  |             | :{"$in":["success"]} |
    // |                      |                  |             | }"                   |
    // |                      |                  |             |                      |
    // |                      |                  |             |                      |
    // +----------------------+------------------+-------------+----------------------+
    // | primary_entity.id    | $in              | Any         | The system-generated |
    // |                      |                  |             | IDs of the primary   |
    // |                      |                  |             | entities affected by |
    // |                      |                  |             | the activity.        |
    // |                      |                  |             | For example,         |
    // |                      |                  |             |                      |
    // |                      |                  |             | filter={"primary_ent |
    // |                      |                  |             | ity.id":{"$in":["9c2 |
    // |                      |                  |             | 934fc-ff4d-11e9-     |
    // |                      |                  |             | 8e11-                |
    // |                      |                  |             | 76706df7fe01"]}}     |
    // |                      |                  |             |                      |
    // |                      |                  |             |                      |
    // +----------------------+------------------+-------------+----------------------+
    // | primary_entity.type  | $eq              | Any         | The type of primary  |
    // |                      |                  |             | entities affected by |
    // |                      |                  |             | the activity.        |
    // |                      |                  |             | Examples of primary  |
    // |                      |                  |             | entity types include |
    // |                      |                  |             | "aws_ebs_volume",    |
    // |                      |                  |             | "aws_ec2_instance",  |
    // |                      |                  |             | "microsoft365_mailbo |
    // |                      |                  |             | x". For example,     |
    // |                      |                  |             |                      |
    // |                      |                  |             | filter={"primary_ent |
    // |                      |                  |             | ity.type":{"$in":["a |
    // |                      |                  |             | ws_ebs_volume"]}}    |
    // |                      |                  |             |                      |
    // |                      |                  |             |                      |
    // +----------------------+------------------+-------------+----------------------+
    // | primary_entity.value | $in              | Any         | The value or name    |
    // |                      |                  |             | associated with the  |
    // |                      |                  |             | primary entities     |
    // |                      |                  |             | affected by          |
    // |                      |                  |             | the compliance       |
    // |                      |                  |             | event. For example,  |
    // |                      |                  |             | the primary entity   |
    // |                      |                  |             | value associated     |
    // |                      |                  |             | with                 |
    // |                      |                  |             | primary entity type  |
    // |                      |                  |             | "aws_ebs_volume" is  |
    // |                      |                  |             | "vol-                |
    // |                      |                  |             | 0a5f2e52d6decd664"   |
    // |                      |                  |             | representing         |
    // |                      |                  |             | the name of the EBS  |
    // |                      |                  |             | volume. The filter   |
    // |                      |                  |             | supports substring   |
    // |                      |                  |             | search for all       |
    // |                      |                  |             | elements in the      |
    // |                      |                  |             | array For example,   |
    // |                      |                  |             |                      |
    // |                      |                  |             | filter={"primary_ent |
    // |                      |                  |             | ity.value":{"$in":[" |
    // |                      |                  |             | vol-0a"]}}           |
    // |                      |                  |             |                      |
    // |                      |                  |             |                      |
    // +----------------------+------------------+-------------+----------------------+
    // | parent_entity.type   | $in              | Any         |  The types of the    |
    // |                      |                  |             | parent entities      |
    // |                      |                  |             | which are associated |
    // |                      |                  |             | with the primary     |
    // |                      |                  |             | entity affected by   |
    // |                      |                  |             | the activity.        |
    // |                      |                  |             | Examples of the      |
    // |                      |                  |             | parent entity types  |
    // |                      |                  |             | include              |
    // |                      |                  |             | "aws_environment",   |
    // |                      |                  |             | "microsoft365_domain |
    // |                      |                  |             | ". For example,      |
    // |                      |                  |             |                      |
    // |                      |                  |             | filter={"parent_enti |
    // |                      |                  |             | ty.type":{"$in":["aw |
    // |                      |                  |             | s_environment"]}}    |
    // |                      |                  |             |                      |
    // |                      |                  |             |                      |
    // +----------------------+------------------+-------------+----------------------+
    // | parent_entity.id     | $in              | Any         |                      |
    // |                      |                  |             | The value or name    |
    // |                      |                  |             | associated with the  |
    // |                      |                  |             | parent entities      |
    // |                      |                  |             | affected by          |
    // |                      |                  |             | the compliance       |
    // |                      |                  |             | event. For example,  |
    // |                      |                  |             | the parent entity    |
    // |                      |                  |             | value associated     |
    // |                      |                  |             | with                 |
    // |                      |                  |             | primary entity type  |
    // |                      |                  |             | "aws_ebs_volume" is  |
    // |                      |                  |             | "891106093485/us-    |
    // |                      |                  |             | west-2" representing |
    // |                      |                  |             | the name of the AWS  |
    // |                      |                  |             | Account Region. For  |
    // |                      |                  |             | example,             |
    // |                      |                  |             |                      |
    // |                      |                  |             | filter={"parent_enti |
    // |                      |                  |             | ty.id":{"$in":["9c29 |
    // |                      |                  |             | 34fc-ff4d-11e9-8e11- |
    // |                      |                  |             | 76706df7fe01"]}}     |
    // |                      |                  |             |                      |
    // |                      |                  |             |                      |
    // +----------------------+------------------+-------------+----------------------+
    //  
    // For more information about filtering, refer to the
    // Filtering section of this guide.
    Filter     *string `json:"filter"`
    // The report type. Examples of report types include, "activity", "audit", and "consumption".
    ClumioType *string `json:"type"`
}

// RestoreAwsDynamodbTableV1Request represents a custom type struct
type RestoreAwsDynamodbTableV1Request struct {
    // Filters based on which DynamoDB backup records are filtered.
    QueryFilter *DynamoDBRestoreQueryFilter `json:"query_filter"`
    // The DynamoDB table restore backup or point-in-time restore options. Only one of these fields should be set.
    Source      *DynamoDBTableRestoreSource `json:"source"`
    // The configuration of the restored DynamoDB table.
    // For restore from snapshot, use the DynamoDB table configurations present at time of snapshot obtained from
    // [GET /backups/aws/dynamodb-tables/{backup_id}](#operation/read-backup-aws-dynamodb-table) and for restoring point-in-time,
    // use the current configuration of the table from [GET /datasources/aws/dynamodb-tables/{table_id}](#operation/read-aws-dynamodb-table).
    // The table properties are set to empty or to their default values if they are specified as `null`.
    Target      *DynamoDBTableRestoreTarget `json:"target"`
}

// RestoreRecordsAwsDynamodbTableV1Request represents a custom type struct
type RestoreRecordsAwsDynamodbTableV1Request struct {
    // Filters based on which DynamoDB backup records are filtered in the query.
    QueryFilter *DynamoDBGRRQueryFilter `json:"query_filter"`
    // The parameters for initiating a DynamoDB table backup query from a backup.
    Source      *DynamoDBGrrSource      `json:"source"`
    // The destination information for the operation, representing the access option
    // for the query results. Users can access the query results by direct download or by
    // email. The direct download (`direct_download`) option allows users to directly download
    // the restored file from the Clumio UI. The email (`email`) option allows users to access
    // the restored file using a downloadable link they receive by email. If a target is not
    // specified, then `target` defaults to `direct_download`.
    Target      *DynamoDBGrrTarget      `json:"target"`
}

// RestoreAwsEbsVolumeV1Request represents a custom type struct
type RestoreAwsEbsVolumeV1Request struct {
    // The EBS volume backup to be restored.
    Source *EBSRestoreSourceV1 `json:"source"`
    // The configuration of the EBS volume to be restored.
    Target *EBSRestoreTargetV1 `json:"target"`
}

// RestoreAwsEbsVolumeV2Request represents a custom type struct
type RestoreAwsEbsVolumeV2Request struct {
    // The EBS volume backup to be restored.
    Source *EBSRestoreSource `json:"source"`
    // The configuration of the EBS volume to be restored.
    Target *EBSRestoreTarget `json:"target"`
}

// RestoreAwsEc2InstanceV1Request represents a custom type struct
type RestoreAwsEc2InstanceV1Request struct {
    // The EC2 instance backup to be restored.
    Source *EC2RestoreSource `json:"source"`
    // The target configuration per EC2 restore type. Only one of these fields should be set.
    Target *EC2RestoreTarget `json:"target"`
}

// RestoreEc2MssqlDatabaseV1Request represents a custom type struct
type RestoreEc2MssqlDatabaseV1Request struct {
    // The EC2 MSSQL database backup to be restored. Only one of `backup` or `pitr`
    // should be set.
    // `pitr` A database backup at a specific point in time to be restored.
    Source *EC2MSSQLRestoreSource `json:"source"`
    // The configuration of the EC2 MSSQL database to which the data has to be restored.
    Target *EC2MSSQLRestoreTarget `json:"target"`
}

// RestoreAwsRdsResourceV1Request represents a custom type struct
type RestoreAwsRdsResourceV1Request struct {
    // The RDS resource backup or snapshot to be restored.  Only one of these fields should be set.
    Source *RdsResourceRestoreSource `json:"source"`
    // The configuration of the RDS resource to be restored.
    Target *RdsResourceRestoreTarget `json:"target"`
}

// RestoreRdsRecordV1Request represents a custom type struct
type RestoreRdsRecordV1Request struct {
    // The RDS database backup to be queried.
    Source *GrrSource `json:"source"`
    // The query to perform on the source RDS database.
    Target *GrrTarget `json:"target"`
}

// RestoreAwsS3BucketV1Request represents a custom type struct
type RestoreAwsS3BucketV1Request struct {
    // The parameters for initiating an S3 bucket restore.
    Source     *RestoreS3BucketSource `json:"source"`
    // The destination where the S3 bucket will be restored.
    Target     *RestoreS3BucketTarget `json:"target"`
    // Type to be used during restore. If not specified, the default value is "Rollback".
    // "Rollback" and "Undo delete marker" value will be deprecated. Use "rollback" and "undo_delete_marker" respectively.
    ClumioType *string                `json:"type"`
}

// PreviewAwsS3BucketV1Request represents a custom type struct
type PreviewAwsS3BucketV1Request struct {
    // Search for or restore only objects that pass the source object filter.
    ObjectFilters           *SourceObjectFiltersV2                        `json:"object_filters"`
    // This field is required when the request type is 'Rollback'. The parameters for initiating a point in time restore.
    PointInTimeRecovery     *RestoreS3BucketSourcePitrOptions             `json:"point_in_time_recovery"`
    // Type to be used during preview. If not specified, the default value is "Rollback".
    // "Rollback" and "Undo delete marker" value will be deprecated. Use "rollback" and "undo_delete_marker" respectively.
    ClumioType              *string                                       `json:"type"`
    // This field is required when the request type is 'Undo delete marker'.
    UndoDeleteMarkerOptions *RestoreS3BucketSourceUndoDeleteMarkerOptions `json:"undo_delete_marker_options"`
}

// RestoreFilesV1Request represents a custom type struct
type RestoreFilesV1Request struct {
    // The files to be restored and from which backup they are to be restored from.
    Source *FileRestoreSource `json:"source"`
    // The destination information for the file level restore, representing the access option
    // for the restored file. Users can access the restored file by direct download or by
    // email. The direct download (`direct_download`) option allows users to directly download
    // the restored file from the Clumio UI. The email (`email`) option allows users to access
    // the restored file using a downloadable link they receive by email. If a target is not
    // specified, then `target` defaults to `direct_download`.
    Target *FileRestoreTarget `json:"target"`
}

// DownloadSharedFileV1Request represents a custom type struct
type DownloadSharedFileV1Request struct {
    // The download link that was sent to you by email. To get the download link,
    // open the email message, click "Download File" to launch the Clumio "Access
    // Requested File" page, and copy the URL.
    EmailLink *string `json:"email_link"`
    // The passcode used to access the restored file. Obtain the passcode from the
    // user who generated the restored file.
    Passcode  *string `json:"passcode"`
}

// ShareRestoredFileV1Request represents a custom type struct
type ShareRestoredFileV1Request struct {
    // The email address of the user who will receive the download link to the restored
    // file.
    EmailAddress *string `json:"email_address"`
    // The optional message sent as part of the email.
    Message      *string `json:"message"`
}

// RestoreProtectionGroupV1Request represents a custom type struct
type RestoreProtectionGroupV1Request struct {
    // The parameters for initiating a protection group restore from a backup.
    Source *ProtectionGroupRestoreSource `json:"source"`
    // The destination where the protection group will be restored.
    Target *ProtectionGroupRestoreTarget `json:"target"`
}

// CreateProtectionGroupInstantAccessEndpointV1Request represents a custom type struct
type CreateProtectionGroupInstantAccessEndpointV1Request struct {
    // The time that this endpoint expires at in RFC-3339 format.
    ExpiryTimestamp *string                `json:"expiry_timestamp"`
    // The user-assigned name of the S3 instant access endpoint.
    Name            *string                `json:"name"`
    // The parameters for creating a S3 instant access endpoint.
    Source          *S3InstantAccessSource `json:"source"`
}

// CostEstimatesProtectionGroupInstantAccessEndpointV1Request represents a custom type struct
type CostEstimatesProtectionGroupInstantAccessEndpointV1Request struct {
    // The Clumio-assigned ID of the protection group S3 asset backup or protection group backup to
    // be restored. Use the endpoints
    // [GET /backups/protection-groups/s3-assets](#operation/list-backup-protection-group-s3-assets)
    // for protection group S3 asset backup, and
    // [GET /backups/protection-groups](#operation/list-backup-protection-groups)
    // for protection group backups to fetch valid values. 
    // Note that only one of either `backup_id` or `pitr` must be provided.
    BackupId                 *string                           `json:"backup_id"`
    // Whether to wait for the operation to complete.
    IsSync                   *bool                             `json:"is_sync"`
    // Search for or restore only objects that pass the source object filter.
    ObjectFilters            *SourceObjectFilters              `json:"object_filters"`
    // The parameters to initiate a point-in-time creation of S3 instant access endpoint.
    // Note that only one of either `backup_id` or `pitr` must be provided.
    Pitr                     *S3InstantAccessSourcePitrOptions `json:"pitr"`
    // Clumio-assigned ID of protection group S3 asset, representing the
    // bucket within the protection group to restore from. Use the
    // [GET /datasources/protection-groups/s3-assets](#operation/list-protection-group-s3-assets)
    // endpoint to fetch valid values.
    ProtectionGroupS3AssetId *string                           `json:"protection_group_s3_asset_id"`
}

// UpdateProtectionGroupInstantAccessEndpointV1Request represents a custom type struct
type UpdateProtectionGroupInstantAccessEndpointV1Request struct {
    // The time that this endpoint expires, in RFC-3339 format. This will revert to default if no
    // state passed.
    ExpiryTimestamp *string `json:"expiry_timestamp"`
    // The user-assigned name of the S3 instant access endpoint. This will be removed if left empty.
    Name            *string `json:"name"`
}

// AddProtectionGroupInstantAccessEndpointRoleV1Request represents a custom type struct
type AddProtectionGroupInstantAccessEndpointRoleV1Request struct {
    // Allow the addition of a role from an external account. This requires a feature flag to be enabled, contact support@clumio.com.
    IsAllowExternalAccount *bool   `json:"is_allow_external_account"`
    // Descriptive alias of the IAM role.
    RoleAlias              *string `json:"role_alias"`
    // ARN of the IAM role to allow access the endpoint. The role must be accessible from AWS account
    // registered with Clumio.
    RoleArn                *string `json:"role_arn"`
}

// UpdateProtectionGroupInstantAccessEndpointRoleV1Request represents a custom type struct
type UpdateProtectionGroupInstantAccessEndpointRoleV1Request struct {
    // The updated descriptive alias of the IAM role. The current alias will be retained if
    // empty updated_role_alias is passed.
    UpdatedRoleAlias *string `json:"updated_role_alias"`
    // The updated ARN of the IAM role to allow access to the endpoint. The role must be
    // accessible from an AWS account registered with Clumio. The current ARN will be retained
    // if empty updated_role_arn is passed.
    UpdatedRoleArn   *string `json:"updated_role_arn"`
}

// RestoreProtectionGroupS3AssetV1Request represents a custom type struct
type RestoreProtectionGroupS3AssetV1Request struct {
    // The parameters for initiating a protection group S3 asset restore
    // or creation of an instant access endpoint from a backup.
    Source *ProtectionGroupS3AssetRestoreSource `json:"source"`
    // The destination where the protection group will be restored.
    Target *ProtectionGroupRestoreTarget        `json:"target"`
}

// PreviewProtectionGroupS3AssetV1Request represents a custom type struct
type PreviewProtectionGroupS3AssetV1Request struct {
    // Backup ID.
    BackupId           *string              `json:"backup_id"`
    // Response type to be sync
    IsSyncPreview      *bool                `json:"is_sync_preview"`
    // Search for or restore only objects that pass the source object filter.
    ObjectFilters      *SourceObjectFilters `json:"object_filters"`
    // The end timestamp of the period within which objects are to be restored, in RFC-3339
    // format. Clumio searches for objects modified before the given time. If `pitr_end_timestamp`
    // is empty, Clumio searches for objects modified up to the current time of the restore request.
    //  If `pitr_end_timestamp` is given without `pitr_start_timestamp`,
    // it is the same as point in time preview.
    PitrEndTimestamp   *string              `json:"pitr_end_timestamp"`
    // The start timestamp of the period within which objects are to be restored, in RFC-3339
    // format. Clumio searches for objects modified since the given time. If `pitr_start_timestamp`
    // is empty, Clumio searches for objects from the beginning of the first backup.
    PitrStartTimestamp *string              `json:"pitr_start_timestamp"`
}

// PreviewProtectionGroupV1Request represents a custom type struct
type PreviewProtectionGroupV1Request struct {
    // The Clumio-assigned ID of the protection group backup to be restored. Use the
    // [GET /backups/protection-groups](#operation/list-backup-protection-groups)
    // endpoint to fetch valid values.
    BackupId                  *string              `json:"backup_id"`
    // Whether to wait for the preview task.
    IsSyncPreview             *bool                `json:"is_sync_preview"`
    // Search for or restore only objects that pass the source object filter.
    ObjectFilters             *SourceObjectFilters `json:"object_filters"`
    // The end timestamp of the period within which objects are to be restored, in RFC-3339
    // format. Clumio searches for objects modified before the given time. If `pitr_end_timestamp`
    // is empty, Clumio searches for objects modified up to the current time of the restore request.
    //  If `pitr_end_timestamp` is given without `pitr_start_timestamp`,
    // it is the same as point in time preview.
    PitrEndTimestamp          *string              `json:"pitr_end_timestamp"`
    // The start timestamp of the period within which objects are to be restored, in RFC-3339
    // format. Clumio searches for objects modified since the given time. If `pitr_start_timestamp`
    // is empty, Clumio searches for objects from the beginning of the first backup.
    PitrStartTimestamp        *string              `json:"pitr_start_timestamp"`
    // A list of Clumio-assigned IDs of protection group S3 assets, representing the
    // buckets within the protection group to restore from. Use the
    // [GET /datasources/protection-groups/s3-assets](#operation/list-protection-group-s3-assets)
    // endpoint to fetch valid values.
    ProtectionGroupS3AssetIds []*string            `json:"protection_group_s3_asset_ids"`
}

// RestoreProtectionGroupS3ObjectsV1Request represents a custom type struct
type RestoreProtectionGroupS3ObjectsV1Request struct {
    // Object defines one object to restore
    Source []*Object                     `json:"source"`
    // The destination where the protection group will be restored.
    Target *ProtectionGroupRestoreTarget `json:"target"`
}

// UpdateAutoUserProvisioningSettingV1Request represents a custom type struct
type UpdateAutoUserProvisioningSettingV1Request struct {
    // Whether auto user provisioning is enabled or not.
    IsEnabled *bool `json:"is_enabled"`
}

// CreateAutoUserProvisioningRuleV1Request represents a custom type struct
type CreateAutoUserProvisioningRuleV1Request struct {
    // The following table describes the possible conditions for a rule.
    // 
    // +--------------------------+-------------------------+-------------------------+
    // |     Group Selection      |     Rule Condition      |       Description       |
    // +==========================+=========================+=========================+
    // | This group               |                         | User must belong to the |
    // |                          |                         | specified group.        |
    // |                          | {"user.groups":{"$eq":" |                         |
    // |                          | Admin"}}                |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // | ANY of these groups      |                         | User must belong to at  |
    // |                          |                         | least one of the        |
    // |                          | {"user.groups":{"$in":[ | specified groups.       |
    // |                          | "Admin", "Eng",         |                         |
    // |                          | "Sales"]}}              |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // | ALL of these groups      |                         | User must belong to all |
    // |                          |                         | the specified groups.   |
    // |                          | {"user.groups":{"$all": |                         |
    // |                          | ["Admin", "Eng",        |                         |
    // |                          | "Sales"]}}              |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // | Group CONTAINS this      |                         | User's group must       |
    // | keyword                  |                         | contain the specified   |
    // |                          | {"user.groups":{"$conta | keyword.                |
    // |                          | ins":{"$in":["Admin"]}} |                         |
    // |                          | }                       |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // | Group CONTAINS ANY of    |                         | User's group must       |
    // | these keywords           |                         | contain at least one of |
    // |                          | {"user.groups":{"$conta | the specified keywords. |
    // |                          | ins":{"$in":["Admin",   |                         |
    // |                          | "Eng", "Sales"]}}}      |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // | Group CONTAINS ALL of    |                         | User's group must       |
    // | these keywords           |                         | contain all the         |
    // |                          | {"user.groups":{"$conta | specified keywords.     |
    // |                          | ins":{"$all":["Admin",  |                         |
    // |                          | "Eng", "Sales"]}}}      |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // 
    Condition *string        `json:"condition"`
    // Unique name assigned to the rule.
    Name      *string        `json:"name"`
    // Specifies the role and the organizational units to be assigned to the user subject to the rule criteria.
    Provision *RuleProvision `json:"provision"`
}

// UpdateAutoUserProvisioningRuleV1Request represents a custom type struct
type UpdateAutoUserProvisioningRuleV1Request struct {
    // The following table describes the possible conditions for a rule.
    // 
    // +--------------------------+-------------------------+-------------------------+
    // |     Group Selection      |     Rule Condition      |       Description       |
    // +==========================+=========================+=========================+
    // | This group               |                         | User must belong to the |
    // |                          |                         | specified group.        |
    // |                          | {"user.groups":{"$eq":" |                         |
    // |                          | Admin"}}                |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // | ANY of these groups      |                         | User must belong to at  |
    // |                          |                         | least one of the        |
    // |                          | {"user.groups":{"$in":[ | specified groups.       |
    // |                          | "Admin", "Eng",         |                         |
    // |                          | "Sales"]}}              |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // | ALL of these groups      |                         | User must belong to all |
    // |                          |                         | the specified groups.   |
    // |                          | {"user.groups":{"$all": |                         |
    // |                          | ["Admin", "Eng",        |                         |
    // |                          | "Sales"]}}              |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // | Group CONTAINS this      |                         | User's group must       |
    // | keyword                  |                         | contain the specified   |
    // |                          | {"user.groups":{"$conta | keyword.                |
    // |                          | ins":{"$in":["Admin"]}} |                         |
    // |                          | }                       |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // | Group CONTAINS ANY of    |                         | User's group must       |
    // | these keywords           |                         | contain at least one of |
    // |                          | {"user.groups":{"$conta | the specified keywords. |
    // |                          | ins":{"$in":["Admin",   |                         |
    // |                          | "Eng", "Sales"]}}}      |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // | Group CONTAINS ALL of    |                         | User's group must       |
    // | these keywords           |                         | contain all the         |
    // |                          | {"user.groups":{"$conta | specified keywords.     |
    // |                          | ins":{"$all":["Admin",  |                         |
    // |                          | "Eng", "Sales"]}}}      |                         |
    // |                          |                         |                         |
    // |                          |                         |                         |
    // +--------------------------+-------------------------+-------------------------+
    // 
    Condition *string        `json:"condition"`
    // Unique name assigned to the rule.
    Name      *string        `json:"name"`
    // Specifies the role and the organizational units to be assigned to the user subject to the rule criteria.
    Provision *RuleProvision `json:"provision"`
}

// UpdateGeneralSettingsV2Request represents a custom type struct
type UpdateGeneralSettingsV2Request struct {
    // The length of time before a user is logged out of the Clumio system due to inactivity. Measured in seconds.
    // The valid range is between 600 seconds (10 minutes) and 3600 seconds (60 minutes).
    // If not configured, the value defaults to 900 seconds (15 minutes).
    AutoLogoutDuration           *int64              `json:"auto_logout_duration"`
    // The designated range of IP addresses that are allowed to access the Clumio REST API.
    // API requests that originate from outside this list will be blocked.
    // The IP address of the server from which this request is being made must be in this list; otherwise, the request will fail.
    // Set the parameter to individual IP addresses and/or a range of IP addresses in CIDR notation.
    // For example, ["193.168.1.0/24", "193.172.1.1"].
    // If not configured, the value defaults to ["0.0.0.0/0"] meaning all addresses will be allowed.
    IpAllowlist                  []*string           `json:"ip_allowlist"`
    // The grouping criteria for each datasource type.
    // These can only be edited for datasource types which do not have any
    // organizational units configured.
    // Deprecated: This struct is deprecated and will be removed in the future.
    // It is being kept for backward compatibility.
    OrganizationalUnitDataGroups *OUGroupingCriteria `json:"organizational_unit_data_groups"`
    // The length of time a user password is valid before it must be changed. Measured in seconds.
    // The valid range is between 2592000 seconds (30 days) and 15552000 seconds (180 days).
    // If not configured, the value defaults to 7776000 seconds (90 days).
    PasswordExpirationDuration   *int64              `json:"password_expiration_duration"`
}

// UpdateTaskV1Request represents a custom type struct
type UpdateTaskV1Request struct {
    // The task status. Set this parameter to `aborting` to abort a task
    // that is in queue ("queued") or in progress ("in_progress").
    // Tasks with other statuses cannot be aborted.
    Status *string `json:"status"`
}

// CreateUserV2Request represents a custom type struct
type CreateUserV2Request struct {
    // The organizational units assigned to the user, with the specified role.
    AccessControlConfiguration []*RoleForOrganizationalUnits `json:"access_control_configuration"`
    // The email address of the user to be added to Clumio.
    Email                      *string                       `json:"email"`
    // The full name of the user to be added to Clumio. For example, type the user's first name and last name.
    // The name displays on the User Management screen and in the body of the email invitation.
    FullName                   *string                       `json:"full_name"`
}

// CreateUserV1Request represents a custom type struct
type CreateUserV1Request struct {
    // The Clumio-assigned ID of the role to assign to the user.
    // The available roles can be retrieved via the [GET /roles](#operation/list-roles) API.
    // When not set, the role is determined as follows
    // 
    // +-------------+---------------------------+----------------+
    // |   Inviter   |       Assigned OUs        | Resulting Role |
    // +=============+===========================+================+
    // | Super Admin | Global OU is assigned     | Super Admin    |
    // +-------------+---------------------------+----------------+
    // | Super Admin | Global OU is not assigned | OU Admin       |
    // +-------------+---------------------------+----------------+
    // | OU Admin    | Any                       | OU Admin       |
    // +-------------+---------------------------+----------------+
    // 
    AssignedRole          *string   `json:"assigned_role"`
    // The email address of the user to be added to Clumio.
    Email                 *string   `json:"email"`
    // The full name of the user to be added to Clumio. For example, type the user's first name and last name.
    // The name displays on the User Management screen and in the body of the email invitation.
    FullName              *string   `json:"full_name"`
    // The Clumio-assigned IDs of the organizational units to be assigned to the user.
    OrganizationalUnitIds []*string `json:"organizational_unit_ids"`
}

// ChangePasswordV2Request represents a custom type struct
type ChangePasswordV2Request struct {
    // The user's current password.
    CurrentPassword *string `json:"current_password"`
    // The new password that is to replace the user's current password. Passwords must be between 14 and 64 characters
    // and include the following: one uppercase character, one lowercase character, one number, and one special character.
    // Spaces are not allowed.
    NewPassword     *string `json:"new_password"`
}

// UpdateUserProfileV2Request represents a custom type struct
type UpdateUserProfileV2Request struct {
    // The full name of the user that is to replace the existing full name.
    // For example, enter the user's first name and last name.
    FullName *string `json:"full_name"`
}

// UpdateUserProfileV1Request represents a custom type struct
type UpdateUserProfileV1Request struct {
    // The full name of the user that is to replace the existing full name.
    // For example, enter the user's first name and last name.
    FullName *string `json:"full_name"`
}

// UpdateUserV2Request represents a custom type struct
type UpdateUserV2Request struct {
    // Updates to the organizational units along with the role assigned to the user.
    AccessControlConfigurationUpdates *EntityGroupAssignmentUpdates `json:"access_control_configuration_updates"`
    // The full name of the user that is to replace the existing full name.
    // For example, enter the user's first name and last name.
    FullName                          *string                       `json:"full_name"`
    // If `true`, enables a user who has been disabled from Clumio,
    // returning the user to its previous "Activated" or "Invited" status. If `false`, disables a user from Clumio.
    // The user will not be able log in to Clumio until another Clumio user re-enables the account.
    IsEnabled                         *bool                         `json:"is_enabled"`
}

// UpdateUserV1Request represents a custom type struct
type UpdateUserV1Request struct {
    // Updates the role assigned to the user.
    // The available roles can be retrieved via the [GET /roles](#operation/list-roles) API.
    AssignedRole                        *string                         `json:"assigned_role"`
    // The full name of the user that is to replace the existing full name.
    // For example, enter the user's first name and last name.
    FullName                            *string                         `json:"full_name"`
    // If `true`, enables a user who has been disabled from Clumio,
    // returning the user to its previous "Activated" or "Invited" status. If `false`, disables a user from Clumio.
    // The user will not be able log in to Clumio until another Clumio user re-enables the account.
    IsEnabled                           *bool                           `json:"is_enabled"`
    // Updates to the organizational unit assignments.
    OrganizationalUnitAssignmentUpdates *EntityGroupAssignmentUpdatesV1 `json:"organizational_unit_assignment_updates"`
}

// ChangePasswordV1Request represents a custom type struct
type ChangePasswordV1Request struct {
    // The user's current password.
    CurrentPassword *string `json:"current_password"`
    // The new password that is to replace the user's current password. Passwords must be between 14 and 64 characters
    // and include the following: one uppercase character, one lowercase character, one number, and one special character.
    // Spaces are not allowed.
    NewPassword     *string `json:"new_password"`
}

// CreateWalletV1Request represents a custom type struct
type CreateWalletV1Request struct {
    // AWS Account ID to associate with the wallet.
    AccountNativeId *string `json:"account_native_id"`
    // The AWS Region to associate with the wallet.
    AwsRegion       *string `json:"aws_region"`
}

// PostProcessKmsV1Request represents a custom type struct.
// The body of the request.
type PostProcessKmsV1Request struct {
    // The AWS-assigned ID of the account associated with the connection.
    AccountNativeId       *string `json:"account_native_id"`
    // The AWS region associated with the connection. For example, `us-east-1`.
    AwsRegion             *string `json:"aws_region"`
    // Whether the CMK was created or an existing CMK was used.
    CreatedMultiRegionCmk *bool   `json:"created_multi_region_cmk"`
    // Role arn to be assumed before accessing ClumioRole in customer account.
    IntermediateRoleArn   *string `json:"intermediate_role_arn"`
    // The multi-region CMK Key ID.
    MultiRegionCmkKeyId   *string `json:"multi_region_cmk_key_id"`
    // Indicates whether this is a Create, Update or Delete request.
    RequestType           *string `json:"request_type"`
    // The ARN of the role.
    RoleArn               *string `json:"role_arn"`
    // The external ID to use with the role.
    RoleExternalId        *string `json:"role_external_id"`
    // The ID of the role.
    RoleId                *string `json:"role_id"`
    // The 36-character Clumio AWS integration ID token used to identify the
    // installation of the CloudFormation/Terraform template on the account.
    Token                 *string `json:"token"`
    // The cloudformation/terraform template version used.
    Version               *uint64 `json:"version"`
}
