package controllers

import (
	"encoding/json"
	"fmt"
	"learn_go/global"
	"learn_go/models"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/validation"
	//"time"
	//"strconv"
)

// MASTER LEARN
type MasterController struct {
	beego.Controller
}

// @Title Get Flow
// @Description pengambilan semua data GetDataLEARN
// @Param	Authorization header string  true "Authorization Token"
// @Param	TRX_DATE path string  true "TRX_DATE"
// @Param	KODE_UNIT path string  true "KODE_UNIT"
// @Param	SEGMENT path string  false "SEGMENT"
// @Param	IMP_STAGE path string  false "IMP_STAGE"
// @Param	search path string  false "search by Reff Number"
// @Success 200 {object} models.GetDataLEARN {"total": int,"data": array value}
// @Failure 403 "Error Token"
// @router /get/:TRX_DATE/:KODE_UNIT/:SEGMENT/:IMP_STAGE/:search [get]
func (c *MasterController) Get() {
	var paramstring = make(map[string]string)
	paramstring["KODE_UNIT"] = c.Ctx.Input.Param(":KODE_UNIT")
	paramstring["SEGMENT"] = c.Ctx.Input.Param(":SEGMENT")
	paramstring["IMP_STAGE"] = c.Ctx.Input.Param(":IMP_STAGE")
	paramstring["search"] = c.Ctx.Input.Param(":search")
	paramstring["TRX_DATE"] = c.Ctx.Input.Param(":TRX_DATE")

	fmt.Println(paramstring)

	R, err := models.PagingDataLEARN(paramstring)

	if err != nil {
		global.LogSQL(global.GetCurrentFuncName() + " - " + err.Error())
		global.SetHttpError(500, c.Ctx, global.APIResponse{500, err.Error()})
	} else {
		c.Data["json"] = R
	}

	c.ServeJSON()
}

// @Title MASTER_DATA
// @Description master_data
// @Param	Authorization header string  false "Authorization Token"
// @Param	KODE_WILAYAH path string true "KODE_WILAYAH"
// @Success 200 {object} models.CabangDiperiksa
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /GETCABANG/:KODE_WILAYAH [get]
func (c *MasterController) GETCABANG() {
	KODE_WILAYAH := c.Ctx.Input.Param(":KODE_WILAYAH")
	datagetcabang, err := models.GetCabangDiperiksa(KODE_WILAYAH)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data := datagetcabang

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
// @Param   data body models.FieldData true "simpan data"
// @Success 200 {object} global.APILoginResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql / validasi"}
// @router /UpdateNilai [put]
func (c *MasterController) UpdateNilai() {
	var data models.FieldData
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &data)
	fmt.Println(err)

	if err != nil {
		c.Ctx.Output.SetStatus(405)
		c.Data["json"] = global.APILoginResponse{Code: 405, Message: err.Error()}
		c.ServeJSON()
	}

	response, err := models.UpdateNilaiLEARN(data)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APILoginResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	jsondata, err := json.Marshal(&response)
	if err != nil {
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = global.APIGetResponse{Code: 403, Message: err.Error()}
		c.ServeJSON()
	}

	c.Data["json"] = global.APIGetResponse{Code: 200, Message: "Data Berhasil Di Update", Data: jsondata}
	c.ServeJSON()
}

// @Title MASTER_DATA
// @Description master_data
// @Param	Authorization header string  false "Authorization Token"
// @Success 200 {object} models.MASTERSEGMENTGET
// @Failure 403 "Error"
// @Failure 500 "Error"
// @router /GETMASTERSEGMENTS [get]
func (c *MasterController) GETMASTERSEGMENTS() {
	datagetcabang, err := models.GETMASTERSEGMENT()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = global.APIGetResponse{Code: 500, Message: err.Error()}
		c.ServeJSON()
	}

	data := datagetcabang

	jsondata, err := json.Marshal(&data)
	if err != nil {
		c.Ctx.Output.SetStatus(403)
		c.Data["json"] = global.APIGetResponse{Code: 403, Message: err.Error()}
		c.ServeJSON()
	}

	c.Data["json"] = global.APIGetResponse{Code: 200, Message: "Data Berhasil dimuat", Data: jsondata}
	c.ServeJSON()
}
