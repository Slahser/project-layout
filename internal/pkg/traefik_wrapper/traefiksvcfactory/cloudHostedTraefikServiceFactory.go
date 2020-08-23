package traefiksvcfactory

import (
	"fmt"
	traefikwrapper "github.com/Slahser/coup-de-grace/internal/pkg/traefik_wrapper"
	traefikv1alpha1 "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func init() {
	Register("cloudhosted", &cloudHostedTraefikServiceFactory{})
}

// testDriverFactory implements the traefiksvcfactory.StorageDriverFactory interface.
type cloudHostedTraefikServiceFactory struct{}

func (factory *cloudHostedTraefikServiceFactory) Create(parameters map[string]interface{}) (*traefikv1alpha1.TraefikService, error) {

	return NewCloudHostedFromParameters(parameters)
}

func NewCloudHostedFromParameters(parameters map[string]interface{}) (*traefikv1alpha1.TraefikService, error) {

	tt := parameters["yy"]
	if tt == nil || fmt.Sprint(tt) == "" {
		return nil, fmt.Errorf("no region parameter provided")
	}
	ttStr := fmt.Sprint(tt)
	print(ttStr)

	return &traefikv1alpha1.TraefikService{
		TypeMeta: k8smetav1.TypeMeta{
			Kind:       "TraefikService",
			APIVersion: "traefik.containo.us/v1alpha1",
		},
		ObjectMeta: k8smetav1.ObjectMeta{
			Name:      "gen-cloud-hosted-svc",
			Namespace: traefikwrapper.TraefikPlayground,
			Labels: map[string]string{
				"type":    "cloud-hosted",
				"org":     "ttorg",
				"project": "ttpro",
				"env":     "ttenv",
			},
			Annotations: nil,
		},
	}, nil
}
