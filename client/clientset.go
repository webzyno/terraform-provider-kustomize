package client

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/kustomize/api/krusty"
)

type ClientSet struct {
	Kustomizer      *krusty.Kustomizer
	DynamicClient   dynamic.Interface
	DiscoveryClient discovery.DiscoveryInterface
}
