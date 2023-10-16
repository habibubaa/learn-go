package global

import (
	"database/sql"
	"errors"

	"github.com/astaxie/beego"
	_ "github.com/denisenkom/go-mssqldb"

	// "fmt"
	"github.com/funujikai/godb"
	"github.com/funujikai/godb/adapters/mssql"
)

func Conn() (*godb.DB, error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("sqlconnLEARN"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func ConnSSO() (*godb.DB, error) {

	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("sqlconnSSO"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// func ConnMKRIntegrasi() (*sql.DB,error) {
// 	conn,err := sql.Open("mssql", beego.AppConfig.String("sqlconnMKRIntegrasi"))
// 	if err != nil {
// 	 	return nil,err
// 	}

// 	err = conn.Ping()
// 	if err != nil {
// 	 	return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
// 	}

// 	return conn,nil
// }

// func ConnHRIS() (*sql.DB,error) {
// 	conn,err := sql.Open("mssql", beego.AppConfig.String("sqlconnHRIS"))
// 	if err != nil {
// 	 	return nil,err
// 	}

// 	err = conn.Ping()
// 	if err != nil {
// 	 	return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
// 	}

// 	return conn,nil
// }

func ConnSimWas() (*sql.DB, error) {
	conn, err := sql.Open("mssql", beego.AppConfig.String("sqlconnLEARN"))
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}

	return conn, nil
}

func ConnSimWas_GODB() (*godb.DB, error) {
	conn, err := godb.Open(mssql.Adapter, beego.AppConfig.String("sqlconnLEARN"))
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}

	return conn, nil
}
