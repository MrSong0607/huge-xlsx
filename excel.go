package main

import (
	"strconv"
	"sync"

	"github.com/xuri/excelize/v2"
)

type ExcelWriter struct {
	fileName   string
	excelFile  *excelize.File
	mutex      sync.Mutex
	rowID      int
	sheetIndex int
	writer     *excelize.StreamWriter
	header     []interface{}
}

func CreateExcelWriter(fileName string) *ExcelWriter {
	return &ExcelWriter{fileName: fileName, excelFile: excelize.NewFile(), mutex: sync.Mutex{}, rowID: 1, sheetIndex: 1}
}

func (_e *ExcelWriter) SetHeader(v []interface{}) {
	_e.header = v
}

func (_e *ExcelWriter) WriteRow(val []interface{}) (e error) {
	_e.mutex.Lock()
	defer _e.mutex.Unlock()

	if _e.rowID > excelize.TotalRows {
		if e = _e.writer.Flush(); e != nil {
			return
		}

		_e.rowID, _e.writer = 1, nil
		_e.sheetIndex++
	}

	if _e.writer == nil {
		_e.excelFile.NewSheet(_e.currentSheetName())
		_e.writer, e = _e.excelFile.NewStreamWriter(_e.currentSheetName())
		if e != nil {
			return
		}
	}

	if _e.header != nil && _e.rowID == 1 {
		if e = _e.doSingle(_e.header); e != nil {
			return
		}
	}

	e = _e.doSingle(val)
	return
}

func (_e *ExcelWriter) doSingle(val []interface{}) (e error) {
	cell, _ := excelize.CoordinatesToCellName(1, _e.rowID)
	if e = _e.writer.SetRow(cell, val); e != nil {
		return
	}

	_e.rowID++
	return
}

func (_e *ExcelWriter) currentSheetName() string {
	return "Sheet" + strconv.Itoa(_e.sheetIndex)
}

func (_e *ExcelWriter) Close() (e error) {
	if _e.writer != nil {
		if e = _e.writer.Flush(); e != nil {
			return
		}
		_e.writer = nil
	}

	return _e.excelFile.SaveAs(_e.fileName)
}
