package watcher

import (
	acs "github.com/appscode/k8s-addons/client/clientset"
	kapi "k8s.io/kubernetes/pkg/api"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/watch"
)

func ingressListFunc(c clientset.Interface) func(kapi.ListOptions) (runtime.Object, error) {
	return func(opts kapi.ListOptions) (runtime.Object, error) {
		return c.Extensions().Ingresses(kapi.NamespaceAll).List(opts)
	}
}

func ingressWatchFunc(c clientset.Interface) func(options kapi.ListOptions) (watch.Interface, error) {
	return func(options kapi.ListOptions) (watch.Interface, error) {
		return c.Extensions().Ingresses(kapi.NamespaceAll).Watch(options)
	}
}

func extIngressListFunc(c acs.AppsCodeExtensionInterface) func(kapi.ListOptions) (runtime.Object, error) {
	return func(opts kapi.ListOptions) (runtime.Object, error) {
		return c.Ingress(kapi.NamespaceAll).List(opts)
	}
}

func extIngressWatchFunc(c acs.AppsCodeExtensionInterface) func(options kapi.ListOptions) (watch.Interface, error) {
	return func(options kapi.ListOptions) (watch.Interface, error) {
		return c.Ingress(kapi.NamespaceAll).Watch(options)
	}
}

func daemonSetListFunc(c clientset.Interface) func(kapi.ListOptions) (runtime.Object, error) {
	return func(opts kapi.ListOptions) (runtime.Object, error) {
		return c.Extensions().DaemonSets(kapi.NamespaceAll).List(opts)
	}
}

func daemonSetWatchFunc(c clientset.Interface) func(options kapi.ListOptions) (watch.Interface, error) {
	return func(options kapi.ListOptions) (watch.Interface, error) {
		return c.Extensions().DaemonSets(kapi.NamespaceAll).Watch(options)
	}
}

func replicaSetListFunc(c clientset.Interface) func(kapi.ListOptions) (runtime.Object, error) {
	return func(opts kapi.ListOptions) (runtime.Object, error) {
		return c.Extensions().ReplicaSets(kapi.NamespaceAll).List(opts)
	}
}

func replicaSetWatchFunc(c clientset.Interface) func(options kapi.ListOptions) (watch.Interface, error) {
	return func(options kapi.ListOptions) (watch.Interface, error) {
		return c.Extensions().ReplicaSets(kapi.NamespaceAll).Watch(options)
	}
}

func petSetListFunc(c clientset.Interface) func(kapi.ListOptions) (runtime.Object, error) {
	return func(opts kapi.ListOptions) (runtime.Object, error) {
		return c.Apps().StatefulSets(kapi.NamespaceAll).List(opts)
	}
}

func petSetWatchFunc(c clientset.Interface) func(options kapi.ListOptions) (watch.Interface, error) {
	return func(options kapi.ListOptions) (watch.Interface, error) {
		return c.Apps().StatefulSets(kapi.NamespaceAll).Watch(options)
	}
}

func alertListFunc(c acs.AppsCodeExtensionInterface) func(kapi.ListOptions) (runtime.Object, error) {
	return func(opts kapi.ListOptions) (runtime.Object, error) {
		return c.Alert(kapi.NamespaceAll).List(opts)
	}
}

func alertWatchFunc(c acs.AppsCodeExtensionInterface) func(options kapi.ListOptions) (watch.Interface, error) {
	return func(options kapi.ListOptions) (watch.Interface, error) {
		return c.Alert(kapi.NamespaceAll).Watch(options)
	}
}

func certificateListFunc(c acs.AppsCodeExtensionInterface) func(kapi.ListOptions) (runtime.Object, error) {
	return func(opts kapi.ListOptions) (runtime.Object, error) {
		return c.Certificate(kapi.NamespaceAll).List(opts)
	}
}

func certificateWatchFunc(c acs.AppsCodeExtensionInterface) func(options kapi.ListOptions) (watch.Interface, error) {
	return func(options kapi.ListOptions) (watch.Interface, error) {
		return c.Certificate(kapi.NamespaceAll).Watch(options)
	}
}
