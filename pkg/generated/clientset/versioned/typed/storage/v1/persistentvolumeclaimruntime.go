/*
 * Tencent is pleased to support the open source community by making TKEStack available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "tkestack.io/volume-decorator/pkg/apis/storage/v1"
	scheme "tkestack.io/volume-decorator/pkg/generated/clientset/versioned/scheme"
)

// PersistentVolumeClaimRuntimesGetter has a method to return a PersistentVolumeClaimRuntimeInterface.
// A group's client should implement this interface.
type PersistentVolumeClaimRuntimesGetter interface {
	PersistentVolumeClaimRuntimes(namespace string) PersistentVolumeClaimRuntimeInterface
}

// PersistentVolumeClaimRuntimeInterface has methods to work with PersistentVolumeClaimRuntime resources.
type PersistentVolumeClaimRuntimeInterface interface {
	Create(*v1.PersistentVolumeClaimRuntime) (*v1.PersistentVolumeClaimRuntime, error)
	Update(*v1.PersistentVolumeClaimRuntime) (*v1.PersistentVolumeClaimRuntime, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.PersistentVolumeClaimRuntime, error)
	List(opts metav1.ListOptions) (*v1.PersistentVolumeClaimRuntimeList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PersistentVolumeClaimRuntime, err error)
	PersistentVolumeClaimRuntimeExpansion
}

// persistentVolumeClaimRuntimes implements PersistentVolumeClaimRuntimeInterface
type persistentVolumeClaimRuntimes struct {
	client rest.Interface
	ns     string
}

// newPersistentVolumeClaimRuntimes returns a PersistentVolumeClaimRuntimes
func newPersistentVolumeClaimRuntimes(c *StorageV1Client, namespace string) *persistentVolumeClaimRuntimes {
	return &persistentVolumeClaimRuntimes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the persistentVolumeClaimRuntime, and returns the corresponding persistentVolumeClaimRuntime object, and an error if there is any.
func (c *persistentVolumeClaimRuntimes) Get(name string, options metav1.GetOptions) (result *v1.PersistentVolumeClaimRuntime, err error) {
	result = &v1.PersistentVolumeClaimRuntime{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("persistentvolumeclaimruntimes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PersistentVolumeClaimRuntimes that match those selectors.
func (c *persistentVolumeClaimRuntimes) List(opts metav1.ListOptions) (result *v1.PersistentVolumeClaimRuntimeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PersistentVolumeClaimRuntimeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("persistentvolumeclaimruntimes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested persistentVolumeClaimRuntimes.
func (c *persistentVolumeClaimRuntimes) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("persistentvolumeclaimruntimes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a persistentVolumeClaimRuntime and creates it.  Returns the server's representation of the persistentVolumeClaimRuntime, and an error, if there is any.
func (c *persistentVolumeClaimRuntimes) Create(persistentVolumeClaimRuntime *v1.PersistentVolumeClaimRuntime) (result *v1.PersistentVolumeClaimRuntime, err error) {
	result = &v1.PersistentVolumeClaimRuntime{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("persistentvolumeclaimruntimes").
		Body(persistentVolumeClaimRuntime).
		Do().
		Into(result)
	return
}

// Update takes the representation of a persistentVolumeClaimRuntime and updates it. Returns the server's representation of the persistentVolumeClaimRuntime, and an error, if there is any.
func (c *persistentVolumeClaimRuntimes) Update(persistentVolumeClaimRuntime *v1.PersistentVolumeClaimRuntime) (result *v1.PersistentVolumeClaimRuntime, err error) {
	result = &v1.PersistentVolumeClaimRuntime{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("persistentvolumeclaimruntimes").
		Name(persistentVolumeClaimRuntime.Name).
		Body(persistentVolumeClaimRuntime).
		Do().
		Into(result)
	return
}

// Delete takes name of the persistentVolumeClaimRuntime and deletes it. Returns an error if one occurs.
func (c *persistentVolumeClaimRuntimes) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("persistentvolumeclaimruntimes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *persistentVolumeClaimRuntimes) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("persistentvolumeclaimruntimes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched persistentVolumeClaimRuntime.
func (c *persistentVolumeClaimRuntimes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PersistentVolumeClaimRuntime, err error) {
	result = &v1.PersistentVolumeClaimRuntime{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("persistentvolumeclaimruntimes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}