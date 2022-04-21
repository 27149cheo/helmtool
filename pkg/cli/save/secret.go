package save

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// secretReader implements a corev1.SecretsInterface
type secretReader struct {
	v1.SecretInterface
	client    client.Client
	namespace string
}

func NewSecretReader(ns string, cl client.Client) *secretReader {
	return newSecretReader(ns, cl)
}

func newSecretReader(ns string, cl client.Client) *secretReader {
	return &secretReader{client: cl, namespace: ns}
}

func (s *secretReader) Get(ctx context.Context, name string, _ metav1.GetOptions) (*corev1.Secret, error) {
	obj := &corev1.Secret{}
	err := s.client.Get(ctx, types.NamespacedName{Namespace: s.namespace, Name: name}, obj)

	return obj, err
}

func (s *secretReader) List(ctx context.Context, opts metav1.ListOptions) (*corev1.SecretList, error) {
	objList := &corev1.SecretList{}
	err := s.client.List(ctx, objList, asListOptions(s.namespace, &opts))

	return objList, err
}

func asListOptions(ns string, o *metav1.ListOptions) *client.ListOptions {
	res := &client.ListOptions{}

	if o == nil {
		return res
	}

	res.Raw = o
	res.Limit = o.Limit
	res.Continue = o.Continue
	res.Namespace = ns

	if o.LabelSelector != "" {
		ls, _ := labels.Parse(o.LabelSelector)
		res.LabelSelector = ls
	}
	if o.FieldSelector != "" {
		fs, _ := fields.ParseSelector(o.FieldSelector)
		res.FieldSelector = fs
	}

	return res
}
