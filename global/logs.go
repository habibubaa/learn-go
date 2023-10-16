package global

import (
	"github.com/astaxie/beego/logs"
	"fmt"
	"log"
)

// log juga ada di main.go untuk record semua hanya saja kurang detail

func LogSQL(message string) {	
	l := logs.NewLogger(10000)
	err := l.SetLogger(logs.AdapterFile,`{"filename":"logs/error_sql.log","level":7,"maxlines":0,"maxsize":1000000,"daily":true,"maxdays":10,"color":true}`)
	if err != nil {
    	log.Fatal(err)
	}
	l.Error(message)	
	fmt.Println(message)
}

func LogUploadFile(message string) {	
	l := logs.NewLogger(10000)
	err := l.SetLogger(logs.AdapterFile,`{"filename":"logs/upload_file.log","level":7,"maxlines":0,"maxsize":1000000,"daily":true,"maxdays":10,"color":true}`)
	if err != nil {
    	log.Fatal(err)
	}
	l.Error(message)	
	fmt.Println(message)
}

func LogWebSocket(message string) {	
	l := logs.NewLogger(10000)
	err := l.SetLogger(logs.AdapterFile,`{"filename":"logs/web_socket.log","level":7,"maxlines":0,"maxsize":1000000,"daily":true,"maxdays":10,"color":true}`)
	if err != nil {
    	log.Fatal(err)
	}
	l.Error(message)	
	fmt.Println(message)
}