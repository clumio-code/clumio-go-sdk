// Copyright (c) 2021 Clumio All Rights Reserved

// Package models has the structs representing definitions
package models

// AWSConnection represents a custom type struct
type AWSConnection struct {
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
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS",
    // "IcebergOnGlue", "IcebergOnS3Tables", "FSX"].
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

// AWSConnectionGroupListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type AWSConnectionGroupListEmbedded struct {
    // TODO: Add struct field description
    Items []*ConnectionGroupWithETag `json:"items"`
}

// AWSConnectionLinks represents a custom type struct.
// URLs to pages related to the resource.
type AWSConnectionLinks struct {
    // The HATEOAS link to this resource.
    Self                   *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    DeleteConnectionAws    *HateoasLink     `json:"delete-connection-aws"`
    // A resource-specific HATEOAS link.
    ReadOrganizationalUnit *HateoasLink     `json:"read-organizational-unit"`
    // A resource-specific HATEOAS link.
    UpdateConnectionAws    *HateoasLink     `json:"update-connection-aws"`
}

// AWSConnectionListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type AWSConnectionListEmbedded struct {
    // TODO: Add struct field description
    Items []*AWSConnection `json:"items"`
}

// AWSConnectionListLinks represents a custom type struct.
// URLs to pages related to the resource.
type AWSConnectionListLinks struct {
    // The HATEOAS link to the first page of results.
    First               *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next                *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self                *HateoasSelfLink  `json:"_self"`
    // A resource-specific HATEOAS link.
    CreateAwsConnection *HateoasLink      `json:"create-aws-connection"`
}

// AWSEnvironment represents a custom type struct
type AWSEnvironment struct {
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
    // Type of permissions in the template deployed for the environment
    TemplatePermissionSet      *string                 `json:"template_permission_set"`
    // The Clumio CloudFormation template version used to deploy the CloudFormation stack.
    TemplateVersion            *int64                  `json:"template_version"`
}

// AWSEnvironmentEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type AWSEnvironmentEmbedded struct {
    // Backup statistics for each AWS environment.
    ReadAwsEnvironmentsBackupStatusStats interface{} `json:"read-aws-environments-backup-status-stats"`
}

// AWSEnvironmentLinks represents a custom type struct.
// URLs to pages related to the resource.
type AWSEnvironmentLinks struct {
    // The HATEOAS link to this resource.
    Self                   *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    ReadOrganizationalUnit *HateoasLink     `json:"read-organizational-unit"`
}

// AWSEnvironmentListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type AWSEnvironmentListEmbedded struct {
    // TODO: Add struct field description
    Items []*AWSEnvironment `json:"items"`
}

// AWSEnvironmentListLinks represents a custom type struct.
// URLs to pages related to the resource.
type AWSEnvironmentListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// AlertEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type AlertEmbedded struct {
    // Embeds the associated consolidated alert in the response.
    ReadConsolidatedAlert interface{} `json:"read-consolidated-alert"`
}

// AlertLinks represents a custom type struct.
// URLs to pages related to the resource.
type AlertLinks struct {
    // The HATEOAS link to this resource.
    Self                  *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    ReadConsolidatedAlert *HateoasLink     `json:"read-consolidated-alert"`
    // A resource-specific HATEOAS link.
    UpdateIndividualAlert *HateoasLink     `json:"update-individual-alert"`
}

// AlertListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type AlertListEmbedded struct {
    // TODO: Add struct field description
    Items []*ConsolidatedAlertWithETag `json:"items"`
}

// AlertListLinks represents a custom type struct.
// URLs to pages related to the resource.
type AlertListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// AlertParentEntity represents a custom type struct.
// The parent object of the primary entity associated with or affected by the alert. For example, "aws_environment" is the parent entity of primary entity "aws_ebs_volume".
type AlertParentEntity struct {
    // A system-generated ID assigned to this entity.
    Id         *string `json:"id"`
    // Type is mostly an asset type or the type of Entity. Some examples are
    // "restored_file", "aws_ebs_volume",  etc.
    ClumioType *string `json:"type"`
    // A system-generated value assigned to the entity. For example, if the primary entity type is "aws_ebs_volume", then the value is the name of the EBS.
    Value      *string `json:"value"`
}

// AlertPrimaryEntity represents a custom type struct.
// The primary object associated with or affected by the alert. Examples of primary entities include "aws_connection", "aws_ebs_volume".
type AlertPrimaryEntity struct {
    // A system-generated ID assigned to this entity.
    Id         *string `json:"id"`
    // Type is mostly an asset type or the type of Entity. Some examples are
    // "restored_file", "aws_ebs_volume",  etc.
    ClumioType *string `json:"type"`
    // A system-generated value assigned to the entity. For example, if the primary entity type is "aws_ebs_volume", then the value is the name of the EBS.
    Value      *string `json:"value"`
}

// AmiModel represents a custom type struct.
// An Amazon Machine Image is a supported and maintained image provided by AWS
// that provides the information required to launch an instance.
type AmiModel struct {
    // The AWS-assigned ID of the AMI.
    AmiNativeId              *string `json:"ami_native_id"`
    // The architecture of the AMI. Possible values include 'i386', 'x86_64', and 'arm64'.
    Architecture             *string `json:"architecture"`
    // Specifies whether enhanced networking with ENA is enabled.
    HasEnaSupport            *bool   `json:"has_ena_support"`
    // The hypervisor type of the AMI. Possible values include 'ovm' and 'xen'.
    HypervisorType           *string `json:"hypervisor_type"`
    // Type of Image (machine | kernel | ramdisk ).
    ImageType                *string `json:"image_type"`
    // A Boolean that indicates whether the image is public.
    IsPublic                 *bool   `json:"is_public"`
    // The name of the AMI.
    Name                     *string `json:"name"`
    // Number of ebs volumes.
    NumberOfEbsVolumes       *int64  `json:"number_of_ebs_volumes"`
    // Number of ephemeral volumes.
    NumberOfEphemeralVolumes *int64  `json:"number_of_ephemeral_volumes"`
    // The ID of the Amazon Web Services account that owns the image.
    OwnerId                  *string `json:"owner_id"`
    // The platform of the AMI. Possible values include "windows" and "linux".
    Platform                 *string `json:"platform"`
    // The name of the root device used by the AMI.
    RootDeviceName           *string `json:"root_device_name"`
    // The type of root device used by the AMI.
    RootDeviceType           *string `json:"root_device_type"`
    // A value of simple indicates that enhanced networking with the Intel 82599 VF
    // interface is enabled.
    SriovNetSupport          *string `json:"sriov_net_support"`
    // The type of virtualization of the AMI. Possible values include 'hvm' and 'paravirtual.'
    VirtualizationType       *string `json:"virtualization_type"`
}

// AssetBackupControl represents a custom type struct.
// The control evaluating whether assets have at least one backup within each window of the specified look back period,
// with retention meeting the minimum required duration.
// For example, a look_back_period of 7 days, window_size of 1 day, and retention_duration of 1 month means that
// there should be a backup every day for the past week and that the retention of that backup should be at least 1 month.
type AssetBackupControl struct {
    // The duration prior to the compliance evaluation point to look back.
    LookBackPeriod           *TimeUnitParamLookbackPeriod                  `json:"look_back_period"`
    // The minimum required retention duration for a backup to be considered compliant.
    MinimumRetentionDuration *TimeUnitParamAssetBackupMinRetentionDuration `json:"minimum_retention_duration"`
    // The size of each evaluation window within the look back period in which at least one compliant backup must exist.
    WindowSize               *TimeUnitParamWindowSize                      `json:"window_size"`
}

// AssetFilter represents a custom type struct.
// The filter for asset. This will be applied to asset backup and asset protection controls.
type AssetFilter struct {
    // The asset groups to be filtered.
    Groups    []*AssetGroupFilter `json:"groups"`
    // The tag filter operation to be applied to the given tags. This is supported for AWS assets only.
    TagOpMode *string             `json:"tag_op_mode"`
    // The asset tags to be filtered. This is supported for AWS assets only.
    Tags      []*Tag              `json:"tags"`
}

// AssetGroupFilter represents a custom type struct.
// The asset groups to be filtered.
type AssetGroupFilter struct {
    // The id of asset group.
    Id         *string `json:"id"`
    // The region of asset group. For example, `us-west-2`.
    // This is supported for AWS asset groups only.
    Region     *string `json:"region"`
    // The type of asset group.
    ClumioType *string `json:"type"`
}

// AssetProtectionControl represents a custom type struct.
// The control evaluating if all assets are protected with a policy or not.
type AssetProtectionControl struct {
    // Treat deactivated policies as compliant if true.
    ShouldIgnoreDeactivatedPolicy *bool `json:"should_ignore_deactivated_policy"`
}

// AssignPolicyAction represents a custom type struct.
// Apply a policy to assets.
type AssignPolicyAction struct {
    // The policy to be applied to the assets.
    PolicyId *string `json:"policy_id"`
}

// AssignmentEntity represents a custom type struct.
// An entity being assigned or unassigned a policy.
type AssignmentEntity struct {
    // A system-generated ID assigned of an entity being assigned or unassigned to a policy.
    Id         *string `json:"id"`
    // 
    // The type of an entity being associated or disassociated with a policy.
    // Valid primary entity types include the following:
    // 
    // +------------------------+-------------------------+
    // |  Primary Entity Type   |         Details         |
    // +========================+=========================+
    // | aws_ebs_volume         | AWS EBS volume.         |
    // +------------------------+-------------------------+
    // | aws_ec2_instance       | AWS EC2 instance.       |
    // +------------------------+-------------------------+
    // | aws_rds_cluster        | AWS RDS cluster.        |
    // +------------------------+-------------------------+
    // | aws_rds_instance       | AWS RDS instance.       |
    // +------------------------+-------------------------+
    // | aws_dynamodb_table     | AWS DynamoDB table.     |
    // +------------------------+-------------------------+
    // | protection_group       | Protection group.       |
    // +------------------------+-------------------------+
    // | aws_iceberg_glue_table | AWS Iceberg Glue table. |
    // +------------------------+-------------------------+
    // | aws_iceberg_s3_table   | AWS Iceberg S3 table.   |
    // +------------------------+-------------------------+
    // 
    ClumioType *string `json:"type"`
}

// AssignmentInputModel represents a custom type struct.
// The portion of the policy assignment used for updates/creations
type AssignmentInputModel struct {
    // The action to be executed by this request.
    // Possible values include "assign" and "unassign".
    Action   *string           `json:"action"`
    // An entity being assigned or unassigned a policy.
    Entity   *AssignmentEntity `json:"entity"`
    // The Clumio-assigned ID of the policy to be applied to the requested entities.
    // If `action: assign`, then this parameter is required.
    // Otherwise, it must not be provided.
    PolicyId *string           `json:"policy_id"`
}

// AttachedEBSVolumeFullModel represents a custom type struct
type AttachedEBSVolumeFullModel struct {
    // The Clumio-assigned id for the volume.
    Id                  *string              `json:"id"`
    // Determines whether this device is the root for the instance.
    IsRoot              *bool                `json:"is_root"`
    // The AWS-assigned id of the KMS key used for encryption of the volume.
    KmsKeyNativeId      *string              `json:"kms_key_native_id"`
    // The device name for the EBS volume. For example, '/dev/xvda'.
    Name                *string              `json:"name"`
    // The size of the EBS volume. Measured in bytes (B).
    Size                *int64               `json:"size"`
    // The status of the EBS volume. Possible values include 'attaching', 'attached',
    // 'detaching', 'detached'.
    Status              *string              `json:"status"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags                []*AwsTagCommonModel `json:"tags"`
    // The type of the volume. Possible values include 'gp2', 'io1', 'st1', 'sc1', and 'standard'.
    ClumioType          *string              `json:"type"`
    // The number of bytes written in the EBS volume.
    UtilizedSizeInBytes *uint64              `json:"utilized_size_in_bytes"`
    // The AWS-assigned ID of the EBS volume.
    VolumeNativeId      *string              `json:"volume_native_id"`
}

// AttributeDefinition represents a custom type struct.
// Represents the attributes within a DynamoDB table by the name
// and the data type (`S` for string, `N` for number, `B` for binary).
type AttributeDefinition struct {
    // TODO: Add struct field description
    Name       *string `json:"name"`
    // TODO: Add struct field description
    ClumioType *string `json:"type"`
}

// AuditTrails represents a custom type struct
type AuditTrails struct {
    // The action performed by the user.
    // 
    // +-------------------------+----------------------------------------------------+
    // |         Action          |                      Details                       |
    // +=========================+====================================================+
    // | create                  | Creating or adding new entities like new policy,   |
    // |                         | configuration, user, etc                           |
    // +-------------------------+----------------------------------------------------+
    // | update                  | Updating an existing entity like policy, settings, |
    // |                         | passwords, etc                                     |
    // +-------------------------+----------------------------------------------------+
    // | delete                  | Delete an existing entity like policy, settings,   |
    // |                         | users, etc                                         |
    // +-------------------------+----------------------------------------------------+
    // | enable                  | Enabling a feature like single sign on or multi    |
    // |                         | factor authentication settings                     |
    // +-------------------------+----------------------------------------------------+
    // | disable                 | Disabling features like single sign on or multi    |
    // |                         | factor authentication settings                     |
    // +-------------------------+----------------------------------------------------+
    // | browse                  | Browsing through entities in the system like       |
    // |                         | mailboxes or backups, etc                          |
    // +-------------------------+----------------------------------------------------+
    // | search                  | Searching through entities in the system like      |
    // |                         | mailboxes or backups, etc                          |
    // +-------------------------+----------------------------------------------------+
    // | login                   | User logs in or tries to login                     |
    // +-------------------------+----------------------------------------------------+
    // | logout                  | User explicitly logged out.                        |
    // +-------------------------+----------------------------------------------------+
    // | register                | When new registrations happen like new             |
    // |                         | datasource registration or user registering for    |
    // |                         | MFA                                                |
    // +-------------------------+----------------------------------------------------+
    // | unregister              | When unregistering like unregistering              |
    // |                         | datasource or user unregistering MFA               |
    // +-------------------------+----------------------------------------------------+
    // | apply                   | Apply policy to protect entities, tags, etc        |
    // +-------------------------+----------------------------------------------------+
    // | remove                  | Remove protection for entities, tags, etc          |
    // +-------------------------+----------------------------------------------------+
    // | invite                  | Inviting a user                                    |
    // +-------------------------+----------------------------------------------------+
    // | suspend                 | Suspend an existing user                           |
    // +-------------------------+----------------------------------------------------+
    // | full_restore            | Full restore of the VM, volume, mailbox, database  |
    // |                         | or other entities                                  |
    // +-------------------------+----------------------------------------------------+
    // | granular_retrieval      | Restoring individual files, mails or records       |
    // +-------------------------+----------------------------------------------------+
    // | redirected              | When cross region restore occurs.                  |
    // +-------------------------+----------------------------------------------------+
    // | unapply                 | Assets removed from a rule.                        |
    // +-------------------------+----------------------------------------------------+
    // | batch_activate          | Activate multiple policies.                        |
    // +-------------------------+----------------------------------------------------+
    // | batch_deactivate        | Deactivate multiple policies.                      |
    // +-------------------------+----------------------------------------------------+
    // | grant_email_access      | Grant email access for a file level object. This   |
    // |                         | is mutually exclusive with grant_download_access   |
    // +-------------------------+----------------------------------------------------+
    // | download                | File was download.                                 |
    // +-------------------------+----------------------------------------------------+
    // | validate_tda_passcode   | Validate passcode that is entered for a download.  |
    // +-------------------------+----------------------------------------------------+
    // | regenerate_tda_passcode | Regenerate a new passcode used for download.       |
    // +-------------------------+----------------------------------------------------+
    // 
    Action          *string             `json:"action"`
    // The category of the auditable action performed by the user.
    // 
    // +-------------------------+----------------------------------------------------+
    // |        Category         |                      Details                       |
    // +=========================+====================================================+
    // | authentication          | Activities related to Authentication               |
    // +-------------------------+----------------------------------------------------+
    // | data_source             | Data source changes                                |
    // +-------------------------+----------------------------------------------------+
    // | policy                  | Policy related actions                             |
    // +-------------------------+----------------------------------------------------+
    // | protection              | Applying and removing protection                   |
    // +-------------------------+----------------------------------------------------+
    // | restore                 | Restore related operations                         |
    // +-------------------------+----------------------------------------------------+
    // | tasks                   | Tasks                                              |
    // +-------------------------+----------------------------------------------------+
    // | backup                  | Backup related operations                          |
    // +-------------------------+----------------------------------------------------+
    // | users                   | User related operations                            |
    // +-------------------------+----------------------------------------------------+
    // | api_tokens              | API Token related operations like creating,        |
    // |                         | revoking or deleting tokens                        |
    // +-------------------------+----------------------------------------------------+
    // | kms_config              | Key Management Service(KMS) related operations     |
    // +-------------------------+----------------------------------------------------+
    // | sso                     | Single sign-on (SSO) related operations            |
    // +-------------------------+----------------------------------------------------+
    // | mfa                     | Multi Factor Authentication(MFA) related           |
    // |                         | operations                                         |
    // +-------------------------+----------------------------------------------------+
    // | reports                 | Reports related operations                         |
    // +-------------------------+----------------------------------------------------+
    // | alerts                  | Alerts related operations                          |
    // +-------------------------+----------------------------------------------------+
    // | cloud_connector         | Cloud connector related operations                 |
    // +-------------------------+----------------------------------------------------+
    // | cloudformation_template | Cloud Formation Template related operations        |
    // +-------------------------+----------------------------------------------------+
    // | bandwidth_config        | Bandwidth configuration related changes            |
    // +-------------------------+----------------------------------------------------+
    // | partner_ecosystem       | Changes to partner ecosystem                       |
    // +-------------------------+----------------------------------------------------+
    // | ecosystem_changes       | Changes in the ecosystem like adding or removing   |
    // |                         | VMs                                                |
    // +-------------------------+----------------------------------------------------+
    // | organizational_unit     | Changes in the Organizational Unit/Entity group    |
    // |                         | such as creation, deletion, patch.                 |
    // +-------------------------+----------------------------------------------------+
    // 
    Category        *string             `json:"category"`
    // Additional details about the activity provided in JSON format.
    // Deprecated
    Details         *string             `json:"details"`
    // TODO: Add struct field description
    DetailsJson     *Details            `json:"details_json"`
    // The Clumio-assigned ID of the audit event.
    Id              *string             `json:"id"`
    // The interface used to make the request i.e. 'UI','API'
    ClumioInterface *string             `json:"interface"`
    // The IP address from which the activity was requested.
    IpAddress       *string             `json:"ip_address"`
    // The parent object of the primary entity associated with or affected by the audit.
    ParentEntity    *AuditParentEntity  `json:"parent_entity"`
    // The primary object associated with the audit event. Examples of primary entities
    // include "aws_connection", "aws_ebs_volume" and "aws_ec2_instance". In some cases like
    // global settings, the primary entity may be null.
    PrimaryEntity   *AuditPrimaryEntity `json:"primary_entity"`
    // The status of the performed action. 'success', 'failure', 'partial_success'
    Status          *string             `json:"status"`
    // The Timestamp of when the activity began. Represented in RFC-3339 format.
    Timestamp       *string             `json:"timestamp"`
    // The email address of the logged in user making the request.
    UserEmail       *string             `json:"user_email"`
}

// AuditParentEntity represents a custom type struct.
// The parent object of the primary entity associated with or affected by the audit.
type AuditParentEntity struct {
    // A system-generated ID assigned to this entity.
    Id         *string `json:"id"`
    // Type is mostly an asset type or the type of Entity. Some examples are
    // "restored_file", "aws_ebs_volume",  etc.
    ClumioType *string `json:"type"`
    // A system-generated value assigned to the entity. For example, if the primary entity type is "aws_ebs_volume", then the value is the name of the EBS.
    Value      *string `json:"value"`
}

// AuditPrimaryEntity represents a custom type struct.
// The primary object associated with the audit event. Examples of primary entities
// include "aws_connection", "aws_ebs_volume" and "aws_ec2_instance". In some cases like
// global settings, the primary entity may be null.
type AuditPrimaryEntity struct {
    // A system-generated ID assigned to this entity.
    Id         *string `json:"id"`
    // Type is mostly an asset type or the type of Entity. Some examples are
    // "restored_file", "aws_ebs_volume",  etc.
    ClumioType *string `json:"type"`
    // A system-generated value assigned to the entity. For example, if the primary entity type is "aws_ebs_volume", then the value is the name of the EBS.
    Value      *string `json:"value"`
}

// AuditTrailListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type AuditTrailListEmbedded struct {
    // TODO: Add struct field description
    Items []*AuditTrails `json:"items"`
}

// AuditTrailListLinks represents a custom type struct.
// URLs to pages related to the resource.
type AuditTrailListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// AutoUserProvisioningRuleEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type AutoUserProvisioningRuleEmbedded struct {
    // Embeds the associated organizational units for the OU UUIDs in the response
    // if requested using the `embed` query parameter.
    ReadOrganizationalUnit interface{} `json:"read-organizational-unit"`
    // Embeds the associated role for the role UUID in the response if requested using the `embed` query parameter.
    ReadRole               interface{} `json:"read-role"`
}

// AutoUserProvisioningRuleLinks represents a custom type struct.
// URLs to pages related to the resource.
type AutoUserProvisioningRuleLinks struct {
    // The HATEOAS link to this resource.
    Self                           *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    DeleteAutoUserProvisioningRule *HateoasLink     `json:"delete-auto-user-provisioning-rule"`
    // A resource-specific HATEOAS link.
    UpdateAutoUserProvisioningRule *HateoasLink     `json:"update-auto-user-provisioning-rule"`
}

// AutoUserProvisioningRuleListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type AutoUserProvisioningRuleListEmbedded struct {
    // AutoUserProvisioningRuleWithETag to support etag string to be calculated
    Items []*AutoUserProvisioningRuleWithETag `json:"items"`
}

// AutoUserProvisioningRuleListLinks represents a custom type struct.
// URLs to pages related to the resource.
type AutoUserProvisioningRuleListLinks struct {
    // The HATEOAS link to the first page of results.
    First                          *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next                           *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self                           *HateoasSelfLink  `json:"_self"`
    // A resource-specific HATEOAS link.
    CreateAutoUserProvisioningRule *HateoasLink      `json:"create-auto-user-provisioning-rule"`
}

// AutoUserProvisioningRuleWithETag represents a custom type struct.
// AutoUserProvisioningRuleWithETag to support etag string to be calculated
type AutoUserProvisioningRuleWithETag struct {
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

// AutoUserProvisioningSettingLinks represents a custom type struct.
// URLs to pages related to the resource.
type AutoUserProvisioningSettingLinks struct {
    // The HATEOAS link to this resource.
    Self                              *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    UpdateAutoUserProvisioningSetting *HateoasLink     `json:"update-auto-user-provisioning-setting"`
}

// AwsDsGroupingCriteria represents a custom type struct.
// The entity type used to group organizational units for AWS resources.
type AwsDsGroupingCriteria struct {
    // Determines whether or not this data group is editable. If false, then an
    // organizational unit uses this data group.
    // To edit this data group, all organizational units using it must be deleted.
    IsEditable *bool   `json:"is_editable"`
    // 
    // +-----------------+-------------------------+
    // |   Entity Type   |         Details         |
    // +=================+=========================+
    // | aws_environment | AWS account and region. |
    // +-----------------+-------------------------+
    // 
    ClumioType *string `json:"type"`
}

// AwsTag represents a custom type struct
type AwsTag struct {
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

// AwsTagCommonModel represents a custom type struct.
// A tag created through AWS Console which can be applied to EBS volumes.
type AwsTagCommonModel struct {
    // The AWS-assigned tag key.
    Key   *string `json:"key"`
    // The AWS-assigned tag value.
    Value *string `json:"value"`
}

// AwsTagEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type AwsTagEmbedded struct {
    // Embedded AWS Backup statistics for each tag.
    ReadAwsEnvironmentTagBackupStatusStats               interface{} `json:"read-aws-environment-tag-backup-status-stats"`
    // Embedded AWS DynamoDB statistics for each tag.
    ReadAwsEnvironmentTagDynamodbTablesProtectionStats   interface{} `json:"read-aws-environment-tag-dynamodb-tables-protection-stats"`
    // Embedded AWS EBS statistics for each tag.
    ReadAwsEnvironmentTagEbsVolumesProtectionStats       interface{} `json:"read-aws-environment-tag-ebs-volumes-protection-stats"`
    // Embedded AWS EC2 statistics for each tag.
    ReadAwsEnvironmentTagEc2InstancesProtectionStats     interface{} `json:"read-aws-environment-tag-ec2-instances-protection-stats"`
    // Embedded Protection Group statistics for each tag.
    ReadAwsEnvironmentTagProtectionGroupsProtectionStats interface{} `json:"read-aws-environment-tag-protection-groups-protection-stats"`
    // Embedded AWS RDS statistics for each tag.
    ReadAwsEnvironmentTagRdsResourcesProtectionStats     interface{} `json:"read-aws-environment-tag-rds-resources-protection-stats"`
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition                                 interface{} `json:"read-policy-definition"`
}

// AwsTagLinks represents a custom type struct.
// URLs to pages related to the resource.
type AwsTagLinks struct {
    // The HATEOAS link to this resource.
    Self                                                 *HateoasSelfLink                 `json:"_self"`
    // A HATEOAS link to protect the entities.
    ProtectEntities                                      *ProtectEntitiesHateoasLink      `json:"protect-entities"`
    // A resource-specific HATEOAS link.
    ReadAwsEnvironmentTagBackupStatusStats               *HateoasLink                     `json:"read-aws-environment-tag-backup-status-stats"`
    // A resource-specific HATEOAS link.
    ReadAwsEnvironmentTagDynamodbTablesProtectionStats   *HateoasLink                     `json:"read-aws-environment-tag-dynamodb-tables-protection-stats"`
    // A resource-specific HATEOAS link.
    ReadAwsEnvironmentTagEbsVolumesProtectionStats       *HateoasLink                     `json:"read-aws-environment-tag-ebs-volumes-protection-stats"`
    // A resource-specific HATEOAS link.
    ReadAwsEnvironmentTagEc2InstancesProtectionStats     *HateoasLink                     `json:"read-aws-environment-tag-ec2-instances-protection-stats"`
    // A resource-specific HATEOAS link.
    ReadAwsEnvironmentTagProtectionGroupsProtectionStats *HateoasLink                     `json:"read-aws-environment-tag-protection-groups-protection-stats"`
    // A resource-specific HATEOAS link.
    ReadAwsEnvironmentTagRdsResourcesProtectionStats     *HateoasLink                     `json:"read-aws-environment-tag-rds-resources-protection-stats"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition                                 *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
    // A HATEOAS link to unprotect the entities.
    UnprotectEntities                                    *UnprotectEntitiesHateoasLink    `json:"unprotect-entities"`
}

// AwsTagListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type AwsTagListEmbedded struct {
    // TODO: Add struct field description
    Items []*AwsTag `json:"items"`
}

// AwsTagListLinks represents a custom type struct.
// URLs to pages related to the resource.
type AwsTagListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// AwsTagModel represents a custom type struct.
// A tag created through AWS console which can be applied to EBS volumes.
type AwsTagModel struct {
    // The Clumio-assigned ID of the AWS tag.
    Id    *string `json:"id"`
    // The AWS-assigned tag key.
    Key   *string `json:"key"`
    // The Clumio-assigned ID of the AWS key.
    KeyId *string `json:"key_id"`
    // The AWS-assigned tag value.
    Value *string `json:"value"`
}

// BackupSLA represents a custom type struct.
// backup_sla captures the SLA parameters
// backup_sla captures the SLA parameters
type BackupSLA struct {
    // The retention time for this SLA. For example, to retain the backup for 1 month, set `unit="months"` and `value=1`.
    RetentionDuration *RetentionBackupSLAParam `json:"retention_duration"`
    // The minimum frequency between backups for this SLA. Also known as the recovery point objective (RPO) interval.
    // For example, to configure the minimum frequency between backups to be every 2 days, set `unit="days"` and `value=2`.
    // To configure the SLA for on-demand backups, set `unit="on_demand"` and leave the `value` field empty.
    RpoFrequency      *RPOBackupSLAParam       `json:"rpo_frequency"`
}

// BackupStatusInfo represents a custom type struct.
// The backup status information applied to this resource.
type BackupStatusInfo struct {
    // BackupStatus is the status of the backup. Possible values are
    // `success`, `partial_success`, `failure`, `no_backup`, and `unknown`. This value
    // depends on `lookback_days`. If not specified, then this field has a value of `unknown`.
    BackupStatus                       *string          `json:"backup_status"`
    // The last failed policy start time. Represented in RFC-3339 format.
    LastFailedPolicyStartTimestamp     *string          `json:"last_failed_policy_start_timestamp"`
    // The last successful policy start time. Represented in RFC-3339 format.
    LastSuccessfulPolicyStartTimestamp *string          `json:"last_successful_policy_start_timestamp"`
    // TODO: Add struct field description
    OperationInfoList                  []*OperationInfo `json:"operation_info_list"`
}

// BackupStatusStats represents a custom type struct.
// Represents the aggregated stats for backup status.
type BackupStatusStats struct {
    // The total number of entities that have a backup status of `failure`.
    FailureCount        *int64 `json:"failure_count"`
    // The total number of entities that have a backup status of `no_backup`.
    NoBackupCount       *int64 `json:"no_backup_count"`
    // The total number of entities that have a backup status of `partial_success`.
    PartialSuccessCount *int64 `json:"partial_success_count"`
    // The total number of entities that have a backup status of `success`.
    SuccessCount        *int64 `json:"success_count"`
}

// BackupTierStat represents a custom type struct.
// BackupTierStat
type BackupTierStat struct {
    // The backup tier name.
    BackupTier               *string `json:"backup_tier"`
    // Cumulative count of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpObjectCount *int64  `json:"total_backed_up_object_count"`
    // Cumulative size of all unexpired objects in each backup (any new or updated since
    // the last backup) that have been backed up as part of this protection group
    TotalBackedUpSizeBytes   *int64  `json:"total_backed_up_size_bytes"`
}

