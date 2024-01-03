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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	upgradev1 "github.com/arshukla98/sample-controller/generated/applyconfiguration/upgrade/v1"
	scheme "github.com/arshukla98/sample-controller/generated/clientset/versioned/scheme"
	v1 "github.com/arshukla98/sample-controller/pkg/apis/upgrade/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeployServicesGetter has a method to return a DeployServiceInterface.
// A group's client should implement this interface.
type DeployServicesGetter interface {
	DeployServices(namespace string) DeployServiceInterface
}

// DeployServiceInterface has methods to work with DeployService resources.
type DeployServiceInterface interface {
	Create(ctx context.Context, deployService *v1.DeployService, opts metav1.CreateOptions) (*v1.DeployService, error)
	Update(ctx context.Context, deployService *v1.DeployService, opts metav1.UpdateOptions) (*v1.DeployService, error)
	UpdateStatus(ctx context.Context, deployService *v1.DeployService, opts metav1.UpdateOptions) (*v1.DeployService, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DeployService, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DeployServiceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DeployService, err error)
	Apply(ctx context.Context, deployService *upgradev1.DeployServiceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DeployService, err error)
	ApplyStatus(ctx context.Context, deployService *upgradev1.DeployServiceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DeployService, err error)
	DeployServiceExpansion
}

// deployServices implements DeployServiceInterface
type deployServices struct {
	client rest.Interface
	ns     string
}

// newDeployServices returns a DeployServices
func newDeployServices(c *UpgradeV1Client, namespace string) *deployServices {
	return &deployServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deployService, and returns the corresponding deployService object, and an error if there is any.
func (c *deployServices) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DeployService, err error) {
	result = &v1.DeployService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeployServices that match those selectors.
func (c *deployServices) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DeployServiceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DeployServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deployServices.
func (c *deployServices) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deployservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deployService and creates it.  Returns the server's representation of the deployService, and an error, if there is any.
func (c *deployServices) Create(ctx context.Context, deployService *v1.DeployService, opts metav1.CreateOptions) (result *v1.DeployService, err error) {
	result = &v1.DeployService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deployservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deployService).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deployService and updates it. Returns the server's representation of the deployService, and an error, if there is any.
func (c *deployServices) Update(ctx context.Context, deployService *v1.DeployService, opts metav1.UpdateOptions) (result *v1.DeployService, err error) {
	result = &v1.DeployService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployservices").
		Name(deployService.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deployService).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *deployServices) UpdateStatus(ctx context.Context, deployService *v1.DeployService, opts metav1.UpdateOptions) (result *v1.DeployService, err error) {
	result = &v1.DeployService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployservices").
		Name(deployService.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deployService).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deployService and deletes it. Returns an error if one occurs.
func (c *deployServices) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployservices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deployServices) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployservices").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deployService.
func (c *deployServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DeployService, err error) {
	result = &v1.DeployService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deployservices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied deployService.
func (c *deployServices) Apply(ctx context.Context, deployService *upgradev1.DeployServiceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DeployService, err error) {
	if deployService == nil {
		return nil, fmt.Errorf("deployService provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(deployService)
	if err != nil {
		return nil, err
	}
	name := deployService.Name
	if name == nil {
		return nil, fmt.Errorf("deployService.Name must be provided to Apply")
	}
	result = &v1.DeployService{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("deployservices").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *deployServices) ApplyStatus(ctx context.Context, deployService *upgradev1.DeployServiceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.DeployService, err error) {
	if deployService == nil {
		return nil, fmt.Errorf("deployService provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(deployService)
	if err != nil {
		return nil, err
	}

	name := deployService.Name
	if name == nil {
		return nil, fmt.Errorf("deployService.Name must be provided to Apply")
	}

	result = &v1.DeployService{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("deployservices").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
