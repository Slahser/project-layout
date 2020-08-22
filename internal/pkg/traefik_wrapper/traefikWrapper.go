package traefikwrapper

import (
	"sync"

	k8swrapper "github.com/Slahser/coup-de-grace/internal/pkg/k8s_wrapper"
	"go.uber.org/zap"

	//clientset "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/generated/clientset/versioned"
	traefikv1alpha1 "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/generated/clientset/versioned/typed/traefik/v1alpha1"
	//faketraefikv1alpha1 "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/generated/clientset/versioned/typed/traefik/v1alpha1/fake"
)

var (
	once          sync.Once
	TraefikClient *traefikv1alpha1.TraefikV1alpha1Client

	traefikPlayground = "traefikv1alpha1-playground"
)

func init() {
	InitTraefikClient()
}

func InitTraefikClient() *traefikv1alpha1.TraefikV1alpha1Client {

	TraefikClient := traefikv1alpha1.NewForConfigOrDie(k8swrapper.K8sConfig)

	zap.S().Info("traefikv1alpha1 client init error")

	return TraefikClient

}

func Middlewares() traefikv1alpha1.MiddlewareInterface {
	return TraefikClient.Middlewares(traefikPlayground)
}

func IngressRoutes() traefikv1alpha1.IngressRouteInterface {
	return TraefikClient.IngressRoutes(traefikPlayground)
}

func TraefikServices() traefikv1alpha1.TraefikServiceInterface {
	return TraefikClient.TraefikServices(traefikPlayground)
}

func GetTraefikClient() *traefikv1alpha1.TraefikV1alpha1Client {

	once.Do(func() {
		TraefikClient = InitTraefikClient()
	})
	return TraefikClient

}
