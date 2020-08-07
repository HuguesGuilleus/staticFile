# staticFile
Minify files and save them into a golang source code.

## Installation
```bash
$ go get -u github.com/HuguesGuilleus/staticFile
```
Or you can use release binary.

## Goal
The goal is to make a functions that return the content of a file.
```bash
staticFile front/app.js
```
Create a file assets.go with a FrontApp function that return the minify cached front/app.js file. You can custom the name with `name=path` in place of just the path.

```bash
staticFile Js=front/app.js
```
The fonction is call Js and contain the file front/app.js

If the path is a directory, all files are imported.

## Flag
- `-dev`: enable dev mode
- `-out.file`: The output go source file. (default "assets.go")
- `-out.pkg`: The package of the source file. (default "main")
- `-out.trim`: Remove prefix in the name function.
