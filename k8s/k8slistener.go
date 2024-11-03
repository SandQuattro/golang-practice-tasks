package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// Load kubeconfig
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// Define the namespace you want to watch
	var namespace string
	flag.StringVar(&namespace, "namespace", "dbs-antifraud", "Kubernetes namespace to watch")

	flag.Parse()

	// Build the configuration from kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	// Create an informer factory for the specified namespace
	factory := informers.NewSharedInformerFactoryWithOptions(
		clientset,
		0,
		informers.WithNamespace(namespace),
	)

	// Create an event informer for the namespace
	eventInformer := factory.Core().V1().Events().Informer()

	// Set up event handlers
	eventInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			event := obj.(*v1.Event)
			fmt.Printf("[Add] Time: %s, Involved Object: %s, Message: %s\n", event.LastTimestamp, event.InvolvedObject.Name, event.Message)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			event := newObj.(*v1.Event)
			fmt.Printf("[Update] Time: %s, Involved Object: %s, Message: %s\n", event.LastTimestamp, event.InvolvedObject.Name, event.Message)
		},
		DeleteFunc: func(obj interface{}) {
			event := obj.(*v1.Event)
			fmt.Printf("[Delete] Time: %s, Involved Object: %s, Message: %s\n", event.LastTimestamp, event.InvolvedObject.Name, event.Message)
		},
	})

	stopCh := make(chan struct{})
	defer close(stopCh)

	// Start the informer
	factory.Start(stopCh)

	// Wait for the caches to be synced before processing items
	synced := cache.WaitForCacheSync(stopCh, eventInformer.HasSynced)
	if !synced {
		log.Fatalf("Error syncing caches")
	}

	fmt.Printf("Listening to events in namespace: %s\n", namespace)

	// Run until the context is cancelled
	<-stopCh
}
