package controllers

import (
	"encoding/json"
	"learn_beego/global"
	"learn_beego/models"

	"github.com/astaxie/beego"
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
