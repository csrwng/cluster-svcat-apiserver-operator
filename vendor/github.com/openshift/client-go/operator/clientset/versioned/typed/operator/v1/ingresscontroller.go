// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/openshift/api/operator/v1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IngressControllersGetter has a method to return a IngressControllerInterface.
// A group's client should implement this interface.
type IngressControllersGetter interface {
	IngressControllers(namespace string) IngressControllerInterface
}

// IngressControllerInterface has methods to work with IngressController resources.
type IngressControllerInterface interface {
	Create(*v1.IngressController) (*v1.IngressController, error)
	Update(*v1.IngressController) (*v1.IngressController, error)
	UpdateStatus(*v1.IngressController) (*v1.IngressController, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.IngressController, error)
	List(opts metav1.ListOptions) (*v1.IngressControllerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.IngressController, err error)
	IngressControllerExpansion
}

// ingressControllers implements IngressControllerInterface
type ingressControllers struct {
	client rest.Interface
	ns     string
}

// newIngressControllers returns a IngressControllers
func newIngressControllers(c *OperatorV1Client, namespace string) *ingressControllers {
	return &ingressControllers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ingressController, and returns the corresponding ingressController object, and an error if there is any.
func (c *ingressControllers) Get(name string, options metav1.GetOptions) (result *v1.IngressController, err error) {
	result = &v1.IngressController{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ingresscontrollers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IngressControllers that match those selectors.
func (c *ingressControllers) List(opts metav1.ListOptions) (result *v1.IngressControllerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.IngressControllerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ingresscontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ingressControllers.
func (c *ingressControllers) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ingresscontrollers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ingressController and creates it.  Returns the server's representation of the ingressController, and an error, if there is any.
func (c *ingressControllers) Create(ingressController *v1.IngressController) (result *v1.IngressController, err error) {
	result = &v1.IngressController{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ingresscontrollers").
		Body(ingressController).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ingressController and updates it. Returns the server's representation of the ingressController, and an error, if there is any.
func (c *ingressControllers) Update(ingressController *v1.IngressController) (result *v1.IngressController, err error) {
	result = &v1.IngressController{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ingresscontrollers").
		Name(ingressController.Name).
		Body(ingressController).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ingressControllers) UpdateStatus(ingressController *v1.IngressController) (result *v1.IngressController, err error) {
	result = &v1.IngressController{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ingresscontrollers").
		Name(ingressController.Name).
		SubResource("status").
		Body(ingressController).
		Do().
		Into(result)
	return
}

// Delete takes name of the ingressController and deletes it. Returns an error if one occurs.
func (c *ingressControllers) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ingresscontrollers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ingressControllers) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ingresscontrollers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ingressController.
func (c *ingressControllers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.IngressController, err error) {
	result = &v1.IngressController{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ingresscontrollers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
