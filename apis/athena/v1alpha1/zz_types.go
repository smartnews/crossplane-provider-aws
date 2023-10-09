/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type ACLConfiguration struct {
	S3ACLOption *string `json:"s3ACLOption,omitempty"`
}

// +kubebuilder:skipversion
type ApplicationDPUSizes struct {
	ApplicationRuntimeID *string `json:"applicationRuntimeID,omitempty"`
}

// +kubebuilder:skipversion
type AthenaError struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kubebuilder:skipversion
type CalculationStatus struct {
	CompletionDateTime *metav1.Time `json:"completionDateTime,omitempty"`

	SubmissionDateTime *metav1.Time `json:"submissionDateTime,omitempty"`
}

// +kubebuilder:skipversion
type CapacityAllocation struct {
	StatusMessage *string `json:"statusMessage,omitempty"`
}

// +kubebuilder:skipversion
type Column struct {
	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type ColumnInfo struct {
	CatalogName *string `json:"catalogName,omitempty"`

	Label *string `json:"label,omitempty"`

	Name *string `json:"name,omitempty"`

	SchemaName *string `json:"schemaName,omitempty"`

	TableName *string `json:"tableName,omitempty"`

	Type *string `json:"type_,omitempty"`
}

// +kubebuilder:skipversion
type CustomerContentEncryptionConfiguration struct {
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kubebuilder:skipversion
type Database struct {
	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type EncryptionConfiguration struct {
	EncryptionOption *string `json:"encryptionOption,omitempty"`

	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kubebuilder:skipversion
type EngineVersion struct {
	EffectiveEngineVersion *string `json:"effectiveEngineVersion,omitempty"`

	SelectedEngineVersion *string `json:"selectedEngineVersion,omitempty"`
}

// +kubebuilder:skipversion
type NamedQuery struct {
	Name *string `json:"name,omitempty"`

	WorkGroup *string `json:"workGroup,omitempty"`
}

// +kubebuilder:skipversion
type NotebookMetadata struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`

	WorkGroup *string `json:"workGroup,omitempty"`
}

// +kubebuilder:skipversion
type NotebookSessionSummary struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`
}

// +kubebuilder:skipversion
type PreparedStatement struct {
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`

	WorkGroupName *string `json:"workGroupName,omitempty"`
}

// +kubebuilder:skipversion
type PreparedStatementSummary struct {
	LastModifiedTime *metav1.Time `json:"lastModifiedTime,omitempty"`
}

// +kubebuilder:skipversion
type QueryExecution struct {
	// The Athena engine version for running queries, or the PySpark engine version
	// for running sessions.
	EngineVersion *EngineVersion `json:"engineVersion,omitempty"`
	// The location in Amazon S3 where query and calculation results are stored
	// and the encryption option, if any, used for query and calculation results.
	// These are known as "client-side settings". If workgroup settings override
	// client-side settings, then the query uses the workgroup settings.
	ResultConfiguration *ResultConfiguration `json:"resultConfiguration,omitempty"`

	SubstatementType *string `json:"substatementType,omitempty"`

	WorkGroup *string `json:"workGroup,omitempty"`
}

// +kubebuilder:skipversion
type QueryExecutionStatistics struct {
	DataManifestLocation *string `json:"dataManifestLocation,omitempty"`
}

// +kubebuilder:skipversion
type QueryExecutionStatus struct {
	CompletionDateTime *metav1.Time `json:"completionDateTime,omitempty"`

	StateChangeReason *string `json:"stateChangeReason,omitempty"`

	SubmissionDateTime *metav1.Time `json:"submissionDateTime,omitempty"`
}

// +kubebuilder:skipversion
type QueryStage struct {
	State *string `json:"state,omitempty"`
}

// +kubebuilder:skipversion
type QueryStagePlanNode struct {
	Identifier *string `json:"identifier,omitempty"`

	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type ResultConfiguration struct {
	// Indicates that an Amazon S3 canned ACL should be set to control ownership
	// of stored query results. When Athena stores query results in Amazon S3, the
	// canned ACL is set with the x-amz-acl request header. For more information
	// about S3 Object Ownership, see Object Ownership settings (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html#object-ownership-overview)
	// in the Amazon S3 User Guide.
	ACLConfiguration *ACLConfiguration `json:"aclConfiguration,omitempty"`
	// If query and calculation results are encrypted in Amazon S3, indicates the
	// encryption option used (for example, SSE_KMS or CSE_KMS) and key information.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`

	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty"`

	OutputLocation *string `json:"outputLocation,omitempty"`
}

// +kubebuilder:skipversion
type ResultConfigurationUpdates struct {
	// Indicates that an Amazon S3 canned ACL should be set to control ownership
	// of stored query results. When Athena stores query results in Amazon S3, the
	// canned ACL is set with the x-amz-acl request header. For more information
	// about S3 Object Ownership, see Object Ownership settings (https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html#object-ownership-overview)
	// in the Amazon S3 User Guide.
	ACLConfiguration *ACLConfiguration `json:"aclConfiguration,omitempty"`
	// If query and calculation results are encrypted in Amazon S3, indicates the
	// encryption option used (for example, SSE_KMS or CSE_KMS) and key information.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`

	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty"`

	OutputLocation *string `json:"outputLocation,omitempty"`

	RemoveACLConfiguration *bool `json:"removeACLConfiguration,omitempty"`

	RemoveEncryptionConfiguration *bool `json:"removeEncryptionConfiguration,omitempty"`

	RemoveExpectedBucketOwner *bool `json:"removeExpectedBucketOwner,omitempty"`

	RemoveOutputLocation *bool `json:"removeOutputLocation,omitempty"`
}

// +kubebuilder:skipversion
type SessionConfiguration struct {
	// If query and calculation results are encrypted in Amazon S3, indicates the
	// encryption option used (for example, SSE_KMS or CSE_KMS) and key information.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`

	ExecutionRole *string `json:"executionRole,omitempty"`

	WorkingDirectory *string `json:"workingDirectory,omitempty"`
}

// +kubebuilder:skipversion
type SessionStatus struct {
	EndDateTime *metav1.Time `json:"endDateTime,omitempty"`

	IdleSinceDateTime *metav1.Time `json:"idleSinceDateTime,omitempty"`

	LastModifiedDateTime *metav1.Time `json:"lastModifiedDateTime,omitempty"`

	StartDateTime *metav1.Time `json:"startDateTime,omitempty"`
}

// +kubebuilder:skipversion
type SessionSummary struct {
	// The Athena engine version for running queries, or the PySpark engine version
	// for running sessions.
	EngineVersion *EngineVersion `json:"engineVersion,omitempty"`

	NotebookVersion *string `json:"notebookVersion,omitempty"`
}

// +kubebuilder:skipversion
type TableMetadata struct {
	Name *string `json:"name,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kubebuilder:skipversion
type WorkGroupConfiguration struct {
	AdditionalConfiguration *string `json:"additionalConfiguration,omitempty"`

	BytesScannedCutoffPerQuery *int64 `json:"bytesScannedCutoffPerQuery,omitempty"`
	// Specifies the KMS key that is used to encrypt the user's data stores in Athena.
	// This setting does not apply to Athena SQL workgroups.
	CustomerContentEncryptionConfiguration *CustomerContentEncryptionConfiguration `json:"customerContentEncryptionConfiguration,omitempty"`

	EnableMinimumEncryptionConfiguration *bool `json:"enableMinimumEncryptionConfiguration,omitempty"`

	EnforceWorkGroupConfiguration *bool `json:"enforceWorkGroupConfiguration,omitempty"`
	// The Athena engine version for running queries, or the PySpark engine version
	// for running sessions.
	EngineVersion *EngineVersion `json:"engineVersion,omitempty"`

	ExecutionRole *string `json:"executionRole,omitempty"`

	PublishCloudWatchMetricsEnabled *bool `json:"publishCloudWatchMetricsEnabled,omitempty"`

	RequesterPaysEnabled *bool `json:"requesterPaysEnabled,omitempty"`
	// The location in Amazon S3 where query and calculation results are stored
	// and the encryption option, if any, used for query and calculation results.
	// These are known as "client-side settings". If workgroup settings override
	// client-side settings, then the query uses the workgroup settings.
	ResultConfiguration *ResultConfiguration `json:"resultConfiguration,omitempty"`
}

// +kubebuilder:skipversion
type WorkGroupConfigurationUpdates struct {
	AdditionalConfiguration *string `json:"additionalConfiguration,omitempty"`

	BytesScannedCutoffPerQuery *int64 `json:"bytesScannedCutoffPerQuery,omitempty"`
	// Specifies the KMS key that is used to encrypt the user's data stores in Athena.
	// This setting does not apply to Athena SQL workgroups.
	CustomerContentEncryptionConfiguration *CustomerContentEncryptionConfiguration `json:"customerContentEncryptionConfiguration,omitempty"`

	EnableMinimumEncryptionConfiguration *bool `json:"enableMinimumEncryptionConfiguration,omitempty"`

	EnforceWorkGroupConfiguration *bool `json:"enforceWorkGroupConfiguration,omitempty"`
	// The Athena engine version for running queries, or the PySpark engine version
	// for running sessions.
	EngineVersion *EngineVersion `json:"engineVersion,omitempty"`

	ExecutionRole *string `json:"executionRole,omitempty"`

	PublishCloudWatchMetricsEnabled *bool `json:"publishCloudWatchMetricsEnabled,omitempty"`

	RemoveBytesScannedCutoffPerQuery *bool `json:"removeBytesScannedCutoffPerQuery,omitempty"`

	RemoveCustomerContentEncryptionConfiguration *bool `json:"removeCustomerContentEncryptionConfiguration,omitempty"`

	RequesterPaysEnabled *bool `json:"requesterPaysEnabled,omitempty"`
	// The information about the updates in the query results, such as output location
	// and encryption configuration for the query results.
	ResultConfigurationUpdates *ResultConfigurationUpdates `json:"resultConfigurationUpdates,omitempty"`
}

// +kubebuilder:skipversion
type WorkGroupSummary struct {
	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	Description *string `json:"description,omitempty"`
	// The Athena engine version for running queries, or the PySpark engine version
	// for running sessions.
	EngineVersion *EngineVersion `json:"engineVersion,omitempty"`

	Name *string `json:"name,omitempty"`

	State *string `json:"state,omitempty"`
}

// +kubebuilder:skipversion
type WorkGroup_SDK struct {
	// The configuration of the workgroup, which includes the location in Amazon
	// S3 where query and calculation results are stored, the encryption option,
	// if any, used for query and calculation results, whether the Amazon CloudWatch
	// Metrics are enabled for the workgroup and whether workgroup settings override
	// query settings, and the data usage limits for the amount of data scanned
	// per query or per workgroup. The workgroup settings override is specified
	// in EnforceWorkGroupConfiguration (true/false) in the WorkGroupConfiguration.
	// See WorkGroupConfiguration$EnforceWorkGroupConfiguration.
	Configuration *WorkGroupConfiguration `json:"configuration,omitempty"`

	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	State *string `json:"state,omitempty"`
}
