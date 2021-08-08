package student

func AtoiBase(s string, base string) int {
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
	return result
}
