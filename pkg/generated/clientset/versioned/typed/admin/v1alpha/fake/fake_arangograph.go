//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha "github.com/arangodb/kube-arangodb/pkg/apis/admin/v1alpha"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArangoGraphs implements ArangoGraphInterface
type FakeArangoGraphs struct {
	Fake *FakeDatabaseadminV1alpha
	ns   string
}

var arangographsResource = schema.GroupVersionResource{Group: "databaseadmin.arangodb.com", Version: "v1alpha", Resource: "arangographs"}

var arangographsKind = schema.GroupVersionKind{Group: "databaseadmin.arangodb.com", Version: "v1alpha", Kind: "ArangoGraph"}

// Get takes name of the arangoGraph, and returns the corresponding arangoGraph object, and an error if there is any.
func (c *FakeArangoGraphs) Get(name string, options v1.GetOptions) (result *v1alpha.ArangoGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(arangographsResource, c.ns, name), &v1alpha.ArangoGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoGraph), err
}

// List takes label and field selectors, and returns the list of ArangoGraphs that match those selectors.
func (c *FakeArangoGraphs) List(opts v1.ListOptions) (result *v1alpha.ArangoGraphList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(arangographsResource, arangographsKind, c.ns, opts), &v1alpha.ArangoGraphList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha.ArangoGraphList{ListMeta: obj.(*v1alpha.ArangoGraphList).ListMeta}
	for _, item := range obj.(*v1alpha.ArangoGraphList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested arangoGraphs.
func (c *FakeArangoGraphs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(arangographsResource, c.ns, opts))

}

// Create takes the representation of a arangoGraph and creates it.  Returns the server's representation of the arangoGraph, and an error, if there is any.
func (c *FakeArangoGraphs) Create(arangoGraph *v1alpha.ArangoGraph) (result *v1alpha.ArangoGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(arangographsResource, c.ns, arangoGraph), &v1alpha.ArangoGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoGraph), err
}

// Update takes the representation of a arangoGraph and updates it. Returns the server's representation of the arangoGraph, and an error, if there is any.
func (c *FakeArangoGraphs) Update(arangoGraph *v1alpha.ArangoGraph) (result *v1alpha.ArangoGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(arangographsResource, c.ns, arangoGraph), &v1alpha.ArangoGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoGraph), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArangoGraphs) UpdateStatus(arangoGraph *v1alpha.ArangoGraph) (*v1alpha.ArangoGraph, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(arangographsResource, "status", c.ns, arangoGraph), &v1alpha.ArangoGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoGraph), err
}

// Delete takes name of the arangoGraph and deletes it. Returns an error if one occurs.
func (c *FakeArangoGraphs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(arangographsResource, c.ns, name), &v1alpha.ArangoGraph{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArangoGraphs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(arangographsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha.ArangoGraphList{})
	return err
}

// Patch applies the patch and returns the patched arangoGraph.
func (c *FakeArangoGraphs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.ArangoGraph, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(arangographsResource, c.ns, name, data, subresources...), &v1alpha.ArangoGraph{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoGraph), err
}