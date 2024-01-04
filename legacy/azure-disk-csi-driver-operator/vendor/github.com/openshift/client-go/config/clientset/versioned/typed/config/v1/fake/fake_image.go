// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openshift/api/config/v1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImages implements ImageInterface
type FakeImages struct {
	Fake *FakeConfigV1
}

var imagesResource = v1.SchemeGroupVersion.WithResource("images")

var imagesKind = v1.SchemeGroupVersion.WithKind("Image")

// Get takes name of the image, and returns the corresponding image object, and an error if there is any.
func (c *FakeImages) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Image, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(imagesResource, name), &v1.Image{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Image), err
}

// List takes label and field selectors, and returns the list of Images that match those selectors.
func (c *FakeImages) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ImageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(imagesResource, imagesKind, opts), &v1.ImageList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ImageList{ListMeta: obj.(*v1.ImageList).ListMeta}
	for _, item := range obj.(*v1.ImageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested images.
func (c *FakeImages) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(imagesResource, opts))
}

// Create takes the representation of a image and creates it.  Returns the server's representation of the image, and an error, if there is any.
func (c *FakeImages) Create(ctx context.Context, image *v1.Image, opts metav1.CreateOptions) (result *v1.Image, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(imagesResource, image), &v1.Image{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Image), err
}

// Update takes the representation of a image and updates it. Returns the server's representation of the image, and an error, if there is any.
func (c *FakeImages) Update(ctx context.Context, image *v1.Image, opts metav1.UpdateOptions) (result *v1.Image, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(imagesResource, image), &v1.Image{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Image), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImages) UpdateStatus(ctx context.Context, image *v1.Image, opts metav1.UpdateOptions) (*v1.Image, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(imagesResource, "status", image), &v1.Image{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Image), err
}

// Delete takes name of the image and deletes it. Returns an error if one occurs.
func (c *FakeImages) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(imagesResource, name, opts), &v1.Image{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImages) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(imagesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ImageList{})
	return err
}

// Patch applies the patch and returns the patched image.
func (c *FakeImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Image, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(imagesResource, name, pt, data, subresources...), &v1.Image{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Image), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied image.
func (c *FakeImages) Apply(ctx context.Context, image *configv1.ImageApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Image, err error) {
	if image == nil {
		return nil, fmt.Errorf("image provided to Apply must not be nil")
	}
	data, err := json.Marshal(image)
	if err != nil {
		return nil, err
	}
	name := image.Name
	if name == nil {
		return nil, fmt.Errorf("image.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(imagesResource, *name, types.ApplyPatchType, data), &v1.Image{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Image), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeImages) ApplyStatus(ctx context.Context, image *configv1.ImageApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Image, err error) {
	if image == nil {
		return nil, fmt.Errorf("image provided to Apply must not be nil")
	}
	data, err := json.Marshal(image)
	if err != nil {
		return nil, err
	}
	name := image.Name
	if name == nil {
		return nil, fmt.Errorf("image.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(imagesResource, *name, types.ApplyPatchType, data, "status"), &v1.Image{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Image), err
}