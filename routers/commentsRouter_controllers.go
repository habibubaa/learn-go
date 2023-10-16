package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	beego.GlobalControllerRouter["learn_beego/controllers:LoginController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "GetSSO",
			Router:           "/user",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/get/:TRX_DATE/:KODE_UNIT/:SEGMENT/:IMP_STAGE/:search",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "GETCABANG",
			Router:           "/GETCABANG/:KODE_WILAYAH",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "GETIMPSTAGE",
			Router:           "/GETIMPSTAGE",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "GETMASTERSEGMENTS",
			Router:           "/GETMASTERSEGMENTS",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "GETSEGMENTS",
			Router:           "/GETSEGMENTS/:KODE_UNIT",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "GETUNIT",
			Router:           "/GETUNIT/:KODE_CABANG",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "GETWILAYAH",
			Router:           "/GETWILAYAH",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "UpdateNilai",
			Router:           "/UpdateNilai",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:MasterController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:MasterController"],
		beego.ControllerComments{
			Method:           "UpdateNilaiExcel",
			Router:           "/UpdateNilaiExcel",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "GETDASHBOARDSEBARAN",
			Router:           "/GETDASHBOARDSEBARAN",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "GETDASHBOARDTOP",
			Router:           "/GETDASHBOARDTOP",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "GETDASHBOARDSTAGE",
			Router:           "/GETDASHBOARDSTAGE/:MONTH",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "GETDASHBOARDPROGRESSION",
			Router:           "/GETDASHBOARDPROGRESSION",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "GETDASHBOARDTIPETRANSAKSILEARNS",
			Router:           "/GETDASHBOARDTIPETRANSAKSILEARNS",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"] = append(beego.GlobalControllerRouter["learn_beego/controllers:DashboardController"],
		beego.ControllerComments{
			Method:           "GETDASHBOARDKOLEKTIBILITASS",
			Router:           "/GETDASHBOARDKOLEKTIBILITAS/:MONTH",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})
}
