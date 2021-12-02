package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	m *string
	h *bool
	H *bool
	i *string
	o *string
	c *string
	s *string
)

func init() {
	m = flag.String("m", "", "the import mode<csv, sql>")
	h = flag.Bool("h", false, "show help")
	i = flag.String("i", "", "input file path")
	o = flag.String("o", "", "output file name,default: huge.xlsx")
	c = flag.String("c", "", "the field delimiter(in csv mode only),default: ,")
	s = flag.String("s", "", "connection string in sql mode,like: username:password@tcp(ip:port)/database?charset=utf8mb4")
	H = flag.Bool("H", false, "set first line of csv file as header in every sheet")
}

func main() {
	flag.Parse()

	if *h {
		flag.Usage()
		return
	}

	if len(*o) == 0 {
		*o = "huge.xlsx"
	}

	if !strings.HasSuffix(*o, ".xlsx") {
		*o += ".xlsx"
	}

	comma := ','
	for _, c := range *c {
		comma = c
		break
	}

	var e error
	switch *m {
	case "csv":
		e = ReadFromCsv(*i, *o, comma, *H)
	case "sql":
		e = ReadFromDatabase(*s, *i, *o)
	default:
		fmt.Println("unsupported import mode")
	}

	if e != nil {
		fmt.Println("error: " + e.Error())
	}
}
