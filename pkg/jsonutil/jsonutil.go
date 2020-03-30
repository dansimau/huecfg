package jsonutil

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// FprintBytes writes formatted JSON to w. It returns an error if JSON parsing
// or writing to w fails.
func FprintBytes(w io.Writer, data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	output, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	_, err = w.Write(output)
	_, err = fmt.Fprintln(w, "")
	return err
}

// PrintBytes takes JSON bytes and prints the formatted JSON as a string to
// stdout.
func PrintBytes(data []byte) error {
	return FprintBytes(os.Stdout, data)
}
