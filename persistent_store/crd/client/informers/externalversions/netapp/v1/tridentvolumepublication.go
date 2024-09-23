// Copyright 2023 NetApp, Inc. All Rights Reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	netappv1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	versioned "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned"
	internalinterfaces "github.com/netapp/trident/persistent_store/crd/client/informers/externalversions/internalinterfaces"
	v1 "github.com/netapp/trident/persistent_store/crd/client/listers/netapp/v1"
)

// TridentVolumePublicationInformer provides access to a shared informer and lister for
// TridentVolumePublications.
type TridentVolumePublicationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TridentVolumePublicationLister
}

type tridentVolumePublicationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTridentVolumePublicationInformer constructs a new informer for TridentVolumePublication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTridentVolumePublicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTridentVolumePublicationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTridentVolumePublicationInformer constructs a new informer for TridentVolumePublication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTridentVolumePublicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TridentV1().TridentVolumePublications(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TridentV1().TridentVolumePublications(namespace).Watch(context.TODO(), options)
			},
		},
		&netappv1.TridentVolumePublication{},
		resyncPeriod,
		indexers,
	)
}

func (f *tridentVolumePublicationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTridentVolumePublicationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tridentVolumePublicationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&netappv1.TridentVolumePublication{}, f.defaultInformer)
}

func (f *tridentVolumePublicationInformer) Lister() v1.TridentVolumePublicationLister {
	return v1.NewTridentVolumePublicationLister(f.Informer().GetIndexer())
}
