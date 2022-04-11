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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta2

import (
	v1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackingImageManagerLister helps list BackingImageManagers.
type BackingImageManagerLister interface {
	// List lists all BackingImageManagers in the indexer.
	List(selector labels.Selector) (ret []*v1beta2.BackingImageManager, err error)
	// BackingImageManagers returns an object that can list and get BackingImageManagers.
	BackingImageManagers(namespace string) BackingImageManagerNamespaceLister
	BackingImageManagerListerExpansion
}

// backingImageManagerLister implements the BackingImageManagerLister interface.
type backingImageManagerLister struct {
	indexer cache.Indexer
}

// NewBackingImageManagerLister returns a new BackingImageManagerLister.
func NewBackingImageManagerLister(indexer cache.Indexer) BackingImageManagerLister {
	return &backingImageManagerLister{indexer: indexer}
}

// List lists all BackingImageManagers in the indexer.
func (s *backingImageManagerLister) List(selector labels.Selector) (ret []*v1beta2.BackingImageManager, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta2.BackingImageManager))
	})
	return ret, err
}

// BackingImageManagers returns an object that can list and get BackingImageManagers.
func (s *backingImageManagerLister) BackingImageManagers(namespace string) BackingImageManagerNamespaceLister {
	return backingImageManagerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackingImageManagerNamespaceLister helps list and get BackingImageManagers.
type BackingImageManagerNamespaceLister interface {
	// List lists all BackingImageManagers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta2.BackingImageManager, err error)
	// Get retrieves the BackingImageManager from the indexer for a given namespace and name.
	Get(name string) (*v1beta2.BackingImageManager, error)
	BackingImageManagerNamespaceListerExpansion
}

// backingImageManagerNamespaceLister implements the BackingImageManagerNamespaceLister
// interface.
type backingImageManagerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BackingImageManagers in the indexer for a given namespace.
func (s backingImageManagerNamespaceLister) List(selector labels.Selector) (ret []*v1beta2.BackingImageManager, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta2.BackingImageManager))
	})
	return ret, err
}

// Get retrieves the BackingImageManager from the indexer for a given namespace and name.
func (s backingImageManagerNamespaceLister) Get(name string) (*v1beta2.BackingImageManager, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta2.Resource("backingimagemanager"), name)
	}
	return obj.(*v1beta2.BackingImageManager), nil
}