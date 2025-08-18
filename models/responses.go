// Copyright (c) 2021 Clumio All Rights Reserved

// Package models has the structs representing responses
package models

// CreateOrganizationalUnitResponseWrapper represents a custom type struct wrapper for different success responses
type CreateOrganizationalUnitResponseWrapper struct {
    // HTTP status code of response
    StatusCode int
    // CreateOrganizationalUnitNoTaskResponse represents Success response of status code Http200
    Http200 *CreateOrganizationalUnitNoTaskResponse 
    // CreateOrganizationalUnitResponse represents Success response of status code Http202
    Http202 *CreateOrganizationalUnitResponse 
}

// CreateOrganizationalUnitV1ResponseWrapper represents a custom type struct wrapper for different success responses
type CreateOrganizationalUnitV1ResponseWrapper struct {
    // HTTP status code of response
    StatusCode int
    // CreateOrganizationalUnitNoTaskResponseV1 represents Success response of status code Http200
    Http200 *CreateOrganizationalUnitNoTaskResponseV1 
    // CreateOrganizationalUnitResponseV1 represents Success response of status code Http202
    Http202 *CreateOrganizationalUnitResponseV1 
}

// PatchOrganizationalUnitResponseWrapper represents a custom type struct wrapper for different success responses
type PatchOrganizationalUnitResponseWrapper struct {
    // HTTP status code of response
    StatusCode int
    // PatchOrganizationalUnitNoTaskResponse represents Success response of status code Http200
    Http200 *PatchOrganizationalUnitNoTaskResponse 
    // PatchOrganizationalUnitResponse represents Success response of status code Http202
    Http202 *PatchOrganizationalUnitResponse 
}

// PatchOrganizationalUnitV1ResponseWrapper represents a custom type struct wrapper for different success responses
type PatchOrganizationalUnitV1ResponseWrapper struct {
    // HTTP status code of response
    StatusCode int
    // PatchOrganizationalUnitNoTaskResponseV1 represents Success response of status code Http200
    Http200 *PatchOrganizationalUnitNoTaskResponseV1 
    // PatchOrganizationalUnitResponseV1 represents Success response of status code Http202
    Http202 *PatchOrganizationalUnitResponseV1 
}

// RestoreRecordsAwsDynamodbTableResponseWrapper represents a custom type struct wrapper for different success responses
type RestoreRecordsAwsDynamodbTableResponseWrapper struct {
    // HTTP status code of response
    StatusCode int
    // RestoreRecordsResponseSync represents Success response of status code Http200
    Http200 *RestoreRecordsResponseSync 
    // RestoreRecordsResponseAsync represents Success response of status code Http202
    Http202 *RestoreRecordsResponseAsync 
}

// RestoreRdsRecordResponseWrapper represents a custom type struct wrapper for different success responses
type RestoreRdsRecordResponseWrapper struct {
    // HTTP status code of response
    StatusCode int
    // RestoreRecordPreviewResponse represents Success response of status code Http200
    Http200 *RestoreRecordPreviewResponse 
    // RestoreRecordResponse represents Success response of status code Http202
    Http202 *RestoreRecordResponse 
}

// CostEstimatesProtectionGroupInstantAccessEndpointResponseWrapper represents a custom type struct wrapper for different success responses
type CostEstimatesProtectionGroupInstantAccessEndpointResponseWrapper struct {
    // HTTP status code of response
    StatusCode int
    // EstimateCostS3InstantAccessEndpointSyncResponse represents Success response of status code Http200
    Http200 *EstimateCostS3InstantAccessEndpointSyncResponse 
    // EstimateCostS3InstantAccessEndpointAsyncResponse represents Success response of status code Http202
    Http202 *EstimateCostS3InstantAccessEndpointAsyncResponse 
}

// PreviewProtectionGroupS3AssetResponseWrapper represents a custom type struct wrapper for different success responses
type PreviewProtectionGroupS3AssetResponseWrapper struct {
    // HTTP status code of response
    StatusCode int
    // PreviewProtectionGroupS3AssetSyncResponse represents Success response of status code Http200
    Http200 *PreviewProtectionGroupS3AssetSyncResponse 
    // PreviewProtectionGroupS3AssetAsyncResponse represents Success response of status code Http202
    Http202 *PreviewProtectionGroupS3AssetAsyncResponse 
}

// PreviewProtectionGroupResponseWrapper represents a custom type struct wrapper for different success responses
type PreviewProtectionGroupResponseWrapper struct {
    // HTTP status code of response
    StatusCode int
    // PreviewProtectionGroupSyncResponse represents Success response of status code Http200
    Http200 *PreviewProtectionGroupSyncResponse 
    // PreviewProtectionGroupAsyncResponse represents Success response of status code Http202
    Http202 *PreviewProtectionGroupAsyncResponse 
}

// AddBucketToProtectionGroupResponse represents a custom type struct for Success
type AddBucketToProtectionGroupResponse struct {
    // The AWS-assigned ID of the account associated with the DynamoDB table.
    AccountNativeId               *string           `json:"account_native_id"`
    // Whether this bucket was added to this protection group by the bucket rule
    AddedByBucketRule             *bool             `json:"added_by_bucket_rule"`
    // Whether this bucket was added to this protection group by the user
    AddedByUser                   *bool             `json:"added_by_user"`
    // The AWS region associated with the DynamoDB table.
    AwsRegion                     *string           `json:"aws_region"`
    // The backup target AWS region associated with the protection group S3 asset.
    BackupTargetAwsRegion         *string           `json:"backup_target_aws_region"`
    // BackupTierStat
    BackupTierStats               []*BackupTierStat `json:"backup_tier_stats"`
    // The Clumio-assigned ID of the bucket
    BucketId                      *string           `json:"bucket_id"`
    // The name of the bucket
    BucketName                    *string           `json:"bucket_name"`
    // Creation time of the protection group in RFC-3339 format.
    CreatedTimestamp              *string           `json:"created_timestamp"`
    // The Clumio-assigned ID of the AWS environment associated with the protection group.
    EnvironmentId                 *string           `json:"environment_id"`
    // The Clumio-assigned ID of the protection group
    GroupId                       *string           `json:"group_id"`
    // The name of the protection group
    GroupName                     *string           `json:"group_name"`
    // The Clumio-assigned ID that represents the bucket within the protection group.
    Id                            *string           `json:"id"`
    // Determines whether the protection group bucket has been deleted
    IsDeleted                     *bool             `json:"is_deleted"`
    // Time of the last backup in RFC-3339 format.
    LastBackupTimestamp           *string           `json:"last_backup_timestamp"`
    // Time of the last successful continuous backup in RFC-3339 format.
    LastContinuousBackupTimestamp *string           `json:"last_continuous_backup_timestamp"`
    // The Clumio-assigned ID of the organizational unit associated with the protection group.
    OrganizationalUnitId          *string           `json:"organizational_unit_id"`
    // Cumulative count of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpObjectCount      *int64            `json:"total_backed_up_object_count"`
    // Cumulative size of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpSizeBytes        *int64            `json:"total_backed_up_size_bytes"`
    // The unsupported reason for the S3 bucket.
    UnsupportedReason             *string           `json:"unsupported_reason"`
}

// AddS3InstantAccessEndpointRoleResponse represents a custom type struct for Success
type AddS3InstantAccessEndpointRoleResponse struct {
    // Embedded responses related to the resource.
    Embedded *S3InstantAccessEndpointEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *S3InstantAccessEndpointLinks    `json:"_links"`
    // The issued ID to the added role
    RoleId   *string                          `json:"role_id"`
}

// ChangePasswordResponse represents a custom type struct for Success
type ChangePasswordResponse struct {
    // HateoasCommonLinks are the common fields for HATEOAS response.
    Links *HateoasCommonLinks `json:"_links"`
}

// CreateAWSConnectionResponse represents a custom type struct for Success
type CreateAWSConnectionResponse struct {
    // URLs to pages related to the resource.
    Links                      *AWSConnectionLinks      `json:"_links"`
    // The alias given to the account on AWS.
    AccountName                *string                  `json:"account_name"`
    // The AWS-assigned ID of the account associated with the connection.
    AccountNativeId            *string                  `json:"account_native_id"`
    // The AWS region associated with the connection. For example, `us-east-1`.
    AwsRegion                  *string                  `json:"aws_region"`
    // AWS account ID of the Clumio control plane.
    ClumioAwsAccountId         *string                  `json:"clumio_aws_account_id"`
    // AWS region of the Clumio control plane
    ClumioAwsRegion            *string                  `json:"clumio_aws_region"`
    // The consolidated configuration of the Clumio Cloud Protect and Clumio Cloud Discover products for this connection.
    // If this connection is deprecated to use unconsolidated configuration, then this field has a
    // value of `null`.
    Config                     *ConsolidatedConfig      `json:"config"`
    // Clumio assigned ID of the associated connection group.
    ConnectionGroupId          *string                  `json:"connection_group_id"`
    // Management status of connection.
    ConnectionManagementStatus *string                  `json:"connection_management_status"`
    // The status of the connection considering all the deployments made for it.
    ConnectionStatus           *string                  `json:"connection_status"`
    // The timestamp of when the connection was created.
    CreatedTimestamp           *string                  `json:"created_timestamp"`
    // AWS account ID of the data plane for the connection.
    DataPlaneAccountId         *string                  `json:"data_plane_account_id"`
    // The deployment method with which the currently active connection was established.
    DeploymentType             *string                  `json:"deployment_type"`
    // The user-provided description for this connection.
    Description                *string                  `json:"description"`
    // The configuration of the Clumio Discover product for this connection.
    // If this connection is not configured for Clumio Discover, then this field has a
    // value of `null`.
    Discover                   *DiscoverConfig          `json:"discover"`
    // Clumio assigned external ID of the connection or of the associated connection group.
    ExternalId                 *string                  `json:"external_id"`
    // The Clumio-assigned ID of the connection.
    Id                         *string                  `json:"id"`
    // Status denoting whether Ingestion has started for the connection.
    // Valid values are "initial", "in_progress", "failed", "completed".
    IngestionStatus            *string                  `json:"ingestion_status"`
    // K8S Namespace
    Namespace                  *string                  `json:"namespace"`
    // The Clumio-assigned ID of the organizational unit associated with the
    // AWS environment. If this parameter is not provided, then the value
    // defaults to the first organizational unit assigned to the requesting
    // user. For more information about organizational units, refer to the
    // Organizational-Units documentation.
    OrganizationalUnitId       *string                  `json:"organizational_unit_id"`
    // The configuration of the Clumio Cloud Protect product for this connection.
    // If this connection is not configured for Clumio Cloud Protect, then this field has a
    // value of `null`.
    Protect                    *ProtectConfig           `json:"protect"`
    // The asset types enabled for protect.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS", "IcebergOnGlue", "IcebergOnS3Tables"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    ProtectAssetTypesEnabled   []*string                `json:"protect_asset_types_enabled"`
    // TODO: Add struct field description
    Resources                  *ConnectionResourcesResp `json:"resources"`
    // The Amazon Resource Name of the stale CloudFormation stack when the connection was migrated to connection groups.
    // NOTE - This has to be removed from AWS as well to delete the connection completely.
    RetiredStackArn            *string                  `json:"retired_stack_arn"`
    // The services to be enabled for this configuration. Valid values are
    // ["discover"], ["discover", "protect"]. This is only set when the
    // registration is created, the enabled services are obtained directly from
    // the installed template after that. (Deprecated as all connections will have
    // both discover and protect enabled)
    ServicesEnabled            []*string                `json:"services_enabled"`
    // The Amazon Resource Name of the installed and active CloudFormation stack(if any) in AWS.
    StackArn                   *string                  `json:"stack_arn"`
    // The name given to the installed and active CloudFormation stack(if any) in AWS.
    StackName                  *string                  `json:"stack_name"`
    // Status denoting whether Target Setup has started for the connection.
    // Valid values are "initial", "in_progress", "failed", "completed".
    TargetSetupStatus          *string                  `json:"target_setup_status"`
    // TODO: Add struct field description
    TemplatePermissionSet      *string                  `json:"template_permission_set"`
    // The 36-character Clumio AWS integration ID token used to identify the
    // installation of the CloudFormation template on the account. This value
    // will be pasted into the ClumioToken field when creating the
    // CloudFormation stack.
    Token                      *string                  `json:"token"`
}

// CreateAWSTemplateV2Response represents a custom type struct for Success
type CreateAWSTemplateV2Response struct {
    // URLs to pages related to the resource.
    Links                       *TemplateLinks           `json:"_links"`
    // The latest available URL for the template.
    CloudformationUrl           *string                  `json:"cloudformation_url"`
    // The configuration of the given template
    Config                      *TemplateConfigurationV2 `json:"config"`
    // The latest available URL for the deployable template.
    DeployableCloudformationUrl *string                  `json:"deployable_cloudformation_url"`
    // swagger: ignore
    GroupToken                  *string                  `json:"group_token"`
    // Categorised Resources, based on the generated template, to be created manually by the user
    Resources                   *CategorisedResources    `json:"resources"`
    // The latest available URL for the terraform template.
    TerraformUrl                *string                  `json:"terraform_url"`
}

// CreateAutoUserProvisioningRuleResponse represents a custom type struct for Success
type CreateAutoUserProvisioningRuleResponse struct {
    // Embedded responses related to the resource.
    Embedded  *AutoUserProvisioningRuleEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links     *AutoUserProvisioningRuleLinks    `json:"_links"`
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
    Condition *string                           `json:"condition"`
    // Unique name assigned to the rule.
    Name      *string                           `json:"name"`
    // Specifies the role and the organizational units to be assigned to the user subject to the rule criteria.
    Provision *RuleProvision                    `json:"provision"`
    // The Clumio-assigned ID of the rule.
    RuleId    *string                           `json:"rule_id"`
}

// CreateComplianceConfigurationResponse represents a custom type struct for Success
type CreateComplianceConfigurationResponse struct {
    // If the `embed` query parameter is set, displays the responses of the related resource,
    // as defined by the embeddable link specified.
    Embedded     interface{}                   `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *ComplianceConfigurationLinks `json:"_links"`
    // The RFC3339 format time when the report configuration was created.
    Created      *string                       `json:"created"`
    // The user-provided description of the compliance report configuration.
    Description  *string                       `json:"description"`
    // The unique identifier of the report configuration.
    Id           *string                       `json:"id"`
    // Most recent report run generated from the report configuration.
    LatestRun    *LatestRun                    `json:"latest_run"`
    // The user-provided name of the compliance report configuration.
    Name         *string                       `json:"name"`
    // Notification channels to send the generated report runs.
    Notification *NotificationSetting          `json:"notification"`
    // Filter and control parameters of compliance report.
    Parameter    *Parameter                    `json:"parameter"`
    // When the report will be generated and sent. If the schedule is not provided then a
    // default value will be used.
    Schedule     *ScheduleSetting              `json:"schedule"`
}

// CreateComplianceRunResponse represents a custom type struct for Success
type CreateComplianceRunResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded    `json:"_embedded"`
    // CreateComplianceRunHateoasLinks
    // URLs to pages related to the resource.
    Links    *CreateComplianceRunHateoasLinks `json:"_links"`
    // The Clumio-assigned ID of the task created for compliance report run generation.
    // The progress of the task can be monitored using the `GET /tasks/{task_id}` endpoint.
    TaskId   *string                          `json:"task_id"`
}

// CreateConnectionGroupResponse represents a custom type struct for Success
type CreateConnectionGroupResponse struct {
    // Embedded responses related to the resource.
    Embedded                 interface{}           `json:"_embedded"`
    // The ETag value.
    Etag                     *string               `json:"_etag"`
    // URLs to pages related to the resource.
    Links                    *ConnectionGroupLinks `json:"_links"`
    // The alias given to the associated account in AWS.
    AccountName              *string               `json:"account_name"`
    // The AWS-assigned IDs of the accounts associated with the Connection Group.
    AccountNativeIds         []*string             `json:"account_native_ids"`
    // List of asset types connected via the connection-group.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS", "IcebergOnGlue", "IcebergOnS3Tables"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    AssetTypesEnabled        []*string             `json:"asset_types_enabled"`
    // The AWS regions associated with the with the Connection Group.
    AwsRegions               []*string             `json:"aws_regions"`
    // The consolidated configuration of the Clumio Cloud Protect and Clumio Cloud Discover products for this connection.
    // If this connection is deprecated to use unconsolidated configuration, then this field has a
    // value of `null`.
    Config                   *ConsolidatedConfig   `json:"config"`
    // The timestamp of when the connection was created.
    CreatedTimestamp         *string               `json:"created_timestamp"`
    // Clumio's S3 URL that contains the template to create the required resources in the
    // given account(s) according to the request.
    DeploymentTemplateUrl    *string               `json:"deployment_template_url"`
    // User-provided description for this connection group.
    Description              *string               `json:"description"`
    // Clumio assigned external ID for the connection group, should be used while creating the AWS stack.
    ExternalId               *string               `json:"external_id"`
    // The Clumio-assigned ID of the Connection Group, should be used as the token while creating the stack in AWS.
    Id                       *string               `json:"id"`
    // The AWS Account IDs that are intended to be associated with the Connection Group.
    IntendedAccountNativeIds []*string             `json:"intended_account_native_ids"`
    // THe asset types that are intended to be connected via connection-group.
    IntendedAssetTypes       []*string             `json:"intended_asset_types"`
    // The AWS regions that are intended to be connected with the Connection Group.
    IntendedAwsRegions       []*string             `json:"intended_aws_regions"`
    // The master account which manages the connection-group's stack.
    MasterAwsAccountId       *string               `json:"master_aws_account_id"`
    // The master region which manages the connection-group's stack.
    MasterRegion             *string               `json:"master_region"`
    // Ongoing Operation of the deployed and active stack of ConnectionGroup.
    OngoingStackOperation    *string               `json:"ongoing_stack_operation"`
    // The Clumio-assigned ID of the organizational unit associated with the
    // AWS environment. If this parameter is not provided, then the value
    // defaults to the first organizational unit assigned to the requesting
    // user. For more information about organizational units, refer to the
    // Organizational-Units documentation.
    OrganizationalUnitId     *string               `json:"organizational_unit_id"`
    // The Amazon Resource Name of the installed CloudFormation stack in AWS.
    StackArn                 *string               `json:"stack_arn"`
    // The name given to the installed CloudFormation stack in AWS.
    StackName                *string               `json:"stack_name"`
    // The status of the Connection Group based on the stack in associated AWS account.
    Status                   *string               `json:"status"`
}

// CreateEC2MSSQLDatabaseRestoreResponse represents a custom type struct for Success
type CreateEC2MSSQLDatabaseRestoreResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded               `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *CreateEC2MSSQLDatabaseRestoreResponseLinks `json:"_links"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                                     `json:"task_id"`
}

// CreateOrganizationalUnitNoTaskResponse represents a custom type struct for Success
type CreateOrganizationalUnitNoTaskResponse struct {
    // URLs to pages related to the resource.
    Links                     *OULinks        `json:"_links"`
    // Number of immediate children of the organizational unit.
    ChildrenCount             *int64          `json:"children_count"`
    // Datasource types configured in this organizational unit. Possible values include `aws`, `microsoft365`, `vmware`, or `mssql`.
    ConfiguredDatasourceTypes []*string       `json:"configured_datasource_types"`
    // List of all recursive descendant organizational units of this OU.
    DescendantIds             []*string       `json:"descendant_ids"`
    // A description of the organizational unit.
    Description               *string         `json:"description"`
    // The Clumio assigned ID of the organizational unit.
    Id                        *string         `json:"id"`
    // Unique name assigned to the organizational unit.
    Name                      *string         `json:"name"`
    // The Clumio assigned ID of the parent organizational unit.
    // The parent organizational unit contains the entities in this organizational unit and can update this organizational unit.
    // If this organizational unit is the global organizational unit, then this field has a value of `null`.
    ParentId                  *string         `json:"parent_id"`
    // Number of users to whom this organizational unit or any of its descendants have been assigned.
    UserCount                 *int64          `json:"user_count"`
    // A user along with a role.
    Users                     []*UserWithRole `json:"users"`
}

// CreateOrganizationalUnitNoTaskResponseV1 represents a custom type struct for Success
type CreateOrganizationalUnitNoTaskResponseV1 struct {
    // URLs to pages related to the resource.
    Links                     *OULinks  `json:"_links"`
    // Number of immediate children of the organizational unit.
    ChildrenCount             *int64    `json:"children_count"`
    // Datasource types configured in this organizational unit. Possible values include `aws`, `microsoft365`, `vmware`, or `mssql`.
    ConfiguredDatasourceTypes []*string `json:"configured_datasource_types"`
    // List of all recursive descendant organizational units of this OU.
    DescendantIds             []*string `json:"descendant_ids"`
    // A description of the organizational unit.
    Description               *string   `json:"description"`
    // The Clumio assigned ID of the organizational unit.
    Id                        *string   `json:"id"`
    // Unique name assigned to the organizational unit.
    Name                      *string   `json:"name"`
    // The Clumio assigned ID of the parent organizational unit.
    // The parent organizational unit contains the entities in this organizational unit and can update this organizational unit.
    // If this organizational unit is the global organizational unit, then this field has a value of `null`.
    ParentId                  *string   `json:"parent_id"`
    // Number of users to whom this organizational unit or any of its descendants have been assigned.
    UserCount                 *int64    `json:"user_count"`
    // Users IDs to whom the organizational unit has been assigned.
    // This attribute will be available when reading a single OU and not when listing OUs.
    Users                     []*string `json:"users"`
}

// CreateOrganizationalUnitResponse represents a custom type struct.
// Accepted
type CreateOrganizationalUnitResponse struct {
    // Embedded responses related to the resource.
    Embedded                  *EntityGroupEmbedded     `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                     *OrganizationalUnitLinks `json:"_links"`
    // Number of immediate children of the organizational unit.
    ChildrenCount             *int64                   `json:"children_count"`
    // Datasource types configured in this organizational unit. Possible values include `aws`, `microsoft365`, `vmware`, or `mssql`.
    ConfiguredDatasourceTypes []*string                `json:"configured_datasource_types"`
    // List of all recursive descendant organizational units of this OU.
    DescendantIds             []*string                `json:"descendant_ids"`
    // A description of the organizational unit.
    Description               *string                  `json:"description"`
    // The Clumio assigned ID of the organizational unit.
    Id                        *string                  `json:"id"`
    // Unique name assigned to the organizational unit.
    Name                      *string                  `json:"name"`
    // The Clumio assigned ID of the parent organizational unit.
    // The parent organizational unit contains the entities in this organizational unit and can update this organizational unit.
    // If this organizational unit is the global organizational unit, then this field has a value of `null`.
    ParentId                  *string                  `json:"parent_id"`
    // The Clumio-assigned ID of the task associated with this organizational unit.
    // The progress of the task can be monitored using the
    // [GET /tasks/{task_id}](#operation/read-task) endpoint.
    TaskId                    *string                  `json:"task_id"`
    // Number of users to whom this organizational unit or any of its descendants have been assigned.
    UserCount                 *int64                   `json:"user_count"`
    // A user along with a role.
    Users                     []*UserWithRole          `json:"users"`
}

// CreateOrganizationalUnitResponseV1 represents a custom type struct.
// Accepted
type CreateOrganizationalUnitResponseV1 struct {
    // Embedded responses related to the resource.
    Embedded                  *EntityGroupEmbedded     `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                     *OrganizationalUnitLinks `json:"_links"`
    // Number of immediate children of the organizational unit.
    ChildrenCount             *int64                   `json:"children_count"`
    // Datasource types configured in this organizational unit. Possible values include `aws`, `microsoft365`, `vmware`, or `mssql`.
    ConfiguredDatasourceTypes []*string                `json:"configured_datasource_types"`
    // List of all recursive descendant organizational units of this OU.
    DescendantIds             []*string                `json:"descendant_ids"`
    // A description of the organizational unit.
    Description               *string                  `json:"description"`
    // The Clumio assigned ID of the organizational unit.
    Id                        *string                  `json:"id"`
    // Unique name assigned to the organizational unit.
    Name                      *string                  `json:"name"`
    // The Clumio assigned ID of the parent organizational unit.
    // The parent organizational unit contains the entities in this organizational unit and can update this organizational unit.
    // If this organizational unit is the global organizational unit, then this field has a value of `null`.
    ParentId                  *string                  `json:"parent_id"`
    // The Clumio-assigned ID of the task associated with this organizational unit.
    // The progress of the task can be monitored using the
    // [GET /tasks/{task_id}](#operation/read-task) endpoint.
    TaskId                    *string                  `json:"task_id"`
    // Number of users to whom this organizational unit or any of its descendants have been assigned.
    UserCount                 *int64                   `json:"user_count"`
    // Users IDs to whom the organizational unit has been assigned.
    // This attribute will be available when reading a single OU and not when listing OUs.
    Users                     []*string                `json:"users"`
}

// CreatePolicyResponse represents a custom type struct for Success
type CreatePolicyResponse struct {
    // If the `embed` query parameter is set, displays the responses of the related resource,
    // as defined by the embeddable link specified.
    Embedded             *PolicyEmbedded    `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                *PolicyLinks       `json:"_links"`
    // The status of the policy.
    // Refer to the Policy Activation Status table for a complete list of policy statuses.
    ActivationStatus     *string            `json:"activation_status"`
    // The created time of the policy in unix time.
    CreatedTime          *int64             `json:"created_time"`
    // The Clumio-assigned ID of the policy.
    Id                   *string            `json:"id"`
    // The following table describes the possible lock statuses of a policy.
    // 
    // +----------+-------------------------------------------------------------------+
    // |  Status  |                            Description                            |
    // +==========+===================================================================+
    // | unlocked | Policies are unlocked until an update or deletion task is queued. |
    // +----------+-------------------------------------------------------------------+
    // | updating | During a policy edit, concurrent edits or deletion requests will  |
    // |          | be rejected.                                                      |
    // +----------+-------------------------------------------------------------------+
    // | deleting | During policy deletion, concurrent edits or deletion requests     |
    // |          | will be rejected.                                                 |
    // +----------+-------------------------------------------------------------------+
    // 
    LockStatus           *string            `json:"lock_status"`
    // The user-provided name of the policy.
    Name                 *string            `json:"name"`
    // TODO: Add struct field description
    Operations           []*PolicyOperation `json:"operations"`
    // The Clumio-assigned ID of the organizational unit associated with the policy.
    OrganizationalUnitId *string            `json:"organizational_unit_id"`
    // The policy-level timezone is deprecated, as the operation-level timezone should be used instead.
    // The timezone must be a valid location name from the IANA Time Zone database.
    // For instance, "America/New_York", "US/Central", "UTC".
    // Deprecated: true
    Timezone             *string            `json:"timezone"`
    // The updated time of the policy in unix time.
    UpdatedTime          *int64             `json:"updated_time"`
}

