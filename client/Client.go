package client

import (
	"os"

	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Clientset todo
type Clientset interface {
	GetConfig() *rest.Config
	GetClient() (*kubernetes.Clientset, error)
	GetExtClient() (*apiextensionsclient.Clientset, error)
	GetRestClient(*schema.GroupVersion, bool) (*rest.RESTClient, error)
}

// ClientsetImpl todo
type ClientsetImpl struct {
	Clientset
	Config *rest.Config
}

// GetConfig todo
func (s *ClientsetImpl) GetConfig() *rest.Config {
	return s.Config
}

// GetClient todo
func (s *ClientsetImpl) GetClient() (*kubernetes.Clientset, error) {
	clientset, err := kubernetes.NewForConfig(s.Config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

// GetExtClient todo
func (s *ClientsetImpl) GetExtClient() (*apiextensionsclient.Clientset, error) {
	clientset, err := apiextensionsclient.NewForConfig(s.Config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

// GetRestClient todo
func (s *ClientsetImpl) GetRestClient(groupVersion *schema.GroupVersion, unversion bool) (*rest.RESTClient, error) {
	cfg := s.GetConfig()
	cfg.ContentConfig.GroupVersion = groupVersion
	cfg.APIPath = "/apis"
	cfg.NegotiatedSerializer = scheme.Codecs
	cfg.UserAgent = rest.DefaultKubernetesUserAgent()
	if unversion {
		return rest.UnversionedRESTClientFor(cfg)
	}
	return rest.RESTClientFor(cfg)
}

// GetK8sClient todo
func GetK8sClient() (Clientset, error) {
	var config *rest.Config
	var err error

	// get the config, either from kubeconfig or using our
	// in-cluster service account
	kubeConfig := os.Getenv("KUBECONFIG")
	if len(kubeConfig) != 0 {
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return nil, err
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	return &ClientsetImpl{
		Config: config,
	}, nil
}
