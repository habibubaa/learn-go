package models

import (
	"encoding/json"
	"learn_go/global"

	//"strconv"

	"fmt"
)

type GetDataLEARN struct {
	ID                    int         `json:"ID" db:"ID"`
	TRX_DATE              interface{} `json:"TRX_DATE" db:"TRX_DATE"`
	LEARN_EXPECTED        float64     `json:"LEARN_EXPECTED" db:"LEARN_EXPECTED"`
	REFF_NUMBER           string      `json:"REFF_NUMBER" db:"REFF_NUMBER"`
	IMP_STAGE             string      `json:"IMP_STAGE" db:"IMP_STAGE"`
	SEGMENT               string      `json:"SEGMENT" db:"SEGMENT"`
	LEARN_ADJUST_EXPECTED float64     `json:"LEARN_ADJUST_EXPECTED" db:"LEARN_ADJUST_EXPECTED"`
	LEARN_TOTAL           float64     `json:"LEARN_TOTAL" db:"LEARN_TOTAL"`
	KODE_UNIT             interface{} `json:"KODE_UNIT" db:"KODE_UNIT"`
	NAMA_UNIT             interface{} `json:"NAMA_UNIT" db:"NAMA_UNIT"`
	INISIAL               interface{} `json:"INISIAL" db:"INISIAL"`
	KODE_CABANG           interface{} `json:"KODE_CABANG" db:"KODE_CABANG"`
	NAMA_CABANG           interface{} `json:"NAMA_CABANG" db:"NAMA_CABANG"`
	KODE_WILAYAH          interface{} `json:"KODE_WILAYAH" db:"KODE_WILAYAH"`
	NAMA_WILAYAH          interface{} `json:"NAMA_WILAYAH" db:"NAMA_WILAYAH"`
	IS_SYARIAH            interface{} `json:"IS_SYARIAH" db:"IS_SYARIAH"`
	IS_ULAMM              interface{} `json:"IS_ULAMM" db:"IS_ULAMM"`
	IS_VC                 interface{} `json:"IS_VC" db:"IS_VC"`
	IS_VS                 interface{} `json:"IS_VS" db:"IS_VS"`
	LAST_BALANCE          float64     `json:"LAST_BALANCE" db:"LAST_BALANCE"`
	ORIG_COLLECT          interface{} `json:"ORIG_COLLECT" db:"ORIG_COLLECT"`
	NAMA_NASABAH          interface{} `json:"NAMA_NASABAH" db:"NAMA_NASABAH"`
	TIPE_KREDIT           interface{} `json:"TIPE_KREDIT" db:"TIPE_KREDIT"`
}

