package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"
var colorOrange = "\x1b[38;2;255;128;m"
var counter int
var color string

func StrColor(text, str string) {

	switch color := str; color {
	case "red":
		fmt.Println(string(colorRed), text, string(colorReset))
	case "green":
		fmt.Println(string(colorGreen), text, string(colorReset))
	case "yellow":
		fmt.Println(string(colorYellow), text, string(colorReset))
	case "blue":
		fmt.Println(string(colorBlue), text, string(colorReset))
	case "purple":
		fmt.Println(string(colorPurple), text, string(colorReset))
	case "white":
		fmt.Println(string(colorWhite), text, string(colorReset))
	case "cyan":
		fmt.Println(string(colorCyan), text, string(colorReset))
	case "orange":
		fmt.Println(string(colorOrange), text, string(colorReset))
	}
}

func SwitchColor(str string) string {
	switch color := str; color {
	case "red":
		return string(colorRed)
	case "green":
		return string(colorGreen)
	case "yellow":
		return string(colorYellow)
	case "blue":
		return string(colorBlue)
	case "purple":
		return string(colorPurple)
	case "white":
		return string(colorWhite)
	case "cyan":
		return string(colorCyan)
	case "orange":
		return string(colorOrange)
	}
	return "\033[0m"
}

func ReadFiles(str []string) []string {
	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return str
}

func CheckIndex(str string, index int, soz int) bool {
	// fmt.Println(str, len(str))
	if str[:8] == "--index[" {

		array := strings.Split(str[8:len(str)-1], ":")
		// fmt.Println(array[0], array[1])
		if len(array) == 2 {
			var num1, num2 int
			var val1, val2 error
			if array[0] == "" && array[1] == "" {
				// fmt.Println("I m here")
				num1 = 0
				num2 = soz
			} else if array[1] == "" {
				num1, val1 = strconv.Atoi(array[0])
				if val1 != nil {
					return false
				}
				num2 = soz
			} else if array[0] == "" {
				num1 = 0
				num2, val2 = strconv.Atoi(array[1])
				if val1 != nil {
					return false
				}
			} else {
				num1, val1 = strconv.Atoi(array[0])
				num2, val2 = strconv.Atoi(array[1])
				if val1 != nil || val2 != nil {
					return false
				}
			}
			if index >= num1 && index <= num2 {
				return true
			}
		} else {
			num, val := strconv.Atoi(str[8 : len(str)-1])
			if val != nil {
				return false
			}
			if num == index {
				return true
			}
		}
	}
	return false
}

func CheckLetter(argument []string, letter string, letter_index int, soz int) bool {
	for index := 0; index < len(argument); index++ {
		if len(argument[index]) != 1 && CheckIndex(argument[index], letter_index, soz) {
			return true
		}
		if len(argument[index]) == 1 {
			for _, val := range argument[index] {
				if letter == string(val) {
					return true
				}
			}
		}
	}
	return false
}

func StrByLines(arr [][8]string, length int, color string, argumets []string, word string, lenarg int, soz string) {
	for index1 := 0; index1 < 8; index1++ {
		str := ""
		for index2 := 0; index2 < length; index2++ {
			for _, val := range arr[index2][index1] {
				if val != '\n' && index2 != length-1 {
					if CheckLetter(argumets, word[index2:index2+1], index2, len(soz)) {
						str = str + (SwitchColor(color) + string(val) + "\033[0m")
					} else {
						str = str + string(val)
					}
				} else if index2 == length-1 {
					if CheckLetter(argumets, word[index2:index2+1], index2, len(soz)) {
						str = str + (SwitchColor(color) + string(val) + "\033[0m")
					} else {
						str = str + string(val)
					}
				}
			}
		}
		if lenarg == 2 {
			StrColor(str, color)
		} else {
			fmt.Println(str)
		}
	}
}

func CheckSecondArg(str string) string {
	if str[0:8] == "--color=" {
		return str[8:]
	}
	return str
}

func Ascii(number int, str []string, index int) string {
	start := number*9 + 2 + index - 1
	nothing := ""
	for index, val := range str {
		if index == start {
			return val
		}
	}
	return nothing
}

func main() {
	arguments := os.Args[1:]
	if len(arguments) < 2 {
		return
	}
	color = color + CheckSecondArg(arguments[1])

	sozder := strings.Split(arguments[0], "\\n")
	for index := 0; index < len(sozder); index++ {
		var standard []string
		if len(sozder[index]) > 0 {
			standard = ReadFiles(standard)
		}
		result := make([][8]string, len(sozder[index]))
		for index1, val := range sozder[index] {
			for index2 := 0; index2 < 8; index2++ {
				result[index1][index2] = Ascii(int(val-32), standard, index2)
			}
		}
		StrByLines(result, len(sozder[index]), color, arguments[2:], sozder[index], len(arguments), sozder[index])
	}
	return
}