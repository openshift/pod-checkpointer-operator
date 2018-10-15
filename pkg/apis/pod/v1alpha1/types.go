package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PodCheckpointerOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []PodCheckpointerOperator `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PodCheckpointerOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              PodCheckpointerOperatorSpec   `json:"spec"`
	Status            PodCheckpointerOperatorStatus `json:"status,omitempty"`
}

type PodCheckpointerOperatorSpec struct {
	// Fill me
}
type PodCheckpointerOperatorStatus struct {
	// Fill me
}
