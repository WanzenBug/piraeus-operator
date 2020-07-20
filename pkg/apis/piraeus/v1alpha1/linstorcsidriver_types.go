/*
Piraeus Operator
Copyright 2019 LINBIT USA, LLC.

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

package v1alpha1

import (
	v1 "github.com/piraeusdatastore/piraeus-operator/pkg/apis/piraeus/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LinstorCSIDriverSpec defines the desired state of LinstorCSIDriver
type LinstorCSIDriverSpec struct {
	v1.LinstorCSIDriverSpec `json:",inline"`
}

// LinstorCSIDriverStatus defines the observed state of LinstorCSIDriver
type LinstorCSIDriverStatus struct {
	v1.LinstorCSIDriverStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinstorCSIDriver is the Schema for the linstorcsidrivers API.
// DEPRECATED: use the piraeus.linstor.com/v1 version instead.
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=linstorcsidrivers,scope=Namespaced
// +kubebuilder:printcolumn:name="NodeReady",type="boolean",JSONPath=".status.NodeReady"
// +kubebuilder:printcolumn:name="ControllerReady",type="boolean",JSONPath=".status.ControllerReady"
type LinstorCSIDriver struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LinstorCSIDriverSpec   `json:"spec,omitempty"`
	Status LinstorCSIDriverStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinstorCSIDriverList contains a list of LinstorCSIDriver
// DEPRECATED: use the piraeus.linstor.com/v1 version instead.
type LinstorCSIDriverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinstorCSIDriver `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LinstorCSIDriver{}, &LinstorCSIDriverList{})
}
