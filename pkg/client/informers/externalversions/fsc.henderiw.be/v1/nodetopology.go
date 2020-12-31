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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	fschenderiwbev1 "github.com/henderiw/fsc-lib-go/pkg/apis/fsc.henderiw.be/v1"
	versioned "github.com/henderiw/fsc-lib-go/pkg/client/clientset/versioned"
	internalinterfaces "github.com/henderiw/fsc-lib-go/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/henderiw/fsc-lib-go/pkg/client/listers/fsc.henderiw.be/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NodeTopologyInformer provides access to a shared informer and lister for
// NodeTopologies.
type NodeTopologyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.NodeTopologyLister
}

type nodeTopologyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNodeTopologyInformer constructs a new informer for NodeTopology type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodeTopologyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodeTopologyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNodeTopologyInformer constructs a new informer for NodeTopology type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodeTopologyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FscV1().NodeTopologies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FscV1().NodeTopologies(namespace).Watch(context.TODO(), options)
			},
		},
		&fschenderiwbev1.NodeTopology{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodeTopologyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodeTopologyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodeTopologyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&fschenderiwbev1.NodeTopology{}, f.defaultInformer)
}

func (f *nodeTopologyInformer) Lister() v1.NodeTopologyLister {
	return v1.NewNodeTopologyLister(f.Informer().GetIndexer())
}
