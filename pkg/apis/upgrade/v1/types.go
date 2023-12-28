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
	DepTemp DeploymentTemplate `json:"deploymentTemplate"`
	ServiceTemp ServiceTemplate `json:"serviceTemplate"`
}

type DeploymentTemplate struct{
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Replicas int32 `json:"replicas"`
	ImageName string `json:"imageName"`
}

type ServiceTemplate struct{
	Name string `json:"name"`
	Type string `json:"type"`
	ServicePort int32 `json:"servicePort"`
}

// UpgradeKubeStatus is the status for a UpgradeKube resource
type UpgradeKubeStatus struct {
	DepCreated bool `json:"DepCreated"`
	SvcCreated bool `json:"SvcCreated"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UpgradeKubeList is a list of UpgradeKube resources
type UpgradeKubeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []UpgradeKube `json:"items"`
}
