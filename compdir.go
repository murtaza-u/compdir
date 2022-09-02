/*
Package compdir is a completion driver for Bonzai command trees and
fulfills the bonzai.Completer package interface. It's like compfile, but
only completes directories.
*/
package compdir

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/rwxrob/bonzai"
	"github.com/rwxrob/fn/filt"
)

// New returns a pointer to a struct that fulfills bonzai.Completer.
// This can be called from within Z.Cmd assignment:
//
// var Cmd = &Z.Cmd{
// 	Name: `some`,
// 	Comp: compdir.New(),
// }

func New() *comp { return new(comp) }

// comp fulfills the bonzai.Completer interface.
type comp struct{}

// Complete returns all directory names within a parent directory
// matching a prefix. If nothing is passed assumes the parent to be the
// current working directory.
func (comp) Complete(_ bonzai.Command, args ...string) []string {
	if len(args) == 0 {
		return recurse(".", "")
	}

	arg := args[len(args)-1]
	arg = expandHome(arg)

	path, sub := filepath.Split(arg)
	if path == "" {
		path = "."
	}

	return recurse(path, sub)
}

var seperator string

func init() {
	// "/" on unix, "\" on windows
	seperator = string(filepath.Separator)
}

func dirsWithSlash(parent string) (dirs []string) {
	items, err := os.ReadDir(parent)
	if err != nil {
		return
	}

	for _, i := range items {
		if !i.IsDir() {
			continue
		}

		name := filepath.Join(parent, i.Name()) + seperator
		dirs = append(dirs, name)
	}

	return
}

// a parent containing only files and no directories is also considered
// an empty directory
func isDirEmpty(dir string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return true
	}

	for _, e := range entries {
		if e.IsDir() {
			return false
		}
	}

	return true
}

func expandHome(path string) string {
	if !strings.HasPrefix(path, "~") {
		return path
	}

	path = strings.TrimPrefix(path, "~")
	home, err := os.UserHomeDir()
	if err != nil {
		// panicking or logging errors in a completion driver will ruin
		// UX
		return path
	}

	return home + path
}

func recurse(dir, sub string) []string {
	entries := dirsWithSlash(dir)

	// complete relative paths
	if sub == "." || sub == ".." {
		sep := seperator
		entries = append(entries, sub+sep)
	}

	list := filt.BaseHasPrefix(entries, sub)
	if len(list) == 1 && !isDirEmpty(list[0]) {
		return recurse(list[0], "")
	}

	return list
}
