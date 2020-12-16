package recursion

//Factorial
//Given n of 1 or more, return the factorial of n, which is n * (n-1) * (n-2) ... 1.
//ex:
// factorial(1) -> 1
// factorial(2) -> 2
// factorial(3) -> 6
func Factorial(n int) int {
	// base case: n = 0

	if n == 0 {
		return 1
	}

	return Factorial(n-1) * n
}

//BunnyEars computes the total number of ears across all the bunnies
//ex:
// BunnyEars(0) -> 0
// BunnyEars(2) -> 4
func BunnyEars(bunnies int) int {
	// base case: bunnies = 0
	if bunnies == 0 {
		return 0
	}

	return BunnyEars(bunnies-1) + 2
}

func SwapStr(str, sub string) string {
	if len(str) < len(sub) || len(str) == 1 {
		return str
	} else if str[0:len(sub)] == sub { // if the first letters math the target string
		str = "bingo" + (str[len(sub):])
		return SwapStr(str[0:1], sub) + SwapStr(str[1:], str)
	} else {
		return SwapStr(str[0:1], sub) + SwapStr(str[1:], sub)
	}
}

//CleanStr returns a string where adjacent characters that are the same
//have been reduced to a single character.
//ex:
// CleanStr("yyzzzza") -> "yza"
// CleanStr("hello") -> "helo"
func CleanStr(str string) string {
	if len(str) == 1 {
		return str
	} else if str[0] == str[1] {
		return CleanStr(str[0:1] + str[2:])
	} else {
		return CleanStr(str[0:1]) + CleanStr(string(str[1]))
	}

}

//NestParen returns true if it is a nesting of zero or more pairs of parenthesis
//ex:
// NestParen("((()))") -> true
// NestParen("(((x))") -> false
func NestParen(str string) bool {
	// base case: len(str) == 0

	if len(str) == 0 {
		return true
	} else if str[0] == '(' && str[len(str)-1] == ')' {
		return NestParen(str[1 : len(str)-1])
	} else {
		return false
	}

}
