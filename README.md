<h1 align="center">Cognilog</h1>
<br>
![home](https://github.com/cognifly/cognibot/blob/master/img/cognilog.png)

Package cognilog is a simple color logging package for Go. The colors supported are 
red, green, yellow, white, cyan and magenta

## Example

Let's have a look at the examples which demonstrates most
of the features found in this library.

Log glues Prefix E.G 2016/07/22 13:37:08 to to msg before printing to stdout.
```go
package main

import (

	"github.com/cognifly/cognilog"
)

func main() {
	msg := "Hello am cognilog"
	cognilog.Log("blue", msg)
}
```

LogINFO glues Prefix E.G 2016/07/22 13:37:08 and [infor] to msg before 
printing to stdout.
```go
package main

import (
	"net/url"
	"github.com/cognifly/cognilog"
)

func main() {
	msg := *url.URL
	cognilog.LogInfo("green", "200 OK", msg)
}
```
Fatal glues Prefix E.G 2016/07/22 13:37:08 to to msg before printing to stdout.
Exit with status 1.
```go
package main

import (
	"github.com/cognifly/cognilog"
)

func main() {
	msg := "very bad error"
	cognilog.Fatal("red", msg)
}
```
FatalINFO glues Prefix E.G 2016/07/22 13:37:08 and [infor] to msg before 
printing to stdout. Exit with status 1.
```go
package main

import (
	"github.com/cognifly/cognilog"
)

func main() {
	msg := "very very bad error"
	cognilog.FatalInfo("red", "panic", msg)
}
```

## Installation

To install, simply run in a terminal:

    go get github.com/cognifly/cognilog

## License

The [BSD-style license] found in the LICENSE file

Name *kampsy kampamba chanda*
Email *kampsycode@gmail.com*
Github *https://github.com/kampsy*
Social *google.com/+kampambachanda*
