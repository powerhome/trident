// Copyright 2024 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1 "github.com/netapp/trident/operator/crd/apis/netapp/v1"
	scheme "github.com/netapp/trident/operator/crd/client/clientset/versioned/scheme"
)

// TridentOrchestratorsGetter has a method to return a TridentOrchestratorInterface.
// A group's client should implement this interface.
type TridentOrchestratorsGetter interface {
	TridentOrchestrators() TridentOrchestratorInterface
}

// TridentOrchestratorInterface has methods to work with TridentOrchestrator resources.
type TridentOrchestratorInterface interface {
	Create(ctx context.Context, tridentOrchestrator *v1.TridentOrchestrator, opts metav1.CreateOptions) (*v1.TridentOrchestrator, error)
	Update(ctx context.Context, tridentOrchestrator *v1.TridentOrchestrator, opts metav1.UpdateOptions) (*v1.TridentOrchestrator, error)
	UpdateStatus(ctx context.Context, tridentOrchestrator *v1.TridentOrchestrator, opts metav1.UpdateOptions) (*v1.TridentOrchestrator, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TridentOrchestrator, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TridentOrchestratorList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentOrchestrator, err error)
	TridentOrchestratorExpansion
}

// tridentOrchestrators implements TridentOrchestratorInterface
type tridentOrchestrators struct {
	client rest.Interface
}

// newTridentOrchestrators returns a TridentOrchestrators
func newTridentOrchestrators(c *TridentV1Client) *tridentOrchestrators {
	return &tridentOrchestrators{
		client: c.RESTClient(),
	}
}

// Get takes name of the tridentOrchestrator, and returns the corresponding tridentOrchestrator object, and an error if there is any.
func (c *tridentOrchestrators) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TridentOrchestrator, err error) {
	result = &v1.TridentOrchestrator{}
	err = c.client.Get().
		Resource("tridentorchestrators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TridentOrchestrators that match those selectors.
func (c *tridentOrchestrators) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TridentOrchestratorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TridentOrchestratorList{}
	err = c.client.Get().
		Resource("tridentorchestrators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tridentOrchestrators.
func (c *tridentOrchestrators) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("tridentorchestrators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tridentOrchestrator and creates it.  Returns the server's representation of the tridentOrchestrator, and an error, if there is any.
func (c *tridentOrchestrators) Create(ctx context.Context, tridentOrchestrator *v1.TridentOrchestrator, opts metav1.CreateOptions) (result *v1.TridentOrchestrator, err error) {
	result = &v1.TridentOrchestrator{}
	err = c.client.Post().
		Resource("tridentorchestrators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentOrchestrator).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tridentOrchestrator and updates it. Returns the server's representation of the tridentOrchestrator, and an error, if there is any.
func (c *tridentOrchestrators) Update(ctx context.Context, tridentOrchestrator *v1.TridentOrchestrator, opts metav1.UpdateOptions) (result *v1.TridentOrchestrator, err error) {
	result = &v1.TridentOrchestrator{}
	err = c.client.Put().
		Resource("tridentorchestrators").
		Name(tridentOrchestrator.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentOrchestrator).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tridentOrchestrators) UpdateStatus(ctx context.Context, tridentOrchestrator *v1.TridentOrchestrator, opts metav1.UpdateOptions) (result *v1.TridentOrchestrator, err error) {
	result = &v1.TridentOrchestrator{}
	err = c.client.Put().
		Resource("tridentorchestrators").
		Name(tridentOrchestrator.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tridentOrchestrator).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tridentOrchestrator and deletes it. Returns an error if one occurs.
func (c *tridentOrchestrators) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("tridentorchestrators").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tridentOrchestrators) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("tridentorchestrators").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tridentOrchestrator.
func (c *tridentOrchestrators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TridentOrchestrator, err error) {
	result = &v1.TridentOrchestrator{}
	err = c.client.Patch(pt).
		Resource("tridentorchestrators").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
