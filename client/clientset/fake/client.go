package fake

import (
	_ "github.com/appscode/k8s-addons/api/install"
	"k8s.io/kubernetes/pkg/client/clientset_generated/release_1_5/fake"
	"k8s.io/kubernetes/pkg/runtime"
)

type ClientSets struct {
	*fake.Clientset
	ACExtensionClient *FakeExtensionClient
}

func NewFakeClient(objects ...runtime.Object) *ClientSets {
	return &ClientSets{
		Clientset:         fake.NewSimpleClientset(objects...),
		ACExtensionClient: NewFakeExtensionClient(objects...),
	}
}
