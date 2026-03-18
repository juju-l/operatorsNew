package main

import (
			"context"
			"encoding/json"
			"sync"
			"k8s.io/apimachinery/pkg/runtime/schema"
			"k8s.io/client-go/dynamic"
			"k8s.io/client-go/dynamic/dynamicinformer"
			"k8s.io/client-go/rest"
			"k8s.io/client-go/tools/cache"
			"k8s.io/client-go/tools/clientcmd"
			"k8s.io/client-go/util/workqueue"
			"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
			"time"
			"os"
)

var gvr = schema.GroupVersionResource { Group: "vipex.cc", Version: "v1alpha1", Resource: "hlms" } //

type HlmController struct{
			ctx context.Context
			cancel context.CancelFunc
			cli dynamic.Interface
			inf cache.SharedIndexInformer
			que workqueue.RateLimitingInterface
			objLock sync.Map
			wkCnt int
}

func (ctl *HlmController)run() {
			hdl(ctl.ctx,ctl.inf,ctl.que)
			go ctl.inf.Run(ctl.ctx.Done()) /**/
			if cache.WaitForCacheSync(ctl.ctx.Done(), ctl.inf.HasSynced /**/) {
			for i := 0; i < ctl.wkCnt; i++ {
			go ctl.runWorker() /**/
			}
			}
			/*go*/<-ctl.ctx.Done()/**/
			ctl.que.ShutDown()
}

func (ctl *HlmController)runWorker() {
	for {
			o, dwn := ctl.que.Get()
			if dwn { break }
			defer ctl.que.Done(o) /**/

			k, ok := o.(string)
			if !ok {
			/**/ctl.que.Forget(o);continue
			}
			//

			l, _ := ctl.
			objLock.
			LoadOrStore(
			k, &sync.Mutex{}, /**/
			)
			l.(*sync.Mutex).Lock()
			defer func() {
			//
			l.(*sync.Mutex).Unlock()
			ctl.objLock.Delete(k)
			}()

			if err := ctl.reconcile(k); err != nil {
			ctl.que.AddRateLimited(k)
			} else {
			ctl.que.Forget(o)
			}
			
			continue
	}
}

func (ctl *HlmController)reconcile(
			key string,
	) error {
	var err error
			s,n,e := cache.SplitMetaNamespaceKey(key)
			if e != nil {
			return e
			}
			o,i,e := ctl.inf.GetIndexer().GetByKey(key)
			if e != nil {
			return e
			}
			if ! i {
			return nil
			}
			hlm := &Hlm{} //o.(*Hlm)
			b,_ := json.Marshal(o)
			err = json.Unmarshal(b, hlm /**/)
			helms(
			"tpl_vipex_cc-0.1.0.tgz",n,s,"sta",hlm,
			)
	return err
}

func hdl (
			ctx context.Context,
			inf cache.SharedIndexInformer,
			que workqueue.RateLimitingInterface,
	) {
	inf.AddEventHandler(
	cache.ResourceEventHandlerFuncs{
    AddFunc: func(obj interface{}) {
			if u,ok:=obj.
			(*unstructured.Unstructured);
			ok {
			que.Add(
			vmust[string](
			cache.MetaNamespaceKeyFunc(u),
			),
			)
			}
	},
    UpdateFunc: func(old, new interface{}) {
			if n,ok:=new.
			(*unstructured.Unstructured);
			ok {
			if o,ok:=old.
			(*unstructured.Unstructured);
			ok {
			if n.
			GetResourceVersion()!=o.
			GetResourceVersion() {
			que.Add(
			vmust[string](
			cache.MetaNamespaceKeyFunc(n),
			),
			)
			}
			}
			}
	},
    DeleteFunc: func(obj interface{}) {
			if u,ok:=obj.
			(*unstructured.Unstructured);
			ok {
			que.Add(
			vmust[string](
			cache.MetaNamespaceKeyFunc(u),
			),
			)
			}else{
			if t,ok:=obj.
			(cache.DeletedFinalStateUnknown);
			ok {
			if u,ok:=t.Obj.
			(*unstructured.Unstructured);
			ok {
			if true {
			que.Add(
			vmust[string](
			cache.MetaNamespaceKeyFunc(u),
			),
			)
			}
			}
			}
			}
	},
	},
	)
}

func New (
			ctx context.Context,
			cancel context.CancelFunc,
			wkCnt int,
	) *HlmController{
	var ctl = &HlmController{}
			ctl.ctx = ctx
			ctl.cancel = cancel
			ctl.cli = cli()
			ctl.inf = inf(ctl.cli)
			ctl.que = que()
			ctl.objLock = sync.Map{}
			ctl.wkCnt = wkCnt
	return ctl
}

func inf (
			cli dynamic.Interface,
	) cache.SharedIndexInformer {
	return dynamicinformer.
			NewDynamicSharedInformerFactory(
			cli, 30 * time.Second, /**/
			).
			ForResource(gvr).
			Informer()//.
	// AddEventHandler ///
}

func cli() *dynamic.DynamicClient {
			// inPod ->rest.InClusterConfig()
			return vmust[*dynamic.DynamicClient](dynamic.NewForConfig(vmust[*rest.Config](clientcmd.BuildConfigFromFlags("",os.Getenv("KUBECONFIG")/**/))))
			///
}

func que() workqueue.RateLimitingInterface {
			return workqueue.
			NewNamedRateLimitingQueue(
			/**/ workqueue.DefaultControllerRateLimiter(), "hlmOperator",
			)//.
			/////
}

func vmust[T any] (
			v T, err error, /**/
			) T {
			//
			if err != nil { panic(err) }
			return v
}