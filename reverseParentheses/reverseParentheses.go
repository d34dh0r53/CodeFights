package reverseParentheses

import (
	"fmt"
)

func reverseParentheses(s string) string {
	fmt.Println(s)
	return string(reverseBytes([]byte(s)))
}

func reverseBytes(in []byte) []byte {
	fmt.Println("In: ", in)
	for i := 0; i < len(in); i++ {
		r := in[i]
		if r == '(' {
			in = append(in[:i], reverseBytes(in[i+1:])...)
		}
		if r == ')' {
			simpleReverse(in[:i])
			return append(in[:i], in[i+1:]...)
		}
	}
	return in

}

func simpleReverse(in []byte) {
	fmt.Println("Simple Reverse In: ", in)
	i := len(in)
	for j := 0; j < i/2; j++ {
		x := in[j]
		in[j] = in[i-j-1]
		in[i-j-1] = x
	}
}