func PagingDataLEARN(paramstring map[string]string) (global.PagingResult, error) {
	data := make([]GetDataLEARN, 0)
	var count global.PagingCount
	var Rdata global.PagingResult
	var sumdata global.ALLLEARN
	var sumadjust global.LEARNADJUST
	var sumexpected global.LEARNEXPECTED
	db, err := global.Conn()
	if err != nil {
		return Rdata, err
	}
	var qmain, qcount, qsum, qsumadjust, qsumexpected string

	if paramstring["search"] != "undefined" {
		paramstring["search"] = " and REFF_NUMBER like '" + paramstring["search"] + "%' "
	} else {
		paramstring["search"] = ""
	}

	if paramstring["SEGMENT"] != "undefined" {
		paramstring["SEGMENT"] = " and SEGMENT = '" + paramstring["SEGMENT"] + "' "
	} else {
		paramstring["SEGMENT"] = ""
	}

	if paramstring["IMP_STAGE"] != "undefined" {
		paramstring["IMP_STAGE"] = " and IMP_STAGE = '" + paramstring["IMP_STAGE"] + "' "
	} else {
		paramstring["IMP_STAGE"] = ""
	}

	if paramstring["KODE_UNIT"] == "NULL" {
		paramstring["KODE_UNIT"] = " and KODE_UNIT is null "
	} else {
		paramstring["KODE_UNIT"] = " and KODE_UNIT = '" + paramstring["KODE_UNIT"] + "' "
	}

	qcount = "SELECT COUNT(*) as total FROM m_LEARN_EXPECTED_DEV where TRX_DATE = '" + paramstring["TRX_DATE"] + "' and KODE_WILAYAH is not null " + paramstring["KODE_UNIT"] + " " +
		"" + paramstring["SEGMENT"] + " " +
		"" + paramstring["IMP_STAGE"] + " " +
		"" + paramstring["search"] + " "

	err = db.RawSQL(qcount).Do(&count)

	if err != nil {
		return Rdata, err
	}

	if count.TOTAL > 0 {
		qmain = "SELECT	ID,TRX_DATE,ISNULL(LEARN_EXPECTED, 0) as LEARN_EXPECTED,REFF_NUMBER,IMP_STAGE,SEGMENT,ISNULL(LEARN_ADJUST_EXPECTED, 0) as LEARN_ADJUST_EXPECTED, " +
			"(ISNULL(LEARN_EXPECTED, 0) + ISNULL(LEARN_ADJUST_EXPECTED, 0)) as LEARN_TOTAL,KODE_UNIT,NAMA_UNIT,INISIAL,KODE_CABANG " +
			",NAMA_CABANG,KODE_WILAYAH,NAMA_WILAYAH,IS_SYARIAH,IS_ULAMM,IS_VC,IS_VS,(ISNULL(LAST_BALANCE, 0)) as LAST_BALANCE, ORIG_COLLECT, NAMA_NASABAH, TIPE_KREDIT " +
			"FROM m_LEARN_EXPECTED_DEV where TRX_DATE = '" + paramstring["TRX_DATE"] + "' and KODE_WILAYAH is not null " + paramstring["KODE_UNIT"] + " " +
			"" + paramstring["SEGMENT"] + "" +
			"" + paramstring["IMP_STAGE"] + "" +
			"" + paramstring["search"] + " " +
			"Order By TRX_DATE DESC "

		err = db.RawSQL(qmain).Do(&data)

		if err != nil {
			return Rdata, err
		}

		qsum = " SELECT sum(ISNULL(LEARN_EXPECTED, 0) + ISNULL(LEARN_ADJUST_EXPECTED, 0)) as total FROM m_LEARN_EXPECTED_DEV where TRX_DATE = '" + paramstring["TRX_DATE"] + "' and KODE_WILAYAH is not null " + paramstring["KODE_UNIT"] + " " +
			"" + paramstring["SEGMENT"] + " " +
			"" + paramstring["IMP_STAGE"] + " " +
			"" + paramstring["search"] + " "

		err = db.RawSQL(qsum).Do(&sumdata)

		if err != nil {
			return Rdata, err
		}

		qsumexpected = " SELECT sum(ISNULL(LEARN_EXPECTED, 0)) as totalexpected FROM m_LEARN_EXPECTED_DEV where TRX_DATE = '" + paramstring["TRX_DATE"] + "' and KODE_WILAYAH is not null " + paramstring["KODE_UNIT"] + " " +
			"" + paramstring["SEGMENT"] + " " +
			"" + paramstring["IMP_STAGE"] + " " +
			"" + paramstring["search"] + " "

		err = db.RawSQL(qsumexpected).Do(&sumexpected)

		if err != nil {
			return Rdata, err
		}

		qsumadjust = " SELECT sum(ISNULL(LEARN_ADJUST_EXPECTED, 0)) as totaladjust FROM m_LEARN_EXPECTED_DEV where TRX_DATE = '" + paramstring["TRX_DATE"] + "' and KODE_WILAYAH is not null " + paramstring["KODE_UNIT"] + " " +
			"" + paramstring["SEGMENT"] + " " +
			"" + paramstring["IMP_STAGE"] + " " +
			"" + paramstring["search"] + " "

		err = db.RawSQL(qsumadjust).Do(&sumadjust)

		if err != nil {
			return Rdata, err
		}
	}

	jsondata, err := json.Marshal(&data)
	if err != nil {
		return Rdata, err
	}

	Rdata = global.PagingResult{TOTAL: count.TOTAL, SUM: sumdata.TOTAL, SUMEXPECTED: sumexpected.TOTALEXPECTED, SUMADJUST: sumadjust.TOTALADJUST, DATA: jsondata}

	err = db.Close()
	if err != nil {
		return Rdata, err
	}

	return Rdata, nil
}

