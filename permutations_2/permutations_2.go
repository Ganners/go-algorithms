package main

import "fmt"

func findPerms(letters string) []string {

	return findPermsHelper(letters, 0)
}

// Complexity is O(e*n!), or O(n!)
func findPermsHelper(letters string, index int) []string {

	if letters == "" {
		return []string{""}
	}

	permutations := []string{}

	for i := index; i < len(letters); i++ {
		fmt.Println(index*10 + i)

		// Swap the letters (easy as a rune slice) from index to i
		temp := []rune(letters)
		temp[index], temp[i] = temp[i], temp[index]

		if index == len(letters)-1 {
			permutations = append(permutations, string(temp))
		} else {
			permutations = append(
				permutations,
				findPermsHelper(string(temp), index+1)...)
		}
	}

	return permutations
}

func main() {
	fmt.Println(findPerms("abcdefghij"))
}
