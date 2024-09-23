// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
)

// TridentVersionLister helps list TridentVersions.
type TridentVersionLister interface {
	// List lists all TridentVersions in the indexer.
	List(selector labels.Selector) (ret []*v1.TridentVersion, err error)
	// TridentVersions returns an object that can list and get TridentVersions.
	TridentVersions(namespace string) TridentVersionNamespaceLister
	TridentVersionListerExpansion
}

// tridentVersionLister implements the TridentVersionLister interface.
type tridentVersionLister struct {
	indexer cache.Indexer
}

// NewTridentVersionLister returns a new TridentVersionLister.
func NewTridentVersionLister(indexer cache.Indexer) TridentVersionLister {
	return &tridentVersionLister{indexer: indexer}
}

// List lists all TridentVersions in the indexer.
func (s *tridentVersionLister) List(selector labels.Selector) (ret []*v1.TridentVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentVersion))
	})
	return ret, err
}

// TridentVersions returns an object that can list and get TridentVersions.
func (s *tridentVersionLister) TridentVersions(namespace string) TridentVersionNamespaceLister {
	return tridentVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TridentVersionNamespaceLister helps list and get TridentVersions.
type TridentVersionNamespaceLister interface {
	// List lists all TridentVersions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TridentVersion, err error)
	// Get retrieves the TridentVersion from the indexer for a given namespace and name.
	Get(name string) (*v1.TridentVersion, error)
	TridentVersionNamespaceListerExpansion
}

// tridentVersionNamespaceLister implements the TridentVersionNamespaceLister
// interface.
type tridentVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TridentVersions in the indexer for a given namespace.
func (s tridentVersionNamespaceLister) List(selector labels.Selector) (ret []*v1.TridentVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentVersion))
	})
	return ret, err
}

// Get retrieves the TridentVersion from the indexer for a given namespace and name.
func (s tridentVersionNamespaceLister) Get(name string) (*v1.TridentVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tridentversion"), name)
	}
	return obj.(*v1.TridentVersion), nil
}
