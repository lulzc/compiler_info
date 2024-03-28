/*

Author: https://github.com/lulzc

ToDo:
- using bufio to speed up
-- cache stdout (fmt = slow)

*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func searchFile(filePath string, pattern string) (bool, error) {
	// read the file contents
	content, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	// convert the byte slice to a string
	fileContent := string(content)

	lines := strings.Split(fileContent, "\n")
	for _, line := range lines {
		// check if the line contains the pattern
		if strings.Contains(line, pattern) {
			return true, nil
		}
	}
	return false, nil
}

func iterate(path string, patterns map[string][]string) (bool, error) {
	var patternFound bool

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			for key, patternList := range patterns {
				for _, pattern := range patternList {
					found, err := searchFile(filePath, pattern)
					if err != nil {
						return err
					}
					if found {
						patternFound = true
						fmt.Printf("pattern %s found on %s \n", key, filePath)
						break
					}
				}
			}
		}
		return nil
	})

	return patternFound, err
}

func main() {

	start := time.Now()

	directoryPath := os.Args[1]

	// define patterns for different cases
	// todo: string[] review
	patterns := map[string][]string{
		"rust":  {"RUST_BACKTRACE=1", "Option::unwrap()", "Result::unwrap()"},
		"go":    {"Go build ID:", "go.buildid", "runtime.gcWork"},
		"zig":   {"ZIG_DEBUG_COLOR", "\\\\.\\pipe\\zig-childprocess-{d}-{d}"},
		"mingw": {"Mingw runtime failure:", "_Jv_RegisterClasses"},
	}

	patternFound, err := iterate(directoryPath, patterns)
	if err != nil {
		fmt.Printf("Error searching file for %s\n", err)
	}
	fmt.Println(patternFound) // for testing ... boolean value for patterns

	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)

}
