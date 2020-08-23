package traefiksvcfactory

import (
	"fmt"
	traefikv1alpha1 "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
	"github.com/go-errors/errors"
)

var svcFactories = make(map[string]TraefikServiceFactory)

/**
registry/storage/driver/traefiksvcfactory/traefiksvcfactory.go:22
registry/storage/driver/s3-aws/s3.go:171
*/
type TraefikServiceFactory interface {
	Create(parameters map[string]interface{}) (*traefikv1alpha1.TraefikService, error)
}

func Register(name string, factory TraefikServiceFactory) {
	if factory == nil {
		panic("Must not provide nil TraefikServiceFactory")
	}
	_, registered := svcFactories[name]
	if registered {
		panic(fmt.Sprintf("TraefikServiceFactory named %s already registered", name))
	}

	svcFactories[name] = factory
}

func Create(name string, parameters map[string]interface{}) (*traefikv1alpha1.TraefikService, error) {
	svcFactory, ok := svcFactories[name]
	if !ok {
		return nil, errors.Errorf(name)
	}
	return svcFactory.Create(parameters)
}
