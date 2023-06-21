# G

Run the following Git command to download the `git` package:

```sh
$ git clone https://github.com/filmangullo/Golang-Project-Structure.git MyProjectGolang
```

﻿# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)
*Gopher image by [Renee French][rf], licensed under [Creative Commons 3.0 Attributions license][cc3-by].*

Our canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Unless otherwise noted, the Go source files are distributed under the
BSD-style license found in the LICENSE file.

### Download and Install

#### Getting Gin

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/gin-gonic/gin"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `gin` package:

```sh
$ go get -u github.com/gin-gonic/gin
```

#### Getting Gorm

* GORM Guides [https://gorm.io](https://gorm.io)
* Gen Guides [https://gorm.io/gen/index.html](https://gorm.io/gen/index.html)

Intall gorm documentation

```sh
$ go get -u gorm.io/gorm
```

With [Connecting to a Database](https://gorm.io/docs/connecting_to_the_database.html) GORM officially supports the databases MySQL, PostgreSQL, SQLite, SQL Server, and TiDB

Gorm install with :
* Mysql database 

```sh
$ go get -u gorm.io/driver/mysql
```

* SQLite database 

```sh
$ go get -u gorm.io/driver/sqlite
```

#### Binary Distributions

Official binary distributions are available at https://go.dev/dl/.

After downloading a binary release, visit https://go.dev/doc/install
for installation instructions.

#### Install From Source

If a binary distribution is not available for your combination of
operating system and architecture, visit
https://go.dev/doc/install/source
for source installation instructions.

### Note

