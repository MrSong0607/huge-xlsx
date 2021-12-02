package main

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ReadFromDatabase(ConnectString string, fileName, dest string) (e error) {
	Db, e := gorm.Open(mysql.Open(ConnectString), &gorm.Config{})
	if e != nil {
		return
	}

	Db = Db.Debug()

	sqlFile, e := os.ReadFile(fileName)
	if e != nil {
		return
	}

	rows, e := Db.Raw(string(sqlFile)).Rows()
	if e != nil {
		return
	}

	cols, e := rows.Columns()
	if e != nil {
		return
	}

	writer := CreateExcelWriter(dest)
	var header []interface{}
	for _, v := range cols {
		header = append(header, v)
	}
	writer.SetHeader(header)

	for rows.Next() {
		r := make([]interface{}, len(cols))
		for i := range r {
			r[i] = &r[i]
		}

		if e = rows.Scan(r...); e != nil {
			return
		}

		if e = writer.WriteRow(r); e != nil {
			return
		}
	}

	e = writer.Close()
	return
}
