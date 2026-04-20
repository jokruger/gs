# Installing

## Install the CLI with Go tooling

Requires Go 1.26 or later.

```sh
go install github.com/jokruger/gs/cmd/gs@latest
```

Make sure `$GOPATH/bin` (or `$HOME/go/bin`) is on your `PATH`.

## Using the CLI

### REPL

Launch the interactive REPL by running `gs` with no arguments:

```sh
gs
```

You will see the `>> ` prompt. Type any GS expression and press Enter to evaluate it. Press `Ctrl+D` to exit.

### Running a script

Pass a `.gs` source file as the first argument:

```sh
gs hello.gs
```

### CLI flags

- `-o <file>`: write compiled bytecode to a file
- `-version`: print the CLI version
- `-strict-assign`: require variables to already exist for plain `=` assignment

By default, GS uses smart `=` assignment at compile time (`x = expr` declares `x` in current scope if unresolved). Use `-strict-assign` to enforce explicit declaration before `=`.

```sh
gs -strict-assign hello.gs
```

### Hashbang / shebang scripts

Add a hashbang line as the first line of your script to make it directly executable:

```go
#!/usr/bin/env gs

fmt := import("fmt")
fmt.println("Hello GS!")
```

Then make the file executable and run it:

```sh
chmod +x hello.gs
./hello.gs
```

## Building from source

The project uses [just](https://github.com/casey/just) as its build tool. Install it before proceeding (`brew install just` on macOS, or see the just documentation for other platforms).

Clone the repository and enter the project directory:

```sh
git clone https://github.com/jokruger/gs.git
cd gs
```

### Common recipes

| Command | Description |
|---------|-------------|
| `just build` | Generate sources and compile the `gs` binary into `./build/gs` |
| `just install` | Build and copy the binary to `$HOME/bin/` |
| `just test` | Run the full test suite |
| `just clean` | Remove build artefacts and profiling files |

Run `just` with no arguments to list all available recipes.