// BackupWindow represents a custom type struct.
// The start and end times of the customized backup window. Use of `backup_window` is deprecated, use `backup_window_tz` instead.
type BackupWindow struct {
    // The time when the backup window closes. Specify the end time in the format `hh:mm`, where `hh` represents the hour of the day and `mm` represents the minute of the day, based on a 24 hour clock. Leave empty if you do not want to specify an end time. If the backup window closes while a backup is in progress, the entire backup process is aborted. Clumio will perform the next backup when the backup window re-opens.
    EndTime   *string `json:"end_time"`
    // The time when the backup window opens. Specify the start time in the format `hh:mm`, where `hh` represents the hour of the day and `mm` represents the minute of the day based on the 24 hour clock.
    StartTime *string `json:"start_time"`
}

// Bucket represents a custom type struct
type Bucket struct {
    // Embedded responses related to the resource.
    Embedded                      interface{}                                   `json:"_embedded"`
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

// BucketLinks represents a custom type struct.
// URLs to pages related to the resource.
type BucketLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// BucketListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type BucketListEmbedded struct {
    // TODO: Add struct field description
    Items []*Bucket `json:"items"`
}

// BucketListLinks represents a custom type struct.
// URLs to pages related to the resource.
type BucketListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// CategorisedResources represents a custom type struct.
// Categorised Resources, based on the generated template, to be created manually by the user
type CategorisedResources struct {
    // TODO: Add struct field description
    Policies     map[string]*PolicyDetails             `json:"policies"`
    // Details for the IAM Role
    Roles        map[string]*ClumioRoleResource        `json:"roles"`
    // TODO: Add struct field description
    Rules        map[string]*ClumioRuleResource        `json:"rules"`
    // Details for the ssm document attached to any resource
    SsmDocuments map[string]*ClumioSsmDocumentResource `json:"ssm_documents"`
    // Details for the SNS Topic
    Topics       map[string]*ClumioTopicResource       `json:"topics"`
}

// ClumioRoleResource represents a custom type struct.
// Details for the IAM Role
type ClumioRoleResource struct {
    // TODO: Add struct field description
    Description     *string          `json:"description"`
    // TODO: Add struct field description
    InlinePolicies  []*PolicyDetails `json:"inline_policies"`
    // TODO: Add struct field description
    ManagedPolicies []*PolicyDetails `json:"managed_policies"`
    // TODO: Add struct field description
    Steps           *string          `json:"steps"`
    // "trust_policy" stores the Trust Relationship policy for the role. It is a stringified JSON blob.
    // The user has to JSONify it and then paste the JSONified blob in aws console while creating the role.
    TrustPolicy     interface{}      `json:"trust_policy"`
}

// ClumioRuleResource represents a custom type struct
type ClumioRuleResource struct {
    // "description" is optional
    Description  *string       `json:"description"`
    // "event_pattern" has stringified JSON blob. The user has to JSONify it and then paste
    // the JSONified blob in aws console while creating the rule.
    EventPattern interface{}   `json:"event_pattern"`
    // "steps" refers to commands to be executed
    Steps        *string       `json:"steps"`
    // "targets" is a string that essentially stores the target for the rule. It generally is an ARN.
    Targets      []interface{} `json:"targets"`
}

// ClumioSsmDocumentInputs represents a custom type struct
type ClumioSsmDocumentInputs struct {
    // "runCommand" is an array of stringified commands.
    Runcommand     []*string `json:"runCommand"`
    // "timeoutSeconds" is a stringified number denoting the timeout for command execution
    Timeoutseconds *string   `json:"timeoutSeconds"`
}

// ClumioSsmDocumentParameterValue represents a custom type struct.
// Details for each parameters of the ssm document
type ClumioSsmDocumentParameterValue struct {
    // "allowedPattern" refers to the pattern that must be satisfied by the parameter
    Allowedpattern *string `json:"allowedPattern"`
    // "default" refers to the default value for that parameter
    ClumioDefault  *string `json:"default"`
    // "description" is optional
    Description    *string `json:"description"`
    // "type" refers to the parameter type
    ClumioType     *string `json:"type"`
}

// ClumioSsmDocumentResource represents a custom type struct.
// Details for the ssm document attached to any resource
type ClumioSsmDocumentResource struct {
    // "description" must contain the version being followed
    Description   *string                                     `json:"description"`
    // Details for each step present inside an ssm document
    Mainsteps     []*ClumioSsmDocumentStep                    `json:"mainSteps"`
    // Details for each parameters of the ssm document
    Parameters    map[string]*ClumioSsmDocumentParameterValue `json:"parameters"`
    // "schemaVersion" is an AWS value for versioning
    Schemaversion *string                                     `json:"schemaVersion"`
}

// ClumioSsmDocumentStep represents a custom type struct.
// Details for each step present inside an ssm document
type ClumioSsmDocumentStep struct {
    // "action" refers to a unique action identified for this step
    Action       *string                  `json:"action"`
    // TODO: Add struct field description
    Inputs       *ClumioSsmDocumentInputs `json:"inputs"`
    // "name" refers to name of that step
    Name         *string                  `json:"name"`
    // "precondition" is used for targeting a OS or validating input parameters
    Precondition map[string]*[]string     `json:"precondition"`
}

// ClumioTopicResource represents a custom type struct.
// Details for the SNS Topic
type ClumioTopicResource struct {
    // TODO: Add struct field description
    Policies []*PolicyDetails `json:"policies"`
    // "steps" refers to commands to be executed
    Steps    *string          `json:"steps"`
}

// CommonFilter represents a custom type struct.
// The common filter which will be applied to all controls.
type CommonFilter struct {
    // The asset types to be included in the report.
    // For example, ["aws_ec2_instance", "microsoft365_drive"].
    AssetTypes          []*string `json:"asset_types"`
    // The data sources to be included in the report. Possible values include `aws`, `microsoft365` or `vmware`.
    DataSources         []*string `json:"data_sources"`
    // The organizational units to be included in the report.
    OrganizationalUnits []*string `json:"organizational_units"`
}

// ComplianceConfiguration represents a custom type struct
type ComplianceConfiguration struct {
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

// ComplianceConfigurationLinks represents a custom type struct.
// URLs to pages related to the resource.
type ComplianceConfigurationLinks struct {
    // The HATEOAS link to this resource.
    Self                                *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    DeleteComplianceReportConfiguration *HateoasLink     `json:"delete-compliance-report-configuration"`
    // A resource-specific HATEOAS link.
    UpdateComplianceReportConfiguration *HateoasLink     `json:"update-compliance-report-configuration"`
}

// ComplianceConfigurationListEmbedded represents a custom type struct.
// An array of embedded resources related to this resource.
type ComplianceConfigurationListEmbedded struct {
    // TODO: Add struct field description
    Items []*ComplianceConfiguration `json:"items"`
}

// ComplianceConfigurationListLinks represents a custom type struct.
// URLs to pages related to the resource.
type ComplianceConfigurationListLinks struct {
    // The HATEOAS link to the first page of results.
    First                               *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next                                *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self                                *HateoasSelfLink  `json:"_self"`
    // A resource-specific HATEOAS link.
    CreateComplianceReportConfiguration *HateoasLink      `json:"create-compliance-report-configuration"`
}

// ComplianceControls represents a custom type struct.
// Compliance controls to evaluate policy or assets for compliance.
type ComplianceControls struct {
    // The control evaluating whether assets have at least one backup within each window of the specified look back period,
    // with retention meeting the minimum required duration.
    // For example, a look_back_period of 7 days, window_size of 1 day, and retention_duration of 1 month means that
    // there should be a backup every day for the past week and that the retention of that backup should be at least 1 month.
    AssetBackup     *AssetBackupControl     `json:"asset_backup"`
    // The control evaluating if all assets are protected with a policy or not.
    AssetProtection *AssetProtectionControl `json:"asset_protection"`
    // The control evaluating if policies have a minimum backup retention and frequency.
    Policy          *PolicyControl          `json:"policy"`
}

// ComplianceFilters represents a custom type struct.
// The set of filters supported in compliance report.
type ComplianceFilters struct {
    // The filter for asset. This will be applied to asset backup and asset protection controls.
    Asset  *AssetFilter  `json:"asset"`
    // The common filter which will be applied to all controls.
    Common *CommonFilter `json:"common"`
}

// ComplianceInfo represents a custom type struct.
// The status per controls in the compliance report created by the report run.
type ComplianceInfo struct {
    // The compliance status of the report run.
    ComplianceStatus  *string        `json:"compliance_status"`
    // The count of compliant items of the report run.
    CompliantCount    *int64         `json:"compliant_count"`
    // The status per controls in the compliance report created by the report run.
    Controls          []*ControlInfo `json:"controls"`
    // The items covered in the compliance report created by the report run.
    ItemsCovered      *ItemsCovered  `json:"items_covered"`
    // The count of non-compliant items of the report run.
    NonCompliantCount *int64         `json:"non_compliant_count"`
    // The count of unknown items of the report run.
    UnknownCount      *int64         `json:"unknown_count"`
}

// ComplianceRun represents a custom type struct
type ComplianceRun struct {
    // Embedded responses related to the resource.
    Embedded           interface{}                `json:"_embedded"`
    // URLs to pages related to the resource.
    Links              *ComplianceRunHateoasLinks `json:"_links"`
    // The status per controls in the compliance report created by the report run.
    ComplianceInfo     *ComplianceInfo            `json:"compliance_info"`
    // The RFC3339 format time when the report run was created.
    Created            *string                    `json:"created"`
    // The RFC3339 format time when the report run was expired.
    Expired            *string                    `json:"expired"`
    // The unique identifier of the report run.
    Id                 *string                    `json:"id"`
    // The name of the report run.
    Name               *string                    `json:"name"`
    // Filter and control parameters of compliance report.
    Parameter          *Parameter                 `json:"parameter"`
    // The unique identifier of the report configuration from which the report run was generated.
    ReportConfigId     *string                    `json:"report_config_id"`
    // The link to download the report CSV.
    ReportDownloadLink *string                    `json:"report_download_link"`
    // The generation status of the report run.
    Status             *string                    `json:"status"`
    // The ID of the report run generation task.
    TaskId             *string                    `json:"task_id"`
    // The RFC3339 format time when the report run was updated.
    Updated            *string                    `json:"updated"`
}

// ComplianceRunHateoasLinks represents a custom type struct.
// URLs to pages related to the resource.
type ComplianceRunHateoasLinks struct {
    // The HATEOAS link to this resource.
    Self                         *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    DeleteComplianceReportRun    *HateoasLink     `json:"delete-compliance-report-run"`
    // A resource-specific HATEOAS link.
    SendComplianceReportRunEmail *HateoasLink     `json:"send-compliance-report-run-email"`
}

// ComplianceRunListHateoasEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ComplianceRunListHateoasEmbedded struct {
    // TODO: Add struct field description
    Items []*ComplianceRun `json:"items"`
}

// ComplianceRunListHateoasLinks represents a custom type struct.
// URLs to pages related to the resource.
type ComplianceRunListHateoasLinks struct {
    // The HATEOAS link to the first page of results.
    First                               *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next                                *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self                                *HateoasSelfLink  `json:"_self"`
    // A resource-specific HATEOAS link.
    CreateComplianceReportConfiguration *HateoasLink      `json:"create-compliance-report-configuration"`
}

// ConnectionGroupLinks represents a custom type struct.
// URLs to pages related to the resource.
type ConnectionGroupLinks struct {
    // The HATEOAS link to this resource.
    Self                     *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    ReadOrganizationalUnit   *HateoasLink     `json:"read-organizational-unit"`
    // A resource-specific HATEOAS link.
    UpdateAwsConnectionGroup *HateoasLink     `json:"update-aws-connection-group"`
}

// ConnectionGroupListLinks represents a custom type struct.
// URLs to pages related to the resource.
type ConnectionGroupListLinks struct {
    // The HATEOAS link to the first page of results.
    First                    *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next                     *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self                     *HateoasSelfLink  `json:"_self"`
    // A resource-specific HATEOAS link.
    CreateAwsConnection      *HateoasLink      `json:"create-aws-connection"`
    // A resource-specific HATEOAS link.
    CreateAwsConnectionGroup *HateoasLink      `json:"create-aws-connection-group"`
}

// ConnectionGroupWithETag represents a custom type struct
type ConnectionGroupWithETag struct {
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
    // Valid values are any of ["EC2/EBS", "RDS", "DynamoDB", "EC2MSSQL", "S3", "EBS",
    // "IcebergOnGlue", "IcebergOnS3Tables", "FSX"].
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
    // TODO: Add struct field description
    TemplatePermissionSet    *string               `json:"template_permission_set"`
}

// ConnectionRegion represents a custom type struct
type ConnectionRegion struct {
    // The ID of the aws region.
    Id                *string `json:"id"`
    // Boolean that notes which regions can be designated as targets
    IsDataPlaneRegion *bool   `json:"is_data_plane_region"`
    // Name is a user friendly name of the AWS region.
    Name              *string `json:"name"`
}

// ConnectionRegionListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ConnectionRegionListEmbedded struct {
    // TODO: Add struct field description
    Items []*ConnectionRegion `json:"items"`
}

// ConnectionRegionListLinks represents a custom type struct.
// URLs to pages related to the resource.
type ConnectionRegionListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// ConnectionResourcesResp represents a custom type struct
type ConnectionResourcesResp struct {
    // SNS topic created in the account to receive relevant events.
    ClumioEventPubArn    *string       `json:"clumio_event_pub_arn"`
    // ARN of the IAM role created in the account, which will be assumed by Clumio.
    ClumioIamRoleArn     *string       `json:"clumio_iam_role_arn"`
    // ARN of the support role which will be used by the Clumio support team.
    ClumioSupportRoleArn *string       `json:"clumio_support_role_arn"`
    // TODO: Add struct field description
    EventRules           *EventRules   `json:"event_rules"`
    // TODO: Add struct field description
    ServiceRoles         *ServiceRoles `json:"service_roles"`
}

// ConsolidatedAlertDetails represents a custom type struct.
// Additional information about the consolidated alert.
type ConsolidatedAlertDetails struct {
    // A brief description of the condition that caused the alert. Examples include "Size Limit Exceeded" and "Insufficient Cloud Connector Capacity".
    Cause      *string `json:"cause"`
    // The general alert category. Examples include "Policy Violated" and "Restore Failed".
    ClumioType *string `json:"type"`
}

// ConsolidatedAlertLinks represents a custom type struct.
// URLs to pages related to the resource.
type ConsolidatedAlertLinks struct {
    // The HATEOAS link to this resource.
    Self                    *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    UpdateConsolidatedAlert *HateoasLink     `json:"update-consolidated-alert"`
}

// ConsolidatedAlertListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ConsolidatedAlertListEmbedded struct {
    // TODO: Add struct field description
    Items []*ConsolidatedAlertWithETag `json:"items"`
}

// ConsolidatedAlertListLinks represents a custom type struct.
// URLs to pages related to the resource.
type ConsolidatedAlertListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// ConsolidatedAlertParentEntity represents a custom type struct.
// The entity associated with or affected by the alert.
type ConsolidatedAlertParentEntity struct {
    // A system-generated ID assigned to this entity.
    Id         *string `json:"id"`
    // Type is mostly an asset type or the type of Entity. Some examples are
    // "restored_file", "aws_ebs_volume",  etc.
    ClumioType *string `json:"type"`
    // A system-generated value assigned to the entity. For example, if the primary entity type is "aws_ebs_volume", then the value is the name of the EBS.
    Value      *string `json:"value"`
}

// ConsolidatedAlertWithETag represents a custom type struct
type ConsolidatedAlertWithETag struct {
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

// ConsolidatedConfig represents a custom type struct.
// The consolidated configuration of the Clumio Cloud Protect and Clumio Cloud Discover products for this connection.
// If this connection is deprecated to use unconsolidated configuration, then this field has a
// value of `null`.
type ConsolidatedConfig struct {
    // The asset types supported on the current version of the feature
    AssetTypesEnabled        []*string                   `json:"asset_types_enabled"`
    // DynamodbAssetInfo
    // The installed information for the DynamoDB feature.
    Dynamodb                 *DynamodbAssetInfo          `json:"dynamodb"`
    // EbsAssetInfo
    // The installed information for the EBS feature.
    Ebs                      *EbsAssetInfo               `json:"ebs"`
    // Ec2AssetInfo
    // The installed information for the EC2 feature.
    Ec2                      *Ec2AssetInfo               `json:"ec2"`
    // EC2MSSQLProtectConfig
    // The installed information for the EC2_MSSQL feature.
    Ec2Mssql                 *EC2MSSQLProtectConfig      `json:"ec2_mssql"`
    // IcebergOnGlueAssetInfo
    // The installed information for the Iceberg on AWS Glue feature.
    IcebergOnGlue            *IcebergOnGlueAssetInfo     `json:"iceberg_on_glue"`
    // IcebergOnS3TablesAssetInfo
    // The installed information for the Iceberg on S3 Tables feature.
    IcebergOnS3Tables        *IcebergOnS3TablesAssetInfo `json:"iceberg_on_s3_tables"`
    // The current version of the feature.
    InstalledTemplateVersion *string                     `json:"installed_template_version"`
    // RdsAssetInfo
    // The installed information for the RDS feature.
    Rds                      *RdsAssetInfo               `json:"rds"`
    // S3AssetInfo
    // The installed information for the S3 feature.
    S3                       *S3AssetInfo                `json:"s3"`
    // The configuration of the Clumio Cloud Warm-Tier Protect product for this connection.
    WarmTierProtect          *WarmTierProtectConfig      `json:"warm_tier_protect"`
}

// ControlInfo represents a custom type struct.
// The status per controls in the compliance report created by the report run.
type ControlInfo struct {
    // The count of compliant items of the control.
    CompliantCount    *int64  `json:"compliant_count"`
    // The compliance status of the control.
    ControlStatus     *string `json:"control_status"`
    // The name of the control.
    Name              *string `json:"name"`
    // The count of non-compliant items of the control.
    NonCompliantCount *int64  `json:"non_compliant_count"`
    // The count of unknown items of the control.
    UnknownCount      *int64  `json:"unknown_count"`
}

// CreateComplianceRunHateoasLinks represents a custom type struct.
// CreateComplianceRunHateoasLinks
// URLs to pages related to the resource.
type CreateComplianceRunHateoasLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// CreateEC2MSSQLDatabaseRestoreResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type CreateEC2MSSQLDatabaseRestoreResponseLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// CreateOnDemandEC2MSSQLDatabaseBackupResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type CreateOnDemandEC2MSSQLDatabaseBackupResponseLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// CreateRdsDatabaseRestoreResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type CreateRdsDatabaseRestoreResponseLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// CreateRestoreRecordResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type CreateRestoreRecordResponseLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// CreateRuleResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type CreateRuleResponseLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// CreateS3InstantAccessEndpointResponseEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type CreateS3InstantAccessEndpointResponseEmbedded struct {
    // Embeds the associated task of a resource in the response if requested using the `embed` query parameter.
    ReadTask interface{} `json:"read-task"`
}

// CreateS3InstantAccessEndpointResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type CreateS3InstantAccessEndpointResponseLinks struct {
    // The HATEOAS link to this resource.
    Self                                     *HateoasSelfLink     `json:"_self"`
    // URL to the detailed information of the instant access endpoint
    ReadProtectionGroupInstantAccessEndpoint interface{}          `json:"read-protection-group-instant-access-endpoint"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask                                 *ReadTaskHateoasLink `json:"read-task"`
}

// DataAccessObject represents a custom type struct.
// Specifies information about the data-access method for accessing the restored
// file.
type DataAccessObject struct {
    // The details used to access the restored file if it was shared by direct download. If
    // the restored file was shared by email (and not by direct download), then this field
    // has a value of `null`.
    DirectDownload *DirectDownloadDataAccessObject `json:"direct_download"`
    // The details used to access the restored file, if it was shared by email. If the
    // restored file was shared by direct download (and not email), then this field has a
    // value of `null`.
    Email          *EmailDownloadDataAccessObject  `json:"email"`
}

// DeletePolicyResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type DeletePolicyResponseLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// DeleteRuleResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type DeleteRuleResponseLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// Details represents a custom type struct
type Details struct {
    // The request body of the API call
    RequestBody interface{}   `json:"request_body"`
    // The request url of the API call. This is omitempty because only internal
    // namespace should return a request_url and prod should not see that
    // a request_url is being returned.
    RequestUrl  *string       `json:"request_url"`
    // TODO: Add struct field description
    Tags        []*RestEntity `json:"tags"`
}

// DirectDownloadDataAccessObject represents a custom type struct.
// The details used to access the restored file if it was shared by direct download. If
// the restored file was shared by email (and not by direct download), then this field
// has a value of `null`.
type DirectDownloadDataAccessObject struct {
    // The download link used to access the restored file through direct download.
    DownloadLink *string `json:"download_link"`
    // The email address of the user who initiated the file level restore.
    RetrievedBy  *string `json:"retrieved_by"`
}

// Directory represents a custom type struct
type Directory struct {
    // URLs to pages related to the resource.
    Links             *DirectoryLinks `json:"_links"`
    // The directory ID of the file. If the file is not a directory, then this field has
    // a value of `null`.
    DirectoryId       *string         `json:"directory_id"`
    // Determines whether or not this file is a directory. If true, then this file
    // is a directory.
    IsDirectory       *bool           `json:"is_directory"`
    // The timestamp of when this file was last modified.
    ModifiedTimestamp *string         `json:"modified_timestamp"`
    // Name of this file.
    Name              *string         `json:"name"`
    // Size of this file, in bytes.
    Size              *uint64         `json:"size"`
}

// DirectoryBrowseEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type DirectoryBrowseEmbedded struct {
    // TODO: Add struct field description
    Items []*Directory `json:"items"`
}

// DirectoryBrowseLinks represents a custom type struct.
// URLs to pages related to the resource.
type DirectoryBrowseLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// DirectoryLinks represents a custom type struct.
// URLs to pages related to the resource.
type DirectoryLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// DiscoverConfig represents a custom type struct.
// The configuration of the Clumio Discover product for this connection.
// If this connection is not configured for Clumio Discover, then this field has a
// value of `null`.
type DiscoverConfig struct {
    // The asset types supported on the current version of the feature
    AssetTypesEnabled        []*string `json:"asset_types_enabled"`
    // The current version of the feature.
    InstalledTemplateVersion *string   `json:"installed_template_version"`
}

// DownloadSharedFileLinks represents a custom type struct.
// URLs to pages related to the resource.
type DownloadSharedFileLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// DynamoDBGRRAttributeFilter represents a custom type struct
type DynamoDBGRRAttributeFilter struct {
    // Filter condition on the DynamoDB attribute.
    // 
    // +----------------------+-------------------------------------------------------+
    // |      Condition       |                         Usage                         |
    // +======================+=======================================================+
    // | EqualTo              | Compares the filter attribute to be equal to the      |
    // |                      | operand value.                                        |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number, Binary and Boolean.              |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | NotEqualTo           | Compares the filter attribute to not be equal to the  |
    // |                      | operand value.                                        |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number, Binary and Boolean.              |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | LessThanOrEqualTo    | Compares the filter attribute to be less than or      |
    // |                      | equal to the operand value.                           |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | LessThan             | Compares the filter attribute to be less than the     |
    // |                      | operand value.                                        |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | GreaterThanOrEqualTo | Compares the filter attribute to be greater than or   |
    // |                      | equal to the operand value.                           |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | GreaterThan          | Compares the filter attribute to be greater than the  |
    // |                      | operand value.                                        |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | Between              | Compares the filter attribute to be between the       |
    // |                      | operand values.                                       |
    // |                      | It expects two operand values. Supported types are:   |
    // |                      | String, Number and Binary.                            |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | Exists               | Checks the filter attribute to exist.                 |
    // |                      | It does not expect any operand value. Supported types |
    // |                      | are: String, Number, Binary, Boolean and Null.        |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | NotExists            | Checks the filter attribute to not exist.             |
    // |                      | It does not expect any operand value. Supported types |
    // |                      | are: String, Number, Binary, Boolean and Null.        |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | Contains             | Checks the filter attribute to contain the operand    |
    // |                      | value.                                                |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String and Binary.                               |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | NotContains          | Checks the filter attribute to not contain the        |
    // |                      | operand value.                                        |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String and Binary.                               |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | BeginsWith           | Checks the filter attribute to begin with the operand |
    // |                      | value.                                                |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String and Binary.                               |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // 
    Condition  *string   `json:"condition"`
    // DynamoDB attribute name.
    Name       *string   `json:"name"`
    // Data-type of the DynamoDB attribute.
    // 
    // +----------------+
    // | Allowed values |
    // +================+
    // | String         |
    // +----------------+
    // | Number         |
    // +----------------+
    // | Binary         |
    // +----------------+
    // | Boolean        |
    // +----------------+
    // | Null           |
    // +----------------+
    // 
    ClumioType *string   `json:"type"`
    // Values for the attribute filter.
    Values     []*string `json:"values"`
}

// DynamoDBGRRQueryFilter represents a custom type struct.
// Filters based on which DynamoDB backup records are filtered in the query.
type DynamoDBGRRQueryFilter struct {
    // TODO: Add struct field description
    AttributeFilters   []*DynamoDBGRRAttributeFilter `json:"attribute_filters"`
    // Key filters based on which DynamoDB backup records are filtered.
    KeyFilters         []*DynamoDBRestoreKeyFilters  `json:"key_filters"`
    // Partition Key value of the DynamoDB table.
    // Deprecated: Use PartitionKeyFilter instead.
    PartitionKey       *string                       `json:"partition_key"`
    // Key filter of the DynamoDB table.
    PartitionKeyFilter *DynamoDBKeyFilter            `json:"partition_key_filter"`
    // Key filter of the DynamoDB table.
    SortKeyFilter      *DynamoDBKeyFilter            `json:"sort_key_filter"`
}

// DynamoDBGrrSource represents a custom type struct.
// The parameters for initiating a DynamoDB table backup query from a backup.
type DynamoDBGrrSource struct {
    // Performs the operation on a DynamoDB table within the specified backup.
    // Use the [GET /backups/aws/dynamodb-tables](#operation/list-backup-aws-dynamodb-tables)
    // endpoint to fetch valid values.
    BackupId *string `json:"backup_id"`
}

// DynamoDBGrrTarget represents a custom type struct.
// The destination information for the operation, representing the access option
// for the query results. Users can access the query results by direct download or by
// email. The direct download (`direct_download`) option allows users to directly download
// the restored file from the Clumio UI. The email (`email`) option allows users to access
// the restored file using a downloadable link they receive by email. If a target is not
// specified, then `target` defaults to `direct_download`.
type DynamoDBGrrTarget struct {
    // Specifies the Clumio UI as the restore target for direct download. Optionally set
    // `direct_download: {}`. If a target is not specified, then `target` defaults to
    // `direct_download`.
    DirectDownload interface{}                      `json:"direct_download"`
    // Specifies a download link (accessible via emails) as the restore target. If not
    // specified, `target` defaults to `direct_download`.
    Email          *EmailRecipientsDataAccessOption `json:"email"`
    // Determines whether the query is preview only. If `true`, a preview of the
    // query results will be provided in the response immediately.
    // If `false` or omitted, a task will be queued to make results
    // of the query available for asynchronous download.
    Preview        *bool                            `json:"preview"`
}

// DynamoDBKeyFilter represents a custom type struct.
// Key filter of the DynamoDB table.
type DynamoDBKeyFilter struct {
    // Filter condition on the DynamoDB key.
    // 
    // +----------------------+-------------------------------------------------------+
    // |      Condition       |                         Usage                         |
    // +======================+=======================================================+
    // | EqualTo              | Compares the filter attribute to be equal to the      |
    // |                      | operand value.                                        |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | NotEqualTo           | Compares the filter attribute to not be equal to the  |
    // |                      | operand value.                                        |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | LessThanOrEqualTo    | Compares the filter attribute to be less than or      |
    // |                      | equal to the operand value.                           |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | LessThan             | Compares the filter attribute to be less than the     |
    // |                      | operand value.                                        |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | GreaterThanOrEqualTo | Compares the filter attribute to be greater than or   |
    // |                      | equal to the operand value.                           |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | GreaterThan          | Compares the filter attribute to be greater than the  |
    // |                      | operand value.                                        |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | Between              | Compares the filter attribute to be between the       |
    // |                      | operand values.                                       |
    // |                      | It expects two operand values. Supported types are:   |
    // |                      | String, Number and Binary.                            |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | Exists               | Checks the filter attribute to exist.                 |
    // |                      | It does not expect any operand value. Supported types |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | NotExists            | Checks the filter attribute to not exist.             |
    // |                      | It does not expect any operand value. Supported types |
    // |                      | are: String, Number and Binary.                       |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | Contains             | Checks the filter attribute to contain the operand    |
    // |                      | value.                                                |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String and Binary.                               |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | NotContains          | Checks the filter attribute to not contain the        |
    // |                      | operand value.                                        |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String and Binary.                               |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // | BeginsWith           | Checks the filter attribute to begin with the operand |
    // |                      | value.                                                |
    // |                      | It expects only one operand value. Supported types    |
    // |                      | are: String and Binary.                               |
    // |                      |                                                       |
    // +----------------------+-------------------------------------------------------+
    // 
    Condition *string   `json:"condition"`
    // Values for the attribute filter.
    Values    []*string `json:"values"`
}

// DynamoDBKeys represents a custom type struct.
// Represents the DynamoDB table keys.
type DynamoDBKeys struct {
    // Represents the attributes within a DynamoDB table by the name
    // and the data type (`S` for string, `N` for number, `B` for binary).
    PartitionKey *AttributeDefinition `json:"partition_key"`
    // Represents the attributes within a DynamoDB table by the name
    // and the data type (`S` for string, `N` for number, `B` for binary).
    SortKey      *AttributeDefinition `json:"sort_key"`
}

// DynamoDBQueryPreviewResult represents a custom type struct.
// If preview was not set to true in the request, then the result of the query will be
// available for download asynchronously and this field has a value of `null`.
type DynamoDBQueryPreviewResult struct {
    // The columns of the previewed query result.
    Attributes []*string   `json:"attributes"`
    // The rows of the previewed query result.
    Items      []*[]string `json:"items"`
}

// DynamoDBRestoreKeyFilters represents a custom type struct.
// Key filters based on which DynamoDB backup records are filtered.
type DynamoDBRestoreKeyFilters struct {
    // Key filter of the DynamoDB table.
    PartitionKeyFilter *DynamoDBKeyFilter `json:"partition_key_filter"`
    // Key filter of the DynamoDB table.
    SortKeyFilter      *DynamoDBKeyFilter `json:"sort_key_filter"`
}

// DynamoDBRestoreQueryFilter represents a custom type struct.
// Filters based on which DynamoDB backup records are filtered.
type DynamoDBRestoreQueryFilter struct {
    // TODO: Add struct field description
    AttributeFilters   []*DynamoDBGRRAttributeFilter `json:"attribute_filters"`
    // Key filters based on which DynamoDB backup records are filtered.
    KeyFilters         []*DynamoDBRestoreKeyFilters  `json:"key_filters"`
    // Key filter of the DynamoDB table.
    PartitionKeyFilter *DynamoDBKeyFilter            `json:"partition_key_filter"`
    // Key filter of the DynamoDB table.
    SortKeyFilter      *DynamoDBKeyFilter            `json:"sort_key_filter"`
}

// DynamoDBRestoreSourceBackupOptions represents a custom type struct.
// The parameters for initiating a DynamoDB table restore from a backup.
type DynamoDBRestoreSourceBackupOptions struct {
    // The Clumio-assigned ID of the DynamoDB table backup to be restored. Use the
    // [GET /backups/aws/dynamodb-tables](#operation/list-backup-aws-dynamodb-tables)
    // endpoint to fetch valid values.
    BackupId *string `json:"backup_id"`
}

// DynamoDBRestoreSourcePitrOptions represents a custom type struct.
// The parameters for initiating a DynamoDB table point-in-time restore.
// Only one of `timestamp` or `use_latest_restorable_time` should be set.
type DynamoDBRestoreSourcePitrOptions struct {
    // The Clumio-assigned ID of the DynamoDB table to be restored.
    // Use the [GET /datasources/aws/dynamodb-tables](#operation/list-aws-dynamodb-tables)
    // endpoint to fetch valid values.
    TableId                 *string `json:"table_id"`
    // A point in time to be restored in RFC-3339 format.
    Timestamp               *string `json:"timestamp"`
    // The type of the backup. Possible values - `clumio_pitr`, `aws_pitr`. The
    // type will be assumed as `aws_pitr` if the field is left empty.
    ClumioType              *string `json:"type"`
    // Restore the table to the latest possible time.
    UseLatestRestorableTime *bool   `json:"use_latest_restorable_time"`
}

// DynamoDBTable represents a custom type struct
type DynamoDBTable struct {
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

// DynamoDBTableBackupLinks represents a custom type struct.
// URLs to pages related to the resource.
type DynamoDBTableBackupLinks struct {
    // The HATEOAS link to this resource.
    Self                      *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    RestoreAwsDynamodbRecords *HateoasLink     `json:"restore-aws-dynamodb-records"`
    // A resource-specific HATEOAS link.
    RestoreAwsDynamodbTable   *HateoasLink     `json:"restore-aws-dynamodb-table"`
}

// DynamoDBTableBackupListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type DynamoDBTableBackupListEmbedded struct {
    // TODO: Add struct field description
    Items []*DynamoDBTableBackupWithETag `json:"items"`
}

// DynamoDBTableBackupListLinks represents a custom type struct.
// URLs to pages related to the resource.
type DynamoDBTableBackupListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// DynamoDBTableBackupWithETag represents a custom type struct
type DynamoDBTableBackupWithETag struct {
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

// DynamoDBTableEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type DynamoDBTableEmbedded struct {
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition interface{} `json:"read-policy-definition"`
}

// DynamoDBTableLinks represents a custom type struct.
// URLs to pages related to the resource.
type DynamoDBTableLinks struct {
    // The HATEOAS link to this resource.
    Self                         *HateoasSelfLink                 `json:"_self"`
    // A resource-specific HATEOAS link.
    CreateBackupAwsDynamodbTable *HateoasLink                     `json:"create-backup-aws-dynamodb-table"`
    // A resource-specific HATEOAS link.
    ListBackupAwsDynamodbTables  *HateoasLink                     `json:"list-backup-aws-dynamodb-tables"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition         *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
    // A resource-specific HATEOAS link.
    RestoreAwsDynamodbTable      *HateoasLink                     `json:"restore-aws-dynamodb-table"`
}

// DynamoDBTableListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type DynamoDBTableListEmbedded struct {
    // TODO: Add struct field description
    Items []*DynamoDBTable `json:"items"`
}

// DynamoDBTableListLinks represents a custom type struct.
// URLs to pages related to the resource.
type DynamoDBTableListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// DynamoDBTableRestoreSource represents a custom type struct.
// The DynamoDB table restore backup or point-in-time restore options. Only one of these fields should be set.
type DynamoDBTableRestoreSource struct {
    // The parameters for initiating a DynamoDB table point-in-time restore.
    // Only one of `timestamp` or `use_latest_restorable_time` should be set.
    ContinuousBackup  *DynamoDBRestoreSourcePitrOptions   `json:"continuous_backup"`
    // The parameters for initiating a DynamoDB table restore from a backup.
    SecurevaultBackup *DynamoDBRestoreSourceBackupOptions `json:"securevault_backup"`
}

// DynamoDBTableRestoreTarget represents a custom type struct.
// The configuration of the restored DynamoDB table.
// For restore from snapshot, use the DynamoDB table configurations present at time of snapshot obtained from
// [GET /backups/aws/dynamodb-tables/{backup_id}](#operation/read-backup-aws-dynamodb-table) and for restoring point-in-time,
// use the current configuration of the table from [GET /datasources/aws/dynamodb-tables/{table_id}](#operation/read-aws-dynamodb-table).
// The table properties are set to empty or to their default values if they are specified as `null`.
type DynamoDBTableRestoreTarget struct {
    // The billing mode of the DynamoDB table. Possible values are PROVISIONED or PAY_PER_REQUEST.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // configuration of source table if both 'billing_mode' and 'provisioned_throughput' are empty or `null`.
    BillingMode               *string                 `json:"billing_mode"`
    // Indicates whether DynamoDB Contributor Insights is enabled (true) or disabled (false)
    // on the table.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // value set in backup if `null`.
    ContributorInsightsStatus *bool                   `json:"contributor_insights_status"`
    // Indicates whether DynamoDB Deletion Protection is enabled (true) or disabled (false)
    // on the table.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // value set in backup if `null`.
    DeletionProtectionEnabled *bool                   `json:"deletion_protection_enabled"`
    // The Clumio-assigned ID of the AWS environment to be used as the restore destination.
    // Use the [GET /datasources/aws/environments](#operation/list-aws-environments) endpoint
    // to fetch valid values.
    EnvironmentId             *string                 `json:"environment_id"`
    // Represents the properties of a global secondary index.
    GlobalSecondaryIndexes    []*GlobalSecondaryIndex `json:"global_secondary_indexes"`
    // Describes the version of global tables in use, if the table is replicated across AWS Regions. If the table
    // is not a global table, then this field has a value of `null`. Possible values are 2017.11.29 or 2019.11.21.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), the version is defaulted to 2019.11.21.
    GlobalTableVersion        *string                 `json:"global_table_version"`
    // Represents the properties of a local secondary index.
    LocalSecondaryIndexes     []*LocalSecondaryIndex  `json:"local_secondary_indexes"`
    // Indicates whether DynamoDB Continuous Backup (PITR) is enabled (true) or disabled (false)
    // on the table.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // value set in backup if `null`.
    PitrStatus                *bool                   `json:"pitr_status"`
    // Represents the provisioned throughput settings for a DynamoDB table.
    ProvisionedThroughput     *ProvisionedThroughput  `json:"provisioned_throughput"`
    // Contains the details of the replica.
    Replicas                  []*ReplicaDescription   `json:"replicas"`
    // Specifies the Write Capacity Units to use during table restore operations.
    // This is not used for restore to a new table without local secondary indexes.
    RestoreWcu                *int64                  `json:"restore_wcu"`
    // Represents the server-side encryption settings for a table.
    SseSpecification          *SSESpecification       `json:"sse_specification"`
    // Represents the DynamoDB Streams configuration for a table in DynamoDB.
    // and the data type (`S` for string, `N` for number, `B` for binary).
    StreamSpecification       *StreamSpecification    `json:"stream_specification"`
    // The table class of the DynamoDB table. Possible values are STANDARD or STANDARD_INFREQUENT_ACCESS.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // STANDARD storage class if empty.
    TableClass                *string                 `json:"table_class"`
    // The name of the new table to which the backup must be restored.
    TableName                 *string                 `json:"table_name"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags                      []*AwsTagCommonModel    `json:"tags"`
}

// DynamodbAssetInfo represents a custom type struct.
// DynamodbAssetInfo
// The installed information for the DynamoDB feature.
type DynamodbAssetInfo struct {
    // The current version of the feature.
    InstalledTemplateVersion *string `json:"installed_template_version"`
}

// DynamodbTemplateInfo represents a custom type struct.
// The latest available information for the DynamoDB feature.
type DynamodbTemplateInfo struct {
    // The latest available feature version for the asset.
    AvailableTemplateVersion *string `json:"available_template_version"`
}

// EBS represents a custom type struct
type EBS struct {
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

// EBSBackup represents a custom type struct
type EBSBackup struct {
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

// EBSBackupAdvancedSetting represents a custom type struct.
// Advanced settings for EBS backup.
type EBSBackupAdvancedSetting struct {
    // Backup tier to store the backup in. Valid values are: (empty) equivalent to standard, `standard`, and `lite`.
    BackupTier *string `json:"backup_tier"`
}

// EBSBackupLinks represents a custom type struct.
// URLs to pages related to the resource.
type EBSBackupLinks struct {
    // The HATEOAS link to this resource.
    Self                *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    RestoreAwsEbsVolume *HateoasLink     `json:"restore-aws-ebs-volume"`
}

// EBSBackupLinksV1 represents a custom type struct.
// URLs to pages related to the resource.
type EBSBackupLinksV1 struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// EBSBackupListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EBSBackupListEmbedded struct {
    // TODO: Add struct field description
    Items []*EBSBackup `json:"items"`
}

// EBSBackupListEmbeddedV1 represents a custom type struct.
// Embedded responses related to the resource.
type EBSBackupListEmbeddedV1 struct {
    // TODO: Add struct field description
    Items []*EBSBackupV1 `json:"items"`
}

// EBSBackupListLinks represents a custom type struct.
// URLs to pages related to the resource.
type EBSBackupListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// EBSBackupListLinksV1 represents a custom type struct.
// URLs to pages related to the resource.
type EBSBackupListLinksV1 struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// EBSBackupV1 represents a custom type struct
type EBSBackupV1 struct {
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

// EBSRestoreSource represents a custom type struct.
// The EBS volume backup to be restored.
type EBSRestoreSource struct {
    // The Clumio-assigned ID of the EBS volume backup to be restored. Use the [GET /backups/aws/ebs-volumes](#operation/list-aws-ebs-volumes) endpoint to fetch valid values.
    BackupId *string `json:"backup_id"`
}

// EBSRestoreSourceV1 represents a custom type struct.
// The EBS volume backup to be restored.
type EBSRestoreSourceV1 struct {
    // The Clumio-assigned ID of the EBS volume backup to be restored. Use the [GET /backups/aws/ebs-volumes](#operation/list-aws-ebs-volumes) endpoint to fetch valid values.
    BackupId *string `json:"backup_id"`
}

// EBSRestoreTarget represents a custom type struct.
// The configuration of the EBS volume to be restored.
type EBSRestoreTarget struct {
    // The availability zone into which the EBS volume is restored. For example, `us-west-2a`.
    // Use the [GET /datasources/aws/environments](#operation/list-aws-environments) endpoint to fetch valid values.
    AwsAz          *string              `json:"aws_az"`
    // The Clumio-assigned ID of the AWS environment to be used as the restore destination. Use the [GET /datasources/aws/environments](#operation/list-aws-environments) endpoint to fetch valid values.
    EnvironmentId  *string              `json:"environment_id"`
    // Iops of the volume to be restored.
    // Iops field is only applicable if volume_type is gp3, io1, io2.
    Iops           *int64               `json:"iops"`
    // The KMS encryption key ID used to encrypt the EBS volume data. The KMS encryption key ID is stored in the AWS cloud as part of your AWS account.
    KmsKeyNativeId *string              `json:"kms_key_native_id"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags           []*AwsTagCommonModel `json:"tags"`
    // Type of the volume to restore as.
    // Allowed Values: gp2, gp3, io1, io2, sc1, st1, standard.
    ClumioType     *string              `json:"type"`
}

// EBSRestoreTargetV1 represents a custom type struct.
// The configuration of the EBS volume to be restored.
type EBSRestoreTargetV1 struct {
    // The availability zone into which the EBS volume is restored. For example, `us-west-2a`.
    // Use the [GET /datasources/aws/environments](#operation/list-aws-environments) endpoint to fetch valid values.
    AwsAz          *string              `json:"aws_az"`
    // The Clumio-assigned ID of the AWS environment to be used as the restore destination. Use the [GET /datasources/aws/environments](#operation/list-aws-environments) endpoint to fetch valid values.
    EnvironmentId  *string              `json:"environment_id"`
    // The KMS encryption key ID used to encrypt the EBS volume data. The KMS encryption key ID is stored in the AWS cloud as part of your AWS account.
    KmsKeyNativeId *string              `json:"kms_key_native_id"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags           []*AwsTagCommonModel `json:"tags"`
}

// EC2 represents a custom type struct
type EC2 struct {
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

// EC2AMIRestoreTarget represents a custom type struct.
// The configuration for the restore to AMI.
type EC2AMIRestoreTarget struct {
    // The description for the AMI.
    Description            *string                            `json:"description"`
    // TODO: Add struct field description
    EbsBlockDeviceMappings []*EC2RestoreEbsBlockDeviceMapping `json:"ebs_block_device_mappings"`
    // The Clumio-assigned ID of the AWS environment to be used as the restore destination.
    // Use the [GET /datasources/aws/environments](#operation/list-aws-environments) endpoint
    // to fetch valid values.
    EnvironmentId          *string                            `json:"environment_id"`
    // The name for the AMI.
    Name                   *string                            `json:"name"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags                   []*AwsTagCommonModel               `json:"tags"`
}

// EC2Backup represents a custom type struct
type EC2Backup struct {
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

// EC2BackupAdvancedSetting represents a custom type struct.
// Advanced settings for EC2 backup.
type EC2BackupAdvancedSetting struct {
    // Backup tier to store the backup in. Valid values are: (empty) equivalent to standard, `standard`, and `lite`.
    BackupTier *string `json:"backup_tier"`
}

// EC2BackupLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2BackupLinks struct {
    // The HATEOAS link to this resource.
    Self                  *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    RestoreAwsEc2Instance *HateoasLink     `json:"restore-aws-ec2-instance"`
}

// EC2BackupListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2BackupListEmbedded struct {
    // TODO: Add struct field description
    Items []*EC2Backup `json:"items"`
}

// EC2BackupListLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2BackupListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// EC2InstanceListLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2InstanceListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// EC2InstanceRestoreTarget represents a custom type struct.
// The configuration of an EC2 instance to be restored.
type EC2InstanceRestoreTarget struct {
    // The AWS-assigned ID of the Amazon Machine Image (AMI) used to launch the EC2 instance.
    AmiNativeId            *string                            `json:"ami_native_id"`
    // The availability zone for the instance. This is determined by the subnet chosen to
    // restore the EC2 instance into.
    AwsAz                  *string                            `json:"aws_az"`
    // TODO: Add struct field description
    EbsBlockDeviceMappings []*EC2RestoreEbsBlockDeviceMapping `json:"ebs_block_device_mappings"`
    // The Clumio-assigned ID of the AWS environment to be used as the restore destination.
    // Use the [GET /datasources/aws/environments](#operation/list-aws-environments) endpoint
    // to fetch valid values.
    EnvironmentId          *string                            `json:"environment_id"`
    // The name of IAM instance profile to launch the instance with.
    IamInstanceProfileName *string                            `json:"iam_instance_profile_name"`
    // The name of SSH KeyPair to be used.
    KeyPairName            *string                            `json:"key_pair_name"`
    // TODO: Add struct field description
    NetworkInterfaces      []*EC2RestoreNetworkInterface      `json:"network_interfaces"`
    // Whether or not to power the instance on at the end of restore. If this is set
    // to true, the instance state will be 'running.' If it is set to false, the state
    // will be 'stopped.'
    ShouldPowerOn          *bool                              `json:"should_power_on"`
    // The AWS-assigned ID of the subnet to launch the instance into.
    SubnetNativeId         *string                            `json:"subnet_native_id"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags                   []*AwsTagCommonModel               `json:"tags"`
    // The AWS-assigned ID of the vpc to launch the instance into.
    VpcNativeId            *string                            `json:"vpc_native_id"`
}

// EC2MSSQLAG represents a custom type struct
type EC2MSSQLAG struct {
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

// EC2MSSQLAGEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MSSQLAGEmbedded struct {
    // availability group level backup status stats
    GetMssqlEc2AvailabilityGroupBackupStatusStats interface{} `json:"get-mssql-ec2-availability-group-backup-status-stats"`
    // availability group level protection stats
    GetMssqlEc2AvailabilityGroupStats             interface{} `json:"get-mssql-ec2-availability-group-stats"`
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition                          interface{} `json:"read-policy-definition"`
}

// EC2MSSQLAGLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MSSQLAGLinks struct {
    // The HATEOAS link to this resource.
    Self                                          *HateoasSelfLink                 `json:"_self"`
    // A resource-specific HATEOAS link.
    GetMssqlEc2AvailabilityGroupBackupStatusStats *HateoasLink                     `json:"get-mssql-ec2-availability-group-backup-status-stats"`
    // A resource-specific HATEOAS link.
    GetMssqlEc2AvailabilityGroupStats             *HateoasLink                     `json:"get-mssql-ec2-availability-group-stats"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition                          *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
}

// EC2MSSQLAGListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MSSQLAGListEmbedded struct {
    // TODO: Add struct field description
    Items []*EC2MSSQLAG `json:"items"`
}

// EC2MSSQLAGListLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MSSQLAGListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// EC2MSSQLDatabase represents a custom type struct
type EC2MSSQLDatabase struct {
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

// EC2MSSQLDatabaseBackup represents a custom type struct
type EC2MSSQLDatabaseBackup struct {
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

// EC2MSSQLDatabaseBackupAdvancedSetting represents a custom type struct.
// Additional policy configuration settings for the `ec2_mssql_database_backup` operation. If this operation is not of type `ec2_mssql_database_backup`, then this field is omitted from the response.
type EC2MSSQLDatabaseBackupAdvancedSetting struct {
    // The alternative replica for MSSQL database backups. This setting only applies to Availability Group databases. Possible values include `"primary"`, `"sync_secondary"`, and `"stop"`. If `"stop"` is provided, then backups will not attempt to switch to a different replica when the preferred replica is unavailable. Otherwise, recurring backups will attempt to use either the primary replica or the secondary replica accordingly.
    AlternativeReplica *string `json:"alternative_replica"`
    // The primary preferred replica for MSSQL database backups. This setting only applies to Availability Group databases. Possible values include `"primary"` and `"sync_secondary"`. Recurring backup will first attempt to use either the primary replica or the secondary replica accordingly.
    PreferredReplica   *string `json:"preferred_replica"`
}

// EC2MSSQLDatabaseBackupEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MSSQLDatabaseBackupEmbedded struct {
    // Embed information for AWS Environment details
    ReadAwsEnvironment interface{} `json:"read-aws-environment"`
}

// EC2MSSQLDatabaseBackupLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MSSQLDatabaseBackupLinks struct {
    // The HATEOAS link to this resource.
    Self                    *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    ReadAwsEnvironment      *HateoasLink     `json:"read-aws-environment"`
    // A resource-specific HATEOAS link.
    RestoreEc2MssqlDatabase *HateoasLink     `json:"restore-ec2-mssql-database"`
}

// EC2MSSQLDatabaseBackupListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MSSQLDatabaseBackupListEmbedded struct {
    // TODO: Add struct field description
    Items []*EC2MSSQLDatabaseBackup `json:"items"`
}

// EC2MSSQLDatabaseBackupListLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MSSQLDatabaseBackupListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// EC2MSSQLDatabaseEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MSSQLDatabaseEmbedded struct {
    // Embed information about the Hosts part of FCI databases
    GetEc2MssqlFailoverClustersHostsInfo interface{} `json:"get-ec2-mssql-failover-clusters-hosts-info"`
    // AWS inventory EC2 Instance embed
    ReadAwsEc2Instance                   interface{} `json:"read-aws-ec2-instance"`
    // Embed information for AWS Environment details
    ReadAwsEnvironment                   interface{} `json:"read-aws-environment"`
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition                 interface{} `json:"read-policy-definition"`
}

// EC2MSSQLDatabaseLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MSSQLDatabaseLinks struct {
    // The HATEOAS link to this resource.
    Self                         *HateoasSelfLink                 `json:"_self"`
    // A resource-specific HATEOAS link.
    CreateBackupEc2MssqlDatabase *HateoasLink                     `json:"create-backup-ec2-mssql-database"`
    // A resource-specific HATEOAS link.
    ListBackupEc2MssqlDatabases  *HateoasLink                     `json:"list-backup-ec2-mssql-databases"`
    // A resource-specific HATEOAS link.
    ReadAwsEc2Instance           *HateoasLink                     `json:"read-aws-ec2-instance"`
    // A resource-specific HATEOAS link.
    ReadAwsEnvironment           *HateoasLink                     `json:"read-aws-environment"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition         *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
}

// EC2MSSQLDatabaseListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MSSQLDatabaseListEmbedded struct {
    // TODO: Add struct field description
    Items []*EC2MSSQLDatabase `json:"items"`
}

// EC2MSSQLDatabaseListLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MSSQLDatabaseListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// EC2MSSQLFCI represents a custom type struct
type EC2MSSQLFCI struct {
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

// EC2MSSQLFCIEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MSSQLFCIEmbedded struct {
    // FCIBackupStatusStats contain information about the backup status of the databases in the cluster
    GetEc2MssqlFailoverClusterBackupStatusStats interface{} `json:"get-ec2-mssql-failover-cluster-backup-status-stats"`
    // ConnectedHostsInfo contains information about the hosts associated with the cluster
    GetEc2MssqlFailoverClusterHostsInfo         interface{} `json:"get-ec2-mssql-failover-cluster-hosts-info"`
    // FCIStats contain information about the compliant databases in the cluster
    GetEc2MssqlFailoverClusterStats             interface{} `json:"get-ec2-mssql-failover-cluster-stats"`
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition                        interface{} `json:"read-policy-definition"`
}

// EC2MSSQLFCILinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MSSQLFCILinks struct {
    // The HATEOAS link to this resource.
    Self                 *HateoasSelfLink                 `json:"_self"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
}

// EC2MSSQLFCIListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MSSQLFCIListEmbedded struct {
    // TODO: Add struct field description
    Items []*EC2MSSQLFCI `json:"items"`
}

// EC2MSSQLFCIListLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MSSQLFCIListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// EC2MSSQLInstance represents a custom type struct
type EC2MSSQLInstance struct {
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

// EC2MSSQLInstanceEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MSSQLInstanceEmbedded struct {
    // Stats pertaining to the backup status of the EC2 MSSQL Instance.
    GetEc2MssqlInstanceBackupStatusStats interface{} `json:"get-ec2-mssql-instance-backup-status-stats"`
    // Stats pertaining to the EC2 MSSQL Instance.
    GetEc2MssqlInstanceStats             interface{} `json:"get-ec2-mssql-instance-stats"`
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition                 interface{} `json:"read-policy-definition"`
}

// EC2MSSQLInstanceLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MSSQLInstanceLinks struct {
    // The HATEOAS link to this resource.
    Self                 *HateoasSelfLink                 `json:"_self"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
}

// EC2MSSQLInstanceListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MSSQLInstanceListEmbedded struct {
    // TODO: Add struct field description
    Items []*EC2MSSQLInstance `json:"items"`
}

// EC2MSSQLInstanceListLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MSSQLInstanceListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// EC2MSSQLInvHost represents a custom type struct
type EC2MSSQLInvHost struct {
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

// EC2MSSQLInvHostLinks represents a custom type struct.
// EC2MSSQLInvHostLinks contains links related to ec2 mssql host
// URLs to pages related to the resource.
type EC2MSSQLInvHostLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// EC2MSSQLInvHostListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MSSQLInvHostListEmbedded struct {
    // TODO: Add struct field description
    Items []*EC2MSSQLInvHost `json:"items"`
}

// EC2MSSQLInvHostListLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MSSQLInvHostListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// EC2MSSQLLogBackupAdvancedSetting represents a custom type struct.
// Additional policy configuration settings for the `ec2_mssql_log_backup` operation. If this operation is not of type `ec2_mssql_log_backup`, then this field is omitted from the response.
type EC2MSSQLLogBackupAdvancedSetting struct {
    // The alternative replica for MSSQL log backups. This setting only applies to Availability Group databases. Possible values include `"primary"`, `"sync_secondary"`, and `"stop"`. If `"stop"` is provided, then backups will not attempt to switch to a different replica when the preferred replica is unavailable. Otherwise, recurring backups will attempt to use either the primary replica or the secondary replica accordingly.
    AlternativeReplica *string `json:"alternative_replica"`
    // The primary preferred replica for MSSQL log backups. This setting only applies to Availability Group databases. Possible values include `"primary"` and `"sync_secondary"`. Recurring backup will first attempt to use either the primary replica or the secondary replica accordingly.
    PreferredReplica   *string `json:"preferred_replica"`
}

// EC2MSSQLPITROptions represents a custom type struct.
// A database backup at a specific point-in-time to be restored.
type EC2MSSQLPITROptions struct {
    // The Clumio-assigned ID of the MSSQL database to be restored.
    // Use the [GET /datasources/aws/ec2-mssql/databases](#operation/list-ec2-mssql-databases)
    // endpoint to fetch valid values.
    DatabaseId      *string `json:"database_id"`
    // If enabled, performs PITR till the latest possible time.
    // Either timestamp or restore_to_latest must be provided, but not both.
    RestoreToLatest *bool   `json:"restore_to_latest"`
    // The point in time to be restored in RFC-3339 format.
    // Either timestamp or restore_to_latest must be provided, but not both.
    Timestamp       *string `json:"timestamp"`
}

// EC2MSSQLProtectConfig represents a custom type struct.
// EC2MSSQLProtectConfig
// The installed information for the EC2_MSSQL feature.
type EC2MSSQLProtectConfig struct {
    // The current version of the feature.
    InstalledTemplateVersion *string `json:"installed_template_version"`
}

// EC2MSSQLRestoreFromBackupOptions represents a custom type struct.
// The EC2 MSSQL database backup to be restored.
type EC2MSSQLRestoreFromBackupOptions struct {
    // The Clumio-assigned ID of the backup to be restored.
    // Use the [GET /backups/aws/ec2-mssql/databases](#operation/list-backup-ec2-mssql-databases)
    // endpoint to fetch valid values.
    BackupId *string `json:"backup_id"`
}

// EC2MSSQLRestoreSource represents a custom type struct.
// The EC2 MSSQL database backup to be restored. Only one of `backup` or `pitr`
// should be set.
// `pitr` A database backup at a specific point in time to be restored.
type EC2MSSQLRestoreSource struct {
    // The EC2 MSSQL database backup to be restored.
    Backup       *EC2MSSQLRestoreFromBackupOptions `json:"backup"`
    // A database backup at a specific point-in-time to be restored.
    Pitr         *EC2MSSQLPITROptions              `json:"pitr"`
    // An AG database to be restored to an AAG.
    RestoreToAag *EC2MSSQLRestoreToAAGOptions      `json:"restore_to_aag"`
}

// EC2MSSQLRestoreTarget represents a custom type struct.
// The configuration of the EC2 MSSQL database to which the data has to be restored.
type EC2MSSQLRestoreTarget struct {
    // The target location within the instance to restore data files. For example,
    // `C:\\Programe Files\clumio\restored-data-files\`. If this field is empty, we
    // will restore data files into the same location as the source database.
    DataFilesPath        *string `json:"data_files_path"`
    // The user-assigned name of the database.
    DatabaseName         *string `json:"database_name"`
    // Final database state after clumio restored the database. If final_database_state
    // is set to empty then clumio will make database in online state.
    // Possible vales are `RESTORING` or `ONLINE`
    FinalDatabaseState   *string `json:"final_database_state"`
    // The Clumio-assigned ID of the instance to restore the database into.
    // Use the [GET /datasources/aws/ec2-mssql/instances](#operation/list-ec2-mssql-instances) to fetch valid values.
    InstanceId           *string `json:"instance_id"`
    // The target location within the instance to restore log files. For example,
    // `C:\\Programe Files\clumio\restored-log-files\`. If this field is empty, we
    // will restore log files into the same location as the source database.
    LogFilesPath         *string `json:"log_files_path"`
    // The boolean value representing if the database has to be restored as new database.
    RestoreAsNewDatabase *bool   `json:"restore_as_new_database"`
}

// EC2MSSQLRestoreToAAGOptions represents a custom type struct.
// An AG database to be restored to an AAG.
type EC2MSSQLRestoreToAAGOptions struct {
    // The Clumio-assigned ID of the MSSQL database to be restored.
    // Use the [GET /datasources/aws/ec2-mssql/databases](#operation/list-ec2-mssql-databases)
    // endpoint to fetch valid values.
    DatabaseId *string `json:"database_id"`
}

// EC2MSSQLTemplateInfo represents a custom type struct.
// The latest available information for the EC2 MSSQL feature.
type EC2MSSQLTemplateInfo struct {
    // The latest available feature version for the asset.
    AvailableTemplateVersion *string `json:"available_template_version"`
}

// EC2MssqlDatabasePitrInterval represents a custom type struct
type EC2MssqlDatabasePitrInterval struct {
    // Embedded responses related to the resource.
    Embedded       interface{}                        `json:"_embedded"`
    // URLs to pages related to the resource.
    Links          *EC2MssqlDatabasePitrIntervalLinks `json:"_links"`
    // End timestamp of the interval. Represented in RFC-3339 format.
    EndTimestamp   *string                            `json:"end_timestamp"`
    // Start timestamp of the interval. Represented in RFC-3339 format.
    StartTimestamp *string                            `json:"start_timestamp"`
}

// EC2MssqlDatabasePitrIntervalLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MssqlDatabasePitrIntervalLinks struct {
    // A resource-specific HATEOAS link.
    RestoreEc2MssqlDatabase *HateoasLink `json:"restore-ec2-mssql-database"`
}

// EC2MssqlDatabasePitrIntervalListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EC2MssqlDatabasePitrIntervalListEmbedded struct {
    // TODO: Add struct field description
    Items []*EC2MssqlDatabasePitrInterval `json:"items"`
}

// EC2MssqlDatabasePitrIntervalListLinks represents a custom type struct.
// URLs to pages related to the resource.
type EC2MssqlDatabasePitrIntervalListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// EC2RestoreEbsBlockDeviceMapping represents a custom type struct
type EC2RestoreEbsBlockDeviceMapping struct {
    // The AWS-assigned ID for a customer managed KMS key under which the
    // EBS volume is encrypted.
    KmsKeyNativeId *string              `json:"kms_key_native_id"`
    // The device name where the EBS volume is attached to the instance, needed by
    // instance_restore_target and ami_restore_target restore type and by volumes_restore_target
    // when target_instance_native_id is provided.
    Name           *string              `json:"name"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags           []*AwsTagCommonModel `json:"tags"`
    // The AWS-assigned ID of the backed-up volume.
    VolumeNativeId *string              `json:"volume_native_id"`
}

