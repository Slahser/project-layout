package main

import (
	"fmt"
	"github.com/Slahser/coup-de-grace/internal/app/helper"
	nucliowrapper "github.com/Slahser/coup-de-grace/internal/pkg/nuclio_wrapper"
	traefikwrapper "github.com/Slahser/coup-de-grace/internal/pkg/traefik_wrapper"
	traefikv1alpha1 "github.com/containous/traefik/v2/pkg/provider/kubernetes/crd/traefik/v1alpha1"
	"github.com/go-errors/errors"
	"go.uber.org/zap"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
	"time"
)

var ()

func main() {

	//全局logger配置
	logger := helper.InitLogger()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	middlewareList, _ := traefikwrapper.Middlewares().List(k8smetav1.ListOptions{})
	for i, mw := range middlewareList.Items {
		fmt.Println("mw" + strconv.Itoa(i) + "=>" + mw.Name)
	}

	routeList, _ := traefikwrapper.IngressRoutes().List(k8smetav1.ListOptions{})
	for i, mw := range routeList.Items {
		fmt.Println("route" + strconv.Itoa(i) + "=>" + mw.Name)
	}

	svcList, _ := traefikwrapper.TraefikServices().List(k8smetav1.ListOptions{})
	for i, mw := range svcList.Items {
		fmt.Println("svc" + strconv.Itoa(i) + "=>" + mw.Name)
	}

	//gen svc self hosted
	selfHostedSvc := &traefikv1alpha1.TraefikService{
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
	}

	//gen svc cloud hosted
	cloudHostedSvc := &traefikv1alpha1.TraefikService{
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
	}

	//gen svc aggr
	aggrSvc := &traefikv1alpha1.TraefikService{
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
	}

	print(selfHostedSvc)
	print(cloudHostedSvc)
	print(aggrSvc)

	execedSvc, createErr := traefikwrapper.TraefikServices().Create(cloudHostedSvc)
	if createErr != nil {
		zap.S().Error(errors.Errorf("create err",createErr))
	} else {
		zap.S().Info("created svc " + execedSvc.Name)

		time.Sleep(5 * 1e8)

		if deletedErr := traefikwrapper.TraefikServices().Delete(execedSvc.Name, &k8smetav1.DeleteOptions{
			TypeMeta: k8smetav1.TypeMeta{
				Kind:       "TraefikService",
				APIVersion: "traefik.containo.us/v1alpha1",
			},
		}); deletedErr != nil {
			zap.S().Error(errors.Errorf("delete err",deletedErr))
		}else {
			zap.S().Info("deleted svc " + execedSvc.Name)
		}
	}

}
