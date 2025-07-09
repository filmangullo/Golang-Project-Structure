package DatetimeFunctions

import (
	"strings"
	"time"
)

// Fungsi utama: input date + php-style format
func FormatDateStyle(inputDate string, phpFormat string) (string, error) {
	inputLayout := "2006-01-02 15:04:05" // asumsi inputnya begini
	t, err := time.Parse(inputLayout, inputDate)
	if err != nil {
		return "", err
	}
	layout := convertLayout(phpFormat)
	return t.Format(layout), nil
}

func convertLayout(phpFormat string) string {
	replacer := strings.NewReplacer(
		"Y", "2006", // Tahun 4 digit
		"y", "06", // Tahun 2 digit
		"m", "01", // Bulan 2 digit
		"n", "1", // Bulan tanpa nol
		"d", "02", // Hari 2 digit
		"j", "2", // Hari tanpa nol
		"H", "15", // Jam 24 jam
		"h", "03", // Jam 12 jam
		"i", "04", // Menit
		"s", "05", // Detik
		"A", "PM", // AM/PM
		"a", "pm", // am/pm
	)
	return replacer.Replace(phpFormat)
}
