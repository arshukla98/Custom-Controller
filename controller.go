package main

import (
	"context"
	"fmt"
	kubeinformers "k8s.io/client-go/informers"
	"reflect"
	"time"

	clientset "github.com/arshukla98/sample-controller/generated/clientset/versioned"
	informers "github.com/arshukla98/sample-controller/generated/informers/externalversions"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	wait "k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	upgradev1 "github.com/arshukla98/sample-controller/pkg/apis/upgrade/v1"
)

type Controller struct {
	// pods gives cached access to pods.
	//pods informers.UpgradeKubeLister
	podsSynced cache.InformerSynced
	// queue is where incoming work is placed to de-dup and to allow "easy"
	// rate limited requeues on errors
	queue workqueue.RateLimitingInterface

	informer cache.SharedIndexInformer

	kubeInformerFactory kubeinformers.SharedInformerFactory

	kubeClient *kubernetes.Clientset

	newClient *clientset.Clientset
}

func NewController(client *clientset.Clientset, kubeClient *kubernetes.Clientset) *Controller {
	fmt.Println("passing queue")
	queue := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "controller-name")

	fmt.Println("passing informers")
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*60)

	exampleInformerFactory := informers.NewSharedInformerFactory(client, time.Second*200)
	upgradeInformer := exampleInformerFactory.Upgrade().V1().UpgradeKubes()
	sharedIndexUpgradeInfomer := upgradeInformer.Informer()

	fmt.Println("Setting up event handlers")
	sharedIndexUpgradeInfomer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			fmt.Println("Adding key:", key)
			fmt.Println("obj:", obj)
			time.Sleep(30 * time.Second)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old interface{}, newo interface{}) {
			fmt.Println("old:", old)
			fmt.Println("newo:", newo)
			time.Sleep(30 * time.Second)
			if !reflect.DeepEqual(old, newo) {
				key, err := cache.MetaNamespaceKeyFunc(newo)
				if err == nil {
					queue.Add(key)

					fmt.Println("Updating key:", key)
					fmt.Println("UpgradeKube Updated: ")
				}
			} else {
				fmt.Println("not updated")
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
	kubeInformerFactory.Start(context.Background().Done())

	c := &Controller{}
	c.newClient = client
	c.kubeClient = kubeClient
	c.kubeInformerFactory = kubeInformerFactory
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
	time.Sleep(15 * time.Second)
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)

	fmt.Println("processNextWorkItem key:", key.(string))
	time.Sleep(15 * time.Second)
	// do your work on the key.  This method will contains your "do stuff" logic
	err := c.syncHandler(key.(string))
	if err == nil {
		// No error, reset the ratelimit counters

		// if you had no error, tell the queue to stop tracking history for your
		// key. This will reset things like failure counts for per-item rate
		// limiting
		time.Sleep(15 * time.Second)
		fmt.Println("processNextWorkItem Forget key:", key.(string))
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < 5 {
		// err != nil and retry
		fmt.Println("processNextWorkItem if 2nd block retry")
		time.Sleep(15 * time.Second)
		c.queue.AddRateLimited(key)
	} else {
		// err != nil and too many retries
		fmt.Println("processNextWorkItem final else block")
		c.queue.Forget(key)
		time.Sleep(15 * time.Second)
		utilruntime.HandleError(fmt.Errorf("%v failed with : %v", key, err))
	}
	time.Sleep(15 * time.Second)
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
	fmt.Println("SyncHandler Before ProcessKubeUpgrade:", uk)
	uk, err = c.ProcessKubeUpgrade(context.Background(), uk.DeepCopy())
	if err != nil {
		return err
	}
	if uk != nil {
		fmt.Println("SyncHandler Returning from ProcessKubeUpgrade:", uk)
		_, err = c.newClient.UpgradeV1().UpgradeKubes(uk.Namespace).UpdateStatus(context.Background(), uk, metav1.UpdateOptions{})
		if err != nil {
			fmt.Printf("SyncHandler: Error Updating UpgradeKubes")
			return err
		}
		fmt.Println("SyncHandler: After Update ukDeep.Status:", uk.Status)
	}
	return nil
}

func (c *Controller) ProcessKubeUpgrade(ctx context.Context, uk *upgradev1.UpgradeKube) (*upgradev1.UpgradeKube, error) {
	defer fmt.Println("Exiting ProcessKubeUpgrade")
	fmt.Println("Entering ProcessKubeUpgrade")

	// nodes, err := c.kubeInformerFactory.Core().V1().Nodes().Lister().List()
	// if err != nil {
	// 	return nil, err
	// }
	// if nodes == nil {
	// 	return nil, fmt.Errorf("No nodes found")
	// }
	ukDeep := uk
	updated := false
	depTemp := &ukDeep.Spec.DepTemp

	fmt.Println("ProcessKubeUpgrade ukDeep.Status:", ukDeep.Status)
	fmt.Println("ProcessKubeUpgrade ukDeep.Status.DepCreated:", ukDeep.Status.DepCreated)

	if !ukDeep.Status.DepCreated && depTemp != nil {
		obj := newDeployment(depTemp)
		deployment, err := c.kubeClient.AppsV1().Deployments(depTemp.Namespace).Create(ctx, obj, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
		if deployment == nil {
			return nil, fmt.Errorf("nil deployment")
		}
		ukDeep.Status.DepCreated = true
		updated = true
		fmt.Println("ProcessKubeUpgrade Updating ukDeep.Status.DepCreated:", ukDeep.Status.DepCreated)
	}

	servTemp := &ukDeep.Spec.ServiceTemp

	fmt.Println("ProcessKubeUpgrade ukDeep.Status:", ukDeep.Status)
	fmt.Println("ProcessKubeUpgrade ukDeep.Status.SvcCreated:", ukDeep.Status.SvcCreated)
	if !ukDeep.Status.SvcCreated && servTemp != nil {
		obj := newService(servTemp, depTemp)
		service, err := c.kubeClient.CoreV1().Services(depTemp.Namespace).Create(ctx, obj, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
		if service == nil {
			return nil, fmt.Errorf("nil service")
		}
		ukDeep.Status.SvcCreated = true
		updated = true
		fmt.Println("ProcessKubeUpgrade Updating ukDeep.Status.SvcCreated:", ukDeep.Status.SvcCreated)
	}
	if updated {
		fmt.Println("Resource needs to be updated.")
		return ukDeep, nil
	}
	fmt.Println("Resource did not need to be updated.")
	return nil, nil
}

func newDeployment(dep *upgradev1.DeploymentTemplate) *appsv1.Deployment {
	labels := map[string]string{
		"app":        dep.Name,
		"controller": "sample-controller",
	}
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: dep.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &dep.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  dep.Name,
							Image: dep.ImageName,
						},
					},
				},
			},
		},
	}
}

func newService(serviceTemp *upgradev1.ServiceTemplate, dep *upgradev1.DeploymentTemplate) *corev1.Service {
	labels := map[string]string{
		"app":        dep.Name,
		"controller": "sample-controller",
	}
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceTemp.Name,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Type:     corev1.ServiceType(serviceTemp.Type),
			Ports: []corev1.ServicePort{
				corev1.ServicePort{
					Port:     serviceTemp.ServicePort,
					Protocol: corev1.ProtocolTCP,
				},
			},
		},
	}
}

/*
spec:
  deployment:
    namespace:
	name:
	image:
	replicas:
  serviceName:
    name:
	type:
	servicePort:
status:
   DevCreated:
   SvcCreated:
*/
