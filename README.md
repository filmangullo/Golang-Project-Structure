# G

Run the following Git command to download the `git` package:

```sh
$ git clone https://github.com/filmangullo/Golang-Project-Structure.git MyProjectGolang
```

# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)
_Gopher image by [Renee French][rf], licensed under [Creative Commons 3.0 Attributions license][cc3-by]._

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

#### Getting GoDotEnv

A Go (golang) port of the Ruby dotenv project (which loads env vars from a .env file).

- Env Guides [https://github.com/joho/godotenv](https://github.com/joho/godotenv)

Intall GoDotEnv documentation

```sh
$ go get github.com/joho/godotenv
```

#### Getting Gorm

- GORM Guides [https://gorm.io](https://gorm.io)
- Gen Guides [https://gorm.io/gen/index.html](https://gorm.io/gen/index.html)

Intall gorm documentation

```sh
$ go get -u gorm.io/gorm
```

With [Connecting to a Database](https://gorm.io/docs/connecting_to_the_database.html) GORM officially supports the databases MySQL, PostgreSQL, SQLite, SQL Server, and TiDB

Gorm install with :

- Mysql database

```sh
$ go get -u gorm.io/driver/mysql
```

- PostgreSQL database

```sh
$ go get gorm.io/driver/postgres
```

- SQLite database

```sh
$ go get -u gorm.io/driver/sqlite
```

#### Getting Air - Live reload for Go apps

Air is yet another live-reloading command line utility for developing Go applications. Run air in your project root directory, leave it alone, and focus on your code.

- Air Guides [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)

Intall Air documentation

```sh
$ go install github.com/cosmtrek/air@latest
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

### Run

After installing "air" you don't need to run "$ go run ." but just "air"

For less typing, you could add `alias air='~/.air'` to your `.bashrc` or `.zshrc`.

First enter into your project

```bash
cd /path/to/your_project
```

The simplest usage is run

```bash
# firstly find `.air.toml` in current directory, if not found, use defaults
air -c .air.toml
```

You can initialize the `.air.toml` configuration file to the current directory with the default settings running the following command.

```bash
air init
```

After this, you can just run the `air` command without additional arguments and it will use the `.air.toml` file for configuration.

```bash
air
```

For modifying the configuration refer to the [air_example.toml](air_example.toml) file.
