package helpers

import (
	"strings"
	"time"
)

type ConversionModel struct {
	Date time.Time
}

type ConversionResult struct {
	Date  int
	Day   string
	Month string
	Year  int
}

type DateConversion interface {
	Translate() ConversionResult
}

var indonesiaMonth = map[string]string{
	"january":   "Januari",
	"february":  "Februari",
	"march":     "Maret",
	"april":     "April",
	"may":       "Mei",
	"june":      "Juni",
	"july":      "Juli",
	"august":    "Agustus",
	"september": "September",
	"october":   "Oktober",
	"november":  "November",
	"december":  "Desember",
}

var indonesiaDay = map[string]string{
	"monday":    "Senin",
	"tuesday":   "Selasa",
	"wednesday": "Rabu",
	"thursday":  "Kamis",
	"friday":    "Jumat",
	"saturday":  "Sabtu",
	"sunday":    "Minggu",
}

func (m *ConversionModel) Translate() ConversionResult {
	year, rmonth, date := m.Date.Date()
	month := strings.ToLower(rmonth.String())
	day := strings.ToLower(m.Date.Weekday().String())
	res := ConversionResult{
		Year:  year,
		Month: indonesiaMonth[month],
		Date:  date,
		Day:   indonesiaDay[day],
	}

	return res
}
