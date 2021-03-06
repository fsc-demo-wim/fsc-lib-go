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
	"time"

	v1 "github.com/fsc-demo-wim/fsc-lib-go/pkg/apis/fsc.henderiw.be/v1"
	scheme "github.com/fsc-demo-wim/fsc-lib-go/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WorkLoadsGetter has a method to return a WorkLoadInterface.
// A group's client should implement this interface.
type WorkLoadsGetter interface {
	WorkLoads(namespace string) WorkLoadInterface
}

// WorkLoadInterface has methods to work with WorkLoad resources.
type WorkLoadInterface interface {
	Create(ctx context.Context, workLoad *v1.WorkLoad, opts metav1.CreateOptions) (*v1.WorkLoad, error)
	Update(ctx context.Context, workLoad *v1.WorkLoad, opts metav1.UpdateOptions) (*v1.WorkLoad, error)
	UpdateStatus(ctx context.Context, workLoad *v1.WorkLoad, opts metav1.UpdateOptions) (*v1.WorkLoad, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.WorkLoad, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.WorkLoadList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.WorkLoad, err error)
	WorkLoadExpansion
}

// workLoads implements WorkLoadInterface
type workLoads struct {
	client rest.Interface
	ns     string
}

// newWorkLoads returns a WorkLoads
func newWorkLoads(c *FscV1Client, namespace string) *workLoads {
	return &workLoads{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the workLoad, and returns the corresponding workLoad object, and an error if there is any.
func (c *workLoads) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.WorkLoad, err error) {
	result = &v1.WorkLoad{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workloads").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of WorkLoads that match those selectors.
func (c *workLoads) List(ctx context.Context, opts metav1.ListOptions) (result *v1.WorkLoadList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.WorkLoadList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("workloads").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested workLoads.
func (c *workLoads) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("workloads").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a workLoad and creates it.  Returns the server's representation of the workLoad, and an error, if there is any.
func (c *workLoads) Create(ctx context.Context, workLoad *v1.WorkLoad, opts metav1.CreateOptions) (result *v1.WorkLoad, err error) {
	result = &v1.WorkLoad{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("workloads").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workLoad).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a workLoad and updates it. Returns the server's representation of the workLoad, and an error, if there is any.
func (c *workLoads) Update(ctx context.Context, workLoad *v1.WorkLoad, opts metav1.UpdateOptions) (result *v1.WorkLoad, err error) {
	result = &v1.WorkLoad{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("workloads").
		Name(workLoad.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workLoad).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *workLoads) UpdateStatus(ctx context.Context, workLoad *v1.WorkLoad, opts metav1.UpdateOptions) (result *v1.WorkLoad, err error) {
	result = &v1.WorkLoad{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("workloads").
		Name(workLoad.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(workLoad).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the workLoad and deletes it. Returns an error if one occurs.
func (c *workLoads) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workloads").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *workLoads) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("workloads").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched workLoad.
func (c *workLoads) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.WorkLoad, err error) {
	result = &v1.WorkLoad{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("workloads").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