// CreateProtectionGroupResponse represents a custom type struct for Success
type CreateProtectionGroupResponse struct {
    // Embedded responses related to the resource.
    Embedded                       interface{}                  `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                          *ProtectionGroupVersionLinks `json:"_links"`
    // The backup target AWS region associated with the protection group, empty if
    // in-region or not configured.
    BackupTargetAwsRegion          *string                      `json:"backup_target_aws_region"`
    // BackupTierStat
    BackupTierStats                []*BackupTierStat            `json:"backup_tier_stats"`
    // Number of buckets
    BucketCount                    *int64                       `json:"bucket_count"`
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
    BucketRule                     *string                      `json:"bucket_rule"`
    // Creation time of the protection group in RFC-3339 format.
    CreatedTimestamp               *string                      `json:"created_timestamp"`
    // The user-assigned description of the protection group.
    Description                    *string                      `json:"description"`
    // The Clumio-assigned ID of the protection group.
    Id                             *string                      `json:"id"`
    // Whether the protection group already has a backup target configured by a policy, or
    // is open to be protected by an in-region or out-of-region S3 policy.
    IsBackupTargetRegionConfigured *bool                        `json:"is_backup_target_region_configured"`
    // Determines whether the protection group is active or has been deleted. Deleted protection
    // groups may be purged after some time once there are no active backups associated with it.
    IsDeleted                      *bool                        `json:"is_deleted"`
    // Time of the last backup in RFC-3339 format.
    LastBackupTimestamp            *string                      `json:"last_backup_timestamp"`
    // Time of the last successful continuous backup in RFC-3339 format.
    LastContinuousBackupTimestamp  *string                      `json:"last_continuous_backup_timestamp"`
    // Modified time of the protection group in RFC-3339 format.
    ModifiedTimestamp              *string                      `json:"modified_timestamp"`
    // The user-assigned name of the protection group.
    Name                           *string                      `json:"name"`
    // ObjectFilter
    // defines which objects will be backed up.
    ObjectFilter                   *ObjectFilter                `json:"object_filter"`
    // The Clumio-assigned ID of the organizational unit associated with the Protection Group.
    OrganizationalUnitId           *string                      `json:"organizational_unit_id"`
    // Cumulative count of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpObjectCount       *int64                       `json:"total_backed_up_object_count"`
    // Cumulative size of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpSizeBytes         *int64                       `json:"total_backed_up_size_bytes"`
    // Version of the protection group. The version number is incremented every time
    // a change is made to the protection group.
    Version                        *int64                       `json:"version"`
}

// CreateRdsResourceRestoreResponse represents a custom type struct for Success
type CreateRdsResourceRestoreResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded          `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *CreateRdsDatabaseRestoreResponseLinks `json:"_links"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the `GET /tasks/{task_id}` endpoint.
    TaskId   *string                                `json:"task_id"`
}

// CreateReportDownloadResponse represents a custom type struct for Success
type CreateReportDownloadResponse struct {
    // _links provides URLs to the related resources of a report CSV download
    Links  *ReportDownloadLinks `json:"_links"`
    // The Clumio-assigned ID of the task created by the request.
    // The progress of the task can be monitored using the
    // [`GET /tasks/{task_id}`](#operation/list-tasks) endpoint.
    TaskId *string              `json:"task_id"`
}

// CreateRuleResponse represents a custom type struct for Success
type CreateRuleResponse struct {
    // URLs to pages related to the resource.
    Links  *CreateRuleResponseLinks `json:"_links"`
    // A rule applies an action subject to a condition criteria.
    Rule   *Rule                    `json:"rule"`
    // The Clumio-assigned ID of the task generated by this request.
    TaskId *string                  `json:"task_id"`
}

// CreateS3InstantAccessEndpointResponse represents a custom type struct for Success
type CreateS3InstantAccessEndpointResponse struct {
    // Embedded responses related to the resource.
    Embedded *CreateS3InstantAccessEndpointResponseEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *CreateS3InstantAccessEndpointResponseLinks    `json:"_links"`
    // The Clumio-assigned ID of the S3 instant access endpoint.
    Id       *string                                        `json:"id"`
    // The Clumio-assigned ID of the task created by this instant access creation request.
    // The progress of the task can be monitored using the `GET /tasks/{task_id}` endpoint.
    TaskId   *string                                        `json:"task_id"`
}

// CreateUserResponse represents a custom type struct for Success
type CreateUserResponse struct {
    // Embedded responses related to the resource.
    Embedded                   *UserEmbedded                 `json:"_embedded"`
    // ETag value
    Etag                       *string                       `json:"_etag"`
    // URLs to pages related to the resource.
    Links                      *UserLinks                    `json:"_links"`
    // The organizational units assigned to the user, with the specified role.
    AccessControlConfiguration []*RoleForOrganizationalUnits `json:"access_control_configuration"`
    // The email address of the Clumio user.
    Email                      *string                       `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                   *string                       `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                         *string                       `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                    *string                       `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                *bool                         `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                  *bool                         `json:"is_enabled"`
    // The timestamp of when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp      *string                       `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount    *int64                        `json:"organizational_unit_count"`
}

// CreateUserResponseV1 represents a custom type struct for Success
type CreateUserResponseV1 struct {
    // Embedded responses related to the resource.
    Embedded                      *UserEmbeddedV1 `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                         *UserLinks      `json:"_links"`
    // The list of organizational unit IDs assigned to the user.
    // This attribute will be available when reading a single user and not when listing users.
    AssignedOrganizationalUnitIds []*string       `json:"assigned_organizational_unit_ids"`
    // Assigned Role for the user.
    AssignedRole                  *string         `json:"assigned_role"`
    // The email address of the Clumio user.
    Email                         *string         `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                      *string         `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                            *string         `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                       *string         `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                   *bool           `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                     *bool           `json:"is_enabled"`
    // The timestamp of when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp         *string         `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount       *int64          `json:"organizational_unit_count"`
}

// CreateWalletResponse represents a custom type struct for Success
type CreateWalletResponse struct {
    // Embedded responses related to the resource.
    Embedded           interface{}            `json:"_embedded"`
    // URLs to pages related to the resource.
    Links              *WalletLinks           `json:"_links"`
    // AWS Account ID associated with the wallet.
    AccountNativeId    *string                `json:"account_native_id"`
    // Version of the template available
    AvailableVersion   *int64                 `json:"available_version"`
    // The AWS region associated with the wallet.
    AwsRegion          *string                `json:"aws_region"`
    // Clumio AWS Account ID.
    ClumioAwsAccountId *string                `json:"clumio_aws_account_id"`
    // DeploymentURL is an (external) link to an AWS console page for quick-creation
    // of the stack.
    DeploymentUrl      *string                `json:"deployment_url"`
    // ErrorCode is a short string describing the error, if any.
    ErrorCode          *string                `json:"error_code"`
    // ErrorMessage is a longer description explaining the error, if any, and how to
    // fix it.
    ErrorMessage       *string                `json:"error_message"`
    // The Clumio-assigned ID of the wallet.
    Id                 *string                `json:"id"`
    // The regions where the wallet is installed.
    InstalledRegions   []*string              `json:"installed_regions"`
    // TODO: Add struct field description
    KeyErrors          map[string]*ErrorModel `json:"key_errors"`
    // RoleArn is the AWS Resource Name of the IAM Role created by the stack.
    RoleArn            *string                `json:"role_arn"`
    // The version of the stack used or being used.
    StackVersion       *int64                 `json:"stack_version"`
    // State describes the state of the wallet. Valid states are:
    // Waiting: The wallet has been created, but a stack hasn't been created. The
    // wallet can't be used in this state.
    // Enabled: The wallet has been created and a stack has been created for the
    // wallet. This is the normal expected state of a wallet in use.
    // Error:   The wallet is inaccessible. See ErrorCode and ErrorMessage fields for
    // additional details.
    State              *string                `json:"state"`
    // The supported regions for the wallet.
    SupportedRegions   []*string              `json:"supported_regions"`
    // TemplateURL is the URL to the CloudFormation template to be used to create the
    // CloudFormation stack.
    TemplateUrl        *string                `json:"template_url"`
    // Token is used to identify and authenticate the CloudFormation stack creation.
    Token              *string                `json:"token"`
}

// DeleteBucketFromProtectionGroupResponse represents a custom type struct for Success
type DeleteBucketFromProtectionGroupResponse struct {
    // The AWS-assigned ID of the account associated with the DynamoDB table.
    AccountNativeId               *string           `json:"account_native_id"`
    // Whether this bucket was added to this protection group by the bucket rule
    AddedByBucketRule             *bool             `json:"added_by_bucket_rule"`
    // Whether this bucket was added to this protection group by the user
    AddedByUser                   *bool             `json:"added_by_user"`
    // The AWS region associated with the DynamoDB table.
    AwsRegion                     *string           `json:"aws_region"`
    // The backup target AWS region associated with the protection group S3 asset.
    BackupTargetAwsRegion         *string           `json:"backup_target_aws_region"`
    // BackupTierStat
    BackupTierStats               []*BackupTierStat `json:"backup_tier_stats"`
    // The Clumio-assigned ID of the bucket
    BucketId                      *string           `json:"bucket_id"`
    // The name of the bucket
    BucketName                    *string           `json:"bucket_name"`
    // Creation time of the protection group in RFC-3339 format.
    CreatedTimestamp              *string           `json:"created_timestamp"`
    // The Clumio-assigned ID of the AWS environment associated with the protection group.
    EnvironmentId                 *string           `json:"environment_id"`
    // The Clumio-assigned ID of the protection group
    GroupId                       *string           `json:"group_id"`
    // The name of the protection group
    GroupName                     *string           `json:"group_name"`
    // The Clumio-assigned ID that represents the bucket within the protection group.
    Id                            *string           `json:"id"`
    // Determines whether the protection group bucket has been deleted
    IsDeleted                     *bool             `json:"is_deleted"`
    // Time of the last backup in RFC-3339 format.
    LastBackupTimestamp           *string           `json:"last_backup_timestamp"`
    // Time of the last successful continuous backup in RFC-3339 format.
    LastContinuousBackupTimestamp *string           `json:"last_continuous_backup_timestamp"`
    // The Clumio-assigned ID of the organizational unit associated with the protection group.
    OrganizationalUnitId          *string           `json:"organizational_unit_id"`
    // Cumulative count of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpObjectCount      *int64            `json:"total_backed_up_object_count"`
    // Cumulative size of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpSizeBytes        *int64            `json:"total_backed_up_size_bytes"`
    // The unsupported reason for the S3 bucket.
    UnsupportedReason             *string           `json:"unsupported_reason"`
}

// DeleteOrganizationalUnitResponse represents a custom type struct.
// Accepted
type DeleteOrganizationalUnitResponse struct {
    // Embedded responses related to the resource.
    Embedded *EntityGroupEmbedded              `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *OrganizationalUnitLinksForDelete `json:"_links"`
    // The Clumio-assigned ID of the task associated with this organizational unit.
    // The progress of the task can be monitored using the
    // [GET /tasks/{task_id}](#operation/read-task) endpoint.
    TaskId   *string                           `json:"task_id"`
}

// DeletePolicyResponse represents a custom type struct for Success
type DeletePolicyResponse struct {
    // URLs to pages related to the resource.
    Links  *DeletePolicyResponseLinks `json:"_links"`
    // The Clumio-assigned ID of the task generated by this request.
    TaskId *string                    `json:"task_id"`
}

// DeleteRuleResponse represents a custom type struct for Success
type DeleteRuleResponse struct {
    // URLs to pages related to the resource.
    Links     *DeleteRuleResponseLinks `json:"_links"`
    // The Clumio-assigned ID of the preview generated by this request. Only valid if
    // `execution_type` is set to `dryrun`.
    PreviewId *string                  `json:"preview_id"`
    // The Clumio-assigned ID of the task generated by this request.
    TaskId    *string                  `json:"task_id"`
}

// DeleteS3InstantAccessEndpointRoleResponse represents a custom type struct for Success
type DeleteS3InstantAccessEndpointRoleResponse struct {
    // Embedded responses related to the resource.
    Embedded *S3InstantAccessEndpointEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *S3InstantAccessEndpointLinks    `json:"_links"`
}

// DeleteUserResponseV1 represents a custom type struct for Success
type DeleteUserResponseV1 struct {
    // HateoasCommonLinks are the common fields for HATEOAS response.
    Links *HateoasCommonLinks `json:"_links"`
}

// DownloadSharedFileResponse represents a custom type struct for Success
type DownloadSharedFileResponse struct {
    // URLs to pages related to the resource.
    Links       *DownloadSharedFileLinks `json:"_links"`
    // A download link that lets you directly download the file. The link expires
    // 24 hours after file restore.
    DownloadUrl *string                  `json:"download_url"`
}

// EditProfileResponse represents a custom type struct for Success
type EditProfileResponse struct {
    // Embedded responses related to the resource.
    Embedded                   *UserEmbedded                 `json:"_embedded"`
    // ETag value
    Etag                       *string                       `json:"_etag"`
    // URLs to pages related to the resource.
    Links                      *UserLinks                    `json:"_links"`
    // The organizational units assigned to the user, with the specified role.
    AccessControlConfiguration []*RoleForOrganizationalUnits `json:"access_control_configuration"`
    // The email address of the Clumio user.
    Email                      *string                       `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                   *string                       `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                         *string                       `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                    *string                       `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                *bool                         `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                  *bool                         `json:"is_enabled"`
    // The timestamp of when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp      *string                       `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount    *int64                        `json:"organizational_unit_count"`
}

// EditProfileResponseV1 represents a custom type struct for Success
type EditProfileResponseV1 struct {
    // Embedded responses related to the resource.
    Embedded                      *UserEmbeddedV1 `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                         *UserLinks      `json:"_links"`
    // The list of organizational unit IDs assigned to the user.
    // This attribute will be available when reading a single user and not when listing users.
    AssignedOrganizationalUnitIds []*string       `json:"assigned_organizational_unit_ids"`
    // Assigned Role for the user.
    AssignedRole                  *string         `json:"assigned_role"`
    // The email address of the Clumio user.
    Email                         *string         `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                      *string         `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                            *string         `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                       *string         `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                   *bool           `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                     *bool           `json:"is_enabled"`
    // The timestamp of when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp         *string         `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount       *int64          `json:"organizational_unit_count"`
}

// Error represents a custom type struct.
// Error
type Error struct {
    // TODO: Add struct field description
    Errors []*SingleErrorResponse `json:"errors"`
}

// EstimateCostDetailsS3InstantAccessEndpointResponse represents a custom type struct for Success
type EstimateCostDetailsS3InstantAccessEndpointResponse struct {
    // The ETag value.
    Etag             *string                                                  `json:"_etag"`
    // URLs to pages related to the resource.
    Links            *EstimateCostDetailsS3InstantAccessEndpointResponseLinks `json:"_links"`
    // The estimated cost for instant access endpoint.
    EstimatedCost    *float64                                                 `json:"estimated_cost"`
    // The count of objects to be restored.
    TotalObjectCount *int64                                                   `json:"total_object_count"`
    // The total size in bytes of objects to be restored.
    TotalObjectSize  *int64                                                   `json:"total_object_size"`
}

// EstimateCostS3InstantAccessEndpointAsyncResponse represents a custom type struct.
// Success (Async)
type EstimateCostS3InstantAccessEndpointAsyncResponse struct {
    // URLs to pages related to the resource.
    Links      *EstimateCostS3InstantAccessEndpointAsyncResponseLinks `json:"_links"`
    // The identifier for the requested estimate which is used to fetch results.
    EstimateId *string                                                `json:"estimate_id"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    // Note that this field is given only for async request.
    TaskId     *string                                                `json:"task_id"`
}

// EstimateCostS3InstantAccessEndpointSyncResponse represents a custom type struct.
// Success (Sync)
type EstimateCostS3InstantAccessEndpointSyncResponse struct {
    // URLs to pages related to the resource.
    Links            *EstimateCostS3InstantAccessEndpointSyncResponseLinks `json:"_links"`
    // The estimated cost for instant access endpoint.
    EstimatedCost    *float64                                              `json:"estimated_cost"`
    // The count of objects to be restored.
    TotalObjectCount *int64                                                `json:"total_object_count"`
    // The total size in bytes of objects to be restored.
    TotalObjectSize  *int64                                                `json:"total_object_size"`
}

// FileListResponse represents a custom type struct for Success
type FileListResponse struct {
    // Embedded responses related to the resource.
    Embedded     *FileVersionsListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *FileVersionsListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64                    `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64                    `json:"limit"`
    // The page token used to get this response.
    Start        *string                   `json:"start"`
}

// FileSearchResponse represents a custom type struct for Success
type FileSearchResponse struct {
    // Embedded responses related to the resource.
    Embedded     *FileSearchListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *FileSearchListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64                  `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64                  `json:"limit"`
    // The page token used to get this response.
    Start        *string                 `json:"start"`
}

// GenerateRestoredFilePasscodeResponse represents a custom type struct for Success
type GenerateRestoredFilePasscodeResponse struct {
    // URLs to pages related to the resource.
    Links    *GenerateRestoredFilePasscodeLinks `json:"_links"`
    // The new passcode that has been generated for the restored file. Send the
    // passcode to the email recipient, who must use it to access the restored file.
    Passcode *string                            `json:"passcode"`
}

// ListAWSConnectionsResponse represents a custom type struct for Success
type ListAWSConnectionsResponse struct {
    // Embedded responses related to the resource.
    Embedded      *AWSConnectionListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links         *AWSConnectionListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount  *int64                     `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied *string                    `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit         *int64                     `json:"limit"`
    // The page token used to get this response.
    Start         *string                    `json:"start"`
}

// ListAWSEnvironmentsResponse represents a custom type struct for Success
type ListAWSEnvironmentsResponse struct {
    // Embedded responses related to the resource.
    Embedded     *AWSEnvironmentListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *AWSEnvironmentListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64                      `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64                      `json:"limit"`
    // The page token used to get this response.
    Start        *string                     `json:"start"`
}

// ListAWSRegionsResponse represents a custom type struct for Success
type ListAWSRegionsResponse struct {
    // Embedded responses related to the resource.
    Embedded     *ConnectionRegionListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *ConnectionRegionListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64                        `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64                        `json:"limit"`
    // The page token used to get this response.
    Start        *string                       `json:"start"`
}

// ListAlertsResponse represents a custom type struct for Success
type ListAlertsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *AlertListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *AlertListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64             `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string            `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64             `json:"limit"`
    // The sort order used in the request.
    SortApplied     *string            `json:"sort_applied"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string            `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64             `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64             `json:"total_pages_count"`
}

// ListAuditTrailsResponse represents a custom type struct for Success
type ListAuditTrailsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *AuditTrailListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *AuditTrailListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                  `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                 `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                  `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                 `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                  `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                  `json:"total_pages_count"`
}

// ListAutoUserProvisioningRulesResponse represents a custom type struct for Success
type ListAutoUserProvisioningRulesResponse struct {
    // Embedded responses related to the resource.
    Embedded     *AutoUserProvisioningRuleListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *AutoUserProvisioningRuleListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64                                `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64                                `json:"limit"`
    // The page token used to get this response.
    Start        *string                               `json:"start"`
}

// ListAwsTagsResponse represents a custom type struct for Success
type ListAwsTagsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *AwsTagListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *AwsTagListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64              `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string             `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64              `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string             `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64              `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64              `json:"total_pages_count"`
}

// ListBucketsResponse represents a custom type struct for Success
type ListBucketsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *BucketListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *BucketListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64              `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64              `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string             `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64              `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64              `json:"total_pages_count"`
}

// ListComplianceConfigurationsResponse represents a custom type struct for Success
type ListComplianceConfigurationsResponse struct {
    // An array of embedded resources related to this resource.
    Embedded      *ComplianceConfigurationListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links         *ComplianceConfigurationListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount  *int64                               `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied *string                              `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit         *int64                               `json:"limit"`
    // The page token used to get this response.
    Start         *string                              `json:"start"`
}

// ListComplianceRunsResponse represents a custom type struct for Success
type ListComplianceRunsResponse struct {
    // Embedded responses related to the resource.
    Embedded      *ComplianceRunListHateoasEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links         *ComplianceRunListHateoasLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount  *int64                            `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied *string                           `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit         *int64                            `json:"limit"`
    // The page token used to get this response.
    Start         *string                           `json:"start"`
}

// ListConnectionGroupsResponse represents a custom type struct for Success
type ListConnectionGroupsResponse struct {
    // Embedded responses related to the resource.
    Embedded      *AWSConnectionGroupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links         *ConnectionGroupListLinks       `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount  *int64                          `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied *string                         `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit         *int64                          `json:"limit"`
    // The page token used to get this response.
    Start         *string                         `json:"start"`
}

// ListConsolidatedAlertsResponse represents a custom type struct for Success
type ListConsolidatedAlertsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *ConsolidatedAlertListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *ConsolidatedAlertListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                         `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                        `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                         `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                        `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                         `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                         `json:"total_pages_count"`
}

// ListDynamoDBTableBackupsResponse represents a custom type struct for Success
type ListDynamoDBTableBackupsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *DynamoDBTableBackupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *DynamoDBTableBackupListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                           `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                          `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                           `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                          `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                           `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                           `json:"total_pages_count"`
}

// ListDynamoDBTableResponse represents a custom type struct for Success
type ListDynamoDBTableResponse struct {
    // Embedded responses related to the resource.
    Embedded        *DynamoDBTableListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *DynamoDBTableListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                     `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                     `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                    `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                     `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                     `json:"total_pages_count"`
}

// ListEBSBackupsResponse represents a custom type struct for Success
type ListEBSBackupsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *EBSBackupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *EBSBackupListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                 `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                 `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                 `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                 `json:"total_pages_count"`
}

// ListEBSBackupsResponseV1 represents a custom type struct for Success
type ListEBSBackupsResponseV1 struct {
    // Embedded responses related to the resource.
    Embedded        *EBSBackupListEmbeddedV1 `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *EBSBackupListLinksV1    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                   `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                  `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                   `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                  `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                   `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                   `json:"total_pages_count"`
}

// ListEC2BackupsResponse represents a custom type struct for Success
type ListEC2BackupsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *EC2BackupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *EC2BackupListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                 `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                 `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                 `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                 `json:"total_pages_count"`
}

// ListEC2MSSQLDatabaseBackupsResponse represents a custom type struct for Success
type ListEC2MSSQLDatabaseBackupsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *EC2MSSQLDatabaseBackupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *EC2MSSQLDatabaseBackupListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                              `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                             `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                              `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                             `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                              `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                              `json:"total_pages_count"`
}

// ListEC2MSSQLDatabasesResponse represents a custom type struct for Success
type ListEC2MSSQLDatabasesResponse struct {
    // Embedded responses related to the resource.
    Embedded        *EC2MSSQLDatabaseListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *EC2MSSQLDatabaseListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                        `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                       `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                        `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                       `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                        `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                        `json:"total_pages_count"`
}

// ListEC2MSSQLFCIsResponse represents a custom type struct for Success
type ListEC2MSSQLFCIsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *EC2MSSQLFCIListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *EC2MSSQLFCIListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                   `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                  `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                   `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                  `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                   `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                   `json:"total_pages_count"`
}

// ListEC2MSSQLInstancesResponse represents a custom type struct for Success
type ListEC2MSSQLInstancesResponse struct {
    // Embedded responses related to the resource.
    Embedded        *EC2MSSQLInstanceListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *EC2MSSQLInstanceListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                        `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                       `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                        `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                       `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                        `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                        `json:"total_pages_count"`
}

// ListEC2MSSQLInvHostsResponse represents a custom type struct for Success
type ListEC2MSSQLInvHostsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *EC2MSSQLInvHostListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *EC2MSSQLInvHostListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                       `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                      `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                       `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                      `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                       `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                       `json:"total_pages_count"`
}

// ListEC2MssqlAGsResponse represents a custom type struct for Success
type ListEC2MssqlAGsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *EC2MSSQLAGListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *EC2MSSQLAGListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                  `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                 `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                  `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                 `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                  `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                  `json:"total_pages_count"`
}

// ListEC2MssqlDatabasePitrIntervalsResponse represents a custom type struct.
// ListEC2MssqlDatabasePitrIntervalsResponse represents the success response
type ListEC2MssqlDatabasePitrIntervalsResponse struct {
    // Embedded responses related to the resource.
    Embedded      *EC2MssqlDatabasePitrIntervalListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links         *EC2MssqlDatabasePitrIntervalListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount  *int64                                    `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied *string                                   `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit         *int64                                    `json:"limit"`
    // The page token used to get this response.
    Start         *string                                   `json:"start"`
}

// ListEbsVolumesResponse represents a custom type struct for Success
type ListEbsVolumesResponse struct {
    // Embedded responses related to the resource.
    Embedded        *EbsVolumeListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *EbsVolumeListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                 `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                 `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                 `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                 `json:"total_pages_count"`
}

// ListEc2InstancesResponse represents a custom type struct for Success
type ListEc2InstancesResponse struct {
    // Embedded responses related to the resource.
    Embedded        *Ec2InstanceListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *EC2InstanceListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                   `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                   `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                  `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                   `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                   `json:"total_pages_count"`
}

// ListFileSystemsResponse represents a custom type struct for Success
type ListFileSystemsResponse struct {
    // _embedded contains the list of items of a file_system list query
    Embedded        *FileSystemListEmbedded `json:"_embedded"`
    // _links provides URLs to related navigable pages of a file_system list query
    Links           *FileSystemListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                  `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                  `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                 `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                  `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                  `json:"total_pages_count"`
}

// ListManagementGroupsResponse represents a custom type struct for Success
type ListManagementGroupsResponse struct {
    // Embedded responses related to the resource.
    Embedded     *ManagementGroupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *ManagementGroupListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64                       `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64                       `json:"limit"`
    // The total count of groups upto 10. Any number of groups beyond 10 will
    // still be returned as 10.
    MinCount     *int64                       `json:"min_count"`
    // The page token used to get this response.
    Start        *string                      `json:"start"`
}

