package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	delimiter string
	field     fieldValue
)

type fieldValue struct {
	value []int
}

func (f fieldValue) String() string {
	var s string

	for i, v := range f.value {
		s += fmt.Sprintf("field:%d=%d\t", i+1, v)
	}

	return s
}

func (f *fieldValue) Set(s string) error {
	var digits []string

	if strings.Contains(s, ",") {
		digits = strings.Split(s, ",")
	} else if strings.Contains(s, " ") {
		digits = strings.Split(s, " ")
	} else {
		digits = append(digits, s)
	}

	for _, d := range digits {
		i, err := strconv.Atoi(strings.TrimSpace(d))
		if err != nil {
			return err
		}

		f.value = append(f.value, i)
	}

	return nil
}

func init() {
	flag.Var(&field, "f", "the field to be cut from file, ex:f2")
	flag.StringVar(&delimiter, "d", "\t", "the delimiter to be used to cut")
}

func main() {
	flag.Parse()
	fileName := flag.Arg(0)

	f, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := bytes.Split(f, []byte("\n"))

	for i, l := range lines {
		fields := bytes.Split(l, []byte(delimiter))

		var out string
		for _, f := range field.value {
			if len(fields) < f {
				fmt.Printf("field %d missing on line: %d, (%d)\n", f, i, len(fields))
				continue
			}

			out += fmt.Sprintf("%s\t", fields[f-1])
		}

		fmt.Println(out)
	}
}
