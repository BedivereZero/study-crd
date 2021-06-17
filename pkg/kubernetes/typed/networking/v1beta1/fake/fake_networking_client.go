package fake

import (
	v1beta1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeNetworkingV1beta1 struct {
	*testing.Fake
}

func (c *FakeNetworkingV1beta1) IngressClasses() v1beta1.IngressClassInterface {
	return &FakeIngressClasses{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNetworkingV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
