/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SupersetDatabaseSpec defines the desired state of a managed Superset database
// connection such as how to reach the database and how Superset should treat it.
type SupersetDatabaseSpec struct {
	// HostPortUserPassword holds the structured connection details (host, port,
	// database, username, password) and their Secret references.
	HostPortUserPassword *ConnectionSpec `json:"hpup,omitempty"`

	// UUID is the Superset-side identity of this database. When set, the import
	// updates the existing Superset object with this UUID. When omitted, a new
	// Superset object is created instead.
	UUID *string `json:"uuid,omitempty"`

	// EncryptedExtra is Superset's inline encrypted_extra payload for this database
	// (engine-specific credentials such as BigQuery/Snowflake key-pair auth).
	EncryptedExtra *string `json:"encryptedExtra,omitempty"`

	// EncryptedExtraFrom references a Secret key holding the encrypted_extra payload.
	// Mutually exclusive with EncryptedExtra.
	EncryptedExtraFrom *corev1.SecretKeySelector `json:"encryptedExtraFrom,omitempty"`

	// SshTunnel is inline SSH tunnel configuration for reaching the database.
	SshTunnel *string `json:"sshTunnel,omitempty"`

	// SshTunnelFrom references a Secret key holding the SSH tunnel configuration.
	// Mutually exclusive with SshTunnel.
	// +optional
	SshTunnelFrom *corev1.SecretKeySelector `json:"sshTunnelFrom,omitempty"`

	// Below are various security flags (Each flag and its *From selector are mutually exclusive).

	// AllowDml permits write statements (INSERT/UPDATE/DELETE) from SQL Lab on this
	// database. When unset, Superset's default behavior applies.
	// +optional
	AllowDml *bool `json:"dml,omitempty"`

	// AllowDmlFrom references a Secret key containing the AllowDml flag.
	// Mutually exclusive with AllowDml.
	// +optional
	AllowDmlFrom *corev1.SecretKeySelector `json:"dmlFrom,omitempty"`

	// FileUpload permits uploading data files (CSV/Excel) to this database.
	// +optional
	FileUpload *bool `json:"upload,omitempty"`

	// FileUploadFrom references a Secret key containing the FileUpload flag.
	// Mutually exclusive with FileUpload.
	// +optional
	FileUploadFrom *corev1.SecretKeySelector `json:"uploadFrom,omitempty"`

	// ExposeInSqllab makes this database selectable in SQL Lab.
	// +optional
	ExposeInSqllab *bool `json:"exposeInSqlLab,omitempty"`

	// ExposeInSqllabFrom references a Secret key containing the ExposeInSqllab flag.
	// Mutually exclusive with ExposeInSqllab.
	// +optional
	ExposeInSqllabFrom *corev1.SecretKeySelector `json:"exposeInSqlLabFrom,omitempty"`

	// ImpersonateUser runs queries as the logged-in Superset user rather than the
	// connection user. Only applicable where the database supports impersonation.
	// +optional
	ImpersonateUser *bool `json:"impersonate,omitempty"`

	// ImpersonateUserFrom references a Secret key containing the ImpersonateUser flag.
	// Mutually exclusive with ImpersonateUser.
	// +optional
	ImpersonateUserFrom *corev1.SecretKeySelector `json:"impersonateFrom,omitempty"`
}

// SupersetDatabaseStatus defines the observed state of a managed Superset database.
type SupersetDatabaseStatus struct {
	// ExternalManagement reports whether the object is marked externally managed in
	// Superset, which hides edit controls in the UI.
	ExternalManagement *bool `json:"externalManagement,omitempty"`

	// ExternalUrl is the external_url recorded on the Superset object.
	ExternalUrl *string `json:"externalUrl,omitempty"`

	// Conditions give a detailed breakdown of resource state via standard fields
	// (Type, Status, Reason, etc.). Reveals object status and potential issues.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SupersetDatabase is a database connection managed in a Superset instance.
type SupersetDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// SupersetRef identifies the Superset this database is imported into.
	SupersetRef *SupersetRef `json:"ref,omitempty"`

	// Spec defines the desired state of the database connection.
	Spec SupersetDatabaseSpec `json:"spec,omitempty"`

	// Status is the observed state of the database connection.
	Status SupersetDatabaseStatus `json:"status,omitempty"`

	// DeletionPolicy controls whether the Superset-side database is retained or
	// removed when this CR is deleted.
	DeletionPolicy DeletionPolicyEnum `json:"deletion,omitempty"`
}

// +kubebuilder:object:root=true

// SupersetDatabaseList contains a list of SupersetDatabase.
type SupersetDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SupersetDatabase `json:"items"`
}