package flag

import (
	"strconv"
	"strings"

	"github.com/abaehre/NumbersToWords/converter"
)

// IntSlice - implements flag interface to allow for multiple values to be passed in via command line
type IntSlice []int

// String - returns formatted string of int slice with conversions to English
func (i *IntSlice) String() string {
	builder := strings.Builder{}
	// fail out early if no values in the slice
	if len(*i) == 0 {
		builder.WriteString("No numbers were inputted!\n")
		return builder.String()
	}
	builder.WriteString("Input\tOutput\n--------------\n")
	for _, val := range *i {
		builder.WriteString(strconv.Itoa(val))
		builder.WriteString("\t")
		word := converter.ToWord(val)
		builder.WriteString(strings.ToUpper(word[:1]) + word[1:])
		builder.WriteString("\n")
	}
	return builder.String()
}

// Set - appends the string value to the IntSlice as an int
func (i *IntSlice) Set(value string) error {
	val, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*i = append(*i, val)
	return nil
}
