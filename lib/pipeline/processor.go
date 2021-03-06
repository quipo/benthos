// Copyright (c) 2017 Ashley Jeffs
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

package pipeline

import (
	"sync/atomic"
	"time"

	"github.com/Jeffail/benthos/lib/processor"
	"github.com/Jeffail/benthos/lib/types"
	"github.com/Jeffail/benthos/lib/util/service/log"
	"github.com/Jeffail/benthos/lib/util/service/metrics"
)

//------------------------------------------------------------------------------

// Processor is a pipeline that supports both Consumer and Producer interfaces.
// The processor will read from a source, perform some processing, and then
// either propagate a new message or drop it.
type Processor struct {
	running int32

	log   log.Modular
	stats metrics.Type

	msgProcessors []processor.Type

	messagesOut chan types.Transaction
	responsesIn chan types.Response

	messagesIn <-chan types.Transaction

	closeChan chan struct{}
	closed    chan struct{}
}

// NewProcessor returns a new message processing pipeline.
func NewProcessor(
	log log.Modular,
	stats metrics.Type,
	msgProcessors ...processor.Type,
) *Processor {
	return &Processor{
		running:       1,
		msgProcessors: msgProcessors,
		log:           log.NewModule(".pipeline.processor"),
		stats:         stats,
		messagesOut:   make(chan types.Transaction),
		responsesIn:   make(chan types.Response),
		closeChan:     make(chan struct{}),
		closed:        make(chan struct{}),
	}
}

//------------------------------------------------------------------------------

// loop is the processing loop of this pipeline.
func (p *Processor) loop() {
	defer func() {
		atomic.StoreInt32(&p.running, 0)

		close(p.messagesOut)
		close(p.closed)
	}()

	var open bool
	for atomic.LoadInt32(&p.running) == 1 {
		var tran types.Transaction
		select {
		case tran, open = <-p.messagesIn:
			if !open {
				return
			}
		case <-p.closeChan:
			return
		}
		p.stats.Incr("pipeline.processor.count", 1)

		resultMsgs := []types.Message{tran.Payload}
		var resultRes types.Response
		for i := 0; len(resultMsgs) > 0 && i < len(p.msgProcessors); i++ {
			var nextResultMsgs []types.Message
			for _, m := range resultMsgs {
				var rMsgs []types.Message
				rMsgs, resultRes = p.msgProcessors[i].ProcessMessage(m)
				nextResultMsgs = append(nextResultMsgs, rMsgs...)
			}
			resultMsgs = nextResultMsgs
		}

		if len(resultMsgs) == 0 {
			p.stats.Incr("pipeline.processor.dropped", 1)
			select {
			case tran.ResponseChan <- resultRes:
			case <-p.closeChan:
				return
			}
			continue
		}

		var res types.Response

		remainingResponses := len(resultMsgs)
		currentMsg := 0

	responsesLoop:
		for remainingResponses > 0 {
			var tsOut chan<- types.Transaction
			tsIndex := 0
			if currentMsg < len(resultMsgs) {
				tsOut = p.messagesOut
				tsIndex = currentMsg
			}
			select {
			case tsOut <- types.NewTransaction(resultMsgs[tsIndex], p.responsesIn):
				currentMsg++
			case res, open = <-p.responsesIn:
				if !open {
					return
				}
				if res.Error() == nil {
					p.stats.Incr("pipeline.processor.send.success", 1)
					remainingResponses--
				} else {
					p.stats.Incr("pipeline.processor.send.error", 1)
					break responsesLoop
				}
			case <-p.closeChan:
				return
			}
		}

		select {
		case tran.ResponseChan <- res:
		case <-p.closeChan:
			return
		}
	}
}

//------------------------------------------------------------------------------

// StartReceiving assigns a messages channel for the pipeline to read.
func (p *Processor) StartReceiving(msgs <-chan types.Transaction) error {
	if p.messagesIn != nil {
		return types.ErrAlreadyStarted
	}
	p.messagesIn = msgs
	go p.loop()
	return nil
}

// TransactionChan returns the channel used for consuming messages from this
// pipeline.
func (p *Processor) TransactionChan() <-chan types.Transaction {
	return p.messagesOut
}

// CloseAsync shuts down the pipeline and stops processing messages.
func (p *Processor) CloseAsync() {
	if atomic.CompareAndSwapInt32(&p.running, 1, 0) {
		close(p.closeChan)
	}
}

// WaitForClose blocks until the StackBuffer output has closed down.
func (p *Processor) WaitForClose(timeout time.Duration) error {
	select {
	case <-p.closed:
	case <-time.After(timeout):
		return types.ErrTimeout
	}
	return nil
}

//------------------------------------------------------------------------------
