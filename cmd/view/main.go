package main

import (
	"context"
	"log"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/BedivereZero/study-crd/pkg/kubernetes"
	"k8s.io/client-go/rest"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		ingressClasses, err := clientset.NetworkingV1beta1().IngressClasses().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Print(err)
		}
		log.Printf("There are %d ingressClasses in the cluster", len(ingressClasses.Items))

		// Examples for error handling:
		// - Use helper functions e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		// _, err = clientset.CoreV1().Pods("default").Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		_, err = clientset.NetworkingV1beta1().IngressClasses().Get(context.TODO(), "example-xxxxx", metav1.GetOptions{})
		if errors.IsNotFound(err) {
			log.Print("IngressClass example-xxxxx not found in default namespace")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			log.Printf("Error getting ingressClass %v", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			log.Print("Found example-xxxxx ingressClass")
		}

		time.Sleep(10 * time.Second)
	}
}
