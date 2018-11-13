package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OrangesysSpec defines the desired state of Orangesys
type OrangesysSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	Label string `json:"label`
	Image string `json:"image"`
}

// OrangesysStatus defines the observed state of Orangesys
type OrangesysStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	Count int32 `json:"count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Orangesys is the Schema for the orangesys API
// +k8s:openapi-gen=true
type Orangesys struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrangesysSpec   `json:"spec,omitempty"`
	Status OrangesysStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OrangesysList contains a list of Orangesys
type OrangesysList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Orangesys `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Orangesys{}, &OrangesysList{})
}
