package number_parsing

import (
	"fmt"
	"strings"
)

// Max and min possible values constants
const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

// Dictionary with english numbers names
var Dict = map[string][]string{
	"one":  {"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"},
	"teen": {"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"},
	"ten":  {"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"},
	"big":  {"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion"},
}

// ConvertNumbers convert given int to human-readable english
func ConvertNumbers(number int) (output string) {
	// Hardcode checking MinInt
	if number == MinInt {
		return "negative nine quintillion two hundred twenty three quadrillion three hundred seventy two trillion thirty six billion eight hundred fifty four million seven hundred seventy five thousand eight hundred seven"
	}

	// Hardcode checking Zero
	if number == 0 {
		return Dict["one"][0]
	}

	// Strings array used for append parsed numbers
	var parsed []string

	// Check if number is negative
	if number < 0 {
		// Append "negative" to the start of array
		parsed = append(parsed, "negative")
		number *= -1
	}

	// Dividing input number by three digits segments
	var segments []int
	for number > 0 {
		segments = append(segments, number%1000)
		number = number / 1000
	}

	fmt.Println(segments)

	// Get english worlds from segments
	for idx := len(segments) - 1; idx >= 0; idx-- {
		segment := segments[idx]

		fmt.Println(idx)
		fmt.Println(segment)

		// Pass empty segment
		if segment == 0 {
			continue
		}

		// Divide segment for numeral unit
		hundreds := segment / 100 % 10
		tens := segment / 10 % 10
		units := segment % 10
		if hundreds > 0 {
			parsed = append(parsed, Dict["one"][hundreds], "hundred")
		}

		// If segment has only hundreds values - go to next iteration
		if tens == 0 && units == 0 {
			goto segmentsEnd
		}

		// Check if 'tens' unit more than 20
		switch tens {
		case 0:
			parsed = append(parsed, Dict["one"][units])
		case 1:
			parsed = append(parsed, Dict["teen"][units])
			break
		default:
			// Check if ten not like 20, 30, ..., 90
			if units > 0 {
				num := fmt.Sprintf("%s %s", Dict["ten"][tens], Dict["one"][units])
				parsed = append(parsed, num)
			} else {
				parsed = append(parsed, Dict["ten"][tens])
			}
			break
		}

		// Add 'big' suffix to segment
	segmentsEnd:
		if big := Dict["big"][idx]; big != "" {
			parsed = append(parsed, big+",")
		}
	}

	// Get string from array
	output = strings.Join(parsed, " ")
	// Get rid of the last last ','
	output = strings.TrimSuffix(output, ",")
	return
}
