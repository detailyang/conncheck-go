package conncheck

import "syscall"

func SyscallRead(fd uintptr, buf []byte) (int, error) {
	return syscall.Read(syscall.Handle(fd), buf[:])
}
