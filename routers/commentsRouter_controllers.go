package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["k8-web-terminal/controllers:MainController"] = append(beego.GlobalControllerRouter["k8-web-terminal/controllers:MainController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["k8-web-terminal/controllers:MainController"] = append(beego.GlobalControllerRouter["k8-web-terminal/controllers:MainController"],
        beego.ControllerComments{
            Method: "Nodes",
            Router: `/api/nodes`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["k8-web-terminal/controllers:MainController"] = append(beego.GlobalControllerRouter["k8-web-terminal/controllers:MainController"],
        beego.ControllerComments{
            Method: "NodePods",
            Router: `/api/nodes/containers`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["k8-web-terminal/controllers:MainController"] = append(beego.GlobalControllerRouter["k8-web-terminal/controllers:MainController"],
        beego.ControllerComments{
            Method: "ContainerTerminal",
            Router: `/container/terminal`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
