package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/BedivereZero/study-crd/api/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// IngressClassesGetter has a method to return a IngressClassInterface.
// A group's client should implement this interface.
type IngressClassesGetter interface {
	IngressClasses() IngressClassInterface
}

// IngressClassInterface has methods to work with IngressClass resources.
type IngressClassInterface interface {
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.IngressClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.IngressClassList, error)
	IngressClassExpansion
}

// ingressClasses implements IngressClassInterface
type ingressClasses struct {
	client rest.Interface
}

// newIngressClasses returns a IngressClasses
func newIngressClasses(c *NetworkingV1beta1Client) *ingressClasses {
	return &ingressClasses{
		client: c.RESTClient(),
	}
}

// Get takes name of the ingressClass, and returns the corresponding ingressClass object, and an error if there is any.
func (c *ingressClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.IngressClass, err error) {
	result = &v1beta1.IngressClass{}
	err = c.client.Get().
		Resource("ingressclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IngressClasses that match those selectors.
func (c *ingressClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.IngressClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.IngressClassList{}
	err = c.client.Get().
		Resource("ingressclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}
