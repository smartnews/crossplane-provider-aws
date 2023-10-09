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
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// JobRunParameters defines the desired state of JobRun
type JobRunParameters struct {
	// Region is which region the JobRun will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`

	ConfigurationOverrides *string `json:"configurationOverrides,omitempty"`
	// The execution role ARN for the job run.
	ExecutionRoleARN *string `json:"executionRoleARN,omitempty"`
	// The job driver for the job run.
	JobDriver *JobDriver `json:"jobDriver,omitempty"`
	// The job template ID to be used to start the job run.
	JobTemplateID *string `json:"jobTemplateID,omitempty"`
	// The values of job template parameters to start a job run.
	JobTemplateParameters map[string]*string `json:"jobTemplateParameters,omitempty"`
	// The Amazon EMR release version to use for the job run.
	ReleaseLabel *string `json:"releaseLabel,omitempty"`
	// The retry policy configuration for the job run.
	RetryPolicyConfiguration *RetryPolicyConfiguration `json:"retryPolicyConfiguration,omitempty"`
	// The tags assigned to job runs.
	Tags                   map[string]*string `json:"tags,omitempty"`
	CustomJobRunParameters `json:",inline"`
}

// JobRunSpec defines the desired state of JobRun
type JobRunSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       JobRunParameters `json:"forProvider"`
}

// JobRunObservation defines the observed state of JobRun
type JobRunObservation struct {
	// This output lists the ARN of job run.
	ARN *string `json:"arn,omitempty"`
	// The reasons why the job run has failed.
	FailureReason *string `json:"failureReason,omitempty"`
	// This output displays the started job run ID.
	ID *string `json:"id,omitempty"`
	// This output displays the name of the started job run.
	Name *string `json:"name,omitempty"`
	// The state of the job run.
	State *string `json:"state,omitempty"`
	// Additional details of the job run state.
	StateDetails *string `json:"stateDetails,omitempty"`
	// This output displays the virtual cluster ID for which the job run was submitted.
	VirtualClusterID *string `json:"virtualClusterID,omitempty"`
}

// JobRunStatus defines the observed state of JobRun.
type JobRunStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          JobRunObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// JobRun is the Schema for the JobRuns API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type JobRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobRunSpec   `json:"spec"`
	Status            JobRunStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JobRunList contains a list of JobRuns
type JobRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JobRun `json:"items"`
}

// Repository type metadata.
var (
	JobRunKind             = "JobRun"
	JobRunGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: JobRunKind}.String()
	JobRunKindAPIVersion   = JobRunKind + "." + GroupVersion.String()
	JobRunGroupVersionKind = GroupVersion.WithKind(JobRunKind)
)

func init() {
	SchemeBuilder.Register(&JobRun{}, &JobRunList{})
}
