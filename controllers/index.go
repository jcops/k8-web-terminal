package controllers

import (
	"github.com/astaxie/beego"
	"log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"fmt"
	"strings"
	"k8s.io/api/core/v1"
	"time"
)


type MainController struct {
	beego.Controller
}


func (this *MainController) URLMapping() {
	this.Mapping("Nodes",this.Nodes)
	this.Mapping("NodePods",this.NodePods)
	this.Mapping("ContainerTerminal",this.ContainerTerminal)
}

// @router / [get]
func (c *MainController) Get() {
	c.TplName = "index.html"


}


// @Title 获取所有节点信息
// @Description 获取所有节点信息
// @Success 200 {string}
// @Failure 404 body is empty
// @router /api/nodes [get]
func (c *MainController) Nodes()  {
	resp, err := Clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	c.Data["json"] =  resp.Items
	c.ServeJSON()
}

// @Title 获取节点上的所有容器
// @Description 获取节点上的所有容器
// @Success 200 {string}
// @Failure 404 body is empty
// @router /api/nodes/containers [get]
func (c *MainController) NodePods()  {
	nodeip := c.GetString("node")
	fmt.Println("sssss",nodeip)
	//resp, err := Clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	nodespod111, err := Clientset.CoreV1().Pods(v1.NamespaceAll).List(metav1.ListOptions{
		FieldSelector: "spec.nodeName=" + nodeip,
	})
	if err != nil {
		log.Fatal(err)
	}
	var mapss []interface{}
	for _, pod := range  nodespod111.Items {
		pods := pod
		var ImageID string
		var ContainerID string
		var Created int64
		var Status string
		if pods.Status.ContainerStatuses != nil {
			ImageID = pods.Status.ContainerStatuses[0].ImageID
			ContainerID = strings.Join(strings.Split(pods.Status.ContainerStatuses[0].ContainerID,"docker://"),"")

			if pods.Status.ContainerStatuses[0].State.Running != nil {
				Created = pods.Status.ContainerStatuses[0].State.Running.StartedAt.Unix()
			}else {
				Created = time.Now().Unix()
			}
		} else  {
			ImageID = ""
			ContainerID = ""

		}
		if pods.Status.Conditions != nil {
			if pods.Status.Conditions[1].Status == "True" {
				Status = "Ready"
			} else {
				Status = pods.Status.Conditions[1].Reason
			}
		} else  {
			Status = "NotReady"
		}
		maps := map[string]interface{}{
			"Name":        pods.Name,
			"Namespace":  pods.Namespace,
			"NodeName":    pods.Spec.NodeName,
			"Labels":      pods.ObjectMeta.Labels,
			"SelfLink":    pods.ObjectMeta.SelfLink,
			"Uid":         pods.ObjectMeta.UID,
			"Status":      Status,
			"IP":          pods.Status.PodIP,
			"Image":       pods.Spec.Containers[0].Image,
			"AppName": 	   pods.Spec.Containers[0].Name,
			"ImageID":     ImageID,
			"Id":          ContainerID,
			"Created":     Created,
			"Command":     pod.Spec.Containers[0].Command,

		}
		mapss = append(mapss, maps)
	}
	fmt.Println(mapss)
	c.Data["json"] =  mapss
	c.ServeJSON()
}


// @router /container/terminal [get]
func (this *MainController) ContainerTerminal() {
	this.TplName = "terminal.html"
}
