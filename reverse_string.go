package main

func main() {
	println(ReverseString("Hello World"))
}

func ReverseString(input string) (output string) {
	for _, v := range input {
		output = string(v) + output
	}
	return output
}
