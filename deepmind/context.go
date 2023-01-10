package deepmind

import (
	"fmt"
	"github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/data/bookkeeping"
	"github.com/algorand/go-algorand/protocol"
	"go.uber.org/atomic"
	"os"
	"time"
)

// NoOpContext can be used when no recording should happen for a given code path
var NoOpContext *Context

var syncContext *Context = NewContext(&DelegateToWriterPrinter{writer: os.Stdout})

// MaybeSyncContext returns the global deepmind.Context
func MaybeSyncContext(local config.Local) *Context {
	if !local.DeepMindEnabled {
		return nil
	}
	return syncContext
}

// SyncContext returns the sync context without any checking if deep mind is enabled or not. Use
// it only for specific cases and ensure you only use it when it's strictly correct to do so as this
// will print stdout lines.
func SyncContext() *Context {
	return syncContext
}

// Context is a block level data container used throughout deep mind instrumentation to
// keep active state about current instrumentation. This contains method to deal with
// block, transaction and call metadata required for proper functioning of Deep Mind
// code.
type Context struct {
	printer   Printer
	seenBlock *atomic.Bool
}

func NewContext(printer Printer) *Context {
	ctx := &Context{
		printer:   printer,
		seenBlock: atomic.NewBool(false),
	}

	return ctx
}

func (ctx *Context) Enabled() bool {
	return ctx != nil
}

func (ctx *Context) DeepMindLog() []byte {
	if ctx == nil {
		return nil
	}

	if v, ok := ctx.printer.(*ToBufferPrinter); ok {
		return v.buffer.Bytes()
	}

	return nil
}

// Block methods

func (ctx *Context) LogBlock(block *bookkeeping.Block) {
	if block.BlockHeader.Round == 1 {
		// Sleep 5 seconds if we're on the first block so that firehose has time to start.
		time.Sleep(time.Second * 5)
	}
	protocol.JSONHandle.Indent = 0
	ctx.printer.Print(fmt.Sprintf("BLOCK %v", string(protocol.EncodeJSON(block.BlockHeader))))
}
