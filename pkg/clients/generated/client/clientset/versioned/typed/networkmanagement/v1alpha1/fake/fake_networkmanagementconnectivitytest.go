// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkmanagement/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkManagementConnectivityTests implements NetworkManagementConnectivityTestInterface
type FakeNetworkManagementConnectivityTests struct {
	Fake *FakeNetworkmanagementV1alpha1
	ns   string
}

var networkmanagementconnectivitytestsResource = v1alpha1.SchemeGroupVersion.WithResource("networkmanagementconnectivitytests")

var networkmanagementconnectivitytestsKind = v1alpha1.SchemeGroupVersion.WithKind("NetworkManagementConnectivityTest")

// Get takes name of the networkManagementConnectivityTest, and returns the corresponding networkManagementConnectivityTest object, and an error if there is any.
func (c *FakeNetworkManagementConnectivityTests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkManagementConnectivityTest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkmanagementconnectivitytestsResource, c.ns, name), &v1alpha1.NetworkManagementConnectivityTest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkManagementConnectivityTest), err
}

// List takes label and field selectors, and returns the list of NetworkManagementConnectivityTests that match those selectors.
func (c *FakeNetworkManagementConnectivityTests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkManagementConnectivityTestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkmanagementconnectivitytestsResource, networkmanagementconnectivitytestsKind, c.ns, opts), &v1alpha1.NetworkManagementConnectivityTestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkManagementConnectivityTestList{ListMeta: obj.(*v1alpha1.NetworkManagementConnectivityTestList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkManagementConnectivityTestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkManagementConnectivityTests.
func (c *FakeNetworkManagementConnectivityTests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkmanagementconnectivitytestsResource, c.ns, opts))

}

// Create takes the representation of a networkManagementConnectivityTest and creates it.  Returns the server's representation of the networkManagementConnectivityTest, and an error, if there is any.
func (c *FakeNetworkManagementConnectivityTests) Create(ctx context.Context, networkManagementConnectivityTest *v1alpha1.NetworkManagementConnectivityTest, opts v1.CreateOptions) (result *v1alpha1.NetworkManagementConnectivityTest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkmanagementconnectivitytestsResource, c.ns, networkManagementConnectivityTest), &v1alpha1.NetworkManagementConnectivityTest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkManagementConnectivityTest), err
}

// Update takes the representation of a networkManagementConnectivityTest and updates it. Returns the server's representation of the networkManagementConnectivityTest, and an error, if there is any.
func (c *FakeNetworkManagementConnectivityTests) Update(ctx context.Context, networkManagementConnectivityTest *v1alpha1.NetworkManagementConnectivityTest, opts v1.UpdateOptions) (result *v1alpha1.NetworkManagementConnectivityTest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkmanagementconnectivitytestsResource, c.ns, networkManagementConnectivityTest), &v1alpha1.NetworkManagementConnectivityTest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkManagementConnectivityTest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkManagementConnectivityTests) UpdateStatus(ctx context.Context, networkManagementConnectivityTest *v1alpha1.NetworkManagementConnectivityTest, opts v1.UpdateOptions) (*v1alpha1.NetworkManagementConnectivityTest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkmanagementconnectivitytestsResource, "status", c.ns, networkManagementConnectivityTest), &v1alpha1.NetworkManagementConnectivityTest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkManagementConnectivityTest), err
}

// Delete takes name of the networkManagementConnectivityTest and deletes it. Returns an error if one occurs.
func (c *FakeNetworkManagementConnectivityTests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkmanagementconnectivitytestsResource, c.ns, name, opts), &v1alpha1.NetworkManagementConnectivityTest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkManagementConnectivityTests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkmanagementconnectivitytestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkManagementConnectivityTestList{})
	return err
}

// Patch applies the patch and returns the patched networkManagementConnectivityTest.
func (c *FakeNetworkManagementConnectivityTests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkManagementConnectivityTest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkmanagementconnectivitytestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkManagementConnectivityTest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkManagementConnectivityTest), err
}
