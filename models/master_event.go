package models

import (
	"learn_go/global"

	//"strconv"

	"fmt"
	"time"
)

// CREATE
type MASTERMATERIADD struct {
	id string `json:"id" `
	id_materi string `json:"id_materi" `
	id_klasifikasi_materi string `json:"id_klasifikasi_materi" `
	nama_materi string `json:"nama_materi" `
	is_active string `json:"is_active" `
	created_by string `json:"created_by" `
	created_date string `json:"created_date" `
	modified_by string `json:"modified_by" `
	modified_date string `json:"modified_date" `
	stat string `json:"stat" `
}

func ADDMASTERMATERI(data_arr []MASTERMATERIADD)(string,error){
	var query,logs string
	
	db,err := global.ConnSimWas_GODB()
	if err != nil {
		return logs,err
	} 
	defer db.Close()

	err = db.Begin() //Begin Transaction
	if err != nil {return logs,err}		

	// date now
	now := time.Now()
	nows := now.Format("2006-01-02")
	fmt.Println(nows)

	for _, data := range data_arr {

		// if data.stat == "1"{
			query = "INSERT INTO em_materi (id_materi,id_klasifikasi_materi,nama_materi,is_active,created_date,created_by,modified_date,modified_by) "+
					"VALUES ('"+data.id_materi+"','"+data.id_klasifikasi_materi+"','"+data.nama_materi+"', 'active','2023-10-17','2','2023-10-17', '3')"
			
			fmt.Println(query)

		// } else {
		// 	query = "UPDATE em_materi"+
		// 	"	SET id_materi = "+data.id_materi+""+
		// 	"   ,id_klasifikasi_materi = '"+data.id_klasifikasi_materi+"')"+
		// 	"   ,is_active = '"+data.is_active+"')"+
		// 	"   ,modified_date = CONVERT(Datetime, '"+nows+"')"+
		// 	"   ,modified_by = 1"+
		// 	"	WHERE id = "+data.id+""
		// }
		logs += " models.ADDMASTERMATERI "+query+" \n\n"
		_ , err = db.RawSQL(query).DoWithIterator()
		if err != nil {
			errsql := err
			err = db.Rollback() //Rollback Transaction
			if err != nil {return logs,err}
			return logs,errsql
		}
	}

	err = db.Commit() //Commit Transaction
	if err != nil {return logs,err}
   
	return logs,nil
}

// READ
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

// UPDATE

// DELETE