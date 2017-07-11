package date

import (
	"fmt"
)

const (
	YYYYMMDDHHmmSS = "2006-01-02 15:04:05"
)

func Formater(formater string) string {
	r := ""
	for _, v := range formater {
		s := fmt.Sprintf("%c", v)
		switch s {
		case "y", "Y":
			r += s

		case "-", " ":
			r += s
		}
	}
	return r
}
