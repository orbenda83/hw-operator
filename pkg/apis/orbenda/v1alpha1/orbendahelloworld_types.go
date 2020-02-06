package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OrBendaHelloWorldSpec defines the desired state of OrBendaHelloWorld
type OrBendaHelloWorldSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Message  string `json:"message"`
	Replicas int32  `json:"replicas"`
}

// OrBendaHelloWorldStatus defines the observed state of OrBendaHelloWorld
type OrBendaHelloWorldStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OrBendaHelloWorld is the Schema for the orbendahelloworlds API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=orbendahelloworlds,scope=Namespaced
type OrBendaHelloWorld struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrBendaHelloWorldSpec   `json:"spec,omitempty"`
	Status OrBendaHelloWorldStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OrBendaHelloWorldList contains a list of OrBendaHelloWorld
type OrBendaHelloWorldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrBendaHelloWorld `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OrBendaHelloWorld{}, &OrBendaHelloWorldList{})
}
