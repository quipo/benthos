// Copyright (c) 2018 Ashley Jeffs
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package processor

import (
	"bytes"
	"compress/gzip"
	"os"
	"reflect"
	"testing"

	"github.com/Jeffail/benthos/lib/types"
	"github.com/Jeffail/benthos/lib/util/service/log"
	"github.com/Jeffail/benthos/lib/util/service/metrics"
)

func TestDecompressBadAlgo(t *testing.T) {
	conf := NewConfig()
	conf.Decompress.Algorithm = "does not exist"

	testLog := log.NewLogger(os.Stdout, log.LoggerConfig{LogLevel: "NONE"})

	_, err := NewDecompress(conf, nil, testLog, metrics.DudType{})
	if err == nil {
		t.Error("Expected error from bad algo")
	}
}

func TestDecompressGZIP(t *testing.T) {
	conf := NewConfig()
	conf.Decompress.Algorithm = "gzip"

	testLog := log.NewLogger(os.Stdout, log.LoggerConfig{LogLevel: "NONE"})

	input := [][]byte{
		[]byte("hello world first part"),
		[]byte("hello world second part"),
		[]byte("third part"),
		[]byte("fourth"),
		[]byte("5"),
	}

	exp := [][]byte{}

	for i := range input {
		exp = append(exp, input[i])

		var buf bytes.Buffer

		zw := gzip.NewWriter(&buf)
		zw.Write(input[i])
		zw.Close()

		input[i] = buf.Bytes()
	}

	if reflect.DeepEqual(input, exp) {
		t.Fatal("Input and exp output are the same")
	}

	proc, err := NewDecompress(conf, nil, testLog, metrics.DudType{})
	if err != nil {
		t.Fatal(err)
	}

	msgs, res := proc.ProcessMessage(types.NewMessage(input))
	if len(msgs) != 1 {
		t.Error("Decompress failed")
	} else if res != nil {
		t.Errorf("Expected nil response: %v", res)
	}
	if act := msgs[0].GetAll(); !reflect.DeepEqual(exp, act) {
		t.Errorf("Unexpected output: %s != %s", act, exp)
	}
}

func TestDecompressIndexBounds(t *testing.T) {
	conf := NewConfig()

	testLog := log.NewLogger(os.Stdout, log.LoggerConfig{LogLevel: "NONE"})

	input := [][]byte{
		[]byte("0"),
		[]byte("1"),
		[]byte("2"),
		[]byte("3"),
		[]byte("4"),
	}

	for i := range input {
		var buf bytes.Buffer

		zw := gzip.NewWriter(&buf)
		zw.Write(input[i])
		zw.Close()

		input[i] = buf.Bytes()
	}

	type result struct {
		index int
		value string
	}

	tests := map[int]result{
		-5: {
			index: 0,
			value: "0",
		},
		-4: {
			index: 1,
			value: "1",
		},
		-3: {
			index: 2,
			value: "2",
		},
		-2: {
			index: 3,
			value: "3",
		},
		-1: {
			index: 4,
			value: "4",
		},
		0: {
			index: 0,
			value: "0",
		},
		1: {
			index: 1,
			value: "1",
		},
		2: {
			index: 2,
			value: "2",
		},
		3: {
			index: 3,
			value: "3",
		},
		4: {
			index: 4,
			value: "4",
		},
	}

	for i, result := range tests {
		conf.Decompress.Parts = []int{i}
		proc, err := NewDecompress(conf, nil, testLog, metrics.DudType{})
		if err != nil {
			t.Fatal(err)
		}

		msgs, res := proc.ProcessMessage(types.NewMessage(input))
		if len(msgs) != 1 {
			t.Errorf("Decompress failed on index: %v", i)
		} else if res != nil {
			t.Errorf("Expected nil response: %v", res)
		}
		if exp, act := result.value, string(msgs[0].GetAll()[result.index]); exp != act {
			t.Errorf("Unexpected output for index %v: %v != %v", i, act, exp)
		}
	}
}

func TestDecompressEmpty(t *testing.T) {
	conf := NewConfig()
	conf.Decompress.Parts = []int{0, 1}

	testLog := log.NewLogger(os.Stdout, log.LoggerConfig{LogLevel: "NONE"})
	proc, err := NewDecompress(conf, nil, testLog, metrics.DudType{})
	if err != nil {
		t.Error(err)
		return
	}

	msgs, _ := proc.ProcessMessage(types.NewMessage([][]byte{}))
	if len(msgs) > 0 {
		t.Error("Expected failure with zero part message")
	}

	msgs, _ = proc.ProcessMessage(types.NewMessage(
		[][]byte{[]byte("first"), []byte("second")},
	))
	if len(msgs) > 0 {
		t.Error("Expected failure with bad data")
	}
}
