# Go

This page is my reference to Go programming language 

### Installation:

Download from [here](https://golang.org/dl/)

### Configuration:

To work with Go-lang, we need to have a single workspace structure. Thus, create a folder following the below structure

```bash
gocode <This is the workspace folder>
	- bin
	- src
		- github.com
			- <your github username, mine is biswas-n>
				- project_1
				- project_2
	- pkg
```

*Note: Later we need to add the **GOCODE** workspace folder to **GOPATH** variable using 
→ System variables on windows (or)
→ `export GOPATH=<your path to GOCODE>` on unix based systems*

### Data Types:

These are the following datatypes available in Go programming language:

1. String - string
2. Boolean - bool
3. Integer - int, int8, int16, int32, int64
4. Unsigned Integer - uint, uint8, uint16, uint32, uint64, uintptr
5. Float - float32, float64
6. Complex - complex64, complex128
7. Alias
    1. byte - for uint8
    2. rune - for int32