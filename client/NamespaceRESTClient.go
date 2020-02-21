package client

import (
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
)

// NamespaceRESTClient todo
type NamespaceRESTClient struct {
	rest.Interface
	RESTClient rest.Interface
	Namespace  string
}

// Post todo
func (s *NamespaceRESTClient) Post() *rest.Request {
	return s.RESTClient.Post().Namespace(s.Namespace)
}

// Put todo
func (s *NamespaceRESTClient) Put() *rest.Request {
	return s.RESTClient.Put().Namespace(s.Namespace)
}

// Patch todo
func (s *NamespaceRESTClient) Patch(pt types.PatchType) *rest.Request {
	return s.RESTClient.Patch(pt).Namespace(s.Namespace)
}

// Get todo
func (s *NamespaceRESTClient) Get() *rest.Request {
	return s.RESTClient.Get().Namespace(s.Namespace)
}

// Delete todo
func (s *NamespaceRESTClient) Delete() *rest.Request {
	return s.RESTClient.Delete().Namespace(s.Namespace)
}
