/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SupersetDatasetSpec defines the desired state of a managed Superset dataset
// connection such as how to reach the dataset and how Superset should treat it.
type SupersetDatasetSpec struct {
	// SupersetRef identifies the Superset this dataset is imported into.
	SupersetRef *SupersetRef `json:"supersetRef"`

	// DatabaseRef identifies the SupersetDatabase resource for this dataset.
	DatabaseRef *SupersetDatabaseRef `json:"databaseRef"`

	// UUID is the Superset-side identity of this dataset. When set, the import
	// updates the existing Superset object with this UUID. When omitted, a new
	// Superset object is created instead.
	UUID       *string `json:"uuid,omitempty"`

	// TableName is the display name of this dataset connection in Superset.
	TableName  *string `json:"datasetName,omitempty"`

	// TableQuery is the SQL query for this dataset connection in Superset.
	TableQuery *string `json:"query,omitempty"`

	// DeletionPolicy controls whether the Superset-side dataset is retained or
	// removed when this CR is deleted. Defaults to Retain.
	// +kubebuilder:validation:Enum=0;1
	// +kubebuilder:validation:Default=0
	DeletionPolicy DeletionPolicyEnum `json:"deletion,omitempty"`
}

// SupersetDatasetStatus defines the observed state of a managed Superset dataset.
type SupersetDatasetStatus struct {
	Existence bool `json:"exists,omitempty"`
	// will be array in case sorting or search operations go along one or the other dimension
	Size       []int              `json:"size,omitempty"`
	Sorted     bool               `json:"sorted,omitempty"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SupersetDataset is a dataset managed in a Superset instance.
type SupersetDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of the dataset connection.
	Spec              SupersetDatasetSpec   `json:"spec,omitempty"`

	// Status is the observed state of the dataset connection.
	Status            SupersetDatasetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SupersetDatasetList contains a list of SupersetDataset resources.
type SupersetDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SupersetDataset `json:"items"`
}
