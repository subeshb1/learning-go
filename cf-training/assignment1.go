package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("madam"))  //true
	fmt.Println(isPalindrome("maam"))   //true
	fmt.Println(isPalindrome("masdam")) //false
}

func isPalindrome(value string) bool {
	var strLength = len(value)
	for i := 0; i < strLength/2; i++ {
		if value[i] != value[strLength-1-i] {
			return false
		}
	}
	return true
}
