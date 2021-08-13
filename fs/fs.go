// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
// )

// func ReadFiles(str []string, name string) []string {
// 	filename := name + ".txt"
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		str = append(str, scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	return str
// }

// func StrByLines(arr [][8]string, length int) {
// 	for index1 := 0; index1 < 8; index1++ {
// 		str := ""
// 		for index2 := 0; index2 < length; index2++ {
// 			for _, val := range arr[index2][index1] {
// 				if val != '\n' && index2 != length-1 {
// 					str = str + string(val)
// 				} else if index2 == length-1 {
// 					str = str + string(val)
// 				}
// 			}
// 		}
// 		fmt.Println(str)
// 	}
// }

// func Ascii(number int, str []string, index int) string {
// 	start := number*9 + 2 + index - 1
// 	nothing := "Some Problems Here"
// 	for index, val := range str {
// 		if index == start {
// 			return val
// 		}
// 	}
// 	return nothing
// }

// func main() {
// 	arguments := os.Args[1:]
// 	sozder := strings.Split(arguments[0], "\\n")
// 	for index := 0; index < len(sozder); index++ {
// 		var standard []string
// 		if len(sozder[index]) > 0 && len(arguments) == 2 {
// 			standard = ReadFiles(standard, arguments[1])
// 		} else {
// 			return
// 		}
// 		result := make([][8]string, len(sozder[index]))
// 		for index1, val := range sozder[index] {
// 			for index2 := 0; index2 < 8; index2++ {
// 				result[index1][index2] = Ascii(int(val-32), standard, index2)
// 			}
// 		}
// 		StrByLines(result, len(sozder[index]))
// 	}
// 	return
// }

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:] // take from terminal arguments
	if len(arg) == 0 { // check on more or les arguments
		fmt.Println("No argument")
		os.Exit(1)
	} else if len(arg) > 2 {
		fmt.Println("Too many args")
		return
	}
	var fileName string
	if len(arg) == 1 {
		fileName = "standard.txt"
	} else {
		fileName = arg[1] + ".txt"
	}

	file, err := os.Open(fileName)
	if err != nil { // check on error
		fmt.Println(err)
	} else {
		file2, _ := ioutil.ReadAll(file)         // read content
		var arrayletters []string = []string{""} // creat array for separation letters
		var count int
		for i := 0; i < len(file2); i++ {
			if file2[i] == 10 { // we consider '\n'
				count++
			}
			arrayletters[len(arrayletters)-1] += string(file2[i]) // add rune in string of extreme array
			if count == 9 {                                       // in the end letter we create new string and counter reset to zero
				count = 0
				if i != len(file2)-1 {
					arrayletters = append(arrayletters, "")
				}
			}
		}

		for i := 0; i < len(arg[0]); i++ { // check content of argument
			if arg[0][i] < 32 || arg[0][i] > 126 {
				fmt.Println("Not correct symbols!")
				return
			}
			if len(arg[0]) != i+1 && arg[0][i] == 92 && arg[0][i+1] == 110 { // check on '\n'
				print2(arg[0][:i], arrayletters, 1) // start of run before '\n'
				arg[0] = arg[0][i+2:]               // cut string
				i = -1                              // cycle start with the beginning
			}
		}
		if len(arg[0]) != 0 { // if in array left letters
			print2(arg[0], arrayletters, 1)
		} else {
			fmt.Println()
		}
	}
	file.Close()
}

func print2(arg string, array []string, depth int) {
	if len(arg) == 0 { // if there is nothing before '\n'
		fmt.Println()
		return
	}
	if depth == 9 {
		return
	}
	elem(arg, array, 0, depth) // print string of letter depending on the depth
	fmt.Println()
	print2(arg, array, depth+1)
}

func elem(arg string, array []string, l int, depth int) {
	var count int
	if len(arg) == l {
		return
	}
	symbolIdx := arg[l] - 32 // index in array equals number in ascii if you take away '32'
	for asciiElIdx := 0; asciiElIdx < len(array[symbolIdx]); asciiElIdx++ {
		if array[symbolIdx][asciiElIdx] == 10 {
			count++ // counter need for find string
			continue
		}
		if count == depth { // when found necessary string, we it print
			fmt.Print(string(array[symbolIdx][asciiElIdx]))
		} else if depth < count {
			break
		}
	}
	count = 0
	elem(arg, array, l+1, depth)
}
