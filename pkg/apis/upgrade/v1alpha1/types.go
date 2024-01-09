package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DeployService is the Schema for the deployservices API
type DeployService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeployServiceSpec   `json:"spec,omitempty"`
	Status DeployServiceStatus `json:"status,omitempty"`
}

// DeployServiceSpec is the spec for a DeployService resource
type DeployServiceSpec struct {
	Template []Config `json:"templates"`
}

type Config struct {
	Name        string             `json:"name"`
	DepTemp     DeploymentTemplate `json:"deploymentTemplate"`
	ServiceTemp ServiceTemplate    `json:"serviceTemplate"`
}

type DeploymentTemplate struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Replicas  int32  `json:"replicas"`
	ImageName string `json:"imageName"`
}

type ServiceTemplate struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	ServicePort int32  `json:"servicePort"`
}

// DeployServiceStatus is the status for a DeployService resource
type DeployServiceStatus struct {
	ConfigStates []ConfigState `json:"configStates"`
}

type ConfigState struct {
	Name       string   `json:"name"`
	DepCreated bool     `json:"DepCreated"`
	SvcCreated bool     `json:"SvcCreated"`
	EndPoints  []string `json:"endpoints"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DeployServiceList contains a list of DeployService
type DeployServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeployService `json:"items"`
}
