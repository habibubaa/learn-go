package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	beego.GlobalControllerRouter["learn_go/controllers:LoginController"] = append(beego.GlobalControllerRouter["learn_go/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "GetSSO",
			Router:           "/user",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_go/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_go/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/get/:TRX_DATE/:KODE_UNIT/:SEGMENT/:IMP_STAGE/:search",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_go/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_go/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "GETCABANG",
			Router:           "/GETCABANG/:KODE_WILAYAH",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_go/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_go/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "GETMASTERSEGMENTS",
			Router:           "/GETMASTERSEGMENTS",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_go/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_go/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "UpdateNilai",
			Router:           "/UpdateNilai",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	// Event section

	beego.GlobalControllerRouter["learn_go/controllers:MasterEventController"] = append(beego.GlobalControllerRouter["learn_go/controllers:MasterEventController"],
		beego.ControllerComments{
			Method:           "GETMASTERMATERIS",
			Router:           "/GETMASTERMATERIS",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	
	beego.GlobalControllerRouter["learn_go/controllers:MasterEventController"] = append(beego.GlobalControllerRouter["learn_go/controllers:MasterEventController"],
		beego.ControllerComments{
			Method: "MASTERMATERIADD",
			Router: "/MASTERMATERIADD",
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Filters: nil,
			Params: nil})
}
