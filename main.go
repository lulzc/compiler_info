/*

Author: https://github.com/lulzc

ToDo:
- cache stdout / add progress bar
- write output to db

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func scanDir(filePath string) {

	files, err := os.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		compilerInfov2(filePath + file.Name())

	}

}

// using buffio NewReader instead of NewScanner
// NewScanner can be problematic for large files and adapted to size
func compilerInfov2(file string) error {
	// define patterns for different cases
	patterns := map[string][]string{
		"rust":  {"RUST_BACKTRACE=1", "Option::unwrap()", "Result::unwrap()"},
		"go":    {"Go build ID:", "go.buildid", "runtime.gcWork"},
		"zig":   {"ZIG_DEBUG_COLOR", "\\\\.\\pipe\\zig-childprocess-{d}-{d}"},
		"mingw": {"Mingw runtime failure:", "_Jv_RegisterClasses"},
	}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		for key, pat := range patterns {
			for _, p := range pat {
				if strings.Contains(line, p) {
					fmt.Printf("%s found pattern: %s\n", file, key)
				}
				break
			}
		}
	}

	return nil
}

func main() {

	start := time.Now()

	filePath := os.Args[1]

	scanDir(filePath)

	// run on single file
	//compilerInfov2(filePath)

	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)

}
