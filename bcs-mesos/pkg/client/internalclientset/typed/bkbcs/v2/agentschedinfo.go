/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/Tencent/bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"
	scheme "github.com/Tencent/bk-bcs/bcs-mesos/pkg/client/internalclientset/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AgentSchedInfosGetter has a method to return a AgentSchedInfoInterface.
// A group's client should implement this interface.
type AgentSchedInfosGetter interface {
	AgentSchedInfos(namespace string) AgentSchedInfoInterface
}

// AgentSchedInfoInterface has methods to work with AgentSchedInfo resources.
type AgentSchedInfoInterface interface {
	Create(*v2.AgentSchedInfo) (*v2.AgentSchedInfo, error)
	Update(*v2.AgentSchedInfo) (*v2.AgentSchedInfo, error)
	UpdateStatus(*v2.AgentSchedInfo) (*v2.AgentSchedInfo, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v2.AgentSchedInfo, error)
	List(opts v1.ListOptions) (*v2.AgentSchedInfoList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.AgentSchedInfo, err error)
	AgentSchedInfoExpansion
}

// agentSchedInfos implements AgentSchedInfoInterface
type agentSchedInfos struct {
	client rest.Interface
	ns     string
}

// newAgentSchedInfos returns a AgentSchedInfos
func newAgentSchedInfos(c *BkbcsV2Client, namespace string) *agentSchedInfos {
	return &agentSchedInfos{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the agentSchedInfo, and returns the corresponding agentSchedInfo object, and an error if there is any.
func (c *agentSchedInfos) Get(name string, options v1.GetOptions) (result *v2.AgentSchedInfo, err error) {
	result = &v2.AgentSchedInfo{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("agentschedinfos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AgentSchedInfos that match those selectors.
func (c *agentSchedInfos) List(opts v1.ListOptions) (result *v2.AgentSchedInfoList, err error) {
	result = &v2.AgentSchedInfoList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("agentschedinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested agentSchedInfos.
func (c *agentSchedInfos) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("agentschedinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a agentSchedInfo and creates it.  Returns the server's representation of the agentSchedInfo, and an error, if there is any.
func (c *agentSchedInfos) Create(agentSchedInfo *v2.AgentSchedInfo) (result *v2.AgentSchedInfo, err error) {
	result = &v2.AgentSchedInfo{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("agentschedinfos").
		Body(agentSchedInfo).
		Do().
		Into(result)
	return
}

// Update takes the representation of a agentSchedInfo and updates it. Returns the server's representation of the agentSchedInfo, and an error, if there is any.
func (c *agentSchedInfos) Update(agentSchedInfo *v2.AgentSchedInfo) (result *v2.AgentSchedInfo, err error) {
	result = &v2.AgentSchedInfo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("agentschedinfos").
		Name(agentSchedInfo.Name).
		Body(agentSchedInfo).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *agentSchedInfos) UpdateStatus(agentSchedInfo *v2.AgentSchedInfo) (result *v2.AgentSchedInfo, err error) {
	result = &v2.AgentSchedInfo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("agentschedinfos").
		Name(agentSchedInfo.Name).
		SubResource("status").
		Body(agentSchedInfo).
		Do().
		Into(result)
	return
}

// Delete takes name of the agentSchedInfo and deletes it. Returns an error if one occurs.
func (c *agentSchedInfos) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("agentschedinfos").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *agentSchedInfos) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("agentschedinfos").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched agentSchedInfo.
func (c *agentSchedInfos) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.AgentSchedInfo, err error) {
	result = &v2.AgentSchedInfo{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("agentschedinfos").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
