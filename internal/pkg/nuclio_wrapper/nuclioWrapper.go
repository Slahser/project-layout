package nucliowrapper

import (
	"os"
	"sync"

	"github.com/nuclio/nuclio/pkg/containerimagebuilderpusher"
	"github.com/nuclio/nuclio/pkg/platform"
	"github.com/nuclio/nuclio/pkg/platform/config"
	"github.com/nuclio/nuclio/pkg/platform/factory"
	nucliozap "github.com/nuclio/zap"

	k8swrapper "github.com/Slahser/coup-de-grace/internal/pkg/k8s_wrapper"
	//nucliov1beta1 "github.com/Slahser/coup-de-grace/third_party/nuclio/pkg/platform/kube/client/clientset/versioned/typed/nuclio.io/v1beta1"

	nucliov1beta1 "github.com/nuclio/nuclio/pkg/platform/kube/client/clientset/versioned/typed/nuclio.io/v1beta1"
	"go.uber.org/zap"
)

var (
	once             sync.Once
	NuclioClient     *nucliov1beta1.NuclioV1beta1Client
	NuclioPlatform   platform.Platform
	NuclioPlayground = "nuclio-playground"
)

func init() {
	InitNuclioClient()
	InitNuclioPlatform()
}

func InitNuclioClient() *nucliov1beta1.NuclioV1beta1Client {
	NuclioClient = nucliov1beta1.NewForConfigOrDie(k8swrapper.K8sConfig)

	zap.S().Info("nuclio client init succeed")

	return NuclioClient
}

func InitNuclioPlatform() platform.Platform {

	nuclioZapLogger, _ := nucliozap.NewNuclioZap("name", "console", nil, os.Stdout, os.Stdout, nucliozap.DebugLevel)

	NuclioPlatform, err := factory.CreatePlatform(nuclioZapLogger,
		"kube",
		&config.Configuration{
			KubeconfigPath: k8swrapper.GetK8sConfigPath(),
			ContainerBuilderConfiguration: containerimagebuilderpusher.ContainerBuilderConfiguration{
				Kind:                                 "docker",
				BusyBoxImage:                         "",
				KanikoImage:                          "",
				KanikoImagePullPolicy:                "",
				JobPrefix:                            "",
				DefaultRegistryCredentialsSecretName: "",
				DefaultBaseRegistryURL:               "",
				DefaultOnbuildRegistryURL:            "",
				CacheRepo:                            "",
				InsecurePushRegistry:                 false,
				InsecurePullRegistry:                 false,
			},
		},
		NuclioPlayground)

	if err != nil {
		zap.S().Fatalf("init nuclio platform error: %v", err)
	}

	return NuclioPlatform
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

func Tt() {
	NuclioPlatform.CreateFunction(nil)
}
