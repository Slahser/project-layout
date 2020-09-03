package k8swrapper

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	K8sClient *kubernetes.Clientset
	K8sConfig *rest.Config
)

func init() {
	InitK8sClient()
}

func InitK8sClient() *kubernetes.Clientset {
	K8sConfig = GetK8sConfig()
	K8sClient, err := kubernetes.NewForConfig(K8sConfig)
	if err != nil {
		panic(err.Error())
	}

	return K8sClient

}

//GetK8sConfig
func GetK8sConfig() *rest.Config {
	kubeconfig := GetK8sConfigPath()

	k8sConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		zap.S().Errorf("k8s client init error %w", err)
		panic(err.Error())
	}
	return k8sConfig
}

//GetK8sConfigPath
func GetK8sConfigPath() string {

	if home, err := homedir.Dir(); err != nil {
		zap.S().Errorf("k8s client init error %w", err)
		panic(err.Error())
	} else {
		return filepath.Join(home, ".kube", "config")
	}

}
