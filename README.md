# readline

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.md)
[![Build Status](https://travis-ci.org/chzyer/readline.svg?branch=master)](https://travis-ci.org/chzyer/readline)
[![GoDoc](https://godoc.org/github.com/chzyer/readline?status.svg)](https://godoc.org/github.com/chzyer/readline)  

Readline is a pure go implementation for GNU-Readline kind library.

**WHY:**
Readline will support most of features which GNU Readline is supported, and provide a pure go environment and a MIT license.

# Demo

![demo](https://raw.githubusercontent.com/chzyer/readline/master/example/demo.gif)

You can read the source code in [example/main.go](https://github.com/chzyer/readline/blob/master/example/main.go).

# Usage

* Simplest example

```go
import "github.com/chzyer/readline"

rl, err := readline.New("> ")
if err != nil {
	panic(err)
}
defer rl.Close()

for {
	line, err := rl.Readline()
	if err != nil { // io.EOF
		break
	}
	println(line)
}
```

* Example with durable history

```go
rl, err := readline.NewEx(&readline.Config{
	Prompt: "> ",
	HistoryFile: "/tmp/readline.tmp",
})
if err != nil {
	panic(err)
}
defer rl.Close()

for {
	line, err := rl.Readline()
	if err != nil { // io.EOF
		break
	}
	println(line)
}
```

* Example with auto refresh

```go
import (
	"log"
	"github.com/chzyer/readline"
)

rl, err := readline.New("> ")
if err != nil {
	panic(err)
}
defer rl.Close()
log.SetOutput(l.Stderr()) // let "log" write to l.Stderr instead of os.Stderr

go func() {
	for _ = range time.Tick(time.Second) {
		log.Println("hello")
	}
}()

for {
	line, err := rl.Readline()
	if err != nil { // io.EOF
		break
	}
	println(line)
}
```


# Shortcut

`Meta`+`B` means press `Esc` and `n` separately.  
Users can change that in terminal simulator(i.e. iTerm2) to `Alt`+`B`

| Shortcut           | Comment                           | Support |
|--------------------|-----------------------------------|---------|
| `Ctrl`+`A`         | Beginning of line                 | Yes     |
| `Ctrl`+`B` / `←`   | Backward one character            | Yes     |
| `Meta`+`B`         | Backward one word                 | Yes     |
| `Ctrl`+`C`         | Send io.EOF                       | Yes     |
| `Ctrl`+`D`         | Delete one character              | Yes     |
| `Meta`+`D`         | Delete one word                   | Yes     |
| `Ctrl`+`E`         | End of line                       | Yes     |
| `Ctrl`+`F` / `→`   | Forward one character             | Yes     |
| `Meta`+`F`         | Forward one word                  | Yes     |
| `Ctrl`+`G`         | Cancel                            | Yes     |
| `Ctrl`+`H`         | Delete previous character         | Yes     |
| `Ctrl`+`I` / `Tab` | Command line completion           | NoYet   |
| `Ctrl`+`J`         | Line feed                         | Yes     |
| `Ctrl`+`K`         | Cut text to the end of line       | Yes     |
| `Ctrl`+`L`         | Clean screen                      | NoYet   |
| `Ctrl`+`M`         | Same as Enter key                 | Yes     |
| `Ctrl`+`N` / `↓`   | Next line (in history)            | Yes     |
| `Ctrl`+`P` / `↑`   | Prev line (in history)            | Yes     |
| `Ctrl`+`R`         | Search backwards in history       | Yes     |
| `Ctrl`+`S`         | Search forwards in history        | Yes     |
| `Ctrl`+`T`         | Transpose characters              | Yes     |
| `Meta`+`T`         | Transpose words                   | NoYet   |
| `Ctrl`+`U`         | Cut text to the beginning of line | NoYet   |
| `Ctrl`+`W`         | Cut previous word                 | Yes     |
| `Backspace`        | Delete previous character         | Yes     |
| `Meta`+`Backspace` | Cut previous word                 | Yes     |
| `Enter`            | Line feed                         | Yes     |


# Feedback

If you have any question, please submit an GitHub Issues and any pull request is welcomed :)