// ListOrganizationalUnitsResponse represents a custom type struct for Success
type ListOrganizationalUnitsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *OrganizationalUnitListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *OrganizationalUnitListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                          `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                         `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                          `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                         `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                          `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                          `json:"total_pages_count"`
}

// ListOrganizationalUnitsResponseV1 represents a custom type struct for Success
type ListOrganizationalUnitsResponseV1 struct {
    // Embedded responses related to the resource.
    Embedded        *OrganizationalUnitListEmbeddedV1 `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *OrganizationalUnitListLinks      `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                            `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                           `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                            `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                           `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                            `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                            `json:"total_pages_count"`
}

// ListPoliciesResponse represents a custom type struct for Success
type ListPoliciesResponse struct {
    // An array of embedded resources related to this resource.
    Embedded      *PolicyListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links         *PolicyListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount  *int64              `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied *string             `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit         *int64              `json:"limit"`
    // The page token used to get this response.
    Start         *string             `json:"start"`
}

// ListProtectionGroupBackupsResponse represents a custom type struct for Success
type ListProtectionGroupBackupsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *ProtectionGroupBackupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *ProtectionGroupBackupListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                             `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                            `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                             `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                            `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                             `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                             `json:"total_pages_count"`
}

// ListProtectionGroupS3AssetBackupsResponse represents a custom type struct for Success
type ListProtectionGroupS3AssetBackupsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *ProtectionGroupS3AssetBackupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *ProtectionGroupS3AssetBackupListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                                    `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                                   `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                                    `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                                   `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                                    `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                                    `json:"total_pages_count"`
}

// ListProtectionGroupS3AssetPitrIntervalsResponse represents a custom type struct.
// ListProtectionGroupS3AssetPitrIntervalsResponse represents the success response
type ListProtectionGroupS3AssetPitrIntervalsResponse struct {
    // Embedded responses related to the resource.
    Embedded      *ProtectionGroupS3AssetPitrIntervalListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links         *ProtectionGroupS3AssetPitrIntervalListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount  *int64                                          `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied *string                                         `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit         *int64                                          `json:"limit"`
    // The page token used to get this response.
    Start         *string                                         `json:"start"`
}

// ListProtectionGroupS3AssetsResponse represents a custom type struct for Success
type ListProtectionGroupS3AssetsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *ProtectionGroupBucketListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *ProtectionGroupBucketListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                             `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                             `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                            `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                             `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                             `json:"total_pages_count"`
}

// ListProtectionGroupsResponse represents a custom type struct for Success
type ListProtectionGroupsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *ProtectionGroupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *ProtectionGroupListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                       `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                       `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                      `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                       `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                       `json:"total_pages_count"`
}

// ListRDSBackupDatabasesResponse represents a custom type struct for Success
type ListRDSBackupDatabasesResponse struct {
    // Embedded responses related to the resource.
    Embedded     *RDSBackupDatabaseListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *RDSBackupDatabaseListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64                         `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64                         `json:"limit"`
    // The page token used to get this response.
    Start        *string                        `json:"start"`
}

// ListRDSDatabaseTablesResponse represents a custom type struct for Success
type ListRDSDatabaseTablesResponse struct {
    // Embedded responses related to the resource.
    Embedded     *RDSDatabaseTableListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *RDSDatabaseTableListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64                        `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64                        `json:"limit"`
    // The page token used to get this response.
    Start        *string                       `json:"start"`
}

// ListRdsDatabaseBackupsResponse represents a custom type struct for Success
type ListRdsDatabaseBackupsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *RdsDatabaseBackupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *RdsDatabaseBackupListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                         `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                        `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                         `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                        `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                         `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                         `json:"total_pages_count"`
}

// ListRdsOptionGroupsResponse represents a custom type struct for Success
type ListRdsOptionGroupsResponse struct {
    // Embedded responses related to the resource.
    Embedded      *OptionGroupsListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links         *OptionGroupsListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount  *int64                    `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied *string                   `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit         *int64                    `json:"limit"`
    // The page token used to get this response.
    Start         *string                   `json:"start"`
}

// ListRdsResourcesResponse represents a custom type struct for Success
type ListRdsResourcesResponse struct {
    // Embedded responses related to the resource.
    Embedded        *RdsResourceListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *RdsResourceListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                   `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                  `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                   `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                  `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                   `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                   `json:"total_pages_count"`
}

// ListReportDownloadsResponse represents a custom type struct for Success
type ListReportDownloadsResponse struct {
    // _embedded contains the list of items of a list report CSV download query
    Embedded        *ReportDownloadListEmbedded `json:"_embedded"`
    // _links provides URLs to related navigable pages of a list report CSV download query
    Links           *ReportDownloadListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                      `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                     `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                      `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                     `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                      `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                      `json:"total_pages_count"`
}

// ListRestoredRecordsResponse represents a custom type struct for Success
type ListRestoredRecordsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *RestoredRecordListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *RestoredRecordListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                      `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                      `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                     `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                      `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                      `json:"total_pages_count"`
}

// ListRolesResponse represents a custom type struct for Success
type ListRolesResponse struct {
    // Embedded responses related to the resource.
    Embedded        *RoleListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *RoleListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64            `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string           `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64            `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string           `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64            `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64            `json:"total_pages_count"`
}

// ListRulesResponse represents a custom type struct for Success
type ListRulesResponse struct {
    // An array of embedded resources related to this resource.
    Embedded     *RuleListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *RuleListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64            `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64            `json:"limit"`
    // The page token used to get this response.
    Start        *string           `json:"start"`
}

// ListS3InstantAccessEndpointsResponse represents a custom type struct for Success
type ListS3InstantAccessEndpointsResponse struct {
    // Embedded responses related to the resource.
    Embedded     *S3InstantAccessEndpointListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *S3InstantAccessEndpointListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64                               `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64                               `json:"limit"`
    // The page token used to get this response.
    Start        *string                              `json:"start"`
}

// ListTasksResponse represents a custom type struct for Success
type ListTasksResponse struct {
    // Embedded responses related to the resource.
    Embedded        *TaskListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *TaskListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64            `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string           `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64            `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string           `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64            `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64            `json:"total_pages_count"`
}

// ListUsersResponse represents a custom type struct for Success
type ListUsersResponse struct {
    // Embedded responses related to the resource.
    Embedded        *UserListEmbedded     `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *UserListHateoasLinks `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string               `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string               `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                `json:"total_pages_count"`
}

// ListUsersResponseV1 represents a custom type struct for Success
type ListUsersResponseV1 struct {
    // Embedded responses related to the resource.
    Embedded        *UserListEmbeddedV1   `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *UserListHateoasLinks `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string               `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string               `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                `json:"total_pages_count"`
}

// ListWalletsResponse represents a custom type struct for Success
type ListWalletsResponse struct {
    // Embedded responses related to the resource.
    Embedded     *WalletListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *WalletListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64              `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64              `json:"limit"`
    // The page token used to get this response.
    Start        *string             `json:"start"`
}

// OnDemandDynamoDBBackupResponse represents a custom type struct for Success
type OnDemandDynamoDBBackupResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *OnDemandDynamoDBBackupLinks  `json:"_links"`
    // The Clumio-assigned ID of the task created for DynamoDB backup.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                       `json:"task_id"`
}

// OnDemandEBSBackupResponse represents a custom type struct for Success
type OnDemandEBSBackupResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *OnDemandEBSBackupLinks       `json:"_links"`
    // The Clumio-assigned ID of the task created for EBS backup.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                       `json:"task_id"`
}

// OnDemandEBSBackupResponseV1 represents a custom type struct for Success
type OnDemandEBSBackupResponseV1 struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *OnDemandEBSBackupLinks       `json:"_links"`
    // The Clumio-assigned ID of the task created for EBS backup.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                       `json:"task_id"`
}

// OnDemandEC2BackupResponse represents a custom type struct for Success
type OnDemandEC2BackupResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *OnDemandEC2BackupLinks       `json:"_links"`
    // The Clumio-assigned ID of the task created for EC2 backup.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                       `json:"task_id"`
}

// OnDemandEC2MSSQLDatabaseBackupResponse represents a custom type struct for Success
type OnDemandEC2MSSQLDatabaseBackupResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded                      `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *CreateOnDemandEC2MSSQLDatabaseBackupResponseLinks `json:"_links"`
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink                                   `json:"_self"`
    // Task Id
    TaskId   *string                                            `json:"task_id"`
}

// PatchGeneralSettingsResponseV2 represents a custom type struct for Success
type PatchGeneralSettingsResponseV2 struct {
    // URLs to pages related to the resource.
    Links                        *GeneralSettingsLinks `json:"_links"`
    // The length of time before a user is logged out of the Clumio system due to inactivity. Measured in seconds.
    // The valid range is between 600 seconds (10 minutes) and 3600 seconds (60 minutes).
    // If not configured, the value defaults to 900 seconds (15 minutes).
    AutoLogoutDuration           *int64                `json:"auto_logout_duration"`
    // The designated range of IP addresses that are allowed to access the Clumio REST API.
    // API requests that originate from outside this list will be blocked.
    // The IP address of the server from which this request is being made must be in this list; otherwise, the request will fail.
    // Set the parameter to individual IP addresses and/or a range of IP addresses in CIDR notation.
    // For example, ["193.168.1.0/24", "193.172.1.1"].
    // If not configured, the value defaults to ["0.0.0.0/0"] meaning all addresses will be allowed.
    IpAllowlist                  []*string             `json:"ip_allowlist"`
    // The grouping criteria for each datasource type.
    // These can only be edited for datasource types which do not have any
    // organizational units configured.
    OrganizationalUnitDataGroups *OUGroupingCriteria   `json:"organizational_unit_data_groups"`
    // The length of time a user password is valid before it must be changed. Measured in seconds.
    // The valid range is between 2592000 seconds (30 days) and 15552000 seconds (180 days).
    // If not configured, the value defaults to 7776000 seconds (90 days).
    PasswordExpirationDuration   *int64                `json:"password_expiration_duration"`
}

// PatchOrganizationalUnitNoTaskResponse represents a custom type struct for Success
type PatchOrganizationalUnitNoTaskResponse struct {
    // URLs to pages related to the resource.
    Links                     *OULinks        `json:"_links"`
    // Number of immediate children of the organizational unit.
    ChildrenCount             *int64          `json:"children_count"`
    // Datasource types configured in this organizational unit. Possible values include `aws`, `microsoft365`, `vmware`, or `mssql`.
    ConfiguredDatasourceTypes []*string       `json:"configured_datasource_types"`
    // List of all recursive descendant organizational units of this OU.
    DescendantIds             []*string       `json:"descendant_ids"`
    // A description of the organizational unit.
    Description               *string         `json:"description"`
    // The Clumio assigned ID of the organizational unit.
    Id                        *string         `json:"id"`
    // Unique name assigned to the organizational unit.
    Name                      *string         `json:"name"`
    // The Clumio assigned ID of the parent organizational unit.
    // The parent organizational unit contains the entities in this organizational unit and can update this organizational unit.
    // If this organizational unit is the global organizational unit, then this field has a value of `null`.
    ParentId                  *string         `json:"parent_id"`
    // Number of users to whom this organizational unit or any of its descendants have been assigned.
    UserCount                 *int64          `json:"user_count"`
    // A user along with a role.
    Users                     []*UserWithRole `json:"users"`
}

// PatchOrganizationalUnitNoTaskResponseV1 represents a custom type struct for Success
type PatchOrganizationalUnitNoTaskResponseV1 struct {
    // URLs to pages related to the resource.
    Links                     *OULinks  `json:"_links"`
    // Number of immediate children of the organizational unit.
    ChildrenCount             *int64    `json:"children_count"`
    // Datasource types configured in this organizational unit. Possible values include `aws`, `microsoft365`, `vmware`, or `mssql`.
    ConfiguredDatasourceTypes []*string `json:"configured_datasource_types"`
    // List of all recursive descendant organizational units of this OU.
    DescendantIds             []*string `json:"descendant_ids"`
    // A description of the organizational unit.
    Description               *string   `json:"description"`
    // The Clumio assigned ID of the organizational unit.
    Id                        *string   `json:"id"`
    // Unique name assigned to the organizational unit.
    Name                      *string   `json:"name"`
    // The Clumio assigned ID of the parent organizational unit.
    // The parent organizational unit contains the entities in this organizational unit and can update this organizational unit.
    // If this organizational unit is the global organizational unit, then this field has a value of `null`.
    ParentId                  *string   `json:"parent_id"`
    // Number of users to whom this organizational unit or any of its descendants have been assigned.
    UserCount                 *int64    `json:"user_count"`
    // Users IDs to whom the organizational unit has been assigned.
    // This attribute will be available when reading a single OU and not when listing OUs.
    Users                     []*string `json:"users"`
}

// PatchOrganizationalUnitResponse represents a custom type struct.
// Accepted
type PatchOrganizationalUnitResponse struct {
    // Embedded responses related to the resource.
    Embedded                  *EntityGroupEmbedded     `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                     *OrganizationalUnitLinks `json:"_links"`
    // Number of immediate children of the organizational unit.
    ChildrenCount             *int64                   `json:"children_count"`
    // Datasource types configured in this organizational unit. Possible values include `aws`, `microsoft365`, `vmware`, or `mssql`.
    ConfiguredDatasourceTypes []*string                `json:"configured_datasource_types"`
    // List of all recursive descendant organizational units of this OU.
    DescendantIds             []*string                `json:"descendant_ids"`
    // A description of the organizational unit.
    Description               *string                  `json:"description"`
    // The Clumio assigned ID of the organizational unit.
    Id                        *string                  `json:"id"`
    // Unique name assigned to the organizational unit.
    Name                      *string                  `json:"name"`
    // The Clumio assigned ID of the parent organizational unit.
    // The parent organizational unit contains the entities in this organizational unit and can update this organizational unit.
    // If this organizational unit is the global organizational unit, then this field has a value of `null`.
    ParentId                  *string                  `json:"parent_id"`
    // The Clumio-assigned ID of the task associated with this organizational unit.
    // The progress of the task can be monitored using the
    // [GET /tasks/{task_id}](#operation/read-task) endpoint.
    TaskId                    *string                  `json:"task_id"`
    // Number of users to whom this organizational unit or any of its descendants have been assigned.
    UserCount                 *int64                   `json:"user_count"`
    // A user along with a role.
    Users                     []*UserWithRole          `json:"users"`
}

// PatchOrganizationalUnitResponseV1 represents a custom type struct.
// Accepted
type PatchOrganizationalUnitResponseV1 struct {
    // Embedded responses related to the resource.
    Embedded                  *EntityGroupEmbedded     `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                     *OrganizationalUnitLinks `json:"_links"`
    // Number of immediate children of the organizational unit.
    ChildrenCount             *int64                   `json:"children_count"`
    // Datasource types configured in this organizational unit. Possible values include `aws`, `microsoft365`, `vmware`, or `mssql`.
    ConfiguredDatasourceTypes []*string                `json:"configured_datasource_types"`
    // List of all recursive descendant organizational units of this OU.
    DescendantIds             []*string                `json:"descendant_ids"`
    // A description of the organizational unit.
    Description               *string                  `json:"description"`
    // The Clumio assigned ID of the organizational unit.
    Id                        *string                  `json:"id"`
    // Unique name assigned to the organizational unit.
    Name                      *string                  `json:"name"`
    // The Clumio assigned ID of the parent organizational unit.
    // The parent organizational unit contains the entities in this organizational unit and can update this organizational unit.
    // If this organizational unit is the global organizational unit, then this field has a value of `null`.
    ParentId                  *string                  `json:"parent_id"`
    // The Clumio-assigned ID of the task associated with this organizational unit.
    // The progress of the task can be monitored using the
    // [GET /tasks/{task_id}](#operation/read-task) endpoint.
    TaskId                    *string                  `json:"task_id"`
    // Number of users to whom this organizational unit or any of its descendants have been assigned.
    UserCount                 *int64                   `json:"user_count"`
    // Users IDs to whom the organizational unit has been assigned.
    // This attribute will be available when reading a single OU and not when listing OUs.
    Users                     []*string                `json:"users"`
}

// PreviewDetailsProtectionGroupResponse represents a custom type struct for Success
type PreviewDetailsProtectionGroupResponse struct {
    // The ETag value.
    Etag    *string                             `json:"_etag"`
    // URLs to pages related to the resource.
    Links   *PreviewDetailsProtectionGroupLinks `json:"_links"`
    // Object defines one object to restore
    Objects []*Object                           `json:"objects"`
}

// PreviewDetailsS3BucketResponse represents a custom type struct for Success
type PreviewDetailsS3BucketResponse struct {
    // URLs to pages related to the resource.
    Links   *PreviewDetailsS3BucketLinks `json:"_links"`
    // ObjectV2 defines one object to restore
    Objects []*ObjectV2                  `json:"objects"`
}

// PreviewProtectionGroupAsyncResponse represents a custom type struct.
// Success (Async)
type PreviewProtectionGroupAsyncResponse struct {
    // URLs to pages related to the resource.
    Links     *PreviewProtectionGroupAsyncLinks `json:"_links"`
    // The identifier for the requested preview which is used to fetch results of the preview. 
    // Note that this field is given only for async request.
    PreviewId *string                           `json:"preview_id"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    // Note that this field is given only for async request.
    TaskId    *string                           `json:"task_id"`
}

// PreviewProtectionGroupS3AssetAsyncResponse represents a custom type struct.
// Success (Async)
type PreviewProtectionGroupS3AssetAsyncResponse struct {
    // URLs to pages related to the resource.
    Links     *PreviewProtectionGroupS3AssetAsyncLinks `json:"_links"`
    // The identifier for the requested preview which is used to fetch results of the preview.
    PreviewId *string                                  `json:"preview_id"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    // Note that this field is given only for async request.
    TaskId    *string                                  `json:"task_id"`
}

// PreviewProtectionGroupS3AssetDetailsResponse represents a custom type struct for Success
type PreviewProtectionGroupS3AssetDetailsResponse struct {
    // The ETag value.
    Etag    *string                                    `json:"_etag"`
    // URLs to pages related to the resource.
    Links   *PreviewProtectionGroupS3AssetDetailsLinks `json:"_links"`
    // Object defines one object to restore
    Objects []*Object                                  `json:"objects"`
}

// PreviewProtectionGroupS3AssetSyncResponse represents a custom type struct.
// Success (Sync)
type PreviewProtectionGroupS3AssetSyncResponse struct {
    // URLs to pages related to the resource.
    Links   *PreviewProtectionGroupS3AssetSyncLinks `json:"_links"`
    // Object defines one object to restore
    Objects []*Object                               `json:"objects"`
}

// PreviewProtectionGroupSyncResponse represents a custom type struct.
// Success (Sync)
type PreviewProtectionGroupSyncResponse struct {
    // URLs to pages related to the resource.
    Links   *PreviewProtectionGroupSyncLinks `json:"_links"`
    // Object defines one object to restore
    Objects []*Object                        `json:"objects"`
}

// PreviewS3BucketResponse represents a custom type struct for Success
type PreviewS3BucketResponse struct {
    // URLs to pages related to the resource.
    Links     *PreviewS3BucketLinks `json:"_links"`
    // The identifier for the requested preview which is used to fetch results of the preview.
    PreviewId *string               `json:"preview_id"`
    // The Clumio-assigned ID of the task created by this preview request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId    *string               `json:"task_id"`
}

// ReadAWSConnectionResponse represents a custom type struct for Success
type ReadAWSConnectionResponse struct {
    // The ETag value.
    Etag                       *string                  `json:"_etag"`
    // URLs to pages related to the resource.
    Links                      *AWSConnectionLinks      `json:"_links"`
    // The alias given to the account on AWS.
    AccountName                *string                  `json:"account_name"`
    // The AWS-assigned ID of the account associated with the connection.
    AccountNativeId            *string                  `json:"account_native_id"`
    // The AWS region associated with the connection. For example, `us-east-1`.
    AwsRegion                  *string                  `json:"aws_region"`
    // AWS account ID of the Clumio control plane.
    ClumioAwsAccountId         *string                  `json:"clumio_aws_account_id"`
    // AWS region of the Clumio control plane
    ClumioAwsRegion            *string                  `json:"clumio_aws_region"`
    // The consolidated configuration of the Clumio Cloud Protect and Clumio Cloud Discover products for this connection.
    // If this connection is deprecated to use unconsolidated configuration, then this field has a
    // value of `null`.
    Config                     *ConsolidatedConfig      `json:"config"`
    // Clumio assigned ID of the associated connection group.
    ConnectionGroupId          *string                  `json:"connection_group_id"`
    // Management status of connection.
    ConnectionManagementStatus *string                  `json:"connection_management_status"`
    // The status of the connection considering all the deployments made for it.
    ConnectionStatus           *string                  `json:"connection_status"`
    // The timestamp of when the connection was created.
    CreatedTimestamp           *string                  `json:"created_timestamp"`
    // AWS account ID of the data plane for the connection.
    DataPlaneAccountId         *string                  `json:"data_plane_account_id"`
    // The deployment method with which the currently active connection was established.
    DeploymentType             *string                  `json:"deployment_type"`
    // The user-provided description for this connection.
    Description                *string                  `json:"description"`
    // The configuration of the Clumio Discover product for this connection.
    // If this connection is not configured for Clumio Discover, then this field has a
    // value of `null`.
    Discover                   *DiscoverConfig          `json:"discover"`
    // Clumio assigned external ID of the connection or of the associated connection group.
    ExternalId                 *string                  `json:"external_id"`
    // The Clumio-assigned ID of the connection.
    Id                         *string                  `json:"id"`
    // Status denoting whether Ingestion has started for the connection.
    // Valid values are "initial", "in_progress", "failed", "completed".
    IngestionStatus            *string                  `json:"ingestion_status"`
    // K8S Namespace
    Namespace                  *string                  `json:"namespace"`
    // The Clumio-assigned ID of the organizational unit associated with the
    // AWS environment. If this parameter is not provided, then the value
    // defaults to the first organizational unit assigned to the requesting
    // user. For more information about organizational units, refer to the
    // Organizational-Units documentation.
    OrganizationalUnitId       *string                  `json:"organizational_unit_id"`
    // The configuration of the Clumio Cloud Protect product for this connection.
    // If this connection is not configured for Clumio Cloud Protect, then this field has a
    // value of `null`.
    Protect                    *ProtectConfig           `json:"protect"`
    // The asset types enabled for protect.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS", "IcebergOnGlue", "IcebergOnS3Tables"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    ProtectAssetTypesEnabled   []*string                `json:"protect_asset_types_enabled"`
    // TODO: Add struct field description
    Resources                  *ConnectionResourcesResp `json:"resources"`
    // The Amazon Resource Name of the stale CloudFormation stack when the connection was migrated to connection groups.
    // NOTE - This has to be removed from AWS as well to delete the connection completely.
    RetiredStackArn            *string                  `json:"retired_stack_arn"`
    // The services to be enabled for this configuration. Valid values are
    // ["discover"], ["discover", "protect"]. This is only set when the
    // registration is created, the enabled services are obtained directly from
    // the installed template after that. (Deprecated as all connections will have
    // both discover and protect enabled)
    ServicesEnabled            []*string                `json:"services_enabled"`
    // The Amazon Resource Name of the installed and active CloudFormation stack(if any) in AWS.
    StackArn                   *string                  `json:"stack_arn"`
    // The name given to the installed and active CloudFormation stack(if any) in AWS.
    StackName                  *string                  `json:"stack_name"`
    // Status denoting whether Target Setup has started for the connection.
    // Valid values are "initial", "in_progress", "failed", "completed".
    TargetSetupStatus          *string                  `json:"target_setup_status"`
    // TODO: Add struct field description
    TemplatePermissionSet      *string                  `json:"template_permission_set"`
    // The 36-character Clumio AWS integration ID token used to identify the
    // installation of the CloudFormation template on the account. This value
    // will be pasted into the ClumioToken field when creating the
    // CloudFormation stack.
    Token                      *string                  `json:"token"`
}

// ReadAWSEnvironmentResponse represents a custom type struct for Success
type ReadAWSEnvironmentResponse struct {
    // Embedded responses related to the resource.
    Embedded                   *AWSEnvironmentEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                      *AWSEnvironmentLinks    `json:"_links"`
    // The name given to the account.
    AccountName                *string                 `json:"account_name"`
    // The AWS-assigned ID of the account associated with the environment.
    AccountNativeId            *string                 `json:"account_native_id"`
    // The valid AWS availability zones for the environment. For example, `us_west-2a`.
    AwsAz                      []*string               `json:"aws_az"`
    // The AWS region associated with the environment. For example, `us-west-2`.
    AwsRegion                  *string                 `json:"aws_region"`
    // The consolidated configuration of the Clumio Cloud Protect and Clumio Cloud Discover products for this connection.
    // If this connection is deprecated to use unconsolidated configuration, then this field has a
    // value of `null`.
    Config                     *ConsolidatedConfig     `json:"config"`
    // Clumio assigned ID of the associated connection group (if any).
    ConnectionGroupId          *string                 `json:"connection_group_id"`
    // The Clumio-assigned ID of the connection associated with the environment.
    ConnectionId               *string                 `json:"connection_id"`
    // Management status of connection for the environment.
    ConnectionManagementStatus *string                 `json:"connection_management_status"`
    // The status of the connection to the environment.
    ConnectionStatus           *string                 `json:"connection_status"`
    // The user-provided account description.
    Description                *string                 `json:"description"`
    // The Clumio-assigned ID of the environment.
    Id                         *string                 `json:"id"`
    // The Clumio-assigned ID of the organizational unit associated with the environment.
    OrganizationalUnitId       *string                 `json:"organizational_unit_id"`
    // The AWS services enabled for this environment. Possible values include "ebs", "rds" and "dynamodb".
    ServicesEnabled            []*string               `json:"services_enabled"`
    // The Clumio CloudFormation template version used to deploy the CloudFormation stack.
    TemplateVersion            *int64                  `json:"template_version"`
}

// ReadAWSTemplatesV2Response represents a custom type struct for Success
type ReadAWSTemplatesV2Response struct {
    // URLs to pages related to the resource.
    Links  *TemplateLinks           `json:"_links"`
    // The configuration of the given template
    Config *TemplateConfigurationV2 `json:"config"`
}

// ReadAlertResponse represents a custom type struct for Success
type ReadAlertResponse struct {
    // Embedded responses related to the resource.
    Embedded            *AlertEmbedded          `json:"_embedded"`
    // The ETag value.
    Etag                *string                 `json:"_etag"`
    // URLs to pages related to the resource.
    Links               *AlertLinks             `json:"_links"`
    // The issue that generated the alert. Each cause belongs to an alert type.
    Cause               *string                 `json:"cause"`
    // The timestamp of when the alert was cleared, either automatically by Clumio or manually by a Clumio user.
    // Represented in RFC-3339 format. If this alert has not been cleared, then this field has a value of `null`.
    ClearedTimestamp    *string                 `json:"cleared_timestamp"`
    // The Clumio-assigned ID of the consolidated alert associated with this individual alert. Alerts are consolidated based on matching parent entity, alert type, and alert cause.
    ConsolidatedAlertId *string                 `json:"consolidated_alert_id"`
    // Additional information about the alert.
    Details             *IndividualAlertDetails `json:"details"`
    // The Clumio-assigned ID of the individual alert.
    Id                  *string                 `json:"id"`
    // A record of user-provided information about the alert.
    Notes               *string                 `json:"notes"`
    // The parent object of the primary entity associated with or affected by the alert. For example, "aws_environment" is the parent entity of primary entity "aws_ebs_volume".
    ParentEntity        *AlertParentEntity      `json:"parent_entity"`
    // The primary object associated with or affected by the alert. Examples of primary entities include "aws_connection", "aws_ebs_volume".
    PrimaryEntity       *AlertPrimaryEntity     `json:"primary_entity"`
    // The number of times the alert has recurred for this primary entity.
    RaisedCount         *uint64                 `json:"raised_count"`
    // The timestamp of when the alert was raised. Represented in RFC-3339 format.
    RaisedTimestamp     *string                 `json:"raised_timestamp"`
    // The alert severity level. Values include "error" and "warning".
    Severity            *string                 `json:"severity"`
    // The individual alert status. An individual alert that is in "active" status is one that is still open and has yet to be addressed.
    // An individual alert that is in "cleared" status is one that has been cleared, either automatically by Clumio or manually by a Clumio user.
    Status              *string                 `json:"status"`
    // TODO: Add struct field description
    Tags                []*RestEntity           `json:"tags"`
    // The general alert category. Some alert types may be associated with multiple causes.
    // Refer to the Alert Type table for a complete list of alert types.
    ClumioType          *string                 `json:"type"`
    // The timestamp of when the alert was last updated. Represented in RFC-3339 format.
    // The alert is updated whenever there is a new occurrence of the same alert within the same entity.
    UpdatedTimestamp    *string                 `json:"updated_timestamp"`
}

// ReadAutoUserProvisioningRuleResponse represents a custom type struct for Success
type ReadAutoUserProvisioningRuleResponse struct {
    // Embedded responses related to the resource.
    Embedded  *AutoUserProvisioningRuleEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links     *AutoUserProvisioningRuleLinks    `json:"_links"`
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
    Condition *string                           `json:"condition"`
    // Unique name assigned to the rule.
    Name      *string                           `json:"name"`
    // Specifies the role and the organizational units to be assigned to the user subject to the rule criteria.
    Provision *RuleProvision                    `json:"provision"`
    // The Clumio-assigned ID of the rule.
    RuleId    *string                           `json:"rule_id"`
}

// ReadAutoUserProvisioningSettingResponse represents a custom type struct for Success
type ReadAutoUserProvisioningSettingResponse struct {
    // URLs to pages related to the resource.
    Links     *AutoUserProvisioningSettingLinks `json:"_links"`
    // Whether auto user provisioning is enabled or not.
    IsEnabled *bool                             `json:"is_enabled"`
}

// ReadAwsTagResponse represents a custom type struct for Success
type ReadAwsTagResponse struct {
    // Embedded responses related to the resource.
    Embedded             *AwsTagEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                *AwsTagLinks    `json:"_links"`
    // The Clumio-assigned ID of the AWS tag.
    Id                   *string         `json:"id"`
    // The AWS-assigned tag key.
    Key                  *string         `json:"key"`
    // The Clumio-assigned ID of the AWS key.
    KeyId                *string         `json:"key_id"`
    // The Clumio-assigned ID of the organizational unit associated with the tag.
    OrganizationalUnitId *string         `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo       *ProtectionInfo `json:"protection_info"`
    // The protection status of the EBS volume. Possible values include "protected" and "unprotected".
    ProtectionStatus     *string         `json:"protection_status"`
    // The AWS-assigned tag value.
    Value                *string         `json:"value"`
}

