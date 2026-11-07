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

type SupersetDatasetSpec struct {
	UUID       *string `json:"uuid,omitempty"`
	TableName  *string `json:"datasetName,omitempty"`
	TableQuery *string `json:"query,omitempty"`
}

type SupersetDatasetStatus struct {
	Existence bool `json:"exists,omitempty"`
	// will be array in case sorting or search operations go along one or the other dimension
	Size       []int              `json:"size,omitempty"`
	Sorted     bool               `json:"sorted,omitempty"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

type SupersetDataset struct {
	SupersetRef           *Superset             `json:"superset,omitempty"`
	DatabaseRef           *SupersetDatabase     `json:"database,omitempty"`
	SupersetDatasetSpec   SupersetDatasetSpec   `json:"spec,omitempty"`
	SupersetDatasetStatus SupersetDatasetStatus `json:"status,omitempty"`
	DeletionPolicy        *enum                 `json:"deletion,omitempty"`
}

type SupersetDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SupersetDataset `json:"items"`
}
