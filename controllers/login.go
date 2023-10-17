package controllers

import (
	"fmt"
	"io/ioutil"
	"learn_go/global"
	"learn_go/models"
	"net/http"
	"net/http/cookiejar"
	"strconv"

	"github.com/astaxie/beego"
	//"strings"
)

// Login
type LoginController struct {
	beego.Controller
}

// @Title LoginSSO
// @Description Get login sso kode aplikasi
// @Param	username formData string  true "username sso"
// @Param	password formData string  true "password sso"
// @Success 200 {object} global.APILoginResponse {"code": 200,"message": "Data berhasil disimpan"}
// @Failure 404 {"code": 404,"message": "Error not found"}
// @Failure 405 {"code": 405,"message": "Error json"}
// @Failure 500 {"code": 500,"message": "Error sql / validasi"}
// @router /user [post]
func (c *LoginController) GetSSO() {
	username := c.GetString("username")
	password := c.GetString("password")

	//START API SSO GET BY USER, PASS & APP_CODE
	data := "?user=" + username + "&pass=" + password + "&app_code="

	token := global.GenerateToken(data)

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar:           jar,
		CheckRedirect: global.RedirectPolicyFunc,
	}

	req, _ := http.NewRequest("GET", beego.AppConfig.String("sso")+data+beego.AppConfig.String("KEY"), nil)
	req.Header.Add("Authorization", "Basic "+global.BasicAuth("event", "event"))

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("err: ", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		fmt.Println("err: ", err)
	}

	RESPONSE, _ := strconv.ParseBool(global.CheckNil(body, "login.[0].response"))
	IDSDM := ""
	NAMA := ""
	POSISI_NAMA := ""
	LOKASI_KERJA := ""
	USERNAME := ""
	MESSAGE := ""
	NIK := ""
	//BISNIS := ""

	// fmt.Println("err: ", global.CheckNil(body, "login.[0].data.[0].idsdm"))

	if RESPONSE == true {
		IDSDM = global.CheckNil(body, "login.[0].data.[0].idsdm")
		NIK = global.CheckNil(body, "login.[0].data.[0].nik")
		NAMA = global.CheckNil(body, "login.[0].data.[0].nama")
		POSISI_NAMA = global.CheckNil(body, "login.[0].data.[0].posisi_nama")
		LOKASI_KERJA = global.CheckNil(body, "login.[0].data.[0].lokasi_kerja")
		USERNAME = global.CheckNil(body, "login.[0].data.[0].username")
		MESSAGE = global.CheckNil(body, "login.[0].message")
	} else {
		RESPONSE = false
		MESSAGE = "Username atau Password tidak valid / User belum didaftarkan"
	}

	if err != nil {
		global.LogSQL(global.GetCurrentFuncName() + " - " + err.Error())
		global.SetHttpError(500, c.Ctx, global.APILoginResponse{500, err.Error()})
	}

	if RESPONSE == false {
		c.Data["json"] = models.LoginFalse{RESPONSE: RESPONSE, MESSAGE: MESSAGE}
	} else {
		// Data := models.LoginUser{NAMA: NAMA,USERNAME: USERNAME,USER_ID: checkUser[0].User_id,POSITION_NAME: checkUser[0].Position_name,POSITION_CODE: checkUser[0].Position_code,DIVISION_NAME: checkUser[0].Division_name,DIVISION_CODE: checkUser[0].Division_code,EMAIL: checkUser[0].Email}
		Data := models.LoginUser{NAMA: NAMA, USERNAME: USERNAME, USER_ID: IDSDM, POSITION_NAME: POSISI_NAMA, NIK: NIK, LOKASI_KERJA: LOKASI_KERJA}
		// c.Data["json"] = models.LoginTrue{RESPONSE: RESPONSE,MESSAGE: MESSAGE,TOKEN: token,DATA: Data}
		c.Data["json"] = models.LoginTrue{RESPONSE: RESPONSE, MESSAGE: MESSAGE, TOKEN: token, DATA: Data}
	}

	c.ServeJSON()
}
