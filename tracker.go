package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// This program lists the pods in a cluster equivalent to
//
// kubectl get pods
//
func main() {
	var ns string
	flag.StringVar(&ns, "namespace", "", "namespace")

	// Bootstrap k8s configuration from local       Kubernetes config file
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	log.Println("Using kubeconfig file: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create an rest client not targeting specific API version
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

        nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get nodes:", err)
	}
	for _, node := range nodes.Items {
		fmt.Printf("%s\n", node.Name)
		for _, condition := range node.Status.Conditions {
			fmt.Printf("\t%s: %s\n", condition.Type, condition.Status)
		}
	}
	pods, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
                log.Fatalln("failed to get pods:", err)
        }
        for _, pod := range pods.Items {
                fmt.Printf("%s\n", pod.Name)
                for _, condition := range pod.Status.Conditions {
                        fmt.Printf("\t%s: %s\n", condition.Type, condition.Status)
                }
	}
	Deployments, err := clientset.AppsV1().Deployments("kube-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get dps:", err)
        }
	for _, Deployment := range Deployments.Items {
		fmt.Printf("%s\n", Deployment.Name)
                for _, condition := range Deployment.Status.Conditions {
                        fmt.Printf("\t%s: %s\n", condition.Type, condition.Status)
                }

	}


//	pods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{})
//	if err != nil {
//		log.Fatalln("failed to get pods:", err)
//	}
//	// print pods
//	for i, pod := range pods.Items {
//		fmt.Printf("[%d] %s\n", i, pod.GetName())
//	}
}
