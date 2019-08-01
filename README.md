# goenvdir 

golang envdir replica

## Build

Use `go build` or `go install` to build

## Usage

`goenvdir [-i] <path to directory> <program to run>`

Use `i` flag to run program with env variables ONLY from directory.

## Examples

With env override

`[titkovd@localhost]$ goenvdir -i ./example env
BAR=cowsaregreat
FOO=123
MOO=FOOD`

Without env override

`[titkovd@localhost]$ goenvdir ./example env
BAR=cowsaregreat
FOO=123
MOO=FOOD
XDG_VTNR=1
TERM_PROGRAM=vscode
<...>
DISPLAY=:0
COLORTERM=truecolor
_=/usr/local/go/bin/go`
