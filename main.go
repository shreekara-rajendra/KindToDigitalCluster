package main

import (
	"flag"
	"fmt"
	"log"

	kclient "github.com/shreekara-rajendra/KindToDigitalOcean/pkg/client/clientset/versioned"
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
}
