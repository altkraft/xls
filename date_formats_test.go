package xls

import (
	"reflect"
	"testing"
)

var (
	expectedCells = [][]string{
		[]string{"Format", "Date", "Date + time", "Time + date"},
		[]string{"YYYY-MM-DD", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
		[]string{"YYYY/MM/DD", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
		[]string{"YYYY.MM.DD", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
		[]string{"DD-MM-YYYY", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
		[]string{"DD/MM/YYYY", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
		[]string{"DD.MM.YYYY", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
		[]string{"MM-DD-YYYY", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
		[]string{"MM/DD/YYYY", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
		[]string{"MM.DD.YYYY", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z", "2006-01-02T15:04:05Z"},
	}
)

func TestDateFormats(t *testing.T) {
	if xlFile, err := Open("DateFormats.xls", "utf-8"); err == nil {
		if cells := xlFile.ReadAllCells(1024); !reflect.DeepEqual(cells, expectedCells) {
			t.Fail()
		}
	}
}
