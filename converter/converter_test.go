package converter

import "testing"

func TestToWord(t *testing.T) {
	tests := []struct {
		description string
		input       int
		output      string
	}{
		{
			description: "should successfully return the english equivalent",
			input:       0,
			output:      "Zero",
		},
		{
			description: "should successfully return the english equivalent",
			input:       13,
			output:      "thirteen",
		},
		{
			description: "should successfully return the english equivalent",
			input:       100,
			output:      "one hundred",
		},
		{
			description: "should successfully return the english equivalent",
			input:       1771,
			output:      "one thousand seven hundred and seventy one",
		},
		{
			description: "should successfully return the english equivalent",
			input:       62626,
			output:      "sixty two thousand six hundred and twenty six",
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			if got := ToWord(test.input); got != test.output {
				t.Fatalf("Got %v but expected %v", got, test.output)
			}
		})
	}
}
