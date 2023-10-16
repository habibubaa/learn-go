package global

import (
	"github.com/astaxie/beego"
	"fmt"
    "runtime"
	"encoding/base64"
	"net/http"
	"github.com/thedevsaddam/gojsonq"
	"strconv"
	"time"	
	"crypto/md5"
	"github.com/dgrijalva/jwt-go"	
	"io"
	"log"		
	"encoding/json"	
	"github.com/astaxie/beego/context"
    //"os"
	"strings"
	"github.com/beevik/guid"	
	"errors"
	"io/ioutil"
)


func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func RedirectPolicyFunc(req *http.Request, via []*http.Request) error{
	req.Header.Add("Authorization","Basic " + BasicAuth("event","event"))
	return nil
}

func CheckNil(body []byte,data string)string{
	err := gojsonq.New().FromString(string(body)).Find(data)
	if err != nil {
		return gojsonq.New().FromString(string(body)).Find(data).(string)
	}else{
		return ""
	}
}

func GenerateToken(d string) string {
	var uid int = 0
	currentTimestamp := time.Now().UTC().Unix()
	//var ttl int64 = 3600 //satu jam
	//var ttl int64 = 43200 //dua belas jam
	h := md5.New()
	_,err := io.WriteString(h, strconv.Itoa(uid))

	if err != nil {
    	log.Fatal(err)
	}

	_,err = io.WriteString(h, strconv.FormatInt(int64(currentTimestamp), 10))

	if err != nil {
    	log.Fatal(err)
	}	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": uid,
		"iss": d,
		"jti": h.Sum(nil),
	})

	tokenString, err := token.SignedString([]byte(beego.AppConfig.String("KEY")))

	if err != nil {
    	log.Fatal(err)
	}

	return (tokenString)
}

func SetHttpError(Code int,ctx *context.Context,data interface{}){	
	ctx.Output.SetStatus(Code)
	resBytes, err := json.Marshal(data)
	err = ctx.Output.Body(resBytes)
	if err != nil {
		panic(err)
	}
}

func UploadFile(file string,path string)(string,error){

		onlyfile := file[strings.IndexByte(file, ',')+1:]

		mediatype := ""
		
		switch file[0:strings.IndexByte(file,';')+1]{
			case "data:application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;":
				mediatype=".xlsx"
			case "data:application/pdf;":
				mediatype=".pdf"
			case "data:application/vnd.openxmlformats-officedocument.wordprocessingml.document;":
				mediatype=".docx"
			default:
				mediatype="error"	
		}

		// fmt.Println(media)
		if mediatype == "error"{
			return "",errors.New("file tidak di izinkan untuk di upload")
		}

		pathfile := path+guid.New().String()+mediatype
	
		// f, err := os.Create(beego.AppConfig.String("upload_path")+pathfile)
		// if err != nil {
		// 	return "",err
		// }
		// err = f.Close()
		// if err != nil {
		// 	return "",err
		// }		

		dec, err := base64.StdEncoding.DecodeString(onlyfile)
		if err != nil {
			return "",err
		}		
	
		// if _, err := f.Write(dec); err != nil {
		// 	return "",err
		// }
		// if err = f.Sync(); err != nil {
		// 	return "",err
		// }

		err = ioutil.WriteFile(beego.AppConfig.String("upload_path")+pathfile, []byte(dec), 0775);

		return pathfile,nil
}
