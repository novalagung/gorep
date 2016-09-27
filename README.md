# gorep

Simple tools to replace imported package name in golang project. Useful to manage a project which deployed into multiple application on one server.

Please continue reading, you'll understand what is the purpose of this library

## Story

This package is only useful when:

 1. You are not using [golang vendor](https://blog.gopheracademy.com/advent-2015/vendor-folder/)
 2. Your project is deployed as multiple app in one server

Say that you have a project which the package path is `github.com/novalagung/project`. This project deployed to server as 3 different environment (one for *testing*, one for *dev*, one for *production*).

<img width="476" alt="screen shot 2016-09-26 at 10 03 12 pm" src="https://cloud.githubusercontent.com/assets/982868/18839396/11ca5fdc-8435-11e6-83d1-7d7331e00268.png" alt="Gorep - Simple tools to replace imported package name in golang project" align="center">

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

It'll waste your time so much, espesially if there are tons of file inside. TONS!. This simple library could be your life-saver.

## Installation

```
go get github.com/novalagung/gorep
```

Make sure `$GOPATH/bin` is added to your `$PATH` variable.

## Usage

Go to your project folder, then run this

```bash
cd $GOPATH/src/github.com/novalagung/project-dev
gorep -from="github.com/novalagung/project" \ 
      -to="github.com/novalagung/project-dev"
```

Or use flag `-path` to specify the project location

```bash
gorep -path="/User/novalagung/goapp/src/github.com/novalagung/project-dev" \
      -from="github.com/novalagung/project" \
      -to="github.com/novalagung/project-dev"
```

All files which contains `github.com/novalagung/project` inside those choosen folder will be replaced with `github.com/novalagung/project-dev`.

 - Flag `-from` filled with package name you want to replace
 - Flag `-to` is the new name
 - Flag `-flag` (optional) is the project path you want to hack. By default it'll be current active directory 

## Author

Noval Agung Prayogo

## License

MIT
