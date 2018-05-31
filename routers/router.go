package routers

import (
	"github.com/astaxie/beego"
	"github.com/haborhuang/weixin-comp-go-demo/controllers"
)

func init() {
	beego.AddNamespace(
		beego.NewNamespace("/wx",
			beego.NSNamespace("/auth",
				beego.NSInclude(
					&controllers.CallbackController{},
				),
			),
			beego.NSNamespace("/msg",
				beego.NSInclude(
					&controllers.MsgController{},
				),
			),
		),
	)
}
