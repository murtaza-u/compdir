# 🌳 Go Bonzai Directory Completer

[![GoDoc](https://godoc.org/github.com/murtaza-u/compdir?status.svg)](https://godoc.org/github.com/murtaza-u/compdir)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

Like [compfile](https://github.com/rwxrob/compfile), but only completes
directories.

## Style Guidelines

* Everything through `go fmt` or equiv, no exceptions
* In Vim `set textwidth=72` (not 80 to line numbers fit)
* Use `/* */` for package documentation comment, `//` elsewhere
* Smallest possible names for given scope while still clear
* Favor additional packages (possibly in `internal`) over long names
* Package globals that will be used a lot can be single capital
* Must be good reason to use more than 4 character pkg name
* Avoid unnecessary comments
* Use "deciduous tree" emoji 🌳 to mark Bonzai stuff

## Legal

"Bonzai" and "bonzai" are legal trademarks of Robert S. Muhlestein but
can be used freely to refer to the Bonzai™ project
<https://github.com/rwxrob/bonzai> without limitation. To avoid
potential developer confusion, intentionally using these trademarks to
refer to other projects --- free or proprietary --- is prohibited.