// ReadBucketResponse represents a custom type struct for Success
type ReadBucketResponse struct {
    // Embedded responses related to the resource.
    Embedded                      interface{}                                   `json:"_embedded"`
    // The ETag value.
    Etag                          *string                                       `json:"_etag"`
    // URLs to pages related to the resource.
    Links                         *BucketLinks                                  `json:"_links"`
    // The AWS-assigned ID of the account associated with the S3 bucket.
    AccountNativeId               *string                                       `json:"account_native_id"`
    // The AWS region associated with the S3 bucket.
    AwsRegion                     *string                                       `json:"aws_region"`
    // The total size breakdown of S3 buckets in bytes per storage class. This parameter aggregates relevant fields from cloudwatch_metrics > size_bytes_per_storage_class
    BucketSizeBytesBreakdown      *S3BucketsInventorySummaryBucketSizeBreakdown `json:"bucket_size_bytes_breakdown"`
    // The Cloudwatch metrics of the bucket.
    CloudwatchMetrics             *S3CloudwatchMetrics                          `json:"cloudwatch_metrics"`
    // The timestamp of when the bucket was created. Represented in RFC-3339 format.
    CreationTimestamp             *string                                       `json:"creation_timestamp"`
    // The AWS encryption output of the bucket.
    EncryptionSetting             *S3EncryptionOutput                           `json:"encryption_setting"`
    // The Clumio-assigned ID of the AWS environment associated with the S3 bucket.
    EnvironmentId                 *string                                       `json:"environment_id"`
    // The EventBridge enablement state for the S3 bucket.
    EventBridgeEnabled            *bool                                         `json:"event_bridge_enabled"`
    // The Clumio-assigned ID of the bucket.
    Id                            *string                                       `json:"id"`
    // The Encryption enablement state for the S3 bucket.
    IsEncryptionEnabled           *bool                                         `json:"is_encryption_enabled"`
    // The Replication enablement state for the S3 bucket.
    IsReplicationEnabled          *bool                                         `json:"is_replication_enabled"`
    // The Versioning enablement state for the S3 bucket.
    IsVersioningEnabled           *bool                                         `json:"is_versioning_enabled"`
    // Time of the last S3 Backtrack sync in RFC-3339 format.
    LastBacktrackSyncTimestamp    *string                                       `json:"last_backtrack_sync_timestamp"`
    // Time of the last backup in RFC-3339 format.
    LastBackupTimestamp           *string                                       `json:"last_backup_timestamp"`
    // Time of the last continuous backup in RFC-3339 format.
    LastContinuousBackupTimestamp *string                                       `json:"last_continuous_backup_timestamp"`
    // The AWS-assigned name of the bucket.
    Name                          *string                                       `json:"name"`
    // Number of objects in bucket.
    ObjectCount                   *int64                                        `json:"object_count"`
    // The Clumio-assigned ID of the organizational unit associated with the S3 bucket.
    OrganizationalUnitId          *string                                       `json:"organizational_unit_id"`
    // Protection group count reflects how many protection groups are linked to this
    // bucket.
    ProtectionGroupCount          *int64                                        `json:"protection_group_count"`
    // The AWS replication output of the bucket.
    ReplicationSetting            *S3ReplicationOutput                          `json:"replication_setting"`
    // Size of bucket in bytes.
    SizeBytes                     *int64                                        `json:"size_bytes"`
    // A tag created through AWS console which can be applied to EBS volumes.
    Tags                          []*AwsTagModel                                `json:"tags"`
    // The AWS versioning output of the bucket.
    VersioningSetting             *S3VersioningOutput                           `json:"versioning_setting"`
}

// ReadComplianceConfigurationResponse represents a custom type struct for Success
type ReadComplianceConfigurationResponse struct {
    // If the `embed` query parameter is set, displays the responses of the related resource,
    // as defined by the embeddable link specified.
    Embedded     interface{}                   `json:"_embedded"`
    // ETag value
    Etag         *string                       `json:"_etag"`
    // URLs to pages related to the resource.
    Links        *ComplianceConfigurationLinks `json:"_links"`
    // The RFC3339 format time when the report configuration was created.
    Created      *string                       `json:"created"`
    // The user-provided description of the compliance report configuration.
    Description  *string                       `json:"description"`
    // The unique identifier of the report configuration.
    Id           *string                       `json:"id"`
    // Most recent report run generated from the report configuration.
    LatestRun    *LatestRun                    `json:"latest_run"`
    // The user-provided name of the compliance report configuration.
    Name         *string                       `json:"name"`
    // Notification channels to send the generated report runs.
    Notification *NotificationSetting          `json:"notification"`
    // Filter and control parameters of compliance report.
    Parameter    *Parameter                    `json:"parameter"`
    // When the report will be generated and sent. If the schedule is not provided then a
    // default value will be used.
    Schedule     *ScheduleSetting              `json:"schedule"`
}

// ReadConnectionGroupResponse represents a custom type struct for Success
type ReadConnectionGroupResponse struct {
    // Embedded responses related to the resource.
    Embedded                 interface{}           `json:"_embedded"`
    // The ETag value.
    Etag                     *string               `json:"_etag"`
    // URLs to pages related to the resource.
    Links                    *ConnectionGroupLinks `json:"_links"`
    // The alias given to the associated account in AWS.
    AccountName              *string               `json:"account_name"`
    // The AWS-assigned IDs of the accounts associated with the Connection Group.
    AccountNativeIds         []*string             `json:"account_native_ids"`
    // List of asset types connected via the connection-group.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS", "IcebergOnGlue", "IcebergOnS3Tables"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    AssetTypesEnabled        []*string             `json:"asset_types_enabled"`
    // The AWS regions associated with the with the Connection Group.
    AwsRegions               []*string             `json:"aws_regions"`
    // The consolidated configuration of the Clumio Cloud Protect and Clumio Cloud Discover products for this connection.
    // If this connection is deprecated to use unconsolidated configuration, then this field has a
    // value of `null`.
    Config                   *ConsolidatedConfig   `json:"config"`
    // The timestamp of when the connection was created.
    CreatedTimestamp         *string               `json:"created_timestamp"`
    // Clumio's S3 URL that contains the template to create the required resources in the
    // given account(s) according to the request.
    DeploymentTemplateUrl    *string               `json:"deployment_template_url"`
    // User-provided description for this connection group.
    Description              *string               `json:"description"`
    // Clumio assigned external ID for the connection group, should be used while creating the AWS stack.
    ExternalId               *string               `json:"external_id"`
    // The Clumio-assigned ID of the Connection Group, should be used as the token while creating the stack in AWS.
    Id                       *string               `json:"id"`
    // The AWS Account IDs that are intended to be associated with the Connection Group.
    IntendedAccountNativeIds []*string             `json:"intended_account_native_ids"`
    // THe asset types that are intended to be connected via connection-group.
    IntendedAssetTypes       []*string             `json:"intended_asset_types"`
    // The AWS regions that are intended to be connected with the Connection Group.
    IntendedAwsRegions       []*string             `json:"intended_aws_regions"`
    // The master account which manages the connection-group's stack.
    MasterAwsAccountId       *string               `json:"master_aws_account_id"`
    // The master region which manages the connection-group's stack.
    MasterRegion             *string               `json:"master_region"`
    // Ongoing Operation of the deployed and active stack of ConnectionGroup.
    OngoingStackOperation    *string               `json:"ongoing_stack_operation"`
    // The Clumio-assigned ID of the organizational unit associated with the
    // AWS environment. If this parameter is not provided, then the value
    // defaults to the first organizational unit assigned to the requesting
    // user. For more information about organizational units, refer to the
    // Organizational-Units documentation.
    OrganizationalUnitId     *string               `json:"organizational_unit_id"`
    // The Amazon Resource Name of the installed CloudFormation stack in AWS.
    StackArn                 *string               `json:"stack_arn"`
    // The name given to the installed CloudFormation stack in AWS.
    StackName                *string               `json:"stack_name"`
    // The status of the Connection Group based on the stack in associated AWS account.
    Status                   *string               `json:"status"`
}

// ReadConsolidatedAlertResponse represents a custom type struct for Success
type ReadConsolidatedAlertResponse struct {
    // The ETag value.
    Etag               *string                        `json:"_etag"`
    // URLs to pages related to the resource.
    Links              *ConsolidatedAlertLinks        `json:"_links"`
    // The number of currently active individual alerts associated with the consolidated alert.
    ActiveEntityCount  *int64                         `json:"active_entity_count"`
    // The issue that generated the alert. Each alert cause is associated with an alert type.
    Cause              *string                        `json:"cause"`
    // The number of cleared individual alerts associated with the consolidated alert.
    ClearedEntityCount *int64                         `json:"cleared_entity_count"`
    // The timestamp of when the consolidated alert was cleared, if ever. Represented in RFC-3339 format. If this alert has not been cleared, this field will have a value of `null`.
    // A consolidated alert goes into "cleared" status when all of its associated individual alerts are in "cleared" status or when a Clumio user manually clears it.
    ClearedTimestamp   *string                        `json:"cleared_timestamp"`
    // Additional information about the consolidated alert.
    Details            *ConsolidatedAlertDetails      `json:"details"`
    // The Clumio-assigned ID of the consolidated alert.
    Id                 *string                        `json:"id"`
    // A record of user-provided information about the alert.
    Notes              *string                        `json:"notes"`
    // The entity associated with or affected by the alert.
    ParentEntity       *ConsolidatedAlertParentEntity `json:"parent_entity"`
    // The timestamp of when the consolidated alert was initially raised. Represented in RFC-3339 format.
    RaisedTimestamp    *string                        `json:"raised_timestamp"`
    // The alert severity level. Values include "error" and "warning".
    Severity           *string                        `json:"severity"`
    // The consolidated alert status. A consolidated alert is in "active" status if one or more of its associated individual alerts is in "active" status.
    // A consolidated alert goes into "cleared" status when all of its associated individual alerts are in "cleared" status or when a Clumio user manually clears it.
    Status             *string                        `json:"status"`
    // The general alert category. An alert type may be associated with multiple alert causes. Examples of alert types include "tag_conflict" and "policy_violated".
    // Refer to the Alert Type table for a complete list of alert types.
    ClumioType         *string                        `json:"type"`
    // The timestamp of when the consolidated alert was last updated. Represented in RFC-3339 format.
    // Raising a new individual alert will update its associated consolidated alert.
    UpdatedTimestamp   *string                        `json:"updated_timestamp"`
}

// ReadDirectoryResponse represents a custom type struct for Success
type ReadDirectoryResponse struct {
    // Embedded responses related to the resource.
    Embedded     *DirectoryBrowseEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *DirectoryBrowseLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64                   `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64                   `json:"limit"`
    // The page token used to get this response.
    Start        *string                  `json:"start"`
}

// ReadDynamoDBTableBackupResponse represents a custom type struct for Success
type ReadDynamoDBTableBackupResponse struct {
    // The ETag value.
    Etag                      *string                   `json:"_etag"`
    // URLs to pages related to the resource.
    Links                     *DynamoDBTableBackupLinks `json:"_links"`
    // The AWS-assigned ID of the account associated with this database at the time of backup.
    AccountNativeId           *string                   `json:"account_native_id"`
    // The AWS region associated with this environment.
    AwsRegion                 *string                   `json:"aws_region"`
    // The billing mode of the DynamoDB table. Possible values are PROVISIONED or PAY_PER_REQUEST.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // configuration of source table if both 'billing_mode' and 'provisioned_throughput' are empty or `null`.
    BillingMode               *string                   `json:"billing_mode"`
    // Indicates whether DynamoDB Contributor Insights is enabled (true) or disabled (false)
    // on the table.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // value set in backup if `null`.
    ContributorInsightsStatus *bool                     `json:"contributor_insights_status"`
    // Indicates whether DynamoDB Deletion Protection is enabled (true) or disabled (false)
    // on the table.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // value set in backup if `null`.
    DeletionProtectionEnabled *bool                     `json:"deletion_protection_enabled"`
    // The timestamp of when this backup expires. Represented in RFC-3339 format.
    ExpirationTimestamp       *string                   `json:"expiration_timestamp"`
    // Represents the properties of a global secondary index.
    GlobalSecondaryIndexes    []*GlobalSecondaryIndex   `json:"global_secondary_indexes"`
    // Describes the version of global tables in use, if the table is replicated across AWS Regions. If the table
    // is not a global table, then this field has a value of `null`. Possible values are 2017.11.29 or 2019.11.21.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), the version is defaulted to 2019.11.21.
    GlobalTableVersion        *string                   `json:"global_table_version"`
    // The Clumio-assigned ID of the backup.
    Id                        *string                   `json:"id"`
    // The number of items in DynamoDB table backup.
    ItemCount                 *int64                    `json:"item_count"`
    // Represents the properties of a local secondary index.
    LocalSecondaryIndexes     []*LocalSecondaryIndex    `json:"local_secondary_indexes"`
    // Indicates whether DynamoDB Continuous Backup (PITR) is enabled (true) or disabled (false)
    // on the table.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // value set in backup if `null`.
    PitrStatus                *bool                     `json:"pitr_status"`
    // Represents the provisioned throughput settings for a DynamoDB table.
    ProvisionedThroughput     *ProvisionedThroughput    `json:"provisioned_throughput"`
    // Contains the details of the replica.
    Replicas                  []*ReplicaDescription     `json:"replicas"`
    // The size of the DynamoDB table backup in bytes.
    Size                      *int64                    `json:"size"`
    // Represents the server-side encryption settings for a table.
    SseSpecification          *SSESpecification         `json:"sse_specification"`
    // The timestamp of when this backup started. Represented in RFC-3339 format.
    StartTimestamp            *string                   `json:"start_timestamp"`
    // Represents the DynamoDB Streams configuration for a table in DynamoDB.
    // and the data type (`S` for string, `N` for number, `B` for binary).
    StreamSpecification       *StreamSpecification      `json:"stream_specification"`
    // The table class of the DynamoDB table. Possible values are STANDARD or STANDARD_INFREQUENT_ACCESS.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // STANDARD storage class if empty.
    TableClass                *string                   `json:"table_class"`
    // The Clumio-assigned ID of the DynamoDB table.
    TableId                   *string                   `json:"table_id"`
    // The name of the DynamoDB table.
    TableName                 *string                   `json:"table_name"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags                      []*AwsTagCommonModel      `json:"tags"`
    // The type of backup. Possible values include `clumio_backup` and `aws_snapshot`.
    ClumioType                *string                   `json:"type"`
}

// ReadDynamoDBTableResponse represents a custom type struct for Success
type ReadDynamoDBTableResponse struct {
    // Embedded responses related to the resource.
    Embedded                                      *DynamoDBTableEmbedded  `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                                         *DynamoDBTableLinks     `json:"_links"`
    // The AWS-assigned ID of the account associated with the DynamoDB table.
    AccountNativeId                               *string                 `json:"account_native_id"`
    // The AWS region associated with the DynamoDB table.
    AwsRegion                                     *string                 `json:"aws_region"`
    // The backup status information applied to this resource.
    BackupStatusInfo                              *BackupStatusInfo       `json:"backup_status_info"`
    // The billing mode of the DynamoDB table. Possible values are PROVISIONED or PAY_PER_REQUEST.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // configuration of source table if both 'billing_mode' and 'provisioned_throughput' are empty or `null`.
    BillingMode                                   *string                 `json:"billing_mode"`
    // Indicates whether DynamoDB Contributor Insights is enabled (true) or disabled (false)
    // on the table.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // value set in backup if `null`.
    ContributorInsightsStatus                     *bool                   `json:"contributor_insights_status"`
    // Indicates whether DynamoDB Deletion Protection is enabled (true) or disabled (false)
    // on the table.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // value set in backup if `null`.
    DeletionProtectionEnabled                     *bool                   `json:"deletion_protection_enabled"`
    // The timestamp of when the table was deleted. Represented in RFC-3339 format.
    // If this table has not been deleted, then this field has a value of `null`.
    DeletionTimestamp                             *string                 `json:"deletion_timestamp"`
    // The Clumio-assigned ID of the policy directly assigned to the entity.
    DirectAssignmentPolicyId                      *string                 `json:"direct_assignment_policy_id"`
    // The earliest continuous snapshot restorable time of the DynamoDB table for Point-in-time restore.
    // Represented in RFC-3339 format. If PITR is not enabled for the table, then this field has a value of `null`.
    EarliestContinuousSnapshotRestorableTimestamp *string                 `json:"earliest_continuous_snapshot_restorable_timestamp"`
    // The Clumio-assigned ID of the AWS environment associated with the DynamoDB table.
    EnvironmentId                                 *string                 `json:"environment_id"`
    // Represents the properties of a global secondary index.
    GlobalSecondaryIndexes                        []*GlobalSecondaryIndex `json:"global_secondary_indexes"`
    // Describes the version of global tables in use, if the table is replicated across AWS Regions. If the table
    // is not a global table, then this field has a value of `null`. Possible values are 2017.11.29 or 2019.11.21.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), the version is defaulted to 2019.11.21.
    GlobalTableVersion                            *string                 `json:"global_table_version"`
    // Determines whether the table has a direct assignment.
    HasDirectAssignment                           *bool                   `json:"has_direct_assignment"`
    // The Clumio-assigned ID of the DynamoDB table.
    Id                                            *string                 `json:"id"`
    // Determines whether the DynamoDB table has been deleted. If `true`, the table has been deleted.
    IsDeleted                                     *bool                   `json:"is_deleted"`
    // Determines whether the DynamoDB table is supported for backups.
    IsSupported                                   *bool                   `json:"is_supported"`
    // The number of items in the DynamoDB table.
    ItemCount                                     *int64                  `json:"item_count"`
    // The timestamp of the most recent snapshot of the DynamoDB table taken as part of
    // AwsSnapMgr. Represented in RFC-3339 format. If the table has never been
    // snapshotted, then this field has a value of `null`.
    LastSnapshotTimestamp                         *string                 `json:"last_snapshot_timestamp"`
    // The latest continuous snapshot restorable time of the DynamoDB table for Point-in-time restore.
    // Represented in RFC-3339 format. If PITR is not enabled for the table, then this field has a value of `null`.
    LatestContinuousSnapshotRestorableTimestamp   *string                 `json:"latest_continuous_snapshot_restorable_timestamp"`
    // Represents the properties of a local secondary index.
    LocalSecondaryIndexes                         []*LocalSecondaryIndex  `json:"local_secondary_indexes"`
    // The AWS-assigned name of the DynamoDB table.
    Name                                          *string                 `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the DynamoDB table.
    OrganizationalUnitId                          *string                 `json:"organizational_unit_id"`
    // Indicates whether DynamoDB Continuous Backup (PITR) is enabled (true) or disabled (false)
    // on the table.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // value set in backup if `null`.
    PitrStatus                                    *bool                   `json:"pitr_status"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo                                *ProtectionInfoWithRule `json:"protection_info"`
    // The protection status of the DynamoDB table. Possible values include "protected",
    // "unprotected", and "unsupported". If the DynamoDB table does not support backups, then
    // this field has a value of `unsupported`.
    ProtectionStatus                              *string                 `json:"protection_status"`
    // Represents the provisioned throughput settings for a DynamoDB table.
    ProvisionedThroughput                         *ProvisionedThroughput  `json:"provisioned_throughput"`
    // Contains the details of the replica.
    Replicas                                      []*ReplicaDescription   `json:"replicas"`
    // The size of the DynamoDB table. Measured in bytes (B).
    Size                                          *int64                  `json:"size"`
    // Represents the server-side encryption settings for a table.
    SseSpecification                              *SSESpecification       `json:"sse_specification"`
    // Represents the DynamoDB Streams configuration for a table in DynamoDB.
    // and the data type (`S` for string, `N` for number, `B` for binary).
    StreamSpecification                           *StreamSpecification    `json:"stream_specification"`
    // The AWS-assigned ARN of the DynamoDB table.
    TableArn                                      *string                 `json:"table_arn"`
    // The table class of the DynamoDB table. Possible values are STANDARD or STANDARD_INFREQUENT_ACCESS.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // STANDARD storage class if empty.
    TableClass                                    *string                 `json:"table_class"`
    // Represents the DynamoDB table keys.
    TableKeys                                     *DynamoDBKeys           `json:"table_keys"`
    // The AWS-assigned ID of the DynamoDB table.
    TableNativeId                                 *string                 `json:"table_native_id"`
    // The current state of the table.
    TableStatus                                   *string                 `json:"table_status"`
    // A tag created through AWS console which can be applied to EBS volumes.
    Tags                                          []*AwsTagModel          `json:"tags"`
    // The reason why protection is not available. If the table is supported,
    // then this field has a value of `null`.
    UnsupportedReason                             *string                 `json:"unsupported_reason"`
}

// ReadEBSBackupResponse represents a custom type struct for Success
type ReadEBSBackupResponse struct {
    // URLs to pages related to the resource.
    Links                *EBSBackupLinks      `json:"_links"`
    // The AWS-assigned ID of the account associated with the backup.
    AccountNativeId      *string              `json:"account_native_id"`
    // The availability zone associated with the volume backup. For example, `us-west-2a`.
    AwsAz                *string              `json:"aws_az"`
    // The AWS region in which the volume backup resides. For example, `us-west-2`.
    AwsRegion            *string              `json:"aws_region"`
    // Backup Tier
    BackupTier           *string              `json:"backup_tier"`
    // The reason that browsing is unavailable for the backup. Possible values include "file_limit_exceeded" and
    // "browsing_unavailable". If browse indexing is successful, then this field has a value of `null`.
    BrowsingFailedReason *string              `json:"browsing_failed_reason"`
    // The timestamp of when this backup expires. Represented in RFC-3339 format.
    ExpirationTimestamp  *string              `json:"expiration_timestamp"`
    // The Clumio-assigned ID of the volume backup.
    Id                   *string              `json:"id"`
    // Iops of the volume.
    Iops                 *int64               `json:"iops"`
    // Determines whether browsing is available for the backup. If `true`, then browsing is available for the backup.
    IsBrowsable          *bool                `json:"is_browsable"`
    // Determines whether the EBS volume backup is encrypted. If `true`, the volume backup is encrypted.
    IsEncrypted          *bool                `json:"is_encrypted"`
    // The AWS-assigned ID of the KMS key encrypting this EBS volume backup. If the volume is not encrypted, this field has a value of `null`.
    KmsKeyNativeId       *string              `json:"kms_key_native_id"`
    // The timestamp of when the migration was triggered. This field will be set only for
    // migration backups. Represented in RFC-3339 format.
    MigrationTimestamp   *string              `json:"migration_timestamp"`
    // The size of the volume backup. Measured in gigabytes (GB).
    Size                 *int64               `json:"size"`
    // The timestamp of when this backup started. Represented in RFC-3339 format.
    StartTimestamp       *string              `json:"start_timestamp"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags                 []*AwsTagCommonModel `json:"tags"`
    // The type of the backup. Possible values - `clumio_backup`, `aws_snapshot`.
    ClumioType           *string              `json:"type"`
    // Utilized size
    UtilizedSizeInBytes  *uint64              `json:"utilized_size_in_bytes"`
    // The Clumio-assigned ID of the EBS volume associated with the volume backup.
    VolumeId             *string              `json:"volume_id"`
    // The AWS-assigned ID of the EBS volume associated with the volume backup.
    VolumeNativeId       *string              `json:"volume_native_id"`
    // The volume type of the original EBS volume before backup. Possible values include `gp2`, `io1`, `st1`, `sc1`, `standard`.
    VolumeType           *string              `json:"volume_type"`
}

