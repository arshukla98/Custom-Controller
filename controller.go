package main

import (
	"context"
	"fmt"
	"time"

	clientset "github.com/arshukla98/sample-controller/generated/clientset/versioned"
	informers "github.com/arshukla98/sample-controller/generated/informers/externalversions"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	wait "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	upgradev1 "github.com/arshukla98/sample-controller/pkg/apis/upgrade/v1"
	controller "github.com/arshukla98/sample-controller/pkg/controller"
)

type Controller struct {
	// pods gives cached access to pods.
	//pods informers.UpgradeKubeLister
	podsSynced cache.InformerSynced
	// queue is where incoming work is placed to de-dup and to allow "easy"
	// rate limited requeues on errors
	queue workqueue.RateLimitingInterface

	informer cache.SharedIndexInformer
}

func NewController(client *clientset.Clientset) *Controller {
	fmt.Println("passing queue")
	queue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "controller-name")

	fmt.Println("passing informers")
	exampleInformerFactory := informers.NewSharedInformerFactory(client, time.Second*5)
	upgradeInformer := exampleInformerFactory.Upgrade().V1().UpgradeKubes()
	sharedIndexUpgradeInfomer := upgradeInformer.Informer()

	fmt.Println("Setting up event handlers")
	sharedIndexUpgradeInfomer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			fmt.Println("Adding key:", key)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old interface{}, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			fmt.Println("Updating key:", key)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			// IndexerInformer uses a delta nodeQueue, therefore for deletes we have to use this
			// key function.

			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			fmt.Println("Deleting key:", key)
			if err == nil {
				queue.Add(key)
			}
		},
	})

	exampleInformerFactory.Start(context.Background().Done())

	c := &Controller{}
	c.podsSynced = sharedIndexUpgradeInfomer.HasSynced
	c.queue = queue
	c.informer = sharedIndexUpgradeInfomer
	return c
}

func (c *Controller) Run(ctx context.Context, threadiness int) error {
	// don't let panics crash the process
	defer utilruntime.HandleCrash()
	// make sure the work queue is shutdown which will trigger workers to end
	defer c.queue.ShutDown()

	fmt.Println("Starting <NAME> controller")

	// wait for your secondary caches to fill before starting your work
	if !cache.WaitForCacheSync(ctx.Done(), c.podsSynced) {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	// start up your worker threads based on threadiness.  Some controllers
	// have multiple kinds of workers
	for i := 0; i < threadiness; i++ {
		// runWorker will loop until "something bad" happens.  The .Until will
		// then rekick the worker after one second
		go wait.UntilWithContext(ctx, c.runWorker, time.Second)
	}

	fmt.Println("Started workers")
	<-ctx.Done()
	fmt.Println("Shutting down workers")

	return nil
}

func (c *Controller) runWorker(ctx context.Context) {
	// hot loop until we're told to stop.  processNextWorkItem will
	// automatically wait until there's work available, so we don't worry
	// about secondary waits
	i := 0
	fmt.Println("Begin runWorker")
	for c.processNextWorkItem(ctx) {
		fmt.Println(time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
		i += 1
		fmt.Println("i:", i)
	}
	fmt.Println("End runWorker")
}

// processNextWorkItem deals with one key off the queue.  It returns false
// when it's time to quit.
func (c *Controller) processNextWorkItem(ctx context.Context) bool {
	fmt.Println("Begin processNextWorkItem")
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	fmt.Println("key:", key.(string))
	// do your work on the key.  This method will contains your "do stuff" logic
	err := c.syncHandler(key.(string))
	if err == nil {
		// No error, reset the ratelimit counters

		// if you had no error, tell the queue to stop tracking history for your
		// key. This will reset things like failure counts for per-item rate
		// limiting
		fmt.Println("Forget key:", key.(string))
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < 5 {
		// err != nil and retry
		c.queue.AddRateLimited(key)
	} else {
		// err != nil and too many retries
		c.queue.Forget(key)
		utilruntime.HandleError(fmt.Errorf("%v failed with : %v", key, err))
	}
	fmt.Println("End processNextWorkItem")
	return true
}

func (c *Controller) syncHandler(key string) error {
	fmt.Println("In syncHandler: ", key)

	obj, exists, err := c.informer.GetIndexer().GetByKey(key)
	if err != nil {
		return fmt.Errorf("Error fetching object with key %s from store: %v", key, err)
	}
	if !exists {
		fmt.Println("Object Deleted")
		return nil
	}

	uk, ok := obj.(*upgradev1.UpgradeKube)
	if !ok || uk == nil {
		return nil
	}

	uk, err = controller.ProcessKubeUpgrade(context.Background(), uk.DeepCopy())
	if err != nil {
		return err
	}
	if uk != nil {
		// Needs to update
		// return err -> this if error in updating.
	}
	return nil
}
