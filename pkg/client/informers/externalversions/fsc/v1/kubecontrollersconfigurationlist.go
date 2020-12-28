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

	fscv1 "github.com/henderiw/fsc-lib-go/pkg/apis/fsc/v1"
	versioned "github.com/henderiw/fsc-lib-go/pkg/client/clientset/versioned"
	internalinterfaces "github.com/henderiw/fsc-lib-go/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/henderiw/fsc-lib-go/pkg/client/listers/fsc/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubeControllersConfigurationListInformer provides access to a shared informer and lister for
// KubeControllersConfigurationLists.
type KubeControllersConfigurationListInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.KubeControllersConfigurationListLister
}

type kubeControllersConfigurationListInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKubeControllersConfigurationListInformer constructs a new informer for KubeControllersConfigurationList type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubeControllersConfigurationListInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubeControllersConfigurationListInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKubeControllersConfigurationListInformer constructs a new informer for KubeControllersConfigurationList type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubeControllersConfigurationListInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FscV1().KubeControllersConfigurationLists(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FscV1().KubeControllersConfigurationLists(namespace).Watch(context.TODO(), options)
			},
		},
		&fscv1.KubeControllersConfigurationList{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubeControllersConfigurationListInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubeControllersConfigurationListInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubeControllersConfigurationListInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&fscv1.KubeControllersConfigurationList{}, f.defaultInformer)
}

func (f *kubeControllersConfigurationListInformer) Lister() v1.KubeControllersConfigurationListLister {
	return v1.NewKubeControllersConfigurationListLister(f.Informer().GetIndexer())
}
