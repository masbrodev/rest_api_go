package models

import (
	"net/http"
	"rest_api/db"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

func FetchAllPegawai() (Response, error) {
	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StorePegawai(nama, alamat, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()
	sqlStatement := "INSERT pegawai (nama, alamat, telepon) VALUES (?,?,?)"
	sql, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := sql.Exec(nama, alamat, telepon)
	if err != nil {
		return res, err
	}

	lii, err := result.LastInsertId()
	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_intesert_id": lii,
	}

	return res, nil
}

func UpdatePegawai(id int, nama, alamat, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE pegawai SET nama = ?, alamat = ?, telepon, WHERE id = ?"
	sql, err := con.Prepare(sqlStatement)
	if sql != nil {
		return res, err
	}

	result, err := sql.Exec(nama, alamat, telepon, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
