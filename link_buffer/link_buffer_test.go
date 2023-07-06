package link_buffer

import (
	"fmt"
	"os"
	"syscall"
	"testing"
	"unsafe"
)

func TestLinkBuffer(t *testing.T) {
	buf1 := []byte("hello, ")
	buf2 := []byte("world!\n")

	ivs := []syscall.Iovec{
		{Base: &buf1[0], Len: uint64(len(buf1))},
		{Base: &buf2[0], Len: uint64(len(buf2))},
	}
	n, _, errNo := syscall.Syscall(syscall.SYS_WRITEV, os.Stdout.Fd(), uintptr(unsafe.Pointer(&ivs[0])), uintptr(len(ivs)))
	fmt.Println(n, errNo)
}
