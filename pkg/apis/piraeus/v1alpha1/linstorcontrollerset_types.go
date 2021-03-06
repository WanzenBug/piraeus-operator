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

// LinstorControllerSetSpec defines the desired state of LinstorControllerSet
type LinstorControllerSetSpec struct {

	// priorityClassName is the name of the PriorityClass for the controller pods
	PriorityClassName PriorityClassName `json:"priorityClassName"`

	// DBConnectionURL is the URL of the ETCD endpoint for LINSTOR Controller
	DBConnectionURL string `json:"dbConnectionURL"`

	// DBCertSecret is the name of the kubernetes secret that holds the CA certificate used to verify
	// the datatbase connection. The secret must contain a key "ca.pem" which holds the certificate in
	// PEM format
	// +nullable
	// +optional
	DBCertSecret string `json:"dbCertSecret"`

	// Use a TLS client certificate for authentication with the database (etcd). If set to true,
	// `dbCertSecret` must be set and contain two additional entries "client.cert" (PEM encoded)
	// and "client.key" (PKCS8 encoded, without passphrase).
	// +optional
	DBUseClientCert bool `json:"dbUseClientCert"`

	// Name of the secret containing the master passphrase for LUKS devices as `MASTER_PASSPHRASE`
	// +nullable
	// +optional
	LuksSecret string `json:"luksSecret"`

	// Name of k8s secret that holds the SSL key for a node (called `keystore.jks`) and the
	// trusted certificates (called `certificates.jks`)
	// +nullable
	// +optional
	SslConfig *LinstorSSLConfig `json:"sslSecret"`

	// DrbdRepoCred is the name of the kubernetes secret that holds the credential for the
	// DRBD repositories
	DrbdRepoCred string `json:"drbdRepoCred"`

	// controllerImage is the image (location + tag) for the LINSTOR controller/server container
	ControllerImage string `json:"controllerImage"`

	// Name of the secret containing the java keystore (`keystore.jks`) used to enable HTTPS on the
	// controller. The controller will create a secured https endpoint on port 3371 with the key
	// stored in `keystore.jks`. The keystore must be secured using the passphrase "linstor". Also
	// needs to contain a truststore `truststore.jks`, which will be used to authenticate clients.
	// +optional
	LinstorHttpsControllerSecret string `json:"linstorHttpsControllerSecret"`

	LinstorClientConfig `json:",inline"`
}

// LinstorControllerSetStatus defines the observed state of LinstorControllerSet
type LinstorControllerSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Keep the first letters in the json for Errors lowercase and all other
	// statuses uppercase. This causes `kubectl get TYPE NAME -oyaml` to sort
	// errors to below the other potentially very long Statuses which is the preferred UX.

	// Errors remaining that will trigger reconciliations.
	Errors []string `json:"errors"`
	// ControllerStatus information.
	ControllerStatus *NodeStatus `json:"ControllerStatus"`
	// SatelliteStatuses by hostname.
	SatelliteStatuses []*SatelliteStatus `json:"SatelliteStatuses"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinstorControllerSet is the Schema for the linstorcontrollersets API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=linstorcontrollersets,scope=Namespaced
type LinstorControllerSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LinstorControllerSetSpec   `json:"spec,omitempty"`
	Status LinstorControllerSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinstorControllerSetList contains a list of LinstorControllerSet
type LinstorControllerSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinstorControllerSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LinstorControllerSet{}, &LinstorControllerSetList{})
}
