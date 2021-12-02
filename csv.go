package main

import (
	"encoding/csv"
	"io"
	"os"
)

func ReadFromCsv(fileName, dest string, comma rune, header bool) (e error) {
	f, e := os.Open(fileName)
	if e != nil {
		return
	}

	r := csv.NewReader(f)
	if comma > 0 {
		r.Comma = comma
	}

	writer := CreateExcelWriter(dest)
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			e = err
			return
		}

		var val []interface{}
		for index := range row {
			val = append(val, row[index])
		}

		if header && writer.header == nil {
			writer.SetHeader(val)
			continue
		}

		if e = writer.WriteRow(val); e != nil {
			return
		}
	}

	e = writer.Close()
	return
}
