package nucliowrapper

import (
	"sync"

	k8swrapper "github.com/Slahser/coup-de-grace/internal/pkg/k8s_wrapper"
	nucliov1beta1 "github.com/nuclio/nuclio/pkg/platform/kube/client/clientset/versioned/typed/nuclio.io/v1beta1"

	//nucliov1beta1 "github.com/nuclio/nuclio/pkg/platform/kube/client/clientset/versioned/typed/nuclio.io/v1beta1"
	"go.uber.org/zap"
)

var (
	once         sync.Once
	NuclioClient *nucliov1beta1.NuclioV1beta1Client

	NuclioPlayground = "nuclio-playground"
)

func init() {
	InitNuclioClient()
}

func InitNuclioClient() *nucliov1beta1.NuclioV1beta1Client {
	NuclioClient = nucliov1beta1.NewForConfigOrDie(k8swrapper.K8sConfig)

	zap.S().Info("nuclio client init error")

	return NuclioClient
}

func Functions() nucliov1beta1.NuclioFunctionInterface {
	return NuclioClient.NuclioFunctions(NuclioPlayground)
}

func Projects() nucliov1beta1.NuclioProjectInterface {
	return NuclioClient.NuclioProjects(NuclioPlayground)
}

func FunctionEvents() nucliov1beta1.NuclioFunctionEventInterface {
	return NuclioClient.NuclioFunctionEvents(NuclioPlayground)
}

func GetNuclioClient() *nucliov1beta1.NuclioV1beta1Client {
	once.Do(func() {
		NuclioClient = InitNuclioClient()
	})
	return NuclioClient

}