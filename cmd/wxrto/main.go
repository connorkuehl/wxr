package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/connorkuehl/wxr"
)

func main() {
	buf := bytes.Buffer{}
	_, err := io.Copy(&buf, os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	var wxr wxr.RSS

	err = xml.Unmarshal(buf.Bytes(), &wxr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", wxr)
}
