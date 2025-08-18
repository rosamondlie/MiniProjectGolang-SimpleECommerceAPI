package utils

import (
	"fmt"
	"time"
)

func CekStatusUser(status bool) (bool, string) {
	switch status {
	case false:
		return false, "Tidak Aktif"
	case true:
		return true, "Aktif"
	default:
		return false, "Aktif"
	}
}

func FormatTanggal(t time.Time) string {
	var bulanIndonesia = map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}

	day := t.Day()
	month := bulanIndonesia[t.Format("January")]
	year := t.Year()
	return fmt.Sprintf("%d %s %d", day, month, year)
}