// ReadEBSBackupResponseV1 represents a custom type struct for Success
type ReadEBSBackupResponseV1 struct {
    // URLs to pages related to the resource.
    Links                *EBSBackupLinksV1    `json:"_links"`
    // The AWS-assigned ID of the account associated with the backup.
    AccountNativeId      *string              `json:"account_native_id"`
    // The availability zone associated with the volume backup. For example, `us-west-2a`.
    AwsAz                *string              `json:"aws_az"`
    // The AWS region in which the volume backup resides. For example, `us-west-2`.
    AwsRegion            *string              `json:"aws_region"`
    // The reason that browsing is unavailable for the backup. Possible values include "file_limit_exceeded" and
    // "browsing_unavailable". If browse indexing is successful, then this field has a value of `null`.
    BrowsingFailedReason *string              `json:"browsing_failed_reason"`
    // The timestamp of when this backup expires. Represented in RFC-3339 format.
    ExpirationTimestamp  *string              `json:"expiration_timestamp"`
    // The Clumio-assigned ID of the volume backup.
    Id                   *string              `json:"id"`
    // Determines whether browsing is available for the backup. If `true`, then browsing is available for the backup.
    IsBrowsable          *bool                `json:"is_browsable"`
    // Determines whether the EBS volume backup is encrypted. If `true`, the volume backup is encrypted.
    IsEncrypted          *bool                `json:"is_encrypted"`
    // The AWS-assigned ID of the KMS key encrypting this EBS volume backup. If the volume is not encrypted, this field has a value of `null`.
    KmsKeyNativeId       *string              `json:"kms_key_native_id"`
    // The size of the volume backup. Measured in gigabytes (GB).
    Size                 *int64               `json:"size"`
    // The timestamp of when this backup started. Represented in RFC-3339 format.
    StartTimestamp       *string              `json:"start_timestamp"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags                 []*AwsTagCommonModel `json:"tags"`
    // The volume type of the original EBS volume before backup. Possible values include `gp2`, `io1`, `st1`, `sc1`, `standard`.
    ClumioType           *string              `json:"type"`
    // The Clumio-assigned ID of the EBS volume associated with the volume backup.
    VolumeId             *string              `json:"volume_id"`
    // The AWS-assigned ID of the EBS volume associated with the volume backup.
    VolumeNativeId       *string              `json:"volume_native_id"`
}

// ReadEC2BackupResponse represents a custom type struct for Success
type ReadEC2BackupResponse struct {
    // URLs to pages related to the resource.
    Links                            *EC2BackupLinks                    `json:"_links"`
    // The AWS-assigned ID of the account associated with the backup.
    AccountNativeId                  *string                            `json:"account_native_id"`
    // An Amazon Machine Image is a supported and maintained image provided by AWS
    // that provides the information required to launch an instance.
    Ami                              *AmiModel                          `json:"ami"`
    // TODO: Add struct field description
    AttachedBackupEbsVolumes         []*AttachedEBSVolumeFullModel      `json:"attached_backup_ebs_volumes"`
    // The availability zone of the instance.
    AwsAz                            *string                            `json:"aws_az"`
    // The AWS region in which the instance backup resides. For example, `us-west-2`.
    AwsRegion                        *string                            `json:"aws_region"`
    // An Amazon Machine Image is a supported and maintained image provided by AWS
    // that provides the information required to launch an instance.
    BackupAmi                        *AmiModel                          `json:"backup_ami"`
    // The tier to which the backup is tagged to.
    BackupTier                       *string                            `json:"backup_tier"`
    // The reason that browsing is unavailable for the backup.
    // If browse indexing is successful, then this field has a value of `null`.
    BrowsingFailedReason             *string                            `json:"browsing_failed_reason"`
    // The timestamp of when this backup expires. Represented in RFC-3339 format.
    ExpirationTimestamp              *string                            `json:"expiration_timestamp"`
    // Denotes an IAM instance profile. An instance profile is a container for an
    // IAM role that you can use to pass role information to an EC2 instance when
    // the instance starts.
    IamInstanceProfile               *IamInstanceProfileModel           `json:"iam_instance_profile"`
    // The Clumio-assigned ID of the instance backup.
    Id                               *string                            `json:"id"`
    // The Clumio-assigned ID of the EC2 instance associated with the instance backup.
    InstanceId                       *string                            `json:"instance_id"`
    // The AWS-assigned ID of the EC2 instance associated with the instance backup.
    InstanceNativeId                 *string                            `json:"instance_native_id"`
    // TODO: Add struct field description
    InstanceStoreBlockDeviceMappings []*InstanceStoreBlockDeviceMapping `json:"instance_store_block_device_mappings"`
    // The instance type of the original EC2 instance before backup. For example, `m5.large`.
    InstanceType                     *string                            `json:"instance_type"`
    // Determines whether browsing is available for the backup. If `true`, then browsing is available for the backup.
    IsBrowsable                      *bool                              `json:"is_browsable"`
    // The name of the key pair associated with this instance. If this instance was not launched with an associated key pair, then this field has a value of `null`.
    KeyPairName                      *string                            `json:"key_pair_name"`
    // The ID of the key pair associated with this instance. If this instance was not launched with an associated key pair, then this field has a value of `null`.
    KeyPairNativeId                  *string                            `json:"key_pair_native_id"`
    // The timestamp of when the migration was triggered. This field will be set only for
    // migration backups. Represented in RFC-3339 format.
    MigrationTimestamp               *string                            `json:"migration_timestamp"`
    // TODO: Add struct field description
    NetworkInterfaces                []*NetworkInterface                `json:"network_interfaces"`
    // The public IP v4 address of the instance if one was assigned.
    PublicIpAddress                  *string                            `json:"public_ip_address"`
    // The size of the instance backup. This is the sum of all the EBS volumes attached to the EC2 measured in gigabytes (GB).
    Size                             *int64                             `json:"size"`
    // The timestamp of when this backup started. Represented in RFC-3339 format.
    StartTimestamp                   *string                            `json:"start_timestamp"`
    // The AWS-assigned Subnet ID of the EC2 instance.
    SubnetNativeId                   *string                            `json:"subnet_native_id"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags                             []*AwsTagCommonModel               `json:"tags"`
    // The type of the backup.
    ClumioType                       *string                            `json:"type"`
    // The total number of bytes written in all the disks of the EC2 instance.
    UtilizedSizeInBytes              *uint64                            `json:"utilized_size_in_bytes"`
    // The AWS-assigned ID of the VPC associated with the EC2 instance.
    VpcNativeId                      *string                            `json:"vpc_native_id"`
}

// ReadEC2MSSQLDatabaseBackupResponse represents a custom type struct for Success
type ReadEC2MSSQLDatabaseBackupResponse struct {
    // Embedded responses related to the resource.
    Embedded            *EC2MSSQLDatabaseBackupEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links               *EC2MSSQLDatabaseBackupLinks    `json:"_links"`
    // TODO: Add struct field description
    DatabaseFiles       []*MssqlDatabaseFile            `json:"database_files"`
    // The Clumio-assigned ID of the database associated with this backup.
    DatabaseId          *string                         `json:"database_id"`
    // The Microsoft SQL database engine at the time of backup.
    Engine              *string                         `json:"engine"`
    // The Microsoft SQL database engine version at the time of backup.
    EngineVersion       *string                         `json:"engine_version"`
    // The Clumio-assigned ID of the AWS environment associated with the database at the time of backup.
    EnvironmentId       *string                         `json:"environment_id"`
    // The timestamp of when this backup expires. Represented in RFC-3339 format.
    ExpirationTimestamp *string                         `json:"expiration_timestamp"`
    // The user-provided endpoint of the host containing the given database at the time of backup.
    HostEndpoint        *string                         `json:"host_endpoint"`
    // The Clumio-assigned ID of the host associated with the database at the time of backup.
    HostId              *string                         `json:"host_id"`
    // The Clumio-assigned ID of the backup.
    Id                  *string                         `json:"id"`
    // The Clumio-assigned instance id at the time of backup.
    InstanceId          *string                         `json:"instance_id"`
    // The instance name at the time of backup.
    InstanceName        *string                         `json:"instance_name"`
    // The timestamp of when this backup started. Represented in RFC-3339 format.
    StartTimestamp      *string                         `json:"start_timestamp"`
    // The type of backup.
    ClumioType          *string                         `json:"type"`
}

// ReadEC2MSSQLDatabaseResponse represents a custom type struct for Success
type ReadEC2MSSQLDatabaseResponse struct {
    // Embedded responses related to the resource.
    Embedded                                *EC2MSSQLDatabaseEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                                   *EC2MSSQLDatabaseLinks    `json:"_links"`
    // The AWS-assigned ID of the account associated with the EC2 instance the database resides in.
    AccountNativeId                         *string                   `json:"account_native_id"`
    // The Clumio-assigned ID of the availability group. It is null in case of a standalone database.
    AvailabilityGroupId                     *string                   `json:"availability_group_id"`
    // The Microsoft SQL assigned name of the availability group. It is null in case of a standalone database.
    AvailabilityGroupName                   *string                   `json:"availability_group_name"`
    // The AWS region associated with the EC2 instance the database resides in.
    AwsRegion                               *string                   `json:"aws_region"`
    // The backup status information applied to this resource.
    BackupStatusInfo                        *BackupStatusInfo         `json:"backup_status_info"`
    // The Clumio-assigned ID of the AWS environment associated with the EC2 MSSQL database.
    EnvironmentId                           *string                   `json:"environment_id"`
    // The Clumio-assigned ID of the failover cluster.
    FailoverClusterId                       *string                   `json:"failover_cluster_id"`
    // The Microsoft SQL assigned name of the Failover Cluster
    FailoverClusterName                     *string                   `json:"failover_cluster_name"`
    // Failovercluster Protection Status is used to indicate the fci protection status associated with the
    // fci database
    FailoverClusterProtectionStatus         *string                   `json:"failover_cluster_protection_status"`
    // The Clumio-assigned ID of the host connection containing the given database.
    HostConnectionId                        *string                   `json:"host_connection_id"`
    // The user-provided endpoint of the host containing the given database.
    HostEndpoint                            *string                   `json:"host_endpoint"`
    // The Clumio-assigned ID of the host containing the given database.
    HostId                                  *string                   `json:"host_id"`
    // The Clumio-assigned ID of the Database.
    Id                                      *string                   `json:"id"`
    // The Clumio-assigned ID of the instance containing the given database.
    InstanceId                              *string                   `json:"instance_id"`
    // The name of the Microsoft SQL instance containing the given database.
    InstanceName                            *string                   `json:"instance_name"`
    // is_supported is true if Clumio supports backup of the database.
    IsSupported                             *bool                     `json:"is_supported"`
    // The timestamp of the last time this database was full backed up.
    // Represented in RFC-3339 format. If this database has never been backed up,
    // this field has a value of `null`.
    LastBackupTimestamp                     *string                   `json:"last_backup_timestamp"`
    // The timestamp of the last time this database was log backed up in Bulk Recovery Model.
    // Represented in RFC-3339 format. If this database has never been backed up,
    // this field has a value of `null`.
    LastBulkRecoveryModelLogBackupTimestamp *string                   `json:"last_bulk_recovery_model_log_backup_timestamp"`
    // The timestamp of the last time this database was log backed up in Full Recovery Model.
    // Represented in RFC-3339 format. If this database has never been backed up,
    // this field has a value of `null`.
    LastFullRecoveryModelLogBackupTimestamp *string                   `json:"last_full_recovery_model_log_backup_timestamp"`
    // The name of the Database.
    Name                                    *string                   `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the database.
    OrganizationalUnitId                    *string                   `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo                          *ProtectionInfo           `json:"protection_info"`
    // recovery_model is the recovery model of the database. Possible values include 'simple_recovery_model',
    // 'bulk_recovery_model', and 'full_recovery_model'.
    RecoveryModel                           *string                   `json:"recovery_model"`
    // The size of the Database.
    Size                                    *float64                  `json:"size"`
    // The status of the database, Possible values include 'active' and 'inactive'.
    Status                                  *string                   `json:"status"`
    // The type of the database. Possible values include 'availability_group_database' and 'standalone_database'.
    ClumioType                              *string                   `json:"type"`
    // unsupported_reason is the reason why Clumio doesn't support backup of such database,
    // possible values include 'filestream_enabled_database'.
    UnsupportedReason                       *string                   `json:"unsupported_reason"`
}

// ReadEC2MSSQLFCIResponse represents a custom type struct for Success
type ReadEC2MSSQLFCIResponse struct {
    // Embedded responses related to the resource.
    Embedded             *EC2MSSQLFCIEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                *EC2MSSQLFCILinks    `json:"_links"`
    // The Clumio-assigned ID of the failover cluster.
    Id                   *string              `json:"id"`
    // The Microsoft SQL-assigned name of the failover cluster.
    Name                 *string              `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the FCI.
    OrganizationalUnitId *string              `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo       *ProtectionInfo      `json:"protection_info"`
    // ProtectionStatus of the FCI
    ProtectionStatus     *string              `json:"protection_status"`
    // The status of the FCI, Possible values include 'active' and 'inactive'.
    Status               *string              `json:"status"`
}

// ReadEC2MSSQLInstanceResponse represents a custom type struct for Success
type ReadEC2MSSQLInstanceResponse struct {
    // Embedded responses related to the resource.
    Embedded                       *EC2MSSQLInstanceEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                          *EC2MSSQLInstanceLinks    `json:"_links"`
    // The AWS-assigned ID of the account associated with the EC2 instance of the MSSQL instance.
    AccountNativeId                *string                   `json:"account_native_id"`
    // The AWS region associated with the EC2 instance of the MSSQL instance.
    AwsRegion                      *string                   `json:"aws_region"`
    // The Clumio-assigned ID of the AWS environment associated with the EC2 MSSQL instance.
    EnvironmentId                  *string                   `json:"environment_id"`
    // The boolean value represents if availability group is present in the instance.
    HasAssociatedAvailabilityGroup *bool                     `json:"has_associated_availability_group"`
    // The Clumio-assigned ID of the host connection containing the given mssql instance.
    HostConnectionId               *string                   `json:"host_connection_id"`
    // The user-provided endpoint of the host containing the given database.
    HostEndpoint                   *string                   `json:"host_endpoint"`
    // The Clumio-assigned ID of the host, containing the instance.
    HostId                         *string                   `json:"host_id"`
    // The Clumio-assigned ID of the Instance.
    Id                             *string                   `json:"id"`
    // The Microsoft SQL assigned name of the instance.
    Name                           *string                   `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the instance.
    OrganizationalUnitId           *string                   `json:"organizational_unit_id"`
    // Product Version of the instance.
    ProductVersion                 *string                   `json:"product_version"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo                 *ProtectionInfo           `json:"protection_info"`
    // The Microsoft SQL assigned server name of the instance.
    ServerName                     *string                   `json:"server_name"`
    // The status of the Instance, Possible values include 'active' and 'inactive'.
    Status                         *string                   `json:"status"`
}

// ReadEC2MSSQLInvHostResponse represents a custom type struct for Success
type ReadEC2MSSQLInvHostResponse struct {
    // Embedded responses related to the resource.
    Embedded                       interface{}           `json:"_embedded"`
    // EC2MSSQLInvHostLinks contains links related to ec2 mssql host
    // URLs to pages related to the resource.
    Links                          *EC2MSSQLInvHostLinks `json:"_links"`
    // The AWS-assigned ID of the account associated with the EC2 instance of the MSSQL host.
    AccountNativeId                *string               `json:"account_native_id"`
    // The AWS region associated with the EC2 instance of the MSSQL host.
    AwsRegion                      *string               `json:"aws_region"`
    // The Clumio-assigned ID of the host connection.
    ConnectionId                   *string               `json:"connection_id"`
    // The user-provided endpoint of the host containing the given database.
    Endpoint                       *string               `json:"endpoint"`
    // The Clumio-assigned ID of the AWS environment associated with the EC2 MSSQL host.
    EnvironmentId                  *string               `json:"environment_id"`
    // Determines whether or not an availability group is present in the host.
    HasAssociatedAvailabilityGroup *bool                 `json:"has_associated_availability_group"`
    // The Clumio-assigned ID of the Host.
    Id                             *string               `json:"id"`
    // The number of instances present in the host.
    InstanceCount                  *int64                `json:"instance_count"`
    // IsPartOfFCI is a boolean field representing if the Host is part of Failover Cluster
    IsPartOfFci                    *bool                 `json:"is_part_of_fci"`
    // The Clumio-assigned ID of the organizational unit associated with the host.
    OrganizationalUnitId           *string               `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo                 *ProtectionInfo       `json:"protection_info"`
    // The status of the Host, Possible values include 'active' and 'inactive'.
    Status                         *string               `json:"status"`
}

// ReadEC2MssqlAGResponse represents a custom type struct for Success
type ReadEC2MssqlAGResponse struct {
    // Embedded responses related to the resource.
    Embedded             *EC2MSSQLAGEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                *EC2MSSQLAGLinks    `json:"_links"`
    // The Clumio-assigned ID of the availability group.
    Id                   *string             `json:"id"`
    // The Microsoft SQL-assigned name of the availability group.
    Name                 *string             `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the availability group.
    OrganizationalUnitId *string             `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo       *ProtectionInfo     `json:"protection_info"`
    // The status of the availability group, Possible values include 'active' and 'inactive'.
    Status               *string             `json:"status"`
}

// ReadEbsVolumeResponse represents a custom type struct for Success
type ReadEbsVolumeResponse struct {
    // Embedded responses related to the resource.
    Embedded                 *EbsVolumeEmbedded      `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                    *EbsVolumeLinks         `json:"_links"`
    // The AWS-assigned ID of the account associated with the EBS volume.
    AccountNativeId          *string                 `json:"account_native_id"`
    // The AWS availability zone in which the EBS volume resides. For example,
    // `us-west-2a`.
    AwsAz                    *string                 `json:"aws_az"`
    // The AWS region associated with the EBS volume.
    AwsRegion                *string                 `json:"aws_region"`
    // The backup status information applied to this resource.
    BackupStatusInfo         *BackupStatusInfo       `json:"backup_status_info"`
    // The timestamp of when the volume was deleted. Represented in RFC-3339 format. If
    // this volume has not been deleted, then this field has a value of `null`.
    DeletionTimestamp        *string                 `json:"deletion_timestamp"`
    // The Clumio-assigned ID of the policy directly assigned to the entity.
    DirectAssignmentPolicyId *string                 `json:"direct_assignment_policy_id"`
    // The Clumio-assigned ID of the AWS environment associated with the EBS volume.
    EnvironmentId            *string                 `json:"environment_id"`
    // Determines whether the table has a direct assignment.
    HasDirectAssignment      *bool                   `json:"has_direct_assignment"`
    // The Clumio-assigned ID of the EBS volume.
    Id                       *string                 `json:"id"`
    // Iops of the volume.
    Iops                     *int64                  `json:"iops"`
    // Determines whether the EBS volume has been deleted. If `true`, the volume has been
    // deleted.
    IsDeleted                *bool                   `json:"is_deleted"`
    // Determines whether the EBS volume is encrypted. If `true`, the volume is encrypted.
    IsEncrypted              *bool                   `json:"is_encrypted"`
    // Determines whether the EBS volume is supported for backups.
    IsSupported              *bool                   `json:"is_supported"`
    // The AWS-assigned ID of the KMS key encrypting the EBS volume. If the volume is
    // unencrypted, then this field has a value of `null`.
    KmsKeyNativeId           *string                 `json:"kms_key_native_id"`
    // The timestamp of the most recent backup of the EBS volume. Represented in RFC-3339
    // format. If the volume has never been backed up, then this field has a value of
    // `null`.
    LastBackupTimestamp      *string                 `json:"last_backup_timestamp"`
    // The timestamp of the most recent snapshot of the EBS volume taken as part of
    // Snapshot Manager. Represented in RFC-3339 format. If the volume has never been
    // snapshotted, then this field has a value of `null`.
    LastSnapshotTimestamp    *string                 `json:"last_snapshot_timestamp"`
    // The AWS-assigned name of the EBS volume.
    Name                     *string                 `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the EBS volume.
    OrganizationalUnitId     *string                 `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo           *ProtectionInfoWithRule `json:"protection_info"`
    // The protection status of the EBS volume. Possible values include "protected",
    // "unprotected", and "unsupported". If the EBS volume does not support backups, then
    // this field has a value of `unsupported`. If the volume has been deleted, then this
    // field has a value of `null`.
    ProtectionStatus         *string                 `json:"protection_status"`
    // The size of the EBS volume. Measured in bytes (B).
    Size                     *int64                  `json:"size"`
    // A tag created through AWS console which can be applied to EBS volumes.
    Tags                     []*AwsTagModel          `json:"tags"`
    // The type of EBS volume. Possible values include "gp2", "io1", "st1", "sc1", and
    // "standard".
    ClumioType               *string                 `json:"type"`
    // The reason why protection is not available. If the volume is supported, then this
    // field has a value of `null`.
    UnsupportedReason        *string                 `json:"unsupported_reason"`
    // The AWS-assigned ID of the EBS volume.
    VolumeNativeId           *string                 `json:"volume_native_id"`
}

