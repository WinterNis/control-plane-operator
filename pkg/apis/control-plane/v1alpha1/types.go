package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ControlPlane `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ControlPlaneSpec   `json:"spec"`
	Status            ControlPlaneStatus `json:"status,omitempty"`
}

type ControlPlaneSpec struct {
	Size int32 `json:"size`
}
type ControlPlaneStatus struct {
	Nodes []string `json:"nodes"`
}
