package fuzzer

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/fatih/color"
)

// Check : parse specs generated by fuzzer and check for ones that do not error
func TestFuzzedSpecs(t *testing.T) {
	red := color.New(color.Bold, color.FgRed).PrintlnFunc()
	yellow := color.New(color.Bold, color.FgYellow).PrintlnFunc()
	cyan := color.New(color.Bold, color.FgCyan).PrintlnFunc()

	files, err := ioutil.ReadDir("data")
	if err != nil {
		t.Error(err)
	}

	fmt.Println("=============================")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("    FAULT FUZZ TESTING")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	fmt.Printf("Parsing %d fuzzed spec files...\n\n", len(files))
	failed := 0

	for _, f := range files {
		data, err := ioutil.ReadFile(fmt.Sprintf("data/%s", f.Name()))
		if err != nil {
			t.Error(err)
		}
		spec := string(data)
		code, err := evaluate(spec)
		switch code {
		case 0:
			fmt.Print("\n\n")
			cyan(f.Name())
			fmt.Println("--------------------")
			fmt.Println(spec)
			yellow("No syntax error caught, inspect spec to confirm valid")
		case 1: // Go Panic detected
			failed++
			fmt.Print("\n\n")
			cyan(f.Name())
			fmt.Println("--------------------")
			fmt.Println(spec)
			red("Panic detected: ", err)

		}
	}
	fmt.Printf("\nTest Complete: %d out of %d files with confirmed errors", failed, len(files))

}