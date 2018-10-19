package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PodCheckpointerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []PodCheckpointer `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type PodCheckpointer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              PodCheckpointerSpec   `json:"spec"`
	Status            PodCheckpointerStatus `json:"status,omitempty"`
}

type PodCheckpointerSpec struct {
	// Fill me
}
type PodCheckpointerStatus struct {
	// Fill me
}
