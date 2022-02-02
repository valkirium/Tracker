package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
//	"time"
//	appsv1 "k8s.io/api/apps/v1"
//	corev1 "k8s.io/api/core/v1"
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
		log.Fatalln("failed to get deployments:", err)
        }
	for _, Deployment := range Deployments.Items {
		fmt.Printf("%s\n", Deployment.Name)
		fmt.Print("Ready: ")
		fmt.Printf("%d\n", Deployment.Status.ReadyReplicas)
		fmt.Print("Up to date: ") 
		fmt.Printf("%d\n", Deployment.Status.UpdatedReplicas) 
		if Deployment.Status.ReadyReplicas == Deployment.Status.UpdatedReplicas {
		fmt.Println("Ok")
		}
	}
	DaemonSets, err := clientset.AppsV1().DaemonSets("kube-system").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
                log.Fatalln("failed to get daemonset:", err)
        }
	for _, DaemonSet := range DaemonSets.Items {
                fmt.Printf("%s\n", DaemonSet.Name)
		fmt.Print("Ready: ")
		fmt.Printf("%d\n", DaemonSet.Status.NumberReady)
		fmt.Print("Up to date: ")
		fmt.Printf("%d\n", DaemonSet.Status.UpdatedNumberScheduled)
		if  DaemonSet.Status.NumberReady == DaemonSet.Status.UpdatedNumberScheduled {
		fmt.Println("Ok")
		}
	}

}

