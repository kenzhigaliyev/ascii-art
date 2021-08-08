package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

func StrByLines(arr [][8]string, length int) {
	for index1 := 0; index1 < 8; index1++ {
		str := ""
		for index2 := 0; index2 < length; index2++ {
			for _, val := range arr[index2][index1] {
				if val != '\n' && index2 != length-1 {
					str = str + string(val)
				} else if index2 == length-1 {
					str = str + string(val)
				}
			}
		}
		fmt.Println(str)
	}
}

func Ascii(number int, str []string, index int) string {
	start := number*9 + 2 + index - 1
	nothing := "Some Problems Here"
	for index, val := range str {
		if index == start {
			return val
		}
	}
	return nothing
}

func main() {
	arguments := os.Args[1:]
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
		StrByLines(result, len(sozder[index]))
	}
	return
}
