package fake

import (
	"flag"

	_ "github.com/appscode/k8s-addons/api/install"
	"github.com/appscode/log"
	"k8s.io/kubernetes/pkg/client/clientset_generated/release_1_5/fake"
	"k8s.io/kubernetes/pkg/runtime"
)

func init() {
	if f := flag.Lookup("test.v"); f == nil {
		log.Fatalln("Unable to execute fake package while not in Test Envorinment")
	}
}

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
