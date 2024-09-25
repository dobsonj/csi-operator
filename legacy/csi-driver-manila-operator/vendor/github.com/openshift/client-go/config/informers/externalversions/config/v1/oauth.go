// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	configv1 "github.com/openshift/api/config/v1"
	versioned "github.com/openshift/client-go/config/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/config/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/config/listers/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OAuthInformer provides access to a shared informer and lister for
// OAuths.
type OAuthInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.OAuthLister
}

type oAuthInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewOAuthInformer constructs a new informer for OAuth type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOAuthInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOAuthInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredOAuthInformer constructs a new informer for OAuth type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOAuthInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1().OAuths().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1().OAuths().Watch(context.TODO(), options)
			},
		},
		&configv1.OAuth{},
		resyncPeriod,
		indexers,
	)
}

func (f *oAuthInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOAuthInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *oAuthInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1.OAuth{}, f.defaultInformer)
}

func (f *oAuthInformer) Lister() v1.OAuthLister {
	return v1.NewOAuthLister(f.Informer().GetIndexer())
}
