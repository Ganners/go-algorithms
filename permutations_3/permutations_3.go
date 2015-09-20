package main

import "fmt"

func getPerms(str string) []string {

	return getPermsHelper("", str)
}

func getPermsHelper(prefix, remainder string) []string {

	result := make([]string, 0)

	if len(remainder) == 0 {
		result = append(result, prefix)
	}

	length := len(remainder)
	for i := 0; i < length; i++ {

		before := remainder[:i]
		after := remainder[i+1:]
		c := string(remainder[i])

		result = append(result, getPermsHelper(prefix+c, before+after)...)
	}

	return result
}

func main() {
	fmt.Println(getPerms("abcdefg"))
}
