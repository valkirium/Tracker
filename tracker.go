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
//	"golang.org/x/build/kubernetes/api"
	"k8s.io/client-go/tools/clientcmd"
)

// This program lists the pods in a cluster equivalent toда
//
// kubectl get pods
//
func main() {
	var ns string
	var i int = 0
	var a int = 0
	var b int = 0
	var c int = 0
	var d int = 0
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
	for index, node := range nodes.Items {
		fmt.Printf("%d\t %s\n", index, node.Name)
		for _, condition := range node.Status.Conditions {
			fmt.Printf("\t%s: %s\n", condition.Type, condition.Status)
			if condition.Type == "NetworkUnavailable" && condition.Status == "False" {
		fmt.Println("Ok")
		i++
			}
			if condition.Type == "MemoryPressure" && condition.Status == "False" {
                fmt.Println("Ok")
		i++
                        }
			if condition.Type == "DiskPressure" && condition.Status == "False" {
                fmt.Println("Ok")
		i++
                        }
			if condition.Type == "PIDPressure" && condition.Status == "False" {
                fmt.Println("Ok")
		i++
                        }
			if condition.Type == "Ready" && condition.Status == "True" {
                fmt.Println("Ok")
		i++
                        }
		a = (index + 1) * 5
		}
	}

	pods, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
                log.Fatalln("failed to get pods:", err)
        }
        for index, pod := range pods.Items {
                fmt.Printf("%d\t %s\n", index,  pod.Name)
                for _, condition := range pod.Status.Conditions {
                        fmt.Printf("\t%s: %s\n", condition.Type, condition.Status)
			if condition.Type == "Initialized" && condition.Status == "True" {
                fmt.Println("Ok")
                i++
                        }
			if condition.Type == "Ready" && condition.Status == "True" {
                fmt.Println("Ok")
                i++
                        }
			if condition.Type == "ContainersReady" && condition.Status == "True" {
                fmt.Println("Ok")
                i++
                        }
			 if condition.Type == "PodScheduled" && condition.Status == "True" {
                fmt.Println("Ok")
                i++
                        }
		b = (index + 1) * 4
                }
	}
	Deployments, err := clientset.AppsV1().Deployments("kube-system").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get deployments:", err)
        }
	for index, Deployment := range Deployments.Items {
		fmt.Printf("%d\t %s\n", index, Deployment.Name)
		fmt.Print("Ready: ")
		fmt.Printf("%d\n", Deployment.Status.ReadyReplicas)
		fmt.Print("Up to date: ")
		fmt.Printf("%d\n", Deployment.Status.UpdatedReplicas)
		if Deployment.Status.ReadyReplicas == Deployment.Status.UpdatedReplicas {
		fmt.Println("Ok")
		i++
		}
		c = index + 1
	}
	DaemonSets, err := clientset.AppsV1().DaemonSets("kube-system").List(context.TODO(), metav1.ListOptions{})
        if err != nil {
                log.Fatalln("failed to get daemonset:", err)
        }
	for index, DaemonSet := range DaemonSets.Items {
                fmt.Printf("%d\t %s\n", index, DaemonSet.Name)
		fmt.Print("Ready: ")
		fmt.Printf("%d\n", DaemonSet.Status.NumberReady)
		fmt.Print("Up to date: ")
		fmt.Printf("%d\n", DaemonSet.Status.UpdatedNumberScheduled)
		if  DaemonSet.Status.NumberReady == DaemonSet.Status.UpdatedNumberScheduled {
		fmt.Println("Ok")
		i++
		}
		d = index + 1
	}
        if i == (a + b + c + d) {
		fmt.Println("OK")
	}else{fmt.Println("Cluster is failed")
}
}

