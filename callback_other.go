// +build !windows

package conncheck

import "syscall"

func callback(ctx *context) func(fd uintptr) bool {
	return func(fd uintptr) bool {
		var buff [1]byte
		ctx.n, ctx.err = syscall.Read(int(fd), buff[:])
		return true
	}
}
