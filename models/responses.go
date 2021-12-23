// Copyright (c) 2021 Clumio All Rights Reserved

// Package models has the structs representing responses
package models

// CreateAWSConnectionResponse represents a custom type struct for Success
type CreateAWSConnectionResponse struct {
    // URLs to pages related to the resource.
    Links                    *AWSConnectionLinks `json:"_links"`
    // The alias given to the account on AWS.
    AccountName              *string             `json:"account_name"`
    // The AWS-assigned ID of the account associated with the connection.
    AccountNativeId          *string             `json:"account_native_id"`
    // The AWS region associated with the connection. For example, `us-east-1`.
    AwsRegion                *string             `json:"aws_region"`
    // Clumio AWS AccountId
    ClumioAwsAccountId       *string             `json:"clumio_aws_account_id"`
    // Clumio AWS Region
    ClumioAwsRegion          *string             `json:"clumio_aws_region"`
    // The status of the connection. Possible values include "connecting",
    // "connected", and "unlinked".
    ConnectionStatus         *string             `json:"connection_status"`
    // The timestamp of when the connection was created.
    CreatedTimestamp         *string             `json:"created_timestamp"`
    // An optional, user-provided description for this connection.
    Description              *string             `json:"description"`
    // The configuration of the Clumio Discover product for this connection.
    // If this connection is not configured for Clumio Discover, then this field has a
    // value of `null`.
    Discover                 *DiscoverConfig     `json:"discover"`
    // The Clumio-assigned ID of the connection.
    Id                       *string             `json:"id"`
    // Lambda Bucket Prefix required for configuring Clumio Protect Warm-tier AWS Lambda.
    LambdaBucketPrefix       *string             `json:"lambda_bucket_prefix"`
    // K8S Namespace
    Namespace                *string             `json:"namespace"`
    // The Clumio-assigned ID of the organizational unit associated with the
    // AWS environment. If this parameter is not provided, then the value
    // defaults to the first organizational unit assigned to the requesting
    // user. For more information about organizational units, refer to the
    // Organizational-Units documentation.
    OrganizationalUnitId     *string             `json:"organizational_unit_id"`
    // The configuration of the Clumio Cloud Protect product for this connection.
    // If this connection is not configured for Clumio Cloud Protect, then this field has a
    // value of `null`.
    Protect                  *ProtectConfig      `json:"protect"`
    // The asset types enabled for protect. This is only populated if "protect"
    // is enabled. Valid values are any of ["EBS", "RDS", "DynamoDB", "EC2MSSQL"].
    // EBS and RDS are mandatory datasources.
    ProtectAssetTypesEnabled []*string           `json:"protect_asset_types_enabled"`
    // The services to be enabled for this configuration. Valid values are
    // ["discover"], ["discover", "protect"]. This is only set when the
    // registration is created, the enabled services are obtained directly from
    // the installed template after that.
    ServicesEnabled          []*string           `json:"services_enabled"`
    // The Amazon Resource Name of the installed CloudFormation stack in this AWS account
    StackArn                 *string             `json:"stack_arn"`
    // The name given to the installed CloudFormation stack on AWS.
    StackName                *string             `json:"stack_name"`
    // Subnets Lambda Object Key required for configuring Clumio Protect Warm-tier AWS Lambda.
    SubnetsLambdaObjectKey   *string             `json:"subnets_lambda_object_key"`
    // The 36-character Clumio AWS integration ID token used to identify the
    // installation of the CloudFormation template on the account. This value
    // will be pasted into the ClumioToken field when creating the
    // CloudFormation stack.
    Token                    *string             `json:"token"`
}

// CreateAWSTemplateV2Response represents a custom type struct for Success
type CreateAWSTemplateV2Response struct {
    // The latest available URL for the template.
    CloudformationUrl *string                `json:"cloudformation_url"`
    // TODO: Add struct field description
    Config            *TemplateConfiguration `json:"config"`
    // The latest available URL for the terraform template.
    TerraformUrl      *string                `json:"terraform_url"`
}

// CreateHcmHostResponse represents a custom type struct for Success
type CreateHcmHostResponse struct {
    // TODO: Add struct field description
    Hosts []*Host `json:"hosts"`
}

// CreateHostECCredentialsResponse represents a custom type struct for Success
type CreateHostECCredentialsResponse struct {
    // URLs to pages related to the resource.
    Links                    *HostLinks `json:"_links"`
    // The edge connector credentials for the host. This token is required during the installation of the MSI.
    EdgeConnectorCredentials *string    `json:"edge_connector_credentials"`
    // The user-provided endpoint used to connect the host.
    Endpoint                 *string    `json:"endpoint"`
    // The Clumio-assigned ID of the management group associated with the host.
    GroupId                  *string    `json:"group_id"`
    // The Clumio-assigned ID of the Host.
    Id                       *string    `json:"id"`
    // The timestamp of the last successful heartbeat of this host. Represented in RFC-3339 format.
    LastHeartbeatTimestamp   *string    `json:"last_heartbeat_timestamp"`
    // Name of the Host.
    Name                     *string    `json:"name"`
    // The connection status of the Host. Possible values include `connected`, `disconnected`, `connection_pending`, and `invalid_token`.
    Status                   *string    `json:"status"`
    // The Clumio-assigned ID of the management subgroup associated with the host.
    SubgroupId               *string    `json:"subgroup_id"`
    // The Clumio-assigned UUID of the host. This UUID is used for filtering hosts during list operations.
    Uuid                     *string    `json:"uuid"`
}

// CreateMssqlDatabaseRestoreResponse represents a custom type struct for Success
type CreateMssqlDatabaseRestoreResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasLinks `json:"_embedded"`
    // Embedded responses related to the resource.
    Links    *ReadTaskHateoasLinks `json:"_links"`
}

// CreateOrganizationalUnitResponse represents a custom type struct for Success
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
    Embedded                      *PolicyEmbedded    `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                         *PolicyLinks       `json:"_links"`
    // The status of the policy.
    // Refer to the Policy Activation Status table for a complete list of policy statuses.
    ActivationStatus              *string            `json:"activation_status"`
    // The Clumio-assigned IDs of the organizational units to whom the policy has been assigned.
    AssignedOrganizationalUnitIds []*string          `json:"assigned_organizational_unit_ids"`
    // The Clumio-assigned ID of the policy.
    Id                            *string            `json:"id"`
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
    LockStatus                    *string            `json:"lock_status"`
    // The user-provided name of the policy.
    Name                          *string            `json:"name"`
    // The SLAs of an individual operation.
    Operations                    []*PolicyOperation `json:"operations"`
    // The Clumio-assigned ID of the organizational unit associated with the policy.
    OrganizationalUnitId          *string            `json:"organizational_unit_id"`
}

// CreateReportDownloadResponse represents a custom type struct for Success
type CreateReportDownloadResponse struct {
    // The Clumio-assigned ID of the task created by the request.
    // The progress of the task can be monitored using the
    // [`GET /tasks/{task_id}`](#operation/list-tasks) endpoint.
    TaskId *string `json:"task_id"`
}

// CreateUserResponse represents a custom type struct for Success
type CreateUserResponse struct {
    // Embedded responses related to the resource.
    Embedded                      *UserEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                         *UserLinks    `json:"_links"`
    // The list of organizational unit IDs assigned to the user.
    // This attribute will be available when reading a single user and not when listing users.
    AssignedOrganizationalUnitIds []*string     `json:"assigned_organizational_unit_ids"`
    // Assigned Role for the user.
    AssignedRole                  *string       `json:"assigned_role"`
    // The email address of the Clumio user.
    Email                         *string       `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                      *string       `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                            *string       `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                       *string       `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                   *bool         `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                     *bool         `json:"is_enabled"`
    // The timestamp of when when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp         *string       `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount       *int64        `json:"organizational_unit_count"`
}

// DeleteHcmHostResponse represents a custom type struct for Success
type DeleteHcmHostResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasLinks `json:"_embedded"`
    // Embedded responses related to the resource.
    Links    *ReadTaskHateoasLinks `json:"_links"`
    // TaskID for DeleteHostsReq
    TaskId   *string               `json:"task_id"`
}

// DeleteOrganizationalUnitResponse represents a custom type struct.
// Accepted
type DeleteOrganizationalUnitResponse struct {
    // Embedded responses related to the resource.
    Embedded *EntityGroupEmbedded  `json:"_embedded"`
    // Embedded responses related to the resource.
    Links    *ReadTaskHateoasLinks `json:"_links"`
}

// EditProfileResponse represents a custom type struct for Success
type EditProfileResponse struct {
    // Embedded responses related to the resource.
    Embedded                      *UserEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                         *UserLinks    `json:"_links"`
    // The list of organizational unit IDs assigned to the user.
    // This attribute will be available when reading a single user and not when listing users.
    AssignedOrganizationalUnitIds []*string     `json:"assigned_organizational_unit_ids"`
    // Assigned Role for the user.
    AssignedRole                  *string       `json:"assigned_role"`
    // The email address of the Clumio user.
    Email                         *string       `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                      *string       `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                            *string       `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                       *string       `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                   *bool         `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                     *bool         `json:"is_enabled"`
    // The timestamp of when when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp         *string       `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount       *int64        `json:"organizational_unit_count"`
}

