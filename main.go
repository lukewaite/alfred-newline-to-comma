// Copyright (c) 2019 Luke Waite <lwaite@gmail.com>
// MIT Licence - http://opensource.org/licenses/MIT

package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := os.Args[1]
	arr := strings.Split(data, "\n")

	var buffer bytes.Buffer

	for _, s := range arr {
		if _, err := strconv.Atoi(s); err == nil {
			buffer.WriteString(s+",")
		} else {
			buffer.WriteString("'" + s + "',")
		}
	}

	out := buffer.String()
	fmt.Print(strings.TrimSuffix(out, ","))
}
