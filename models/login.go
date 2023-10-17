package models

//"learn_go/global"
//"fmt"
//"encoding/json"
//"encoding/base64"
//"crypto/sha1"
//"crypto/md5"

type LoginUser struct {
	NAMA          string `json:"nama"`
	USERNAME      string `json:"username"`
	USER_ID       string `json:"user_id"`
	POSITION_NAME string `json:"position_name"`
	//ROLE string `json:"role"`
	NIK          string `json:"nik"`
	LOKASI_KERJA string `json:"lokasi_kerja"`
}

type LoginTrue struct {
	RESPONSE bool      `json:"response"`
	MESSAGE  string    `json:"message"`
	TOKEN    string    `json:"token"`
	DATA     LoginUser `json:"data"`
}

type LoginFalse struct {
	RESPONSE bool   `json:"response"`
	MESSAGE  string `json:"message"`
}

type LoginMKRIntegrasi struct {
	IDUser    string      `db:"id_user_profile" json:"id_user_profile"`
	NamaDepan string      `db:"nama_depan" json:"nama_depan"`
	Username  string      `db:"username" json:"username"`
	NIK       interface{} `db:"mkr_users_profile.nik" json:"mkr_users_profile.nik"`
	Nama      string      `db:"mkr_users_profile.nama" json:"mkr_users_profile.nama"`
	UnitKerja string      `db:"unit_kerja" json:"unit_kerja"`
}

type DataHRIS struct {
	KaryawanId string `db:"karyawan_id" json:"karyawan_id"`
	NIPHRIS    string `db:"nip_hris" json:"nip_hris"`
	NIP        string `db:"nip" json:"nip"`
	NamaHRIS   string `db:"nama_hris" json:"nama_hris"`
	Nama       string `db:"nama" json:"nama"`
	Inisial    string `db:"inisial" json:"inisial"`
	Posisi     string `db:"posisi" json:"posisi"`
	UnitKerja  string `db:"unit_kerja" json:"unit_kerja"`
	NomorInduk string `db:"nomor_induk" json:"nomor_induk"`
}

type CekUserPusat struct {
	Username string `db:"username" json:"username"`
	Nama     string `db:"nama" json:"nama"`
	Role     string `db:"role" json:"role"`
}

// func CheckUserLogin(Username string, password string) []LoginMKRIntegrasi{
// 	var data []LoginMKRIntegrasi
// 	db,err := global.ConnMKRIntegrasi()
// 	if err != nil {
// 	 	fmt.Println(err.Error())
// 	}
// 	defer db.Close()

// 	hasher := md5.New()
// 	hasher.Write([]byte(password))
// 	mdhash := hasher.Sum(nil)
// 	ss1 := fmt.Sprintf("%x", mdhash)

// 	sha := sha1.New()
// 	   sha.Write([]byte(ss1))
// 	encrypted := sha.Sum(nil)
// 	ss := fmt.Sprintf("%x", encrypted)

// 	var query string

// 	query = "SELECT  id_user_profile, nama_depan,username,mkr_users_profile.nik,mkr_users_profile.nama,unit_kerja "+
// 		"FROM mkr_users WITH (NOLOCK) "+
// 		"INNER JOIN mkr_users_profile ON mkr_users.id_user_profile = mkr_users_profile.id "+
// 		"WHERE username = '"+Username+"' AND password = '"+ss+"'"

// 	rows, err := db.Query(query)
// 	if err != nil {
// 	 	fmt.Println(err.Error())
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var each = LoginMKRIntegrasi{}
// 		err = rows.Scan(&each.IDUser, &each.NamaDepan, &each.Username, &each.NIK, &each.Nama, &each.UnitKerja)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		data = append(data, each)
// 	}
// 	// db.Close()

// 	return data
// }

// func GetDataHris(karyawan_id string) []DataHRIS{
// 	var data []DataHRIS
// 	db,err := global.ConnHRIS()
// 	if err != nil {
// 	 	fmt.Println(err.Error())
// 	}
// 	defer db.Close()

// 	var query string

// 	query = "SELECT karyawan_id, nip_hris, nip, nama_hris, nama, inisial, posisi, unit_kerja, nomor_induk "+
// 		"FROM karyawan "+
// 		"WHERE karyawan_id = '"+karyawan_id+"'"

// 	rows, err := db.Query(query)
// 	if err != nil {
// 	 	fmt.Println(err.Error())
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 	 	var each = DataHRIS{}
// 		err = rows.Scan(&each.KaryawanId, &each.NIPHRIS, &each.NIP, &each.NamaHRIS, &each.Nama, &each.Inisial, &each.Posisi, &each.UnitKerja, &each.NomorInduk)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		data = append(data, each)
// 	}

// 	fmt.Println(data)
// 	// db.Close()

// 	return data
// }

// func CheckDataPusat(username string) []CekUserPusat{
// 	var data []CekUserPusat
// 	db,err := global.ConnHRIS()
// 	if err != nil {
// 	 	fmt.Println(err.Error())
// 	}
// 	defer db.Close()

// 	var query string

// 	query = "SELECT DISTINCT   username, name as nama, keterangan as role "+
// 			"FROM Tb_UserPusat a WITH(NOLOCK) INNER JOIN Tb_UserPusat_Role b ON a.role = b.id "+
// 			"WHERE username = '"+username+"'"

// 	rows, err := db.Query(query)
// 	if err != nil {
// 	 	fmt.Println(err.Error())
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var each = CekUserPusat{}
// 		err = rows.Scan(&each.Username, &each.Nama, &each.Role)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		data = append(data, each)
// 	}

// 	return data
// }

type CekRole struct {
	Kode_role  string `db:"kode_role" json:"kode_role"`
	Keterangan string `db:"keterangan" json:"keterangan"`
	Roleid     string `db:"roleid" json:"roleid"`
}

// func CheckRole(jabatan string) []CekRole{
// 	var data []CekRole
// 	db,err := global.ConnSimWas()
// 	if err != nil {
// 	 	fmt.Println(err.Error())
// 	}
// 	defer db.Close()

// 	var query string

// 	query = "SELECT  kode_role,keterangan,roleid FROM TB_Role_User WHERE nama_role = '"+ jabatan +"' "+
// 			"or keterangan = '"+ jabatan +"' or kode_role = '"+ jabatan +"'"

// 	rows, err := db.Query(query)
// 	if err != nil {
// 	 	fmt.Println(err.Error())
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var each = CekRole{}
// 		err = rows.Scan(&each.Kode_role, &each.Keterangan, &each.Roleid)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		data = append(data, each)
// 	}

// 	return data
// }
