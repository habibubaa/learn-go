package controllers

import (
	"encoding/json"
	"learn_go/global"
	"learn_go/models"

	"github.com/astaxie/beego"
	"fmt"
	// "github.com/astaxie/beego/validation"
	//"time"
	//"strconv"
)

// MASTER EVENT
type MasterEventController struct {
	beego.Controller
}

//======================================= EVENT API SECTION =========================================//

// @Title MASTER_DATA
// @Description master_data
// @Param	Authorization header string  false "Authorization Token"
// @Success 200 {object} models.MASTERMATERIGET
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /GETMASTERMATERIS [get]
func (c *MasterEventController) GETMASTERMATERIS() {
	datagetmateri, err := models.GETMASTERMATERI()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data := datagetmateri

	jsondata, err := json.Marshal(&data)
	if err != nil {
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = global.APIGetResponse{Code: 403, Message: err.Error()}
		c.ServeJSON()
	}

	c.Data["json"] = global.APIGetResponse{Code: 200, Message: "Data Berhasil dimuat", Data: jsondata}
	c.ServeJSON()
}


// @Title MASTER_DATA
// @Description master_data
// @Param	Authorization header string  false "Authorization Token"
// @Param   data body []models.MASTERMATERIADD true "simpan data"
// @Success 200 {object} global.APILoginResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql / validasi"}
// @router /MASTERMATERIADD [post]
func (c *MasterEventController) MASTERMATERIADD() {
	var data []models.MASTERMATERIADD
	fmt.Println(data)
	var Clogs,Mlogs string
	Clogs = " ================ v1/master-event/PostMateri ================  "

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	
	if err != nil {		
		// global.Logging("ERROR",Clogs+" controllers.PostMateri  json.Unmarshal ---> " + err.Error())
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APILoginResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
		return
	}	

	Mlogs,err = models.ADDMASTERMATERI(data)
	if err != nil {
		// global.Logging("ERROR",Clogs+Mlogs+" controllers.PostMateri models.ADDMASTERMATERI ---> " + err.Error())
		c.Ctx.Output.SetStatus(500)	
		c.Data["json"] = global.APILoginResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
		return
	}
	
	Clogs += Mlogs
	Clogs += " ================ v1/pkmb/PostPKMB ================ "+" \n\n"
	// global.Logging("PAYLOAD",Clogs)	
	
	c.Data["json"] = global.APILoginResponse{Code: 200, Message: "Data berhasil disimpan"}
    c.ServeJSON()	
}