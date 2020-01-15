package conncheck

import (
	"errors"
	"io"
	"net"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

var ErrUnexpectedReadOneByte = errors.New("conn: unexpected read one byte")

type closure struct {
	F   uintptr
	n   *int
	err *error
}

type context struct {
	n   int
	err error
	fn  func(fd uintptr) bool
}

var ctxPool = sync.Pool{
	New: func() interface{} {
		return &context{}
	},
}

func acquireCtx() *context {
	ctx, ok := ctxPool.Get().(*context)
	if !ok {
		panic("failed to type casting")
	}

	if ctx.fn == nil {
		ctx.fn = callback(ctx)
	} else {
		// This is unsafe but can reduce the runtime.newobject
		//
		// if the go rutnime breaks the closure function spec, we will be broken
		ptr := *(*closure)(unsafe.Pointer(&ctx.fn))
		ptr.n = &ctx.n
		ptr.err = &ctx.err
	}

	return ctx
}

func releaseCtx(ctx *context) {
	ctx.n = 0
	ctx.err = nil
	ctxPool.Put(ctx)
}

func Check(c net.Conn) error {
	var (
		n   int
		err error
	)

	sconn, ok := c.(syscall.Conn)
	if !ok {
		return nil
	}

	// SyscallConn allocate the raw conn on heap
	rc, err := sconn.SyscallConn()
	if err != nil {
		return err
	}

	ctx := acquireCtx()

	if err = c.SetDeadline(time.Time{}); err != nil {
		return err
	}
	rerr := rc.Read(ctx.fn)

	n = ctx.n
	err = ctx.err
	releaseCtx(ctx)

	switch {
	case rerr != nil:
		return rerr
	case n == 0 && err == nil:
		return io.EOF
	case n > 0:
		return ErrUnexpectedReadOneByte
	case err == syscall.EAGAIN || err == syscall.EWOULDBLOCK:
		return nil
	default:
		return err
	}
}
