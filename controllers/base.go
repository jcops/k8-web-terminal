package controllers

import (
	"os"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"flag"
	"path/filepath"
)

var Clientset *kubernetes.Clientset
var Config *rest.Config

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
func init() {
	// 配置 k8s 集群外 kubeconfig 配置文件，默认位置 $HOME/.kube/config
	//
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	}else {
		kubeconfig = flag.String("kubeconfig", "", "(optional) absolute path to the kubeconfig file")
	}
	flag.Parse()

	//在 kubeconfig 中使用当前上下文环境，config 获取支持 url 和 path 方式
	Config, err =rest.InClusterConfig()
	if err != nil {
		Config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}
		Clientset, err = kubernetes.NewForConfig(Config)
		if err != nil {
			panic(err.Error())
		}
	}else {
		Clientset, err = kubernetes.NewForConfig(Config)
		if err != nil {
			panic(err.Error())
		}
	}



   ////  集群内方式
	//Config, err = rest.InClusterConfig()
	//if err != nil {
	//	panic(err.Error())
	//}
   //
	//Clientset, err = kubernetes.NewForConfig(Config)
	//if err != nil {
	//	panic(err.Error())
	//}


}
