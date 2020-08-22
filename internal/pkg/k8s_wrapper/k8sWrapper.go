package k8swrapper

import (
	"flag"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	K8sClient *kubernetes.Clientset
	K8sConfig *rest.Config
)

func init(){
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

func GetK8sConfig() *rest.Config {
	var kubeconfig *string
	if home, _ := homedir.Dir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()

	k8sConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		print("k8s client init error")
		panic(err.Error())
	}
	return k8sConfig
}