// EC2RestoreNetworkInterface represents a custom type struct
type EC2RestoreNetworkInterface struct {
    // The position of the network interface in the attachment order. A primary
    // network interface has a device index of 0.
    DeviceIndex              *int64    `json:"device_index"`
    // The AWS-assigned ID of the existing network interface to attach to the
    // restored instance. If one wishes to restore this network interface from the backup,
    // then this field should be set to `null`.
    NetworkInterfaceNativeId *string   `json:"network_interface_native_id"`
    // Whether or not a default network interface should be restored. It will not have any of
    // the same configurations as the backup network interface.
    RestoreDefault           *bool     `json:"restore_default"`
    // Whether or not the network interface should be restored the backup network interface.
    // It will be configured with the same configurations as the backup network interface.
    RestoreFromBackup        *bool     `json:"restore_from_backup"`
    // The AWS-assigned IDs for the security groups to associate with this network interface.
    // If one wishes to attach an existing network interface, then this field should be
    // set to `null`.
    SecurityGroupNativeIds   []*string `json:"security_group_native_ids"`
    // The AWS-assigned ID of the subnet associated with the network interface.
    // If one wishes to attach an existing network interface, then this field should be
    // set to `null`.
    SubnetNativeId           *string   `json:"subnet_native_id"`
}

// EC2RestoreSource represents a custom type struct.
// The EC2 instance backup to be restored.
type EC2RestoreSource struct {
    // The Clumio-assigned ID of the EC2 instance backup to be restored. Use the
    // [GET /backups/aws/ec2-instances](#operation/list-aws-ec2-instances)
    // endpoint to fetch valid values.
    BackupId *string `json:"backup_id"`
}

// EC2RestoreTarget represents a custom type struct.
// The target configuration per EC2 restore type. Only one of these fields should be set.
type EC2RestoreTarget struct {
    // The configuration for the restore to AMI.
    AmiRestoreTarget      *EC2AMIRestoreTarget      `json:"ami_restore_target"`
    // The configuration of an EC2 instance to be restored.
    InstanceRestoreTarget *EC2InstanceRestoreTarget `json:"instance_restore_target"`
    // The target configuration for the volumes to be restored.
    VolumesRestoreTarget  *EC2VolumesRestoreTarget  `json:"volumes_restore_target"`
}

// EC2VolumesRestoreTarget represents a custom type struct.
// The target configuration for the volumes to be restored.
type EC2VolumesRestoreTarget struct {
    // The availability zone for restoring the volumes unattached. Either this or
    // target_instance_native_id needs to be specified.
    AwsAz                  *string                            `json:"aws_az"`
    // TODO: Add struct field description
    EbsBlockDeviceMappings []*EC2RestoreEbsBlockDeviceMapping `json:"ebs_block_device_mappings"`
    // The Clumio-assigned ID of the AWS environment to be used as the restore destination.
    // Use the [GET /datasources/aws/environments](#operation/list-aws-environments) endpoint
    // to fetch valid values.
    EnvironmentId          *string                            `json:"environment_id"`
    // The aws native ID of the EC2 instance to be used to attach the restored volumes.
    // If not present, then aws_az need to be specified.
    TargetInstanceNativeId *string                            `json:"target_instance_native_id"`
}

// EbsAssetInfo represents a custom type struct.
// EbsAssetInfo
// The installed information for the EBS feature.
type EbsAssetInfo struct {
    // The current version of the feature.
    InstalledTemplateVersion *string `json:"installed_template_version"`
}

// EbsTemplateInfo represents a custom type struct
type EbsTemplateInfo struct {
    // The latest available feature version for the asset.
    AvailableTemplateVersion *string `json:"available_template_version"`
}

// EbsVolumeEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EbsVolumeEmbedded struct {
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition interface{} `json:"read-policy-definition"`
}

// EbsVolumeLinks represents a custom type struct.
// URLs to pages related to the resource.
type EbsVolumeLinks struct {
    // The HATEOAS link to this resource.
    Self                 *HateoasSelfLink                 `json:"_self"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
}

// EbsVolumeListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EbsVolumeListEmbedded struct {
    // TODO: Add struct field description
    Items []*EBS `json:"items"`
}

// EbsVolumeListLinks represents a custom type struct.
// URLs to pages related to the resource.
type EbsVolumeListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// Ec2AssetInfo represents a custom type struct.
// Ec2AssetInfo
// The installed information for the EC2 feature.
type Ec2AssetInfo struct {
    // The current version of the feature.
    InstalledTemplateVersion *string `json:"installed_template_version"`
}

// Ec2InstanceEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type Ec2InstanceEmbedded struct {
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition interface{} `json:"read-policy-definition"`
}

// Ec2InstanceLinks represents a custom type struct.
// URLs to pages related to the resource.
type Ec2InstanceLinks struct {
    // The HATEOAS link to this resource.
    Self                 *HateoasSelfLink                 `json:"_self"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
}

// Ec2InstanceListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type Ec2InstanceListEmbedded struct {
    // TODO: Add struct field description
    Items []*EC2 `json:"items"`
}

// Ec2TemplateInfo represents a custom type struct
type Ec2TemplateInfo struct {
    // The latest available feature version for the asset.
    AvailableTemplateVersion *string `json:"available_template_version"`
}

// EmailDownloadDataAccessObject represents a custom type struct.
// The details used to access the restored file, if it was shared by email. If the
// restored file was shared by direct download (and not email), then this field has a
// value of `null`.
type EmailDownloadDataAccessObject struct {
    // The optional message that was sent as part of the email.
    EmailMessage *string `json:"email_message"`
    // The email address of the user who initiated the file level restore.
    RetrievedBy  *string `json:"retrieved_by"`
    // The email address of the user who the file was retrieved for.
    RetrievedFor *string `json:"retrieved_for"`
}

// EmailDownloadDataAccessOption represents a custom type struct.
// Specifies a download link (accessible via email) as the restore target. If not
// specified, `target` defaults to `direct_download`. If you specify `email`, also
// send the user the passcode that gets generated from this request (see `passcode` in
// the response). After the user clicks the download link, they must enter the
// passcode to access the files.
type EmailDownloadDataAccessOption struct {
    // The email address of the user who will receive the download link to the restored
    // file.
    EmailAddress *string `json:"email_address"`
    // The optional message sent as part of the email.
    Message      *string `json:"message"`
}

// EmailRecipientsDataAccessOption represents a custom type struct.
// Specifies a download link (accessible via emails) as the restore target. If not
// specified, `target` defaults to `direct_download`.
type EmailRecipientsDataAccessOption struct {
    // The optional message sent as part of the email.
    Message         *string   `json:"message"`
    // The recipient email addresses who will receive the download link to the restored records.
    RecipientEmails []*string `json:"recipient_emails"`
}

// EntityGroupAssignmentUpdates represents a custom type struct.
// Updates to the organizational units along with the role assigned to the user.
type EntityGroupAssignmentUpdates struct {
    // The organizational units assigned to the user, with the specified role.
    Add    []*RoleForOrganizationalUnits `json:"add"`
    // The organizational units assigned to the user, with the specified role.
    Remove []*RoleForOrganizationalUnits `json:"remove"`
}

// EntityGroupAssignmentUpdatesV1 represents a custom type struct.
// Updates to the organizational unit assignments.
type EntityGroupAssignmentUpdatesV1 struct {
    // The Clumio-assigned IDs of the organizational units to be assigned to the user.
    Add    []*string `json:"add"`
    // The Clumio-assigned IDs of the organizational units to be unassigned to the user.
    Remove []*string `json:"remove"`
}

// EntityGroupEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type EntityGroupEmbedded struct {
    // Embeds the associated task of a resource in the response if requested using the `embed` query parameter.
    ReadTask interface{} `json:"read-task"`
}

// EntityModel represents a custom type struct.
// entityModel denotes the entityModel
type EntityModel struct {
    // The parent object of the primary entity associated with the organizational unit.
    // The parent object is optional and can be omitted.
    ParentEntity  *OrganizationalUnitParentEntity  `json:"parent_entity"`
    // The primary object associated with the organizational unit. Examples of primary entities include "aws_environment".
    PrimaryEntity *OrganizationalUnitPrimaryEntity `json:"primary_entity"`
}

// ErrorModel represents a custom type struct
type ErrorModel struct {
    // ErrorCode is a short string describing the error, if any.
    ErrorCode    *string `json:"error_code"`
    // ErrorMessage is a longer description explaining the error, if any, and how to
    // fix it.
    ErrorMessage *string `json:"error_message"`
}

