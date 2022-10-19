package main

import (
	"errors"
	"fmt"

	"github.com/piotrpersona/per"
)

func ReadFile() (err error) {
	err = errors.New("OH NO")
	return
}

func main() {
	err := ReadFile()
	err = per.Error(err, "cannot read file")
	fmt.Println(err)

	err = per.Errorf(ReadFile(), "cannot read file: %s", "file.txt")
	fmt.Println(err)

	per.Format = "Custom MSG: %s, ERR: %s"
	err = per.Errorf(ReadFile(), "cannot read file: %s", "file.txt")
	fmt.Println(err)
}
