# compiler_info

Overall, the program is designed to search for specific patterns within files in a given directory and provide feedback on the presence of those patterns.
My usecase for this is to iterate over (on large-scale) of malware samples to find out in which language they are written.

The provided Golang code is simple and easy to extend to your specific needs.




Example: you can add the string pattern for the UPX packer  

```
	patterns := map[string][]string{
		"upx":  {"UPX!"},
	}
```

## Note
*It's my first golang program so of course this code can be improved (e.g using go routines, buffio) and they are better ways todo it overall.*

- the search is not bullet proof
- they are ways to obfuscate a binary - where this string search approach will not work
- also other tools exists like capa (from Mandiant) which have better results overall


## Usage

Tested with Go v1.22 on Kali Linux and Arch Linux.
The os.Args variable is used to retrieve command-line arguments, specifically the directory path.

Simple run:
`go run main.go directory_to_scan`

OR

Build the program with:
`go build -ldflags "-s -w"`
and execute it: `./compiler_info directory_to_scan`

