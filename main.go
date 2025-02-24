package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	kclient "github.com/shreekara-rajendra/KindToDigitalOcean/pkg/client/clientset/versioned"
	custominformer "github.com/shreekara-rajendra/KindToDigitalOcean/pkg/client/informers/externalversions"
	"github.com/shreekara-rajendra/KindToDigitalOcean/pkg/controller"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	kubeconfig := flag.String("kubeconfig", "/home/shreekara-rajendra/.kube/config", "configurations for client")
	flag.Parse()
	restConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Println("config couldnt be located")
		//in cluster config
		restConfig, err = rest.InClusterConfig()
		if err != nil {
			log.Fatal(err)
		}
	}
	clientset, err := kclient.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("error %s", err.Error())
	}
	fmt.Println(clientset)
	dclist, err := clientset.ShreekararajendraV1alpha1().DigitalClusters("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("error %s", err.Error())
	}
	fmt.Printf("Length of resourcelist: %d", len(dclist.Items))
	for _, item := range dclist.Items {
		fmt.Println(item)
	}
	ch := make(chan struct{})
	informer := custominformer.NewSharedInformerFactory(clientset, 20*time.Minute)
	digitalinformer := informer.Shreekararajendra().V1alpha1().DigitalClusters()
	digitalController := controller.NewController(clientset, digitalinformer)
	informer.Start(ch)
	digitalController.Run(ch)

}
