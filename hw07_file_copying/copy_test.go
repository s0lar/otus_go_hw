package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	// Place your code here.

	t.Run("Errors", func(t *testing.T) {
		tests := []struct {
			error  string
			from   string
			to     string
			offset int64
			limit  int64
		}{
			{error: "unsupported file", from: "", to: ""},
			{error: "unsupported file", from: "a", to: ""},
			{error: "same file", from: "a", to: "a"},
			{error: "open a: no such file or directory", from: "a", to: "aa"},
			{error: "offset exceeds file size", from: "testdata/input.txt", to: "/tmp/aa.txt", offset: 10000},
		}

		for i, test := range tests {
			err := Copy(test.from, test.to, test.offset, test.limit)
			assert.EqualError(t, err, test.error, i)
		}
	})

	t.Run("Success", func(t *testing.T) {
		tests := []struct {
			from   string
			to     string
			offset int64
			limit  int64
		}{
			{from: "testdata/input.txt", to: "/tmp/aa.txt", offset: 0, limit: 10000},
		}

		for _, test := range tests {
			Copy(test.from, test.to, test.offset, test.limit)
		}
	})
}
