package controller

import (
	"context"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Controller struct{}

func NewController(*kubernetes.Clientset) *Controller {
	controller := &Controller{}
	return controller
}

func Execute() error {

	kubeConfig, err := getKubeConfigPath()
	if err != nil {
		return err
	}

	cfg, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return err
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return err
	}

	nss, err := kubeClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	for _, ns := range nss.Items {
		logrus.Info(ns.Name)
	}

	_ = NewController(kubeClient)

	logrus.Info("DONE")

	return nil
}