type CabangDiperiksa struct {
	KODE_CABANG interface{} `json:"KODE_CABANG" db:"KODE_CABANG"`
	NAMA_CABANG interface{} `json:"NAMA_CABANG" db:"NAMA_CABANG"`
}

func GetCabangDiperiksa(KODE_WILAYAH string) ([]CabangDiperiksa, error) {
	var data []CabangDiperiksa
	db, err := global.ConnSimWas()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := "select distinct KODE_CABANG, LTRIM(RTRIM(NAMA_CABANG)) as NAMA_CABANG from m_LEARN_EXPECTED_DEV where KODE_WILAYAH = '" + KODE_WILAYAH + "' and KODE_WILAYAH is not null order by NAMA_CABANG asc"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var each = CabangDiperiksa{}
		err = rows.Scan(&each.KODE_CABANG, &each.NAMA_CABANG)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}

	return data, nil
}

type SEGMENTGET struct {
	SEGMENT interface{} `json:"SEGMENT" db:"SEGMENT"`
}

func GetSEGMENT(KODE_UNIT string) ([]SEGMENTGET, error) {
	var data []SEGMENTGET
	db, err := global.ConnSimWas()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := "select distinct SEGMENT from m_LEARN_EXPECTED_DEV where KODE_WILAYAH is not null and KODE_UNIT = '" + KODE_UNIT + "' order by SEGMENT asc"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var each = SEGMENTGET{}
		err = rows.Scan(&each.SEGMENT)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}

	return data, nil
}

type UpdateDataLEARN struct {
	ID                    int         `json:"ID" db:"ID"`
	TRX_DATE              interface{} `json:"TRX_DATE" db:"TRX_DATE"`
	LEARN_EXPECTED        float64     `json:"LEARN_EXPECTED" db:"LEARN_EXPECTED"`
	REFF_NUMBER           string      `json:"REFF_NUMBER" db:"REFF_NUMBER"`
	IMP_STAGE             string      `json:"IMP_STAGE" db:"IMP_STAGE"`
	SEGMENT               string      `json:"SEGMENT" db:"SEGMENT"`
	LEARN_ADJUST_EXPECTED float64     `json:"LEARN_ADJUST_EXPECTED" db:"LEARN_ADJUST_EXPECTED"`
	LEARN_TOTAL           float64     `json:"LEARN_TOTAL" db:"LEARN_TOTAL"`
	KODE_UNIT             interface{} `json:"KODE_UNIT" db:"KODE_UNIT"`
	NAMA_UNIT             interface{} `json:"NAMA_UNIT" db:"NAMA_UNIT"`
	INISIAL               interface{} `json:"INISIAL" db:"INISIAL"`
	KODE_CABANG           interface{} `json:"KODE_CABANG" db:"KODE_CABANG"`
	NAMA_CABANG           interface{} `json:"NAMA_CABANG" db:"NAMA_CABANG"`
	KODE_WILAYAH          interface{} `json:"KODE_WILAYAH" db:"KODE_WILAYAH"`
	NAMA_WILAYAH          interface{} `json:"NAMA_WILAYAH" db:"NAMA_WILAYAH"`
	IS_SYARIAH            interface{} `json:"IS_SYARIAH" db:"IS_SYARIAH"`
	IS_ULAMM              interface{} `json:"IS_ULAMM" db:"IS_ULAMM"`
	IS_VC                 interface{} `json:"IS_VC" db:"IS_VC"`
	IS_VS                 interface{} `json:"IS_VS" db:"IS_VS"`
	NAMA_USER_EDIT        interface{} `json:"NAMA_USER_EDIT" db:"NAMA_USER_EDIT"`
	NIP_USER_EDIT         interface{} `json:"NIP_USER_EDIT" db:"NIP_USER_EDIT"`
	TANGGAL_EDIT          interface{} `json:"TANGGAL_EDIT" db:"TANGGAL_EDIT"`
	LAST_BALANCE          float64     `json:"LAST_BALANCE" db:"LAST_BALANCE"`
	ORIG_COLLECT          interface{} `json:"ORIG_COLLECT" db:"ORIG_COLLECT"`
}

