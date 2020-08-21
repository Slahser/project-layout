package traefik_wrapper

import (
	"sync"

	"k8s.io/client-go/rest"

	traefik "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/generated/clientset/versioned/typed/traefik/v1alpha1"
	"go.uber.org/zap"
)

var (
	once          sync.Once
	TraefikClient *traefik.TraefikV1alpha1Client

	traefikPlayground = "traefik-playground"
)

func InitTraefikClient() *traefik.TraefikV1alpha1Client {

	var k8sConfig rest.Config

	TraefikClient, err := traefik.NewForConfig(&k8sConfig)

	if err != nil {
		zap.S().Fatalf("traefik client init error")
	}

	zap.S().Info("traefik client init error")

	return TraefikClient

}

func Middlewares() traefik.MiddlewareInterface {
	return TraefikClient.Middlewares(traefikPlayground)
}

func IngressRoutes() traefik.IngressRouteInterface {
	return TraefikClient.IngressRoutes(traefikPlayground)
}

func TraefikServices() traefik.TraefikServiceInterface {
	return TraefikClient.TraefikServices(traefikPlayground)
}

func GetTraefikClient() *traefik.TraefikV1alpha1Client {

	once.Do(func() {
		TraefikClient = InitTraefikClient()
	})
	return TraefikClient

}
