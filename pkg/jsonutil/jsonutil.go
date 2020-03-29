package jsonutil

import (
	"encoding/json"
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
	return err
}

// PrintBytes takes JSON bytes and prints the formatted JSON as a string to
// stdout.
func PrintBytes(data []byte) error {
	return FprintBytes(os.Stdout, data)
}
