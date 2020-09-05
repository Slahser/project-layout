package main

import (
	"fmt"
	"strconv"
	"time"

	nucliowrapper "github.com/Slahser/coup-de-grace/internal/pkg/nuclio_wrapper"

	"github.com/Slahser/coup-de-grace/internal/app/helper"
	traefikwrapper "github.com/Slahser/coup-de-grace/internal/pkg/traefik_wrapper"
	"github.com/Slahser/coup-de-grace/internal/pkg/traefik_wrapper/traefiksvcfactory"
	errors "github.com/go-errors/errors"
	zap "go.uber.org/zap"
	k8smetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ()

func main() {

	//全局logger配置
	logger := helper.InitLogger()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	//ttTraefik()
	ttNuclio()

}

func ttNuclio() {

	funcList, _ := nucliowrapper.Functions().List(k8smetav1.ListOptions{})
	for i, mw := range funcList.Items {
		fmt.Println("func" + strconv.Itoa(i) + "=>" + mw.Name)
	}

	funcEventsList, _ := nucliowrapper.FunctionEvents().List(k8smetav1.ListOptions{})
	for i, mw := range funcEventsList.Items {
		fmt.Println("funcEvent" + strconv.Itoa(i) + "=>" + mw.Name)
	}

	projectsList, _ := nucliowrapper.Projects().List(k8smetav1.ListOptions{})
	for i, mw := range projectsList.Items {
		fmt.Println("project" + strconv.Itoa(i) + "=>" + mw.Name)
	}

}

func ttTraefik() {

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
	selfHostedSvc, _ := traefiksvcfactory.Create(traefiksvcfactory.SELF_HOSTED, make(map[string]interface{}))

	//gen svc cloud hosted
	cloudHostedSvc, _ := traefiksvcfactory.Create(traefiksvcfactory.CLOUD_HOSTED, make(map[string]interface{}))

	//gen svc aggr
	aggrSvc, _ := traefiksvcfactory.Create(traefiksvcfactory.AGGR, make(map[string]interface{}))

	print(selfHostedSvc)
	print(cloudHostedSvc)
	print(aggrSvc)

	execedSvc, createErr := traefikwrapper.TraefikServices().Create(cloudHostedSvc)
	if createErr != nil {
		zap.S().Error(errors.Errorf("create err %w", createErr))
	} else {
		zap.S().Info("created svc " + execedSvc.Name)

		time.Sleep(5 * 1e8)

		if deletedErr := traefikwrapper.TraefikServices().Delete(execedSvc.Name, &k8smetav1.DeleteOptions{
			TypeMeta: k8smetav1.TypeMeta{
				Kind:       "TraefikService",
				APIVersion: "traefik.containo.us/v1alpha1",
			},
		}); deletedErr != nil {
			zap.S().Error(errors.Errorf("delete err %w", deletedErr))
		} else {
			zap.S().Info("deleted svc " + execedSvc.Name)
		}
	}
}
