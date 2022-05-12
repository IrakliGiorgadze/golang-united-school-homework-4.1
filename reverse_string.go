package reverse_string

import "fmt"

func ReverseString(input string) (output string) {
	for _, v := range input {
		output = string(v) + output
	}
	fmt.Println(output)
	return output
}
