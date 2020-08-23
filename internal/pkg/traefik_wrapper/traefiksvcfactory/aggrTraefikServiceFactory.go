package traefiksvcfactory

import (
	traefikwrapper "github.com/Slahser/coup-de-grace/internal/pkg/traefik_wrapper"
	traefikv1alpha1 "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	Register("aggr", &aggrTraefikServiceFactory{})
}

// testDriverFactory implements the traefiksvcfactory.StorageDriverFactory interface.
type aggrTraefikServiceFactory struct{}

func (factory *aggrTraefikServiceFactory) Create(parameters map[string]interface{}) (*traefikv1alpha1.TraefikService, error) {
	return NewAggrFromParameters(parameters)
}

func NewAggrFromParameters(parameters map[string]interface{}) (*traefikv1alpha1.TraefikService, error) {
	return &traefikv1alpha1.TraefikService{
		TypeMeta: k8smetav1.TypeMeta{
			Kind:       "TraefikService",
			APIVersion: "traefik.containo.us/v1alpha1",
		},
		ObjectMeta: k8smetav1.ObjectMeta{
			Name:      "aggr-svc",
			Namespace: traefikwrapper.TraefikPlayground,
			Labels: map[string]string{
				"type":    "aggr",
				"org":     "",
				"project": "",
				"env":     "",
			},
			Annotations: nil,
		},
		Spec: traefikv1alpha1.ServiceSpec{
			Weighted: &traefikv1alpha1.WeightedRoundRobin{
				Services: []traefikv1alpha1.Service{
					{
						LoadBalancerSpec: traefikv1alpha1.LoadBalancerSpec{
							Name:      "self-hosted-svc",
							Kind:      "TraefikService",
							Namespace: traefikwrapper.TraefikPlayground,
							Weight:    nil,
						},
					},
					{
						LoadBalancerSpec: traefikv1alpha1.LoadBalancerSpec{
							Name:      "cloud-hosted-svc",
							Kind:      "TraefikService",
							Namespace: traefikwrapper.TraefikPlayground,
							Weight:    nil,
						},
					},
				},
			},
		},
	}, nil
}
