package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpgradeKube is a specification for a UpgradeKube resource
type UpgradeKube struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UpgradeKubeSpec   `json:"spec"`
	Status UpgradeKubeStatus `json:"status"`
}

// UpgradeKubeSpec is the spec for a UpgradeKube resource
type UpgradeKubeSpec struct {
	
}

// UpgradeKubeStatus is the status for a UpgradeKube resource
type UpgradeKubeStatus struct {

}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpgradeKubeList is a list of UpgradeKube resources
type UpgradeKubeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []UpgradeKube `json:"items"`
}
