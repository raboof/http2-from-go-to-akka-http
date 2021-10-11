package main

import (
	"context"
	//"flag"
	"fmt"
	//"/path/filepath"
	//"time"

	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
	//"k8s.io/client-go/util/homedir"
)

func main() {
	kubeConfig := rest.Config{
		Host: "https://localhost:8001",
		BearerToken: "asdf",
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}

	clientset, err := kubernetes.NewForConfig(&kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// Examples for error handling:
		// - Use helper functions like e.g. errors.IsNotFound()
		// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
		//namespace := "default"
		//pod := "example-xxxxx"
		//_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
		//if errors.IsNotFound(err) {
		//	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
		//} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		//	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
		//		pod, namespace, statusError.ErrStatus.Message)
		//} else if err != nil {
		//	panic(err.Error())
		//} else {
		//	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
		//}

		//time.Sleep(10 * time.Second)
	}
}
