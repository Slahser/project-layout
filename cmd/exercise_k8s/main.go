package main

import (
	"fmt"
	"github.com/Slahser/coup-de-grace/internal/app/helper"
	traefikwrapper "github.com/Slahser/coup-de-grace/internal/pkg/traefik_wrapper"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

func main() {

	//全局logger配置
	logger := helper.InitLogger()
	undo := zap.ReplaceGlobals(logger)
	defer undo()


	middlewareList, _ := traefikwrapper.Middlewares().List(metav1.ListOptions{})

	for i, mw := range middlewareList.Items {
		fmt.Println("mw"+strconv.Itoa(i)+"=>" +mw.Name)
	}

	routeList, _ := traefikwrapper.IngressRoutes().List(metav1.ListOptions{})

	for i, mw := range routeList.Items {
		fmt.Println("route"+strconv.Itoa(i)+"=>" +mw.Name)
	}

	svcList, _ := traefikwrapper.TraefikServices().List(metav1.ListOptions{})

	for i, mw := range svcList.Items {
		fmt.Println("svc"+strconv.Itoa(i)+"=>" +mw.Name)
	}

}
