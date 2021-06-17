package main

import (
	"context"
	"flag"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	clientset "github.com/BedivereZero/study-crd/pkg/client/clientset/versioned"
)

var (
	masterURL  string
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func main() {
	klog.Info("Hello, world!")
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalln("build kubeconfig fail:", err)
	}

	clientSet, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatal("new client set fail:", err)
	}

	for {
		listFoos(clientSet)
		getFoo(clientSet, "test", "example-foo")
		time.Sleep(10 * time.Second)
	}
}

func listFoos(cs *clientset.Clientset) {
	fooList, err := cs.FooV1alpha1().Foos("test").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		klog.Warningln("list foos fail:", err)
	}
	klog.Infof("There are %d Foos in test namespace", len(fooList.Items))
}

func getFoo(cs *clientset.Clientset, ns, name string) {
	foo, err := cs.FooV1alpha1().Foos(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		klog.Infof("Foo %s not found in namespace %s", name, ns)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		klog.Warningf("Error getting ingressClass %v", statusError.ErrStatus.Message)
	} else if err != nil {
		klog.Warningln("Unknown error:", err)
		panic(err.Error())
	} else {
		klog.Infof("Found %s/%s at namespace %s", foo.Kind, foo.Name, foo.Namespace)
	}
}
