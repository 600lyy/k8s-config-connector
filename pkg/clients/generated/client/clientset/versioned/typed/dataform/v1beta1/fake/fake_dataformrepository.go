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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dataform/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataformRepositories implements DataformRepositoryInterface
type FakeDataformRepositories struct {
	Fake *FakeDataformV1beta1
	ns   string
}

var dataformrepositoriesResource = v1beta1.SchemeGroupVersion.WithResource("dataformrepositories")

var dataformrepositoriesKind = v1beta1.SchemeGroupVersion.WithKind("DataformRepository")

// Get takes name of the dataformRepository, and returns the corresponding dataformRepository object, and an error if there is any.
func (c *FakeDataformRepositories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DataformRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dataformrepositoriesResource, c.ns, name), &v1beta1.DataformRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataformRepository), err
}

// List takes label and field selectors, and returns the list of DataformRepositories that match those selectors.
func (c *FakeDataformRepositories) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DataformRepositoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dataformrepositoriesResource, dataformrepositoriesKind, c.ns, opts), &v1beta1.DataformRepositoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DataformRepositoryList{ListMeta: obj.(*v1beta1.DataformRepositoryList).ListMeta}
	for _, item := range obj.(*v1beta1.DataformRepositoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataformRepositories.
func (c *FakeDataformRepositories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dataformrepositoriesResource, c.ns, opts))

}

// Create takes the representation of a dataformRepository and creates it.  Returns the server's representation of the dataformRepository, and an error, if there is any.
func (c *FakeDataformRepositories) Create(ctx context.Context, dataformRepository *v1beta1.DataformRepository, opts v1.CreateOptions) (result *v1beta1.DataformRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dataformrepositoriesResource, c.ns, dataformRepository), &v1beta1.DataformRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataformRepository), err
}

// Update takes the representation of a dataformRepository and updates it. Returns the server's representation of the dataformRepository, and an error, if there is any.
func (c *FakeDataformRepositories) Update(ctx context.Context, dataformRepository *v1beta1.DataformRepository, opts v1.UpdateOptions) (result *v1beta1.DataformRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dataformrepositoriesResource, c.ns, dataformRepository), &v1beta1.DataformRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataformRepository), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataformRepositories) UpdateStatus(ctx context.Context, dataformRepository *v1beta1.DataformRepository, opts v1.UpdateOptions) (*v1beta1.DataformRepository, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dataformrepositoriesResource, "status", c.ns, dataformRepository), &v1beta1.DataformRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataformRepository), err
}

// Delete takes name of the dataformRepository and deletes it. Returns an error if one occurs.
func (c *FakeDataformRepositories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dataformrepositoriesResource, c.ns, name, opts), &v1beta1.DataformRepository{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataformRepositories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dataformrepositoriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DataformRepositoryList{})
	return err
}

// Patch applies the patch and returns the patched dataformRepository.
func (c *FakeDataformRepositories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataformRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dataformrepositoriesResource, c.ns, name, pt, data, subresources...), &v1beta1.DataformRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataformRepository), err
}
