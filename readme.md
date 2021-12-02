a tool to fast transform data from csv or export data from mysql to excel file.

command option
```
  -H    set first line of csv file as header in every sheet
  -c string
        the field delimiter(in csv mode only),default: ,
  -h    show help
  -i string
        input file path
  -m string
        the import mode<csv, sql>
  -o string
        output file name,default: huge.xlsx
  -s string
        connection string in sql mode,like: username:password@tcp(ip:port)/database?charset=utf8mb4
```

export from mysql
```bash
huge-xlsx -m sql -i query.sql -o sql.xlsx -s 'username:password@tcp(ip:port)/database?charset=utf8mb4'
```

transform from csv file
```bash
huge-xlsx -m csv -i test.csv -o test.xlsx -H 
```