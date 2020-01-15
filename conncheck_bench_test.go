package conncheck

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkCheck(b *testing.B) {
	ln, err := net.Listen("tcp4", "127.0.0.1:0")
	require.Nil(b, err)

	doneCh := make(chan struct{})
	go func() {
		_, _ = ln.Accept()
		<-doneCh
	}()

	conn, err := net.Dial("tcp4", ln.Addr().String())
	require.Equal(b, nil, err)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Check(conn)
	}

	doneCh <- struct{}{}
}