// ReadEc2InstanceResponse represents a custom type struct for Success
type ReadEc2InstanceResponse struct {
    // Embedded responses related to the resource.
    Embedded                 *Ec2InstanceEmbedded    `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                    *Ec2InstanceLinks       `json:"_links"`
    // The AWS-assigned ID of the account associated with the EC2 instance.
    AccountNativeId          *string                 `json:"account_native_id"`
    // The AWS availability zone in which the EC2 instance resides. For example,
    // `us-west-2a`.
    AwsAz                    *string                 `json:"aws_az"`
    // Determines whether the EC2 instance has been deleted. If `true`, the instance has
    // been deleted.
    AwsRegion                *string                 `json:"aws_region"`
    // The backup status information applied to this resource.
    BackupStatusInfo         *BackupStatusInfo       `json:"backup_status_info"`
    // The timestamp of when the instance was deleted. Represented in RFC-3339 format.
    // If this instance has not been deleted, then this field has a value of `null`.
    DeletionTimestamp        *string                 `json:"deletion_timestamp"`
    // The Clumio-assigned ID of the policy directly assigned to the entity.
    DirectAssignmentPolicyId *string                 `json:"direct_assignment_policy_id"`
    // EnaSupport indicates whether the Elastic Network Adapter (ENA) is enabled for the
    // instance.
    EnaSupport               *bool                   `json:"ena_support"`
    // The Clumio-assigned ID of the AWS environment associated with the EC2 instance.
    EnvironmentId            *string                 `json:"environment_id"`
    // Determines whether the table has a direct assignment.
    HasDirectAssignment      *bool                   `json:"has_direct_assignment"`
    // The Clumio-assigned ID of the EC2 instance.
    Id                       *string                 `json:"id"`
    // The AWS-assigned ID of the EC2 instance.
    InstanceNativeId         *string                 `json:"instance_native_id"`
    // Determines whether the EC2 instance has been deleted. If `true`, the instance has
    // been deleted.
    IsDeleted                *bool                   `json:"is_deleted"`
    // Determines whether the EC2 instance is supported for backups.
    IsSupported              *bool                   `json:"is_supported"`
    // The timestamp of the most recent backup of the EC2 instance. Represented in
    // RFC-3339 format. If the instance has never been backed up, then this field has a
    // value of `null`.
    LastBackupTimestamp      *string                 `json:"last_backup_timestamp"`
    // The timestamp of the most recent snapshot of the EC2 instance taken as part of the
    // EC2 Snapshot Manager. Represented in RFC-3339 format. If the instance has never
    // been backed up, then this field has a value of `null`.
    LastSnapshotTimestamp    *string                 `json:"last_snapshot_timestamp"`
    // The AWS-assigned name of the EC2 instance.
    Name                     *string                 `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the EC2 instance.
    OrganizationalUnitId     *string                 `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo           *ProtectionInfoWithRule `json:"protection_info"`
    // The protection status of the EC2 instance. Possible values include "protected",
    // "unprotected", and "unsupported". If the EC2 instance does not support backups,
    // then this field has a value of `unsupported`. If the instance has been deleted,
    // then this field has a value of `null`.
    ProtectionStatus         *string                 `json:"protection_status"`
    // The state of the EC2 instance. Possible values include: pending, running,
    // terminated, stopped, stopping, shutting-down, rebooting
    State                    *string                 `json:"state"`
    // The AWS Subnet ID of the EC2 instance
    SubnetId                 *string                 `json:"subnet_id"`
    // A tag created through AWS console which can be applied to EBS volumes.
    Tags                     []*AwsTagModel          `json:"tags"`
    // The AWS region associated with the EC2 instance. Possible instances types can be
    // found in: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html
    ClumioType               *string                 `json:"type"`
    // The reason why protection is not available. If the volume is supported,
    // then this field has a value of `null`.
    UnsupportedReason        *string                 `json:"unsupported_reason"`
    // AWS-assigned ID of the VPC associated with the EC2 instance.
    VpcId                    *string                 `json:"vpc_id"`
}

// ReadFileSystemResponse represents a custom type struct for Success
type ReadFileSystemResponse struct {
    // HateoasCommonLinks are the common fields for HATEOAS response.
    Links                *HateoasCommonLinks `json:"_links"`
    // The amount of available memory on the filesystem in bytes. Does not include
    // reserved memory.
    Available            *uint64             `json:"available"`
    // The filesystem UUID produced by the `lsblk` linux command. If this filesystem
    // was not given a UUID in the host environment, then this field has a value of
    // `null`.
    FilesystemNativeId   *string             `json:"filesystem_native_id"`
    // The Clumio-assigned ID of the filesystem.
    Id                   *string             `json:"id"`
    // The reason why file indexing failed. If file indexing succeeded, then this field
    // has a value of `null`. Possible values include "unsupported" and "encrypted".
    IndexingFailedReason *string             `json:"indexing_failed_reason"`
    // Determines whether the file system was encrypted.
    IsEncrypted          *bool               `json:"is_encrypted"`
    // Determines whether the file system has been indexed.
    // If `true`, file indexing completed successfully.
    IsIndexed            *bool               `json:"is_indexed"`
    // The location of this filesystem in the host environment. Only identifies mount
    // points that correspond to Windows drive letters. All other mount points are
    // identified by a '/'.
    MountPath            *string             `json:"mount_path"`
    // The number of files (including directories) indexed in the file system.
    NumFilesIndexed      *int64              `json:"num_files_indexed"`
    // The total amount of memory available to the filesystem in bytes.
    Size                 *uint64             `json:"size"`
    // The type of the filesystem. This field is populated with values returned from
    // the lsblk command. Possible values include `ntfs`, `xfs`, and `ext3`.
    ClumioType           *string             `json:"type"`
    // The amount of memory used by the filesystem in bytes.
    Used                 *uint64             `json:"used"`
}

// ReadGeneralSettingsResponseV2 represents a custom type struct for Success
type ReadGeneralSettingsResponseV2 struct {
    // URLs to pages related to the resource.
    Links                        *GeneralSettingsLinks `json:"_links"`
    // The length of time before a user is logged out of the Clumio system due to inactivity. Measured in seconds.
    // The valid range is between 600 seconds (10 minutes) and 3600 seconds (60 minutes).
    // If not configured, the value defaults to 900 seconds (15 minutes).
    AutoLogoutDuration           *int64                `json:"auto_logout_duration"`
    // The designated range of IP addresses that are allowed to access the Clumio REST API.
    // API requests that originate from outside this list will be blocked.
    // The IP address of the server from which this request is being made must be in this list; otherwise, the request will fail.
    // Set the parameter to individual IP addresses and/or a range of IP addresses in CIDR notation.
    // For example, ["193.168.1.0/24", "193.172.1.1"].
    // If not configured, the value defaults to ["0.0.0.0/0"] meaning all addresses will be allowed.
    IpAllowlist                  []*string             `json:"ip_allowlist"`
    // The grouping criteria for each datasource type.
    // These can only be edited for datasource types which do not have any
    // organizational units configured.
    OrganizationalUnitDataGroups *OUGroupingCriteria   `json:"organizational_unit_data_groups"`
    // The length of time a user password is valid before it must be changed. Measured in seconds.
    // The valid range is between 2592000 seconds (30 days) and 15552000 seconds (180 days).
    // If not configured, the value defaults to 7776000 seconds (90 days).
    PasswordExpirationDuration   *int64                `json:"password_expiration_duration"`
}

// ReadManagementGroupResponse represents a custom type struct for Success
type ReadManagementGroupResponse struct {
    // Etag
    Etag                  *string               `json:"_etag"`
    // URLs to pages related to the resource.
    Links                 *ManagementGroupLinks `json:"_links"`
    // Determines whether backups are allowed to occur across different subgroups or cloud connectors.
    BackupAcrossSubgroups *bool                 `json:"backup_across_subgroups"`
    // The Clumio-assigned ID of the management group.
    Id                    *string               `json:"id"`
    // The name of the management group.
    Name                  *string               `json:"name"`
    // The type of the management group. Possible values include `on_prem`.
    ClumioType            *string               `json:"type"`
    // The Clumio-assigned ID of the vCenter server associated with the management group.
    // All management groups are associated with a vCenter server.
    VcenterId             *string               `json:"vcenter_id"`
}

// ReadOrganizationalUnitResponse represents a custom type struct for Success
type ReadOrganizationalUnitResponse struct {
    // Embedded responses related to the resource.
    Embedded                  *EntityGroupEmbedded     `json:"_embedded"`
    // ETag value
    Etag                      *string                  `json:"_etag"`
    // URLs to pages related to the resource.
    Links                     *OrganizationalUnitLinks `json:"_links"`
    // Number of immediate children of the organizational unit.
    ChildrenCount             *int64                   `json:"children_count"`
    // Datasource types configured in this organizational unit. Possible values include `aws`, `microsoft365`, `vmware`, or `mssql`.
    ConfiguredDatasourceTypes []*string                `json:"configured_datasource_types"`
    // List of all recursive descendant organizational units of this OU.
    DescendantIds             []*string                `json:"descendant_ids"`
    // A description of the organizational unit.
    Description               *string                  `json:"description"`
    // The Clumio assigned ID of the organizational unit.
    Id                        *string                  `json:"id"`
    // Unique name assigned to the organizational unit.
    Name                      *string                  `json:"name"`
    // The Clumio assigned ID of the parent organizational unit.
    // The parent organizational unit contains the entities in this organizational unit and can update this organizational unit.
    // If this organizational unit is the global organizational unit, then this field has a value of `null`.
    ParentId                  *string                  `json:"parent_id"`
    // The Clumio-assigned ID of the task associated with this organizational unit.
    // The progress of the task can be monitored using the
    // [GET /tasks/{task_id}](#operation/read-task) endpoint.
    TaskId                    *string                  `json:"task_id"`
    // Number of users to whom this organizational unit or any of its descendants have been assigned.
    UserCount                 *int64                   `json:"user_count"`
    // A user along with a role.
    Users                     []*UserWithRole          `json:"users"`
}

// ReadOrganizationalUnitResponseV1 represents a custom type struct for Success
type ReadOrganizationalUnitResponseV1 struct {
    // Embedded responses related to the resource.
    Embedded                  *EntityGroupEmbedded     `json:"_embedded"`
    // ETag value
    Etag                      *string                  `json:"_etag"`
    // URLs to pages related to the resource.
    Links                     *OrganizationalUnitLinks `json:"_links"`
    // Number of immediate children of the organizational unit.
    ChildrenCount             *int64                   `json:"children_count"`
    // Datasource types configured in this organizational unit. Possible values include `aws`, `microsoft365`, `vmware`, or `mssql`.
    ConfiguredDatasourceTypes []*string                `json:"configured_datasource_types"`
    // List of all recursive descendant organizational units of this OU.
    DescendantIds             []*string                `json:"descendant_ids"`
    // A description of the organizational unit.
    Description               *string                  `json:"description"`
    // The Clumio assigned ID of the organizational unit.
    Id                        *string                  `json:"id"`
    // Unique name assigned to the organizational unit.
    Name                      *string                  `json:"name"`
    // The Clumio assigned ID of the parent organizational unit.
    // The parent organizational unit contains the entities in this organizational unit and can update this organizational unit.
    // If this organizational unit is the global organizational unit, then this field has a value of `null`.
    ParentId                  *string                  `json:"parent_id"`
    // The Clumio-assigned ID of the task associated with this organizational unit.
    // The progress of the task can be monitored using the
    // [GET /tasks/{task_id}](#operation/read-task) endpoint.
    TaskId                    *string                  `json:"task_id"`
    // Number of users to whom this organizational unit or any of its descendants have been assigned.
    UserCount                 *int64                   `json:"user_count"`
    // Users IDs to whom the organizational unit has been assigned.
    // This attribute will be available when reading a single OU and not when listing OUs.
    Users                     []*string                `json:"users"`
}

// ReadPolicyResponse represents a custom type struct for Success
type ReadPolicyResponse struct {
    // If the `embed` query parameter is set, displays the responses of the related resource,
    // as defined by the embeddable link specified.
    Embedded             *PolicyEmbedded    `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                *PolicyLinks       `json:"_links"`
    // The status of the policy.
    // Refer to the Policy Activation Status table for a complete list of policy statuses.
    ActivationStatus     *string            `json:"activation_status"`
    // The created time of the policy in unix time.
    CreatedTime          *int64             `json:"created_time"`
    // The Clumio-assigned ID of the policy.
    Id                   *string            `json:"id"`
    // The following table describes the possible lock statuses of a policy.
    // 
    // +----------+-------------------------------------------------------------------+
    // |  Status  |                            Description                            |
    // +==========+===================================================================+
    // | unlocked | Policies are unlocked until an update or deletion task is queued. |
    // +----------+-------------------------------------------------------------------+
    // | updating | During a policy edit, concurrent edits or deletion requests will  |
    // |          | be rejected.                                                      |
    // +----------+-------------------------------------------------------------------+
    // | deleting | During policy deletion, concurrent edits or deletion requests     |
    // |          | will be rejected.                                                 |
    // +----------+-------------------------------------------------------------------+
    // 
    LockStatus           *string            `json:"lock_status"`
    // The user-provided name of the policy.
    Name                 *string            `json:"name"`
    // TODO: Add struct field description
    Operations           []*PolicyOperation `json:"operations"`
    // The Clumio-assigned ID of the organizational unit associated with the policy.
    OrganizationalUnitId *string            `json:"organizational_unit_id"`
    // The policy-level timezone is deprecated, as the operation-level timezone should be used instead.
    // The timezone must be a valid location name from the IANA Time Zone database.
    // For instance, "America/New_York", "US/Central", "UTC".
    // Deprecated: true
    Timezone             *string            `json:"timezone"`
    // The updated time of the policy in unix time.
    UpdatedTime          *int64             `json:"updated_time"`
}

// ReadProtectionGroupBackupResponse represents a custom type struct for Success
type ReadProtectionGroupBackupResponse struct {
    // URLs to pages related to the resource.
    Links                  *ProtectionGroupBackupLinks `json:"_links"`
    // The number of objects in the protection group that were successfully backed up.
    BackedUpObjectCount    *int64                      `json:"backed_up_object_count"`
    // The total size in bytes of objects in the protection group that were
    // successfully backed up.
    BackedUpSizeBytes      *int64                      `json:"backed_up_size_bytes"`
    // The timestamp of when this backup expires. Represented in RFC-3339 format.
    ExpirationTimestamp    *string                     `json:"expiration_timestamp"`
    // The number of objects in the protection group that failed to be backed up.
    FailedObjectCount      *int64                      `json:"failed_object_count"`
    // The total size in bytes of objects in the protection group that failed
    // to be backed up.
    FailedSizeBytes        *int64                      `json:"failed_size_bytes"`
    // The Clumio-assigned ID of the protection group backup.
    Id                     *string                     `json:"id"`
    // The Clumio-assigned ID of the protection group.
    ProtectionGroupId      *string                     `json:"protection_group_id"`
    // The user-assigned name of the protection group.
    ProtectionGroupName    *string                     `json:"protection_group_name"`
    // The version of the protection group at the time the backup was taken.
    ProtectionGroupVersion *int64                      `json:"protection_group_version"`
    // The timestamp of when this backup started. Represented in RFC-3339 format.
    StartTimestamp         *string                     `json:"start_timestamp"`
    // The type of backup. Possible values include `protection_group_backup`.
    ClumioType             *string                     `json:"type"`
}

// ReadProtectionGroupResponse represents a custom type struct for Success
type ReadProtectionGroupResponse struct {
    // Embedded responses related to the resource.
    Embedded                         *ProtectionGroupEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                            *ProtectionGroupLinks    `json:"_links"`
    // Represents the aggregated stats for backup status.
    BackupStatusStats                *BackupStatusStats       `json:"backup_status_stats"`
    // The backup target AWS region associated with the protection group, empty if
    // in-region or not configured.
    BackupTargetAwsRegion            *string                  `json:"backup_target_aws_region"`
    // BackupTierStat
    BackupTierStats                  []*BackupTierStat        `json:"backup_tier_stats"`
    // Number of buckets
    BucketCount                      *int64                   `json:"bucket_count"`
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
    BucketRule                       *string                  `json:"bucket_rule"`
    // Creation time of the protection group in RFC-3339 format.
    CreatedTimestamp                 *string                  `json:"created_timestamp"`
    // The user-assigned description of the protection group.
    Description                      *string                  `json:"description"`
    // Timestamp of the earliest protection group backup which has not expired yet. Represented in
    // RFC-3339 format. Only available for Read API.
    EarliestAvailableBackupTimestamp *string                  `json:"earliest_available_backup_timestamp"`
    // The Clumio-assigned ID of the protection group.
    Id                               *string                  `json:"id"`
    // Whether the protection group already has a backup target configured by a policy, or
    // is open to be protected by an in-region or out-of-region S3 policy.
    IsBackupTargetRegionConfigured   *bool                    `json:"is_backup_target_region_configured"`
    // Determines whether the protection group is active or has been deleted. Deleted protection
    // groups may be purged after some time once there are no active backups associated with it.
    IsDeleted                        *bool                    `json:"is_deleted"`
    // Time of the last backup in RFC-3339 format.
    LastBackupTimestamp              *string                  `json:"last_backup_timestamp"`
    // Time of the last successful continuous backup in RFC-3339 format.
    LastContinuousBackupTimestamp    *string                  `json:"last_continuous_backup_timestamp"`
    // Modified time of the protection group in RFC-3339 format.
    ModifiedTimestamp                *string                  `json:"modified_timestamp"`
    // The user-assigned name of the protection group.
    Name                             *string                  `json:"name"`
    // ObjectFilter
    // defines which objects will be backed up.
    ObjectFilter                     *ObjectFilter            `json:"object_filter"`
    // The Clumio-assigned ID of the organizational unit associated with the Protection Group.
    OrganizationalUnitId             *string                  `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo                   *ProtectionInfoWithRule  `json:"protection_info"`
    // TODO: Add struct field description
    ProtectionStats                  *ProtectionStats         `json:"protection_stats"`
    // The protection status of the protection group. Possible values include "protected",
    // "unprotected", and "unsupported". If the protection group does not support backups, then
    // this field has a value of `unsupported`.
    ProtectionStatus                 *string                  `json:"protection_status"`
    // The list of AWS regions that this protection group is linked to
    Regions                          []*string                `json:"regions"`
    // Cumulative count of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpObjectCount         *int64                   `json:"total_backed_up_object_count"`
    // Cumulative size of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpSizeBytes           *int64                   `json:"total_backed_up_size_bytes"`
    // Version of the protection group. The version number is incremented every time
    // a change is made to the protection group.
    Version                          *int64                   `json:"version"`
}

// ReadProtectionGroupS3AssetBackupResponse represents a custom type struct for Success
type ReadProtectionGroupS3AssetBackupResponse struct {
    // URLs to pages related to the resource.
    Links                    *ProtectionGroupS3AssetBackupLinks `json:"_links"`
    // The AWS region in which the instance backup resides. For example, `us-west-2`.
    AwsRegion                *string                            `json:"aws_region"`
    // The number of objects in the protection group S3 asset that were successfully backed up.
    BackedUpObjectCount      *uint64                            `json:"backed_up_object_count"`
    // The total size in bytes of objects in the protection group S3 asset that were
    // successfully backed up.
    BackedUpSizeBytes        *uint64                            `json:"backed_up_size_bytes"`
    // The Clumio-assigned ID of the bucket.
    BucketId                 *string                            `json:"bucket_id"`
    // The name of the bucket.
    BucketName               *string                            `json:"bucket_name"`
    // The timestamp of when this backup expires. Represented in RFC-3339 format.
    ExpirationTimestamp      *string                            `json:"expiration_timestamp"`
    // The number of objects in the protection group S3 asset that failed to be backed up.
    FailedObjectCount        *uint64                            `json:"failed_object_count"`
    // The total size in bytes of objects in the protection group S3 asset that failed
    // to be backed up.
    FailedSizeBytes          *uint64                            `json:"failed_size_bytes"`
    // The Clumio-assigned ID of the protection group S3 asset backup.
    Id                       *string                            `json:"id"`
    // The Clumio-assigned ID of the protection group.
    ProtectionGroupId        *string                            `json:"protection_group_id"`
    // The Clumio-assigned ID of the protection group S3 asset.
    ProtectionGroupS3AssetId *string                            `json:"protection_group_s3_asset_id"`
    // The version of the protection group at the time the backup was taken.
    ProtectionGroupVersion   *int64                             `json:"protection_group_version"`
    // The timestamp of when this backup started. Represented in RFC-3339 format.
    StartTimestamp           *string                            `json:"start_timestamp"`
    // The type of backup. Possible values include `protection_group_s3_asset_backup`.
    ClumioType               *string                            `json:"type"`
}

// ReadProtectionGroupS3AssetContinuousBackupStatsResponse represents a custom type struct for Success
type ReadProtectionGroupS3AssetContinuousBackupStatsResponse struct {
    // ProtectionGroupBucketContinuousBackupStatsLinks
    // URLs to pages related to the resources.
    Links      *ProtectionGroupBucketContinuousBackupStatsLinks `json:"_links"`
    // ProtectionGroupBucketContinuousBackupStats
    Bins       []*ProtectionGroupBucketContinuousBackupStats    `json:"bins"`
    // ProtectionGroupBucketContinuousBackupStats
    TotalStats *ProtectionGroupBucketContinuousBackupStats      `json:"total_stats"`
}

// ReadProtectionGroupS3AssetResponse represents a custom type struct for Success
type ReadProtectionGroupS3AssetResponse struct {
    // Embedded responses related to the resource.
    Embedded                         *ProtectionGroupBucketEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                            *ProtectionGroupBucketLinks    `json:"_links"`
    // The AWS-assigned ID of the account associated with the DynamoDB table.
    AccountNativeId                  *string                        `json:"account_native_id"`
    // Whether this bucket was added to this protection group by the bucket rule
    AddedByBucketRule                *bool                          `json:"added_by_bucket_rule"`
    // Whether this bucket was added to this protection group by the user
    AddedByUser                      *bool                          `json:"added_by_user"`
    // The AWS region associated with the DynamoDB table.
    AwsRegion                        *string                        `json:"aws_region"`
    // The backup status information applied to this resource.
    BackupStatusInfo                 *BackupStatusInfo              `json:"backup_status_info"`
    // The backup target AWS region associated with the protection group S3 asset.
    BackupTargetAwsRegion            *string                        `json:"backup_target_aws_region"`
    // BackupTierStat
    BackupTierStats                  []*BackupTierStat              `json:"backup_tier_stats"`
    // The Clumio-assigned ID of the bucket
    BucketId                         *string                        `json:"bucket_id"`
    // The name of the bucket
    BucketName                       *string                        `json:"bucket_name"`
    // Creation time of the protection group in RFC-3339 format.
    CreatedTimestamp                 *string                        `json:"created_timestamp"`
    // Timestamp of the earliest protection group backup which has not expired yet. Represented in
    // RFC-3339 format. Only available for Read API.
    EarliestAvailableBackupTimestamp *string                        `json:"earliest_available_backup_timestamp"`
    // The Clumio-assigned ID of the AWS environment associated with the protection group.
    EnvironmentId                    *string                        `json:"environment_id"`
    // The Clumio-assigned ID of the protection group
    GroupId                          *string                        `json:"group_id"`
    // The name of the protection group
    GroupName                        *string                        `json:"group_name"`
    // The Clumio-assigned ID that represents the bucket within the protection group.
    Id                               *string                        `json:"id"`
    // Determines whether the protection group bucket has been deleted
    IsDeleted                        *bool                          `json:"is_deleted"`
    // Time of the last backup in RFC-3339 format.
    LastBackupTimestamp              *string                        `json:"last_backup_timestamp"`
    // Time of the last successful continuous backup in RFC-3339 format.
    LastContinuousBackupTimestamp    *string                        `json:"last_continuous_backup_timestamp"`
    // The Clumio-assigned ID of the organizational unit associated with the protection group.
    OrganizationalUnitId             *string                        `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo                   *ProtectionInfoWithRule        `json:"protection_info"`
    // The protection status of the protection group. Possible values include "protected",
    // "unprotected", and "unsupported". If the protection group does not support backups, then
    // this field has a value of `unsupported`.
    ProtectionStatus                 *string                        `json:"protection_status"`
    // Cumulative count of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpObjectCount         *int64                         `json:"total_backed_up_object_count"`
    // Cumulative size of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpSizeBytes           *int64                         `json:"total_backed_up_size_bytes"`
    // The unsupported reason for the S3 bucket.
    UnsupportedReason                *string                        `json:"unsupported_reason"`
}

// ReadRDSDatabaseTableColumnsResponse represents a custom type struct for Success
type ReadRDSDatabaseTableColumnsResponse struct {
    // URLs to pages related to the resource.
    Links   *RDSDatabaseTableColumnLinks `json:"_links"`
    // RDSDatabaseTableColumn denotes the model for rds database column
    Columns []*RDSDatabaseTableColumn    `json:"columns"`
}

// ReadRDSDatabaseTableResponse represents a custom type struct for Success
type ReadRDSDatabaseTableResponse struct {
    // Embedded responses related to the resource.
    Embedded *RDSDatabaseTableEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *RDSDatabaseTableLinks    `json:"_links"`
    // The name of the table within the specified RDS database.
    Name     *string                   `json:"name"`
}

