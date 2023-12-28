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

package v1

// UpgradeKubeStatusApplyConfiguration represents an declarative configuration of the UpgradeKubeStatus type for use
// with apply.
type UpgradeKubeStatusApplyConfiguration struct {
	DepCreated *bool `json:"DepCreated,omitempty"`
	SvcCreated *bool `json:"SvcCreated,omitempty"`
}

// UpgradeKubeStatusApplyConfiguration constructs an declarative configuration of the UpgradeKubeStatus type for use with
// apply.
func UpgradeKubeStatus() *UpgradeKubeStatusApplyConfiguration {
	return &UpgradeKubeStatusApplyConfiguration{}
}

// WithDepCreated sets the DepCreated field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DepCreated field is set to the value of the last call.
func (b *UpgradeKubeStatusApplyConfiguration) WithDepCreated(value bool) *UpgradeKubeStatusApplyConfiguration {
	b.DepCreated = &value
	return b
}

// WithSvcCreated sets the SvcCreated field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SvcCreated field is set to the value of the last call.
func (b *UpgradeKubeStatusApplyConfiguration) WithSvcCreated(value bool) *UpgradeKubeStatusApplyConfiguration {
	b.SvcCreated = &value
	return b
}
