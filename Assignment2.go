package main
// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm


import (
	"fmt"
	"unicode"
)

func main() {
	NewArr := make([]int,10)
	var count int = 0
	for i := 0; i < len(NewArr); i++ {
		count++ 
		NewArr[i] = count 
	}
	NewArr = append(NewArr, 10)
	
	CreateMap := map[string]int{
		"first" : 1,
		"second" : 2, 
		"third" : 3, 
		"fourth" : 4, 
		"fifth" : 5,
	}
	// Feel free to use the main function for testing your functions
	//fmt.Println("Hello, Ø¯Ù†ÙŠØ§!")
	//fmt.Println(TopCharacters("HelloWorld",2))
	//fmt.Println(ParsePhone("123-456-7890"))
	fmt.Println(Anagram("moona", "noomr"))
	fmt.Println(FindEvens(NewArr))
	fmt.Println(SliceProduct(NewArr))
	fmt.Println(Unique(NewArr))
	fmt.Println(InvertMap(CreateMap))
	fmt.Println(TopCharacters("Hello World", 2))
	fmt.Println(ParsePhone("123-456-7890"))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	var arr string;
	var count int = 0 
	arr = "("
	for i := 0; i < len(phone); i++ {
		if unicode.IsNumber(rune(phone[i])) {
			count++
			if count == 3 { 
				arr += string(phone[i]) + ") "
			} else if count == 6 { 
				arr += string(phone[i]) + "-"
			} else {
				arr += string(phone[i])
			}
		}
	}
	return arr
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if HelperMethodForAnagram(s1) == HelperMethodForAnagram(s2) {
		return true
	}
	return false;
}



func HelperMethodForAnagram(s1 string) string { 
	r := []rune(s1)
	for i := 0; i < len(r) - 1; i++ {
		for j := i + 1;  j < len(r); j++ {
		if r[i] > r[j] { 
			temp := r[j]
			r[j] = r[i]
			r[i] = temp
		}
		}
	}
	return string(r)
}
// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var arr []int 
	for i := 0; i < len(e); i++ { 
		if (e[i] % 2 == 0) { 
			arr = append(arr, e[i])
			fmt.Println(e[i])
		}
	}
	return arr
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	var Product int = 1
	for i := 0; i < len(e); i++ { 
		Product *= e[i]
	}
	return Product
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	var arr []int
	for i := 0; i < len(e); i++ {
		var count int = 0 
		for j := 0; j < len(e); j++ { 
			if (e[i] == e[j]) { 
				count++;
			}
		}
		if count <= 1 { 
			arr = append(arr,e[i])
		}
	}
	return arr
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
		NewMap := make(map[int]string, len(kv)) 
		for k, v := range kv { 
			NewMap[v] = k
		}
	return NewMap 
}



// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(s); i++ { 
		var count int = 0 
		for j := 0; j < len(s); j++ {
			if s[i] == s[j] { 
				count++
			}
		}
		if count > k {
			m[string(s[i])] = count 
		}
	}

	// TODO
	return m
}
