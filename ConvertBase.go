package student

import (
	"github.com/01-edu/z01"
)

func ConvertBase(nbr, baseFrom, baseTo string) string {
	num := AtoiBaseNew(nbr, baseFrom)
	str := PrintNbrBaseNew(num, baseTo)
	return str
}

func AtoiBaseNew(s string, base string) int {
	str := []rune{}
	result := 0
	for index, char1 := range base {
		if char1 == '-' || char1 == '+' {
			return 0
		}
		for _, char2 := range base[index+1:] {
			if char1 == char2 {
				return 0
			}
		}
		str = append(str, char1)
	}
	for j, val := range s {
		for i := 0; i < len(base); i++ {
			if val == str[i] {
				if j == len(s)-1 {
					result = result + i
				} else {
					result = (result + i) * len(base) // All rights all preserved by Adil Kenzhigaliyev!
				}
			}
		}
	}
	// fmt.Println(result)
	return result
}

func PrintNbrBaseNew(nbr int, base string) string {
	positive := true
	result := ""
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return result
	}
	if nbr < 0 {
		positive = false
		nbr = nbr * -1
	}
	for index, char1 := range base {
		if char1 == '-' || char1 == '+' {
			return result
		}
		for _, char2 := range base[index+1:] {
			if char1 == char2 {
				return result
			}
		}
	}
	for nbr >= len(base) {
		val := nbr % len(base)
		result = base[val:val+1] + result
		nbr = nbr / len(base)
	}
	// fmt.Println(result)

	result = base[nbr:nbr+1] + result
	if !positive {
		result = "-" + result
	}
	return result
}
