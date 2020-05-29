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