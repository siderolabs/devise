package discoverers

import (
	"log"
	"time"

	"google.golang.org/grpc/grpclog"

	"github.com/autonomy/devise/devise/renderer"
	"github.com/autonomy/devise/devise/storage/datastore"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/fields"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

// Kubernetes implements the Discoverer interface.
type Kubernetes struct {
	client    *kubernetes.Clientset
	datastore datastore.Datastore
	renderer  *renderer.Renderer
}

// NewKubernetesDiscoverer instantiates and returns and Kubernetes discoverer.
func NewKubernetesDiscoverer() (k Kubernetes) {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Printf("Failed to start Kubernetes discoverer: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("Failed to start Kubernetes discoverer: %v", err)
	}
	k = Kubernetes{
		client: clientset,
	}

	return
}

// Discover implements discoverers.Discoverer
func (k Kubernetes) Discover(d datastore.Datastore, r *renderer.Renderer) {
	k.datastore = d
	k.renderer = r
	getter := k.client.Core().RESTClient()
	watchlist := cache.NewListWatchFromClient(getter, "pods", v1.NamespaceAll, fields.Everything())
	_, controller := cache.NewInformer(
		watchlist,
		&v1.Pod{},
		time.Second*0,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    k.addFunc,
			DeleteFunc: k.delteFunc,
			UpdateFunc: k.updateFunc,
		},
	)

	stop := make(chan struct{})
	go controller.Run(stop)
}

func (k Kubernetes) addFunc(obj interface{}) {
	pod := obj.(*v1.Pod)
	if _, ok := pod.Annotations["autonomy.io/devise"]; !ok {
		log.Printf("Skipping pod %s", pod.Name)
		return
	}

	if pod.Status.PodIP == "" {
		grpclog.Printf("Discover(_) = _, No IP found")
		return
	}
	err := k.datastore.Put(&datastore.Entry{Key: pod.GetName(), Value: []byte(pod.Status.PodIP)})
	if err != nil {
		log.Printf("Failed to insert entry: %v", err)
	}
	k.renderer.Render(pod.GetName(), pod.Status.PodIP, k.datastore)
}

func (k Kubernetes) delteFunc(obj interface{}) {
	pod := obj.(*v1.Pod)
	err := k.datastore.Delete(pod.GetName())
	if err != nil {
		log.Printf("Failed to delete entry: %v", err)
	}
}

func (k Kubernetes) updateFunc(oldObj, newObj interface{}) {
	oldPod := oldObj.(*v1.Pod)
	newPod := newObj.(*v1.Pod)
	if oldPod.Status.PodIP == "" && newPod.Status.PodIP != "" {
		k.renderer.Render(newPod.GetName(), newPod.Status.PodIP, k.datastore)
	}
}
