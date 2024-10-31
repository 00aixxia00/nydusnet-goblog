# GO

### Basics
##### How do we run the code
More information: see Go CLI
- go build: compiles
- go run: compiles and execute, eg `$ go run .`
- go fmt: format code
- go install & go get: handle dependencys
- go test: runs any test

##### package main
Package == Project == Workspace

Files in the same package do not have to be imported into each other! 

Running `go run main.go state.go` on the following files, will work out fine, if they are in the same folder. 🙂

main.go
```go
package main
 
func main() {
    printState()
}
```
state.go
```go
package main
 
import "fmt"
 
func printState() {
    fmt.Println("California")
}
```



- assign to folder it belongs to write `package main` on top of the file
- type of packages
    - executable: generates a file that we can run, compiled and then executed (main.exe) 
        - eg. package main (must have a func called main inside)
        - eg. package apple ...
    - reusable: like dependencies or packages - "helpers"
        - eg. package calculator, package uploader
    
    

##### import fmt
- access to all functionality to package X
- fmt=format - "print out"
-  [standard linbrary](https://pkg.go.dev/std) 
    - math
    - encoding
    - debug
    - crypto
    - io
    - fmt ...

```
I FU****CKN LIKE THIS PAGE
...
better ... it's gonna be your BF
```


##### func


##### variables
- explizit: `var card string = "Ace of Spades"` 
- letting go compiler decide: by `card := "Ace of Spades"`
    - use it once! the first time when assigning the variable 
    - fter that we can just type `card = "Five of Diamonds"`



### Lists
- **array**: fixed legth list of things
- **slice**: an array that can grow or shrink
- elements must be of the identical type
- both must be defined with a dataype

