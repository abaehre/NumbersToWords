package converter

import "strings"

// ToWord - converts a single integer number into its English equivalent
// note on lack of string builder - it doesn't have a prepend equivalent because of its implementation but should be of negligible impact on performance
func ToWord(num int) string {
	if num == 0 {
		return zero
	}
	ret := ""
	prefix := ""
	if num < 0 {
		num *= -1
		prefix = negative
	}

	placeInNum := 0
	for num > 0 {
		if tempNum := num % 1000; tempNum != 0 {
			ret = convertLessThanOneThousand(tempNum) + largeClassifiers[placeInNum] + ret
		}
		placeInNum++
		num /= 1000
	}
	return strings.TrimSpace(prefix + ret)
}

func convertLessThanOneThousand(num int) string {
	ret := ""
	if num%100 < 20 {
		ret = belowTwenty[num%100]
		num /= 100
	} else {
		ret = belowTwenty[num%10]
		num /= 10

		ret = classifiers[num%10] + ret
		num /= 10
	}
	if num == 0 {
		return ret
	}
	hundredString := " hundred and"
	if ret == "" {
		hundredString = " hundred"
	}
	return belowTwenty[num] + hundredString + ret
}