// EstimateCostDetailsS3InstantAccessEndpointResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type EstimateCostDetailsS3InstantAccessEndpointResponseLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// EstimateCostS3InstantAccessEndpointAsyncResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type EstimateCostS3InstantAccessEndpointAsyncResponseLinks struct {
    // The HATEOAS link to this resource.
    Self                                                     *HateoasSelfLink     `json:"_self"`
    // URL to retrieve details of cost estimates
    CostEstimatesDetailsProtectionGroupInstantAccessEndpoint interface{}          `json:"cost-estimates-details-protection-group-instant-access-endpoint"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask                                                 *ReadTaskHateoasLink `json:"read-task"`
}

// EstimateCostS3InstantAccessEndpointSyncResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type EstimateCostS3InstantAccessEndpointSyncResponseLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// EventRules represents a custom type struct
type EventRules struct {
    // Cloudtrail rule for the service if any.
    CloudtrailRuleArn *string `json:"cloudtrail_rule_arn"`
    // Cloudwatch rule for the service if any.
    CloudwatchRuleArn *string `json:"cloudwatch_rule_arn"`
}

// FileDescriptor represents a custom type struct.
// Specifies a file/directory by providing path and file system.
type FileDescriptor struct {
    // The Clumio-assigned ID of the filesystem within which to restore the file. Use
    // [ GET /backups/files/search/{search_result_id}/versions](#operation/list-file-versions)
    // to fetch the value.
    FilesystemId *string `json:"filesystem_id"`
    // The path of the file to be restored. Use
    // [GET /backups/files/search](#operation/list-files) to fetch the value.
    Path         *string `json:"path"`
}

// FileRestoreSource represents a custom type struct.
// The files to be restored and from which backup they are to be restored from.
type FileRestoreSource struct {
    // The Clumio-assigned ID of the backup containing the files you want to restore. Use
    // [ GET /backups/files/search/{search_result_id}/versions](#operation/list-file-versions)
    // to fetch the value.
    BackupId *string           `json:"backup_id"`
    // Specifies a file/directory by providing path and file system.
    Files    []*FileDescriptor `json:"files"`
}

// FileRestoreTarget represents a custom type struct.
// The destination information for the file level restore, representing the access option
// for the restored file. Users can access the restored file by direct download or by
// email. The direct download (`direct_download`) option allows users to directly download
// the restored file from the Clumio UI. The email (`email`) option allows users to access
// the restored file using a downloadable link they receive by email. If a target is not
// specified, then `target` defaults to `direct_download`.
type FileRestoreTarget struct {
    // Specifies the Clumio UI as the restore target for direct download. Optionally set
    // `direct_download: {}`. If a target is not specified, then `target` defaults to
    // `direct_download`.
    DirectDownload interface{}                    `json:"direct_download"`
    // Specifies a download link (accessible via email) as the restore target. If not
    // specified, `target` defaults to `direct_download`. If you specify `email`, also
    // send the user the passcode that gets generated from this request (see `passcode` in
    // the response). After the user clicks the download link, they must enter the
    // passcode to access the files.
    Email          *EmailDownloadDataAccessOption `json:"email"`
}

// FileSearchListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type FileSearchListEmbedded struct {
    // TODO: Add struct field description
    Items []*FileSearchResult `json:"items"`
}

// FileSearchListLinks represents a custom type struct.
// URLs to pages related to the resource.
type FileSearchListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// FileSearchResult represents a custom type struct
type FileSearchResult struct {
    // URLs to pages related to the resource.
    Links          *ListFileVersionsHateoasLinks `json:"_links"`
    // The full file path.
    Path           *string                       `json:"path"`
    // The Clumio-assigned ID representing a collection of one or more versions of the same
    // file backed up at different times. This ID cannot be used to restore the
    // file. To restore the file, use the
    // [GET /backups/files/search/{search_result_id}/versions](#operation/list-file-versions)
    // endpoint to retrieve particular versions of this file that can be restored.
    // Then, use the backup ID, filesystem ID, and file path as parameters for the
    // [POST /restores/files](#operation/restore-files) endpoint.
    SearchResultId *string                       `json:"search_result_id"`
}

// FileSystem represents a custom type struct
type FileSystem struct {
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

// FileSystemListEmbedded represents a custom type struct.
// _embedded contains the list of items of a file_system list query
type FileSystemListEmbedded struct {
    // TODO: Add struct field description
    Items []*FileSystem `json:"items"`
}

// FileSystemListLinks represents a custom type struct.
// _links provides URLs to related navigable pages of a file_system list query
type FileSystemListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// FileVersion represents a custom type struct
type FileVersion struct {
    // URLs to pages related to the resource.
    Links             *FileVersionHateoas `json:"_links"`
    // The Clumio-assigned ID of the backup.
    BackupId          *string             `json:"backup_id"`
    // The Clumio-assigned ID of the filesystem within which to restore the file. Use
    // [ GET /backups/files/search/{search_result_id}/versions](#operation/list-file-versions)
    // to fetch the value.
    FilesystemId      *string             `json:"filesystem_id"`
    // The timestamp of the last time the file was modified. Represented in RFC-3339 format.
    ModifiedTimestamp *string             `json:"modified_timestamp"`
    // The path of the file to be restored. Use
    // [GET /backups/files/search](#operation/list-files) to fetch the value.
    Path              *string             `json:"path"`
    // The size of the file in bytes.
    Size              *int64              `json:"size"`
    // The timestamp of when the backup associated with this file started. Represented in RFC-3339 format.
    StartTimestamp    *string             `json:"start_timestamp"`
}

// FileVersionHateoas represents a custom type struct.
// URLs to pages related to the resource.
type FileVersionHateoas struct {
    // A resource-specific HATEOAS link.
    RestoreFiles *HateoasLink `json:"restore-files"`
}

// FileVersionsListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type FileVersionsListEmbedded struct {
    // TODO: Add struct field description
    Items []*FileVersion `json:"items"`
}

// FileVersionsListLinks represents a custom type struct.
// URLs to pages related to the resource.
type FileVersionsListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// GeneralSettingsLinks represents a custom type struct.
// URLs to pages related to the resource.
type GeneralSettingsLinks struct {
    // The HATEOAS link to this resource.
    Self                  *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    UpdateGeneralSettings *HateoasLink     `json:"update-general-settings"`
}

// GenerateRestoredFilePasscodeLinks represents a custom type struct.
// URLs to pages related to the resource.
type GenerateRestoredFilePasscodeLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// GlobalSecondaryIndex represents a custom type struct.
// Represents the properties of a global secondary index.
type GlobalSecondaryIndex struct {
    // Indicates whether DynamoDB Contributor Insights is enabled (true) or disabled (false)
    // on the index.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this is defaulted to the
    // value set in backup if `null`.
    ContributorInsightsStatus *bool                  `json:"contributor_insights_status"`
    // The name of the global secondary index.
    IndexName                 *string                `json:"index_name"`
    // Represents a single element of a key schema. A key schema specifies the attributes that make up the primary key
    // of a table, or the key attributes of an index.
    KeySchema                 []*KeySchemaElement    `json:"key_schema"`
    // Represents attributes that are copied (projected) from the table into an index. These are in addition to the
    // primary key attributes and index key attributes, which are automatically projected.
    Projection                *Projection            `json:"projection"`
    // Represents the provisioned throughput settings for a DynamoDB table.
    ProvisionedThroughput     *ProvisionedThroughput `json:"provisioned_throughput"`
}

// GrrSource represents a custom type struct.
// The RDS database backup to be queried.
type GrrSource struct {
    // Performs the operation on a database within the specified backup.
    // Use the [GET /backups/aws/rds-resources](#operation/list-backup-aws-rds-resources)
    // endpoint to fetch valid values.
    BackupId     *string `json:"backup_id"`
    // Performs the operation on the database with the specified name.
    // Use the [GET /backups/aws/rds-resources](#operation/list-backup-aws-rds-resource-databases)
    // endpoint to fetch valid values.
    DatabaseName *string `json:"database_name"`
}

// GrrTarget represents a custom type struct.
// The query to perform on the source RDS database.
type GrrTarget struct {
    // Determines whether the query is preview only. If `true`, a preview of the
    // query results will be provided in the response immediately.
    // If `false` or omitted, a task will be queued to make the result
    // of the query available for asynchronous download.
    Preview        *bool   `json:"preview"`
    // The SQL statement that is to be executed on the target database.
    // For example, "SELECT * FROM employee WHERE id > 100"
    QueryStatement *string `json:"query_statement"`
}

// HateoasCommonLinks represents a custom type struct.
// HateoasCommonLinks are the common fields for HATEOAS response.
type HateoasCommonLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// HateoasFirstLink represents a custom type struct.
// The HATEOAS link to the first page of results.
type HateoasFirstLink struct {
    // The URI for the referenced operation.
    Href       *string `json:"href"`
    // Determines whether the "href" link is a URI template. If set to `true`, the "href" link is a URI template.
    Templated  *bool   `json:"templated"`
    // The HTTP method to be used with the "href" link for the referenced operation.
    ClumioType *string `json:"type"`
}

// HateoasLastLink represents a custom type struct.
// The HATEOAS link to the last page of results.
type HateoasLastLink struct {
    // The URI for the referenced operation.
    Href       *string `json:"href"`
    // Determines whether the "href" link is a URI template. If set to `true`, the "href" link is a URI template.
    Templated  *bool   `json:"templated"`
    // The HTTP method to be used with the "href" link for the referenced operation.
    ClumioType *string `json:"type"`
}

// HateoasLink represents a custom type struct.
// A resource-specific HATEOAS link.
type HateoasLink struct {
    // The URI for the referenced operation.
    Href       *string `json:"href"`
    // Determines whether the "href" link is a URI template. If set to `true`, the "href" link is a URI template.
    Templated  *bool   `json:"templated"`
    // The HTTP method to be used with the "href" link for the referenced operation.
    ClumioType *string `json:"type"`
}

// HateoasNextLink represents a custom type struct.
// The HATEOAS link to the next page of results.
type HateoasNextLink struct {
    // The URI for the referenced operation.
    Href       *string `json:"href"`
    // Determines whether the "href" link is a URI template. If set to `true`, the "href" link is a URI template.
    Templated  *bool   `json:"templated"`
    // The HTTP method to be used with the "href" link for the referenced operation.
    ClumioType *string `json:"type"`
}

// HateoasPrevLink represents a custom type struct.
// The HATEOAS link to the previous page of results.
type HateoasPrevLink struct {
    // The URI for the referenced operation.
    Href       *string `json:"href"`
    // Determines whether the "href" link is a URI template. If set to `true`, the "href" link is a URI template.
    Templated  *bool   `json:"templated"`
    // The HTTP method to be used with the "href" link for the referenced operation.
    ClumioType *string `json:"type"`
}

// HateoasSelfLink represents a custom type struct.
// The HATEOAS link to this resource.
type HateoasSelfLink struct {
    // The URI for the referenced operation.
    Href       *string `json:"href"`
    // Determines whether the "href" link is a URI template. If set to `true`, the "href" link is a URI template.
    Templated  *bool   `json:"templated"`
    // The HTTP method to be used with the "href" link for the referenced operation.
    ClumioType *string `json:"type"`
}

// IamInstanceProfileModel represents a custom type struct.
// Denotes an IAM instance profile. An instance profile is a container for an
// IAM role that you can use to pass role information to an EC2 instance when
// the instance starts.
type IamInstanceProfileModel struct {
    // The Amazon Resource Name (ARN) specifying the IAM instance profile.
    Arn      *string `json:"arn"`
    // The AWS-assigned name of the IAM instance profile.
    Name     *string `json:"name"`
    // The AWS-assigned ID of the IAM instance profile.
    NativeId *string `json:"native_id"`
}

// IcebergBackupAdvancedSetting represents a custom type struct.
// IcebergBackupAdvancedSetting defines the advanced settings for Iceberg backup operations
type IcebergBackupAdvancedSetting struct {
    // Backup tier to store the backup in. Valid values are: `standard`.
    BackupTier *string `json:"backup_tier"`
}

// IcebergOnGlueAssetInfo represents a custom type struct.
// IcebergOnGlueAssetInfo
// The installed information for the Iceberg on AWS Glue feature.
type IcebergOnGlueAssetInfo struct {
    // The current version of the feature.
    InstalledTemplateVersion *string `json:"installed_template_version"`
}

// IcebergOnGlueTemplateInfo represents a custom type struct.
// IcebergOnGlueTemplateInfo is the latest available information for the IcebergOnGlue feature.
type IcebergOnGlueTemplateInfo struct {
    // The latest available feature version for the asset.
    AvailableTemplateVersion *string `json:"available_template_version"`
}

// IcebergOnS3TablesAssetInfo represents a custom type struct.
// IcebergOnS3TablesAssetInfo
// The installed information for the Iceberg on S3 Tables feature.
type IcebergOnS3TablesAssetInfo struct {
    // The current version of the feature.
    InstalledTemplateVersion *string `json:"installed_template_version"`
}

// IcebergOnS3TablesTemplateInfo represents a custom type struct.
// IcebergOnS3TablesTemplateInfo is the latest available information for the IcebergOnS3Tables feature.
type IcebergOnS3TablesTemplateInfo struct {
    // The latest available feature version for the asset.
    AvailableTemplateVersion *string `json:"available_template_version"`
}

// IndividualAlertDetails represents a custom type struct.
// Additional information about the alert.
type IndividualAlertDetails struct {
    // A brief description of the condition that caused the alert. Examples include "Size Limit Exceeded" and "Insufficient Cloud Connector Capacity".
    Cause       *string            `json:"cause"`
    // A detailed description of the alert, including the reason why the alert occurred
    // and the steps you must take to resolve the issue.
    Description *string            `json:"description"`
    // The general alert category. Examples include "Policy Violated" and "Restore Failed".
    ClumioType  *string            `json:"type"`
    // Data specific to the alert generated. If the alert has no variables, then this
    // field has a value of `null`.
    Variables   map[string]*string `json:"variables"`
}

// InstanceStoreBlockDeviceMapping represents a custom type struct
type InstanceStoreBlockDeviceMapping struct {
    // Encryption for the instance store volume. Possible values include 'hardware encrypted'
    // and 'Not encrypted'.
    Encryption  *string `json:"encryption"`
    // Determines whether or not the volume is a NVME instance store volume or a
    // non-NVME instance store volume.
    IsNvme      *bool   `json:"is_nvme"`
    // The device name for the instance store volume. For example, '/dev/sdb'.
    Name        *string `json:"name"`
    // The size of the instance store volume. Measured in bytes (B).
    Size        *int64  `json:"size"`
    // The type of the block device. Only possible value is "Instance Store".
    ClumioType  *string `json:"type"`
    // The AWS-assigned name of the instance store volume.
    VirtualName *string `json:"virtual_name"`
}

// ItemsCovered represents a custom type struct.
// The items covered in the compliance report created by the report run.
type ItemsCovered struct {
    // The count of covered assets of the report run.
    AssetCount  *int64 `json:"asset_count"`
    // The count of covered policies of the report run.
    PolicyCount *int64 `json:"policy_count"`
}

// KeySchemaElement represents a custom type struct.
// Represents a single element of a key schema. A key schema specifies the attributes that make up the primary key
// of a table, or the key attributes of an index.
type KeySchemaElement struct {
    // The name of a key attribute.
    AttributeName *string `json:"attribute_name"`
    // The role that this key attribute will assume.
    // Possible values include: `HASH` - partition key and `RANGE` - sort key.
    KeyType       *string `json:"key_type"`
}

// LatestRun represents a custom type struct.
// Most recent report run generated from the report configuration.
type LatestRun struct {
    // The status per controls in the compliance report created by the report run.
    ComplianceInfo     *ComplianceInfo `json:"compliance_info"`
    // The RFC3339 format time when the report run was created.
    Created            *string         `json:"created"`
    // The RFC3339 format time when the report run was expired.
    Expired            *string         `json:"expired"`
    // The unique identifier of the report run.
    Id                 *string         `json:"id"`
    // The name of the report run.
    Name               *string         `json:"name"`
    // Filter and control parameters of compliance report.
    Parameter          *Parameter      `json:"parameter"`
    // The unique identifier of the report configuration from which the report run was generated.
    ReportConfigId     *string         `json:"report_config_id"`
    // The link to download the report CSV.
    ReportDownloadLink *string         `json:"report_download_link"`
    // The generation status of the report run.
    Status             *string         `json:"status"`
    // The ID of the report run generation task.
    TaskId             *string         `json:"task_id"`
    // The RFC3339 format time when the report run was updated.
    Updated            *string         `json:"updated"`
}

// ListFileVersionsHateoasLink represents a custom type struct.
// A HATEOAS link to the file versions associated with this resource.
type ListFileVersionsHateoasLink struct {
    // The URI for the referenced operation.
    Href       *string `json:"href"`
    // Determines whether the "href" link is a URI template. If set to `true`, the "href" link is a URI template.
    Templated  *bool   `json:"templated"`
    // The HTTP method to be used with the "href" link for the referenced operation.
    ClumioType *string `json:"type"`
}

// ListFileVersionsHateoasLinks represents a custom type struct.
// URLs to pages related to the resource.
type ListFileVersionsHateoasLinks struct {
    // A HATEOAS link to the file versions associated with this resource.
    ListFileVersions *ListFileVersionsHateoasLink `json:"list-file-versions"`
}

// LocalSecondaryIndex represents a custom type struct.
// Represents the properties of a local secondary index.
type LocalSecondaryIndex struct {
    // The name of the local secondary index
    IndexName  *string             `json:"index_name"`
    // Represents a single element of a key schema. A key schema specifies the attributes that make up the primary key
    // of a table, or the key attributes of an index.
    KeySchema  []*KeySchemaElement `json:"key_schema"`
    // Represents attributes that are copied (projected) from the table into an index. These are in addition to the
    // primary key attributes and index key attributes, which are automatically projected.
    Projection *Projection         `json:"projection"`
}

// M365GroupingCriteria represents a custom type struct.
// The entity type used to group organizational units for Microsoft 365 resources.
type M365GroupingCriteria struct {
    // Determines whether or not this data group is editable. If false, then an
    // organizational unit uses this data group.
    // To edit this data group, all organizational units using it must be deleted.
    IsEditable *bool   `json:"is_editable"`
    // 
    // +---------------------+------------------------+
    // |     Entity Type     |        Details         |
    // +=====================+========================+
    // | microsoft365_domain | Microsoft 365 account. |
    // +---------------------+------------------------+
    // 
    ClumioType *string `json:"type"`
}

// MSSQLDatabaseBackupAdvancedSetting represents a custom type struct.
// Additional policy configuration settings for the `mssql_database_backup` operation. If this operation is not of type `mssql_database_backup`, then this field is omitted from the response.
type MSSQLDatabaseBackupAdvancedSetting struct {
    // The alternative replica for MSSQL database backups. This setting only applies to Availability Group databases. Possible values include `"primary"`, `"sync_secondary"`, and `"stop"`. If `"stop"` is provided, then backups will not attempt to switch to a different replica when the preferred replica is unavailable. Otherwise, recurring backups will attempt to use either the primary replica or the secondary replica accordingly.
    AlternativeReplica *string `json:"alternative_replica"`
    // The primary preferred replica for MSSQL database backups. This setting only applies to Availability Group databases. Possible values include `"primary"` and `"sync_secondary"`. Recurring backup will first attempt to use either the primary replica or the secondary replica accordingly.
    PreferredReplica   *string `json:"preferred_replica"`
}

// MSSQLLogBackupAdvancedSetting represents a custom type struct.
// Additional policy configuration settings for the `mssql_log_backup` operation. If this operation is not of type `mssql_log_backup`, then this field is omitted from the response.
type MSSQLLogBackupAdvancedSetting struct {
    // The alternative replica for MSSQL log backups. This setting only applies to Availability Group databases. Possible values include `"primary"`, `"sync_secondary"`, and `"stop"`. If `"stop"` is provided, then backups will not attempt to switch to a different replica when the preferred replica is unavailable. Otherwise, recurring backups will attempt to use either the primary replica or the secondary replica accordingly.
    AlternativeReplica *string `json:"alternative_replica"`
    // The primary preferred replica for MSSQL log backups. This setting only applies to Availability Group databases. Possible values include `"primary"` and `"sync_secondary"`. Recurring backup will first attempt to use either the primary replica or the secondary replica accordingly.
    PreferredReplica   *string `json:"preferred_replica"`
}

// ManagementGroup represents a custom type struct
type ManagementGroup struct {
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

// ManagementGroupLinks represents a custom type struct.
// URLs to pages related to the resource.
type ManagementGroupLinks struct {
    // The HATEOAS link to this resource.
    Self                  *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    UpdateManagementGroup *HateoasLink     `json:"update-management-group"`
}

// ManagementGroupListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ManagementGroupListEmbedded struct {
    // TODO: Add struct field description
    Items []*ManagementGroup `json:"items"`
}

// ManagementGroupListLinks represents a custom type struct.
// URLs to pages related to the resource.
type ManagementGroupListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// MssqlDatabaseFile represents a custom type struct
type MssqlDatabaseFile struct {
    // The name of the database file.
    Name       *string `json:"name"`
    // The type of the database file. Possible values include sql_row_file and sql_log_file.
    ClumioType *string `json:"type"`
}

// MssqlServiceInstanceProfiles represents a custom type struct
type MssqlServiceInstanceProfiles struct {
    // Policy created for ec2 instance profile role
    Ec2SsmInstanceProfileArn *string `json:"ec2_ssm_instance_profile_arn"`
}

// MssqlServiceRoles represents a custom type struct
type MssqlServiceRoles struct {
    // Role assumable by ec2 service.
    Ec2InstanceProfileRoleArn *string `json:"ec2_instance_profile_role_arn"`
    // Instance created for ec2 instance profile role
    Ec2SsmInstanceProfileArn  *string `json:"ec2_ssm_instance_profile_arn"`
    // Role assumable by ssm service.
    SsmNotificationRoleArn    *string `json:"ssm_notification_role_arn"`
}

// NetworkInterface represents a custom type struct
type NetworkInterface struct {
    // The device index for the network interface.
    DeviceIndex              *int64    `json:"device_index"`
    // The AWS-assigned ID for the network interface.
    NetworkInterfaceNativeId *string   `json:"network_interface_native_id"`
    // The public IP v4 address of the network interface if one was assigned.
    PublicIp                 *string   `json:"public_ip"`
    // The AWS-assigned IDs for the security groups associated with this network interface.
    SecurityGroupNativeIds   []*string `json:"security_group_native_ids"`
    // The subnet native ID for the network interface.
    SubnetNativeId           *string   `json:"subnet_native_id"`
    // The AWS-assigned name of the network interface. For example, `eth0`.
    VirtualName              *string   `json:"virtual_name"`
}

// NotificationSetting represents a custom type struct.
// Notification channels to send the generated report runs.
type NotificationSetting struct {
    // Email list to send a generated report run.
    EmailList []*string `json:"email_list"`
}

// OUGroupingCriteria represents a custom type struct.
// The grouping criteria for each datasource type.
// These can only be edited for datasource types which do not have any
// organizational units configured.
// Deprecated: This struct is deprecated and will be removed in the future.
// It is being kept for backward compatibility.
type OUGroupingCriteria struct {
    // The entity type used to group organizational units for AWS resources.
    Aws          *AwsDsGroupingCriteria `json:"aws"`
    // The entity type used to group organizational units for Microsoft 365 resources.
    Microsoft365 *M365GroupingCriteria  `json:"microsoft365"`
}

// OULinks represents a custom type struct.
// URLs to pages related to the resource.
type OULinks struct {
    // The HATEOAS link to this resource.
    Self                     *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    DeleteOrganizationalUnit *HateoasLink     `json:"delete-organizational-unit"`
    // A resource-specific HATEOAS link.
    PatchOrganizationalUnit  *HateoasLink     `json:"patch-organizational-unit"`
}

// Object represents a custom type struct.
// Object defines one object to restore
type Object struct {
    // Bucket the object belongs to
    Bucket                 *string `json:"bucket"`
    // Etag of the contents of the object.
    Etag                   *string `json:"etag"`
    // Last time the object was backed up as an RFC3339 string.
    LastBackupTime         *string `json:"last_backup_time"`
    // Last modified time of the object as an RFC3339 string.
    LastModifiedTime       *string `json:"last_modified_time"`
    // Object key
    ObjectKey              *string `json:"object_key"`
    // The Clumio-assigned ID of a protection group S3 asset,
    // which represents the bucket within the protection group to restore from.
    ProtectionGroupAssetId *string `json:"protection_group_asset_id"`
    // region of the backup object
    Region                 *string `json:"region"`
    // Encrypted metadata for the object to be restored 
    // You can get `restore_cookie` via
    // [POST /restores/protection-groups/{protection_group_id}/previews](#operation/preview-protection-group)
    RestoreCookie          *string `json:"restore_cookie"`
    // Size in Bytes
    SizeInBytes            *int64  `json:"size_in_bytes"`
    // Storage class. Valid values are: `S3 Standard`, `S3 Standard-IA`, `S3 Intelligent-Tiering`,
    // and `S3 One Zone-IA`.
    StorageClass           *string `json:"storage_class"`
    // Version ID
    VersionId              *string `json:"version_id"`
}

// ObjectFilter represents a custom type struct.
// ObjectFilter
// defines which objects will be backed up.
type ObjectFilter struct {
    // The cutoff date for inclusion objects from the backup. Any object with a last modified
    // date after or equal  than this value will  be included in the backup. This is useful for
    // filtering out old or irrelevant objects based on their modification timestamps.
    // EarliestLastModifiedTimeStamp support RFC-3339 format.
    EarliestLastModifiedTimestamp *string         `json:"earliest_last_modified_timestamp"`
    // Whether to back up only the latest object version.
    LatestVersionOnly             *bool           `json:"latest_version_only"`
    // PrefixFilter
    PrefixFilters                 []*PrefixFilter `json:"prefix_filters"`
    // Storage class to include in the backup. If not specified, then all objects across all storage
    // classes will be backed up. Valid values are: `S3 Standard`, `S3 Standard-IA`,
    // `S3 Intelligent-Tiering`, and `S3 One Zone-IA`.
    StorageClasses                []*string       `json:"storage_classes"`
}

// ObjectV2 represents a custom type struct.
// ObjectV2 defines one object to restore
type ObjectV2 struct {
    // Bucket the object belongs to
    Bucket                 *string `json:"bucket"`
    // Etag of the contents of the object.
    Etag                   *string `json:"etag"`
    // Last time the object was backed up as an RFC3339 string.
    LastBackupTimestamp    *string `json:"last_backup_timestamp"`
    // Last modified time of the object as an RFC3339 string.
    LastModifiedTimestamp  *string `json:"last_modified_timestamp"`
    // Object key
    ObjectKey              *string `json:"object_key"`
    // The Clumio-assigned ID of a protection group S3 asset,
    // which represents the bucket within the protection group to restore from.
    ProtectionGroupAssetId *string `json:"protection_group_asset_id"`
    // region of the backup object
    Region                 *string `json:"region"`
    // Encrypted metadata for the object to be restored 
    // You can get `restore_cookie` via
    // [POST /restores/protection-groups/{protection_group_id}/previews](#operation/preview-protection-group)
    RestoreCookie          *string `json:"restore_cookie"`
    // Size in Bytes
    SizeInBytes            *int64  `json:"size_in_bytes"`
    // Storage class. Valid values are: `S3 Standard`, `S3 Standard-IA`, `S3 Intelligent-Tiering`,
    // and `S3 One Zone-IA`.
    StorageClass           *string `json:"storage_class"`
    // Version ID
    VersionId              *string `json:"version_id"`
}

// OnDemandDynamoDBBackupLinks represents a custom type struct.
// URLs to pages related to the resource.
type OnDemandDynamoDBBackupLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// OnDemandEBSBackupLinks represents a custom type struct.
// URLs to pages related to the resource.
type OnDemandEBSBackupLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// OnDemandEC2BackupLinks represents a custom type struct.
// URLs to pages related to the resource.
type OnDemandEC2BackupLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// OnDemandSetting represents a custom type struct.
// Settings for requesting on-demand backup directly.
type OnDemandSetting struct {
    // Additional operation-specific policy settings. For operation types which do not support additional settings, this field is `null`.
    AdvancedSettings  *PolicyAdvancedSettings  `json:"advanced_settings"`
    // The region in which this backup is stored. This might be used for cross-region backup.
    BackupAwsRegion   *string                  `json:"backup_aws_region"`
    // The retention time for this SLA. For example, to retain the backup for 1 month, set `unit="months"` and `value=1`.
    RetentionDuration *RetentionBackupSLAParam `json:"retention_duration"`
}

// OnDemandThroughputOverride represents a custom type struct.
// Replica-specific ondemand throughput settings. If not specified, uses the source table's ondemand throughput settings.
type OnDemandThroughputOverride struct {
    // The maximum number of strongly consistent reads consumed per second.
    MaxReadRequestUnits *int64 `json:"max_read_request_units"`
}

// OperationInfo represents a custom type struct
type OperationInfo struct {
    // BackupStatus is the status of the backup. Possible values are
    // `success`, `partial_success`, `failure`, `no_backup`, and `unknown`. This value
    // depends on `lookback_days`. If not specified, then this field has a value of `unknown`.
    BackupStatus                       *string `json:"backup_status"`
    // The last failed policy start time. Represented in RFC-3339 format.
    LastFailedPolicyStartTimestamp     *string `json:"last_failed_policy_start_timestamp"`
    // The last successful policy start time. Represented in RFC-3339 format.
    LastSuccessfulPolicyStartTimestamp *string `json:"last_successful_policy_start_timestamp"`
    // The policy operation type.
    Operation                          *string `json:"operation"`
}

// OptionGroups represents a custom type struct
type OptionGroups struct {
    // Embedded responses related to the resource.
    Embedded                          interface{}        `json:"_embedded"`
    // URLs to pages related to the resource.
    Links                             *OptionGroupsLinks `json:"_links"`
    // The AWS database engine at the time of backup.
    Engine                            *string            `json:"engine"`
    // The AWS database engine version at the time of backup.
    EngineVersion                     *string            `json:"engine_version"`
    // Determines whether this option group contains additional non-permanent options apart from
    // the non-permanent options at the time of backup.
    HasAdditionalNonPermanentOptions  *bool              `json:"has_additional_non_permanent_options"`
    // Determines whether this option group contains additional non-persistent options apart from
    // the non-persistent options at time of backup.
    HasAdditionalNonPersistentOptions *bool              `json:"has_additional_non_persistent_options"`
    // Determines whether this option group contains additional permanent options apart from
    // the permanent options at the time of backup.
    HasAdditionalPermanentOptions     *bool              `json:"has_additional_permanent_options"`
    // Determines whether this option group contains additional persistent options apart from
    // the persistent options at the time of backup.
    HasAdditionalPersistentOptions    *bool              `json:"has_additional_persistent_options"`
    // Compatibility of the Option Group
    IsCompatible                      *bool              `json:"is_compatible"`
    // Minimum required minor engine version for this option-group to be applicable.
    MinimumRequiredMinorEngineVersion *string            `json:"minimum_required_minor_engine_version"`
    // Name of the option group
    Name                              *string            `json:"name"`
    // OptionModel denotes the Model for OptionModel
    Options                           []*OptionModel     `json:"options"`
}

// OptionGroupsLinks represents a custom type struct.
// URLs to pages related to the resource.
type OptionGroupsLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// OptionGroupsListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type OptionGroupsListEmbedded struct {
    // TODO: Add struct field description
    Items []*OptionGroups `json:"items"`
}

// OptionGroupsListLinks represents a custom type struct.
// URLs to pages related to the resource.
type OptionGroupsListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// OptionModel represents a custom type struct.
// OptionModel denotes the Model for OptionModel
type OptionModel struct {
    // Determines whether or not the option is permanent.
    IsPermanent          *bool            `json:"is_permanent"`
    // Determines whether or not the option is persistent.
    IsPersistent         *bool            `json:"is_persistent"`
    // Determines whether the option is required to restore from a given backup.
    IsRequiredForRestore *bool            `json:"is_required_for_restore"`
    // The AWS-assigned name of the RDS option.
    Name                 *string          `json:"name"`
    // OptionSetting denotes the Model for OptionSetting
    OptionSetting        []*OptionSetting `json:"option_setting"`
    // Option version of the option.
    OptionVersion        *string          `json:"option_version"`
}

// OptionSetting represents a custom type struct.
// OptionSetting denotes the Model for OptionSetting
type OptionSetting struct {
    // The AWS-assigned description of the RDS option setting.
    Description *string `json:"description"`
    // The AWS-assigned name of the RDS option setting.
    Name        *string `json:"name"`
    // Value of the option setting
    Value       *string `json:"value"`
}

// OracleDatabaseBackupAdvancedSetting represents a custom type struct.
// Additional policy configuration settings for the `oracle_database_backup` operation. If this operation is not of type `oracle_database_backup`, then this field is omitted from the response.
type OracleDatabaseBackupAdvancedSetting struct {
    // The alternative replica for Oracle database backups. This setting only applies to Data Guard databases. Valid values are: `primary`, `standby`, and `stop`.
    // If `"stop"` is provided, then backups will not attempt to switch to a different replica when the preferred replica is unavailable.
    AlternativeReplica *string `json:"alternative_replica"`
    // The primary preferred replica for Oracle database backups. This setting only applies to Data Guard databases. Valid values are: `primary`, and `standby`.
    // Recurring backup will first attempt to use either the primary replica or the secondary replica accordingly.
    PreferredReplica   *string `json:"preferred_replica"`
}

// OracleLogBackupAdvancedSetting represents a custom type struct.
// Additional policy configuration settings for the `oracle_log_backup` operation. If this operation is not of type `oracle_log_backup`, then this field is omitted from the response.
type OracleLogBackupAdvancedSetting struct {
    // The alternative replica for Oracle database backups. This setting only applies to Data Guard databases. Valid values are: `primary`, `standby`, and `stop`.
    // If `"stop"` is provided, then backups will not attempt to switch to a different replica when the preferred replica is unavailable.
    AlternativeReplica *string `json:"alternative_replica"`
    // The primary preferred replica for Oracle database backups. This setting only applies to Data Guard databases. Valid values are: `primary`, and `standby`.
    // Recurring backup will first attempt to use either the primary replica or the secondary replica accordingly.
    PreferredReplica   *string `json:"preferred_replica"`
}

// OrganizationalUnitLinks represents a custom type struct.
// URLs to pages related to the resource.
type OrganizationalUnitLinks struct {
    // The HATEOAS link to this resource.
    Self                     *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    DeleteOrganizationalUnit *HateoasLink     `json:"delete-organizational-unit"`
    // A resource-specific HATEOAS link.
    PatchOrganizationalUnit  *HateoasLink     `json:"patch-organizational-unit"`
    // A resource-specific HATEOAS link.
    ReadTask                 *HateoasLink     `json:"read-task"`
}

// OrganizationalUnitLinksForDelete represents a custom type struct.
// URLs to pages related to the resource.
type OrganizationalUnitLinksForDelete struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    ReadTask *HateoasLink     `json:"read-task"`
}

// OrganizationalUnitListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type OrganizationalUnitListEmbedded struct {
    // OrganizationalUnitWithETag to support etag string to be calculated
    Items []*OrganizationalUnitWithETag `json:"items"`
}

// OrganizationalUnitListEmbeddedV1 represents a custom type struct.
// Embedded responses related to the resource.
type OrganizationalUnitListEmbeddedV1 struct {
    // OrganizationalUnitWithETagV1 to support etag string to be calculated
    Items []*OrganizationalUnitWithETagV1 `json:"items"`
}

// OrganizationalUnitListLinks represents a custom type struct.
// URLs to pages related to the resource.
type OrganizationalUnitListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// OrganizationalUnitParentEntity represents a custom type struct.
// The parent object of the primary entity associated with the organizational unit.
// The parent object is optional and can be omitted.
type OrganizationalUnitParentEntity struct {
    // The Clumio assigned ID of the entity.
    Id         *string `json:"id"`
    // The entity type.
    ClumioType *string `json:"type"`
}

// OrganizationalUnitPrimaryEntity represents a custom type struct.
// The primary object associated with the organizational unit. Examples of primary entities include "aws_environment".
type OrganizationalUnitPrimaryEntity struct {
    // The Clumio assigned ID of the entity.
    Id         *string `json:"id"`
    // The entity type.
    ClumioType *string `json:"type"`
}

// OrganizationalUnitWithETag represents a custom type struct.
// OrganizationalUnitWithETag to support etag string to be calculated
type OrganizationalUnitWithETag struct {
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

// OrganizationalUnitWithETagV1 represents a custom type struct.
// OrganizationalUnitWithETagV1 to support etag string to be calculated
type OrganizationalUnitWithETagV1 struct {
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

// Parameter represents a custom type struct.
// Filter and control parameters of compliance report.
type Parameter struct {
    // Compliance controls to evaluate policy or assets for compliance.
    Controls *ComplianceControls `json:"controls"`
    // The set of filters supported in compliance report.
    Filters  *ComplianceFilters  `json:"filters"`
}

// PermissionListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type PermissionListEmbedded struct {
    // TODO: Add struct field description
    Items []*PermissionModel `json:"items"`
}

// PermissionListLinks represents a custom type struct.
// URLs to pages related to the resource.
type PermissionListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// PermissionModel represents a custom type struct
type PermissionModel struct {
    // Description of the permission.
    Description *string `json:"description"`
    // The Clumio-assigned ID of the permission.
    Id          *string `json:"id"`
    // Name of the permission.
    // The following table lists the supported permissions for a role:
    // 
    // +----------------------------------------------------+-------------------------+
    // |                  Permission Name                   | Full/Partial Applicable |
    // +====================================================+=========================+
    // | Policy Management                                  | Yes                     |
    // +----------------------------------------------------+-------------------------+
    // | Data Source Management                             | Yes                     |
    // +----------------------------------------------------+-------------------------+
    // | Perform Backup (Scheduled or On-demand)            | Yes                     |
    // +----------------------------------------------------+-------------------------+
    // | Regular Restore                                    | No                      |
    // +----------------------------------------------------+-------------------------+
    // | Redirected Granular Restore (things like GRR &     | Yes                     |
    // | content download)                                  |                         |
    // +----------------------------------------------------+-------------------------+
    // | Dashboards & Reports                               | Yes                     |
    // +----------------------------------------------------+-------------------------+
    // | Some Admin Functions (User mgmt, SSO/MFA, IP       | No                      |
    // | Allow, Password expiry, OU, KMS)                   |                         |
    // +----------------------------------------------------+-------------------------+
    // | Other Admin Functions (API Tokens, Tasks, Alerts   | Yes                     |
    // | and Audit Logs)                                    |                         |
    // +----------------------------------------------------+-------------------------+
    // 
    Name        *string `json:"name"`
}

// Policy represents a custom type struct
type Policy struct {
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

// PolicyAdvancedSettings represents a custom type struct.
// Additional operation-specific policy settings. For operation types which do not support additional settings, this field is `null`.
type PolicyAdvancedSettings struct {
    // Advanced settings for EBS backup.
    AwsEbsVolumeBackup              *EBSBackupAdvancedSetting                       `json:"aws_ebs_volume_backup"`
    // Advanced settings for EC2 backup.
    AwsEc2InstanceBackup            *EC2BackupAdvancedSetting                       `json:"aws_ec2_instance_backup"`
    // IcebergBackupAdvancedSetting defines the advanced settings for Iceberg backup operations
    AwsIcebergTableBackup           *IcebergBackupAdvancedSetting                   `json:"aws_iceberg_table_backup"`
    // Advanced settings for RDS PITR configuration sync.
    AwsRdsConfigSync                *RDSConfigSyncAdvancedSetting                   `json:"aws_rds_config_sync"`
    // Settings for determining if a RDS policy is created with standard or archive tier.
    AwsRdsResourceGranularBackup    *RDSLogicalBackupAdvancedSetting                `json:"aws_rds_resource_granular_backup"`
    // Additional policy configuration settings for the `ec2_mssql_database_backup` operation. If this operation is not of type `ec2_mssql_database_backup`, then this field is omitted from the response.
    Ec2MssqlDatabaseBackup          *EC2MSSQLDatabaseBackupAdvancedSetting          `json:"ec2_mssql_database_backup"`
    // Additional policy configuration settings for the `ec2_mssql_log_backup` operation. If this operation is not of type `ec2_mssql_log_backup`, then this field is omitted from the response.
    Ec2MssqlLogBackup               *EC2MSSQLLogBackupAdvancedSetting               `json:"ec2_mssql_log_backup"`
    // Additional policy configuration settings for the `mssql_database_backup` operation. If this operation is not of type `mssql_database_backup`, then this field is omitted from the response.
    MssqlDatabaseBackup             *MSSQLDatabaseBackupAdvancedSetting             `json:"mssql_database_backup"`
    // Additional policy configuration settings for the `mssql_log_backup` operation. If this operation is not of type `mssql_log_backup`, then this field is omitted from the response.
    MssqlLogBackup                  *MSSQLLogBackupAdvancedSetting                  `json:"mssql_log_backup"`
    // Additional policy configuration settings for the `oracle_database_backup` operation. If this operation is not of type `oracle_database_backup`, then this field is omitted from the response.
    OracleDatabaseBackup            *OracleDatabaseBackupAdvancedSetting            `json:"oracle_database_backup"`
    // Additional policy configuration settings for the `oracle_log_backup` operation. If this operation is not of type `oracle_log_backup`, then this field is omitted from the response.
    OracleLogBackup                 *OracleLogBackupAdvancedSetting                 `json:"oracle_log_backup"`
    // Additional policy configuration settings for the `protection_group_backup` operation. If this operation is not of type `protection_group_backup`, then this field is omitted from the response.
    ProtectionGroupBackup           *ProtectionGroupBackupAdvancedSetting           `json:"protection_group_backup"`
    // Additional policy configuration settings for the `protection_group_continuous_backup` operation. If this operation is not of type `protection_group_continuous_backup`, then this field is omitted from the response.
    ProtectionGroupContinuousBackup *ProtectionGroupContinuousBackupAdvancedSetting `json:"protection_group_continuous_backup"`
}

// PolicyControl represents a custom type struct.
// The control evaluating if policies have a minimum backup retention and frequency.
type PolicyControl struct {
    // The minimum retention duration for policy control.
    MinimumRetentionDuration *TimeUnitParamPolicyMinRetentionDuration `json:"minimum_retention_duration"`
    // The minimum RPO duration for policy control.
    MinimumRpoFrequency      *TimeUnitParamRpoDuration                `json:"minimum_rpo_frequency"`
}

// PolicyDetails represents a custom type struct
type PolicyDetails struct {
    // "description" is a Clumio assigned term. User can choose to ignore it.
    Description    *string     `json:"description"`
    // "name" is a Clumio assigned term. User can choose to ignore it.
    Name           *string     `json:"name"`
    // "policy_document" has stringified JSON blob. The user has to JSONify it and then paste
    // the JSONified blob in aws console while creating the policy.
    PolicyDocument interface{} `json:"policy_document"`
}

// PolicyEmbedded represents a custom type struct.
// If the `embed` query parameter is set, displays the responses of the related resource,
// as defined by the embeddable link specified.
type PolicyEmbedded struct {
    // Embeds the EBS protection statistics into the response.
    ReadPolicyAwsEbsVolumesProtectionStats interface{} `json:"read-policy-aws-ebs-volumes-protection-stats"`
}

// PolicyLinks represents a custom type struct.
// URLs to pages related to the resource.
type PolicyLinks struct {
    // The HATEOAS link to this resource.
    Self                   *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    DeletePolicyDefinition *HateoasLink     `json:"delete-policy-definition"`
    // A resource-specific HATEOAS link.
    UpdatePolicyDefinition *HateoasLink     `json:"update-policy-definition"`
}

// PolicyListEmbedded represents a custom type struct.
// An array of embedded resources related to this resource.
type PolicyListEmbedded struct {
    // TODO: Add struct field description
    Items []*Policy `json:"items"`
}

// PolicyListLinks represents a custom type struct.
// URLs to pages related to the resource.
type PolicyListLinks struct {
    // The HATEOAS link to the first page of results.
    First                  *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next                   *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self                   *HateoasSelfLink  `json:"_self"`
    // A resource-specific HATEOAS link.
    CreatePolicyDefinition *HateoasLink      `json:"create-policy-definition"`
}

// PolicyOperation represents a custom type struct
type PolicyOperation struct {
    // Determines whether the protection policy should take action now or during the specified backup window.
    // If set to `immediate`, Clumio starts the backup process right away. If set to `window`, Clumio starts the backup process when the backup window (`backup_window_tz`) opens.
    // If set to `window` and `operation in ("aws_rds_resource_aws_snapshot", "mssql_log_backup", "ec2_mssql_log_backup")`,
    // the backup window will not be determined by Clumio's backup window.
    ActionSetting     *string                 `json:"action_setting"`
    // Additional operation-specific policy settings. For operation types which do not support additional settings, this field is `null`.
    AdvancedSettings  *PolicyAdvancedSettings `json:"advanced_settings"`
    // The region in which this backup is stored. This might be used for cross-region backup.
    BackupAwsRegion   *string                 `json:"backup_aws_region"`
    // The start and end times of the customized backup window. Use of `backup_window` is deprecated, use `backup_window_tz` instead.
    BackupWindow      *BackupWindow           `json:"backup_window"`
    // The start and end times of the customized backup window. Use of `backup_window` is deprecated, use `backup_window_tz` instead.
    BackupWindowTz    *BackupWindow           `json:"backup_window_tz"`
    // The next start time of this operation in unix time.
    NextStartTime     *int64                  `json:"next_start_time"`
    // The previous start time of this operation in unix time.
    PreviousStartTime *int64                  `json:"previous_start_time"`
    // backup_sla captures the SLA parameters
    // backup_sla captures the SLA parameters
    Slas              []*BackupSLA            `json:"slas"`
    // The timezone for the operation. The timezone must be a valid location name from the IANA Time Zone database.
    // For instance, "America/New_York", "US/Central", "UTC".
    Timezone          *string                 `json:"timezone"`
    // The operation to be performed for this SLA set. Each SLA set corresponds to one and only one operation.
    // Refer to the Policy Operation table for a complete list of policy operations.
    ClumioType        *string                 `json:"type"`
}

// PolicyOperationInput represents a custom type struct
type PolicyOperationInput struct {
    // Determines whether the protection policy should take action now or during the specified backup window.
    // If set to `immediate`, Clumio starts the backup process right away. If set to `window`, Clumio starts the backup process when the backup window (`backup_window_tz`) opens.
    // If set to `window` and `operation in ("aws_rds_resource_aws_snapshot", "mssql_log_backup", "ec2_mssql_log_backup")`,
    // the backup window will not be determined by Clumio's backup window.
    ActionSetting    *string                 `json:"action_setting"`
    // Additional operation-specific policy settings. For operation types which do not support additional settings, this field is `null`.
    AdvancedSettings *PolicyAdvancedSettings `json:"advanced_settings"`
    // The region in which this backup is stored. This might be used for cross-region backup.
    BackupAwsRegion  *string                 `json:"backup_aws_region"`
    // The start and end times of the customized backup window. Use of `backup_window` is deprecated, use `backup_window_tz` instead.
    BackupWindow     *BackupWindow           `json:"backup_window"`
    // The start and end times of the customized backup window. Use of `backup_window` is deprecated, use `backup_window_tz` instead.
    BackupWindowTz   *BackupWindow           `json:"backup_window_tz"`
    // backup_sla captures the SLA parameters
    // backup_sla captures the SLA parameters
    Slas             []*BackupSLA            `json:"slas"`
    // The timezone for the operation. The timezone must be a valid location name from the IANA Time Zone database.
    // For instance, "America/New_York", "US/Central", "UTC".
    Timezone         *string                 `json:"timezone"`
    // The operation to be performed for this SLA set. Each SLA set corresponds to one and only one operation.
    // Refer to the Policy Operation table for a complete list of policy operations.
    ClumioType       *string                 `json:"type"`
}

// PrefixFilter represents a custom type struct.
// PrefixFilter
type PrefixFilter struct {
    // List of subprefixes to exclude from the prefix.
    ExcludedSubPrefixes []*string `json:"excluded_sub_prefixes"`
    // Prefix to include.
    Prefix              *string   `json:"prefix"`
}

// PreviewDetailsProtectionGroupLinks represents a custom type struct.
// URLs to pages related to the resource.
type PreviewDetailsProtectionGroupLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// PreviewDetailsS3BucketLinks represents a custom type struct.
// URLs to pages related to the resource.
type PreviewDetailsS3BucketLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// PreviewProtectionGroupAsyncLinks represents a custom type struct.
// URLs to pages related to the resource.
type PreviewProtectionGroupAsyncLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// PreviewProtectionGroupS3AssetAsyncLinks represents a custom type struct.
// URLs to pages related to the resource.
type PreviewProtectionGroupS3AssetAsyncLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// PreviewProtectionGroupS3AssetDetailsLinks represents a custom type struct.
// URLs to pages related to the resource.
type PreviewProtectionGroupS3AssetDetailsLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// PreviewProtectionGroupS3AssetSyncLinks represents a custom type struct.
// URLs to pages related to the resource.
type PreviewProtectionGroupS3AssetSyncLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// PreviewProtectionGroupSyncLinks represents a custom type struct.
// URLs to pages related to the resource.
type PreviewProtectionGroupSyncLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// PreviewS3BucketLinks represents a custom type struct.
// URLs to pages related to the resource.
type PreviewS3BucketLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// Projection represents a custom type struct.
// Represents attributes that are copied (projected) from the table into an index. These are in addition to the
// primary key attributes and index key attributes, which are automatically projected.
type Projection struct {
    // Represents the non-key attribute names which will be projected into the index.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), this must be empty if
    // 'projection_type' is ALL or KEYS_ONLY, and non-empty if 'projection_type' is INCLUDE.
    NonKeyAttributes []*string `json:"non_key_attributes"`
    // The set of attributes that are projected into the index. Valid Values: ALL, KEYS_ONLY, INCLUDE.
    ProjectionType   *string   `json:"projection_type"`
}

// ProtectConfig represents a custom type struct.
// The configuration of the Clumio Cloud Protect product for this connection.
// If this connection is not configured for Clumio Cloud Protect, then this field has a
// value of `null`.
type ProtectConfig struct {
    // The asset types supported on the current version of the feature
    AssetTypesEnabled        []*string     `json:"asset_types_enabled"`
    // EbsAssetInfo
    // The installed information for the EBS feature.
    Ebs                      *EbsAssetInfo `json:"ebs"`
    // The current version of the feature.
    InstalledTemplateVersion *string       `json:"installed_template_version"`
    // RdsAssetInfo
    // The installed information for the RDS feature.
    Rds                      *RdsAssetInfo `json:"rds"`
}

// ProtectEntitiesHateoasLink represents a custom type struct.
// A HATEOAS link to protect the entities.
type ProtectEntitiesHateoasLink struct {
    // The URI for the referenced operation.
    Href       *string `json:"href"`
    // Determines whether the "href" link is a URI template. If set to `true`, the "href" link is a URI template.
    Templated  *bool   `json:"templated"`
    // The HTTP method to be used with the "href" link for the referenced operation.
    ClumioType *string `json:"type"`
}

// ProtectionGroup represents a custom type struct
type ProtectionGroup struct {
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

// ProtectionGroupBackup represents a custom type struct
type ProtectionGroupBackup struct {
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
    // The number of objects in the protection group that were missing during backup.
    MissingObjectCount     *int64                      `json:"missing_object_count"`
    // The total size in bytes of objects in the protection group that were missing during backup.
    MissingSizeBytes       *int64                      `json:"missing_size_bytes"`
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

// ProtectionGroupBackupAdvancedSetting represents a custom type struct.
// Additional policy configuration settings for the `protection_group_backup` operation. If this operation is not of type `protection_group_backup`, then this field is omitted from the response.
type ProtectionGroupBackupAdvancedSetting struct {
    // Backup tier to store the backup in. Valid values are: `standard`, `archive`
    BackupTier *string `json:"backup_tier"`
}

// ProtectionGroupBackupLinks represents a custom type struct.
// URLs to pages related to the resource.
type ProtectionGroupBackupLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// ProtectionGroupBackupListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ProtectionGroupBackupListEmbedded struct {
    // TODO: Add struct field description
    Items []*ProtectionGroupBackup `json:"items"`
}

// ProtectionGroupBackupListLinks represents a custom type struct.
// URLs to pages related to the resource.
type ProtectionGroupBackupListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// ProtectionGroupBucket represents a custom type struct
type ProtectionGroupBucket struct {
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

// ProtectionGroupBucketContinuousBackupStats represents a custom type struct.
// ProtectionGroupBucketContinuousBackupStats
type ProtectionGroupBucketContinuousBackupStats struct {
    // The end time for the continuous backup stats in RFC-3339 format.
    BackupEndTime                    *string `json:"backup_end_time"`
    // The start time for the continuous backup stats in RFC-3339 format.
    BackupStartTime                  *string `json:"backup_start_time"`
    // The number of objects in the continuous backup task successfully deleted.
    DeletedObjectsCount              *uint64 `json:"deleted_objects_count"`
    // The total size in bytes of objects in the continuous backup task successfully deleted.
    DeletedObjectsSize               *uint64 `json:"deleted_objects_size"`
    // The number of failed continuous backup task executions.
    FailedContinuousBackupsCount     *uint32 `json:"failed_continuous_backups_count"`
    // The number of objects in the continuous backup task failed to be backed up.
    FailedObjectsCount               *uint64 `json:"failed_objects_count"`
    // The total size in bytes of objects in the continuous backup task failed to be backed up.
    FailedObjectsSize                *uint64 `json:"failed_objects_size"`
    // The number of included objects after the protection group filter.
    FilteredInCount                  *uint64 `json:"filtered_in_count"`
    // The total size in bytes of included objects after the protection group filter.
    FilteredInSize                   *uint64 `json:"filtered_in_size"`
    // The number of excluded objects after the protection group filter.
    FilteredOutCount                 *uint64 `json:"filtered_out_count"`
    // The total size in bytes of excluded objects after the protection group filter.
    FilteredOutSize                  *uint64 `json:"filtered_out_size"`
    // The number of objects in the continuous backup task missed to be backed up.
    MissingObjectsCount              *uint64 `json:"missing_objects_count"`
    // The total size in bytes of objects in the continuous backup task missed to be backed up.
    MissingObjectsSize               *uint64 `json:"missing_objects_size"`
    // The number of ongoing continuous backup task executions.
    OngoingContinuousBackupsCount    *uint32 `json:"ongoing_continuous_backups_count"`
    // The number of successful continuous backup task executions.
    SuccessfulContinuousBackupsCount *uint32 `json:"successful_continuous_backups_count"`
    // The number of objects in the continuous backup task successfully backed up.
    SuccessfulObjectsCount           *uint64 `json:"successful_objects_count"`
    // The total size in bytes of objects in the continuous backup task successfully backed up.
    SuccessfulObjectsSize            *uint64 `json:"successful_objects_size"`
    // The number of total continuous backup task executions.
    TotalContinuousBackupsCount      *uint32 `json:"total_continuous_backups_count"`
}

// ProtectionGroupBucketContinuousBackupStatsLinks represents a custom type struct.
// ProtectionGroupBucketContinuousBackupStatsLinks
// URLs to pages related to the resources.
type ProtectionGroupBucketContinuousBackupStatsLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// ProtectionGroupBucketEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ProtectionGroupBucketEmbedded struct {
    // This embed is for internal use only since an embed results in additional HTTP
    // calls. "embeds" can affect the performance of "list" API calls as an embed is
    // processed once per item in the result list.
    ReadOrganizationalUnit interface{} `json:"read-organizational-unit"`
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition   interface{} `json:"read-policy-definition"`
}

// ProtectionGroupBucketLinks represents a custom type struct.
// URLs to pages related to the resource.
type ProtectionGroupBucketLinks struct {
    // The HATEOAS link to this resource.
    Self                              *HateoasSelfLink                 `json:"_self"`
    // A resource-specific HATEOAS link.
    DeleteBucketProtectionGroup       *HateoasLink                     `json:"delete-bucket-protection-group"`
    // A resource-specific HATEOAS link.
    ListBackupProtectionGroupS3Assets *HateoasLink                     `json:"list-backup-protection-group-s3-assets"`
    // A resource-specific HATEOAS link.
    ReadOrganizationalUnit            *HateoasLink                     `json:"read-organizational-unit"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition              *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
}

// ProtectionGroupBucketListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ProtectionGroupBucketListEmbedded struct {
    // TODO: Add struct field description
    Items []*ProtectionGroupBucket `json:"items"`
}

// ProtectionGroupBucketListLinks represents a custom type struct.
// URLs to pages related to the resource.
type ProtectionGroupBucketListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// ProtectionGroupContinuousBackupAdvancedSetting represents a custom type struct.
// Additional policy configuration settings for the `protection_group_continuous_backup` operation. If this operation is not of type `protection_group_continuous_backup`, then this field is omitted from the response.
type ProtectionGroupContinuousBackupAdvancedSetting struct {
    // If true, tries to disable EventBridge notification for the given protection group, when continuous backup no longer conducts.
    // It may override the existing S3 bucket notification configuration in the customer's account.
    // This takes effect only when `event_bridge_enabled` is set to `false`.
    DisableEventbridgeNotification *bool `json:"disable_eventbridge_notification"`
}

// ProtectionGroupEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ProtectionGroupEmbedded struct {
    // This embed is for internal use only since an embed results in additional HTTP
    // calls. "embeds" can affect the performance of "list" API calls as an embed is
    // processed once per item in the result list.
    ReadOrganizationalUnit interface{} `json:"read-organizational-unit"`
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition   interface{} `json:"read-policy-definition"`
}

// ProtectionGroupLinks represents a custom type struct.
// URLs to pages related to the resource.
type ProtectionGroupLinks struct {
    // The HATEOAS link to this resource.
    Self                        *HateoasSelfLink                 `json:"_self"`
    // A resource-specific HATEOAS link.
    AddBucketProtectionGroup    *HateoasLink                     `json:"add-bucket-protection-group"`
    // A resource-specific HATEOAS link.
    DeleteBucketProtectionGroup *HateoasLink                     `json:"delete-bucket-protection-group"`
    // A resource-specific HATEOAS link.
    ListBackupProtectionGroups  *HateoasLink                     `json:"list-backup-protection-groups"`
    // A resource-specific HATEOAS link.
    ReadOrganizationalUnit      *HateoasLink                     `json:"read-organizational-unit"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition        *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
    // A resource-specific HATEOAS link.
    UpdateProtectionGroup       *HateoasLink                     `json:"update-protection-group"`
}

// ProtectionGroupListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ProtectionGroupListEmbedded struct {
    // TODO: Add struct field description
    Items []*ProtectionGroup `json:"items"`
}

// ProtectionGroupListLinks represents a custom type struct.
// URLs to pages related to the resource.
type ProtectionGroupListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// ProtectionGroupRestoreSource represents a custom type struct.
// The parameters for initiating a protection group restore from a backup.
type ProtectionGroupRestoreSource struct {
    // The Clumio-assigned ID of the protection group backup to be restored. Use the
    // [GET /backups/protection-groups](#operation/list-backup-protection-groups)
    // endpoint to fetch valid values. 
    // Note that only one of `backup_id` or `pitr` must be given.
    BackupId                  *string                                  `json:"backup_id"`
    // Search for or restore only objects that pass the source object filter.
    ObjectFilters             *SourceObjectFilters                     `json:"object_filters"`
    // The parameters for initiating a point in time restore.
    // Note that only one of `backup_id` or `pitr` must be given.
    Pitr                      *ProtectionGroupRestoreSourcePitrOptions `json:"pitr"`
    // A list of Clumio-assigned IDs of protection group S3 assets, representing the
    // buckets within the protection group to restore from. Use the
    // [GET /datasources/protection-groups/s3-assets](#operation/list-protection-group-s3-assets)
    // endpoint to fetch valid values.
    ProtectionGroupS3AssetIds []*string                                `json:"protection_group_s3_asset_ids"`
}

// ProtectionGroupRestoreSourcePitrOptions represents a custom type struct.
// The parameters for initiating a point in time restore.<br>
// Note that only one of `backup_id` or `pitr` must be given.
type ProtectionGroupRestoreSourcePitrOptions struct {
    // Clumio-assigned ID of protection group, representing the
    // protection group to restore from. Use the
    // [GET /datasources/protection-groups](#operation/list-protection-group)
    // endpoint to fetch valid values.
    ProtectionGroupId     *string `json:"protection_group_id"`
    // The end timestamp to be restored in RFC-3339 format.
    // Clumio restores last objects modified before the given time. 
    // If `restore_end_timestamp` is given without `restore_start_timestamp`,
    // it is the same as point in time restore.
    RestoreEndTimestamp   *string `json:"restore_end_timestamp"`
    // The start timestamp to be restored in RFC-3339 format.
    // Clumio restores objects modified since the given time.
    RestoreStartTimestamp *string `json:"restore_start_timestamp"`
}

// ProtectionGroupRestoreTarget represents a custom type struct.
// The destination where the protection group will be restored.
type ProtectionGroupRestoreTarget struct {
    // The Clumio-assigned ID of the bucket to which the backup must be restored.
    // Use the [GET /datasources/aws/s3-buckets](#operation/list-aws-s3-buckets) endpoint
    // to fetch valid values.
    BucketId                       *string              `json:"bucket_id"`
    // Default AWS checksum algorithm for restored object.
    // Valid values are: `CRC32`, `CRC32C`, `CRC64NVME`, `SHA1`, and `SHA256. 
    // Note that this will be applied when backup didn't have checksum algorithm information.
    DefaultObjectChecksumAlgorithm *string              `json:"default_object_checksum_algorithm"`
    // The Clumio-assigned ID of the AWS environment to be used as the restore destination.
    // Use the [GET /datasources/aws/s3-buckets/{bucket_id}](#operation/read-aws-s3-bucket) endpoint
    // to fetch the environment ID for a bucket.
    EnvironmentId                  *string              `json:"environment_id"`
    // If overwrite is set to true, we will overwrite an object if it exists. If it's set to false,
    // then we will fail the restore if an object already exists.
    Overwrite                      *bool                `json:"overwrite"`
    // Prefix to restore the objects under. If more than one bucket is restored, the
    // bucket name will be appended to the prefix.
    Prefix                         *string              `json:"prefix"`
    // Whether to restore objects with their original storage class or not. 
    // If it is `true`, `storage_class` must be empty.
    // Otherwise, `storage_class` must be given.
    RestoreOriginalStorageClass    *bool                `json:"restore_original_storage_class"`
    // Storage class for restored objects. Valid values are: `S3 Standard`, `S3 Standard-IA`,
    // `S3 Intelligent-Tiering` and `S3 One Zone-IA`. 
    // Note that this must be given unless `restore_original_storage_class` is `true`.
    StorageClass                   *string              `json:"storage_class"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags                           []*AwsTagCommonModel `json:"tags"`
}

// ProtectionGroupS3AssetBackup represents a custom type struct
type ProtectionGroupS3AssetBackup struct {
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
    // The number of objects in the protection group S3 asset that were missing during backup.
    MissingObjectCount       *int64                             `json:"missing_object_count"`
    // The total size in bytes of objects in the protection group S3 asset that were missing during backup.
    MissingSizeBytes         *int64                             `json:"missing_size_bytes"`
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

// ProtectionGroupS3AssetBackupLinks represents a custom type struct.
// URLs to pages related to the resource.
type ProtectionGroupS3AssetBackupLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// ProtectionGroupS3AssetBackupListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ProtectionGroupS3AssetBackupListEmbedded struct {
    // TODO: Add struct field description
    Items []*ProtectionGroupS3AssetBackup `json:"items"`
}

// ProtectionGroupS3AssetBackupListLinks represents a custom type struct.
// URLs to pages related to the resource.
type ProtectionGroupS3AssetBackupListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// ProtectionGroupS3AssetPitrInterval represents a custom type struct
type ProtectionGroupS3AssetPitrInterval struct {
    // The end time of the interval, represented in RFC3339 format.
    EndTimestamp   *string `json:"end_timestamp"`
    // The start time of the interval, represented in RFC3339 format.
    StartTimestamp *string `json:"start_timestamp"`
}

// ProtectionGroupS3AssetPitrIntervalListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ProtectionGroupS3AssetPitrIntervalListEmbedded struct {
    // TODO: Add struct field description
    Items []*ProtectionGroupS3AssetPitrInterval `json:"items"`
}

// ProtectionGroupS3AssetPitrIntervalListLinks represents a custom type struct.
// URLs to pages related to the resource.
type ProtectionGroupS3AssetPitrIntervalListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// ProtectionGroupS3AssetRestoreSource represents a custom type struct.
// The parameters for initiating a protection group S3 asset restore
// or creation of an instant access endpoint from a backup.
type ProtectionGroupS3AssetRestoreSource struct {
    // The Clumio-assigned ID of the protection group S3 asset backup to be restored. Use the
    // [GET /backups/protection-groups/s3-assets](#operation/list-backup-protection-group-s3-assets)
    // endpoint to fetch valid values. 
    // Note that only one of `backup_id` or `pitr` must be given.
    BackupId      *string                                         `json:"backup_id"`
    // Search for or restore only objects that pass the source object filter.
    ObjectFilters *SourceObjectFilters                            `json:"object_filters"`
    // The parameters for initiating a point in time restore.
    // Note that only one of `backup_id` or `pitr` must be given.
    Pitr          *ProtectionGroupS3AssetRestoreSourcePitrOptions `json:"pitr"`
}

// ProtectionGroupS3AssetRestoreSourcePitrOptions represents a custom type struct.
// The parameters for initiating a point in time restore.<br>
// Note that only one of `backup_id` or `pitr` must be given.
type ProtectionGroupS3AssetRestoreSourcePitrOptions struct {
    // Clumio-assigned ID of protection group S3 asset, representing the
    // bucket within the protection group to restore from. Use the
    // [GET /datasources/protection-groups/s3-assets](#operation/list-protection-group-s3-assets)
    // endpoint to fetch valid values.
    ProtectionGroupS3AssetId *string `json:"protection_group_s3_asset_id"`
    // The ending time to be restored in RFC-3339 format.
    // We will restore last objects modified before the given time. 
    // If `restore_end_timestamp` is given without `restore_start_timestamp`,
    // it is the same as point in time restore.
    RestoreEndTimestamp      *string `json:"restore_end_timestamp"`
    // The starting time to be restored in RFC-3339 format.
    // We will restore objects modified since the given time.
    RestoreStartTimestamp    *string `json:"restore_start_timestamp"`
}

// ProtectionGroupVersionLinks represents a custom type struct.
// URLs to pages related to the resource.
type ProtectionGroupVersionLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// ProtectionInfo represents a custom type struct.
// The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
type ProtectionInfo struct {
    // The ID of the entity from which protection was inherited.
    // If protection was not inherited, then this field has a value of `null`.
    InheritingEntityId   *string `json:"inheriting_entity_id"`
    // The type of entity from which protection was inherited.
    // If protection was not inherited, then this field has a value of `null`.
    // Entities from which protection can be inherited include the following:
    // 
    // +------------------------+----------+
    // | Inheriting Entity Type | Details  |
    // +========================+==========+
    // | aws_tag                | AWS tag. |
    // +------------------------+----------+
    // 
    InheritingEntityType *string `json:"inheriting_entity_type"`
    // A system-generated ID assigned to the policy protecting this resource.
    PolicyId             *string `json:"policy_id"`
}

// ProtectionInfoWithRule represents a custom type struct.
// The protection policy applied to this resource. If the resource is not protected, then this field has a value of `null`.
type ProtectionInfoWithRule struct {
    // The ID of the entity from which protection was inherited.
    // If protection was not inherited, then this field has a value of `null`.
    InheritingEntityId   *string `json:"inheriting_entity_id"`
    // The type of entity from which protection was inherited.
    // If protection was not inherited, then this field has a value of `null`.
    // Entities from which protection can be inherited include the following:
    // 
    // +------------------------+----------+
    // | Inheriting Entity Type | Details  |
    // +========================+==========+
    // | aws_tag                | AWS tag. |
    // +------------------------+----------+
    // 
    InheritingEntityType *string `json:"inheriting_entity_type"`
    // A system-generated ID assigned to the policy protecting this resource.
    PolicyId             *string `json:"policy_id"`
}

// ProtectionStats represents a custom type struct
type ProtectionStats struct {
    // The total number of entities associated with deactivated policies.
    DeactivatedCount *int64 `json:"deactivated_count"`
    // The number of entities with protection applied.
    ProtectedCount   *int64 `json:"protected_count"`
    // The number of entities without protection applied.
    UnprotectedCount *int64 `json:"unprotected_count"`
}

// ProvisionedThroughput represents a custom type struct.
// Represents the provisioned throughput settings for a DynamoDB table.
type ProvisionedThroughput struct {
    // The maximum number of strongly consistent reads consumed per second.
    ReadCapacityUnits  *int64 `json:"read_capacity_units"`
    // The maximum number of writes consumed per second.
    WriteCapacityUnits *int64 `json:"write_capacity_units"`
}

// ProvisionedThroughputOverride represents a custom type struct.
// Replica-specific provisioned throughput settings. If not specified, uses the source table's provisioned throughput settings.
type ProvisionedThroughputOverride struct {
    // The maximum number of strongly consistent reads consumed per second.
    ReadCapacityUnits *int64 `json:"read_capacity_units"`
}

// RDSBackupDatabase represents a custom type struct
type RDSBackupDatabase struct {
    // The name of the database.
    Name *string `json:"name"`
}

// RDSBackupDatabaseListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type RDSBackupDatabaseListEmbedded struct {
    // TODO: Add struct field description
    Items []*RDSBackupDatabase `json:"items"`
}

// RDSBackupDatabaseListLinks represents a custom type struct.
// URLs to pages related to the resource.
type RDSBackupDatabaseListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// RDSConfigSyncAdvancedSetting represents a custom type struct.
// Advanced settings for RDS PITR configuration sync.
type RDSConfigSyncAdvancedSetting struct {
    // Syncs the configuration of RDS PITR in AWS. Valid values are: `immediate`, and `maintenance_window`.
    // Note that applying the setting "immediately" may cause unexpected downtime.
    Apply *string `json:"apply"`
}

// RDSDatabaseTable represents a custom type struct
type RDSDatabaseTable struct {
    // Embedded responses related to the resource.
    Embedded *RDSDatabaseTableEmbedded `json:"_embedded"`
    // URLs to pages related to the resource.
    Links    *RDSDatabaseTableLinks    `json:"_links"`
    // The name of the table within the specified RDS database.
    Name     *string                   `json:"name"`
}

// RDSDatabaseTableColumn represents a custom type struct.
// RDSDatabaseTableColumn denotes the model for rds database column
type RDSDatabaseTableColumn struct {
    // The name of the column.
    Name       *string `json:"name"`
    // The Hive data type of the column. Possible values include `int`, `bigint`, `string`, and `boolean`.
    ClumioType *string `json:"type"`
}

// RDSDatabaseTableColumnLinks represents a custom type struct.
// URLs to pages related to the resource.
type RDSDatabaseTableColumnLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// RDSDatabaseTableEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type RDSDatabaseTableEmbedded struct {
    // Add resource specific HATEOAS embedded
    ReadBackupAwsRdsResourceDatabaseTableColumns interface{} `json:"read-backup-aws-rds-resource-database-table-columns"`
}

// RDSDatabaseTableLinks represents a custom type struct.
// URLs to pages related to the resource.
type RDSDatabaseTableLinks struct {
    // The HATEOAS link to this resource.
    Self                                         *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    ReadBackupAwsRdsResourceDatabaseTableColumns *HateoasLink     `json:"read-backup-aws-rds-resource-database-table-columns"`
}

// RDSDatabaseTableListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type RDSDatabaseTableListEmbedded struct {
    // TODO: Add struct field description
    Items []*RDSDatabaseTable `json:"items"`
}

// RDSDatabaseTableListLinks represents a custom type struct.
// URLs to pages related to the resource.
type RDSDatabaseTableListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// RDSLogicalBackupAdvancedSetting represents a custom type struct.
// Settings for determining if a RDS policy is created with standard or archive tier.
type RDSLogicalBackupAdvancedSetting struct {
    // Backup tier to store the backup in. Valid values are: `standard`, `archive`
    BackupTier *string `json:"backup_tier"`
}

// RDSLogicalPreviewQueryResult represents a custom type struct.
// The preview of the query result, if `preview:true` in the request.
// If preview was not set to true in the request, then the result of the query will be
// available for download asynchronously.
type RDSLogicalPreviewQueryResult struct {
    // RDSDatabaseTableColumn denotes the model for rds database column
    Columns []*RDSDatabaseTableColumn `json:"columns"`
    // The rows of the previewed query result.
    Rows    []*[]string               `json:"rows"`
}

// RPOBackupSLAParam represents a custom type struct.
// The minimum frequency between backups for this SLA. Also known as the recovery point objective (RPO) interval.
// For example, to configure the minimum frequency between backups to be every 2 days, set `unit="days"` and `value=2`.
// To configure the SLA for on-demand backups, set `unit="on_demand"` and leave the `value` field empty.
type RPOBackupSLAParam struct {
    // The weekday in decimal of the Weekly SLA parameter. Valid values are integers from 0 to 6, indicates Sunday, Monday, ..., Saturday. For example, to configure backup on every Monday, set `unit="weekly"`, `value=1`, and `offsets={1}`.
    Offsets []*int64 `json:"offsets"`
    // The measurement unit of the SLA parameter.
    Unit    *string  `json:"unit"`
    // The measurement value of the SLA parameter.
    Value   *int64   `json:"value"`
}

// RdsAssetInfo represents a custom type struct.
// RdsAssetInfo
// The installed information for the RDS feature.
type RdsAssetInfo struct {
    // The current version of the feature.
    InstalledTemplateVersion *string `json:"installed_template_version"`
}

// RdsDatabaseBackup represents a custom type struct
type RdsDatabaseBackup struct {
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

// RdsDatabaseBackupLinks represents a custom type struct.
// URLs to pages related to the resource.
type RdsDatabaseBackupLinks struct {
    // The HATEOAS link to this resource.
    Self                            *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    ListAwsRdsResourcesOptionGroups *HateoasLink     `json:"list-aws-rds-resources-option-groups"`
    // A resource-specific HATEOAS link.
    RestoreAwsRdsResource           *HateoasLink     `json:"restore-aws-rds-resource"`
    // A resource-specific HATEOAS link.
    RestoreRdsRecord                *HateoasLink     `json:"restore-rds-record"`
}

// RdsDatabaseBackupListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type RdsDatabaseBackupListEmbedded struct {
    // TODO: Add struct field description
    Items []*RdsDatabaseBackup `json:"items"`
}

// RdsDatabaseBackupListLinks represents a custom type struct.
// URLs to pages related to the resource.
type RdsDatabaseBackupListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// RdsInstanceModel represents a custom type struct
type RdsInstanceModel struct {
    // The class of the RDS instance at the time of backup. Possible values include `db.r5.2xlarge` and `db.t2.small`.
    // For a full list of possible values, refer to the Amazon documentation at
    // https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html.
    Class                *string `json:"class"`
    // Determines whether the RDS instance had a public IP address in addition to the private IP address at the time of backup.
    // For more information about the public access option, refer to the Amazon documentation at
    // https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html.
    IsPubliclyAccessible *bool   `json:"is_publicly_accessible"`
    // The AWS-assigned name of the RDS instance at the time of backup.
    Name                 *string `json:"name"`
}

// RdsResource represents a custom type struct
type RdsResource struct {
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

// RdsResourceEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type RdsResourceEmbedded struct {
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition interface{} `json:"read-policy-definition"`
}

// RdsResourceLinks represents a custom type struct.
// URLs to pages related to the resource.
type RdsResourceLinks struct {
    // The HATEOAS link to this resource.
    Self                      *HateoasSelfLink                 `json:"_self"`
    // A resource-specific HATEOAS link.
    ListBackupAwsRdsResources *HateoasLink                     `json:"list-backup-aws-rds-resources"`
    // A resource-specific HATEOAS link.
    ListRdsRestoredRecords    *HateoasLink                     `json:"list-rds-restored-records"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition      *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
    // A resource-specific HATEOAS link.
    RestoreAwsRdsResource     *HateoasLink                     `json:"restore-aws-rds-resource"`
}

// RdsResourceListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type RdsResourceListEmbedded struct {
    // TODO: Add struct field description
    Items []*RdsResource `json:"items"`
}

// RdsResourceListLinks represents a custom type struct.
// URLs to pages related to the resource.
type RdsResourceListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// RdsResourceRestoreSource represents a custom type struct.
// The RDS resource backup or snapshot to be restored.  Only one of these fields should be set.
type RdsResourceRestoreSource struct {
    // The parameters for initiating an RDS restore from a backup.
    Backup   *RdsResourceRestoreSourceAirGapOptions `json:"backup"`
    // The parameters for initiating an RDS restore from a snapshot.
    Snapshot *RdsResourceRestoreSourcePitrOptions   `json:"snapshot"`
}

// RdsResourceRestoreSourceAirGapOptions represents a custom type struct.
// The parameters for initiating an RDS restore from a backup.
type RdsResourceRestoreSourceAirGapOptions struct {
    // The Clumio-assigned ID of the RDS backup to be restored.
    // Use the [GET /backups/aws/rds-resources](#operation/list-backup-aws-rds-resources)
    // endpoint to fetch valid values.
    BackupId *string `json:"backup_id"`
}

// RdsResourceRestoreSourcePitrOptions represents a custom type struct.
// The parameters for initiating an RDS restore from a snapshot.
type RdsResourceRestoreSourcePitrOptions struct {
    // The Clumio-assigned ID of the RDS resource to be restored.
    // Use the [GET /datasources/aws/rds-resources](#operation/list-aws-rds-resources)
    // endpoint to fetch valid values.
    ResourceId *string `json:"resource_id"`
    // A point in time to be restored in RFC-3339 format.
    Timestamp  *string `json:"timestamp"`
}

