package challenges

// A string S consisting of N characters is considered to be properly nested
// if any of the following conditions is true:
// S is empty;
// S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
// S has the form "VW" where V and W are properly nested strings.
// For example, the string "{[()()]}" is properly nested but "([)()]" is not.
//
// Write a function:
//
// def solution(S string) bool
//
// that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.
//
// For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]",
// the function should return 0, as explained above.
//
// Write an efficient algorithm for the following assumptions:
//
// N is an integer within the range [0..200,000];
// string S is made only of the following characters: '(', '{', '[', ']', '}' and/or ')'.

var (
	openingBracketChars = [...]rune{'(', '{', '['}
	closingBracketChars = [...]rune{')', '}', ']'}
)

func SolutionBrackets(bracketStr string) bool {
	stack := []rune{}
	for _, v := range bracketStr {
		if v == openingBracketChars[0] || v == openingBracketChars[1] || v == openingBracketChars[2] {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			prev_v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (v == closingBracketChars[0] && prev_v != openingBracketChars[0]) ||
				(v == closingBracketChars[1] && prev_v != openingBracketChars[1]) ||
				(v == closingBracketChars[2] && prev_v != openingBracketChars[2]) {
				return false
			}
		}
	}

	return len(stack) == 0
}
