package middleware

import (
	"encoding/json"
	"fmt"
	"learn_beego/global"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	//"strings"
)

func init() {
}

func Jwt(ctx *context.Context) {
	ctx.Output.Header("Content-Type", "application/json")
	var uri string = ctx.Input.URI()

	//karena login gak menggunakan token
	if uri == "/v1/login/user" {
		return
	}

	// fmt.Println(uri)

	// karena Preflight Request tidak mengirim token
	if ctx.Input.Method() == "OPTIONS" {
		return
	}

	if ctx.Input.Header("Authorization") == "" {
		ctx.Output.SetStatus(403)
		resBody, err := json.Marshal(global.APILoginResponse{403, "notAllowed"})
		err = ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}

	var tokenString string = ctx.Input.Header("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(beego.AppConfig.String("KEY")), nil
	})

	if err != nil {
		ctx.Output.SetStatus(403)
		var responseBody global.APILoginResponse = global.APILoginResponse{403, err.Error()}
		resBytes, err := json.Marshal(responseBody)
		err = ctx.Output.Body(resBytes)
		if err != nil {
			panic(err)
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && claims != nil {
		return
	} else {
		ctx.Output.SetStatus(403)
		resBody, err := json.Marshal(global.APILoginResponse{403, ctx.Input.Header("Authorization")})
		err = ctx.Output.Body(resBody)
		if err != nil {
			panic(err)
		}
	}
}
