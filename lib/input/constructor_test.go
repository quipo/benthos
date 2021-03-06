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

package input

import (
	"reflect"
	"testing"

	"github.com/Jeffail/benthos/lib/processor"
)

func TestSanitise(t *testing.T) {
	exp := map[string]interface{}{
		"type": "amqp",
		"amqp": map[string]interface{}{
			"consumer_tag":   "benthos-consumer",
			"prefetch_count": float64(1),
			"prefetch_size":  float64(0),
			"url":            "amqp://guest:guest@localhost:5672/",
			"exchange":       "benthos-exchange",
			"exchange_type":  "direct",
			"queue":          "benthos-queue",
			"key":            "benthos-key",
		},
	}

	conf := NewConfig()
	conf.Type = "amqp"
	conf.Processors = nil

	act, err := SanitiseConfig(conf)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(act, exp) {
		t.Errorf("Wrong sanitised output: %v != %v", act, exp)
	}

	exp["processors"] = []interface{}{
		map[string]interface{}{
			"type": "combine",
			"combine": map[string]interface{}{
				"parts": float64(2),
			},
		},
		map[string]interface{}{
			"type": "archive",
			"archive": map[string]interface{}{
				"format": "binary",
				"path":   "nope",
			},
		},
	}

	proc := processor.NewConfig()
	proc.Type = "combine"
	conf.Processors = append(conf.Processors, proc)

	proc = processor.NewConfig()
	proc.Type = "archive"
	proc.Archive.Path = "nope"
	conf.Processors = append(conf.Processors, proc)

	act, err = SanitiseConfig(conf)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(act, exp) {
		t.Errorf("Wrong sanitised output: %v != %v", act, exp)
	}
}
