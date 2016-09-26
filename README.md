# gorep

Simple app to replace imported package name in golang project. Useful to manage imported package name of a project which deployed into multiple application on one server.

Please continue reading, you'll understand.

## Story

This package is only useful when:

 1. You are not using [golang vendor](https://blog.gopheracademy.com/advent-2015/vendor-folder/)
 2. Your project is deployed as multiple app in one server

Say that you have a project which the package path is `github.com/novalagung/project`. You deployed this project on server as 3 different app (one for testing, one for dev, one for production).

```bash
+ $GOPATH
   + src
      + github.com
         + novalagung
            + project
                - main.go
                - file_a.go
                - file_b.go
                - file_c.go
                - other files ...
            + project-dev
            + project-test
   + bin
   + pkg
```

When you want to build the `project-dev` and `project-test`, you have to change imported package names in all files inside each folder.

```go
import "github.com/novalagung/project/model"
import "github.com/novalagung/project/controller"
import "github.com/novalagung/project/conf"
import "github.com/novalagung/project/view"
// ... and so oon, in all files
```

need to be replaced into this for the `project-dev`, as well as `project-test`

```go
import "github.com/novalagung/project-dev/model"
import "github.com/novalagung/project-dev/controller"
import "github.com/novalagung/project-dev/conf"
import "github.com/novalagung/project-dev/view"
```

It'll waste your time so much. So this simple library could be your life-saver.

## Installation

```
go get github.com/novalagung/gorep
```

Make sure `$GOPATH/bin` is added to your `$PATH` variable.

## Usage

Go to your project folder, then run this

```bash
cd $GOPATH/src/github.com/novalagung/project-dev
gorep -from="github.com/novalagung/project" -to="github.com/novalagung/project-dev"

# or $GOPATH/bin/gorep path="." -from="eaciit/x10-dev" -to="eaciit/x10"
```

Or use flag `-path` to specify the project location

```bash
gorep -path="/User/novalagung/goapp/src/github.com/novalagung/project-dev" -from="github.com/novalagung/project" -to="github.com/novalagung/project-dev"
```

flag `-from` is the package you want to replace `-to` is the new name. All files which contains `github.com/novalagung/project` inside those folder will be replaced with `github.com/novalagung/project-dev`.

## Author

Noval Agung Prayogo

## License

MIT