// RdsResourceRestoreTarget represents a custom type struct.
// The configuration of the RDS resource to be restored.
type RdsResourceRestoreTarget struct {
    // The Clumio-assigned ID of the AWS environment to be used as the restore destination.
    // Use the [GET /datasources/aws/environments](#operation/list-aws-environments) endpoint to fetch valid values.
    EnvironmentId          *string              `json:"environment_id"`
    // The instance class of the RDS resources to be created. Possible values include `db.r5.2xlarge` and `db.t2.small`.
    InstanceClass          *string              `json:"instance_class"`
    // Designates whether the restored RDS resource also has a public IP address in addition to the private IP address.
    IsPubliclyAccessible   *bool                `json:"is_publicly_accessible"`
    // The AWS-assigned ID of the KMS encryption key used to encrypt data in this RDS resource.
    KmsKeyNativeId         *string              `json:"kms_key_native_id"`
    // The name given to the restored RDS resource.
    // The name must follow AWS RDS naming conventions:
    // https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Limits.html#RDS_Limits.Constraints
    Name                   *string              `json:"name"`
    // Option group name to be added to the restored RDS resource.
    OptionGroupName        *string              `json:"option_group_name"`
    // The name of the parameter group to be associated with the restored RDS resource.
    ParameterGroupName     *string              `json:"parameter_group_name"`
    // The AWS-assigned IDs of the security groups to be associated with the restored RDS resource.
    SecurityGroupNativeIds []*string            `json:"security_group_native_ids"`
    // The AWS-assigned name of the subnet group to be associated with the restored RDS resource.
    SubnetGroupName        *string              `json:"subnet_group_name"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags                   []*AwsTagCommonModel `json:"tags"`
}

// RdsTemplateInfo represents a custom type struct
type RdsTemplateInfo struct {
    // The latest available feature version for the asset.
    AvailableTemplateVersion *string `json:"available_template_version"`
}

// ReadPolicyDefinitionHateoasLink represents a custom type struct.
// A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
type ReadPolicyDefinitionHateoasLink struct {
    // The URI for the referenced operation.
    Href       *string `json:"href"`
    // Determines whether the "href" link is a URI template. If set to `true`, the "href" link is a URI template.
    Templated  *bool   `json:"templated"`
    // The HTTP method to be used with the "href" link for the referenced operation.
    ClumioType *string `json:"type"`
}

// ReadTaskHateoasLink represents a custom type struct.
// A HATEOAS link to the task associated with this resource.
type ReadTaskHateoasLink struct {
    // The URI for the referenced operation.
    Href       *string `json:"href"`
    // Determines whether the "href" link is a URI template. If set to `true`, the "href" link is a URI template.
    Templated  *bool   `json:"templated"`
    // The HTTP method to be used with the "href" link for the referenced operation.
    ClumioType *string `json:"type"`
}

// ReadTaskHateoasOuterEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type ReadTaskHateoasOuterEmbedded struct {
    // Embeds the associated task of a resource in the response if requested using the `embed` query parameter.
    ReadTask interface{} `json:"read-task"`
}

// ReplicaDescription represents a custom type struct.
// Contains the details of the replica.
type ReplicaDescription struct {
    // Represents the properties of a replica global secondary index.
    GlobalSecondaryIndexes        []*ReplicaGlobalSecondaryIndexDescription `json:"global_secondary_indexes"`
    // The AWS KMS key of the replica that will be used for AWS KMS encryption.
    KmsMasterKeyId                *string                                   `json:"kms_master_key_id"`
    // Replica-specific ondemand throughput settings. If not specified, uses the source table's ondemand throughput settings.
    OnDemandThroughputOverride    *OnDemandThroughputOverride               `json:"on_demand_throughput_override"`
    // Replica-specific provisioned throughput settings. If not specified, uses the source table's provisioned throughput settings.
    ProvisionedThroughputOverride *ProvisionedThroughputOverride            `json:"provisioned_throughput_override"`
    // The name of the Region.
    RegionName                    *string                                   `json:"region_name"`
    // Replica-specific table class summary.
    // Possible values are `STANDARD` or `STANDARD_INFREQUENT_ACCESS`. This is defaulted to the
    // `STANDARD` storage class if empty.
    ReplicaTableClass             *string                                   `json:"replica_table_class"`
}

// ReplicaGlobalSecondaryIndexDescription represents a custom type struct.
// Represents the properties of a replica global secondary index.
type ReplicaGlobalSecondaryIndexDescription struct {
    // The name of the global secondary index.
    IndexName                     *string                        `json:"index_name"`
    // Replica-specific ondemand throughput settings. If not specified, uses the source table's ondemand throughput settings.
    OnDemandThroughputOverride    *OnDemandThroughputOverride    `json:"on_demand_throughput_override"`
    // Replica-specific provisioned throughput settings. If not specified, uses the source table's provisioned throughput settings.
    ProvisionedThroughputOverride *ProvisionedThroughputOverride `json:"provisioned_throughput_override"`
}

// ReportDownload represents a custom type struct
type ReportDownload struct {
    // The link to the actual CSV report.
    DownloadLink        *string `json:"download_link"`
    // The time when the request was completed.
    EndTimestamp        *string `json:"end_timestamp"`
    // The time when this report CSV will expire and not be available for download.
    ExpirationTimestamp *string `json:"expiration_timestamp"`
    // The name of CSV file.
    FileName            *string `json:"file_name"`
    // The filters applied to the report when download was initiated.
    Filters             *string `json:"filters"`
    // The id of the report that uniquely identifies the report.
    Id                  *string `json:"id"`
    // The time when the request was made.
    StartTimestamp      *string `json:"start_timestamp"`
    // The Clumio-assigned ID of the task which generated the restored file.
    TaskId              *string `json:"task_id"`
    // The type of report this CSV Download is associated with.
    // The possible values: ["activity"].
    ClumioType          *string `json:"type"`
}

// ReportDownloadLinks represents a custom type struct.
// _links provides URLs to the related resources of a report CSV download
type ReportDownloadLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// ReportDownloadListEmbedded represents a custom type struct.
// _embedded contains the list of items of a list report CSV download query
type ReportDownloadListEmbedded struct {
    // TODO: Add struct field description
    Items []*ReportDownload `json:"items"`
}

// ReportDownloadListLinks represents a custom type struct.
// _links provides URLs to related navigable pages of a list report CSV download query
type ReportDownloadListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// Resources represents a custom type struct.
// Partial updates are not supported, therefore you must provide ARNs for all configured resources,
// including those for resources that are not being updated.
type Resources struct {
    // SNS topic created in the account to receive relevant events.
    ClumioEventPubArn       *string                  `json:"clumio_event_pub_arn"`
    // ARN of the IAM role created in the account, which will be assumed by Clumio.
    ClumioIamRoleArn        *string                  `json:"clumio_iam_role_arn"`
    // ARN of the support role which will be used by the Clumio support team.
    ClumioSupportRoleArn    *string                  `json:"clumio_support_role_arn"`
    // TODO: Add struct field description
    EventRules              *EventRules              `json:"event_rules"`
    // TODO: Add struct field description
    ServiceInstanceProfiles *ServiceInstanceProfiles `json:"service_instance_profiles"`
    // TODO: Add struct field description
    ServiceRoles            *ServiceRoles            `json:"service_roles"`
}

// RestEntity represents a custom type struct
type RestEntity struct {
    // A system-generated ID assigned to this entity.
    Id         *string `json:"id"`
    // Type is mostly an asset type or the type of Entity. Some examples are
    // "restored_file", "aws_ebs_volume",  etc.
    ClumioType *string `json:"type"`
    // A system-generated value assigned to the entity. For example, if the primary entity type is "aws_ebs_volume", then the value is the name of the EBS.
    Value      *string `json:"value"`
}

// RestoreDynamoDBTableLinks represents a custom type struct.
// URLs to pages related to the resource.
type RestoreDynamoDBTableLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// RestoreEBSLinks represents a custom type struct.
// URLs to pages related to the resource.
type RestoreEBSLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// RestoreEC2Links represents a custom type struct.
// URLs to pages related to the resource.
type RestoreEC2Links struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// RestoreFileLinks represents a custom type struct.
// URLs to pages related to the resource.
type RestoreFileLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// RestoreObjectsLinks represents a custom type struct.
// URLs to pages related to the resource.
type RestoreObjectsLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// RestoreProtectionGroupLinks represents a custom type struct.
// URLs to pages related to the resource.
type RestoreProtectionGroupLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// RestoreProtectionGroupS3AssetLinks represents a custom type struct.
// URLs to pages related to the resource.
type RestoreProtectionGroupS3AssetLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// RestoreRecordsLinksAsync represents a custom type struct.
// URLs to pages related to the resource.
type RestoreRecordsLinksAsync struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// RestoreRecordsLinksSync represents a custom type struct.
// URLs to pages related to the resource.
type RestoreRecordsLinksSync struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// RestoreS3BucketLinks represents a custom type struct.
// URLs to pages related to the resource.
type RestoreS3BucketLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// RestoreS3BucketSource represents a custom type struct.
// The parameters for initiating an S3 bucket restore.
type RestoreS3BucketSource struct {
    // The Clumio-assigned ID of the bucket to be restored. Use the
    // [GET /datasources/aws/s3-buckets](#operation/list-aws-s3-buckets) endpoint
    // to fetch valid values.
    BucketId                *string                                       `json:"bucket_id"`
    // Search for or restore only objects that pass the source object filter.
    ObjectFilters           *SourceObjectFiltersV2                        `json:"object_filters"`
    // This field is required when the request type is 'Rollback'. The parameters for initiating a point in time restore.
    PointInTimeRecovery     *RestoreS3BucketSourcePitrOptions             `json:"point_in_time_recovery"`
    // This field is required when the request type is 'Undo delete marker'.
    UndoDeleteMarkerOptions *RestoreS3BucketSourceUndoDeleteMarkerOptions `json:"undo_delete_marker_options"`
}

// RestoreS3BucketSourcePitrOptions represents a custom type struct.
// This field is required when the request type is 'Rollback'. The parameters for initiating a point in time restore.
type RestoreS3BucketSourcePitrOptions struct {
    // The end timestamp to be restored in RFC-3339 format.
    // Clumio restores objects modified before the given timestamp.
    RestoreEndTimestamp *string `json:"restore_end_timestamp"`
}

// RestoreS3BucketSourceUndoDeleteMarkerOptions represents a custom type struct.
// This field is required when the request type is 'Undo delete marker'.
type RestoreS3BucketSourceUndoDeleteMarkerOptions struct {
    // The start timestamp for undo delete marker in RFC-3339 format.
    // Clumio undoes the delete markers created after the given timestamp.
    StartTimestamp *string `json:"start_timestamp"`
}

// RestoreS3BucketTarget represents a custom type struct.
// The destination where the S3 bucket will be restored.
type RestoreS3BucketTarget struct {
    // The name of the destination bucket.
    BucketName    *string              `json:"bucket_name"`
    // The Clumio-assigned ID of the AWS environment to be used as the restore destination. Use the
    // [GET /datasources/aws/s3-buckets/{bucket_id}](#operation/read-aws-s3-bucket) endpoint
    // to fetch the environment ID for a bucket.
    EnvironmentId *string              `json:"environment_id"`
    // Prefix to restore the objects under.
    Prefix        *string              `json:"prefix"`
    // Storage class for restored objects. Valid values are: `S3 Standard`, `S3 Standard-IA`,
    // `S3 Intelligent-Tiering` and `S3 One Zone-IA`. 
    // If not given or empty, objects are restored with their original storage classes.
    StorageClass  *string              `json:"storage_class"`
    // A tag created through AWS Console which can be applied to EBS volumes.
    Tags          []*AwsTagCommonModel `json:"tags"`
}

// RestoredFileInfo represents a custom type struct
type RestoredFileInfo struct {
    // Specifies information about the data-access method for accessing the restored
    // file.
    AccessMethods       []*DataAccessObject `json:"access_methods"`
    // The Clumio-assigned ID of the backup associated with this restored file.
    BackupId            *string             `json:"backup_id"`
    // The timestamp of the when the backup associated with this file started.
    // Represented in RFC-3339 format.
    BackupTimestamp     *string             `json:"backup_timestamp"`
    // The timestamp of when the restored file will expire. Represented in RFC-3339
    // format.
    ExpirationTimestamp *string             `json:"expiration_timestamp"`
    // The Clumio-assigned ID of the restored file.
    Id                  *string             `json:"id"`
    // The Clumio-assigned name of the restored file.
    Name                *string             `json:"name"`
    // The size of the restored file. Measured in bytes (B).
    Size                *int64              `json:"size"`
    // The Clumio-assigned ID of the task which generated the restored file.
    TaskId              *string             `json:"task_id"`
}

// RestoredFilesListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type RestoredFilesListEmbedded struct {
    // TODO: Add struct field description
    Items []*RestoredFileInfo `json:"items"`
}

// RestoredFilesListLinks represents a custom type struct.
// URLs to pages related to the resource.
type RestoredFilesListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// RestoredRecord represents a custom type struct
type RestoredRecord struct {
    // URLs to pages related to the resource.
    Links               *RestoredRecordLinks `json:"_links"`
    // The AWS-assigned ID of the account with this record.
    AccountNativeId     *string              `json:"account_native_id"`
    // The AWS region associated with this record. For example, `us-west-2`.
    AwsRegion           *string              `json:"aws_region"`
    // The Clumio-assigned ID of the backup associated with this record.
    BackupId            *string              `json:"backup_id"`
    // The AWS-assigned name of the database associated with this record.
    DatabaseName        *string              `json:"database_name"`
    // The download link of the query result.
    DownloadLink        *string              `json:"download_link"`
    // The timestamp of when the record will expire. Represented in RFC-3339 format.
    ExpirationTimestamp *string              `json:"expiration_timestamp"`
    // The Clumio-assigned ID of the restored record.
    Id                  *string              `json:"id"`
    // The SQL query statement which produced this record.
    QueryStatement      *string              `json:"query_statement"`
    // The Clumio-assigned ID of the RDS resource associated with this record.
    ResourceId          *string              `json:"resource_id"`
    // The number of rows produced by the query.
    RowCount            *int64               `json:"row_count"`
    // The timestamp of when the query was executed. Represented in RFC-3339 format.
    StartTimestamp      *string              `json:"start_timestamp"`
    // The Clumio-assigned ID of the task which generated the restored record.
    TaskId              *string              `json:"task_id"`
}

// RestoredRecordLinks represents a custom type struct.
// URLs to pages related to the resource.
type RestoredRecordLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// RestoredRecordListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type RestoredRecordListEmbedded struct {
    // TODO: Add struct field description
    Items []*RestoredRecord `json:"items"`
}

// RestoredRecordListLinks represents a custom type struct.
// URLs to pages related to the resource.
type RestoredRecordListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// RetentionBackupSLAParam represents a custom type struct.
// The retention time for this SLA. For example, to retain the backup for 1 month, set `unit="months"` and `value=1`.
type RetentionBackupSLAParam struct {
    // The measurement unit of the SLA parameter.
    Unit  *string `json:"unit"`
    // The measurement value of the SLA parameter.
    Value *int64  `json:"value"`
}

// RoleForOrganizationalUnits represents a custom type struct.
// The organizational units assigned to the user, with the specified role.
type RoleForOrganizationalUnits struct {
    // The Clumio-assigned IDs of the organizational units assigned to the user.
    // Use the [GET /organizational-units](#operation/list-organizational-units) endpoint to fetch
    // valid values.
    OrganizationalUnitIds []*string `json:"organizational_unit_ids"`
    // The Clumio-assigned ID of the role assigned to the user.
    // Use the [GET /roles](#operation/list-roles) endpoint to fetch valid values.
    RoleId                *string   `json:"role_id"`
}

// RoleLinks represents a custom type struct.
// URLs to pages related to the resource.
type RoleLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// RoleListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type RoleListEmbedded struct {
    // RoleWithETag to support etag string to be calculated
    Items []*RoleWithETag `json:"items"`
}

// RoleListLinks represents a custom type struct.
// URLs to pages related to the resource.
type RoleListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// RoleModel represents a custom type struct.
// RoleModel denotes the Model for Role
type RoleModel struct {
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

// RoleWithETag represents a custom type struct.
// RoleWithETag to support etag string to be calculated
type RoleWithETag struct {
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

// Rule represents a custom type struct.
// A rule applies an action subject to a condition criteria.
type Rule struct {
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

// RuleAction represents a custom type struct.
// An action to be applied subject to the rule criteria.
type RuleAction struct {
    // Apply a policy to assets.
    AssignPolicy *AssignPolicyAction `json:"assign_policy"`
}

// RuleEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type RuleEmbedded struct {
    // Embeds the associated policy of a protected resource in the response if requested using the `embed` query parameter. Unprotected resources will not have an associated policy.
    ReadPolicyDefinition interface{} `json:"read-policy-definition"`
}

// RuleLinks represents a custom type struct.
// URLs to pages related to the resource.
type RuleLinks struct {
    // The HATEOAS link to this resource.
    Self                 *HateoasSelfLink                 `json:"_self"`
    // A resource-specific HATEOAS link.
    DeletePolicyRule     *HateoasLink                     `json:"delete-policy-rule"`
    // A HATEOAS link to the policy protecting this resource. Will be omitted for unprotected entities.
    ReadPolicyDefinition *ReadPolicyDefinitionHateoasLink `json:"read-policy-definition"`
    // A resource-specific HATEOAS link.
    UpdatePolicyRule     *HateoasLink                     `json:"update-policy-rule"`
}

// RuleListEmbedded represents a custom type struct.
// An array of embedded resources related to this resource.
type RuleListEmbedded struct {
    // A rule applies an action subject to a condition criteria.
    Items []*Rule `json:"items"`
}

// RuleListLinks represents a custom type struct.
// URLs to pages related to the resource.
type RuleListLinks struct {
    // The HATEOAS link to the first page of results.
    First            *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next             *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self             *HateoasSelfLink  `json:"_self"`
    // A resource-specific HATEOAS link.
    CreatePolicyRule *HateoasLink      `json:"create-policy-rule"`
}

// RulePriority represents a custom type struct.
// A priority relative to other rules.
type RulePriority struct {
    // The rule ID before which this rule should be inserted.
    BeforeRuleId *string `json:"before_rule_id"`
}

// RuleProvision represents a custom type struct.
// Specifies the role and the organizational units to be assigned to the user subject to the rule criteria.
type RuleProvision struct {
    // The Clumio-assigned IDs of the organizational units to be assigned to the user.
    // Use the [GET /organizational-units](#operation/list-organizational-units) endpoint to fetch valid values.
    OrganizationalUnitIds []*string `json:"organizational_unit_ids"`
    // The Clumio-assigned ID of the role to be assigned to the user.
    // Use the [GET /roles](#operation/list-roles) endpoint to fetch valid values.
    RoleId                *string   `json:"role_id"`
}

// S3AccessControlTranslation represents a custom type struct.
// A container for information about access control for replicas.
type S3AccessControlTranslation struct {
    // Specifies the replica ownership.
    Owner *string `json:"owner"`
}

// S3AssetInfo represents a custom type struct.
// S3AssetInfo
// The installed information for the S3 feature.
type S3AssetInfo struct {
    // The current version of the feature.
    InstalledTemplateVersion *string `json:"installed_template_version"`
}

// S3BucketSizeRes represents a custom type struct.
// The size breakdown in bytes with timestamps of a bucket per storage class.
type S3BucketSizeRes struct {
    // Size of Deep Archive Object Overhead in bytes.
    DeepArchiveObjectOverhead                   *int64  `json:"deep_archive_object_overhead"`
    // Timestamp when CloudWatch reported the Deep Archive Object Overhead size.
    DeepArchiveObjectOverheadRetrievedTime      *string `json:"deep_archive_object_overhead_retrieved_time"`
    // Size of Deep Archive S3 Object Overhead in bytes.
    DeepArchiveS3ObjectOverhead                 *int64  `json:"deep_archive_s3_object_overhead"`
    // Timestamp when CloudWatch reported the Deep Archive S3 Object Overhead size.
    DeepArchiveS3ObjectOverheadRetrievedTime    *string `json:"deep_archive_s3_object_overhead_retrieved_time"`
    // Size of Deep Archive Staging Storage objects in bytes.
    DeepArchiveStagingStorage                   *int64  `json:"deep_archive_staging_storage"`
    // Timestamp when CloudWatch reported the Deep Archive Staging Storage objects size.
    DeepArchiveStagingStorageRetrievedTime      *string `json:"deep_archive_staging_storage_retrieved_time"`
    // Size of Deep Archive Storage objects in bytes.
    DeepArchiveStorage                          *int64  `json:"deep_archive_storage"`
    // Timestamp when CloudWatch reported the Deep Archive Storage objects size.
    DeepArchiveStorageRetrievedTime             *string `json:"deep_archive_storage_retrieved_time"`
    // Size of Glacier Instant Retrieval Storage objects in bytes.
    GlacierInstantRetrievalStorage              *int64  `json:"glacier_instant_retrieval_storage"`
    // Timestamp when CloudWatch reported the Glacier Instant Retrieval Storage objects size.
    GlacierInstantRetrievalStorageRetrievedTime *string `json:"glacier_instant_retrieval_storage_retrieved_time"`
    // Size of Glacier Object Overhead in bytes.
    GlacierObjectOverhead                       *int64  `json:"glacier_object_overhead"`
    // Timestamp when CloudWatch reported the Glacier Object Overhead size.
    GlacierObjectOverheadRetrievedTime          *string `json:"glacier_object_overhead_retrieved_time"`
    // Size of Glacier S3 Object Overhead in bytes.
    GlacierS3ObjectOverhead                     *int64  `json:"glacier_s3_object_overhead"`
    // Timestamp when CloudWatch reported the Glacier S3 Object Overhead size.
    GlacierS3ObjectOverheadRetrievedTime        *string `json:"glacier_s3_object_overhead_retrieved_time"`
    // Size of Glacier Staging Storage objects in bytes.
    GlacierStagingStorage                       *int64  `json:"glacier_staging_storage"`
    // Timestamp when CloudWatch reported the Glacier Staging Storage objects size.
    GlacierStagingStorageRetrievedTime          *string `json:"glacier_staging_storage_retrieved_time"`
    // Size of Glacier Storage objects in bytes.
    GlacierStorage                              *int64  `json:"glacier_storage"`
    // Timestamp when CloudWatch reported the Glacier Storage objects size.
    GlacierStorageRetrievedTime                 *string `json:"glacier_storage_retrieved_time"`
    // Size of Intelligent-Tiering AA Storage objects in bytes.
    IntelligentTieringAaStorage                 *int64  `json:"intelligent_tiering_aa_storage"`
    // Timestamp when CloudWatch reported the Intelligent-Tiering AA Storage objects size.
    IntelligentTieringAaStorageRetrievedTime    *string `json:"intelligent_tiering_aa_storage_retrieved_time"`
    // Size of Intelligent-Tiering AIA Storage objects in bytes.
    IntelligentTieringAiaStorage                *int64  `json:"intelligent_tiering_aia_storage"`
    // Timestamp when CloudWatch reported the Intelligent-Tiering AIA Storage objects size.
    IntelligentTieringAiaStorageRetrievedTime   *string `json:"intelligent_tiering_aia_storage_retrieved_time"`
    // Size of Intelligent-Tiering DAA Storage objects in bytes.
    IntelligentTieringDaaStorage                *int64  `json:"intelligent_tiering_daa_storage"`
    // Timestamp when CloudWatch reported the Intelligent-Tiering DAA Storage objects size.
    IntelligentTieringDaaStorageRetrievedTime   *string `json:"intelligent_tiering_daa_storage_retrieved_time"`
    // Size of Intelligent-Tiering FA Storage objects in bytes.
    IntelligentTieringFaStorage                 *int64  `json:"intelligent_tiering_fa_storage"`
    // Timestamp when CloudWatch reported the Intelligent-Tiering FA Storage objects size.
    IntelligentTieringFaStorageRetrievedTime    *string `json:"intelligent_tiering_fa_storage_retrieved_time"`
    // Size of Intelligent-Tiering IA Storage objects in bytes.
    IntelligentTieringIaStorage                 *int64  `json:"intelligent_tiering_ia_storage"`
    // Timestamp when CloudWatch reported the Intelligent-Tiering IA Storage objects size.
    IntelligentTieringIaStorageRetrievedTime    *string `json:"intelligent_tiering_ia_storage_retrieved_time"`
    // Size of OneZone IA Overhead in bytes.
    OneZoneIaSizeOverhead                       *int64  `json:"one_zone_ia_size_overhead"`
    // Timestamp when CloudWatch reported the OneZone IA Overhead size.
    OneZoneIaSizeOverheadRetrievedTime          *string `json:"one_zone_ia_size_overhead_retrieved_time"`
    // Size of OneZone IA Storage objects in bytes.
    OneZoneIaStorage                            *int64  `json:"one_zone_ia_storage"`
    // Timestamp when CloudWatch reported the OneZone IA Storage objects size.
    OneZoneIaStorageRetrievedTime               *string `json:"one_zone_ia_storage_retrieved_time"`
    // Size of Reduced Redundancy Storage objects in bytes.
    ReducedRedundancyStorage                    *int64  `json:"reduced_redundancy_storage"`
    // Timestamp when CloudWatch reported the Reduced Redundancy Storage objects size.
    ReducedRedundancyStorageRetrievedTime       *string `json:"reduced_redundancy_storage_retrieved_time"`
    // Size of Standard IA Object Overhead in bytes.
    StandardIaObjectOverhead                    *int64  `json:"standard_ia_object_overhead"`
    // Timestamp when CloudWatch reported the Standard IA Object Overhead size.
    StandardIaObjectOverheadRetrievedTime       *string `json:"standard_ia_object_overhead_retrieved_time"`
    // Size of Standard IA Overhead in bytes.
    StandardIaSizeOverhead                      *int64  `json:"standard_ia_size_overhead"`
    // Timestamp when CloudWatch reported the Standard IA Overhead size.
    StandardIaSizeOverheadRetrievedTime         *string `json:"standard_ia_size_overhead_retrieved_time"`
    // Size of Standard IA Storage objects in bytes.
    StandardIaStorage                           *int64  `json:"standard_ia_storage"`
    // Timestamp when CloudWatch reported the Standard IA Storage objects size.
    StandardIaStorageRetrievedTime              *string `json:"standard_ia_storage_retrieved_time"`
    // Size of Standard Storage objects in bytes.
    StandardStorage                             *int64  `json:"standard_storage"`
    // Timestamp when CloudWatch reported the Standard Storage objects size.
    StandardStorageRetrievedTime                *string `json:"standard_storage_retrieved_time"`
}

// S3BucketsInventorySummaryBucketSizeBreakdown represents a custom type struct.
// The total size breakdown of S3 buckets in bytes per storage class. This parameter aggregates relevant fields from cloudwatch_metrics > size_bytes_per_storage_class
type S3BucketsInventorySummaryBucketSizeBreakdown struct {
    // Size of Glacier Deep Archive Storage in bytes.
    GlacierDeepArchiveStorageBytes       *int64 `json:"glacier_deep_archive_storage_bytes"`
    // Size of Glacier Flexible Retrieval Storage in bytes.
    GlacierFlexibleRetrievalStorageBytes *int64 `json:"glacier_flexible_retrieval_storage_bytes"`
    // Size of Glacier Instant Retrieval Storage in bytes.
    GlacierInstantRetrievalStorageBytes  *int64 `json:"glacier_instant_retrieval_storage_bytes"`
    // Size of Intelligent-Tiering Storage objects in bytes.
    IntelligentTieringStorageBytes       *int64 `json:"intelligent_tiering_storage_bytes"`
    // Size of OneZone-IA Storage in bytes.
    OneZoneIaStorageBytes                *int64 `json:"one_zone_ia_storage_bytes"`
    // Size of Reduced Redundancy Storage in bytes.
    ReducedRedundancyStorageBytes        *int64 `json:"reduced_redundancy_storage_bytes"`
    // Size of Standard-IA Storage in bytes.
    StandardIaStorageBytes               *int64 `json:"standard_ia_storage_bytes"`
    // Size of Standard Storage in bytes.
    StandardStorageBytes                 *int64 `json:"standard_storage_bytes"`
}

// S3CloudwatchMetrics represents a custom type struct.
// The Cloudwatch metrics of the bucket.
type S3CloudwatchMetrics struct {
    // The average size of object in bucket.
    AverageObjectSizeBytes     *float64         `json:"average_object_size_bytes"`
    // Timestamp when average size of the bucket is calculated.
    AverageObjectSizeBytesTime *string          `json:"average_object_size_bytes_time"`
    // Number of objects in bucket.
    ObjectCount                *int64           `json:"object_count"`
    // Timestamp when CloudWatch reported the bucket object count.
    ObjectCountRetrievedTime   *string          `json:"object_count_retrieved_time"`
    // Size of bucket in bytes.
    SizeBytes                  *int64           `json:"size_bytes"`
    // The size breakdown in bytes with timestamps of a bucket per storage class.
    SizeBytesPerStorageClass   *S3BucketSizeRes `json:"size_bytes_per_storage_class"`
    // Timestamp when CloudWatch reported the bucket size.
    SizeBytesRetrievedTime     *string          `json:"size_bytes_retrieved_time"`
}

// S3DeleteMarkerReplication represents a custom type struct.
// Specifies whether Amazon S3 replicates delete markers.
type S3DeleteMarkerReplication struct {
    // Indicates whether to replicate delete markers.
    Status *string `json:"status"`
}

// S3Destination represents a custom type struct.
// Specifies information about where to publish analysis or configuration results.
type S3Destination struct {
    // A container for information about access control for replicas.
    AccessControlTranslation *S3AccessControlTranslation `json:"access_control_translation"`
    // Destination bucket owner account ID.
    Account                  *string                     `json:"account"`
    // The Amazon Resource Name (ARN) of the bucket where
    // you want Amazon S3 to store the results.
    Bucket                   *string                     `json:"bucket"`
    // Specifies encryption-related information for an Amazon S3 bucket
    // that is a destination for replicated objects.
    EncryptionConfiguration  *S3EncryptionConfiguration  `json:"encryption_configuration"`
    // A container specifying replication metrics-related settings
    // enabling replication metrics and events.
    Metrics                  *S3Metrics                  `json:"metrics"`
    // A container specifying S3 Replication Time Control (S3 RTC)
    // related information.
    ReplicationTime          *S3ReplicationTime          `json:"replication_time"`
    // The storage class to use when replicating objects.
    StorageClass             *string                     `json:"storage_class"`
}

// S3EncryptionConfiguration represents a custom type struct.
// Specifies encryption-related information for an Amazon S3 bucket
// that is a destination for replicated objects.
type S3EncryptionConfiguration struct {
    // Specifies the ID (Key ARN or Alias ARN) of the customer managed
    // AWS KMS key stored in AWS Key Management Service (KMS) for the
    // destination bucket.
    ReplicaKmsKeyId *string `json:"replica_kms_key_id"`
}

// S3EncryptionOutput represents a custom type struct.
// The AWS encryption output of the bucket.
type S3EncryptionOutput struct {
    // Specifies the default server-side-encryption configuration.
    ServerSideEncryptionConfiguration *S3ServerSideEncryptionConfiguration `json:"server_side_encryption_configuration"`
}

// S3ExistingObjectReplication represents a custom type struct.
// Configuration to replicate existing source bucket objects.
type S3ExistingObjectReplication struct {
    // Specifies whether the existing object replication is enabled.
    Status *string `json:"status"`
}

// S3InstantAccessEndpoint represents a custom type struct.
// S3InstantAccessEndpoint
type S3InstantAccessEndpoint struct {
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

// S3InstantAccessEndpointEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type S3InstantAccessEndpointEmbedded struct {
    // Embeds the associated protection group S3 asset
    ReadProtectionGroupS3Asset interface{} `json:"read-protection-group-s3-asset"`
}

// S3InstantAccessEndpointLinks represents a custom type struct.
// URLs to pages related to the resource.
type S3InstantAccessEndpointLinks struct {
    // The HATEOAS link to this resource.
    Self                                                   *HateoasSelfLink `json:"_self"`
    // URL to the detailed information of the instant access endpoint
    ReadProtectionGroupInstantAccessEndpoint               interface{}      `json:"read-protection-group-instant-access-endpoint"`
    // URL to the role permissions of the instant access endpoint
    ReadProtectionGroupInstantAccessEndpointRolePermission interface{}      `json:"read-protection-group-instant-access-endpoint-role-permission"`
    // URL for getting URI of the instant access endpoint
    ReadProtectionGroupInstantAccessEndpointUri            interface{}      `json:"read-protection-group-instant-access-endpoint-uri"`
    // URL to the associated protection group S3 asset
    ReadProtectionGroupS3Asset                             interface{}      `json:"read-protection-group-s3-asset"`
}

// S3InstantAccessEndpointListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type S3InstantAccessEndpointListEmbedded struct {
    // S3InstantAccessEndpoint
    Items []*S3InstantAccessEndpoint `json:"items"`
}

// S3InstantAccessEndpointListLinks represents a custom type struct.
// URLs to pages related to the resource.
type S3InstantAccessEndpointListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// S3InstantAccessEndpointRole represents a custom type struct.
// IAM role which is allowed access to the OLAP endpoint.
type S3InstantAccessEndpointRole struct {
    // The alias of the IAM role given by the user in the UI.
    Alias                 *string `json:"alias"`
    // The ARN of the IAM role.
    Arn                   *string `json:"arn"`
    // The ID of the IAM role. Used as an identifier in the API URL.
    Id                    *string `json:"id"`
    // The time when the role was last modified, in RFC-3339 format.
    LastModifiedTimestamp *string `json:"last_modified_timestamp"`
}

// S3InstantAccessEndpointStat represents a custom type struct.
// S3InstantAccessEndpointStat
// Statistical metric related to the instant access endpoint.
// S3InstantAccessEndpointStat swagger: model S3InstantAccessEndpointStat
type S3InstantAccessEndpointStat struct {
    // The unit counts of the metric.
    Count *int64  `json:"count"`
    // The name of the metric.
    Name  *string `json:"name"`
    // Unit of the metric.
    Unit  *string `json:"unit"`
}

// S3InstantAccessSource represents a custom type struct.
// The parameters for creating a S3 instant access endpoint.
type S3InstantAccessSource struct {
    // The Clumio-assigned ID of the protection group S3 asset backup or protection group backup to
    // be restored. Use the endpoints
    // [GET /backups/protection-groups/s3-assets](#operation/list-backup-protection-group-s3-assets)
    // for protection group S3 asset backup, and
    // [GET /backups/protection-groups](#operation/list-backup-protection-groups)
    // for protection group backups to fetch valid values. 
    // Note that only one of either `backup_id` or `pitr` must be provided.
    BackupId                 *string                           `json:"backup_id"`
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

// S3InstantAccessSourcePitrOptions represents a custom type struct.
// The parameters to initiate a point-in-time creation of S3 instant access endpoint.<br>
// Note that only one of either `backup_id` or `pitr` must be provided.
type S3InstantAccessSourcePitrOptions struct {
    // The end time of the period in which the objects must be restored in RFC-3339 format.
    // Objects modified before this given time will be restored.
    // If `restore_end_timestamp` is provided without `restore_start_timestamp`, then it is the same
    // as a point-in-time restore.
    RestoreEndTimestamp   *string `json:"restore_end_timestamp"`
    // The start time of the period in which the objects must be restored in RFC-3339 format.
    // Objects modified since the given time will be restored.
    RestoreStartTimestamp *string `json:"restore_start_timestamp"`
}

// S3Metrics represents a custom type struct.
// A container specifying replication metrics-related settings
// enabling replication metrics and events.
type S3Metrics struct {
    // A container specifying the time value for S3 Replication Time
    // Control (S3 RTC) and replication metrics EventThreshold.
    EventThreshold *S3ReplicationTimeValue `json:"event_threshold"`
    // Specifies whether the replication metrics are enabled.
    Status         *string                 `json:"status"`
}

// S3ReplicaModifications represents a custom type struct.
// A filter that you can specify for selection for modifications on replicas.
type S3ReplicaModifications struct {
    // Specifies whether Amazon S3 replicates modifications on replicas.
    Status *string `json:"status"`
}

// S3ReplicationConfiguration represents a custom type struct.
// A container for replication rules with a maximum size
// of 2MB and a maximum of 1,000 rules.
type S3ReplicationConfiguration struct {
    // The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
    // role that Amazon S3 assumes when replicating objects.
    Role  *string              `json:"role"`
    // Specifies which Amazon S3 objects to replicate and where to store the replicas.
    Rules []*S3ReplicationRule `json:"rules"`
}

// S3ReplicationOutput represents a custom type struct.
// The AWS replication output of the bucket.
type S3ReplicationOutput struct {
    // A container for replication rules with a maximum size
    // of 2MB and a maximum of 1,000 rules.
    ReplicationConfiguration *S3ReplicationConfiguration `json:"replication_configuration"`
}

// S3ReplicationRule represents a custom type struct.
// Specifies which Amazon S3 objects to replicate and where to store the replicas.
type S3ReplicationRule struct {
    // Specifies whether Amazon S3 replicates delete markers.
    DeleteMarkerReplication   *S3DeleteMarkerReplication   `json:"delete_marker_replication"`
    // Specifies information about where to publish analysis or configuration results.
    Destination               *S3Destination               `json:"destination"`
    // Configuration to replicate existing source bucket objects.
    ExistingObjectReplication *S3ExistingObjectReplication `json:"existing_object_replication"`
    // A filter that identifies the subset of objects
    // to which the replication rule applies.
    Filter                    *S3ReplicationRuleFilter     `json:"filter"`
    // A unique identifier for the rule (max 255 characters).
    Id                        *string                      `json:"id"`
    // The priority indicates which rule has precedence whenever
    // two or more replication rules conflict.
    Priority                  *int64                       `json:"priority"`
    // A container that describes additional filters for identifying
    // the source objects that you want to replicate.
    SourceSelectionCriteria   *S3SourceSelectionCriteria   `json:"source_selection_criteria"`
    // Specifies whether the rule is enabled.
    Status                    *string                      `json:"status"`
}

// S3ReplicationRuleAndOperator represents a custom type struct.
// A container for specifying rule filters. The filters
// determine the subset of objects to which the rule applies.
type S3ReplicationRuleAndOperator struct {
    // An object key name prefix that identifies
    // the subset of objects to which the rule applies.
    Prefix *string  `json:"prefix"`
    // A container of a key value name pair.
    Tags   []*S3Tag `json:"tags"`
}

// S3ReplicationRuleFilter represents a custom type struct.
// A filter that identifies the subset of objects
// to which the replication rule applies.
type S3ReplicationRuleFilter struct {
    // A container for specifying rule filters. The filters
    // determine the subset of objects to which the rule applies.
    And    *S3ReplicationRuleAndOperator `json:"and"`
    // An object key name prefix that identifies the
    // subset of objects to which the rule applies.
    Prefix *string                       `json:"prefix"`
    // A container of a key value name pair.
    Tag    *S3Tag                        `json:"tag"`
}

// S3ReplicationTime represents a custom type struct.
// A container specifying S3 Replication Time Control (S3 RTC)
// related information.
type S3ReplicationTime struct {
    // Specifies whether the replication time is enabled.
    Status *string                 `json:"status"`
    // A container specifying the time value for S3 Replication Time
    // Control (S3 RTC) and replication metrics EventThreshold.
    Time   *S3ReplicationTimeValue `json:"time"`
}

// S3ReplicationTimeValue represents a custom type struct.
// A container specifying the time value for S3 Replication Time
// Control (S3 RTC) and replication metrics EventThreshold.
type S3ReplicationTimeValue struct {
    // Contains an integer specifying time in minutes.
    Minutes *int64 `json:"minutes"`
}

// S3ServerSideEncryptionByDefault represents a custom type struct.
// Describes the default server-side encryption to apply to new objects in the bucket.
type S3ServerSideEncryptionByDefault struct {
    // AWS Key Management Service (KMS) customer AWS KMS key ID to use for the default encryption.
    KmsMasterKeyId *string `json:"kms_master_key_id"`
    // Server-side encryption algorithm to use for the default encryption.
    SseAlgorithm   *string `json:"sse_algorithm"`
}

// S3ServerSideEncryptionConfiguration represents a custom type struct.
// Specifies the default server-side-encryption configuration.
type S3ServerSideEncryptionConfiguration struct {
    // Specifies the default server-side encryption configuration.
    Rules []*S3ServerSideEncryptionRule `json:"rules"`
}

// S3ServerSideEncryptionRule represents a custom type struct.
// Specifies the default server-side encryption configuration.
type S3ServerSideEncryptionRule struct {
    // Describes the default server-side encryption to apply to new objects in the bucket.
    ApplyServerSideEncryptionByDefault *S3ServerSideEncryptionByDefault `json:"apply_server_side_encryption_by_default"`
    // Specifies whether Amazon S3 should use an S3 Bucket Key with server-side
    // encryption using KMS (SSE-KMS) for new objects in the bucket.
    BucketKeyEnabled                   *bool                            `json:"bucket_key_enabled"`
}

// S3ServiceRoles represents a custom type struct
type S3ServiceRoles struct {
    // Role assumable by Event Bridge to send event notifications.
    ContinuousBackupsRoleArn *string `json:"continuous_backups_role_arn"`
}

// S3SourceSelectionCriteria represents a custom type struct.
// A container that describes additional filters for identifying
// the source objects that you want to replicate.
type S3SourceSelectionCriteria struct {
    // A filter that you can specify for selection for modifications on replicas.
    ReplicaModifications   *S3ReplicaModifications   `json:"replica_modifications"`
    // A container for filter information for the selection of
    // S3 objects encrypted with AWS KMS.
    SseKmsEncryptedObjects *S3SseKmsEncryptedObjects `json:"sse_kms_encrypted_objects"`
}

// S3SseKmsEncryptedObjects represents a custom type struct.
// A container for filter information for the selection of
// S3 objects encrypted with AWS KMS.
type S3SseKmsEncryptedObjects struct {
    // Specifies whether Amazon S3 replicates objects created with server-side
    // encryption using an AWS KMS key stored in AWS Key Management Service.
    Status *string `json:"status"`
}

// S3Tag represents a custom type struct.
// A container of a key value name pair.
type S3Tag struct {
    // Name of the object key.
    Key   *string `json:"key"`
    // Value of the tag.
    Value *string `json:"value"`
}

// S3TemplateInfo represents a custom type struct.
// The latest available information for the S3 feature.
type S3TemplateInfo struct {
    // The latest available feature version for the asset.
    AvailableTemplateVersion *string `json:"available_template_version"`
}

// S3VersioningOutput represents a custom type struct.
// The AWS versioning output of the bucket.
type S3VersioningOutput struct {
    // Specifies whether MFA delete is enabled in the bucket versioning configuration.
    MfaDelete *string `json:"mfa_delete"`
    // The versioning state of the bucket.
    Status    *string `json:"status"`
}

// SSESpecification represents a custom type struct.
// Represents the server-side encryption settings for a table.
type SSESpecification struct {
    // The server-side encryption KMS key type.
    // This field will only be populated for [GET /datasources/aws/dynamodb-tables/{table_id}](#operation/read-aws-dynamodb-table)
    // and [GET /backups/aws/dynamodb-tables/{backup_id}](#operation/read-backup-aws-dynamodb-table).
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), `kms_master_key_id` must be specified
    // in case of CUSTOMER_MANAGED. Possible values include: DEFAULT, AWS_MANAGED, CUSTOMER_MANAGED.
    KmsKeyType     *string `json:"kms_key_type"`
    // The AWS KMS customer master key (CMK) ARN that is used to encrypt the table.
    // If this field is `null`, server-side encryption is the default encryption (AWS owned CMK).
    // Otherwise, an AWS-managed or customer-managed CMK exists having these values.
    // For [POST /restores/aws/dynamodb](#operation/restore-aws-dynamodb-table), use key ID, Amazon Resource
    // Name (ARN), alias name or alias ARN to specify a key to be used for encrypting the restored table.
    // In case of default encryption (AWS owned CMK), specify this as `null`.
    KmsMasterKeyId *string `json:"kms_master_key_id"`
}

// ScheduleSetting represents a custom type struct.
// When the report will be generated and sent. If the schedule is not provided then a
// default value will be used.
type ScheduleSetting struct {
    // The day of the month when the report will be sent out. This is required for the 'monthly'
    // report frequency. It has to be >= 1 and <= 28, or '-1', which signifies end of month.
    // If the day_of_month is set to -1 then the report will be sent out at the end of every month.
    DayOfMonth *int64  `json:"day_of_month"`
    // Which day the report will be sent out. This is required for 'weekly' report frequency.
    DayOfWeek  *string `json:"day_of_week"`
    // The unit of frequency in which the report is generated.
    Frequency  *string `json:"frequency"`
    // When the report will be send out. This field should follow the format "HH:MM" based
    // on a 24-hour clock. Only values where HH ranges from 0 to 23 and MM ranges from 0 to
    // 59 are allowed.
    StartTime  *string `json:"start_time"`
    // The timezone for the report schedule. The timezone must be a valid location name from
    // the IANA Time Zone database. For instance, it can be "America/New_York", "US/Central",
    // "UTC", or similar. If empty, then the timezone is considered as UTC.
    Timezone   *string `json:"timezone"`
}

// ServiceInstanceProfiles represents a custom type struct
type ServiceInstanceProfiles struct {
    // TODO: Add struct field description
    Mssql *MssqlServiceInstanceProfiles `json:"mssql"`
}

// ServiceRoles represents a custom type struct
type ServiceRoles struct {
    // TODO: Add struct field description
    Mssql *MssqlServiceRoles `json:"mssql"`
    // TODO: Add struct field description
    S3    *S3ServiceRoles    `json:"s3"`
}

// SetAssignmentsResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type SetAssignmentsResponseLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// SetBucketPropertiesResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type SetBucketPropertiesResponseLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// ShareFileRestoreEmailLinks represents a custom type struct.
// URLs to pages related to the resource.
type ShareFileRestoreEmailLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// SingleErrorResponse represents a custom type struct
type SingleErrorResponse struct {
    // TODO: Add struct field description
    ErrorCode    *uint32 `json:"error_code"`
    // The reason for the error.
    ErrorMessage *string `json:"error_message"`
}

// SourceObjectFilters represents a custom type struct.
// Search for or restore only objects that pass the source object filter.
type SourceObjectFilters struct {
    // Filter for objects with this etag.
    Etag               *string   `json:"etag"`
    // If set to true, filter for latest versions only. Otherwise, all versions will
    // be returned.
    LatestVersionOnly  *bool     `json:"latest_version_only"`
    // Filter for objects with at most this size in bytes.
    MaxObjectSizeBytes *int64    `json:"max_object_size_bytes"`
    // Filter for objects with at least this size in bytes.
    MinObjectSizeBytes *int64    `json:"min_object_size_bytes"`
    // Filter for objects whose key contains this string.
    ObjectKeyContains  *string   `json:"object_key_contains"`
    // Filter for objects whose key matches this string.
    ObjectKeyMatches   *string   `json:"object_key_matches"`
    // Filter for objects that start with this key prefix.
    ObjectKeyPrefix    *string   `json:"object_key_prefix"`
    // Filter for objects that end with this key suffix.
    ObjectKeySuffix    *string   `json:"object_key_suffix"`
    // Storage class to include in the restore. If not specified, then all objects across all storage
    // classes will be restored. Valid values are: `S3 Standard`, `S3 Standard-IA`, `S3 Intelligent-Tiering`
    // and `S3 One Zone-IA`.
    StorageClasses     []*string `json:"storage_classes"`
    // Filter for objects with this version ID.
    VersionId          *string   `json:"version_id"`
}

// SourceObjectFiltersV2 represents a custom type struct.
// Search for or restore only objects that pass the source object filter.
type SourceObjectFiltersV2 struct {
    // Filter for objects with this etag.
    Etag                *string   `json:"etag"`
    // If set to true, filter for latest versions only. Otherwise, all versions will
    // be returned.
    IsLatestVersionOnly *bool     `json:"is_latest_version_only"`
    // Filter for objects with at most this size in bytes.
    MaxObjectSizeBytes  *int64    `json:"max_object_size_bytes"`
    // Filter for objects with at least this size in bytes.
    MinObjectSizeBytes  *int64    `json:"min_object_size_bytes"`
    // Filter for objects whose key contains this string.
    ObjectKeyContains   *string   `json:"object_key_contains"`
    // Filter for objects whose key matches this string.
    ObjectKeyMatches    *string   `json:"object_key_matches"`
    // Filter for objects that start with this key prefix.
    ObjectKeyPrefix     *string   `json:"object_key_prefix"`
    // Filter for objects that end with this key suffix.
    ObjectKeySuffix     *string   `json:"object_key_suffix"`
    // Storage class to include in the restore. If not specified, then all objects across all storage
    // classes will be restored. Valid values are: `S3 Standard`, `S3 Standard-IA`, `S3 Intelligent-Tiering`
    // and `S3 One Zone-IA`.
    StorageClasses      []*string `json:"storage_classes"`
    // Filter for objects with this version ID.
    VersionId           *string   `json:"version_id"`
}

// StreamSpecification represents a custom type struct.
// Represents the DynamoDB Streams configuration for a table in DynamoDB.
// and the data type (`S` for string, `N` for number, `B` for binary).
type StreamSpecification struct {
    // Indicates whether DynamoDB Streams is enabled (true) or disabled (false)
    // on the table.
    Enabled  *bool   `json:"enabled"`
    // When an item in the table is modified, ViewType determines what information
    // is written to the stream for this table. Valid values for ViewType
    // are:
    // 
    // KEYS_ONLY - Only the key attributes of the modified item are written
    // to the stream.
    // 
    // NEW_IMAGE - The entire item, as it appears after it was modified, is
    // written to the stream.
    // 
    // OLD_IMAGE - The entire item, as it appeared before it was modified,
    // is written to the stream.
    // 
    // NEW_AND_OLD_IMAGES - Both the new and the old item images of the item
    // are written to the stream.
    ViewType *string `json:"view_type"`
}

// Tag represents a custom type struct.
// The asset tags to be filtered. This is supported for AWS assets only.
type Tag struct {
    // The key of tag to filter.
    Key   *string `json:"key"`
    // The value of tag to filter.
    Value *string `json:"value"`
}

// TaskLinks represents a custom type struct.
// URLs to pages related to the resource.
type TaskLinks struct {
    // The HATEOAS link to this resource.
    Self                   *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    ReadOrganizationalUnit *HateoasLink     `json:"read-organizational-unit"`
    // A resource-specific HATEOAS link.
    UpdateTask             *HateoasLink     `json:"update-task"`
}

// TaskListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type TaskListEmbedded struct {
    // TODO: Add struct field description
    Items []*TaskWithETag `json:"items"`
}

// TaskListLinks represents a custom type struct.
// URLs to pages related to the resource.
type TaskListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// TaskParentEntity represents a custom type struct.
// The parent entity associated with the task.
type TaskParentEntity struct {
    // A system-generated ID assigned to this entity.
    Id         *string `json:"id"`
    // Type is mostly an asset type or the type of Entity. Some examples are
    // "restored_file", "aws_ebs_volume",  etc.
    ClumioType *string `json:"type"`
    // A system-generated value assigned to the entity. For example, if the primary entity type is "aws_ebs_volume", then the value is the name of the EBS.
    Value      *string `json:"value"`
}

// TaskPrimaryEntity represents a custom type struct.
// The primary entity associated with the task.
type TaskPrimaryEntity struct {
    // A system-generated ID assigned to this entity.
    Id         *string `json:"id"`
    // Type is mostly an asset type or the type of Entity. Some examples are
    // "restored_file", "aws_ebs_volume",  etc.
    ClumioType *string `json:"type"`
    // A system-generated value assigned to the entity. For example, if the primary entity type is "aws_ebs_volume", then the value is the name of the EBS.
    Value      *string `json:"value"`
}

// TaskWithETag represents a custom type struct
type TaskWithETag struct {
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

// TemplateConfigurationV2 represents a custom type struct.
// The configuration of the given template
type TemplateConfigurationV2 struct {
    // The AWS asset types supported with the available version of the template.
    AssetTypesEnabled        []*string                      `json:"asset_types_enabled"`
    // The latest available version for the template.
    AvailableTemplateVersion *string                        `json:"available_template_version"`
    // The latest available information for the DynamoDB feature.
    Dynamodb                 *DynamodbTemplateInfo          `json:"dynamodb"`
    // TODO: Add struct field description
    Ebs                      *EbsTemplateInfo               `json:"ebs"`
    // TODO: Add struct field description
    Ec2                      *Ec2TemplateInfo               `json:"ec2"`
    // The latest available information for the EC2 MSSQL feature.
    Ec2Mssql                 *EC2MSSQLTemplateInfo          `json:"ec2_mssql"`
    // IcebergOnGlueTemplateInfo is the latest available information for the IcebergOnGlue feature.
    IcebergOnGlue            *IcebergOnGlueTemplateInfo     `json:"iceberg_on_glue"`
    // IcebergOnS3TablesTemplateInfo is the latest available information for the IcebergOnS3Tables feature.
    IcebergOnS3Tables        *IcebergOnS3TablesTemplateInfo `json:"iceberg_on_s3_tables"`
    // TODO: Add struct field description
    Rds                      *RdsTemplateInfo               `json:"rds"`
    // The latest available information for the S3 feature.
    S3                       *S3TemplateInfo                `json:"s3"`
    // Configuration information about the Warm-Tier Protect feature of the template.
    WarmTierProtect          *WarmTierProtectTemplateInfo   `json:"warm_tier_protect"`
}

// TemplateLinks represents a custom type struct.
// URLs to pages related to the resource.
type TemplateLinks struct {
    // The HATEOAS link to this resource.
    Self *HateoasSelfLink `json:"_self"`
}

// TimeUnitParamAssetBackupMinRetentionDuration represents a custom type struct.
// The minimum required retention duration for a backup to be considered compliant.
type TimeUnitParamAssetBackupMinRetentionDuration struct {
    // Unit indicates the unit for time unit param.
    Unit  *string `json:"unit"`
    // Value indicates the value for time unit param.
    Value *int32  `json:"value"`
}

// TimeUnitParamLookbackPeriod represents a custom type struct.
// The duration prior to the compliance evaluation point to look back.
type TimeUnitParamLookbackPeriod struct {
    // Unit indicates the unit for time unit param.
    Unit  *string `json:"unit"`
    // Value indicates the value for time unit param.
    Value *int32  `json:"value"`
}

// TimeUnitParamPolicyMinRetentionDuration represents a custom type struct.
// The minimum retention duration for policy control.
type TimeUnitParamPolicyMinRetentionDuration struct {
    // Unit indicates the unit for time unit param.
    Unit  *string `json:"unit"`
    // Value indicates the value for time unit param.
    Value *int32  `json:"value"`
}

// TimeUnitParamRpoDuration represents a custom type struct.
// The minimum RPO duration for policy control.
type TimeUnitParamRpoDuration struct {
    // Unit indicates the unit for time unit param.
    Unit  *string `json:"unit"`
    // Value indicates the value for time unit param.
    Value *int32  `json:"value"`
}

// TimeUnitParamWindowSize represents a custom type struct.
// The size of each evaluation window within the look back period in which at least one compliant backup must exist.
type TimeUnitParamWindowSize struct {
    // Unit indicates the unit for time unit param.
    Unit  *string `json:"unit"`
    // Value indicates the value for time unit param.
    Value *int32  `json:"value"`
}

// UnprotectEntitiesHateoasLink represents a custom type struct.
// A HATEOAS link to unprotect the entities.
type UnprotectEntitiesHateoasLink struct {
    // The URI for the referenced operation.
    Href       *string `json:"href"`
    // Determines whether the "href" link is a URI template. If set to `true`, the "href" link is a URI template.
    Templated  *bool   `json:"templated"`
    // The HTTP method to be used with the "href" link for the referenced operation.
    ClumioType *string `json:"type"`
}

// UpdateEntities represents a custom type struct.
// Updates to the entities in the organizational unit.
// Adding or removing entities from the OU is an asynchronous operation.
// The response has a task ID which can be used to track the progress of the operation.
type UpdateEntities struct {
    // entityModel denotes the entityModel
    Add    []*EntityModel `json:"add"`
    // entityModel denotes the entityModel
    Remove []*EntityModel `json:"remove"`
}

// UpdatePolicyResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type UpdatePolicyResponseLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// UpdateProtectionGroupAssignments represents a custom type struct.
// UpdateProtectionGroupAssignments denotes the protection groups to be assigned or
// unassigned.
// Updates to the protection group assignments.
type UpdateProtectionGroupAssignments struct {
    // List of protection group IDs to assign to this organizational unit.
    Assign   []*string `json:"assign"`
    // List of protection group IDs to un-assign from this organizational unit.
    Unassign []*string `json:"unassign"`
}

// UpdateRuleResponseLinks represents a custom type struct.
// URLs to pages related to the resource.
type UpdateRuleResponseLinks struct {
    // The HATEOAS link to this resource.
    Self     *HateoasSelfLink     `json:"_self"`
    // A HATEOAS link to the task associated with this resource.
    ReadTask *ReadTaskHateoasLink `json:"read-task"`
}

// UpdateUserAssignments represents a custom type struct.
// Updates to the user assignments.
type UpdateUserAssignments struct {
    // List of user IDs to assign this organizational unit.
    Add    []*string `json:"add"`
    // List of user IDs to un-assign this organizational unit.
    Remove []*string `json:"remove"`
}

// UpdateUserAssignmentsWithRole represents a custom type struct.
// UpdateUserAssignmentsWithRole denotes the users to be added or removed along with the role.
type UpdateUserAssignmentsWithRole struct {
    // A user along with a role.
    Add    []*UserWithRole `json:"add"`
    // A user along with a role.
    Remove []*UserWithRole `json:"remove"`
}

// UserEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type UserEmbedded struct {
    // RoleModel denotes the Model for Role
    ReadRole []*RoleModel `json:"read-role"`
}

// UserEmbeddedV1 represents a custom type struct.
// Embedded responses related to the resource.
type UserEmbeddedV1 struct {
    // A description of the role.
    Description *string `json:"description"`
    // Unique name assigned to the role.
    Name        *string `json:"name"`
}

// UserHateoasV1 represents a custom type struct
type UserHateoasV1 struct {
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
    // The first and last name of the Clumio user. The name appears in the
    // User Management screen and is used to identify the user.
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
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and
    // the invitation is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                     *bool           `json:"is_enabled"`
    // The timestamp of when the user was last active in the Clumio system.
    // Represented in RFC-3339 format.
    LastActivityTimestamp         *string         `json:"last_activity_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount       *int64          `json:"organizational_unit_count"`
}

