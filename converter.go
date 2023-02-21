package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// romanToArabic parse roman number to arabic. If validation successful.
// If not return error
func romanToArabic(s string) (int, error) {
	if !validRoman(s) {
		return 0, fmt.Errorf("ronam number not valid")
	}

	keys := []string{"CM", "CD", "XC", "XL", "IX", "IV", "I", "V", "X", "L", "C", "D", "M"}
	romanArabic := map[string]int{
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

	res := 0
	for _, key := range keys {
		res = res + (strings.Count(s, key) * romanArabic[key])
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

// arabicToRoman based on the fact that Roman numbers have a range encode arabic
// to roman.
func arabicToRoman(n int) (string, error) {
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
			res += strings.Repeat("M", m)
			n -= 1000 * m
		}

		if n >= 900 {
			res += "CM"
			n -= 900
		}

		if n >= 500 && n < 900 {
			d1 := n % 500 / 100
			res += "D" + strings.Repeat("C", d1)
			n -= 500 + 100*d1
		}

		if n >= 400 {
			res += "CD"
			n -= 400
		}

		if n >= 100 && n < 400 {
			d2 := n / 100
			res += strings.Repeat("C", d2)
			n %= 100 * d2
		}

		if n >= 90 {
			res += "XC"
			n -= 90
		}

		if n >= 50 && n < 90 {
			l := n/10 - 5
			res += "L" + strings.Repeat("X", l)
			n %= 10
		}

		if n >= 40 {
			res += "XL"
			n -= 40
		}

		if n >= 10 && n <= 40 {
			x := n / 10
			res += strings.Repeat("X", x)
			n %= 10
		}

		if n == 9 {
			res += "IX"
			n -= n
		}

		if n >= 5 && n < 9 {
			v := n % 5
			res += "V" + strings.Repeat("I", v)
			n -= v + 5
		}

		if n == 4 {
			res += "IV"
			n -= n
		}
		if n > 0 && n < 4 {
			res += strings.Repeat("I", n)
			n -= n
		}
	}

	return res, nil
}

// valueParse try parse int
func valueParse(s string) (int, error) {
	v, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("not integer value")
	}
	return int(v), err
}
