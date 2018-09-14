/*
Copyright 2018 The Kube Vault Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/kubevault/operator/apis/core/v1alpha1"
	scheme "github.com/kubevault/operator/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VaultServerVersionsGetter has a method to return a VaultServerVersionInterface.
// A group's client should implement this interface.
type VaultServerVersionsGetter interface {
	VaultServerVersions() VaultServerVersionInterface
}

// VaultServerVersionInterface has methods to work with VaultServerVersion resources.
type VaultServerVersionInterface interface {
	Create(*v1alpha1.VaultServerVersion) (*v1alpha1.VaultServerVersion, error)
	Update(*v1alpha1.VaultServerVersion) (*v1alpha1.VaultServerVersion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VaultServerVersion, error)
	List(opts v1.ListOptions) (*v1alpha1.VaultServerVersionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VaultServerVersion, err error)
	VaultServerVersionExpansion
}

// vaultServerVersions implements VaultServerVersionInterface
type vaultServerVersions struct {
	client rest.Interface
}

// newVaultServerVersions returns a VaultServerVersions
func newVaultServerVersions(c *CoreV1alpha1Client) *vaultServerVersions {
	return &vaultServerVersions{
		client: c.RESTClient(),
	}
}

// Get takes name of the vaultServerVersion, and returns the corresponding vaultServerVersion object, and an error if there is any.
func (c *vaultServerVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.VaultServerVersion, err error) {
	result = &v1alpha1.VaultServerVersion{}
	err = c.client.Get().
		Resource("vaultserverversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VaultServerVersions that match those selectors.
func (c *vaultServerVersions) List(opts v1.ListOptions) (result *v1alpha1.VaultServerVersionList, err error) {
	result = &v1alpha1.VaultServerVersionList{}
	err = c.client.Get().
		Resource("vaultserverversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vaultServerVersions.
func (c *vaultServerVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("vaultserverversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a vaultServerVersion and creates it.  Returns the server's representation of the vaultServerVersion, and an error, if there is any.
func (c *vaultServerVersions) Create(vaultServerVersion *v1alpha1.VaultServerVersion) (result *v1alpha1.VaultServerVersion, err error) {
	result = &v1alpha1.VaultServerVersion{}
	err = c.client.Post().
		Resource("vaultserverversions").
		Body(vaultServerVersion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a vaultServerVersion and updates it. Returns the server's representation of the vaultServerVersion, and an error, if there is any.
func (c *vaultServerVersions) Update(vaultServerVersion *v1alpha1.VaultServerVersion) (result *v1alpha1.VaultServerVersion, err error) {
	result = &v1alpha1.VaultServerVersion{}
	err = c.client.Put().
		Resource("vaultserverversions").
		Name(vaultServerVersion.Name).
		Body(vaultServerVersion).
		Do().
		Into(result)
	return
}

// Delete takes name of the vaultServerVersion and deletes it. Returns an error if one occurs.
func (c *vaultServerVersions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("vaultserverversions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vaultServerVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("vaultserverversions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched vaultServerVersion.
func (c *vaultServerVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VaultServerVersion, err error) {
	result = &v1alpha1.VaultServerVersion{}
	err = c.client.Patch(pt).
		Resource("vaultserverversions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
