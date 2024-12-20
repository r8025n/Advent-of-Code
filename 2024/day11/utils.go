package main


import(
	// "fmt"
	"strconv"
)
func removeLeadingZeros(str string) string {
	leadingZeros := 0

	for _, char := range str {
		if char != '0' {
			break
		}
		leadingZeros++
	}

	if leadingZeros != len(str) {
		return str[leadingZeros:]
	}
	return "0"
}

func transformNumber(input string) (output1, output2 string) {
	output2 = ""

	if input == "0" {
		output1 = "1"
	}else if len(input) % 2 == 0 {
		output1 = removeLeadingZeros(input[:len(input)/2])
		output2 = removeLeadingZeros(input[len(input)/2:])

	} else{
		stoneVal, _ := strconv.Atoi(input)
		stoneVal = stoneVal * 2024;
		output1 = strconv.Itoa(stoneVal)
	}

	return output1, output2
}
