package iteration

import "strings"

const repeatCount = 5

// Repeat takes a character and repeats it 5 times
// func Repeat(character string) string {
// 	var repeated string
/*
	Strings in Go are immutable, meaning every concatenation, such as in our Repeat function,
	involves copying memory to accommodate the new string. This impacts performance, particularly during heavy string concatenation.
*/
// 	for i := 0; i < repeatCount; i++ {
// 		repeated += character
// 	}
// 	return repeated
// }

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
