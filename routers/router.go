// @APIVersion 1.0.0
// @Title LEARN API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"learn_beego/controllers"
	"learn_beego/middleware"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(middleware.Jwt),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/master",
			beego.NSInclude(
				&controllers.MasterController{},
			),
		),
		beego.NSNamespace("/master-event",
			beego.NSInclude(
				&controllers.MasterEventController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
