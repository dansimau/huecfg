package cmd

import (
	"strings"
)

// headerTransform allow us to rename headers from the data to different
// names when displaying output or taking input from the user.
type headerTransform struct {
	output map[string]string
	input  map[string]string
}

func newHeaderTransform(output map[string]string) headerTransform {
	t := headerTransform{}
	t.output = make(map[string]string)
	t.input = make(map[string]string)

	for k, v := range output {
		t.output[strings.ToLower(k)] = v
	}

	// Input transformations are the reverse of the output ones
	for k, v := range output {
		t.input[strings.ToLower(v)] = k
	}

	return t
}

func (t headerTransform) TransformOutput(header string) string {
	v, ok := t.output[strings.ToLower(header)]
	if !ok {
		return header
	}

	return v
}

func (t headerTransform) TransformInput(header string) string {
	v, ok := t.input[strings.ToLower(header)]
	if !ok {
		return header
	}

	return v
}

type fieldTransformFunc = func(field string) (transformed string)

type fieldTransform struct {
	output map[string]fieldTransformFunc
}

func newFieldTransform(output map[string]fieldTransformFunc) fieldTransform {
	t := fieldTransform{}
	t.output = make(map[string]fieldTransformFunc)

	for k, v := range output {
		t.output[strings.ToLower(k)] = v
	}

	return t
}

func (t fieldTransform) TransformOutput(fieldName, fieldValue string) string {
	transFunc, ok := t.output[strings.ToLower(fieldName)]

	if !ok {
		return fieldValue
	}

	return transFunc(fieldValue)
}
