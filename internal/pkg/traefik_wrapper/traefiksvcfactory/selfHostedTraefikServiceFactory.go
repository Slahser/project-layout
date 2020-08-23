package traefiksvcfactory

import (
	nucliowrapper "github.com/Slahser/coup-de-grace/internal/pkg/nuclio_wrapper"
	traefikwrapper "github.com/Slahser/coup-de-grace/internal/pkg/traefik_wrapper"
	traefikv1alpha1 "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const svcType = "selfhosted"

func init() {
	Register(svcType, &selfHostedTraefikServiceFactory{})
}

// testDriverFactory implements the traefiksvcfactory.StorageDriverFactory interface.
type selfHostedTraefikServiceFactory struct{}

func (factory *selfHostedTraefikServiceFactory) Create(parameters map[string]interface{}) (*traefikv1alpha1.TraefikService, error) {
	return NewSelfHostedFromParameters(parameters)
}

func NewSelfHostedFromParameters(parameters map[string]interface{}) (*traefikv1alpha1.TraefikService, error) {
	return &traefikv1alpha1.TraefikService{
		TypeMeta: k8smetav1.TypeMeta{
			Kind:       "TraefikService",
			APIVersion: "traefik.containo.us/v1alpha1",
		},
		ObjectMeta: k8smetav1.ObjectMeta{
			Name:      "gen-self-hosted-svc",
			Namespace: traefikwrapper.TraefikPlayground,
			Labels: map[string]string{
				"type":    "self-hosted",
				"org":     "ttorg",
				"project": "ttpro",
				"env":     "ttenv",
			},
			Annotations: nil,
		},
		Spec: traefikv1alpha1.ServiceSpec{
			Weighted: &traefikv1alpha1.WeightedRoundRobin{
				Services: []traefikv1alpha1.Service{
					{
						LoadBalancerSpec: traefikv1alpha1.LoadBalancerSpec{
							Name:      "",
							Kind:      "Service",
							Namespace: nucliowrapper.NuclioPlayground,
							Port:      80,
							Scheme:    "http",
							Weight:    nil,
						},
					},
				},
			},
		},
	}, nil
}