// ReadRdsDatabaseBackupResponse represents a custom type struct for Success
type ReadRdsDatabaseBackupResponse struct {
    // URLs to pages related to the resource.
    Links                  *RdsDatabaseBackupLinks `json:"_links"`
    // The AWS-assigned ID of the account associated with this database at the time of backup.
    AccountNativeId        *string                 `json:"account_native_id"`
    // The AWS availability zones associated with this database at the time of backup.
    AwsAzs                 []*string               `json:"aws_azs"`
    // The AWS region associated with this environment.
    AwsRegion              *string                 `json:"aws_region"`
    // The AWS-assigned ID of the database at the time of backup.
    DatabaseNativeId       *string                 `json:"database_native_id"`
    // The AWS database engine at the time of backup.
    Engine                 *string                 `json:"engine"`
    // The aws database engine version at the time of backup.
    EngineVersion          *string                 `json:"engine_version"`
    // The timestamp of when this backup expires. Represented in RFC-3339 format.
    ExpirationTimestamp    *string                 `json:"expiration_timestamp"`
    // The Clumio-assigned ID of the backup.
    Id                     *string                 `json:"id"`
    // TODO: Add struct field description
    Instances              []*RdsInstanceModel     `json:"instances"`
    // The AWS-assigned ID of the KMS key associated with this database at the time of backup.
    KmsKeyNativeId         *string                 `json:"kms_key_native_id"`
    // The timestamp of when the migration was triggered. This field will be set only for
    // migration granular backups. Represented in RFC-3339 format.
    MigrationTimestamp     *string                 `json:"migration_timestamp"`
    // Option group name associated with the backed up RDS resource
    OptionGroupName        *string                 `json:"option_group_name"`
    // The Clumio-assigned ID of the database associated with this backup.
    ResourceId             *string                 `json:"resource_id"`
    // The type of the RDS resource associated with this backup. Possible values include `aws_rds_cluster` and `aws_rds_instance`.
    ResourceType           *string                 `json:"resource_type"`
    // The AWS-assigned IDs of the security groups associated with this RDS resource backup.
    SecurityGroupNativeIds []*string               `json:"security_group_native_ids"`
    // The size of the RDS resource backup. Measured in bytes (B).
    Size                   *int64                  `json:"size"`
    // The timestamp of when this backup started. Represented in RFC-3339 format.
    StartTimestamp         *string                 `json:"start_timestamp"`
    // The AWS-assigned name of the subnet group associated with this RDS resource backup.
    SubnetGroupName        *string                 `json:"subnet_group_name"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags                   []*AwsTagCommonModel    `json:"tags"`
    // The type of backup. Possible values include `clumio_snapshot` and `granular_backup`.
    ClumioType             *string                 `json:"type"`
}

// ReadRdsResourceResponse represents a custom type struct for Success
type ReadRdsResourceResponse struct {
    // Embedded responses related to the resource.
    Embedded                               *RdsResourceEmbedded    `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                                  *RdsResourceLinks       `json:"_links"`
    // The AWS-assigned ID of the account associated with this resource.
    AccountNativeId                        *string                 `json:"account_native_id"`
    // The AWS availability zone(s) associated with the resource. For example, `us-west-2a`.
    AwsAzs                                 []*string               `json:"aws_azs"`
    // The AWS region associated with this resource.
    AwsRegion                              *string                 `json:"aws_region"`
    // The backup status information applied to this resource.
    BackupStatusInfo                       *BackupStatusInfo       `json:"backup_status_info"`
    // The timestamp of when the RDS resource was deleted. Represented in RFC-3339 format.
    // If the resource was not deleted, then this field has a value of `null`.
    DeletionTimestamp                      *string                 `json:"deletion_timestamp"`
    // The Clumio-assigned ID of the policy directly assigned to the entity.
    DirectAssignmentPolicyId               *string                 `json:"direct_assignment_policy_id"`
    // The timestamp of the oldest AWS snapshot of the RDS resource. Represented in RFC-3339
    // format. If the resource has no available snapshots, then this field has a value of `null`.
    EarliestAwsSnapshotRestorableTimestamp *string                 `json:"earliest_aws_snapshot_restorable_timestamp"`
    // The database engine of the RDS resource. Possible values include `postgres` and `mysql`.
    // For a full list of possible values, please refer to the AWS documentation at
    // https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html
    Engine                                 *string                 `json:"engine"`
    // The database engine mode of the RDS resource. Possible values include `provisioned`
    // and `serverless`.
    EngineMode                             *string                 `json:"engine_mode"`
    // The database engine version of the RDS resource. For example, `10.12`.
    EngineVersion                          *string                 `json:"engine_version"`
    // The Clumio-assigned ID of the AWS environment associated with this resource.
    EnvironmentId                          *string                 `json:"environment_id"`
    // The timestamp of the first active backup of the database to Clumio. Represented
    // in RFC-3339 format.
    FirstClumioSnapshotTimestamp           *string                 `json:"first_clumio_snapshot_timestamp"`
    // The timestamp of the first active granular backup for the database. Represented in
    // RFC-3339 format.
    FirstGranularBackupTimestamp           *string                 `json:"first_granular_backup_timestamp"`
    // Determines whether the table has a direct assignment.
    HasDirectAssignment                    *bool                   `json:"has_direct_assignment"`
    // The Clumio-assigned ID of the resource.
    Id                                     *string                 `json:"id"`
    // Determines whether an RDS resource is deleted.
    IsDeleted                              *bool                   `json:"is_deleted"`
    // Determines whether an RDS resource is encrypted.
    IsEncrypted                            *bool                   `json:"is_encrypted"`
    // Determines whether the RDS resource is supported for backups.
    IsSupported                            *bool                   `json:"is_supported"`
    // The AWS-assigned ID of the KMS key encrypting this resource. If the resource is
    // unencrypted, then this field has a value of `null`.
    KmsKeyNativeId                         *string                 `json:"kms_key_native_id"`
    // The timestamp of the last time this database was backed up to Clumio. Represented
    // in RFC-3339 format.
    LastClumioSnapshotTimestamp            *string                 `json:"last_clumio_snapshot_timestamp"`
    // The timestamp of the last time this database had granular backup performed.
    // Represented in RFC-3339 format.
    LastGranularBackupTimestamp            *string                 `json:"last_granular_backup_timestamp"`
    // The timestamp of the newest AWS snapshot of the RDS resource. Represented in RFC-3339
    // format. If the resource has no available snapshots, then this field has a value of `null`.
    LatestAwsSnapshotRestorableTimestamp   *string                 `json:"latest_aws_snapshot_restorable_timestamp"`
    // The AWS-assigned name of the RDS resource. For example, `clumio-aurora-dev`.
    Name                                   *string                 `json:"name"`
    // The organizational unit to which this resource belongs.
    OrganizationalUnitId                   *string                 `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo                         *ProtectionInfoWithRule `json:"protection_info"`
    // The protection status of the RDS resource. Possible values include `protected`,
    // `unprotected`, and `unsupported`. If the RDS resource does not support backups, then
    // this field has a value of `unsupported`. If the resource has been deleted, then this
    // field has a value of `null`.
    ProtectionStatus                       *string                 `json:"protection_status"`
    // The AWS-assigned ID of the RDS resource. For example, `cluster-3WW6IXRWO5ZS4PTUIKGZEACISY`.
    ResourceNativeId                       *string                 `json:"resource_native_id"`
    // The AWS-assigned IDs of the security groups associated with this resource
    SecurityGroupNativeIds                 []*string               `json:"security_group_native_ids"`
    // The size of the RDS resource. Measured in bytes (B).
    Size                                   *int64                  `json:"size"`
    // The RDS subnet group name associated with this resource.
    SubnetGroupName                        *string                 `json:"subnet_group_name"`
    // A tag created through AWS console which can be applied to EBS volumes.
    Tags                                   []*AwsTagModel          `json:"tags"`
    // The RDS resource type. Possible values include `aws_rds_cluster` and `aws_rds_instance`.
    ClumioType                             *string                 `json:"type"`
    // The reason why protection is not available on this RDS resource, if any.
    // Possible values include `rds_engine_oracle` and `rds_postgres_9_4`.
    // If the resource is supported, then this field has a value of `null`.
    UnsupportedReason                      *string                 `json:"unsupported_reason"`
}

// ReadRoleResponse represents a custom type struct for Success
type ReadRoleResponse struct {
    // ETag value
    Etag        *string            `json:"_etag"`
    // URLs to pages related to the resource.
    Links       *RoleLinks         `json:"_links"`
    // A description of the role.
    Description *string            `json:"description"`
    // The Clumio-assigned ID of the role.
    Id          *string            `json:"id"`
    // Unique name assigned to the role.
    Name        *string            `json:"name"`
    // TODO: Add struct field description
    Permissions []*PermissionModel `json:"permissions"`
    // Number of users to whom the role has been assigned.
    UserCount   *int64             `json:"user_count"`
}

// ReadRuleResponse represents a custom type struct for Success
type ReadRuleResponse struct {
    // Embedded responses related to the resource.
    Embedded             *RuleEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                *RuleLinks    `json:"_links"`
    // An action to be applied subject to the rule criteria.
    Action               *RuleAction   `json:"action"`
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
    Condition            *string       `json:"condition"`
    // The Clumio-assigned ID of the policy rule.
    Id                   *string       `json:"id"`
    // Name of the rule. Max 100 characters.
    Name                 *string       `json:"name"`
    // The Clumio-assigned ID of the organizational unit (OU) to which the policy rule belongs.
    OrganizationalUnitId *string       `json:"organizational_unit_id"`
    // A priority relative to other rules.
    Priority             *RulePriority `json:"priority"`
}

// ReadS3InstantAccessEndpointResponse represents a custom type struct for Success
type ReadS3InstantAccessEndpointResponse struct {
    // Embedded responses related to the resource.
    Embedded                 *S3InstantAccessEndpointEmbedded `json:"_embedded"`
    // The ETag value.
    Etag                     *string                          `json:"_etag"`
    // URLs to pages related to the resource.
    Links                    *S3InstantAccessEndpointLinks    `json:"_links"`
    // The AWS-assigned ID of the account associated with the S3 instant access endpoint.
    AwsAccountId             *string                          `json:"aws_account_id"`
    // The AWS region of the S3 instant access endpoint and its source backup.
    BackupRegion             *string                          `json:"backup_region"`
    // The name of source bucket.
    BucketName               *string                          `json:"bucket_name"`
    // The time that this endpoint was created, in RFC-3339 format.
    CreatedTimestamp         *string                          `json:"created_timestamp"`
    // The status of the S3 instant access endpoint. Possible values include "preparing",
    // "active", "expiring", and "expired".
    EndpointStatus           *string                          `json:"endpoint_status"`
    // The time that this endpoint expires, in RFC-3339 format.
    ExpiryTimestamp          *string                          `json:"expiry_timestamp"`
    // The Clumio-assigned ID of the S3 instant access endpoint.
    Id                       *string                          `json:"id"`
    // The user-assigned name of the S3 instant access endpoint.
    Name                     *string                          `json:"name"`
    // The time in RFC-3339 format that the restored objects are backed up from.
    ObjectsCreatedAfter      *string                          `json:"objects_created_after"`
    // The time in RFC-3339 format that the restored objects are backed up to.
    ObjectsCreatedBefore     *string                          `json:"objects_created_before"`
    // The Clumio-assigned ID of the organizational unit associated with the S3 instant access endpoint.
    OrganizationalUnitId     *string                          `json:"organizational_unit_id"`
    // The Clumio-assigned ID of the protection group this endpoint is created for.
    ProtectionGroupId        *string                          `json:"protection_group_id"`
    // The user-assigned name of the protection group this endpoints is created for.
    ProtectionGroupName      *string                          `json:"protection_group_name"`
    // The Clumio-assigned ID of the bucket protection group.
    ProtectionGroupS3AssetId *string                          `json:"protection_group_s3_asset_id"`
    // The AWS region of the source bucket.
    Region                   *string                          `json:"region"`
    // The time at which the backup was restored from this endpoint in RFC-3339 format.
    // Deprecated.
    RestoreTimestamp         *string                          `json:"restore_timestamp"`
    // IAM role which is allowed access to the OLAP endpoint.
    Roles                    []*S3InstantAccessEndpointRole   `json:"roles"`
    // S3InstantAccessEndpointStat
    // Statistical metric related to the instant access endpoint.
    // S3InstantAccessEndpointStat swagger: model S3InstantAccessEndpointStat
    Stats                    []*S3InstantAccessEndpointStat   `json:"stats"`
    // The time that this endpoint was last updated, in RFC-3339 format.
    UpdatedTimestamp         *string                          `json:"updated_timestamp"`
}

// ReadS3InstantAccessEndpointRolePermissionResponse represents a custom type struct for Success
type ReadS3InstantAccessEndpointRolePermissionResponse struct {
    // Embedded responses related to the resource.
    Embedded    *S3InstantAccessEndpointEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links       *S3InstantAccessEndpointLinks    `json:"_links"`
    // The permissions JSON string to be attached to the user's IAM role to allow access to the
    // Instant Access endpoint.
    Permissions *string                          `json:"permissions"`
}

// ReadS3InstantAccessEndpointUriResponse represents a custom type struct for Success
type ReadS3InstantAccessEndpointUriResponse struct {
    // Embedded responses related to the resource.
    Embedded                           *S3InstantAccessEndpointEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                              *S3InstantAccessEndpointLinks    `json:"_links"`
    // An alias of the endpoint bucket.
    BucketAlias                        *string                          `json:"bucket_alias"`
    // An Origin Domain form of the endpoint URI for CloudFront distribution.
    CloudfrontDistributionOriginDomain *string                          `json:"cloudfront_distribution_origin_domain"`
    // The URI of the endpoint.
    EndpointUri                        *string                          `json:"endpoint_uri"`
    // The AWS region the endpoint is located in.
    Region                             *string                          `json:"region"`
}

// ReadTaskResponse represents a custom type struct for Success
type ReadTaskResponse struct {
    // The ETag value.
    Etag               *string            `json:"_etag"`
    // URLs to pages related to the resource.
    Links              *TaskLinks         `json:"_links"`
    // The task category. Examples of task types include "backup", "restore", "snapshot", and "system".
    // 
    // +-------------------+----------------------------------------------------------+
    // |     Category      |                       Description                        |
    // +===================+==========================================================+
    // | backup            | Encompasses all modes of backups. This does not include  |
    // |                   | in-account snapshots.                                    |
    // +-------------------+----------------------------------------------------------+
    // | restore           | Encompasses all modes of restores. This does not include |
    // |                   | restores of in-account snapshots.                        |
    // +-------------------+----------------------------------------------------------+
    // | snapshot          | Encompasses all modes of in-account snapshots.           |
    // +-------------------+----------------------------------------------------------+
    // | snapshot_restore  | Encompasses all modes of snapshot restores.              |
    // +-------------------+----------------------------------------------------------+
    // | system            | Encompasses a variety of system-initiated tasks, such as |
    // |                   | aws_rds_backup_target_setup and                          |
    // |                   | aws_ec2_instance_backup_indexing.                        |
    // +-------------------+----------------------------------------------------------+
    // | report_generation | Encompasses task types which generate reports, such as   |
    // |                   | activity_report_file_download.                           |
    // +-------------------+----------------------------------------------------------+
    // | management        | Encompasses user-initiated tasks which manage Clumio     |
    // |                   | resources, such as organizational_unit_update and        |
    // |                   | policy_update.                                           |
    // +-------------------+----------------------------------------------------------+
    // 
    Category           *string            `json:"category"`
    // The timestamp of when the task was created. Represented in RFC-3339 format.
    CreatedTimestamp   *string            `json:"created_timestamp"`
    // The timestamp of when the task ended. If this task has not yet ended,
    // then this field has a value of `null`. Represented in RFC-3339 format.
    EndTimestamp       *string            `json:"end_timestamp"`
    // The task genre. A genre is a high-level collection of task categories.
    // 
    // +----------------+-------------------------------------------------------------+
    // |     Genre      |                         Description                         |
    // +================+=============================================================+
    // | operational    | Encompasses all backup, restore, snapshot, and              |
    // |                | snapshot_restore tasks.                                     |
    // +----------------+-------------------------------------------------------------+
    // | administrative | Encompasses management, system, and report_generation       |
    // |                | tasks.                                                      |
    // +----------------+-------------------------------------------------------------+
    // 
    Genre              *string            `json:"genre"`
    // The Clumio-assigned ID of the task.
    Id                 *string            `json:"id"`
    // Determines whether or not this task can be aborted.
    // A task can be aborted if its status is either "queued" or "in_progress".
    // Tasks of certain types including "aws_ebs_volume_backup_indexing" cannot be aborted.
    IsAbortable        *bool              `json:"is_abortable"`
    // The parent entity associated with the task.
    ParentEntity       *TaskParentEntity  `json:"parent_entity"`
    // The primary entity associated with the task.
    PrimaryEntity      *TaskPrimaryEntity `json:"primary_entity"`
    // The percentage progress of task completion. Measured as an integer value between 0 and 100.
    ProgressPercentage *int64             `json:"progress_percentage"`
    // The timestamp of when the task started. If this task has not started yet,
    // then this field has a value of `null`. Represented in RFC-3339 format.
    StartTimestamp     *string            `json:"start_timestamp"`
    // The task status. Examples of task statuses include, "queued", "in_progress", and "completed".
    // Refer to the Task Status table for a complete list of task statuses.
    Status             *string            `json:"status"`
    // Refer to the Task Type table for a complete list of task types.
    ClumioType         *string            `json:"type"`
}

// ReadUserResponse represents a custom type struct for Success
type ReadUserResponse struct {
    // Embedded responses related to the resource.
    Embedded                   *UserEmbedded                 `json:"_embedded"`
    // ETag value
    Etag                       *string                       `json:"_etag"`
    // URLs to pages related to the resource.
    Links                      *UserLinks                    `json:"_links"`
    // The organizational units assigned to the user, with the specified role.
    AccessControlConfiguration []*RoleForOrganizationalUnits `json:"access_control_configuration"`
    // The email address of the Clumio user.
    Email                      *string                       `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                   *string                       `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                         *string                       `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                    *string                       `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                *bool                         `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                  *bool                         `json:"is_enabled"`
    // The timestamp of when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp      *string                       `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount    *int64                        `json:"organizational_unit_count"`
}

// ReadUserResponseV1 represents a custom type struct for Success
type ReadUserResponseV1 struct {
    // Embedded responses related to the resource.
    Embedded                      *UserEmbeddedV1 `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                         *UserLinks      `json:"_links"`
    // The list of organizational unit IDs assigned to the user.
    // This attribute will be available when reading a single user and not when listing users.
    AssignedOrganizationalUnitIds []*string       `json:"assigned_organizational_unit_ids"`
    // Assigned Role for the user.
    AssignedRole                  *string         `json:"assigned_role"`
    // The email address of the Clumio user.
    Email                         *string         `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                      *string         `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                            *string         `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                       *string         `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                   *bool           `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                     *bool           `json:"is_enabled"`
    // The timestamp of when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp         *string         `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount       *int64          `json:"organizational_unit_count"`
}

// ReadWalletResponse represents a custom type struct for Success
type ReadWalletResponse struct {
    // Embedded responses related to the resource.
    Embedded           interface{}            `json:"_embedded"`
    // URLs to pages related to the resource.
    Links              *WalletLinks           `json:"_links"`
    // AWS Account ID associated with the wallet.
    AccountNativeId    *string                `json:"account_native_id"`
    // Version of the template available
    AvailableVersion   *int64                 `json:"available_version"`
    // The AWS region associated with the wallet.
    AwsRegion          *string                `json:"aws_region"`
    // Clumio AWS Account ID.
    ClumioAwsAccountId *string                `json:"clumio_aws_account_id"`
    // DeploymentURL is an (external) link to an AWS console page for quick-creation
    // of the stack.
    DeploymentUrl      *string                `json:"deployment_url"`
    // ErrorCode is a short string describing the error, if any.
    ErrorCode          *string                `json:"error_code"`
    // ErrorMessage is a longer description explaining the error, if any, and how to
    // fix it.
    ErrorMessage       *string                `json:"error_message"`
    // The Clumio-assigned ID of the wallet.
    Id                 *string                `json:"id"`
    // The regions where the wallet is installed.
    InstalledRegions   []*string              `json:"installed_regions"`
    // TODO: Add struct field description
    KeyErrors          map[string]*ErrorModel `json:"key_errors"`
    // RoleArn is the AWS Resource Name of the IAM Role created by the stack.
    RoleArn            *string                `json:"role_arn"`
    // The version of the stack used or being used.
    StackVersion       *int64                 `json:"stack_version"`
    // State describes the state of the wallet. Valid states are:
    // Waiting: The wallet has been created, but a stack hasn't been created. The
    // wallet can't be used in this state.
    // Enabled: The wallet has been created and a stack has been created for the
    // wallet. This is the normal expected state of a wallet in use.
    // Error:   The wallet is inaccessible. See ErrorCode and ErrorMessage fields for
    // additional details.
    State              *string                `json:"state"`
    // The supported regions for the wallet.
    SupportedRegions   []*string              `json:"supported_regions"`
    // TemplateURL is the URL to the CloudFormation template to be used to create the
    // CloudFormation stack.
    TemplateUrl        *string                `json:"template_url"`
    // Token is used to identify and authenticate the CloudFormation stack creation.
    Token              *string                `json:"token"`
}

// RefreshWalletResponse represents a custom type struct for Success
type RefreshWalletResponse struct {
    // Embedded responses related to the resource.
    Embedded           interface{}            `json:"_embedded"`
    // URLs to pages related to the resource.
    Links              *WalletLinks           `json:"_links"`
    // AWS Account ID associated with the wallet.
    AccountNativeId    *string                `json:"account_native_id"`
    // Version of the template available
    AvailableVersion   *int64                 `json:"available_version"`
    // The AWS region associated with the wallet.
    AwsRegion          *string                `json:"aws_region"`
    // Clumio AWS Account ID.
    ClumioAwsAccountId *string                `json:"clumio_aws_account_id"`
    // DeploymentURL is an (external) link to an AWS console page for quick-creation
    // of the stack.
    DeploymentUrl      *string                `json:"deployment_url"`
    // ErrorCode is a short string describing the error, if any.
    ErrorCode          *string                `json:"error_code"`
    // ErrorMessage is a longer description explaining the error, if any, and how to
    // fix it.
    ErrorMessage       *string                `json:"error_message"`
    // The Clumio-assigned ID of the wallet.
    Id                 *string                `json:"id"`
    // The regions where the wallet is installed.
    InstalledRegions   []*string              `json:"installed_regions"`
    // TODO: Add struct field description
    KeyErrors          map[string]*ErrorModel `json:"key_errors"`
    // RoleArn is the AWS Resource Name of the IAM Role created by the stack.
    RoleArn            *string                `json:"role_arn"`
    // The version of the stack used or being used.
    StackVersion       *int64                 `json:"stack_version"`
    // State describes the state of the wallet. Valid states are:
    // Waiting: The wallet has been created, but a stack hasn't been created. The
    // wallet can't be used in this state.
    // Enabled: The wallet has been created and a stack has been created for the
    // wallet. This is the normal expected state of a wallet in use.
    // Error:   The wallet is inaccessible. See ErrorCode and ErrorMessage fields for
    // additional details.
    State              *string                `json:"state"`
    // The supported regions for the wallet.
    SupportedRegions   []*string              `json:"supported_regions"`
    // TemplateURL is the URL to the CloudFormation template to be used to create the
    // CloudFormation stack.
    TemplateUrl        *string                `json:"template_url"`
    // Token is used to identify and authenticate the CloudFormation stack creation.
    Token              *string                `json:"token"`
}

// RestoreDynamoDBTableResponse represents a custom type struct for Success
type RestoreDynamoDBTableResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *RestoreDynamoDBTableLinks    `json:"_links"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                       `json:"task_id"`
}

// RestoreEBSResponse represents a custom type struct for Success
type RestoreEBSResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *RestoreEBSLinks              `json:"_links"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                       `json:"task_id"`
}

// RestoreEBSResponseV1 represents a custom type struct for Success
type RestoreEBSResponseV1 struct {
    // HateoasCommonLinks are the common fields for HATEOAS response.
    Links *HateoasCommonLinks `json:"_links"`
}

// RestoreEC2Response represents a custom type struct for Success
type RestoreEC2Response struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *RestoreEC2Links              `json:"_links"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                       `json:"task_id"`
}

// RestoreFileResponse represents a custom type struct for Success
type RestoreFileResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *RestoreFileLinks             `json:"_links"`
    // The Clumio-assigned ID of the restored file.
    Id       *string                       `json:"id"`
    // The passcode that the end-user must use to access the restored
    // file, in the case the restored file was emailed to the end-user as part
    // of transparent data access.
    Passcode *string                       `json:"passcode"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                       `json:"task_id"`
}

// RestoreObjectsResponse represents a custom type struct for Success
type RestoreObjectsResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *RestoreObjectsLinks          `json:"_links"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                       `json:"task_id"`
}

// RestoreProtectionGroupResponse represents a custom type struct for Success
type RestoreProtectionGroupResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *RestoreProtectionGroupLinks  `json:"_links"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                       `json:"task_id"`
}

// RestoreProtectionGroupS3AssetResponse represents a custom type struct for Success
type RestoreProtectionGroupS3AssetResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded       `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *RestoreProtectionGroupS3AssetLinks `json:"_links"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                             `json:"task_id"`
}

// RestoreRecordPreviewResponse represents a custom type struct.
// Preview Success
type RestoreRecordPreviewResponse struct {
    // HateoasCommonLinks are the common fields for HATEOAS response.
    Links         *HateoasCommonLinks           `json:"_links"`
    // The preview of the query result, if `preview:true` in the request.
    // If preview was not set to true in the request, then the result of the query will be
    // available for download asynchronously.
    PreviewResult *RDSLogicalPreviewQueryResult `json:"preview_result"`
}

// RestoreRecordResponse represents a custom type struct for Success
type RestoreRecordResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded     `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *CreateRestoreRecordResponseLinks `json:"_links"`
    // The Clumio-assigned ID of the task generated by this request.
    // The requested records will be available for asynchronous download when the Task completes.
    // Use the [GET /restores/aws/rds-resources/records](#operation/list-rds-restored-records)
    // endpoint to list the records available for download. 
    // If `"preview":true` was set in the request, then a preview of the result
    // will be given instead of a Task ID.
    TaskId   *string                           `json:"task_id"`
}

// RestoreRecordsResponseAsync represents a custom type struct.
// Records restore request accepted
type RestoreRecordsResponseAsync struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *RestoreRecordsLinksAsync     `json:"_links"`
    // passcode that the end-user must use to access the restored
    // file, in case when restored file was emailed to the end-user as part
    // of transparent data access.
    Passcode *string                       `json:"passcode"`
    // The Clumio-assigned ID of the task generated by this request.
    // The requested records will be available for asynchronous download when the Task completes.
    // Use the [GET /restores/aws/dynamodb-tables/records](#operation/list-restored-records-aws-dynamodb-table)
    // endpoint to list the records available for download. 
    // If `"preview":true` was set in the request, then a preview of the result
    // will be given instead of a Task ID, and this field will be omitted.
    TaskId   *string                       `json:"task_id"`
}

// RestoreRecordsResponseSync represents a custom type struct.
// Records preview success
type RestoreRecordsResponseSync struct {
    // URLs to pages related to the resource.
    Links         *RestoreRecordsLinksSync    `json:"_links"`
    // If preview was not set to true in the request, then the result of the query will be
    // available for download asynchronously and this field has a value of `null`.
    PreviewResult *DynamoDBQueryPreviewResult `json:"preview_result"`
}

// RestoreS3BucketResponse represents a custom type struct for Success
type RestoreS3BucketResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasOuterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *RestoreS3BucketLinks         `json:"_links"`
    // The Clumio-assigned ID of the task created by this restore request.
    // The progress of the task can be monitored using the
    // `GET /tasks/{task_id}` endpoint.
    TaskId   *string                       `json:"task_id"`
}

// RestoredFilesResponse represents a custom type struct for Success
type RestoredFilesResponse struct {
    // Embedded responses related to the resource.
    Embedded        *RestoredFilesListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *RestoredFilesListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                     `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                     `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                    `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                     `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                     `json:"total_pages_count"`
}

// SendComplianceRunEmailResponse represents a custom type struct for Success
type SendComplianceRunEmailResponse struct {
    // Embedded responses related to the resource.
    Embedded interface{}                `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *ComplianceRunHateoasLinks `json:"_links"`
}

// SetAssignmentsResponse represents a custom type struct for Success
type SetAssignmentsResponse struct {
    // URLs to pages related to the resource.
    Links  *SetAssignmentsResponseLinks `json:"_links"`
    // The Clumio-assigned ID of the task generated by this request.
    TaskId *string                      `json:"task_id"`
}

// SetBucketPropertiesResponse represents a custom type struct.
// Accepted
type SetBucketPropertiesResponse struct {
    // URLs to pages related to the resource.
    Links *SetBucketPropertiesResponseLinks `json:"_links"`
}

// ShareFileRestoreEmailResponse represents a custom type struct for Success
type ShareFileRestoreEmailResponse struct {
    // URLs to pages related to the resource.
    Links *ShareFileRestoreEmailLinks `json:"_links"`
}

// UpdateAWSConnectionResponse represents a custom type struct for Success
type UpdateAWSConnectionResponse struct {
    // URLs to pages related to the resource.
    Links                      *AWSConnectionLinks      `json:"_links"`
    // The alias given to the account on AWS.
    AccountName                *string                  `json:"account_name"`
    // The AWS-assigned ID of the account associated with the connection.
    AccountNativeId            *string                  `json:"account_native_id"`
    // The AWS region associated with the connection. For example, `us-east-1`.
    AwsRegion                  *string                  `json:"aws_region"`
    // AWS account ID of the Clumio control plane.
    ClumioAwsAccountId         *string                  `json:"clumio_aws_account_id"`
    // AWS region of the Clumio control plane
    ClumioAwsRegion            *string                  `json:"clumio_aws_region"`
    // The consolidated configuration of the Clumio Cloud Protect and Clumio Cloud Discover products for this connection.
    // If this connection is deprecated to use unconsolidated configuration, then this field has a
    // value of `null`.
    Config                     *ConsolidatedConfig      `json:"config"`
    // Clumio assigned ID of the associated connection group.
    ConnectionGroupId          *string                  `json:"connection_group_id"`
    // Management status of connection.
    ConnectionManagementStatus *string                  `json:"connection_management_status"`
    // The status of the connection considering all the deployments made for it.
    ConnectionStatus           *string                  `json:"connection_status"`
    // The timestamp of when the connection was created.
    CreatedTimestamp           *string                  `json:"created_timestamp"`
    // AWS account ID of the data plane for the connection.
    DataPlaneAccountId         *string                  `json:"data_plane_account_id"`
    // The deployment method with which the currently active connection was established.
    DeploymentType             *string                  `json:"deployment_type"`
    // The user-provided description for this connection.
    Description                *string                  `json:"description"`
    // The configuration of the Clumio Discover product for this connection.
    // If this connection is not configured for Clumio Discover, then this field has a
    // value of `null`.
    Discover                   *DiscoverConfig          `json:"discover"`
    // Clumio assigned external ID of the connection or of the associated connection group.
    ExternalId                 *string                  `json:"external_id"`
    // The Clumio-assigned ID of the connection.
    Id                         *string                  `json:"id"`
    // Status denoting whether Ingestion has started for the connection.
    // Valid values are "initial", "in_progress", "failed", "completed".
    IngestionStatus            *string                  `json:"ingestion_status"`
    // K8S Namespace
    Namespace                  *string                  `json:"namespace"`
    // The Clumio-assigned ID of the organizational unit associated with the
    // AWS environment. If this parameter is not provided, then the value
    // defaults to the first organizational unit assigned to the requesting
    // user. For more information about organizational units, refer to the
    // Organizational-Units documentation.
    OrganizationalUnitId       *string                  `json:"organizational_unit_id"`
    // The configuration of the Clumio Cloud Protect product for this connection.
    // If this connection is not configured for Clumio Cloud Protect, then this field has a
    // value of `null`.
    Protect                    *ProtectConfig           `json:"protect"`
    // The asset types enabled for protect.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS", "IcebergOnGlue", "IcebergOnS3Tables"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    ProtectAssetTypesEnabled   []*string                `json:"protect_asset_types_enabled"`
    // TODO: Add struct field description
    Resources                  *ConnectionResourcesResp `json:"resources"`
    // The Amazon Resource Name of the stale CloudFormation stack when the connection was migrated to connection groups.
    // NOTE - This has to be removed from AWS as well to delete the connection completely.
    RetiredStackArn            *string                  `json:"retired_stack_arn"`
    // The services to be enabled for this configuration. Valid values are
    // ["discover"], ["discover", "protect"]. This is only set when the
    // registration is created, the enabled services are obtained directly from
    // the installed template after that. (Deprecated as all connections will have
    // both discover and protect enabled)
    ServicesEnabled            []*string                `json:"services_enabled"`
    // The Amazon Resource Name of the installed and active CloudFormation stack(if any) in AWS.
    StackArn                   *string                  `json:"stack_arn"`
    // The name given to the installed and active CloudFormation stack(if any) in AWS.
    StackName                  *string                  `json:"stack_name"`
    // Status denoting whether Target Setup has started for the connection.
    // Valid values are "initial", "in_progress", "failed", "completed".
    TargetSetupStatus          *string                  `json:"target_setup_status"`
    // TODO: Add struct field description
    TemplatePermissionSet      *string                  `json:"template_permission_set"`
    // The 36-character Clumio AWS integration ID token used to identify the
    // installation of the CloudFormation template on the account. This value
    // will be pasted into the ClumioToken field when creating the
    // CloudFormation stack.
    Token                      *string                  `json:"token"`
}

