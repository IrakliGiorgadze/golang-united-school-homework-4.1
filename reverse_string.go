package reverse_string

func ReverseString(input string) (output string) {
	for _, v := range input {
		output = string(v) + output
	}
	return output
}
