/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	upgradev1 "github.com/arshukla98/sample-controller/generated/applyconfiguration/upgrade/v1"
	v1 "github.com/arshukla98/sample-controller/pkg/apis/upgrade/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=upgrade.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("DeploymentTemplate"):
		return &upgradev1.DeploymentTemplateApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DeployService"):
		return &upgradev1.DeployServiceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DeployServiceSpec"):
		return &upgradev1.DeployServiceSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DeployServiceStatus"):
		return &upgradev1.DeployServiceStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServiceTemplate"):
		return &upgradev1.ServiceTemplateApplyConfiguration{}

	}
	return nil
}