// UpdateAlertResponse represents a custom type struct for Success
type UpdateAlertResponse struct {
    // Embedded responses related to the resource.
    Embedded            *AlertEmbedded          `json:"_embedded"`
    // The ETag value.
    Etag                *string                 `json:"_etag"`
    // URLs to pages related to the resource.
    Links               *AlertLinks             `json:"_links"`
    // The issue that generated the alert. Each cause belongs to an alert type.
    Cause               *string                 `json:"cause"`
    // The timestamp of when the alert was cleared, either automatically by Clumio or manually by a Clumio user.
    // Represented in RFC-3339 format. If this alert has not been cleared, then this field has a value of `null`.
    ClearedTimestamp    *string                 `json:"cleared_timestamp"`
    // The Clumio-assigned ID of the consolidated alert associated with this individual alert. Alerts are consolidated based on matching parent entity, alert type, and alert cause.
    ConsolidatedAlertId *string                 `json:"consolidated_alert_id"`
    // Additional information about the alert.
    Details             *IndividualAlertDetails `json:"details"`
    // The Clumio-assigned ID of the individual alert.
    Id                  *string                 `json:"id"`
    // A record of user-provided information about the alert.
    Notes               *string                 `json:"notes"`
    // The parent object of the primary entity associated with or affected by the alert. For example, "aws_environment" is the parent entity of primary entity "aws_ebs_volume".
    ParentEntity        *AlertParentEntity      `json:"parent_entity"`
    // The primary object associated with or affected by the alert. Examples of primary entities include "aws_connection", "aws_ebs_volume".
    PrimaryEntity       *AlertPrimaryEntity     `json:"primary_entity"`
    // The number of times the alert has recurred for this primary entity.
    RaisedCount         *uint64                 `json:"raised_count"`
    // The timestamp of when the alert was raised. Represented in RFC-3339 format.
    RaisedTimestamp     *string                 `json:"raised_timestamp"`
    // The alert severity level. Values include "error" and "warning".
    Severity            *string                 `json:"severity"`
    // The individual alert status. An individual alert that is in "active" status is one that is still open and has yet to be addressed.
    // An individual alert that is in "cleared" status is one that has been cleared, either automatically by Clumio or manually by a Clumio user.
    Status              *string                 `json:"status"`
    // TODO: Add struct field description
    Tags                []*RestEntity           `json:"tags"`
    // The general alert category. Some alert types may be associated with multiple causes.
    // Refer to the Alert Type table for a complete list of alert types.
    ClumioType          *string                 `json:"type"`
    // The timestamp of when the alert was last updated. Represented in RFC-3339 format.
    // The alert is updated whenever there is a new occurrence of the same alert within the same entity.
    UpdatedTimestamp    *string                 `json:"updated_timestamp"`
}

// UpdateAutoUserProvisioningRuleResponse represents a custom type struct for Success
type UpdateAutoUserProvisioningRuleResponse struct {
    // Embedded responses related to the resource.
    Embedded  *AutoUserProvisioningRuleEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links     *AutoUserProvisioningRuleLinks    `json:"_links"`
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
    Condition *string                           `json:"condition"`
    // Unique name assigned to the rule.
    Name      *string                           `json:"name"`
    // Specifies the role and the organizational units to be assigned to the user subject to the rule criteria.
    Provision *RuleProvision                    `json:"provision"`
    // The Clumio-assigned ID of the rule.
    RuleId    *string                           `json:"rule_id"`
}

// UpdateAutoUserProvisioningSettingResponse represents a custom type struct for Success
type UpdateAutoUserProvisioningSettingResponse struct {
    // URLs to pages related to the resource.
    Links     *AutoUserProvisioningSettingLinks `json:"_links"`
    // Whether auto user provisioning is enabled or not.
    IsEnabled *bool                             `json:"is_enabled"`
}

// UpdateComplianceConfigurationResponse represents a custom type struct for Success
type UpdateComplianceConfigurationResponse struct {
    // If the `embed` query parameter is set, displays the responses of the related resource,
    // as defined by the embeddable link specified.
    Embedded     interface{}                   `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *ComplianceConfigurationLinks `json:"_links"`
    // The RFC3339 format time when the report configuration was created.
    Created      *string                       `json:"created"`
    // The user-provided description of the compliance report configuration.
    Description  *string                       `json:"description"`
    // The unique identifier of the report configuration.
    Id           *string                       `json:"id"`
    // Most recent report run generated from the report configuration.
    LatestRun    *LatestRun                    `json:"latest_run"`
    // The user-provided name of the compliance report configuration.
    Name         *string                       `json:"name"`
    // Notification channels to send the generated report runs.
    Notification *NotificationSetting          `json:"notification"`
    // Filter and control parameters of compliance report.
    Parameter    *Parameter                    `json:"parameter"`
    // When the report will be generated and sent. If the schedule is not provided then a
    // default value will be used.
    Schedule     *ScheduleSetting              `json:"schedule"`
}

// UpdateConnectionGroupResponse represents a custom type struct for Success
type UpdateConnectionGroupResponse struct {
    // Embedded responses related to the resource.
    Embedded                 interface{}           `json:"_embedded"`
    // The ETag value.
    Etag                     *string               `json:"_etag"`
    // URLs to pages related to the resource.
    Links                    *ConnectionGroupLinks `json:"_links"`
    // The alias given to the associated account in AWS.
    AccountName              *string               `json:"account_name"`
    // The AWS-assigned IDs of the accounts associated with the Connection Group.
    AccountNativeIds         []*string             `json:"account_native_ids"`
    // List of asset types connected via the connection-group.
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS", "IcebergOnGlue", "IcebergOnS3Tables"].
    // 
    // NOTE -
    // 1. EC2/EBS is required for EC2MSSQL.
    // 2. EBS as a value is deprecated in favor of EC2/EBS.
    AssetTypesEnabled        []*string             `json:"asset_types_enabled"`
    // The AWS regions associated with the with the Connection Group.
    AwsRegions               []*string             `json:"aws_regions"`
    // The consolidated configuration of the Clumio Cloud Protect and Clumio Cloud Discover products for this connection.
    // If this connection is deprecated to use unconsolidated configuration, then this field has a
    // value of `null`.
    Config                   *ConsolidatedConfig   `json:"config"`
    // The timestamp of when the connection was created.
    CreatedTimestamp         *string               `json:"created_timestamp"`
    // Clumio's S3 URL that contains the template to create the required resources in the
    // given account(s) according to the request.
    DeploymentTemplateUrl    *string               `json:"deployment_template_url"`
    // User-provided description for this connection group.
    Description              *string               `json:"description"`
    // Clumio assigned external ID for the connection group, should be used while creating the AWS stack.
    ExternalId               *string               `json:"external_id"`
    // The Clumio-assigned ID of the Connection Group, should be used as the token while creating the stack in AWS.
    Id                       *string               `json:"id"`
    // The AWS Account IDs that are intended to be associated with the Connection Group.
    IntendedAccountNativeIds []*string             `json:"intended_account_native_ids"`
    // THe asset types that are intended to be connected via connection-group.
    IntendedAssetTypes       []*string             `json:"intended_asset_types"`
    // The AWS regions that are intended to be connected with the Connection Group.
    IntendedAwsRegions       []*string             `json:"intended_aws_regions"`
    // The master account which manages the connection-group's stack.
    MasterAwsAccountId       *string               `json:"master_aws_account_id"`
    // The master region which manages the connection-group's stack.
    MasterRegion             *string               `json:"master_region"`
    // Ongoing Operation of the deployed and active stack of ConnectionGroup.
    OngoingStackOperation    *string               `json:"ongoing_stack_operation"`
    // The Clumio-assigned ID of the organizational unit associated with the
    // AWS environment. If this parameter is not provided, then the value
    // defaults to the first organizational unit assigned to the requesting
    // user. For more information about organizational units, refer to the
    // Organizational-Units documentation.
    OrganizationalUnitId     *string               `json:"organizational_unit_id"`
    // The Amazon Resource Name of the installed CloudFormation stack in AWS.
    StackArn                 *string               `json:"stack_arn"`
    // The name given to the installed CloudFormation stack in AWS.
    StackName                *string               `json:"stack_name"`
    // The status of the Connection Group based on the stack in associated AWS account.
    Status                   *string               `json:"status"`
}

// UpdateConsolidatedAlertResponse represents a custom type struct for Success
type UpdateConsolidatedAlertResponse struct {
    // The ETag value.
    Etag               *string                        `json:"_etag"`
    // URLs to pages related to the resource.
    Links              *ConsolidatedAlertLinks        `json:"_links"`
    // The number of currently active individual alerts associated with the consolidated alert.
    ActiveEntityCount  *int64                         `json:"active_entity_count"`
    // The issue that generated the alert. Each alert cause is associated with an alert type.
    Cause              *string                        `json:"cause"`
    // The number of cleared individual alerts associated with the consolidated alert.
    ClearedEntityCount *int64                         `json:"cleared_entity_count"`
    // The timestamp of when the consolidated alert was cleared, if ever. Represented in RFC-3339 format. If this alert has not been cleared, this field will have a value of `null`.
    // A consolidated alert goes into "cleared" status when all of its associated individual alerts are in "cleared" status or when a Clumio user manually clears it.
    ClearedTimestamp   *string                        `json:"cleared_timestamp"`
    // Additional information about the consolidated alert.
    Details            *ConsolidatedAlertDetails      `json:"details"`
    // The Clumio-assigned ID of the consolidated alert.
    Id                 *string                        `json:"id"`
    // A record of user-provided information about the alert.
    Notes              *string                        `json:"notes"`
    // The entity associated with or affected by the alert.
    ParentEntity       *ConsolidatedAlertParentEntity `json:"parent_entity"`
    // The timestamp of when the consolidated alert was initially raised. Represented in RFC-3339 format.
    RaisedTimestamp    *string                        `json:"raised_timestamp"`
    // The alert severity level. Values include "error" and "warning".
    Severity           *string                        `json:"severity"`
    // The consolidated alert status. A consolidated alert is in "active" status if one or more of its associated individual alerts is in "active" status.
    // A consolidated alert goes into "cleared" status when all of its associated individual alerts are in "cleared" status or when a Clumio user manually clears it.
    Status             *string                        `json:"status"`
    // The general alert category. An alert type may be associated with multiple alert causes. Examples of alert types include "tag_conflict" and "policy_violated".
    // Refer to the Alert Type table for a complete list of alert types.
    ClumioType         *string                        `json:"type"`
    // The timestamp of when the consolidated alert was last updated. Represented in RFC-3339 format.
    // Raising a new individual alert will update its associated consolidated alert.
    UpdatedTimestamp   *string                        `json:"updated_timestamp"`
}

// UpdateManagementGroupResponse represents a custom type struct for Success
type UpdateManagementGroupResponse struct {
    // URLs to pages related to the resource.
    Links      *ManagementGroupLinks `json:"_links"`
    // The Clumio-assigned ID of the management group.
    Id         *string               `json:"id"`
    // The name of the management group.
    Name       *string               `json:"name"`
    // The type of the management group. Possible values include `on_prem`.
    ClumioType *string               `json:"type"`
    // The Clumio-assigned ID of the vCenter server associated with the management group.
    // All management groups are associated with a vCenter server.
    VcenterId  *string               `json:"vcenter_id"`
}

// UpdatePolicyResponse represents a custom type struct for Success
type UpdatePolicyResponse struct {
    // If the `embed` query parameter is set, displays the responses of the related resource,
    // as defined by the embeddable link specified.
    Embedded             *PolicyEmbedded            `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                *UpdatePolicyResponseLinks `json:"_links"`
    // The status of the policy.
    // Refer to the Policy Activation Status table for a complete list of policy statuses.
    ActivationStatus     *string                    `json:"activation_status"`
    // The created time of the policy in unix time.
    CreatedTime          *int64                     `json:"created_time"`
    // The Clumio-assigned ID of the policy.
    Id                   *string                    `json:"id"`
    // The following table describes the possible lock statuses of a policy.
    // 
    // +----------+-------------------------------------------------------------------+
    // |  Status  |                            Description                            |
    // +==========+===================================================================+
    // | unlocked | Policies are unlocked until an update or deletion task is queued. |
    // +----------+-------------------------------------------------------------------+
    // | updating | During a policy edit, concurrent edits or deletion requests will  |
    // |          | be rejected.                                                      |
    // +----------+-------------------------------------------------------------------+
    // | deleting | During policy deletion, concurrent edits or deletion requests     |
    // |          | will be rejected.                                                 |
    // +----------+-------------------------------------------------------------------+
    // 
    LockStatus           *string                    `json:"lock_status"`
    // The user-provided name of the policy.
    Name                 *string                    `json:"name"`
    // TODO: Add struct field description
    Operations           []*PolicyOperation         `json:"operations"`
    // The Clumio-assigned ID of the organizational unit associated with the policy.
    OrganizationalUnitId *string                    `json:"organizational_unit_id"`
    // The Clumio-assigned ID of the task generated by this request.
    TaskId               *string                    `json:"task_id"`
    // The policy-level timezone is deprecated, as the operation-level timezone should be used instead.
    // The timezone must be a valid location name from the IANA Time Zone database.
    // For instance, "America/New_York", "US/Central", "UTC".
    // Deprecated: true
    Timezone             *string                    `json:"timezone"`
    // The updated time of the policy in unix time.
    UpdatedTime          *int64                     `json:"updated_time"`
}

// UpdateProtectionGroupResponse represents a custom type struct for Success
type UpdateProtectionGroupResponse struct {
    // Embedded responses related to the resource.
    Embedded                       interface{}                  `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                          *ProtectionGroupVersionLinks `json:"_links"`
    // The backup target AWS region associated with the protection group, empty if
    // in-region or not configured.
    BackupTargetAwsRegion          *string                      `json:"backup_target_aws_region"`
    // BackupTierStat
    BackupTierStats                []*BackupTierStat            `json:"backup_tier_stats"`
    // Number of buckets
    BucketCount                    *int64                       `json:"bucket_count"`
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
    BucketRule                     *string                      `json:"bucket_rule"`
    // Creation time of the protection group in RFC-3339 format.
    CreatedTimestamp               *string                      `json:"created_timestamp"`
    // The user-assigned description of the protection group.
    Description                    *string                      `json:"description"`
    // The Clumio-assigned ID of the protection group.
    Id                             *string                      `json:"id"`
    // Whether the protection group already has a backup target configured by a policy, or
    // is open to be protected by an in-region or out-of-region S3 policy.
    IsBackupTargetRegionConfigured *bool                        `json:"is_backup_target_region_configured"`
    // Determines whether the protection group is active or has been deleted. Deleted protection
    // groups may be purged after some time once there are no active backups associated with it.
    IsDeleted                      *bool                        `json:"is_deleted"`
    // Time of the last backup in RFC-3339 format.
    LastBackupTimestamp            *string                      `json:"last_backup_timestamp"`
    // Time of the last successful continuous backup in RFC-3339 format.
    LastContinuousBackupTimestamp  *string                      `json:"last_continuous_backup_timestamp"`
    // Modified time of the protection group in RFC-3339 format.
    ModifiedTimestamp              *string                      `json:"modified_timestamp"`
    // The user-assigned name of the protection group.
    Name                           *string                      `json:"name"`
    // ObjectFilter
    // defines which objects will be backed up.
    ObjectFilter                   *ObjectFilter                `json:"object_filter"`
    // The Clumio-assigned ID of the organizational unit associated with the Protection Group.
    OrganizationalUnitId           *string                      `json:"organizational_unit_id"`
    // Cumulative count of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpObjectCount       *int64                       `json:"total_backed_up_object_count"`
    // Cumulative size of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpSizeBytes         *int64                       `json:"total_backed_up_size_bytes"`
    // Version of the protection group. The version number is incremented every time
    // a change is made to the protection group.
    Version                        *int64                       `json:"version"`
}

// UpdateRuleResponse represents a custom type struct for Success
type UpdateRuleResponse struct {
    // URLs to pages related to the resource.
    Links  *UpdateRuleResponseLinks `json:"_links"`
    // A rule applies an action subject to a condition criteria.
    Rule   *Rule                    `json:"rule"`
    // The Clumio-assigned ID of the task generated by this request.
    TaskId *string                  `json:"task_id"`
}

// UpdateS3InstantAccessEndpointResponse represents a custom type struct for Success
type UpdateS3InstantAccessEndpointResponse struct {
    // Embedded responses related to the resource.
    Embedded                 *S3InstantAccessEndpointEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                    *S3InstantAccessEndpointLinks    `json:"_links"`
    // The AWS-assigned ID of the account associated with the S3 instant access endpoint.
    AwsAccountId             *string                          `json:"aws_account_id"`
    // The AWS region of the S3 instant access endpoint and its source backup.
    BackupRegion             *string                          `json:"backup_region"`
    // The name of source bucket.
    BucketName               *string                          `json:"bucket_name"`
    // The time that this endpoint was created, in RFC-3339 format.
    CreatedTimestamp         *string                          `json:"created_timestamp"`
    // The status of the S3 instant access endpoint. Possible values include "preparing",
    // "active", "expiring", and "expired".
    EndpointStatus           *string                          `json:"endpoint_status"`
    // The time that this endpoint expires, in RFC-3339 format.
    ExpiryTimestamp          *string                          `json:"expiry_timestamp"`
    // The Clumio-assigned ID of the S3 instant access endpoint.
    Id                       *string                          `json:"id"`
    // The user-assigned name of the S3 instant access endpoint.
    Name                     *string                          `json:"name"`
    // The time in RFC-3339 format that the restored objects are backed up from.
    ObjectsCreatedAfter      *string                          `json:"objects_created_after"`
    // The time in RFC-3339 format that the restored objects are backed up to.
    ObjectsCreatedBefore     *string                          `json:"objects_created_before"`
    // The Clumio-assigned ID of the organizational unit associated with the S3 instant access endpoint.
    OrganizationalUnitId     *string                          `json:"organizational_unit_id"`
    // The Clumio-assigned ID of the protection group this endpoint is created for.
    ProtectionGroupId        *string                          `json:"protection_group_id"`
    // The user-assigned name of the protection group this endpoints is created for.
    ProtectionGroupName      *string                          `json:"protection_group_name"`
    // The Clumio-assigned ID of the bucket protection group.
    ProtectionGroupS3AssetId *string                          `json:"protection_group_s3_asset_id"`
    // The AWS region of the source bucket.
    Region                   *string                          `json:"region"`
    // The time at which the backup was restored from this endpoint in RFC-3339 format.
    // Deprecated.
    RestoreTimestamp         *string                          `json:"restore_timestamp"`
    // The time that this endpoint was last updated, in RFC-3339 format.
    UpdatedTimestamp         *string                          `json:"updated_timestamp"`
}

// UpdateS3InstantAccessEndpointRoleResponse represents a custom type struct for Success
type UpdateS3InstantAccessEndpointRoleResponse struct {
    // Embedded responses related to the resource.
    Embedded *S3InstantAccessEndpointEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *S3InstantAccessEndpointLinks    `json:"_links"`
}

// UpdateTaskResponse represents a custom type struct for Success
type UpdateTaskResponse struct {
    // The ETag value.
    Etag               *string            `json:"_etag"`
    // URLs to pages related to the resource.
    Links              *TaskLinks         `json:"_links"`
    // The task category. Examples of task types include "backup", "restore", "snapshot", and "system".
    // 
    // +-------------------+----------------------------------------------------------+
    // |     Category      |                       Description                        |
    // +===================+==========================================================+
    // | backup            | Encompasses all modes of backups. This does not include  |
    // |                   | in-account snapshots.                                    |
    // +-------------------+----------------------------------------------------------+
    // | restore           | Encompasses all modes of restores. This does not include |
    // |                   | restores of in-account snapshots.                        |
    // +-------------------+----------------------------------------------------------+
    // | snapshot          | Encompasses all modes of in-account snapshots.           |
    // +-------------------+----------------------------------------------------------+
    // | snapshot_restore  | Encompasses all modes of snapshot restores.              |
    // +-------------------+----------------------------------------------------------+
    // | system            | Encompasses a variety of system-initiated tasks, such as |
    // |                   | aws_rds_backup_target_setup and                          |
    // |                   | aws_ec2_instance_backup_indexing.                        |
    // +-------------------+----------------------------------------------------------+
    // | report_generation | Encompasses task types which generate reports, such as   |
    // |                   | activity_report_file_download.                           |
    // +-------------------+----------------------------------------------------------+
    // | management        | Encompasses user-initiated tasks which manage Clumio     |
    // |                   | resources, such as organizational_unit_update and        |
    // |                   | policy_update.                                           |
    // +-------------------+----------------------------------------------------------+
    // 
    Category           *string            `json:"category"`
    // The timestamp of when the task was created. Represented in RFC-3339 format.
    CreatedTimestamp   *string            `json:"created_timestamp"`
    // The timestamp of when the task ended. If this task has not yet ended,
    // then this field has a value of `null`. Represented in RFC-3339 format.
    EndTimestamp       *string            `json:"end_timestamp"`
    // The task genre. A genre is a high-level collection of task categories.
    // 
    // +----------------+-------------------------------------------------------------+
    // |     Genre      |                         Description                         |
    // +================+=============================================================+
    // | operational    | Encompasses all backup, restore, snapshot, and              |
    // |                | snapshot_restore tasks.                                     |
    // +----------------+-------------------------------------------------------------+
    // | administrative | Encompasses management, system, and report_generation       |
    // |                | tasks.                                                      |
    // +----------------+-------------------------------------------------------------+
    // 
    Genre              *string            `json:"genre"`
    // The Clumio-assigned ID of the task.
    Id                 *string            `json:"id"`
    // Determines whether or not this task can be aborted.
    // A task can be aborted if its status is either "queued" or "in_progress".
    // Tasks of certain types including "aws_ebs_volume_backup_indexing" cannot be aborted.
    IsAbortable        *bool              `json:"is_abortable"`
    // The parent entity associated with the task.
    ParentEntity       *TaskParentEntity  `json:"parent_entity"`
    // The primary entity associated with the task.
    PrimaryEntity      *TaskPrimaryEntity `json:"primary_entity"`
    // The percentage progress of task completion. Measured as an integer value between 0 and 100.
    ProgressPercentage *int64             `json:"progress_percentage"`
    // The timestamp of when the task started. If this task has not started yet,
    // then this field has a value of `null`. Represented in RFC-3339 format.
    StartTimestamp     *string            `json:"start_timestamp"`
    // The task status. Examples of task statuses include, "queued", "in_progress", and "completed".
    // Refer to the Task Status table for a complete list of task statuses.
    Status             *string            `json:"status"`
    // Refer to the Task Type table for a complete list of task types.
    ClumioType         *string            `json:"type"`
}

// UpdateUserResponse represents a custom type struct for Success
type UpdateUserResponse struct {
    // Embedded responses related to the resource.
    Embedded                   *UserEmbedded                 `json:"_embedded"`
    // ETag value
    Etag                       *string                       `json:"_etag"`
    // URLs to pages related to the resource.
    Links                      *UserLinks                    `json:"_links"`
    // The organizational units assigned to the user, with the specified role.
    AccessControlConfiguration []*RoleForOrganizationalUnits `json:"access_control_configuration"`
    // The email address of the Clumio user.
    Email                      *string                       `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                   *string                       `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                         *string                       `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                    *string                       `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                *bool                         `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                  *bool                         `json:"is_enabled"`
    // The timestamp of when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp      *string                       `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount    *int64                        `json:"organizational_unit_count"`
}

// UpdateUserResponseV1 represents a custom type struct for Success
type UpdateUserResponseV1 struct {
    // Embedded responses related to the resource.
    Embedded                      *UserEmbeddedV1 `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                         *UserLinks      `json:"_links"`
    // The list of organizational unit IDs assigned to the user.
    // This attribute will be available when reading a single user and not when listing users.
    AssignedOrganizationalUnitIds []*string       `json:"assigned_organizational_unit_ids"`
    // Assigned Role for the user.
    AssignedRole                  *string         `json:"assigned_role"`
    // The email address of the Clumio user.
    Email                         *string         `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                      *string         `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                            *string         `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                       *string         `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                   *bool           `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                     *bool           `json:"is_enabled"`
    // The timestamp of when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp         *string         `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount       *int64          `json:"organizational_unit_count"`
}