// Error represents a custom type struct.
// Error
type Error struct {
    // TODO: Add struct field description
    Errors []*SingleErrorResponse `json:"errors"`
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

// ListComputeResourcesResponse represents a custom type struct for Success
type ListComputeResourcesResponse struct {
    // TODO: Add struct field description
    Computeresourcefolders []*VCenterFolder          `json:"computeResourceFolders"`
    // TODO: Add struct field description
    Computeresources       []*VCenterComputeResource `json:"computeResources"`
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

// ListDatacentersResponse represents a custom type struct for Success
type ListDatacentersResponse struct {
    // Embedded responses related to the resource.
    Embedded        *DatacenterListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *DatacenterListLinks    `json:"_links"`
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

// ListFoldersResponse represents a custom type struct for Success
type ListFoldersResponse struct {
    // Embedded responses related to the resource.
    Embedded        *FolderListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *FolderListLinks    `json:"_links"`
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

// ListHcmHostsResponse represents a custom type struct for Success
type ListHcmHostsResponse struct {
    // Embedded responses related to the resource.
    Embedded      *HostListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links         *HostListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount  *int64            `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied *string           `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit         *int64            `json:"limit"`
    // The page token used to get this response.
    Start         *string           `json:"start"`
}

// ListHostsResponse represents a custom type struct for Success
type ListHostsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *HostListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *HostListLinks    `json:"_links"`
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

// ListMssqlAGsResponse represents a custom type struct for Success
type ListMssqlAGsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *MssqlAGListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *MssqlAGListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64               `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string              `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64               `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string              `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64               `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64               `json:"total_pages_count"`
}

// ListMssqlDatabaseBackupsResponse represents a custom type struct for Success
type ListMssqlDatabaseBackupsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *MssqlDatabaseBackupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *MssqlDatabaseBackupListLinks    `json:"_links"`
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

// ListMssqlDatabasePitrIntervalsResponse represents a custom type struct.
// ListMssqlDatabasePitrIntervalsResponse represents the success response
type ListMssqlDatabasePitrIntervalsResponse struct {
    // Embedded responses related to the resource.
    Embedded      *MssqlDatabasePitrIntervalListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links         *MssqlDatabasePitrIntervalListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount  *int64                                 `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied *string                                `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit         *int64                                 `json:"limit"`
    // The page token used to get this response.
    Start         *string                                `json:"start"`
}

// ListMssqlDatabasesResponse represents a custom type struct for Success
type ListMssqlDatabasesResponse struct {
    // Embedded responses related to the resource.
    Embedded        *MssqlDatabaseListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *MssqlDatabaseListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                     `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                    `json:"filter_applied"`
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

// ListMssqlHostsResponse represents a custom type struct for Success
type ListMssqlHostsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *MssqlHostListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *MssqlHostListLinks    `json:"_links"`
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

// ListMssqlInstancesResponse represents a custom type struct for Success
type ListMssqlInstancesResponse struct {
    // Embedded responses related to the resource.
    Embedded        *MssqlInstanceListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *MssqlInstanceListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                     `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                    `json:"filter_applied"`
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

// ListResourcePoolsResponse represents a custom type struct for Success
type ListResourcePoolsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *ResourcePoolListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *ResourcePoolListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                    `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                   `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                    `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                   `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                    `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                    `json:"total_pages_count"`
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

// ListSubgroupsResponse represents a custom type struct for Success
type ListSubgroupsResponse struct {
    // Embedded responses related to the resource.
    Embedded     *SubgroupListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links        *SubgroupListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount *int64                `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit        *int64                `json:"limit"`
    // The total count of subgroups upto 20. Any number of subgroups beyond 20
    // will still be returned as 20.
    MinCount     *int64                `json:"min_count"`
    // The page token used to get this response.
    Start        *string               `json:"start"`
}

// ListTagCategories2Response represents a custom type struct for Success
type ListTagCategories2Response struct {
    // Embedded responses related to the resource
    Embedded        *TagCategory2ListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource
    Links           *TagCategory2ListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                    `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                    `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                   `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                    `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                    `json:"total_pages_count"`
}

// ListTagsResponse represents a custom type struct for Success
type ListTagsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *Tag2ListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *Tag2ListLinks    `json:"_links"`
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

// ListVMBackupsResponse represents a custom type struct for Success
type ListVMBackupsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *VMBackupListEmbedded     `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *VMBackupListHateoasLinks `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64                    `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string                   `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64                    `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string                   `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64                    `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64                    `json:"total_pages_count"`
}

// ListVMwareDatastoresResponse represents a custom type struct for Success
type ListVMwareDatastoresResponse struct {
    // Embedded responses related to the resource.
    Embedded        *VMwareDatastoreListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *VMwareDatastoreListLinks    `json:"_links"`
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

// ListVMwareVCenterNetworksResponse represents a custom type struct for Success
type ListVMwareVCenterNetworksResponse struct {
    // Embedded responses related to the resource.
    Embedded        *VMwareVCenterNetworkListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *VMwareVCenterNetworkListLinks    `json:"_links"`
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

// ListVcentersResponse represents a custom type struct for Success
type ListVcentersResponse struct {
    // Embedded responses related to the resource.
    Embedded        *VcenterListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *VcenterListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64               `json:"current_count"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64               `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string              `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64               `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64               `json:"total_pages_count"`
}

// ListVmsResponse represents a custom type struct for Success
type ListVmsResponse struct {
    // Embedded responses related to the resource.
    Embedded        *VmListEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links           *VmListLinks    `json:"_links"`
    // The number of items listed on the current page.
    CurrentCount    *int64          `json:"current_count"`
    // The filter used in the request. The filter includes both manually-specified and system-generated filters.
    FilterApplied   *string         `json:"filter_applied"`
    // The maximum number of items displayed per page in the response.
    Limit           *int64          `json:"limit"`
    // The page number used to get this response.
    // Pages are indexed starting from 1 (i.e., `"start": "1"`).
    Start           *string         `json:"start"`
    // The total number of items, summed across all pages.
    TotalCount      *int64          `json:"total_count"`
    // The total number of pages of results.
    TotalPagesCount *int64          `json:"total_pages_count"`
}

// MoveHcmHostsResponse represents a custom type struct for Success
type MoveHcmHostsResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasLinks `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *MoveHostsLinks       `json:"_links"`
    // TaskID for MoveHostsReq
    TaskId   *string               `json:"task_id"`
}

// OnDemandMssqlBackupResponse represents a custom type struct for Success
type OnDemandMssqlBackupResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasLinks `json:"_embedded"`
    // Embedded responses related to the resource.
    Links    *ReadTaskHateoasLinks `json:"_links"`
}

// PatchGeneralSettingsResponseV2 represents a custom type struct for Success
type PatchGeneralSettingsResponseV2 struct {
    // URLs to pages related to the resource.
    Links                        *GeneralSettingsLinks `json:"_links"`
    // The length of time before a user is logged out of the Clumio system due to inactivity. Measured in seconds.
    // The valid range is between `600` seconds (10 minutes) and `3600` seconds (60 minutes).
    // If not configured, the value defaults to `900` seconds (15 minutes).
    AutoLogoutDuration           *int64                `json:"auto_logout_duration"`
    // The designated range of IP addresses that are allowed to access the Clumio REST API.
    // API requests that originate from outside this list will be blocked.
    // The IP address of the server from which this request is being made must be in this list; otherwise, the request will fail.
    // Set the parameter to individual IP addresses and/or a range of IP addresses in CIDR notation.
    // For example, `["193.168.1.0/24", "193.172.1.1"]`.
    // If not configured, the value defaults to ["0.0.0.0/0"] meaning all addresses will be allowed.
    IpAllowlist                  []*string             `json:"ip_allowlist"`
    // The grouping criteria for each datasource type.
    // These can only be edited for datasource types which do not have any
    // organizational units configured.
    OrganizationalUnitDataGroups *OUGroupingCriteria   `json:"organizational_unit_data_groups"`
    // The length of time a user password is valid before it must be changed. Measured in seconds.
    // The valid range is between `2592000` seconds (30 days) and `15552000` seconds (180 days).
    // If not configured, the value defaults to `7776000` seconds (90 days).
    PasswordExpirationDuration   *int64                `json:"password_expiration_duration"`
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
    // Number of users to whom this organizational unit or any of its descendants have been assigned.
    UserCount                 *int64                   `json:"user_count"`
    // Users IDs to whom the organizational unit has been assigned.
    // This attribute will be available when reading a single OU and not when listing OUs.
    Users                     []*string                `json:"users"`
}

// ReadAWSConnectionResponse represents a custom type struct for Success
type ReadAWSConnectionResponse struct {
    // URLs to pages related to the resource.
    Links                    *AWSConnectionLinks `json:"_links"`
    // The alias given to the account on AWS.
    AccountName              *string             `json:"account_name"`
    // The AWS-assigned ID of the account associated with the connection.
    AccountNativeId          *string             `json:"account_native_id"`
    // The AWS region associated with the connection. For example, `us-east-1`.
    AwsRegion                *string             `json:"aws_region"`
    // Clumio AWS AccountId
    ClumioAwsAccountId       *string             `json:"clumio_aws_account_id"`
    // Clumio AWS Region
    ClumioAwsRegion          *string             `json:"clumio_aws_region"`
    // The status of the connection. Possible values include "connecting",
    // "connected", and "unlinked".
    ConnectionStatus         *string             `json:"connection_status"`
    // The timestamp of when the connection was created.
    CreatedTimestamp         *string             `json:"created_timestamp"`
    // An optional, user-provided description for this connection.
    Description              *string             `json:"description"`
    // The configuration of the Clumio Discover product for this connection.
    // If this connection is not configured for Clumio Discover, then this field has a
    // value of `null`.
    Discover                 *DiscoverConfig     `json:"discover"`
    // The Clumio-assigned ID of the connection.
    Id                       *string             `json:"id"`
    // Lambda Bucket Prefix required for configuring Clumio Protect Warm-tier AWS Lambda.
    LambdaBucketPrefix       *string             `json:"lambda_bucket_prefix"`
    // K8S Namespace
    Namespace                *string             `json:"namespace"`
    // The Clumio-assigned ID of the organizational unit associated with the
    // AWS environment. If this parameter is not provided, then the value
    // defaults to the first organizational unit assigned to the requesting
    // user. For more information about organizational units, refer to the
    // Organizational-Units documentation.
    OrganizationalUnitId     *string             `json:"organizational_unit_id"`
    // The configuration of the Clumio Cloud Protect product for this connection.
    // If this connection is not configured for Clumio Cloud Protect, then this field has a
    // value of `null`.
    Protect                  *ProtectConfig      `json:"protect"`
    // The asset types enabled for protect. This is only populated if "protect"
    // is enabled. Valid values are any of ["EBS", "RDS", "DynamoDB", "EC2MSSQL"].
    // EBS and RDS are mandatory datasources.
    ProtectAssetTypesEnabled []*string           `json:"protect_asset_types_enabled"`
    // The services to be enabled for this configuration. Valid values are
    // ["discover"], ["discover", "protect"]. This is only set when the
    // registration is created, the enabled services are obtained directly from
    // the installed template after that.
    ServicesEnabled          []*string           `json:"services_enabled"`
    // The Amazon Resource Name of the installed CloudFormation stack in this AWS account
    StackArn                 *string             `json:"stack_arn"`
    // The name given to the installed CloudFormation stack on AWS.
    StackName                *string             `json:"stack_name"`
    // Subnets Lambda Object Key required for configuring Clumio Protect Warm-tier AWS Lambda.
    SubnetsLambdaObjectKey   *string             `json:"subnets_lambda_object_key"`
    // The 36-character Clumio AWS integration ID token used to identify the
    // installation of the CloudFormation template on the account. This value
    // will be pasted into the ClumioToken field when creating the
    // CloudFormation stack.
    Token                    *string             `json:"token"`
}

// ReadAWSEnvironmentResponse represents a custom type struct for Success
type ReadAWSEnvironmentResponse struct {
    // Embedded responses related to the resource.
    Embedded             *AWSEnvironmentEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                *AWSEnvironmentLinks    `json:"_links"`
    // The name given to the account.
    AccountName          *string                 `json:"account_name"`
    // The AWS-assigned ID of the account associated with the environment.
    AccountNativeId      *string                 `json:"account_native_id"`
    // The valid AWS availability zones for the environment. For example, `us_west-2a`.
    AwsAz                []*string               `json:"aws_az"`
    // The AWS region associated with the environment. For example, `us-west-2`.
    AwsRegion            *string                 `json:"aws_region"`
    // The Clumio-assigned ID of the connection associated with the environment.
    ConnectionId         *string                 `json:"connection_id"`
    // The status of the connection to the environment, which is mediated by a CloudFormation stack.
    ConnectionStatus     *string                 `json:"connection_status"`
    // The user-provided account description.
    Description          *string                 `json:"description"`
    // The Clumio-assigned ID of the environment.
    Id                   *string                 `json:"id"`
    // The Clumio-assigned ID of the organizational unit associated with the environment.
    OrganizationalUnitId *string                 `json:"organizational_unit_id"`
    // The AWS services enabled for this environment. Possible values include "ebs", "rds" and "dynamodb".
    ServicesEnabled      []*string               `json:"services_enabled"`
    // The Clumio CloudFormation template version used to deploy the CloudFormation stack.
    TemplateVersion      *int64                  `json:"template_version"`
}

// ReadAWSTemplatesResponse represents a custom type struct for Success
type ReadAWSTemplatesResponse struct {
    // The latest available CloudFormation template for Clumio Discover.
    Discover *DiscoverTemplateInfo `json:"discover"`
    // The latest available CloudFormation template for Clumio Cloud Protect.
    Protect  *ProtectTemplateInfo  `json:"protect"`
}

// ReadAWSTemplatesV2Response represents a custom type struct for Success
type ReadAWSTemplatesV2Response struct {
    // TODO: Add struct field description
    Config *TemplateConfiguration `json:"config"`
}

// ReadAlertResponse represents a custom type struct for Success
type ReadAlertResponse struct {
    // Embedded responses related to the resource.
    Embedded            *AlertEmbedded          `json:"_embedded"`
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
    // The primary object associated with or affected by the alert. Examples of primary entities include "aws_connection", "aws_ebs_volume" and "vmware_vm".
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

// ReadComputeResourceResponse represents a custom type struct for Success
type ReadComputeResourceResponse struct {
    // Embedded responses related to the resource.
    Embedded              *ComputeResourceEmbedded                     `json:"_embedded"`
    // The ETag value.
    Etag                  *string                                      `json:"_etag"`
    // URLs to pages related to the resource.
    Links                 *ComputeResourceLinks                        `json:"_links"`
    // The compute resource folder in which the compute resource resides.
    ComputeResourceFolder *VMwareVCenterComputeResourceFolderModel     `json:"compute_resource_folder"`
    // The data center associated with this compute resource.
    Datacenter            *VMwareVCenterComputeResourceDatacenterModel `json:"datacenter"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the compute resource.
    Id                    *string                                      `json:"id"`
    // Determines whether the compute resource is a cluster. If `true`, then the compute resource is a cluster.
    IsCluster             *bool                                        `json:"is_cluster"`
    // Determines whether the compute resource has Distributed Resource Scheduler (DRS) enabled. If this field and `"is_cluster":true`, then DRS is enabled in the compute resource cluster.
    IsDrsEnabled          *bool                                        `json:"is_drs_enabled"`
    // The VMware-assigned name of the compute resource.
    Name                  *string                                      `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the compute resource.
    OrganizationalUnitId  *string                                      `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo        *ProtectionInfo                              `json:"protection_info"`
    // The protection status of this compute resource. Refer to the Protection Status table for a complete list of protection statuses.
    ProtectionStatus      *string                                      `json:"protection_status"`
}

// ReadConsolidatedAlertResponse represents a custom type struct for Success
type ReadConsolidatedAlertResponse struct {
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

// ReadDatacenterResponse represents a custom type struct for Success
type ReadDatacenterResponse struct {
    // Embedded responses related to the resource.
    Embedded                  *DatacenterEmbedded                     `json:"_embedded"`
    // The ETag value.
    Etag                      *string                                 `json:"_etag"`
    // URLs to pages related to the resource.
    Links                     *DatacenterLinks                        `json:"_links"`
    // TODO: Add struct field description
    AncestorRefs              []*AncestorRefModel                     `json:"ancestor_refs"`
    // The data center folder in which the data center resides.
    DatacenterFolder          *VMwareDatacenterFolderIDModel          `json:"datacenter_folder"`
    // Determines whether compute resources exist directly under the hidden root compute resource folder. If `true`, then compute resources exist directly under the root compute resource folder.
    HasComputeResources       *bool                                   `json:"has_compute_resources"`
    // Determines whether VMs exist directly under the hidden root VM folder. If `true`, then VMs exist directly under the root VM folder.
    HasVmFolders              *bool                                   `json:"has_vm_folders"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the data center.
    Id                        *string                                 `json:"id"`
    // The VMware-assigned name of this data center.
    Name                      *string                                 `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the datacenter.
    OrganizationalUnitId      *string                                 `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo            *ProtectionInfo                         `json:"protection_info"`
    // The protection status of this data center. Refer to the Protection Status table for a complete list of protection statuses.
    ProtectionStatus          *string                                 `json:"protection_status"`
    // The hidden root compute resource folder of the data center.
    RootComputeResourceFolder *VMwareRootComputeResourceFolderIDModel `json:"root_compute_resource_folder"`
    // The hidden root virtual machine folder of the data center.
    RootVmFolder              *VMwareRootVMFolderIDModel              `json:"root_vm_folder"`
}

// ReadDirectoryResponse represents a custom type struct for Success
type ReadDirectoryResponse struct {
    // TODO: Add struct field description
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

// ReadEbsTagComplianceStatsResponse represents a custom type struct for Success
type ReadEbsTagComplianceStatsResponse struct {
    // The total number of compliant entities.
    CompliantCount     *int64 `json:"compliant_count"`
    // The total number of entities associated with deactivated policies.
    DeactivatedCount   *int64 `json:"deactivated_count"`
    // Determines whether one or more entities is currently seeding or waiting for seeding.
    // If set to `true`, at least one entity is currently seeding or waiting for seeding.
    HasSeedingEntities *bool  `json:"has_seeding_entities"`
    // The total number of non-compliant entities.
    NonCompliantCount  *int64 `json:"non_compliant_count"`
    // The number of entities with protection applied.
    ProtectedCount     *int64 `json:"protected_count"`
    // The number of entities without protection applied.
    UnprotectedCount   *int64 `json:"unprotected_count"`
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
    // The compliance status of the protected EBS volume. Possible values include
    // "compliant" and "noncompliant". If the volume is not protected, then this field has
    // a value of `null`.
    ComplianceStatus         *string                 `json:"compliance_status"`
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

// ReadFileSystemResponse represents a custom type struct for Success
type ReadFileSystemResponse struct {
    // TODO: Add struct field description
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
    // The total amount of memory available to the filesystem in bytes.
    Size                 *uint64             `json:"size"`
    // The type of the filesystem. This field is populated with values returned from
    // the lsblk command. Possible values include `ntfs`, `xfs`, and `ext3`.
    ClumioType           *string             `json:"type"`
    // The amount of memory used by the filesystem in bytes.
    Used                 *uint64             `json:"used"`
}

// ReadFolderResponse represents a custom type struct for Success
type ReadFolderResponse struct {
    // Embedded responses related to the resource.
    Embedded              *FolderEmbedded                     `json:"_embedded"`
    // The ETag value.
    Etag                  *string                             `json:"_etag"`
    // URLs to pages related to the resource.
    Links                 *FolderLinks                        `json:"_links"`
    // The data center associated with this folder.
    Datacenter            *VMwareVCenterFolderDatacenterModel `json:"datacenter"`
    // Count of all descendant folders inside this folder
    DescendantFolderCount *int64                              `json:"descendant_folder_count"`
    // Determines whether the folder has direct child folders.
    HasChildFolders       *bool                               `json:"has_child_folders"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the folder.
    Id                    *string                             `json:"id"`
    // Determines whether the folder is a hidden root folder. If `true`, the folder is a hidden root folder.
    IsRoot                *bool                               `json:"is_root"`
    // Determines whether the folder can be used as a restore destination. If `true`, the folder can be used as a restore destination, and backups can be restored to the folder.
    IsSupported           *bool                               `json:"is_supported"`
    // The VMware-assigned name of the folder.
    Name                  *string                             `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the folder.
    OrganizationalUnitId  *string                             `json:"organizational_unit_id"`
    // The parent folder under which this folder resides.
    ParentFolder          *VMwareVCenterParentFolderModel     `json:"parent_folder"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo        *ProtectionInfo                     `json:"protection_info"`
    // The protection status of this folder. Refer to the Protection Status table for a complete list of protection statuses.
    ProtectionStatus      *string                             `json:"protection_status"`
    // The folder type. Examples of folder types include "datacenter_folder" and "vm_folder". Refer to the Folder Type table for a complete list of folder types.
    ClumioType            *string                             `json:"type"`
}

// ReadGeneralSettingsResponseV2 represents a custom type struct for Success
type ReadGeneralSettingsResponseV2 struct {
    // URLs to pages related to the resource.
    Links                        *GeneralSettingsLinks `json:"_links"`
    // The length of time before a user is logged out of the Clumio system due to inactivity. Measured in seconds.
    // The valid range is between `600` seconds (10 minutes) and `3600` seconds (60 minutes).
    // If not configured, the value defaults to `900` seconds (15 minutes).
    AutoLogoutDuration           *int64                `json:"auto_logout_duration"`
    // The designated range of IP addresses that are allowed to access the Clumio REST API.
    // API requests that originate from outside this list will be blocked.
    // The IP address of the server from which this request is being made must be in this list; otherwise, the request will fail.
    // Set the parameter to individual IP addresses and/or a range of IP addresses in CIDR notation.
    // For example, `["193.168.1.0/24", "193.172.1.1"]`.
    // If not configured, the value defaults to ["0.0.0.0/0"] meaning all addresses will be allowed.
    IpAllowlist                  []*string             `json:"ip_allowlist"`
    // The grouping criteria for each datasource type.
    // These can only be edited for datasource types which do not have any
    // organizational units configured.
    OrganizationalUnitDataGroups *OUGroupingCriteria   `json:"organizational_unit_data_groups"`
    // The length of time a user password is valid before it must be changed. Measured in seconds.
    // The valid range is between `2592000` seconds (30 days) and `15552000` seconds (180 days).
    // If not configured, the value defaults to `7776000` seconds (90 days).
    PasswordExpirationDuration   *int64                `json:"password_expiration_duration"`
}

// ReadHcmHostResponse represents a custom type struct for Success
type ReadHcmHostResponse struct {
    // URLs to pages related to the resource.
    Links                  *HostLinks `json:"_links"`
    // The endpoints discovered post the host connection of the host.
    DiscoveredEndpoints    []*string  `json:"discovered_endpoints"`
    // The Current MSI version of the edge connector installed in the host.
    EdgeConnectorVersion   *string    `json:"edge_connector_version"`
    // The user-provided endpoint used to connect the host.
    Endpoint               *string    `json:"endpoint"`
    // The Clumio-assigned ID of the management group associated with the host.
    GroupId                *string    `json:"group_id"`
    // The Clumio-assigned ID of the Host.
    Id                     *string    `json:"id"`
    // The timestamp of the last successful heartbeat of this host. Represented in RFC-3339 format.
    LastHeartbeatTimestamp *string    `json:"last_heartbeat_timestamp"`
    // Name of the Host.
    Name                   *string    `json:"name"`
    // The operational status of the Host. Possible values include `upgrade_in_progress`, `upgrade_success`, `upgrade_failed`, `delete_in_progress`, `delete_failed`, `move_in_progress`, `move_succeeded` and `move_failed`.
    OperationalStatus      *string    `json:"operational_status"`
    // The connection status of the Host. Possible values include `connected`, `disconnected`, `connection_pending`, and `invalid_token`.
    Status                 *string    `json:"status"`
    // The Clumio-assigned ID of the subgroup associated with the host.
    SubgroupId             *string    `json:"subgroup_id"`
    // The Clumio-assigned UUID of the host. This UUID is used for filtering hosts during list operations.
    Uuid                   *string    `json:"uuid"`
}

// ReadHostResponse represents a custom type struct for Success
type ReadHostResponse struct {
    // The ETag value.
    Etag                *string                                `json:"_etag"`
    // URLs to pages related to the resource.
    Links               *HostLinks                             `json:"_links"`
    // The VMware compute resource representing the host.
    ComputeResource     *VMwareVCenterHostComputeResourceModel `json:"compute_resource"`
    // The connection state of the host as seen through the vCenter server. Examples include "connected", "disconnected", and "not_responding".
    ConnectionState     *string                                `json:"connection_state"`
    // The data center in which the host resides.
    Datacenter          *VMwareVCenterHostDatacenterModel      `json:"datacenter"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the host.
    Id                  *string                                `json:"id"`
    // Determines whether the host has been placed in maintenance mode as seen through the vCenter server. If `true`, the host is in maintenance mode.
    IsInMaintenanceMode *bool                                  `json:"is_in_maintenance_mode"`
    // Determines whether the host has been placed in quarantine mode as seen through the vCenter server. If `true`, the host is in quarantine mode.
    IsInQuarantineMode  *bool                                  `json:"is_in_quarantine_mode"`
    // Determines whether the host is a standalone host. If `true`, the host is a standalone host.
    IsStandalone        *bool                                  `json:"is_standalone"`
    // Determines whether the host can be used as a restore destination. If `true`, the host can be used as a restore destination and backups can be restored to the host.
    IsSupported         *bool                                  `json:"is_supported"`
    // The VMware-assigned name of the host.
    Name                *string                                `json:"name"`
    // The power state of the host as seen through the vCenter server. Examples include "powered_off", "powered_on", and "standby".
    PowerState          *string                                `json:"power_state"`
}

// ReadManagementGroupResponse represents a custom type struct for Success
type ReadManagementGroupResponse struct {
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

// ReadMssqlAGResponse represents a custom type struct for Success
type ReadMssqlAGResponse struct {
    // MssqlAGEmbedded is embed of MSSQL Availability Groups
    // Embedded responses related to the resource.
    Embedded             *MssqlAGEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                *MssqlAGLinks    `json:"_links"`
    // The Clumio-assigned ID of the availability group.
    Id                   *string          `json:"id"`
    // The Microsoft SQL-assigned name of the availability group.
    Name                 *string          `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the availability group.
    OrganizationalUnitId *string          `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo       *ProtectionInfo  `json:"protection_info"`
    // The status of the availability group, Possible values include 'active' and 'inactive'.
    Status               *string          `json:"status"`
}

// ReadMssqlDatabaseBackupResponse represents a custom type struct for Success
type ReadMssqlDatabaseBackupResponse struct {
    // Embedded responses related to the resource.
    Embedded            *MssqlDatabaseBackupEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links               *MssqlDatabaseBackupLinks    `json:"_links"`
    // TODO: Add struct field description
    DatabaseFiles       []*MssqlDatabaseFile         `json:"database_files"`
    // The Clumio-assigned ID of the database associated with this backup.
    DatabaseId          *string                      `json:"database_id"`
    // The Microsoft SQL database engine at the time of backup.
    Engine              *string                      `json:"engine"`
    // The Microsoft SQL database engine version at the time of backup.
    EngineVersion       *string                      `json:"engine_version"`
    // The timestamp of when this backup expires. Represented in RFC-3339 format.
    ExpirationTimestamp *string                      `json:"expiration_timestamp"`
    // The Clumio-assigned ID of the management group associated with the database at the time of backup.
    GroupId             *string                      `json:"group_id"`
    // The user-provided endpoint of the host containing the given database at the time of backup.
    HostEndpoint        *string                      `json:"host_endpoint"`
    // The Clumio-assigned ID of the host associated with the database at the time of backup.
    HostId              *string                      `json:"host_id"`
    // The Clumio-assigned ID of the backup.
    Id                  *string                      `json:"id"`
    // The Clumio-assigned instance id at the time of backup.
    InstanceId          *string                      `json:"instance_id"`
    // The instance name at the time of backup.
    InstanceName        *string                      `json:"instance_name"`
    // The timestamp of when this backup started. Represented in RFC-3339 format.
    StartTimestamp      *string                      `json:"start_timestamp"`
    // The Clumio-assigned ID of the management subgroup associated with the database at the time of backup.
    SubgroupId          *string                      `json:"subgroup_id"`
    // The type of backup. Possible values include `mssql_database_backup`, `mssql_log_backup_full_recovery_model` and `mssql_log_backup_bulk_logged_model`.
    ClumioType          *string                      `json:"type"`
}

// ReadMssqlDatabaseResponse represents a custom type struct for Success
type ReadMssqlDatabaseResponse struct {
    // Embedded responses related to the resource.
    Embedded                                *MssqlDatabaseEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                                   *DatabaseLinks         `json:"_links"`
    // The Clumio-assigned ID of the availability group. It is null in case of a standalone database.
    AvailabilityGroupId                     *string                `json:"availability_group_id"`
    // The Microsoft SQL assigned name of the availability group. It is null in case of a standalone database.
    AvailabilityGroupName                   *string                `json:"availability_group_name"`
    // The policy compliance status of the resource. If the database is not protected,
    // then this field has a value of `null`. Refer to
    // 
    // the Compliance Status table
    // 
    // for a complete list of compliance statuses.
    ComplianceStatus                        *string                `json:"compliance_status"`
    // The Clumio-assigned ID of the group to which the standalone database belongs, in case of an
    // availability group database it will be empty.
    GroupId                                 *string                `json:"group_id"`
    // The user-provided endpoint of the host containing the given database.
    HostEndpoint                            *string                `json:"host_endpoint"`
    // The Clumio-assigned ID of the host containing the given database.
    HostId                                  *string                `json:"host_id"`
    // The Clumio-assigned ID of the Database.
    Id                                      *string                `json:"id"`
    // The Clumio-assigned ID of the instance containing the given database.
    InstanceId                              *string                `json:"instance_id"`
    // The name of the Microsoft SQL instance containing the given database.
    InstanceName                            *string                `json:"instance_name"`
    // is_supported is true if Clumio supports backup of the database.
    IsSupported                             *bool                  `json:"is_supported"`
    // The timestamp of the last time this database was full backed up.
    // Represented in RFC-3339 format. If this database has never been backed up,
    // this field has a value of `null`.
    LastBackupTimestamp                     *string                `json:"last_backup_timestamp"`
    // The timestamp of the last time this database was log backed up in Bulk Recovery Model.
    // Represented in RFC-3339 format. If this database has never been backed up,
    // this field has a value of `null`.
    LastBulkRecoveryModelLogBackupTimestamp *string                `json:"last_bulk_recovery_model_log_backup_timestamp"`
    // The timestamp of the last time this database was log backed up in Full Recovery Model.
    // Represented in RFC-3339 format. If this database has never been backed up,
    // this field has a value of `null`.
    LastFullRecoveryModelLogBackupTimestamp *string                `json:"last_full_recovery_model_log_backup_timestamp"`
    // The name of the Database.
    Name                                    *string                `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the database.
    OrganizationalUnitId                    *string                `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo                          *ProtectionInfo        `json:"protection_info"`
    // recovery_model is the recovery model of the database. Possible values include 'simple_recovery_model',
    // 'bulk_recovery_model', and 'full_recovery_model'.
    RecoveryModel                           *string                `json:"recovery_model"`
    // The size of the Database.
    Size                                    *float64               `json:"size"`
    // The status of the database, Possible values include 'active' and 'inactive'.
    Status                                  *string                `json:"status"`
    // subgroup id is the id of the Subgroup where this database belongs, in case of AG database
    // it will be empty.
    SubgroupId                              *string                `json:"subgroup_id"`
    // The type of the database. Possible values include 'availability_group_database' and 'standalone_database'.
    ClumioType                              *string                `json:"type"`
    // unsupported_reason is the reason why Clumio doesn't support backup of such database,
    // possible values include 'filestream_enabled_database'.
    UnsupportedReason                       *string                `json:"unsupported_reason"`
}

// ReadMssqlHostResponse represents a custom type struct for Success
type ReadMssqlHostResponse struct {
    // Embedded responses related to the resource.
    Embedded                       *MssqlHostEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                          *MssqlHostLinks    `json:"_links"`
    // The user-provided endpoint of the host containing the given database.
    Endpoint                       *string            `json:"endpoint"`
    // The Clumio-assigned ID of the management group to which the host belongs.
    GroupId                        *string            `json:"group_id"`
    // Determines whether or not an availability group is present in the host.
    HasAssociatedAvailabilityGroup *bool              `json:"has_associated_availability_group"`
    // The Clumio-assigned ID of the Host.
    Id                             *string            `json:"id"`
    // The number of instances present in the host.
    InstanceCount                  *int64             `json:"instance_count"`
    // The Clumio-assigned ID of the organizational unit associated with the host.
    OrganizationalUnitId           *string            `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo                 *ProtectionInfo    `json:"protection_info"`
    // The status of the Host, Possible values include 'active' and 'inactive'.
    Status                         *string            `json:"status"`
    // The Clumio-assigned ID of the management subgroup to which the host belongs.
    SubgroupId                     *string            `json:"subgroup_id"`
}

// ReadMssqlInstanceResponse represents a custom type struct for Success
type ReadMssqlInstanceResponse struct {
    // Embedded responses related to the resource.
    Embedded                       *MssqlInstanceEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                          *MssqlInstanceLinks    `json:"_links"`
    // The Clumio-assigned ID of the management group to which the host belongs.
    GroupId                        *string                `json:"group_id"`
    // The boolean value represents if availability group is present in the instance.
    HasAssociatedAvailabilityGroup *bool                  `json:"has_associated_availability_group"`
    // The user-provided endpoint of the host containing the given database.
    HostEndpoint                   *string                `json:"host_endpoint"`
    // The Clumio-assigned ID of the host, containing the instance.
    HostId                         *string                `json:"host_id"`
    // The Clumio-assigned ID of the Instance.
    Id                             *string                `json:"id"`
    // The Microsoft SQL assigned name of the instance.
    Name                           *string                `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the instance.
    OrganizationalUnitId           *string                `json:"organizational_unit_id"`
    // Product Version of the instance.
    ProductVersion                 *string                `json:"product_version"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo                 *ProtectionInfo        `json:"protection_info"`
    // The Microsoft SQL assigned server name of the instance.
    ServerName                     *string                `json:"server_name"`
    // The status of the Instance, Possible values include 'active' and 'inactive'.
    Status                         *string                `json:"status"`
    // The Clumio-assigned ID of the management subgroup to which the host belongs.
    SubgroupId                     *string                `json:"subgroup_id"`
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
    Embedded                      *PolicyEmbedded    `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                         *PolicyLinks       `json:"_links"`
    // The status of the policy.
    // Refer to the Policy Activation Status table for a complete list of policy statuses.
    ActivationStatus              *string            `json:"activation_status"`
    // The Clumio-assigned IDs of the organizational units to whom the policy has been assigned.
    AssignedOrganizationalUnitIds []*string          `json:"assigned_organizational_unit_ids"`
    // The Clumio-assigned ID of the policy.
    Id                            *string            `json:"id"`
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
    LockStatus                    *string            `json:"lock_status"`
    // The user-provided name of the policy.
    Name                          *string            `json:"name"`
    // The SLAs of an individual operation.
    Operations                    []*PolicyOperation `json:"operations"`
    // The Clumio-assigned ID of the organizational unit associated with the policy.
    OrganizationalUnitId          *string            `json:"organizational_unit_id"`
}

// ReadResourcePoolResponse represents a custom type struct for Success
type ReadResourcePoolResponse struct {
    // The ETag value.
    Etag            *string                                 `json:"_etag"`
    // URLs to pages related to the resource.
    Links           *ResourcePoolLinks                      `json:"_links"`
    // The compute resource that the resource pool comprises.
    ComputeResource *VMwareResourcePoolComputeResourceModel `json:"compute_resource"`
    // The data center in which the resource pool resides.
    Datacenter      *ResourcePoolDatacenterModel            `json:"datacenter"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the resource pool.
    Id              *string                                 `json:"id"`
    // Determines whether the resource pool is the default, hidden resource pool.
    IsRoot          *bool                                   `json:"is_root"`
    // Determines whether the resource pool can be used as a restore destination. If `true`, the resource pool can be used as a restore destination and backups can be restored to the resource pool.
    IsSupported     *bool                                   `json:"is_supported"`
    // The VMware-assigned name of the resource pool.
    Name            *string                                 `json:"name"`
    // The vCenter object that is the parent of the resource pool.
    Parent          *VMwareResourcePoolParentModel          `json:"parent"`
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

// ReadSubgroupResponse represents a custom type struct for Success
type ReadSubgroupResponse struct {
    // URLs to pages related to the resource.
    Links                       *SubgroupLinks               `json:"_links"`
    // The number of cloud connectors in this subgroup, aggregated by their status.
    CloudConnectorCountByStatus *CloudConnectorCountByStatus `json:"cloud_connector_count_by_status"`
    // The overall health of cloud connectors in this subgroup. Possible values include: 'healthy', indicating
    // that all cloud connectors in the subgroup are connected; 'degraded' indicating that one or more cloud
    // connectors in the subgroup have connection issues; `none`, indicating that no cloud connectors are in the subgroup.
    CloudConnectorStatus        *string                      `json:"cloud_connector_status"`
    // The Clumio-assigned ID of the management group associated with this subgroup.
    GroupId                     *string                      `json:"group_id"`
    // The Clumio-assigned ID of the management subgroup.
    Id                          *string                      `json:"id"`
    // The name of the management subgroup.
    Name                        *string                      `json:"name"`
}

// ReadTagCategory2Response represents a custom type struct for Success
type ReadTagCategory2Response struct {
    // The ETag value.
    Etag                 *string            `json:"_etag"`
    // URLs to pages related to the resource
    Links                *TagCategory2Links `json:"_links"`
    // A description of the tag category.
    Description          *string            `json:"description"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the tag category.
    Id                   *string            `json:"id"`
    // The VMware-assigned name of the tag category.
    Name                 *string            `json:"name"`
    // The number of tags in the tag category.
    NumberOfTags         *int32             `json:"number_of_tags"`
    // The Clumio-assigned ID of the organizational unit associated with the tag category.
    OrganizationalUnitId *string            `json:"organizational_unit_id"`
}

// ReadTagResponse represents a custom type struct for Success
type ReadTagResponse struct {
    // Embedded responses related to the resource.
    Embedded             *Tag2Embedded           `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                *Tag2Links              `json:"_links"`
    // The tag category associated with the tag.
    Category             *TagParentCategoryModel `json:"category"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the tag.
    Id                   *string                 `json:"id"`
    // The VMware-assigned name of the tag.
    Name                 *string                 `json:"name"`
    // The Clumio-assigned ID of the organizational unit associated with the tag.
    OrganizationalUnitId *string                 `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo       *ProtectionInfo         `json:"protection_info"`
    // The protection status of this tag. Refer to the Protection Status table for a complete list of protection statuses.
    ProtectionStatus     *string                 `json:"protection_status"`
}

// ReadTaskResponse represents a custom type struct for Success
type ReadTaskResponse struct {
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
    // Tasks of certain types including
    // "vmware_vm_backup_indexing" and "aws_ebs_volume_backup_indexing" cannot be aborted.
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
    // The task type. Examples of task types include "vm_backup_seeding", "ebs_indexing", and "file_restore".
    // Refer to the Task Type table for a complete list of task types.
    ClumioType         *string            `json:"type"`
}

// ReadUserResponse represents a custom type struct for Success
type ReadUserResponse struct {
    // Embedded responses related to the resource.
    Embedded                      *UserEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                         *UserLinks    `json:"_links"`
    // The list of organizational unit IDs assigned to the user.
    // This attribute will be available when reading a single user and not when listing users.
    AssignedOrganizationalUnitIds []*string     `json:"assigned_organizational_unit_ids"`
    // Assigned Role for the user.
    AssignedRole                  *string       `json:"assigned_role"`
    // The email address of the Clumio user.
    Email                         *string       `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                      *string       `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                            *string       `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                       *string       `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                   *bool         `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                     *bool         `json:"is_enabled"`
    // The timestamp of when when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp         *string       `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount       *int64        `json:"organizational_unit_count"`
}

// ReadVMBackupResponse represents a custom type struct for Success
type ReadVMBackupResponse struct {
    // URLs to pages related to the resource.
    Links                *VMBackupHateoasLinks     `json:"_links"`
    // The reason that browsing is unavailable for the backup. Possible values include "file_limit_exceeded" and
    // "browsing_unavailable". If browse indexing is successful, then this field has a value of `null`.
    BrowsingFailedReason *string                   `json:"browsing_failed_reason"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the data center
    // associated with this backup.
    DatacenterId         *string                   `json:"datacenter_id"`
    // The timestamp of when this backup expires. Represented in RFC-3339 format.
    ExpirationTimestamp  *string                   `json:"expiration_timestamp"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the
    // host associated with this backup.
    HostId               *string                   `json:"host_id"`
    // The Clumio-assigned ID of the backup.
    Id                   *string                   `json:"id"`
    // Determines whether browsing is available for the backup. If `true`, then browsing is available for the backup.
    IsBrowsable          *bool                     `json:"is_browsable"`
    // TODO: Add struct field description
    Nics                 []*VMNicBackupModel       `json:"nics"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the resource pool
    // associated with this backup.
    ResourcePoolId       *string                   `json:"resource_pool_id"`
    // The timestamp of when this backup started. Represented in RFC-3339 format.
    StartTimestamp       *string                   `json:"start_timestamp"`
    // VMTagWithCategoryModel
    // A tag associated with the VM.
    Tags                 []*VMTagWithCategoryModel `json:"tags"`
    // The IP address or FQDN of the vCenter server associated with this backup.
    // If a backup was initiated before 2020-06-30, when this field was introduced,
    // then this field has a value of `null`.
    VcenterEndpoint      *string                   `json:"vcenter_endpoint"`
    // The Clumio-assigned ID of the vCenter associated with this backup.
    VcenterId            *string                   `json:"vcenter_id"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the
    // VM folder associated with this backup.
    VmFolderId           *string                   `json:"vm_folder_id"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the
    // VM associated with this backup.
    VmId                 *string                   `json:"vm_id"`
    // The name of the virtual machine associated with this backup.
    VmName               *string                   `json:"vm_name"`
}

// ReadVMwareComputeResourceStatsResponse represents a custom type struct for Success
type ReadVMwareComputeResourceStatsResponse struct {
    // The ETag value.
    Etag               *string                                    `json:"_etag"`
    // URLs to pages related to the resource.
    Links              *VMwareComputeResourceComplianceStatsLinks `json:"_links"`
    // The total number of compliant entities.
    CompliantCount     *int64                                     `json:"compliant_count"`
    // The total number of entities associated with deactivated policies.
    DeactivatedCount   *int64                                     `json:"deactivated_count"`
    // Determines whether one or more entities is currently seeding or waiting for seeding.
    // If set to `true`, at least one entity is currently seeding or waiting for seeding.
    HasSeedingEntities *bool                                      `json:"has_seeding_entities"`
    // The total number of non-compliant entities.
    NonCompliantCount  *int64                                     `json:"non_compliant_count"`
    // The number of entities with protection applied.
    ProtectedCount     *int64                                     `json:"protected_count"`
    // The number of entities without protection applied.
    UnprotectedCount   *int64                                     `json:"unprotected_count"`
}

// ReadVMwareDatacenterStatsResponse represents a custom type struct for Success
type ReadVMwareDatacenterStatsResponse struct {
    // The ETag value.
    Etag               *string                     `json:"_etag"`
    // URLs to pages related to the resource.
    Links              *VMwareDatacenterStatsLinks `json:"_links"`
    // The total number of compliant entities.
    CompliantCount     *int64                      `json:"compliant_count"`
    // The total number of entities associated with deactivated policies.
    DeactivatedCount   *int64                      `json:"deactivated_count"`
    // Determines whether one or more entities is currently seeding or waiting for seeding.
    // If set to `true`, at least one entity is currently seeding or waiting for seeding.
    HasSeedingEntities *bool                       `json:"has_seeding_entities"`
    // The total number of non-compliant entities.
    NonCompliantCount  *int64                      `json:"non_compliant_count"`
    // The number of entities with protection applied.
    ProtectedCount     *int64                      `json:"protected_count"`
    // The number of entities without protection applied.
    UnprotectedCount   *int64                      `json:"unprotected_count"`
}

// ReadVMwareDatastoreResponse represents a custom type struct for Success
type ReadVMwareDatastoreResponse struct {
    // The ETag value.
    Etag             *string                                `json:"_etag"`
    // URLs to pages related to the resource.
    Links            *VMwareDatastoreLinks                  `json:"_links"`
    // TODO: Add struct field description
    ComputeResources []*ComputeResourceIDModel              `json:"compute_resources"`
    // The data center in which this datastore resides.
    Datacenter       *VMwareVCenterDatastoreDatacenterModel `json:"datacenter"`
    // VMwareVCenterDatastoreFolderModel
    // The datastore folder in which this datastore resides.
    DatastoreFolder  *VMwareVCenterDatastoreFolderModel     `json:"datastore_folder"`
    // The file system format used for the datastore. Refer to the Supported Datastore Types section for a complete list of datastore types.
    DatastoreType    *string                                `json:"datastore_type"`
    // TODO: Add struct field description
    Hosts            []*HostIDModel                         `json:"hosts"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the datastore.
    Id               *string                                `json:"id"`
    // Determines whether the datastore is shared across multiple hosts. If `true`, the datastore is a multi-host datastore.
    IsMultiHost      *bool                                  `json:"is_multi_host"`
    // Determines whether the datastore can be used as a restore destination. If `true`, the datastore can be used as a restore destination and backups can be restored to the datastore.
    IsSupported      *bool                                  `json:"is_supported"`
    // The VMware-assigned name of this datastore.
    Name             *string                                `json:"name"`
}

// ReadVMwareFolderStatsResponse represents a custom type struct for Success
type ReadVMwareFolderStatsResponse struct {
    // The ETag value.
    Etag               *string                 `json:"_etag"`
    // URLs to pages related to the resource.
    Links              *VMwareFolderStatsLinks `json:"_links"`
    // The total number of compliant entities.
    CompliantCount     *int64                  `json:"compliant_count"`
    // The total number of entities associated with deactivated policies.
    DeactivatedCount   *int64                  `json:"deactivated_count"`
    // Determines whether one or more entities is currently seeding or waiting for seeding.
    // If set to `true`, at least one entity is currently seeding or waiting for seeding.
    HasSeedingEntities *bool                   `json:"has_seeding_entities"`
    // The total number of non-compliant entities.
    NonCompliantCount  *int64                  `json:"non_compliant_count"`
    // The number of entities with protection applied.
    ProtectedCount     *int64                  `json:"protected_count"`
    // The number of entities without protection applied.
    UnprotectedCount   *int64                  `json:"unprotected_count"`
}

// ReadVMwareTagStatsResponse represents a custom type struct for Success
type ReadVMwareTagStatsResponse struct {
    // The ETag value.
    Etag               *string              `json:"_etag"`
    // URLs to pages related to the resource.
    Links              *VMwareTagStatsLinks `json:"_links"`
    // The total number of compliant entities.
    CompliantCount     *int64               `json:"compliant_count"`
    // The total number of entities associated with deactivated policies.
    DeactivatedCount   *int64               `json:"deactivated_count"`
    // Determines whether one or more entities is currently seeding or waiting for seeding.
    // If set to `true`, at least one entity is currently seeding or waiting for seeding.
    HasSeedingEntities *bool                `json:"has_seeding_entities"`
    // The total number of non-compliant entities.
    NonCompliantCount  *int64               `json:"non_compliant_count"`
    // The number of entities with protection applied.
    ProtectedCount     *int64               `json:"protected_count"`
    // The number of entities without protection applied.
    UnprotectedCount   *int64               `json:"unprotected_count"`
}

// ReadVMwareVCenterNetworkResponse represents a custom type struct for Success
type ReadVMwareVCenterNetworkResponse struct {
    // The ETag value.
    Etag          *string                              `json:"_etag"`
    // URLs to pages related to the resource.
    Links         *VMwareVCenterNetworkLinks           `json:"_links"`
    // The data center associated with this network.
    Datacenter    *VMwareVCenterNetworkDatacenterModel `json:"datacenter"`
    // The VMware-assigned ID of this network.
    Id            *string                              `json:"id"`
    // Determines whether VMs can be connected to the network. If `true`, VMs can be connected to the network.
    IsSupported   *bool                                `json:"is_supported"`
    // The name of this network.
    Name          *string                              `json:"name"`
    // The network folder associated with this network.
    NetworkFolder *VMwareVCenterNetworkFolderModel     `json:"network_folder"`
}

// ReadVMwareVCenterProtectionStatsResponse represents a custom type struct for Success
type ReadVMwareVCenterProtectionStatsResponse struct {
    // URLs to pages related to the resource.
    Links              *ReadVMwareVCenterProtectionStatsLinks `json:"_links"`
    // The total number of compliant entities.
    CompliantCount     *int64                                 `json:"compliant_count"`
    // The total number of entities associated with deactivated policies.
    DeactivatedCount   *int64                                 `json:"deactivated_count"`
    // Determines whether one or more entities is currently seeding or waiting for seeding.
    // If set to `true`, at least one entity is currently seeding or waiting for seeding.
    HasSeedingEntities *bool                                  `json:"has_seeding_entities"`
    // The total number of non-compliant entities.
    NonCompliantCount  *int64                                 `json:"non_compliant_count"`
    // The number of entities with protection applied.
    ProtectedCount     *int64                                 `json:"protected_count"`
    // The number of entities without protection applied.
    UnprotectedCount   *int64                                 `json:"unprotected_count"`
}

// ReadVcenterResponse represents a custom type struct for Success
type ReadVcenterResponse struct {
    // Embedded responses related to the resource.
    Embedded                  *VcenterEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                     *VcenterLinks    `json:"_links"`
    // The region to which data is backed-up to for the vCenter server. If the vCenter server's back up region is unavailable, this field has a value of `unavailable`. Refer to the Back up Regions table for a complete list of back up regions.
    BackupRegion              *string          `json:"backup_region"`
    // The URL at which the Clumio Cloud Connector for this vCenter server can be downloaded.
    CloudConnectorDownloadUrl *string          `json:"cloud_connector_download_url"`
    // The IP address or FQDN of the vCenter server.
    Endpoint                  *string          `json:"endpoint"`
    // The Clumio-assigned ID of the vCenter server.
    Id                        *string          `json:"id"`
    // The IP address or FQDN of the vCenter server.
    // This field has been replaced by the `endpoint` field
    // and is being retained for backward compatibility reasons.
    IpAddress                 *string          `json:"ip_address"`
    // The Clumio-assigned ID of the organizational unit associated with the vCenter.
    OrganizationalUnitId      *string          `json:"organizational_unit_id"`
    // The connection status of the Clumio Cloud Connector. Examples include "pending", "connected", "disconnected", "invalid_credentials", "partial", and "unavailable".
    Status                    *string          `json:"status"`
    // The type of vCenter server. If the vCenter server's type is unavailable, this field has a value of `unavailable`. Refer to the vCenter Types table for a complete list of vCenter types.
    ClumioType                *string          `json:"type"`
    // The token given to the Clumio Cloud Connectors to identify the vCenter server.
    VcenterToken              *string          `json:"vcenter_token"`
}

// ReadVmResponse represents a custom type struct for Success
type ReadVmResponse struct {
    // Embedded responses related to the resource.
    Embedded              *VmEmbedded                   `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                 *VmLinks                      `json:"_links"`
    // The policy compliance status of the resource. If the VM is deleted or unprotected, then this field has a value of `null`. Refer to the Compliance Status table for a complete list of compliance statuses.
    ComplianceStatus      *string                       `json:"compliance_status"`
    // The compute resource from which the VM draws. If the VM is deleted, then `compute_resource.id` has a value of `null`.
    ComputeResource       *VMComputeResourceModel       `json:"compute_resource"`
    // The compute resource folder associated with this VM. If the VM is deleted, then this field has a value of `null`.
    ComputeResourceFolder *VMComputeResourceFolderModel `json:"compute_resource_folder"`
    // The data center in which the VM resides. If the VM is deleted, then `datacenter.id` has a value of `null`.
    Datacenter            *VMDatacenterModel            `json:"datacenter"`
    // The data center folder associated with this VM. If the VM is deleted, then this field has a value of `null`.
    DatacenterFolder      *VMDatacenterFolderModel      `json:"datacenter_folder"`
    // The host on which the VM resides. If the VM is deleted, then `host.id` and `host.is_standalone` have values of `null`. The `host.name` field may also have a value of `null`.
    Host                  *VMHostModel                  `json:"host"`
    // The VMware-assigned Managed Object Reference (MoRef) ID of the VM.
    // VMs from different vCenters may have the same VMware-assigned ID.
    Id                    *string                       `json:"id"`
    // Determines whether the VM is deleted. If `true`, the VM is deleted.
    IsDeleted             *bool                         `json:"is_deleted"`
    // Determines whether the VM is supported for backups. If `true`, the VM is supported for backups.
    IsSupported           *bool                         `json:"is_supported"`
    // The timestamp of when the VM was last
    // backed up. If this VM has not been backed up, then this field has a value of `null`.
    LastBackupTimestamp   *string                       `json:"last_backup_timestamp"`
    // The name of the virtual machine.
    Name                  *string                       `json:"name"`
    // The network interface card (NIC) attached to the VM.
    Nics                  []*VMNicModel                 `json:"nics"`
    // The Clumio-assigned ID of the organizational unit associated with the VM.
    OrganizationalUnitId  *string                       `json:"organizational_unit_id"`
    // The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
    ProtectionInfo        *ProtectionInfo               `json:"protection_info"`
    // The protection status of the resource. If the VM has been deleted, then this field has a value of `null`. Refer to the Protection Status table for a complete list of protection statuses.
    ProtectionStatus      *string                       `json:"protection_status"`
    // The resource pool from which the VM draws. If the VM is deleted, then `resource_pool.id` and `resource_pool.is_root` have values of `null`.
    ResourcePool          *VMResourcePoolModel          `json:"resource_pool"`
    // VMTagWithCategoryModel
    // A tag associated with the VM.
    Tags                  []*VMTagWithCategoryModel     `json:"tags"`
    // The reason why the VM cannot be supported. If the VM is supported, then this field has a value of `null`.
    UnsupportedReason     *string                       `json:"unsupported_reason"`
    // The Clumio-assigned ID of the VM.
    // Use this parameter in the [GET /backups/files/search](#operation/list-files) endpoint
    // to search for files in backups of this VM.
    Uuid                  *string                       `json:"uuid"`
    // The VM folder containing the VM. If the VM is deleted, then `vm_folder.id` has a value of `null`.
    VmFolder              *VMFolderModel                `json:"vm_folder"`
}

// RestoreFileResponse represents a custom type struct for Success
type RestoreFileResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasLinks `json:"_embedded"`
    // Embedded responses related to the resource.
    Links    *ReadTaskHateoasLinks `json:"_links"`
    // The Clumio-assigned ID of the restored file.
    Id       *string               `json:"id"`
    // passcode that the end-user must use to access the restored
    // file, in the case the restored file was emailed to the end-user as part
    // of transparent data access.
    Passcode *string               `json:"passcode"`
}

// RestoreVMwareVMResponse represents a custom type struct for Success
type RestoreVMwareVMResponse struct {
    // Embedded responses related to the resource.
    Embedded *ReadTaskHateoasLinks `json:"_embedded"`
    // Embedded responses related to the resource.
    Links    *ReadTaskHateoasLinks `json:"_links"`
    // The Clumio-assigned ID of the task created by this restore request. The progress of the task can be monitored using the `GET /tasks/{task_id}` endpoint.
    TaskId   *string               `json:"task_id"`
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

// UpdateAWSConnectionResponse represents a custom type struct for Success
type UpdateAWSConnectionResponse struct {
    // URLs to pages related to the resource.
    Links                    *AWSConnectionLinks `json:"_links"`
    // The alias given to the account on AWS.
    AccountName              *string             `json:"account_name"`
    // The AWS-assigned ID of the account associated with the connection.
    AccountNativeId          *string             `json:"account_native_id"`
    // The AWS region associated with the connection. For example, `us-east-1`.
    AwsRegion                *string             `json:"aws_region"`
    // Clumio AWS AccountId
    ClumioAwsAccountId       *string             `json:"clumio_aws_account_id"`
    // Clumio AWS Region
    ClumioAwsRegion          *string             `json:"clumio_aws_region"`
    // The status of the connection. Possible values include "connecting",
    // "connected", and "unlinked".
    ConnectionStatus         *string             `json:"connection_status"`
    // The timestamp of when the connection was created.
    CreatedTimestamp         *string             `json:"created_timestamp"`
    // An optional, user-provided description for this connection.
    Description              *string             `json:"description"`
    // The configuration of the Clumio Discover product for this connection.
    // If this connection is not configured for Clumio Discover, then this field has a
    // value of `null`.
    Discover                 *DiscoverConfig     `json:"discover"`
    // The Clumio-assigned ID of the connection.
    Id                       *string             `json:"id"`
    // Lambda Bucket Prefix required for configuring Clumio Protect Warm-tier AWS Lambda.
    LambdaBucketPrefix       *string             `json:"lambda_bucket_prefix"`
    // K8S Namespace
    Namespace                *string             `json:"namespace"`
    // The Clumio-assigned ID of the organizational unit associated with the
    // AWS environment. If this parameter is not provided, then the value
    // defaults to the first organizational unit assigned to the requesting
    // user. For more information about organizational units, refer to the
    // Organizational-Units documentation.
    OrganizationalUnitId     *string             `json:"organizational_unit_id"`
    // The configuration of the Clumio Cloud Protect product for this connection.
    // If this connection is not configured for Clumio Cloud Protect, then this field has a
    // value of `null`.
    Protect                  *ProtectConfig      `json:"protect"`
    // The asset types enabled for protect. This is only populated if "protect"
    // is enabled. Valid values are any of ["EBS", "RDS", "DynamoDB", "EC2MSSQL"].
    // EBS and RDS are mandatory datasources.
    ProtectAssetTypesEnabled []*string           `json:"protect_asset_types_enabled"`
    // The services to be enabled for this configuration. Valid values are
    // ["discover"], ["discover", "protect"]. This is only set when the
    // registration is created, the enabled services are obtained directly from
    // the installed template after that.
    ServicesEnabled          []*string           `json:"services_enabled"`
    // The Amazon Resource Name of the installed CloudFormation stack in this AWS account
    StackArn                 *string             `json:"stack_arn"`
    // The name given to the installed CloudFormation stack on AWS.
    StackName                *string             `json:"stack_name"`
    // Subnets Lambda Object Key required for configuring Clumio Protect Warm-tier AWS Lambda.
    SubnetsLambdaObjectKey   *string             `json:"subnets_lambda_object_key"`
    // The 36-character Clumio AWS integration ID token used to identify the
    // installation of the CloudFormation template on the account. This value
    // will be pasted into the ClumioToken field when creating the
    // CloudFormation stack.
    Token                    *string             `json:"token"`
}

// UpdateAlertResponse represents a custom type struct for Success
type UpdateAlertResponse struct {
    // Embedded responses related to the resource.
    Embedded            *AlertEmbedded          `json:"_embedded"`
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
    // The primary object associated with or affected by the alert. Examples of primary entities include "aws_connection", "aws_ebs_volume" and "vmware_vm".
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

// UpdateConsolidatedAlertResponse represents a custom type struct for Success
type UpdateConsolidatedAlertResponse struct {
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
    Embedded                      *PolicyEmbedded    `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                         *PolicyLinks       `json:"_links"`
    // The status of the policy.
    // Refer to the Policy Activation Status table for a complete list of policy statuses.
    ActivationStatus              *string            `json:"activation_status"`
    // The Clumio-assigned IDs of the organizational units to whom the policy has been assigned.
    AssignedOrganizationalUnitIds []*string          `json:"assigned_organizational_unit_ids"`
    // The Clumio-assigned ID of the policy.
    Id                            *string            `json:"id"`
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
    LockStatus                    *string            `json:"lock_status"`
    // The user-provided name of the policy.
    Name                          *string            `json:"name"`
    // The SLAs of an individual operation.
    Operations                    []*PolicyOperation `json:"operations"`
    // The Clumio-assigned ID of the organizational unit associated with the policy.
    OrganizationalUnitId          *string            `json:"organizational_unit_id"`
}

// UpdateSubgroupResponse represents a custom type struct for Success
type UpdateSubgroupResponse struct {
    // URLs to pages related to the resource.
    Links                       *SubgroupLinks               `json:"_links"`
    // The number of cloud connectors in this subgroup, aggregated by their status.
    CloudConnectorCountByStatus *CloudConnectorCountByStatus `json:"cloud_connector_count_by_status"`
    // The overall health of cloud connectors in this subgroup. Possible values include: 'healthy', indicating
    // that all cloud connectors in the subgroup are connected; 'degraded' indicating that one or more cloud
    // connectors in the subgroup have connection issues; `none`, indicating that no cloud connectors are in the subgroup.
    CloudConnectorStatus        *string                      `json:"cloud_connector_status"`
    // The Clumio-assigned ID of the management group associated with this subgroup.
    GroupId                     *string                      `json:"group_id"`
    // The Clumio-assigned ID of the management subgroup.
    Id                          *string                      `json:"id"`
    // The name of the management subgroup.
    Name                        *string                      `json:"name"`
}

// UpdateTaskResponse represents a custom type struct for Success
type UpdateTaskResponse struct {
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
    // Tasks of certain types including
    // "vmware_vm_backup_indexing" and "aws_ebs_volume_backup_indexing" cannot be aborted.
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
    // The task type. Examples of task types include "vm_backup_seeding", "ebs_indexing", and "file_restore".
    // Refer to the Task Type table for a complete list of task types.
    ClumioType         *string            `json:"type"`
}

// UpdateUserResponse represents a custom type struct for Success
type UpdateUserResponse struct {
    // Embedded responses related to the resource.
    Embedded                      *UserEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                         *UserLinks    `json:"_links"`
    // The list of organizational unit IDs assigned to the user.
    // This attribute will be available when reading a single user and not when listing users.
    AssignedOrganizationalUnitIds []*string     `json:"assigned_organizational_unit_ids"`
    // Assigned Role for the user.
    AssignedRole                  *string       `json:"assigned_role"`
    // The email address of the Clumio user.
    Email                         *string       `json:"email"`
    // The first and last name of the Clumio user. The name appears in the User Management screen and is used to identify the user.
    FullName                      *string       `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                            *string       `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                       *string       `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                   *bool         `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and the invitation
    // is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                     *bool         `json:"is_enabled"`
    // The timestamp of when when the user was last active in the Clumio system. Represented in RFC-3339 format.
    LastActivityTimestamp         *string       `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount       *int64        `json:"organizational_unit_count"`
}
