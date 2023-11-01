/*
Copyright 2023 The KubeVirt Authors.

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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// FakeDataVolumes implements DataVolumeInterface
type FakeDataVolumes struct {
	Fake *FakeCdiV1beta1
	ns   string
}

var datavolumesResource = schema.GroupVersionResource{Group: "cdi.kubevirt.io", Version: "v1beta1", Resource: "datavolumes"}

var datavolumesKind = schema.GroupVersionKind{Group: "cdi.kubevirt.io", Version: "v1beta1", Kind: "DataVolume"}

// Get takes name of the dataVolume, and returns the corresponding dataVolume object, and an error if there is any.
func (c *FakeDataVolumes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DataVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datavolumesResource, c.ns, name), &v1beta1.DataVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataVolume), err
}

// List takes label and field selectors, and returns the list of DataVolumes that match those selectors.
func (c *FakeDataVolumes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DataVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datavolumesResource, datavolumesKind, c.ns, opts), &v1beta1.DataVolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DataVolumeList{ListMeta: obj.(*v1beta1.DataVolumeList).ListMeta}
	for _, item := range obj.(*v1beta1.DataVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataVolumes.
func (c *FakeDataVolumes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datavolumesResource, c.ns, opts))

}

// Create takes the representation of a dataVolume and creates it.  Returns the server's representation of the dataVolume, and an error, if there is any.
func (c *FakeDataVolumes) Create(ctx context.Context, dataVolume *v1beta1.DataVolume, opts v1.CreateOptions) (result *v1beta1.DataVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datavolumesResource, c.ns, dataVolume), &v1beta1.DataVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataVolume), err
}

// Update takes the representation of a dataVolume and updates it. Returns the server's representation of the dataVolume, and an error, if there is any.
func (c *FakeDataVolumes) Update(ctx context.Context, dataVolume *v1beta1.DataVolume, opts v1.UpdateOptions) (result *v1beta1.DataVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datavolumesResource, c.ns, dataVolume), &v1beta1.DataVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataVolumes) UpdateStatus(ctx context.Context, dataVolume *v1beta1.DataVolume, opts v1.UpdateOptions) (*v1beta1.DataVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datavolumesResource, "status", c.ns, dataVolume), &v1beta1.DataVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataVolume), err
}

// Delete takes name of the dataVolume and deletes it. Returns an error if one occurs.
func (c *FakeDataVolumes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datavolumesResource, c.ns, name), &v1beta1.DataVolume{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataVolumes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datavolumesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DataVolumeList{})
	return err
}

// Patch applies the patch and returns the patched dataVolume.
func (c *FakeDataVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datavolumesResource, c.ns, name, pt, data, subresources...), &v1beta1.DataVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataVolume), err
}