type FieldData struct {
	ID                    string `json:"ID" db:"ID"`
	LEARN_ADJUST_EXPECTED string `json:"LEARN_ADJUST_EXPECTED" db:"LEARN_ADJUST_EXPECTED"`
	NAMA_USER_EDIT        string `json:"NAMA_USER_EDIT" db:"NAMA_USER_EDIT"`
	NIP_USER_EDIT         string `json:"NIP_USER_EDIT" db:"NIP_USER_EDIT"`
	SEGMENT               string `json:"SEGMENT" db:"SEGMENT"`
	IMP_STAGE             string `json:"IMP_STAGE" db:"IMP_STAGE"`
	TRX_DATE              string `json:"TRX_DATE" db:"TRX_DATE"`
	KODE_UNIT             string `json:"KODE_UNIT" db:"KODE_UNIT"`
}

type FIELDSUM struct {
	TOTAL         float64 `json:"TOTAL" db:"total"`
	TOTALADJUST   float64 `json:"TOTALADJUST" db:"totaladjust"`
	TOTALEXPECTED float64 `json:"TOTALEXPECTED" db:"totalexpected"`
}

func UpdateNilaiLEARN(data FieldData) (global.PagingResult, error) {
	datass := make([]UpdateDataLEARN, 0)
	var Rdata global.PagingResult
	var sumdata FIELDSUM
	db, err := global.ConnSimWas_GODB()
	if err != nil {
		return Rdata, err
	}
	defer db.Close()

	var query, qsum string

	if data.SEGMENT != "undefined" {
		data.SEGMENT = " and SEGMENT = '" + data.SEGMENT + "' "
	} else {
		data.SEGMENT = ""
	}

	if data.IMP_STAGE != "undefined" {
		data.IMP_STAGE = " and IMP_STAGE = '" + data.IMP_STAGE + "' "
	} else {
		data.IMP_STAGE = ""
	}

	err = db.Begin() //Begin Transaction
	if err != nil {
		return Rdata, err
	}

	query = "EXEC p_update_adjust " +
		"@id=" + data.ID + ", " +
		"@nama='" + data.NAMA_USER_EDIT + "', " +
		"@nip='" + data.NIP_USER_EDIT + "', " +
		"@adjust=" + data.LEARN_ADJUST_EXPECTED + ""

	err = db.RawSQL(query).Do(&datass)

	if err != nil {
		errsql := err
		err = db.Rollback() //Rollback Transaction
		if err != nil {
			return Rdata, err
		}
		return Rdata, errsql
	}

	qsum = " SELECT sum(ISNULL(LEARN_EXPECTED, 0) + ISNULL(LEARN_ADJUST_EXPECTED, 0)) as total, " +
		"sum(ISNULL(LEARN_EXPECTED, 0)) as totalexpected, sum(ISNULL(LEARN_ADJUST_EXPECTED, 0)) as totaladjust " +
		"FROM m_LEARN_EXPECTED_DEV where TRX_DATE = '" + data.TRX_DATE + "' and KODE_WILAYAH is not null and KODE_UNIT = '" + data.KODE_UNIT + "' " +
		"" + data.SEGMENT + " " +
		"" + data.IMP_STAGE + " "

	err = db.RawSQL(qsum).Do(&sumdata)

	if err != nil {
		return Rdata, err
	}

	err = db.Commit() //Commit Transaction
	if err != nil {
		return Rdata, err
	}

	jsondata, err := json.Marshal(&datass)
	if err != nil {
		return Rdata, err
	}

	Rdata = global.PagingResult{SUM: sumdata.TOTAL, SUMEXPECTED: sumdata.TOTALEXPECTED, SUMADJUST: sumdata.TOTALADJUST, DATA: jsondata}

	return Rdata, nil
}

type MASTERSEGMENTGET struct {
	SEGMENT interface{} `json:"SEGMENT" db:"SEGMENT"`
}

func GETMASTERSEGMENT() ([]MASTERSEGMENTGET, error) {
	var data []MASTERSEGMENTGET
	db, err := global.ConnSimWas()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := "select distinct SEGMENT from m_LEARN_EXPECTED_DEV where KODE_WILAYAH is not null order by SEGMENT asc"

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var each = MASTERSEGMENTGET{}
		err = rows.Scan(&each.SEGMENT)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, each)
	}

	return data, nil
}
