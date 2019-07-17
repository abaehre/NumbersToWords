package flag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		description string
		intSlice    []int
		outputFile  string
	}{
		{
			description: "should successfully return with only one number",
			intSlice:    []int{282820},
			outputFile:  "test_data/oneValue.txt",
		},
		{
			description: "should successfully return with multiple numbers",
			intSlice:    []int{0, 13, 85, 5237},
			outputFile:  "test_data/multipleValues.txt",
		},
		{
			description: "should succesfully return with 0 numbers",
			intSlice:    []int{},
			outputFile:  "test_data/noValues.txt",
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			i := IntSlice(test.intSlice)
			got, err := json.Marshal(i.String())
			if err != nil {
				t.Fatalf("Errored while marshalling with error %v", err)
			}
			if want, err := loadFile(test.outputFile); err != nil {
				t.Fatalf("LoadFile failed retrieving outputFile %v", test.outputFile)
			} else if !reflect.DeepEqual(got, want) {
				t.Fatalf("\nExpected %v, but received %v", string(want), string(got))
			}
		})
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		description      string
		intSlice         []int
		setVal           string
		expectedIntSlice []int
		expectErr        bool
	}{
		{
			description:      "should successfully append a valid int",
			intSlice:         []int{},
			setVal:           "1",
			expectedIntSlice: []int{1},
		},
		{
			description: "should error out on non-number string",
			intSlice:    []int{},
			setVal:      "fail",
			expectErr:   true,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			i := IntSlice(test.intSlice)
			if err := i.Set(test.setVal); err != nil && !test.expectErr {
				t.Fatalf("Received err on Set - %v", err)
			} else if test.expectErr && err != nil {
				return
			}

			if !reflect.DeepEqual(i, IntSlice(test.expectedIntSlice)) {
				t.Fatalf("Expected %v, but received %v", test.expectedIntSlice, i)
			}
		})
	}
}

func loadFile(fileName string) ([]byte, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(fmt.Sprintf("%s/%s", pwd, fileName))
}
