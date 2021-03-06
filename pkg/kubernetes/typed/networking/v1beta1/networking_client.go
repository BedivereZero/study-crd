package v1beta1

import (
	v1beta1 "github.com/BedivereZero/study-crd/api/networking/v1beta1"
	"k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type NetworkingV1beta1Interface interface {
	RESTClient() rest.Interface
	IngressClassesGetter
}

// NetworkingV1beta1Client is used to interact with features provided by the networking.k8s.io group.
type NetworkingV1beta1Client struct {
	restClient rest.Interface
}

func (c *NetworkingV1beta1Client) IngressClasses() IngressClassInterface {
	return newIngressClasses(c)
}

// NewForConfig creates a new NetworkingV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*NetworkingV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &NetworkingV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new NetworkingV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *NetworkingV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new NetworkingV1beta1Client for the given RESTClient.
func New(c rest.Interface) *NetworkingV1beta1Client {
	return &NetworkingV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *NetworkingV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
