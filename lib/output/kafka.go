// Copyright (c) 2014 Ashley Jeffs
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

package output

import (
	"github.com/Jeffail/benthos/lib/output/writer"
	"github.com/Jeffail/benthos/lib/types"
	"github.com/Jeffail/benthos/lib/util/service/log"
	"github.com/Jeffail/benthos/lib/util/service/metrics"
)

//------------------------------------------------------------------------------

func init() {
	Constructors["kafka"] = TypeSpec{
		constructor: NewKafka,
		description: `
The kafka output type writes messages to a kafka broker, these messages are
acknowledged, which is propagated back to the input. The config field
'ack_replicas' determines whether we wait for acknowledgement from all replicas
or just a single broker.

If the field 'key' is not empty then each message will be given its contents as
a key. This field can be dynamically set using function interpolations described
[here](../config_interpolation.md#functions).

By default the paritioner will select partitions based on a hash of the key
value. If the key is empty then a partition is chosen at random. You can
alternatively force the partitioner to round-robin partitions with the field
'round_robin_partitions'.`,
	}
}

//------------------------------------------------------------------------------

// NewKafka creates a new Kafka output type.
func NewKafka(conf Config, mgr types.Manager, log log.Modular, stats metrics.Type) (Type, error) {
	k, err := writer.NewKafka(conf.Kafka, log, stats)
	if err != nil {
		return nil, err
	}
	return NewWriter(
		"kafka", k, log, stats,
	)
}

//------------------------------------------------------------------------------
