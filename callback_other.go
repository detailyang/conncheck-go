// +build !windows

package conncheck

import "syscall"

func SyscallRead(fd uintptr, buf []byte) (int, error) {
	return syscall.Read(int(fd), buf[:])
}
