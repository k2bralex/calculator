package converter

import (
	"fmt"
	"regexp"
	"strings"
)

// Decode parse roman number to arabic. If validation successful.
// If not return error
func Decode(s string) (int64, error) {
	if !validRoman(s) {
		return 0, fmt.Errorf("ronam number not valid")
	}

	keys := []string{"CM", "CD", "XC", "XL", "IX", "IV", "I", "V", "X", "L", "C", "D", "M"}
	romanArabic := map[string]int64{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"CM": 900,
		"CD": 400,
		"XC": 90,
		"XL": 40,
		"IX": 9,
		"IV": 4,
	}

	var res int64 = 0
	for _, key := range keys {
		res = res + int64(strings.Count(s, key))*romanArabic[key]
		s = strings.ReplaceAll(s, key, "")
	}

	return res, nil
}

// validRoman use regexp to check roman number. If ok return true
func validRoman(s string) bool {
	pattern := `^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`
	matched, _ := regexp.MatchString(pattern, s)
	return matched
}

// Encode based on the fact that Roman numbers have a range encode arabic
// to roman.
func Encode(n int64) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("roman can't be zero or negative")
	}
	if n > 3999 {
		return "", fmt.Errorf("roman can't be greater than 3999")
	}

	res := ""
	for n > 0 {
		if n >= 1000 {
			m := n / 1000
			res += strings.Repeat("M", int(m))
			n -= 1000 * m
		}

		if n >= 900 {
			res += "CM"
			n -= 900
		}

		if n >= 500 && n < 900 {
			d1 := n % 500 / 100
			res += "D" + strings.Repeat("C", int(d1))
			n -= 500 + 100*d1
		}

		if n >= 400 {
			res += "CD"
			n -= 400
		}

		if n >= 100 && n < 400 {
			d2 := n / 100
			res += strings.Repeat("C", int(d2))
			n %= 100 * d2
		}

		if n >= 90 {
			res += "XC"
			n -= 90
		}

		if n >= 50 && n < 90 {
			l := n/10 - 5
			res += "L" + strings.Repeat("X", int(l))
			n %= 10
		}

		if n >= 40 {
			res += "XL"
			n -= 40
		}

		if n >= 10 && n <= 40 {
			x := n / 10
			res += strings.Repeat("X", int(x))
			n %= 10
		}

		if n == 9 {
			res += "IX"
			n -= n
		}

		if n >= 5 && n < 9 {
			v := n % 5
			res += "V" + strings.Repeat("I", int(v))
			n -= v + 5
		}

		if n == 4 {
			res += "IV"
			n -= n
		}
		if n > 0 && n < 4 {
			res += strings.Repeat("I", int(n))
			n -= n
		}
	}

	return res, nil
}
