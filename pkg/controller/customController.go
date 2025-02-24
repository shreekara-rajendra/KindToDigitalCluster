package controller

import (
	"log"
	"time"

	customclient "github.com/shreekara-rajendra/KindToDigitalOcean/pkg/client/clientset/versioned"
	custominformer "github.com/shreekara-rajendra/KindToDigitalOcean/pkg/client/informers/externalversions/shreekararajendra.dev/v1alpha1"
	customlister "github.com/shreekara-rajendra/KindToDigitalOcean/pkg/client/listers/shreekararajendra.dev/v1alpha1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	//clientset for accessing resource of customtype
	clientset customclient.Interface
	//cache has synced or not
	cachesynced cache.InformerSynced
	//lister interface for accessing resource of custom type
	clister customlister.DigitalClusterLister
	//queue
	queue workqueue.TypedRateLimitingInterface[interface{}]
}

func NewController(clientset customclient.Interface, dcinformer custominformer.DigitalClusterInformer) *Controller {
	c := &Controller{
		clientset:   clientset,
		clister:     dcinformer.Lister(),
		cachesynced: dcinformer.Informer().HasSynced,
		queue:       workqueue.NewTypedRateLimitingQueue(workqueue.DefaultTypedControllerRateLimiter[interface{}]()),
	}
	dcinformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.handleadd,
		DeleteFunc: c.handledelete,
	})
	return c
}

func (c *Controller) worker() {
	for c.processItem() {

	}
}

func (c *Controller) processItem() bool {

	item, shutdown := c.queue.Get()
	if shutdown {
		return false
	}
	defer c.queue.Forget(item)
	key, err := cache.MetaNamespaceKeyFunc(item)
	if err != nil {
		log.Printf("getting key from cache %s\n", err.Error())
		return false
	}
	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		log.Printf("splitting key into ns,name %s\n", err.Error())
		return false
	}
	log.Printf("name: %s; namespace: %s\n", name, ns)
	dgresource, err := c.clister.DigitalClusters(ns).Get(name)
	if err != nil {
		log.Printf("error %s getting resource from lister\n", err.Error())
		return false
	}
	log.Printf("custom resource %+v\n", dgresource)
	return true

}
func (c *Controller) Run(ch <-chan struct{}) {
	log.Println("starting controller")
	if !cache.WaitForCacheSync(ch, c.cachesynced) {
		log.Fatalln("waiting for cache to be synced")
	}
	go wait.Until(c.worker, 1*time.Second, ch)
	<-ch
}

func (c *Controller) handleadd(item interface{}) {
	log.Println("add was called")
	c.queue.Add(item)
}

func (c *Controller) handledelete(item interface{}) {
	log.Println("delete was called")
	c.queue.Add(item)
}
