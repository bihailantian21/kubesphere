/*
Copyright 2019 The KubeSphere authors.

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

package v1alpha2

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha2 "kubesphere.io/kubesphere/pkg/apis/iam/v1alpha2"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// GlobalRoleBindingsGetter has a method to return a GlobalRoleBindingInterface.
// A group's client should implement this interface.
type GlobalRoleBindingsGetter interface {
	GlobalRoleBindings() GlobalRoleBindingInterface
}

// GlobalRoleBindingInterface has methods to work with GlobalRoleBinding resources.
type GlobalRoleBindingInterface interface {
	Create(*v1alpha2.GlobalRoleBinding) (*v1alpha2.GlobalRoleBinding, error)
	Update(*v1alpha2.GlobalRoleBinding) (*v1alpha2.GlobalRoleBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.GlobalRoleBinding, error)
	List(opts v1.ListOptions) (*v1alpha2.GlobalRoleBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.GlobalRoleBinding, err error)
	GlobalRoleBindingExpansion
}

// globalRoleBindings implements GlobalRoleBindingInterface
type globalRoleBindings struct {
	client rest.Interface
}

// newGlobalRoleBindings returns a GlobalRoleBindings
func newGlobalRoleBindings(c *IamV1alpha2Client) *globalRoleBindings {
	return &globalRoleBindings{
		client: c.RESTClient(),
	}
}

// Get takes name of the globalRoleBinding, and returns the corresponding globalRoleBinding object, and an error if there is any.
func (c *globalRoleBindings) Get(name string, options v1.GetOptions) (result *v1alpha2.GlobalRoleBinding, err error) {
	result = &v1alpha2.GlobalRoleBinding{}
	err = c.client.Get().
		Resource("globalrolebindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalRoleBindings that match those selectors.
func (c *globalRoleBindings) List(opts v1.ListOptions) (result *v1alpha2.GlobalRoleBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.GlobalRoleBindingList{}
	err = c.client.Get().
		Resource("globalrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalRoleBindings.
func (c *globalRoleBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("globalrolebindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a globalRoleBinding and creates it.  Returns the server's representation of the globalRoleBinding, and an error, if there is any.
func (c *globalRoleBindings) Create(globalRoleBinding *v1alpha2.GlobalRoleBinding) (result *v1alpha2.GlobalRoleBinding, err error) {
	result = &v1alpha2.GlobalRoleBinding{}
	err = c.client.Post().
		Resource("globalrolebindings").
		Body(globalRoleBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a globalRoleBinding and updates it. Returns the server's representation of the globalRoleBinding, and an error, if there is any.
func (c *globalRoleBindings) Update(globalRoleBinding *v1alpha2.GlobalRoleBinding) (result *v1alpha2.GlobalRoleBinding, err error) {
	result = &v1alpha2.GlobalRoleBinding{}
	err = c.client.Put().
		Resource("globalrolebindings").
		Name(globalRoleBinding.Name).
		Body(globalRoleBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the globalRoleBinding and deletes it. Returns an error if one occurs.
func (c *globalRoleBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("globalrolebindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalRoleBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("globalrolebindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched globalRoleBinding.
func (c *globalRoleBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.GlobalRoleBinding, err error) {
	result = &v1alpha2.GlobalRoleBinding{}
	err = c.client.Patch(pt).
		Resource("globalrolebindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
