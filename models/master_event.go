package models

import (
	"learn_beego/global"

	//"strconv"

	"fmt"
)

type MASTERMATERIGET struct {
	MATERI interface{} `json:"MATERI" db:"MATERI"`
}

func GETMASTERMATERI() ([]MASTERMATERIGET, error) {
	var data []MASTERMATERIGET
	db, err := global.ConnSimWas()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := "select distinct nama_materi from em_materi order by nama_materi asc"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var each = MASTERMATERIGET{}
		err = rows.Scan(&each.MATERI)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}

	return data, nil
}
