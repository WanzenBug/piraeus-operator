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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LinstorNodeSetSpec defines the desired state of LinstorNodeSet
type LinstorNodeSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// priorityClassName is the name of the PriorityClass for the node pods
	PriorityClassName PriorityClassName `json:"priorityClassName"`

	// StoragePools is a list of StoragePools for LinstorNodeSet to manage.
	// +optional
	// +nullable
	StoragePools *StoragePools `json:"storagePools"`

	// If set, the operator will automatically create storage pools of the specified type for all devices that can
	// be found. The name of the storage pools matches the device name. For example, all devices `/dev/sdc` will be
	// part of the `sdc` storage pool.
	// +optional
	// +kubebuilder:validation:Enum=None;LVM;LVMTHIN;ZFS
	AutomaticStorageType string `json:"automaticStorageType"`

	// drbdKernelModuleInjectionMode selects the source for the DRBD kernel module
	// +kubebuilder:validation:Enum=None;Compile;ShippedModules
	DRBDKernelModuleInjectionMode KernelModuleInjectionMode `json:"drbdKernelModuleInjectionMode"`

	// Name of k8s secret that holds the SSL key for a node (called `keystore.jks`) and
	// the trusted certificates (called `certificates.jks`)
	// +optional
	// +nullable
	SslConfig *LinstorSSLConfig `json:"sslSecret"`

	// drbdRepoCred is the name of the kubernetes secret that holds the credential for the DRBD repositories
	DrbdRepoCred string `json:"drbdRepoCred"`

	// satelliteImage is the image (location + tag) for the LINSTOR satellite container
	SatelliteImage string `json:"satelliteImage"`

	// kernelModImage is the image (location + tag) for the LINSTOR/DRBD kernel module injector container
	KernelModImage string `json:"kernelModImage"`

	// Cluster URL of the linstor controller.
	// If not set, will be determined from the current resource name.
	// +optional
	ControllerEndpoint string `json:"controllerEndpoint"`

	LinstorClientConfig `json:",inline"`
}

// KernelModuleInjectionMode describes the source for injecting a kernel module
type KernelModuleInjectionMode string

const (
	// ModuleInjectionNone means that no module will be injected
	ModuleInjectionNone = "None"
	// ModuleInjectionCompile means that the module will be compiled from sources available on the host
	ModuleInjectionCompile = "Compile"
	// ModuleInjectionShippedModules means that a module included in the injector image will be used
	ModuleInjectionShippedModules = "ShippedModules"
)

// LinstorNodeSetStatus defines the observed state of LinstorNodeSet
type LinstorNodeSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Keep the first letters in the json for Errors lowercase and all other
	// statuses uppercase. This causes `kubectl get TYPE NAME -oyaml` to sort
	// errors to below the other potentially very long Statuses which is the preferred UX.

	// Errors remaining that will trigger reconciliations.
	Errors []string `json:"errors"`
	// SatelliteStatuses by hostname.
	SatelliteStatuses []*SatelliteStatus `json:"SatelliteStatuses"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinstorNodeSet is the Schema for the linstornodesets API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=linstornodesets,scope=Namespaced
type LinstorNodeSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LinstorNodeSetSpec   `json:"spec,omitempty"`
	Status LinstorNodeSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinstorNodeSetList contains a list of LinstorNodeSet
type LinstorNodeSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinstorNodeSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LinstorNodeSet{}, &LinstorNodeSetList{})
}