// UserLinks represents a custom type struct.
// URLs to pages related to the resource.
type UserLinks struct {
    // The HATEOAS link to this resource.
    Self       *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    DeleteUser *HateoasLink     `json:"delete-user"`
    // A resource-specific HATEOAS link.
    UpdateUser *HateoasLink     `json:"update-user"`
}

// UserListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type UserListEmbedded struct {
    // UserWithETag to support etag string to be calculated.
    Items []*UserWithETag `json:"items"`
}

// UserListEmbeddedV1 represents a custom type struct.
// Embedded responses related to the resource.
type UserListEmbeddedV1 struct {
    // TODO: Add struct field description
    Items []*UserHateoasV1 `json:"items"`
}

// UserListHateoasLinks represents a custom type struct.
// URLs to pages related to the resource.
type UserListHateoasLinks struct {
    // The HATEOAS link to the first page of results.
    First      *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last       *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next       *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev       *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self       *HateoasSelfLink  `json:"_self"`
    // A resource-specific HATEOAS link.
    CreateUser *HateoasLink      `json:"create-user"`
}

// UserWithETag represents a custom type struct.
// UserWithETag to support etag string to be calculated.
type UserWithETag struct {
    // Embedded responses related to the resource.
    Embedded                    *UserEmbedded                 `json:"_embedded"`
    // ETag value
    Etag                        *string                       `json:"_etag"`
    // URLs to pages related to the resource.
    Links                       *UserLinks                    `json:"_links"`
    // The organizational units assigned to the user, with the specified role.
    AccessControlConfiguration  []*RoleForOrganizationalUnits `json:"access_control_configuration"`
    // The email address of the Clumio user.
    Email                       *string                       `json:"email"`
    // The first and last name of the Clumio user. The name appears in the
    // User Management screen and is used to identify the user.
    FullName                    *string                       `json:"full_name"`
    // The Clumio-assigned ID of the Clumio user.
    Id                          *string                       `json:"id"`
    // The ID number of the user who sent the email invitation.
    Inviter                     *string                       `json:"inviter"`
    // Determines whether the user has activated their Clumio account.
    // If `true`, the user has activated the account.
    IsConfirmed                 *bool                         `json:"is_confirmed"`
    // Determines whether the user is enabled (in "Activated" or "Invited" status) in Clumio.
    // If `true`, the user is in "Activated" or "Invited" status in Clumio.
    // Users in "Activated" status can log in to Clumio.
    // Users in "Invited" status have been invited to log in to Clumio via an email invitation and
    // the invitation is pending acceptance from the user.
    // If `false`, the user has been manually suspended and cannot log in to Clumio
    // until another Clumio user reactivates the account.
    IsEnabled                   *bool                         `json:"is_enabled"`
    // The timestamp of when the user was last active in the Clumio system.
    // Represented in RFC-3339 format.
    LastActivityTimestamp       *string                       `json:"last_activity_timestamp"`
    // The last password change time of the Clumio user. Represented in RFC-3339 format.
    LastPasswordChangeTimestamp *string                       `json:"last_password_change_timestamp"`
    // The number of organizational units accessible to the user.
    OrganizationalUnitCount     *int64                        `json:"organizational_unit_count"`
}

// UserWithRole represents a custom type struct.
// A user along with a role.
type UserWithRole struct {
    // The Clumio-assigned ID of the role to be assigned to the user.
    AssignedRole *string `json:"assigned_role"`
    // The ID of the user.
    UserId       *string `json:"user_id"`
}

// Wallet represents a custom type struct
type Wallet struct {
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

// WalletLinks represents a custom type struct.
// URLs to pages related to the resource.
type WalletLinks struct {
    // The HATEOAS link to this resource.
    Self           *HateoasSelfLink `json:"_self"`
    // A resource-specific HATEOAS link.
    DeleteWallet   *HateoasLink     `json:"delete-wallet"`
    // A resource-specific HATEOAS link.
    ListWalletKeys *HateoasLink     `json:"list-wallet-keys"`
    // A resource-specific HATEOAS link.
    RefreshWallet  *HateoasLink     `json:"refresh-wallet"`
}

// WalletListEmbedded represents a custom type struct.
// Embedded responses related to the resource.
type WalletListEmbedded struct {
    // TODO: Add struct field description
    Items []*Wallet `json:"items"`
}

// WalletListLinks represents a custom type struct.
// URLs to pages related to the resource.
type WalletListLinks struct {
    // The HATEOAS link to the first page of results.
    First *HateoasFirstLink `json:"_first"`
    // The HATEOAS link to the last page of results.
    Last  *HateoasLastLink  `json:"_last"`
    // The HATEOAS link to the next page of results.
    Next  *HateoasNextLink  `json:"_next"`
    // The HATEOAS link to the previous page of results.
    Prev  *HateoasPrevLink  `json:"_prev"`
    // The HATEOAS link to this resource.
    Self  *HateoasSelfLink  `json:"_self"`
}

// WarmTierProtectConfig represents a custom type struct.
// The configuration of the Clumio Cloud Warm-Tier Protect product for this connection.
type WarmTierProtectConfig struct {
    // DynamodbAssetInfo
    // The installed information for the DynamoDB feature.
    Dynamodb                 *DynamodbAssetInfo `json:"dynamodb"`
    // The current version of the feature.
    InstalledTemplateVersion *string            `json:"installed_template_version"`
}

// WarmTierProtectTemplateInfo represents a custom type struct.
// Configuration information about the Warm-Tier Protect feature of the template.
type WarmTierProtectTemplateInfo struct {
    // The AWS asset types supported with the available version of the template.
    AssetTypesEnabled        []*string             `json:"asset_types_enabled"`
    // The latest available version for the template.
    AvailableTemplateVersion *string               `json:"available_template_version"`
    // The latest available information for the DynamoDB feature.
    Dynamodb                 *DynamodbTemplateInfo `json:"dynamodb"`
}
