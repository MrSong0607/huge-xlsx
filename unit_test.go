package main

import "testing"

func TestQuery(t *testing.T) {
	e := ReadFromDatabase("", "db.sql", "sql.xlsx")
	if e != nil {
		t.Error(e)
	}
}

func TestCsv(t *testing.T) {
	e := ReadFromCsv("station.csv", "station.xlsx", 0, true)
	if e != nil {
		t.Error(e)
	}
}
