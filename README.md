# Orders by go rutines

### Requirements

- api runing in http://localhost:3333/

### Usage
```bash
Usage:
  -file string
    	Input file (default "./sources/pendingmin.txt")
  -routines int
    	Number of routines to use (default 10)

```
### Example of use

```
$ go run ./cli/main.go -file=./sources/pending.txt -routines 50                                
```