package main

import (
	"bufio"
	"io"
)

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
