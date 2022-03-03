package main
import "fmt"





func FizzBuzz(i int) string {
	var something string = "";
	if i % 3 == 0 {
		something = "Fizz"
	} else if i % 5 == 0 {
		something = "Buzz"
	} else if i % 3 == 0 && i % 5 == 0 {
		something = "FizzBuzz"
	} else {
		return ""
	}
	
	return something
}

func isPrime(n int) bool {
	var flag bool = true 
	for i := 2; i <= n / 2; i++ {
		if n % i == 0 {
			flag = false
			break
		}
	} 

	return flag
} 

func isPalindrome(s string) bool { 
	var istrue bool = true
	for i := 2; i < len(s); i++ {
		if s[i] != s[len(s) - 1 - i] { 
			istrue = false
		}
	}
	return istrue
}



func main() {

	var i int = 100
	fmt.Println(FizzBuzz(i))
	fmt.Println(isPrime(30))
	fmt.Println(isPalindrome("moom"))


}
















