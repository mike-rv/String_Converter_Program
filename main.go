package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("sample.txt") // read file & store as slice of bytes
	if err != nil {                               // if empty print nil/ blank
		fmt.Println("Err")
	}
	var slice []string
	stringSlice := string(content)      // convert byte slice to string
	slice = strings.Fields(stringSlice) // convert string to slice of string elements
	boolSwitch := true                  // declare boolSwitch for boolean flagging

	for i := 0; i < len(slice); i++ { // for loop over string slice

		if slice[i] == "(cap," { // find "(cap" in slice of string elements
			num, _ := strconv.Atoi(strings.Trim(slice[i+1], ")")) // 1) remove ")" from "6)" 2) convert "6" to int
			for j := len(slice) - 1; j >= 0; j-- {                // loop backwards through whole slice
				if j < i && j > i-num-1 { // if the iterator is less than the index of "cap,"
					slice[j] = capit(slice[j])
				}
			}
			remove(slice, i)
			remove(slice, i)
			slice = slice[:len(slice)-2]
		}
		if slice[i] == "(low," {
			num, _ := strconv.Atoi(strings.Trim(slice[i+1], ")"))
			for j := len(slice) - 1; j >= 0; j-- {
				if j < i && j > i-num-1 {
					slice[j] = low(slice[j])
				}
			}
			remove(slice, i)
			remove(slice, i)
			slice = slice[:len(slice)-2]
		}
		if slice[i] == "(up," {
			num, _ := strconv.Atoi(strings.Trim(slice[i+1], ")"))
			for j := len(slice) - 1; j >= 0; j-- {
				if j < i && j > i-num-1 {
					slice[j] = up(slice[j])
				}
			}
			remove(slice, i)
			remove(slice, i)
			slice = slice[:len(slice)-2]
		}
		if slice[i] == "(cap)" { // find "cap" in sample.txt
			slice[i-1] = capit(slice[i-1]) // capitalize word before "(cap)", it -> It
			remove(slice, i)
			slice = slice[:len(slice)-1]
		}
		if slice[i] == "(hex)" {
			slice[i-1] = hex(slice[i-1])
			remove(slice, i)
			slice = slice[:len(slice)-1]
		}
		if slice[i] == "(bin)" {
			slice[i-1] = bin(slice[i-1])
			remove(slice, i)
			slice = slice[:len(slice)-1]
		}
		if slice[i] == "(up)" {
			slice[i-1] = up(slice[i-1])
			remove(slice, i)
			slice = slice[:len(slice)-1]
		}
		if slice[i] == "(low)" {
			slice[i-1] = low(slice[i-1])
			remove(slice, i)
			slice = slice[:len(slice)-1]
		}
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "a" && slice[i+1][0] == 'o' {
			slice[i] = "an"
		}
		if slice[i] == "A" {
			slice[i] = "An"
		}
	}
	for i := 0; i < len(slice); i++ {
		if slice[i][0] == ',' && len(string(slice[i])) > 1 {
			slice[i] = strings.ReplaceAll(string(slice[i]), ",", "")
			slice[i-1] = slice[i-1] + ","
		} else if slice[i][0] == ',' && len(string(slice[i])) == 1 {
			slice[i-1] = slice[i-1] + ","
			remove(slice, i)
			slice = slice[:len(slice)-1]
		}
		if slice[len(slice)-1][0] == '?' {
			slice[len(slice)-2] = slice[len(slice)-2] + slice[len(slice)-1]
			slice = slice[:len(slice)-1]
		}
		if slice[i][0] == '?' && len(string(slice[i])) > 1 {
			slice[i] = strings.ReplaceAll(string(slice[i]), "?", "")
			slice[i-1] = slice[i-1] + "?"
		} else if slice[i][0] == '?' && len(string(slice[i])) == 1 {
			slice[i-1] = slice[i-1] + "?"
			remove(slice, i)
			slice = slice[:len(slice)-1]
		}
		if slice[i][0] == '.' && len(string(slice[i])) > 1 {
			slice[i] = strings.ReplaceAll(string(slice[i]), ".", "")
			slice[i-1] = slice[i-1] + "."
		} else if slice[i][0] == '.' && len(string(slice[i])) == 1 {
			slice[i-1] = slice[i-1] + "."
			remove(slice, i)
			slice = slice[:len(slice)-1]
		}
		if slice[i][0] == '!' && len(string(slice[i])) > 1 {
			slice[i] = strings.ReplaceAll(string(slice[i]), "!", "")
			slice[i-1] = slice[i-1] + "!"
		} else if slice[i][0] == '!' && len(string(slice[i])) == 1 {
			slice[i-1] = slice[i-1] + "!"
			remove(slice, i)
			slice = slice[:len(slice)-1]
		}
		if slice[i][0] == ':' && len(string(slice[i])) > 1 {
			slice[i] = strings.ReplaceAll(string(slice[i]), ":", "")
			slice[i-1] = slice[i-1] + ":"
		} else if slice[i][0] == ':' && len(string(slice[i])) == 1 {
			slice[i-1] = slice[i-1] + ":"
			remove(slice, i)
			slice = slice[:len(slice)-1]
		}
	}
	for i := 0; i < len(slice); i++ {
		if slice[i][0] == '\'' && i < len(slice)-1 {
			if boolSwitch { // boolSwitch = true / on the first instance of "'", execute following code
				slice[i] = strings.Join(slice[i:i+2], "")
				remove(slice, i+1)
				slice = slice[:len(slice)-1] // remove duplicate
				boolSwitch = false
			} else {
				slice[i-1] = slice[i-1] + slice[i]
				remove(slice, i)
				slice = slice[:len(slice)-1]
				boolSwitch = true
			}
		}
	}
	if slice[len(slice)-1][0] == '\'' {
		slice[len(slice)-2] = slice[len(slice)-2] + slice[len(slice)-1]
		slice = slice[:len(slice)-1]
	}
	result := ""
	for _, stringSlice := range slice {
		if stringSlice != "" {
			result += stringSlice + " " // result = result + slice[i] + " "
		}
	}
	result = strings.TrimRight(result, " ")
	s := []byte(result) // convert string to byte slice
	ioutil.WriteFile("result.txt", s, 0644)
}
func capit(s string) string {
	return strings.Title(s)
}
func hex(hex_num string) string {
	num, err := strconv.ParseInt(hex_num, 16, 64)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(num)
}
func bin(bin_num string) string {
	num, err := strconv.ParseInt(bin_num, 2, 64)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(num)
}
func low(s string) string {
	return strings.ToLower(s)
}
func up(s string) string {
	return strings.ToUpper(s)
}
func remove(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
