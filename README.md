# gotree
A small exercise in Go. 
Prints contents of a directory recursively in the same manner as `tree` does.

# Installation
Clone the repo:
```bash
git clone https://github.com/JakobSachs/gotree
cd gotree
```
Build:
```
go get JakobSachs/gotree
go build
```

and run:
```
./gotree 
```

_(Alternatively you can also just run directly via `go run gotree.go`)_

# Options

Sofar only the only option is the depth of the file system tree to be printed:
```
./gotree -level 5 ~
```
