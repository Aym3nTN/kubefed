/*
Copyright 2018 The Kubernetes Authors.

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
package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/federation/v1alpha1"
	scheme "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FederatedJobsGetter has a method to return a FederatedJobInterface.
// A group's client should implement this interface.
type FederatedJobsGetter interface {
	FederatedJobs(namespace string) FederatedJobInterface
}

// FederatedJobInterface has methods to work with FederatedJob resources.
type FederatedJobInterface interface {
	Create(*v1alpha1.FederatedJob) (*v1alpha1.FederatedJob, error)
	Update(*v1alpha1.FederatedJob) (*v1alpha1.FederatedJob, error)
	UpdateStatus(*v1alpha1.FederatedJob) (*v1alpha1.FederatedJob, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.FederatedJob, error)
	List(opts v1.ListOptions) (*v1alpha1.FederatedJobList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedJob, err error)
	FederatedJobExpansion
}

// federatedJobs implements FederatedJobInterface
type federatedJobs struct {
	client rest.Interface
	ns     string
}

// newFederatedJobs returns a FederatedJobs
func newFederatedJobs(c *FederationV1alpha1Client, namespace string) *federatedJobs {
	return &federatedJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedJob, and returns the corresponding federatedJob object, and an error if there is any.
func (c *federatedJobs) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedJob, err error) {
	result = &v1alpha1.FederatedJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedJobs that match those selectors.
func (c *federatedJobs) List(opts v1.ListOptions) (result *v1alpha1.FederatedJobList, err error) {
	result = &v1alpha1.FederatedJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedJobs.
func (c *federatedJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a federatedJob and creates it.  Returns the server's representation of the federatedJob, and an error, if there is any.
func (c *federatedJobs) Create(federatedJob *v1alpha1.FederatedJob) (result *v1alpha1.FederatedJob, err error) {
	result = &v1alpha1.FederatedJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedjobs").
		Body(federatedJob).
		Do().
		Into(result)
	return
}

// Update takes the representation of a federatedJob and updates it. Returns the server's representation of the federatedJob, and an error, if there is any.
func (c *federatedJobs) Update(federatedJob *v1alpha1.FederatedJob) (result *v1alpha1.FederatedJob, err error) {
	result = &v1alpha1.FederatedJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedjobs").
		Name(federatedJob.Name).
		Body(federatedJob).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *federatedJobs) UpdateStatus(federatedJob *v1alpha1.FederatedJob) (result *v1alpha1.FederatedJob, err error) {
	result = &v1alpha1.FederatedJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedjobs").
		Name(federatedJob.Name).
		SubResource("status").
		Body(federatedJob).
		Do().
		Into(result)
	return
}

// Delete takes name of the federatedJob and deletes it. Returns an error if one occurs.
func (c *federatedJobs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched federatedJob.
func (c *federatedJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedJob, err error) {
	result = &v1alpha1.FederatedJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}