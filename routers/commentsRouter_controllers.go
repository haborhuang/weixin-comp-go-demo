package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:CallbackController"] = append(beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:CallbackController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/events`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:CallbackController"] = append(beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:CallbackController"],
		beego.ControllerComments{
			Method: "AuthLink",
			Router: `/auth_link`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:CallbackController"] = append(beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:CallbackController"],
		beego.ControllerComments{
			Method: "AuthCallback",
			Router: `/auth_cb`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:MsgController"] = append(beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:MsgController"],
		beego.ControllerComments{
			Method: "GetTmpls",
			Router: `/apps/:app/templates`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:MsgController"] = append(beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:MsgController"],
		beego.ControllerComments{
			Method: "HandleEvents",
			Router: `/events/:app`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:MsgController"] = append(beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:MsgController"],
		beego.ControllerComments{
			Method: "SendTmplMsg",
			Router: `/apps/:app/templates/:tmpl/message`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:MsgController"] = append(beego.GlobalControllerRouter["github.com/haborhuang/weixin-comp-go-demo/controllers:MsgController"],
		beego.ControllerComments{
			Method: "GetToken",
			Router: `/apps/:app/token`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
